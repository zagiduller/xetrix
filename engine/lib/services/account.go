package services

import (
	"engine/lib/payments"
	"engine/lib/services/events"
	pb "engine/lib/structs"
	"fmt"
	"github.com/google/uuid"
	"golang.org/x/net/context"
	"log"
)

type IAccountRepository interface {
	CreateAccount(acc *pb.Account) (*pb.Account, error)
	UpdateAccount(acc *pb.Account) (*pb.Account, error)
	GetAllAccount() ([]*pb.Account, error)
	FindAccountByAddress(address string) (*pb.Account, error)
	FindAccountById(address string) (*pb.Account, error)
	FindAccountByOwnerId(ownerId string) ([]*pb.Account, error)
	FindAccountByCurrencyAndOwnerId(id, currencySymbol string) (*pb.Account, error)
}

func NewAccountService(
	repo IAccountRepository,
	srv_c *ServiceCurrency,
	srv_u *ServiceUser,
	srv_t *ServiceTransaction,
	ps *payments.Registry,

) *ServiceAccount {
	return &ServiceAccount{
		serviceCurrency:    srv_c,
		serviceUser:        srv_u,
		serviceTransaction: srv_t,
		repo:               repo,
		payregistry:        ps,
	}
}

type ServiceAccount struct {
	serviceCurrency    *ServiceCurrency
	serviceUser        *ServiceUser
	serviceTransaction *ServiceTransaction
	repo               IAccountRepository
	bus                *events.Bus
	payregistry        *payments.Registry
}

func (s *ServiceAccount) AddEventBus(bus *events.Bus) {
	s.bus = bus

	bus.Subscribe(s,
		&pb.Event{Type: &pb.Event_NewUser{}},
		//&pb.Event{Type: &pb.Event_NewCurrency{}},
		&pb.Event{Type: &pb.Event_PaySystemAdded{}},
		&pb.Event{Type: &pb.Event_CurrencyActivated{}},
		&pb.Event{Type: &pb.Event_CurrencyDeactivated{}},
	)
}

func (s *ServiceAccount) Update(event *pb.Event) {
	switch event.Type.(type) {
	case *pb.Event_NewUser:
		s.GenerateInternalAccounts(context.Background(), event.GetNewUser().User)
	case *pb.Event_PaySystemAdded:
		admCtx := context.WithValue(context.Background(), "admin-id", "Event_machine")
		s.generateInternalByCurrencyForAllUsers(admCtx, event.GetPaySystemAdded().Currency)
	case *pb.Event_CurrencyActivated:
		s.activateByCurrency(event.GetCurrencyActivated().Currency)
	case *pb.Event_CurrencyDeactivated:
		s.activateByCurrency(event.GetCurrencyDeactivated().Currency)
	}
}

func (s *ServiceAccount) Notify(event *pb.Event) {
	s.bus.NewEvent(event)
}

func (s *ServiceAccount) CreateAccount(ctx context.Context, q *pb.Query_CreateAccount) (*pb.Response_Account, error) {
	//Убедиться что такого кошелька в базе нет
	// Проверка валюты
	resp_c, err := s.serviceCurrency.GetCurrency(ctx, &pb.Query_Currency{Id: q.CurrencyId, Name: q.CurrencyName, Symbol: q.CurrencySymbol})
	if resp_c == nil || err != nil {
		return nil, fmt.Errorf("CreateAccount: Currency not found. %s", err)
	}

	// Взять ID участника из контекста. Проверять его существование в middleware
	resp_p, err := s.serviceUser.GetUser(ctx, &pb.Query_User{Id: q.OwnerId})
	if err != nil {
		return nil, fmt.Errorf("CreateAccount: Participant not found. %s", err)
	}

	// Нельзя добавить внешний счет для фиатной валюты
	if q.Type == pb.AccountType_EXTERNAL && resp_c.Object.Type == pb.Currency_FIAT_CURRENCY {
		return nil, fmt.Errorf("CreateAccount: Operation not allowed")
	}

	if q.Type == pb.AccountType_EXTERNAL && len(q.Address) == 0 {
		return nil, fmt.Errorf("CreateAccount: Address not set for external type(%d) account", pb.AccountType_EXTERNAL)
	}

	if len(q.Address) == 0 {
		return nil, fmt.Errorf("CreateAccount: Invalid address")
	}

	var activeStatus pb.Account_Status
	activeStatus = pb.Account_INACTIVE

	if resp_c.Object.Active {
		activeStatus = pb.Account_ACTIVE
	}

	_acc := &pb.Account{
		OwnerId:     resp_p.Object.Id,
		Currency:    resp_c.Object,
		Address:     q.Address,
		Type:        q.Type,
		BlockNumber: q.BlockNumber,
		Status:      activeStatus,
	}

	acc, err := s.repo.CreateAccount(_acc)
	if err != nil {
		return nil, fmt.Errorf("CreateAccount: %s ", err)
	}

	s.Notify(&pb.Event{Type: &pb.Event_NewAccount{
		NewAccount: &pb.EventNewAccount{Account: acc},
	}})

	return &pb.Response_Account{
		Created:     true,
		Object:      acc,
		QueryStatus: pb.QueryStatus_Query_Success,
	}, nil
}

