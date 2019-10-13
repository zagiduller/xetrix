package services

import (
	"engine/lib/helper"
	"engine/lib/services/events"
	pb "engine/lib/structs"
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

type IOrderRepository interface {
	CreateOrder(in *pb.Order) (*pb.Order, error)
	GetOrder(id string) (*pb.Order, error)
	GetOrders(req *pb.Query_Order) ([]*pb.Order, error)
	UpdateOrderAvailable(o *pb.Order) (*pb.Order, error)

	CreateContract(c *pb.Contract) (*pb.Contract, error)
	GetContract(id string) (*pb.Contract, error)
	GetContracts(req *pb.Query_Contract) ([]*pb.Contract, error)

	CreateWithdrawal(wo *pb.WithdrawalOrder) (*pb.WithdrawalOrder, error)
	GetWithdrawal(wo *pb.Query_Withdrawal) ([]*pb.WithdrawalOrder, error)
	UpdateWithdrawal(wo *pb.WithdrawalOrder) (*pb.WithdrawalOrder, error)
}

type ServiceOrder struct {
	serviceCurrency       *ServiceCurrency
	serviceUser           *ServiceUser
	serviceAccount        *ServiceAccount
	serviceAccountBalance *ServiceAccountBalance
	serviceTransaction    *ServiceTransaction
	serviceCommission     *ServiceCommission
	repo                  IOrderRepository
	bus                   *events.Bus
}

func NewOrderService(repo IOrderRepository, srv_c *ServiceCurrency,
	srv_u *ServiceUser, srv_t *ServiceTransaction, srv_a *ServiceAccount, srv_ab *ServiceAccountBalance, srv_cm *ServiceCommission) *ServiceOrder {
	return &ServiceOrder{
		serviceCurrency:       srv_c,
		serviceUser:           srv_u,
		serviceTransaction:    srv_t,
		serviceAccount:        srv_a,
		serviceAccountBalance: srv_ab,
		serviceCommission:     srv_cm,
		repo:                  repo,
	}
}

func (s *ServiceOrder) AddEventBus(bus *events.Bus) {
	s.bus = bus
	bus.Subscribe(s,
		&pb.Event{Type: &pb.Event_TxConfirm{}},
	)
}

func (s *ServiceOrder) Update(event *pb.Event) {

	switch event.Type.(type) {
	// Подтверждена транзакция. Если она связано с заявкой на вывод, то нужно исполнить заявку PerformWithdrawal
	case *pb.Event_TxConfirm:
		tx := event.GetTxConfirm().Tx
		if tx.Reason.Status == pb.TxReason_WITHDRAW_TX {
			ctx := context.Background()
			go s.PerformWithdrawal(context.WithValue(ctx, "admin-id", "Event_machine"), &pb.Query_Withdrawal{
				Id:          tx.Reason.WithdrawalOrderId,
				RelatedTxId: tx.Id,
			})
		}
	}

}

func (s *ServiceOrder) Notify(event *pb.Event) {
	s.bus.NewEvent(event)
}

func (s *ServiceOrder) CreateOrder(ctx context.Context, q *pb.Query_Order) (*pb.Response_Order, error) {
	pid := ctx.Value("pid")
	if pid != nil {
		//Проверка инициатора
		p_resp, err := s.serviceUser.GetUser(ctx, &pb.Query_User{Id: pid.(string)})
		if p_resp == nil || err != nil {
			//if p_resp, err := s.serviceParticipant.GetParticipant(ctx, &pb.Query_Participant{ Id: q.OwnerId }); p_resp == nil || err != nil {
			return nil, fmt.Errorf("CreateOrder: undefined participant %s", pid.(string))
		}
		owner := p_resp.Object

		//Проверка счетов
		sa_resp, _ := s.serviceAccount.GetAccount(ctx, &pb.Query_Account{Address: q.SendingAddress})
		ra_resp, _ := s.serviceAccount.GetAccount(ctx, &pb.Query_Account{Address: q.ReceiveAddress})
		if sa_resp == nil || ra_resp == nil {
			return nil, fmt.Errorf("CreateOrder: account undefined ")
		}

		if sc_resp, _ := s.serviceCurrency.GetCurrency(ctx, &pb.Query_Currency{Id: sa_resp.Object.Currency.Id}); !sc_resp.Object.Active {
			return nil, fmt.Errorf("CreateOrder: currency %s is inactive", sc_resp.Object.Name)
		}

		if rc_resp, _ := s.serviceCurrency.GetCurrency(ctx, &pb.Query_Currency{Id: ra_resp.Object.Currency.Id}); !rc_resp.Object.Active {
			return nil, fmt.Errorf("CreateOrder: currency %s is inactive", rc_resp.Object.Name)
		}

		if sa_resp.Object.Status == pb.Account_INACTIVE || ra_resp.Object.Status == pb.Account_INACTIVE {
			return nil, fmt.Errorf("CreateOrder: cant create order with inactive account or inactive currency. Currency: %t Send acc inactive : %t, Receive acc inactive: %t",
				sa_resp.Object.Status == pb.Account_INACTIVE, ra_resp.Object.Status == pb.Account_INACTIVE)
		}

		if q.Amount <= 0 || q.Price <= 0 {
			return nil, fmt.Errorf("CreateOrder: not avaliable amount or price")
		}

		// Проверка баланса для внутренних кошельков
		if sa_resp.Object.Type == pb.AccountType_INTERNAL {
			if sa_blnc_resp, _ := s.serviceAccountBalance.GetBalance(ctx, &pb.Query_Account{Address: sa_resp.Object.Address}); sa_blnc_resp == nil || sa_blnc_resp.Object.Available < q.Amount {
				return nil, fmt.Errorf("CreateOrder: not enought balance")
			}
		}
		// Установка комиссии для ордера 2%
		// Вопрос: как задавать размер комиссии?

		_order := pb.Order{
			OwnerId:            owner.Id,
			SendingAddress:     q.SendingAddress,
			ReceiveAddress:     q.ReceiveAddress,
			SellCurrencySymbol: sa_resp.Object.Currency.Symbol,
			BuyCurrencySymbol:  ra_resp.Object.Currency.Symbol,
			Amount:             q.Amount,
			Price:              q.Price,
			Available:          q.Amount,
			FrontMetaData:      q.FrontMetaData,
		}

		if _order.Commission, err = s.serviceCommission.Init(ctx, &pb.Query_CalculateCommission{
			User: owner, Order: &_order,
		}); err != nil {
			log.Printf("CreateOrder: Warning! %s", err)
		}

		order, err := s.repo.CreateOrder(&_order)
		if err != nil {
			return nil, fmt.Errorf("CreateOrder: %s ", err)
		}

		s.Notify(&pb.Event{Type: &pb.Event_NewOrder{
			NewOrder: &pb.EventNewOrder{Order: order},
		}})

		return &pb.Response_Order{
			Created:     true,
			Object:      order,
			QueryStatus: pb.QueryStatus_Query_Success,
		}, nil
	}
	return nil, fmt.Errorf("CreateOrder: operation not permited. Empty context")
}

func (s *ServiceOrder) GetOrder(ctx context.Context, req *pb.Query_Order) (*pb.Response_Order, error) {
	o, err := s.repo.GetOrder(req.Id)
	if err != nil {
		return nil, fmt.Errorf("GetOrder: %s", err)
	}

	return &pb.Response_Order{
		Object:      o,
		QueryStatus: pb.QueryStatus_Query_Success,
	}, nil
}

func (s *ServiceOrder) GetOrders(ctx context.Context, req *pb.Query_Order) (*pb.Response_Order, error) {
	ords, err := s.repo.GetOrders(req)
	if err != nil {
		return nil, fmt.Errorf("GetOrders: %s", err)
	}
	return &pb.Response_Order{
		Items:       ords,
		ItemsCount:  uint32(len(ords)),
		QueryStatus: pb.QueryStatus_Query_Success,
	}, nil
}

func (s *ServiceOrder) CancelOrder(ctx context.Context, req *pb.Query_Order) (*pb.Response_Order, error) {
	pid := ctx.Value("pid")
	if pid != nil {
		_o, err := s.repo.GetOrder(req.Id)
		if err != nil {
			return nil, fmt.Errorf("CancelOrder: %s ", err)
		}
		if _o == nil {
			return nil, fmt.Errorf("CancelOrder: not find ")
		}

		if _o.OwnerId != pid {
			return nil, fmt.Errorf("CancelOrder: operation not permitted")
		}

		if _o.Status.Status == pb.DealStatus_PERFORMED || _o.Status.Status == pb.DealStatus_CANCELED {
			return nil, fmt.Errorf("CancelOrder: order cannot be canceled")
		}

		_o.Available = 0
		_o.Status = &pb.DealStatus{
			Status:    pb.DealStatus_CANCELED,
			CreatedAt: helper.CurrentTimestamp(),
		}
		o, err := s.repo.UpdateOrderAvailable(_o)
		if err != nil {
			return nil, fmt.Errorf("CancelOrder: %s ", err)
		}

		s.Notify(&pb.Event{Type: &pb.Event_OrderChange{
			OrderChange: &pb.EventOrderChange{Order: o},
		}})

		s.Notify(&pb.Event{Type: &pb.Event_OrderCanceled{
			OrderCanceled: &pb.EventOrderCanceled{Order: o},
		}})

		return &pb.Response_Order{
			Object:      o,
			Canceled:    true,
			QueryStatus: pb.QueryStatus_Query_Success,
		}, nil
	}
	return nil, fmt.Errorf("CreateContract: operation not permited. Empty context")
}

func (s *ServiceOrder) CreateContract(ctx context.Context, q *pb.Query_CreateContract) (*pb.Response_Contract, error) {
	pid := ctx.Value("pid")
	if pid != nil {
		o_resp, _ := s.GetOrder(ctx, &pb.Query_Order{Id: q.OrderId})
		if o_resp.Object == nil {
			return nil, status.Error(codes.InvalidArgument, "CreateContract: undefined order")
		}
		o := o_resp.Object

		// Order на меньше предполагаемого контракта
		if q.Amount <= 0 || o.Available < q.Amount {
			return nil, status.Error(codes.Unavailable, fmt.Sprintf("CreateContract: not enought order available"))
		}

		p_resp, _ := s.serviceUser.GetUser(ctx, &pb.Query_User{Id: pid.(string)})
		if p_resp == nil || p_resp.Object.Id == o.OwnerId {
			return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("CreateContract: unavailable buyer"))
		}

		// Проверка кошелька на оплату
		sba_resp, err := s.serviceAccount.GetAccount(ctx, &pb.Query_Account{Address: q.SendingAddress})
		if err != nil {
			return nil, status.Error(codes.Unavailable, fmt.Sprintf("CreateContract: %s ", err))
		}
		sbab_resp, err := s.serviceAccountBalance.GetBalance(ctx, &pb.Query_Account{Address: sba_resp.Object.Address})
		if err != nil {
			return nil, status.Error(codes.Internal, fmt.Sprintf("CreateContract: %s ", err))
		}
		if sba_resp == nil {
			return nil, fmt.Errorf("CreateContract: Undefined sending account ")
		}
		if sba_resp.Object.Currency.Symbol != o.BuyCurrencySymbol {
			return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("CreateContract: Unavailable currency(%s) in a buyer sending account(%s) ", o.BuyCurrencySymbol, sba_resp.Object.Currency.Symbol))
		}
		if sba_resp.Object.OwnerId != pid.(string) {
			return nil, status.Error(codes.PermissionDenied, "CreateContract: Unavailable owner in a sending account ")
		}
		if sba_resp.Object.Type != pb.AccountType_INTERNAL {
			return nil, status.Error(codes.InvalidArgument, "CreateContract: Unavailable sending account type ")
		}
		if sbab_resp == nil {
			return nil, status.Error(codes.Internal, "CreateContract: Account balance getting error, nil getting ")
		}
		if sbab_resp.Object.Available < (q.Amount * o.Price) {
			return nil, status.Errorf(codes.InvalidArgument, "CreateContract: Account balance getting error: %f < %f ", sbab_resp.Object.Available, q.Amount*o.Price)
		}

		rba_resp, _ := s.serviceAccount.GetAccount(ctx, &pb.Query_Account{Address: q.ReceiveAddress})
		if rba_resp.Object == nil || rba_resp.Object.OwnerId != p_resp.Object.Id ||
			rba_resp.Object.Currency.Symbol != o.SellCurrencySymbol {
			return nil, fmt.Errorf("CreateContract: Unavailable receive account")
		}

		_c := &pb.Contract{
			OrderId:              o.Id,
			BuyerId:              p_resp.Object.Id,
			SellerId:             o.OwnerId,
			BuyerSendAddress:     sba_resp.Object.Address,
			BuyerReceiveAddress:  rba_resp.Object.Address,
			SellerSendAddress:    o.SendingAddress,
			SellerReceiveAddress: o.ReceiveAddress,
			Amount:               q.Amount,
			Cost:                 q.Amount * o.Price,
			Available:            q.Amount,
			Price:                o.Price,
			FrontMetaData:        q.FrontMetaData,
		}

		// Commission
		if _c.SellerCommission, err = s.serviceCommission.Init(ctx, &pb.Query_CalculateCommission{
			User: &pb.User{Id: o.OwnerId}, Order: o, Contract: _c,
		}); err == nil && o.Commission != nil {
			o.Commission.Remainder -= _c.SellerCommission.Amount
		}

		_c.BuyerCommission, _ = s.serviceCommission.Init(ctx, &pb.Query_CalculateCommission{
			User: p_resp.Object, Order: o, Contract: _c,
		})

		c, err := s.repo.CreateContract(_c)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "CreateContract: %s", err)
		}

		o.Available -= c.Amount
		if o.Available == 0 {
			o.Status.Status = pb.DealStatus_PERFORMED
			o.Status.CreatedAt = helper.CurrentTimestamp()
		}

		if _, err := s.repo.UpdateOrderAvailable(o); err != nil {
			return nil, status.Errorf(codes.Internal, "CreateContract: %s ", err)
		}

		log.Printf("CreateContract: Created. %s. Amount: %f Price: %f Cost: %f \n", c.Id, c.Amount, c.Price, c.Cost)

		// Выбрасываем события
		s.Notify(&pb.Event{Type: &pb.Event_OrderChange{
			OrderChange: &pb.EventOrderChange{Order: o},
		}})

		if o.Status.Status == pb.DealStatus_PERFORMED {
			s.Notify(&pb.Event{Type: &pb.Event_OrderPerformed{
				OrderPerformed: &pb.EventOrderPerformed{Order: o},
			}})
		}

		s.Notify(&pb.Event{Type: &pb.Event_NewContract{
			NewContract: &pb.EventNewContract{Contract: c},
		}})
		//

		return &pb.Response_Contract{
			Created:     true,
			Object:      c,
			QueryStatus: pb.QueryStatus_Query_Success,
		}, nil
	}
	return nil, fmt.Errorf("CreateContract: operation not permited. Empty context")
}

