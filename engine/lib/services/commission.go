package services

import (
	pb "engine/lib/structs"
	"fmt"
	"golang.org/x/net/context"
	"log"
	"os"
)

//db-interface
type ICommissionRepository interface {
	FindAddressCommissions(address string, active bool) ([]*pb.Commission, error)
}

//service
type ServiceCommission struct {
	serviceAccount        *ServiceAccount
	serviceAccountBalance *ServiceAccountBalance
	repo                  ICommissionRepository
	systemUserId          string
	rule                  float64
}

//constructor
func NewCommissionService(repo ICommissionRepository, srv_a *ServiceAccount, srv_ab *ServiceAccountBalance, systemUserId string, rule float64) *ServiceCommission {
	if len(systemUserId) == 0 {
		systemUserId = os.Getenv("SYSTEM_USER_ID")
	}
	if len(systemUserId) == 0 {
		log.Printf("Warning! NewCommissionService not set systemUserId")
	} else {
		log.Printf("Comission service starting with SystemUserUid: %s", systemUserId)
	}

	if rule <= 0 {
		rule = 0.01
	}

	log.Printf("Commission is %f ", rule*100)

	return &ServiceCommission{
		serviceAccount:        srv_a,
		serviceAccountBalance: srv_ab,
		repo:                  repo,
		systemUserId:          systemUserId,
		rule:                  rule,
	}
}

func (s *ServiceCommission) SetSystemUserId(uid string) {
	s.systemUserId = uid
}

// ServiceCommission Calc комиссия рассчитаная от суммы получаемой продавцом и отправляемой покупателем
func (s *ServiceCommission) Calc(ctx context.Context, req *pb.Query_CalculateCommission) (*pb.Commission, error) {
	if req.User == nil {
		return nil, fmt.Errorf("ServiceCommission: Calc: user should not be nil")
	}
	if req.Order == nil {
		return nil, fmt.Errorf("ServiceCommission: Calc: order should not be nil")
	}

	cm := &pb.Commission{
		Currency: req.Order.BuyCurrencySymbol,
		Type:     &pb.Commission_TypeStatus{},
	}

	u := req.User

	o := req.Order
	//rule := float64(0)
	rule := 0.01

	var value float64

	if req.Contract != nil {
		// Расчитываем для контракта
		c := req.Contract
		if len(o.Id) > 0 {
			cm.Type.OrderId = o.Id
		}

		// Продавец. Считает от комиссии контракта
		if u.Id == c.SellerId && o.Commission != nil {
			cm.Type.Status = pb.Commission_TypeStatus_CONTRACT_SELLER
			// Берем столько от комиссии ордера сколько от него составляет контракт
			value = o.Commission.Amount * (c.Amount / o.Amount)
		} else { // Покупатель. Считает процент от суммы которую он собирается потратить
			cm.Type.Status = pb.Commission_TypeStatus_CONTRACT_BUYER
			value = c.Amount * rule
			cm.Currency = o.SellCurrencySymbol
		}

		//value =
	} else {
		cm.Type.Status = pb.Commission_TypeStatus_ORDER
		value = (o.Amount * o.Price) * rule
	}

	cm.Amount = value

	log.Printf("ServiceCommission: Calc: commission calculated: %f %s %s", cm.Amount, cm.Currency, cm.Type.Status.String())

	return cm, nil
}

// ServiceCommission Init
func (s *ServiceCommission) Init(ctx context.Context, req *pb.Query_CalculateCommission) (*pb.Commission, error) {
	cm, err := s.Calc(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("Init: %s", err)
	}

	// Комиссию с покупателя берем с его счета получения
	if cm.Type.Status == pb.Commission_TypeStatus_CONTRACT_SELLER {
		cm.SendingAddress = req.Order.ReceiveAddress
	}

	// Комиссию с продавца берем с его счета получения
	if cm.Type.Status == pb.Commission_TypeStatus_CONTRACT_BUYER {
		cm.SendingAddress = req.Contract.BuyerReceiveAddress
	}

	// Находим счета получения комиссий
	a_resp, err := s.serviceAccount.GetAccount(ctx, &pb.Query_Account{ParticipantId: s.systemUserId, CurrencySymbol: cm.Currency})
	if err != nil || a_resp == nil || a_resp.Object == nil {
		return nil, fmt.Errorf("ServiceCommission: Init: can't find receive acc. Data: { UserID: %s CommissionCurrency: %s }, %s Error: %s", s.systemUserId, cm.Currency, err)
	}

	cm.ReceiveAddress = a_resp.Object.Address

	if (cm.Type.Status == pb.Commission_TypeStatus_CONTRACT_SELLER || cm.Type.Status == pb.Commission_TypeStatus_CONTRACT_BUYER) && (len(cm.SendingAddress) == 0 || len(cm.ReceiveAddress) == 0) {
		return nil, fmt.Errorf("ServiceCommission: Init: receive or send address not set. CM: \n %v", cm)
	}

	log.Printf("ServiceCommission: Init: commission calculated: %f %s %s", cm.Amount, cm.Currency, cm.Type.Status)

	return &pb.Commission{
		Currency:       cm.Currency,
		SendingAddress: cm.SendingAddress,
		ReceiveAddress: cm.ReceiveAddress,
		Amount:         cm.Amount,
		Remainder:      cm.Amount,
	}, nil
}