func (s *ServiceAccount) GetAllUserAccount(ctx context.Context, _ *pb.Empty) (*pb.Response_Account, error) {
	pid := ctx.Value("pid")
	log.Printf("GetAllUserAccount pid: %s", pid)
	if pid != nil {
		if accs, err := s.repo.FindAccountByOwnerId(pid.(string)); len(accs) > 0 && err == nil {
			return &pb.Response_Account{
				Items:       accs,
				ItemsCount:  uint32(len(accs)),
				QueryStatus: pb.QueryStatus_Query_Success,
			}, nil
		}
	}
	return nil, fmt.Errorf("GetAllUserAccount: Not find")
}

func (s *ServiceAccount) GetAllAccount(ctx context.Context, _ *pb.Empty) (*pb.Response_Account, error) {
	aid := ctx.Value("admin-id")
	log.Printf("GetAllAccount Admin-id: %s", aid)
	if aid != nil {
		if accs, err := s.repo.GetAllAccount(); len(accs) > 0 && err == nil {
			return &pb.Response_Account{
				Items:       accs,
				ItemsCount:  uint32(len(accs)),
				QueryStatus: pb.QueryStatus_Query_Success,
			}, nil
		}
	}
	return nil, fmt.Errorf("GetAllAccount: Not find")
}

// TODO функция работает не логично и очень опасно! Нужно привести в адекватный вид логику условий
func (s *ServiceAccount) GetAccount(ctx context.Context, q *pb.Query_Account) (*pb.Response_Account, error) {
	if len(q.AccountId) > 0 {
		if acc, err := s.repo.FindAccountById(q.AccountId); acc != nil && err == nil {
			return &pb.Response_Account{
				Object:      acc,
				QueryStatus: pb.QueryStatus_Query_Success,
			}, nil
		}
	}
	if len(q.Address) > 0 {
		if acc, err := s.repo.FindAccountByAddress(q.Address); acc != nil && err == nil {
			return &pb.Response_Account{
				Object:      acc,
				QueryStatus: pb.QueryStatus_Query_Success,
			}, nil
		}
	}

	if len(q.ParticipantId) > 0 {
		if len(q.CurrencySymbol) > 0 {
			if acc, err := s.repo.FindAccountByCurrencyAndOwnerId(q.ParticipantId, q.CurrencySymbol); acc != nil && err == nil {
				return &pb.Response_Account{
					Object:      acc,
					QueryStatus: pb.QueryStatus_Query_Success,
				}, nil
			} else {
				return nil, err
			}
		}

		if accs, err := s.repo.FindAccountByOwnerId(q.ParticipantId); len(accs) > 0 && err == nil {
			return &pb.Response_Account{
				Items:       accs,
				ItemsCount:  uint32(len(accs)),
				QueryStatus: pb.QueryStatus_Query_Success,
			}, nil
		}
	}

	return nil, fmt.Errorf("GetAccount: Account not found")
}