func (s *ServiceOrder) GetContract(ctx context.Context, req *pb.Query_Contract) (*pb.Response_Contract, error) {
	c, err := s.repo.GetContract(req.Id)
	if err != nil {
		return nil, fmt.Errorf("GetContract: %s", err)
	}

	return &pb.Response_Contract{
		Object:      c,
		QueryStatus: pb.QueryStatus_Query_Success,
	}, nil
}

func (s *ServiceOrder) GetContracts(ctx context.Context, req *pb.Query_Contract) (*pb.Response_Contract, error) {
	cts, err := s.repo.GetContracts(req)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "GetContracts: %s", err)
	}
	return &pb.Response_Contract{
		QueryStatus: pb.QueryStatus_Query_Success,
		Items:       cts,
		ItemsCount:  uint32(len(cts)),
	}, nil
}

func (s *ServiceOrder) IsSellerContractTx(ctx context.Context, tx *pb.Tx) (*pb.Contract, error) {
	c_resp, err := s.GetContracts(ctx, &pb.Query_Contract{
		SellerSendAddress:   tx.FromAddress,
		BuyerReceiveAddress: tx.ToAddress,
		Active:              true,
	})
	if err != nil {
		return nil, fmt.Errorf("IsSellerContractTx: %s", err)
	}

	if c_resp == nil || c_resp.ItemsCount > 1 {
		return nil, fmt.Errorf("IsSellerContractTx: Contract not find or came contradictory data(%d) ", c_resp.ItemsCount)
	} else if c_resp.ItemsCount == 1 {
		return c_resp.Items[0], nil
	}
	return nil, nil
}

