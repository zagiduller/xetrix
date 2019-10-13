package services

import (
	"context"
	"engine/lib/helper"
	"engine/lib/services/events"
	pb "engine/lib/structs"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

type ServiceTransactionProcessing struct {
	serviceAccount     *ServiceAccount
	serviceOrder       *ServiceOrder
	serviceTransaction *ServiceTransaction
	repo               ITransactionRepository
	bus                *events.Bus
}

// Принимает регистр платежных систем
func NewTransactionProcessingService(repo ITransactionRepository, srv_t *ServiceTransaction, srv_a *ServiceAccount, srv_o *ServiceOrder) *ServiceTransactionProcessing {

	return &ServiceTransactionProcessing{
		repo:               repo,
		serviceAccount:     srv_a,
		serviceTransaction: srv_t,
		serviceOrder:       srv_o,
	}
}

func (s *ServiceTransactionProcessing) AddEventBus(bus *events.Bus) {
	s.bus = bus
	bus.Subscribe(s,
		&pb.Event{Type: &pb.Event_PaySystemRawTx{}},
		&pb.Event{Type: &pb.Event_RelatedPaySystemPrepareTx{}},
		&pb.Event{Type: &pb.Event_RelatedPaySystemOutTx{}},
	)
}

func (s *ServiceTransactionProcessing) createPreparedEventFromTx(tx *pb.Tx) *pb.Event_ETHtxPrepared {
	respacc, err := s.serviceAccount.GetAccount(context.Background(), &pb.Query_Account{Address: tx.ToAddress})
	if err != nil {
		log.Printf("createPreparedEventFromTx: %s", err)
		return nil
	}
	usrresp, err := s.serviceAccount.serviceUser.GetUser(context.Background(), &pb.Query_User{Id: respacc.Object.OwnerId})
	if err != nil {
		log.Printf("createPreparedEventFromTx: %s", err)
		return nil
	}

	return &pb.Event_ETHtxPrepared{ETHtxPrepared: &pb.EventETHtxProcces{
		Currency: respacc.Object.Currency,
		Tx:       tx,
		OwnerTo:  usrresp.Object,
		To:       respacc.Object,
	}}
}

func (s *ServiceTransactionProcessing) Update(event *pb.Event) {
	switch event.Type.(type) {
	// Баланс аккаунта изменен
	case *pb.Event_PaySystemRawTx:
		s.fund(context.Background(), event.GetPaySystemRawTx().Raw)
		//case *pb.Event_RelatedPaySystemPrepareTx:
		//	// ethint отправил деньги со счета системы на счет пользователя для оплаты комиссии
		//	// меняем статус и обьявляем транзакцию подготовленной к выводу
		//	if tx, _ := s.repo.GetTx(&pb.Query_Tx{TxId: event.GetRelatedPaySystemPrepareTx().TxId}); tx != nil {
		//		tx.Reason.Status = pb.TxReason_FUND_WAIT_PREPARE_TX
		//		tx.Reason.PreparePStxId = event.GetRelatedPaySystemPrepareTx().RelatedId
		//		if err := s.updateTx(tx); err != nil {
		//			log.Printf("ServiceTransactionProcessing: Update: %v \t %s", err)
		//			return
		//		}
		//		// Не выбрасываем Event_ETHtxPrepared так как транзакция ешё не не подтверждена в блокчейне
		//		//if ev := s.createPreparedEventFromTx(tx); ev != nil {
		//		//	s.Notify(&pb.Event{Type: ev})
		//		//}
		//	}
		//case *pb.Event_RelatedPaySystemOutTx:
		//	// ethint отправил деньги со счета пользователя на счет мастера
		//	// привязать хэш отправления к транзакции в системе и обновить статус
		//	if tx, _ := s.repo.GetTx(&pb.Query_Tx{TxId: event.GetRelatedPaySystemOutTx().TxId}); tx != nil {
		//		tx.Reason.Status = pb.TxReason_FUND_PERFORMED_TX
		//		tx.Reason.OutPStxId = event.GetRelatedPaySystemOutTx().RelatedId
		//		if err := s.updateTx(tx); err != nil {
		//			log.Printf("ServiceTransactionProcessing: Update: %v \t %s", err)
		//			break
		//		}
		//		log.Printf("ServiceTransactionProcessing: Transaction (%s) performed with Out tx: %s", tx.Id, tx.Reason.OutPStxId)
		//	}
	}
}

func (s *ServiceTransactionProcessing) Notify(event *pb.Event) {
	s.bus.NewEvent(event)
}

// Пришла транзакция от продавца
// посчитать процент от контракта и получить эквивалентную комиссию
// для продавца и покупателя, создать под это связанные транзакции
func (s *ServiceTransactionProcessing) contractRelated(tx *pb.Tx, c *pb.Contract) ([]*pb.Tx, error) {
	// Созание структуры связных транзакций возможно только для контрактных транзакций которые не подтверждены
	if (tx.Status == pb.TxStatus_UNCONFIRMED) && (tx.Reason.Status == pb.TxReason_SELLER_CONTRACT_TX || tx.Reason.Status == pb.TxReason_BUYER_CONTRACT_TX) {
		var related []*pb.Tx
		var contractAmount float64
		var cms *pb.Commission

		// У контракта должны быть установлены комиссии
		// Транзакция продавца
		if tx.Reason.Status == pb.TxReason_SELLER_CONTRACT_TX && c.SellerCommission != nil {
			cms = c.SellerCommission
			contractAmount = c.Amount
		}

		if tx.Reason.Status == pb.TxReason_BUYER_CONTRACT_TX && c.BuyerCommission != nil {
			cms = c.BuyerCommission
			contractAmount = c.Amount * c.Price
		}

		if contractAmount > 0 && cms.Amount > 0 {
			cmsValue := cms.Amount * (tx.Amount / contractAmount)
			log.Printf("contractRelated: contractAmount: cms.Amount(%f) * tx.Amount(%f) / contractAmount(%f)", cms.Amount, tx.Amount, contractAmount)
			if cmsValue > cms.Remainder {
				log.Printf("contractRelated: Error! cmsValue > cms.Remainder")
			}
			rtx := pb.Tx{
				FromAddress: cms.SendingAddress,
				ToAddress:   cms.ReceiveAddress,
				Amount:      cmsValue,
				Reason: &pb.TxReason{
					ContractId: c.Id,
					Status:     pb.TxReason_CONTRACT_COMMISSION_TX,
				},
			}
			cms.Remainder -= cmsValue
			related = append(related, &rtx)
		}

		return related, nil
	}
	return nil, fmt.Errorf("contractRelated: can not set related-txs for this tx")
}

func (s *ServiceTransactionProcessing) UnderstandingRawTx(ctx context.Context, q *pb.Query_RawTx) (*pb.Response_Tx, error) {

	_tx := pb.Tx{}
	fa_resp, _ := s.serviceAccount.GetAccount(ctx, &pb.Query_Account{Address: q.FromAddress})
	if fa_resp != nil {
		_tx.FromAddress = fa_resp.Object.Address
		_tx.CurrencySymbol = fa_resp.Object.Currency.Symbol
		_tx.FromAddressOwnerId = fa_resp.Object.OwnerId
	}

	ta_resp, _ := s.serviceAccount.GetAccount(ctx, &pb.Query_Account{Address: q.ToAddress})
	if ta_resp != nil {
		_tx.ToAddress = ta_resp.Object.Address
		_tx.CurrencySymbol = ta_resp.Object.Currency.Symbol
		_tx.ToAddressOwnerId = ta_resp.Object.OwnerId

	}

	if fa_resp == nil && ta_resp == nil {
		return nil, fmt.Errorf("UnderstandingRawTx: Undefined tx")
	}

	// Проверяем сущестувет ли транзакция платежной системы
	// Нужно сделать это до отправки событий Event_ETHtxPrepared
	// Так как это приводит к повторной отправке средств на мастер счета
	if len(q.InPStxId) > 0 { // Это транзакция пополнения
		if ta_resp.Object != nil && q.BlockNumber < ta_resp.Object.BlockNumber {
			return nil, fmt.Errorf("UnderstandingRawTx: PS tx (%s) happened (BN: %d) before (%s)address applied to system (BN: %d)", q.InPStxId, q.BlockNumber, ta_resp.Object.Address, ta_resp.Object.BlockNumber)
		}

		if fromPsRelatedTx, _ := s.serviceTransaction.GetTx(ctx, &pb.Query_Tx{InPStxId: q.InPStxId}); fromPsRelatedTx != nil && fromPsRelatedTx.ItemsCount > 0 {
			orphanTx := fromPsRelatedTx.Items[0]
			//
			// Не актуально! Изменена концепция вывода средств со счета
			// Если статус транзакции не исполненный и если нет OUT транзакции,
			// его следует его поменять
			if orphanTx.Status != pb.TxStatus_CANCELED && orphanTx.Reason.Status != pb.TxReason_FUND_PERFORMED_TX {
				// есть OUT транзакция, помечаем готовой
				if len(orphanTx.Reason.OutPStxId) > 0 {
					orphanTx.Reason.Status = pb.TxReason_FUND_PERFORMED_TX
					if _, err := s.repo.UpdateTx(orphanTx); err != nil {
						log.Printf("UnderstandingRawTx: Unhandled! Error update tx: %v \t %s", err)
					}
				} else {
					// OUT транзакции нет. Оставляем это на ручной режим
				}

			}
			return nil, fmt.Errorf("UnderstandingRawTx: related from InPStxId transaction is exist on"+
				" \n\ttxId: %s \n\tSymbol: %s \n\tStatus: %s, \n\tReasonStatus: %s \n\tinPStxId: %s \n\tBlockNumber: %d\n",
				orphanTx.Id, orphanTx.CurrencySymbol, orphanTx.Status.String(),
				orphanTx.Reason.Status.String(), q.InPStxId, q.BlockNumber)
		}
	}

	_tx.Amount = q.Amount

	if len(_tx.FromAddress) == 0 { //Fund'
		// Если счет пополняется в токенах или эфирах и это не системное пополнение
		if len(q.InPStxId) > 0 && (ta_resp.Object.Currency.Type == pb.Currency_ETH_CONTRACT_TOKEN || ta_resp.Object.Currency.Symbol == "ETH") {

			if respUsr, err := s.serviceAccount.serviceUser.GetUser(ctx, &pb.Query_User{Id: ta_resp.Object.OwnerId}); err == nil && respUsr.Object != nil {
				if ta_resp.Object.Currency.Symbol == "ETH" {
					_tx.Reason = &pb.TxReason{Status: pb.TxReason_FUND_PREPARED_TX, InPStxId: q.InPStxId}
					//Отравляем эвент о готовности перевода денег с этого счета
					s.Notify(&pb.Event{Type: &pb.Event_ETHtxPrepared{ETHtxPrepared: &pb.EventETHtxProcces{
						Currency: ta_resp.Object.Currency,
						Tx:       &_tx,
						OwnerTo:  respUsr.Object,
						To:       ta_resp.Object,
					}}})
				} else {
					_tx.Reason = &pb.TxReason{Status: pb.TxReason_FUND_UNPERFORMED_TX, InPStxId: q.InPStxId}
				}
			}
		} else {
			// Либо фиат, либо пополнение вручную
			_tx.Reason = &pb.TxReason{Status: pb.TxReason_FUND_PERFORMED_TX, InPStxId: q.InPStxId}
		}

	} else if len(_tx.ToAddress) == 0 { //Withdraw
		_tx.Reason = &pb.TxReason{Status: pb.TxReason_WITHDRAW_TX}
	} else { //Contract
		c, err := s.serviceOrder.IsSellerContractTx(ctx, &_tx)
		if err != nil {
			return nil, fmt.Errorf("UnderstandingRawTx: %s", err)
		}
		if c == nil {
			c, err = s.serviceOrder.IsBuyerContractTx(ctx, &_tx)
			if err != nil {
				return nil, fmt.Errorf("UnderstandingRawTx: %s", err)
			}
			if c == nil {
				return nil, fmt.Errorf("UnderstandingRawTx: Tx not have a reason :)")
			}
			_tx.Reason = &pb.TxReason{Status: pb.TxReason_BUYER_CONTRACT_TX, ContractId: c.Id}
			if _tx.Related, err = s.contractRelated(&_tx, c); err != nil {
				log.Printf("UnderstandingRawTx: %s", err)
			}
		} else {
			_tx.Reason = &pb.TxReason{Status: pb.TxReason_SELLER_CONTRACT_TX, ContractId: c.Id}
			if _tx.Related, err = s.contractRelated(&_tx, c); err != nil {
				log.Printf("UnderstandingRawTx: %s", err)
			}
		}
	}

	return &pb.Response_Tx{
		Object:      &_tx,
		QueryStatus: pb.QueryStatus_Query_Success,
	}, nil
}

func (s *ServiceTransactionProcessing) ConfirmTx(ctx context.Context, query *pb.Query_Tx) (*pb.Response_Tx, error) {
	//Проверка. Возможно ли подтверждение?
	if _tx, _ := s.repo.GetTx(query); _tx != nil && _tx.Reason.Status != pb.TxReason_UNREASON_TX {
		if _tx.Status == pb.TxStatus_CONFIRMED || _tx.Status == pb.TxStatus_CANCELED {
			return nil, status.Error(codes.InvalidArgument, "ConfirmTx: tx status error. I can't confirm")
		}
		if ok, err := s.repo.ConfirmTx(_tx); err != nil || !ok {
			return nil, status.Errorf(codes.InvalidArgument, "ConfirmTx: not confirmed. %s", err)
		}

		if _tx.Reason.Status == pb.TxReason_SELLER_CONTRACT_TX {
			log.Println("ConfirmTx: Init contract tx ")
			if _, err := s.ContractTx(ctx, _tx); err != nil {
				return nil, status.Errorf(codes.Aborted, "ConfirmTx: %s", err)
			}
		}

		//Связаные транзакции во вложенной структуре,
		//Подтверждаем
		for _, rtx := range _tx.Related {
			if _, err := s.ConfirmTx(ctx, &pb.Query_Tx{TxId: rtx.Id}); err != nil {
				log.Printf("ConfirmTx: Related tx confirm Error! %s ", err)
			} else {
				log.Printf("ConfirmTx: WARNING! Related tx confirmation error %s", err)
			}

		}

		log.Printf("ConfirmTx: Tx %s confirmed. Reason: %s", _tx.Id, _tx.Reason.Status.String())

		//Подтверждена транзакция
		s.Notify(&pb.Event{Type: &pb.Event_TxConfirm{TxConfirm: &pb.EventTxConfirm{Tx: _tx}}})

		return &pb.Response_Tx{
			Confirmed:   true,
			Object:      _tx,
			QueryStatus: pb.QueryStatus_Query_Success,
		}, nil
	}
	return nil, fmt.Errorf("ConfirmTx: not confirmed")
}

//stx - seller tx
func (s *ServiceTransactionProcessing) ContractTx(ctx context.Context, stx *pb.Tx) (*pb.Tx, error) {

	if stx.Status == pb.TxStatus_CONFIRMED {

		log.Println("ContractTx: Tx creation start")
		c_resp, err := s.serviceOrder.GetContracts(ctx, &pb.Query_Contract{
			SellerSendAddress:   stx.FromAddress,
			BuyerReceiveAddress: stx.ToAddress,
			Active:              true,
		})

		if c_resp == nil || c_resp.ItemsCount != 1 {
			return nil, fmt.Errorf("ContractTx: Contract not find or came contradictory data(%d). %s", c_resp.ItemsCount, err)
		}

		c := c_resp.Items[0]

		var c_amount float64
		c_amount = stx.Amount

		if stx.Amount > c.Available {
			c_amount = c.Available
		}

		q_rtx := &pb.Query_RawTx{
			FromAddress: c.BuyerSendAddress,
			ToAddress:   c.SellerReceiveAddress,
			Amount:      c.Price * c_amount,
		}

		rtx_resp, err := s.UnderstandingRawTx(ctx, q_rtx)
		if err != nil {
			return nil, fmt.Errorf("ContractTx: Understandig raw error: %s ", err)
		}

		btx_resp, err := s.serviceTransaction.CreateTx(ctx, rtx_resp.Object)
		if err != nil {
			return nil, fmt.Errorf("ContractTx: %s", err)
		}

		if btx_resp == nil || btx_resp.Object.Reason.Status != pb.TxReason_BUYER_CONTRACT_TX || btx_resp.Object.Reason.ContractId != c.Id {
			return nil, fmt.Errorf("ContractTx: Error create contract-tx ")
		}
		log.Printf("ContractTx: Contract tx created: %s ", btx_resp.Object.Id)

		resp, err := s.ConfirmTx(ctx, &pb.Query_Tx{TxId: btx_resp.Object.Id})
		if err != nil {
			return nil, fmt.Errorf("ContractTx: Error contract-tx confirm %s ", err)
		}

		log.Println("ContractTx: Contract tx confirmed")

		Mu.Lock()
		c.Available -= c_amount
		if c.Available == 0 {
			c.Status.Status = pb.DealStatus_PERFORMED
			c.Status.CreatedAt = helper.CurrentTimestamp()
		}
		Mu.Unlock()

		if _, err := s.repo.UpdateContractAvailable(c); err != nil {
			return nil, fmt.Errorf("ContractTx: %s ", err)
		}

		log.Println("ContractTx: Buyer contract-tx created")

		return resp.Object, nil

	}
	return nil, fmt.Errorf("ContractTx: Unconfirmed tx ")
}

// Метод создан для автоматического исполнения контрактов
func (s *ServiceTransactionProcessing) CreateInternalContract(ctx context.Context, q *pb.Query_CreateContract) (*pb.Response_Contract, error) {
	pid := ctx.Value("pid")
	if pid != nil {
		c_resp, err := s.serviceOrder.CreateContract(ctx, q)
		if err != nil {
			return nil, fmt.Errorf("CreateInternalContract: %s", err)
		}
		c := c_resp.Object
		_tx_resp, err := s.UnderstandingRawTx(ctx, &pb.Query_RawTx{FromAddress: c.SellerSendAddress, ToAddress: c.BuyerReceiveAddress, Amount: c.Amount})
		if err != nil {
			return nil, fmt.Errorf("CreateInternalContract: %s", err)
		}
		if _tx_resp.Object.Reason.Status != pb.TxReason_SELLER_CONTRACT_TX {
			return nil, fmt.Errorf("CreateInternalContract: Unavailable tx reason")
		}
		tx_resp, err := s.serviceTransaction.CreateTx(ctx, _tx_resp.Object)
		if err != nil {
			return nil, fmt.Errorf("CreateInternalContract: %s", err)
		}
		if _, err := s.ConfirmTx(ctx, &pb.Query_Tx{TxId: tx_resp.Object.Id}); err != nil {
			return nil, fmt.Errorf("CreateInternalContract: %s", err)
		}
		return s.serviceOrder.GetContract(ctx, &pb.Query_Contract{Id: c.Id})
	}
	return nil, fmt.Errorf("CreateInternalContract: operation not permited. Empty context")
}

func (s *ServiceTransactionProcessing) fund(ctx context.Context, qtx *pb.Query_RawTx) bool {
	_tx_resp, err := s.UnderstandingRawTx(ctx, qtx)
	if err != nil {
		log.Printf("ServiceTransactionProcessing: fund: %s", err)
		return false
	}
	tx_resp, err := s.serviceTransaction.CreateTx(ctx, _tx_resp.Object)
	if err != nil {
		log.Printf("ServiceTransactionProcessing: fund: %s", err)
		return false
	}
	_, err = s.ConfirmTx(ctx, &pb.Query_Tx{TxId: tx_resp.Object.Id})
	if err != nil {
		log.Printf("ServiceTransactionProcessing: fund: %s", err)
		return false
	}
	log.Printf("ServiceTransactionProcessing: %s funded %f", qtx.ToAddress, qtx.Amount)
	return true
}

func (s *ServiceTransactionProcessing) updateTx(tx *pb.Tx) error {
	if _, err := s.repo.UpdateTx(tx); err != nil {
		//TODO разберись уже с логами!
		return fmt.Errorf("updateTx: ", err)
	}

	// Изменилась транзакция, выбрасываем эвент
	s.Notify(&pb.Event{Type: &pb.Event_TxProccessUpdate{TxProccessUpdate: &pb.EventTxProccessUpdate{Tx: tx}}})

	return nil
}

func (s *ServiceTransactionProcessing) NeedPrepareTxs(ctx context.Context, query *pb.Query_TxsPrepareProccess) (*pb.Empty, error) {
	for _, txid := range query.TxsId {
		if tx, err := s.repo.GetTx(&pb.Query_Tx{TxId: txid}); err == nil {
			if event := s.createPreparedEventFromTx(tx); event != nil {
				txProccessEvent := event.ETHtxPrepared
				s.Notify(&pb.Event{
					Type: &pb.Event_ETHtxNeedPrepare{
						ETHtxNeedPrepare: txProccessEvent,
					},
				})
			}
		}
	}
	return &pb.Empty{}, nil
}

func (s *ServiceTransactionProcessing) TxsPrepared(ctx context.Context, query *pb.Query_TxsPrepareProccess) (*pb.Empty, error) {
	for _, txid := range query.TxsId {
		// Проверяю на TxReason_FUND_PERFORMED_TX что бы не допустить всяких кошмаров с повторной отправкой
		if tx, err := s.repo.GetTx(&pb.Query_Tx{TxId: txid}); err == nil &&
			(tx.Reason.Status == pb.TxReason_FUND_UNPERFORMED_TX ||
				tx.Reason.Status == pb.TxReason_FUND_WAIT_PREPARE_TX ||
				tx.Reason.Status == pb.TxReason_FUND_PREPARED_TX) && len(tx.Reason.OutPStxId) == 0 {
			tx.Reason.Status = pb.TxReason_FUND_PREPARED_TX
			err := s.updateTx(tx)
			if err != nil {
				return nil, fmt.Errorf("ServiceTransactionProcessing: TxsPrepared: ", err)
			}
			if event := s.createPreparedEventFromTx(tx); event != nil {
				s.Notify(&pb.Event{Type: event})
			}
		}
	}
	return &pb.Empty{}, nil
}