func (s *ServiceAccount) GenerateInternalAccounts(ctx context.Context, p *pb.User) (*pb.Response_Account, error) {
	cs, _ := s.serviceCurrency.GetCurrency(ctx, &pb.Query_Currency{})

	resp := &pb.Response_Account{}

	for _, c := range cs.Items {
		acc, err := s.generateInternal(ctx, c, p)
		if err != nil {
			log.Printf("GenerateInternalAccounts: %s", err)
			continue
		}
		resp.Items = append(resp.Items, acc)
	}
	return resp, nil
}

func (s *ServiceAccount) generateInternalByCurrencyForAllUsers(ctx context.Context, c *pb.Currency) error {
	respusr, err := s.serviceUser.GetAllUsers(ctx, &pb.Empty{})
	if err != nil {
		return fmt.Errorf("generateInternalByCurrencyForAllUsers: %s", err)
	}

	for _, u := range respusr.Items {
		// Можно попробовать запускать в отдельной горутине
		if _, err := s.generateInternal(ctx, c, u); err != nil {
			log.Printf("generateInternalByCurrencyForAllUsers: %s skip", err)
		}
	}
	return nil
}

func (s *ServiceAccount) generateInternal(ctx context.Context, c *pb.Currency, u *pb.User) (*pb.Account, error) {
	if find, _ := s.GetAccount(ctx, &pb.Query_Account{ParticipantId: u.Id, CurrencySymbol: c.Symbol}); find != nil {
		return nil, fmt.Errorf("generateInternal: account by currency %s already exist. Userid: %s. \n %v \n", c.Symbol, u.Id, find)
	}

	addr, bn := s.payregistry.GenerateAddress(u, c.Symbol)
	if len(addr) == 0 {
		u2id := uuid.New()
		addr = c.Symbol + "_" + u2id.String()
	}

	q_acc := &pb.Query_CreateAccount{
		OwnerId:      u.Id,
		CurrencyName: c.Name,
		Address:      addr,
		Type:         pb.AccountType_INTERNAL,
		BlockNumber:  bn,
		Active:       c.Active,
	}
	resp_acc, err := s.CreateAccount(ctx, q_acc)
	if err != nil {
		return nil, fmt.Errorf("generateInternal: %s", err)
	}
	log.Printf("Internal address generated: (%s) %s BN: %d", resp_acc.Object.Currency.Symbol, resp_acc.Object.Address, resp_acc.Object.BlockNumber)
	return resp_acc.Object, nil
}

func (s *ServiceAccount) deactivateByCurrency(c *pb.Currency) error {
	accs, err := s.repo.GetAllAccount()
	if err != nil {
		return fmt.Errorf("deactivateByCurrency: %s", err)
	}

	log.Printf("Deactivate accounts by %s currency ", c.Name)

	for _, a := range accs {
		if a.Currency.Symbol == c.Symbol {
			if err := s.deactivate(a); err != nil {
				log.Printf("deactivateByCurrency: %s skip", err)
				continue
			}
		}
	}

	return nil
}

func (s *ServiceAccount) activateByCurrency(c *pb.Currency) error {
	accs, err := s.repo.GetAllAccount()
	if err != nil {
		return fmt.Errorf("activateByCurrency: %s", err)
	}

	log.Printf("Activate accounts by %s currency ", c.Name)

	for _, a := range accs {
		if a.Currency.Symbol == c.Symbol {
			if err := s.activate(a); err != nil {
				log.Printf("activateByCurrency: %s skip", err)
				continue
			}
		}
	}

	return nil
}

func (s *ServiceAccount) activate(a *pb.Account) error {
	a.Status = pb.Account_ACTIVE
	if _, err := s.repo.UpdateAccount(a); err != nil {
		return err
	}
	return nil
}

func (s *ServiceAccount) deactivate(a *pb.Account) error {
	a.Status = pb.Account_INACTIVE
	if _, err := s.repo.UpdateAccount(a); err != nil {
		return err
	}
	return nil
}