func (s *ServiceCommission) OldCalc(ctx context.Context, req *pb.Query_CalculateCommission) (*pb.Commission, error) {
	if req.User == nil {
		return nil, fmt.Errorf("ServiceCommission: Calc: participant should not be nil")
	}
	if req.Order == nil {
		return nil, fmt.Errorf("ServiceCommission: Calc: order should not be nil")
	}

	cm := &pb.Commission{
		Currency: req.Order.BuyCurrencySymbol,
		Type:     &pb.Commission_TypeStatus{},
	}

	u := req.User

	o := req.Order
	rule := float64(0)
	//rule := 0.02

	var value float64

	if req.Contract != nil {
		// Расчитываем для контракта
		c := req.Contract
		if len(o.Id) > 0 {
			cm.Type.OrderId = o.Id
		}

		// Seller
		if u.Id == c.SellerId {
			value = o.Commission.Amount * (c.Amount / o.Amount)
			cm.Type.Status = pb.Commission_TypeStatus_CONTRACT_SELLER
		} else {
			cm.Type.Status = pb.Commission_TypeStatus_CONTRACT_BUYER
			// Buyer
			value = (c.Amount * c.Price) * rule
		}

	} else {
		cm.Type.Status = pb.Commission_TypeStatus_ORDER
		// Расчитываем для ордера
		value = o.Amount * rule
	}
	cm.Amount = value

	log.Printf("ServiceCommission: Calc: commission calculated: %f %s %s", cm.Amount, cm.Currency, cm.Type.Status.String())

	return cm, nil
}

// ServiceCommission OldInit Инициализация комиссии
// Функция использует константную валюту в которой создастся структура комиссии
// Функция так же проверяет наличие средств на счету участника
func (s *ServiceCommission) OldInit(ctx context.Context, req *pb.Query_CalculateCommission) (*pb.Commission, error) {
	cm, err := s.OldCalc(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("Init: %s", err)
	}

	u := req.User
	a_resp, err := s.serviceAccount.GetAccount(ctx, &pb.Query_Account{ParticipantId: u.Id, CurrencySymbol: cm.Currency})
	if err != nil {
		return nil, fmt.Errorf("ServiceCommission: Init: %s", err)
	}

	a := a_resp.Object
	if ab_resp, err := s.serviceAccountBalance.GetBalance(ctx, &pb.Query_Account{AccountId: a.Id}); err == nil {
		blc := ab_resp.Object.Available
		//if ab_resp.Object.LockedReasons != nil {
		//	if cm.Type.Status == pb.Commission_TypeStatus_CONTRACT_SELLER {
		//		for _,lr := range  ab_resp.Object.LockedReasons {
		//			log.Printf("Reason finded %#v", lr)
		//		}
		//	}
		//}

		if cm.Amount > blc {
			return nil, fmt.Errorf("ServiceCommission: Init: not enought balance in sending account. Info: Cm: %v \n Blc_Resp: %#v ", cm, ab_resp.Object.LockedReasons)
		}
	} else {
		return nil, fmt.Errorf("ServiceCommission: Init: %s", err)
	}

	log.Printf("ServiceCommission: Init: commission calculated: %f %s %s", cm.Amount, cm.Currency, cm.Type.Status)

	return &pb.Commission{
		Currency:       cm.Currency,
		SendingAddress: a.Address,
		//ReceiveAddress: ReceiveSystemAddress,
		Amount:    cm.Amount,
		Remainder: cm.Amount,
	}, nil
}