func (s *ServiceOrder) IsBuyerContractTx(ctx context.Context, tx *pb.Tx) (*pb.Contract, error) {
	c_resp, err := s.GetContracts(ctx, &pb.Query_Contract{
		BuyerSendAddress:     tx.FromAddress,
		SellerReceiveAddress: tx.ToAddress,
		Active:               true,
	})

	if err != nil {
		return nil, fmt.Errorf("IsBuyerContractTx: %s", err)
	}

	if c_resp == nil || c_resp.ItemsCount > 1 {
		return nil, fmt.Errorf("IsBuyerContractTx: Contract not find or came contradictory data(%d) ", c_resp.ItemsCount)
	} else if c_resp.ItemsCount == 1 {
		return c_resp.Items[0], nil
	}
	return nil, nil
}

func (s *ServiceOrder) CreateWithdrawal(ctx context.Context, q *pb.Query_Withdrawal) (*pb.WithdrawalOrder, error) {
	pid := ctx.Value("pid")
	if pid != nil {
		//Проверка инициатора
		p_resp, err := s.serviceUser.GetUser(ctx, &pb.Query_User{Id: pid.(string)})
		if p_resp == nil || err != nil {
			//if p_resp, err := s.serviceParticipant.GetParticipant(ctx, &pb.Query_Participant{ Id: q.OwnerId }); p_resp == nil || err != nil {
			return nil, fmt.Errorf("CreateWithdrawal: Undefined participant %s", pid.(string))
		}
		owner := p_resp.Object

		//Проверка запроса
		if q.Amount <= 0 {
			return nil, fmt.Errorf("CreateWithdrawal: Not avaliable amount")
		}

		//Проверка счетов
		sa_resp, _ := s.serviceAccount.GetAccount(ctx, &pb.Query_Account{Address: q.SendingAddress})
		if sa_resp == nil || sa_resp.Object.OwnerId != owner.Id {
			return nil, fmt.Errorf("CreateWithdrawal: Undefined account")
		}

		// Проверка баланса для внутренних кошельков
		if sa_resp.Object.Type == pb.AccountType_INTERNAL {
			if sa_blnc_resp, _ := s.serviceAccountBalance.GetBalance(ctx, &pb.Query_Account{Address: sa_resp.Object.Address}); sa_blnc_resp == nil || sa_blnc_resp.Object.Available < q.Amount {
				return nil, fmt.Errorf("CreateWithdrawal: Not enought balance")
			}
		}

		_wo := pb.WithdrawalOrder{
			OwnerId:        owner.Id,
			SendingAddress: q.SendingAddress,
			Amount:         q.Amount,
			PaymentSystem:  q.PaymentSystem,
		}

		//
		for _, at := range q.Attributes {
			log.Printf("CreateWithdrawal: Нужна валидация! %s - %s", at.Key, at.Value)
		}

		_wo.Attributes = q.Attributes

		wo, err := s.repo.CreateWithdrawal(&_wo)
		if err != nil {
			return nil, fmt.Errorf("CreateWithdrawal: %s ", err)
		}

		_txresp, err := s.serviceTransaction.CreateTx(ctx, &pb.Tx{
			FromAddress: q.SendingAddress,
			Amount:      q.Amount,
			Reason: &pb.TxReason{
				WithdrawalOrderId: wo.Id,
				Status:            pb.TxReason_WITHDRAW_TX,
			},
		})
		if err != nil {
			return nil, fmt.Errorf("CreateWithdrawal: %s", err)
		}

		wo.RelatedTxId = _txresp.Object.Id

		s.repo.UpdateWithdrawal(wo)

		s.Notify(&pb.Event{Type: &pb.Event_NewWithdrawalOrder{
			NewWithdrawalOrder: &pb.EventNewWithdrawal{WithdrawalOrder: wo},
		}})

		return wo, nil
	}

	return nil, fmt.Errorf("CreateWithdrawal: operation not permited. Empty context")
}