/////
type IAccountBalanceRepository interface {
	GetContracts(req *pb.Query_Contract) ([]*pb.Contract, error)
	GetOrders(req *pb.Query_Order) ([]*pb.Order, error)
	GetWithdrawal(qw *pb.Query_Withdrawal) ([]*pb.WithdrawalOrder, error)
	FindAddressCommissions(address string, active bool) ([]*pb.Commission, error)
}

func NewAccountBalanceService(
	repo IAccountBalanceRepository,
	srv_a *ServiceAccount,
	srv_t *ServiceTransaction,
	bus *events.Bus,
) *ServiceAccountBalance {
	return &ServiceAccountBalance{
		repo:               repo,
		serviceAccount:     srv_a,
		serviceTransaction: srv_t,
		bus:                bus,
	}
}

type ServiceAccountBalance struct {
	repo               IAccountBalanceRepository
	serviceAccount     *ServiceAccount
	serviceTransaction *ServiceTransaction
	bus                *events.Bus
}

func (s *ServiceAccountBalance) AddEventBus(bus *events.Bus) {
	s.bus = bus
	bus.Subscribe(s,
		&pb.Event{Type: &pb.Event_TxConfirm{}},
		&pb.Event{Type: &pb.Event_NewOrder{}},
		&pb.Event{Type: &pb.Event_NewContract{}},
		&pb.Event{Type: &pb.Event_OrderChange{}},
		&pb.Event{Type: &pb.Event_ContractChange{}},
		&pb.Event{Type: &pb.Event_NewWithdrawalOrder{}},
		&pb.Event{Type: &pb.Event_WithdrawalPerformed{}},
	)
}

func (s *ServiceAccountBalance) Update(event *pb.Event) {

	switch event.Type.(type) {
	// Баланс аккаунта изменен
	case *pb.Event_TxConfirm:
		tx := event.GetTxConfirm().Tx
		go s.calcBalanceAndNotify(tx.FromAddress)
		go s.calcBalanceAndNotify(tx.ToAddress)
	case *pb.Event_NewOrder:
		o := event.GetNewOrder().Order
		go s.calcBalanceAndNotify(o.SendingAddress)
	case *pb.Event_OrderChange:
		o := event.GetOrderChange().Order
		go s.calcBalanceAndNotify(o.SendingAddress)
	case *pb.Event_NewContract:
		c := event.GetNewContract().Contract
		go s.calcBalanceAndNotify(c.BuyerSendAddress)
	case *pb.Event_ContractChange:
		c := event.GetContractChange().Contract
		go s.calcBalanceAndNotify(c.BuyerSendAddress)
		go s.calcBalanceAndNotify(c.SellerSendAddress)
	case *pb.Event_NewWithdrawalOrder:
		wo := event.GetNewWithdrawalOrder().WithdrawalOrder
		go s.calcBalanceAndNotify(wo.SendingAddress)
	case *pb.Event_WithdrawalPerformed:
		wo := event.GetWithdrawalPerformed().WithdrawalOrder
		go s.calcBalanceAndNotify(wo.SendingAddress)
	}
}

func (s *ServiceAccountBalance) Notify(event *pb.Event) {
	log.Printf("Account Balance change notify.")
	s.bus.NewEvent(event)
}

func (s *ServiceAccountBalance) calcBalanceAndNotify(addr string) {
	if len(addr) > 0 {
		if rb, err := s.GetBalance(context.Background(), &pb.Query_Account{Address: addr}); err == nil {
			s.Notify(&pb.Event{Type: &pb.Event_BalanceChange{BalanceChange: &pb.EventAccountBalanceChange{Balance: rb.Object}}})
		}
	}
}

//GetBalance является непреодолимой преградой
//(на самом деле нет, ещё важно что бы в CreateOrder и CreateOrder это учитывалось)
//на пути мошенников всех мастей стремящихся создать
//order или contract с sendAddress на котором
//не достаточно баланса available.
func (s *ServiceAccountBalance) GetBalance(ctx context.Context, q *pb.Query_Account) (*pb.Response_AccountBalance, error) {
	a_resp, err := s.serviceAccount.GetAccount(ctx, q)
	if err != nil || a_resp == nil || a_resp.Object == nil {
		return nil, fmt.Errorf("GetBalance: %s", err)
	}

	//Получаю все транзы по счету
	txs, err := s.serviceTransaction.GetTx(ctx, &pb.Query_Tx{Address: a_resp.Object.Address})
	if err != nil {
		return nil, fmt.Errorf("GetBalance: %s", err)
	}

	var balance pb.AccountBalance
	var available float64

	balance.AccountId = a_resp.Object.Id
	balance.AccountOwnerId = a_resp.Object.OwnerId

	// Считаем available
	for _, tx := range txs.Items {
		if tx.Status == pb.TxStatus_CONFIRMED {
			if tx.ToAddress == a_resp.Object.Address {
				available += tx.Amount
			} else {
				available -= tx.Amount
			}
		}
	}

	//Все активные orders
	//orders
	ords, err := s.repo.GetOrders(&pb.Query_Order{
		SendingAddress: a_resp.Object.Address,
		ReceiveAddress: a_resp.Object.Address,
		Active:         true,
	})
	if err != nil {
		return nil, fmt.Errorf("GetBalance: %s ", err)
	}
	//Блокируем средства по активным ордерам
	if len(ords) > 0 {
		for _, o := range ords {

			locked := pb.AccountBalance_LockedReason{
				Reason:  pb.AccountBalance_LockedReason_REASON_ORDER_LOCKED,
				OrderId: o.Id,
			}

			if o.SendingAddress == a_resp.Object.Address {
				locked.Amount = o.Available
			}
			balance.LockedReasons = append(balance.LockedReasons, &locked)
		}
	}

	//Все активные контракты
	//contracts
	cts, err := s.repo.GetContracts(&pb.Query_Contract{
		BuyerSendAddress:  a_resp.Object.Address,
		SellerSendAddress: a_resp.Object.Address,
		Active:            true,
	})
	if err != nil {
		return nil, fmt.Errorf("GetBalance: %s ", err)
	}

	//Блокируем средства по активным контрактам
	if len(cts) > 0 {
		for _, ct := range cts {

			locked := pb.AccountBalance_LockedReason{
				Reason:     pb.AccountBalance_LockedReason_REASON_CONTRACT_LOCKED,
				ContractId: ct.Id,
			}

			if ct.SellerSendAddress == a_resp.Object.Address {
				locked.Amount = ct.Available
			} else if ct.BuyerSendAddress == a_resp.Object.Address {
				locked.Amount = ct.Cost
			}

			balance.LockedReasons = append(balance.LockedReasons, &locked)
		}
	}

	//Блокировка по комиссиям
	cmss, err := s.repo.FindAddressCommissions(a_resp.Object.Address, true)
	if len(cmss) > 0 {
		for _, cm := range cmss {
			if cm.SendingAddress == a_resp.Object.Address {
				locked := pb.AccountBalance_LockedReason{
					Reason: pb.AccountBalance_LockedReason_REASON_COMMISSION_LOCKED,
					Amount: cm.Remainder,
				}
				balance.LockedReasons = append(balance.LockedReasons, &locked)
			}
		}
	}

	//
	wos, err := s.repo.GetWithdrawal(&pb.Query_Withdrawal{
		SendingAddress: q.Address,
		Status:         &pb.DealStatus{Status: pb.DealStatus_CREATED},
	})
	if err != nil {
		return nil, fmt.Errorf("GetBalance: %s ", err)
	}

	if len(wos) > 0 {
		for _, wo := range wos {
			if wo.SendingAddress == a_resp.Object.Address {
				locked := pb.AccountBalance_LockedReason{
					Reason:            pb.AccountBalance_LockedReason_REASON_WITHDRAWAL_LOCKED,
					Amount:            wo.Amount,
					WithdrawalOrderId: wo.Id,
				}
				balance.LockedReasons = append(balance.LockedReasons, &locked)
			}
		}
	}

	var lockedAmount float64
	for _, locked := range balance.LockedReasons {
		lockedAmount += locked.Amount
	}

	balance.Locked = lockedAmount
	balance.Available = available - balance.Locked

	return &pb.Response_AccountBalance{
		Object:      &balance,
		QueryStatus: pb.QueryStatus_Query_Success,
	}, nil
}