func (s *ServiceOrder) PerformWithdrawal(ctx context.Context, q *pb.Query_Withdrawal) (*pb.WithdrawalOrder, error) {
	if aid := ctx.Value("admin-id"); aid == nil {
		return nil, status.Error(codes.PermissionDenied, "PerformWithdrawal: not allowed")
	}
	if len(q.Id) == 0 || len(q.RelatedTxId) == 0 {
		return nil, status.Error(codes.InvalidArgument, "PerformWithdrawal: Bad Request")
	}
	wos, err := s.repo.GetWithdrawal(q)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "PerformWithdrawal: %s", err)
	}
	wo := wos[0]

	rtx, err := s.serviceTransaction.GetTx(ctx, &pb.Query_Tx{TxId: q.RelatedTxId})
	if err != nil || rtx.ItemsCount == 0 || rtx.ItemsCount > 1 {
		return nil, status.Errorf(codes.InvalidArgument, "PerformWithdrawal: serviceTransaction.GetTx error %s", err)
	}
	tx := rtx.Items[0]
	if tx.Reason.Status != pb.TxReason_WITHDRAW_TX || tx.Reason.WithdrawalOrderId != wo.Id {
		return nil, status.Errorf(codes.InvalidArgument, "PerformWithdrawal: Related tx not find or related tx != wo.Tx. %s", err)
	}

	if tx.Status != pb.TxStatus_CONFIRMED {
		return nil, status.Errorf(codes.Aborted, "PerformWithdrawal: Related tx not confirmed")
	}

	wo.Status.Status = pb.DealStatus_PERFORMED

	s.repo.UpdateWithdrawal(wo)

	s.Notify(&pb.Event{Type: &pb.Event_WithdrawalPerformed{
		WithdrawalPerformed: &pb.EventWithdrawalPerformed{WithdrawalOrder: wo},
	}})

	return wo, nil
}

func (s *ServiceOrder) CancelWithdrawal(ctx context.Context, q *pb.Query_Withdrawal) (*pb.WithdrawalOrder, error) {
	return nil, nil
}

func (s *ServiceOrder) GetWithdrawalOrders(ctx context.Context, q *pb.Query_Withdrawal) (*pb.Response_Withdrawal, error) {
	aid := ctx.Value("admin-id")
	pid := ctx.Value("pid")
	if aid == nil && pid == nil {
		return nil, fmt.Errorf("GetWithdrawalOrders: not allowed")
	}

	if aid == nil && pid != nil {
		q.OwnerId = pid.(string)
	}

	wos, err := s.repo.GetWithdrawal(q)
	if err != nil {
		return nil, fmt.Errorf("GetWithdrawalOrders: %s", err)
	}

	return &pb.Response_Withdrawal{
		Items: wos,
	}, nil
}
