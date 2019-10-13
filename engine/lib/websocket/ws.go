package websocket

import (
	"context"
	"engine/lib/helper"
	"engine/lib/services/events"
	"engine/lib/structs"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin/json"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"gopkg.in/olahol/melody.v1"
	"log"
	"net/http"
	"sync"
)

var (
	mu = &sync.Mutex{}
)

type userid string

type WS struct {
	m             *melody.Melody
	mAdmin        *melody.Melody
	bus           *events.Bus
	userSessions  map[userid]*melody.Session
	adminSessions map[userid]*melody.Session
}

func (ws *WS) AddEventBus(bus *events.Bus) {
	// Подписка на события
	bus.Subscribe(ws,
		//&structs.Event{Type: &structs.Event_NewCurrency{}},
		&structs.Event{Type: &structs.Event_CurrencyActivated{}},
		&structs.Event{Type: &structs.Event_CurrencyDeactivated{}},
		&structs.Event{Type: &structs.Event_PaySystemAdded{}},
		&structs.Event{Type: &structs.Event_NewUser{}},
		&structs.Event{Type: &structs.Event_NewAccount{}},
		&structs.Event{Type: &structs.Event_NewTransaction{}},
		&structs.Event{Type: &structs.Event_BalanceChange{}},
		&structs.Event{Type: &structs.Event_TxConfirm{}},
		&structs.Event{Type: &structs.Event_NewOrder{}},
		&structs.Event{Type: &structs.Event_NewContract{}},
		&structs.Event{Type: &structs.Event_OrderChange{}},
		&structs.Event{Type: &structs.Event_OrderPerformed{}},
		&structs.Event{Type: &structs.Event_ContractChange{}},
		&structs.Event{Type: &structs.Event_ContractPerformed{}},
		&structs.Event{Type: &structs.Event_OrderCanceled{}},
		&structs.Event{Type: &structs.Event_NewWithdrawalOrder{}},
		&structs.Event{Type: &structs.Event_WithdrawalPerformed{}},
		&structs.Event{Type: &structs.Event_TxProccessUpdate{}},
	)

	ws.bus = bus
}

//type isWsMessage_Type interface {
//	isWsMessage_Type()
//}
//
//// табличный метод для маршалинга сообщений по типу эвента
//var marchalingTable map[isWsMessage_Type] func(interface{})[]byte

// Пришел эвент. Нужно разослать по сессиям
func (ws *WS) Update(event *structs.Event) {
	mu.Lock()
	defer mu.Unlock()

	var uid userid
	wsMsg := new(structs.WsMessage)

	wsMsg.Namespace = "objects"

	onlyAdmin := false
	broadcast := false

	switch event.Type.(type) {
	case *structs.Event_NewAccount:
		wsMsg.Action = "Event_NewAccount"
		uid = userid(event.GetNewAccount().Account.OwnerId)
		wsMsg.Data = &structs.WsMessage_Account{Account: event.GetNewAccount().Account}
	case *structs.Event_OrderPerformed, *structs.Event_OrderCanceled:
		broadcast = true
	case *structs.Event_NewOrder:
		broadcast = true
		wsMsg.Action = "Event_NewOrder"
		wsMsg.Data = &structs.WsMessage_Order{Order: event.GetNewOrder().Order}
	case *structs.Event_NewContract:
		wsMsg.Action = "Event_NewContract"
		wsMsg.Data = &structs.WsMessage_Contract{Contract: event.GetNewContract().Contract}
	case *structs.Event_OrderChange:
		broadcast = true
		wsMsg.Action = "Event_OrderChange"
		wsMsg.Data = &structs.WsMessage_Order{Order: event.GetOrderChange().Order}
	case *structs.Event_BalanceChange:
		wsMsg.Action = "Event_BalanceChange"
		b := event.GetBalanceChange().Balance
		uid = userid(b.AccountOwnerId)
		wsMsg.Data = &structs.WsMessage_AccountBalance{AccountBalance: b}
		// Только для администраторов
	case *structs.Event_NewUser:
		onlyAdmin = true
		wsMsg.Action = "Event_NewUser"
		u := event.GetNewUser().User
		wsMsg.Data = &structs.WsMessage_User{User: u}
	case *structs.Event_NewTransaction:
		onlyAdmin = true
		wsMsg.Action = "Event_NewTransaction"
		tx := event.GetNewTransaction().Tx
		wsMsg.Data = &structs.WsMessage_Tx{Tx: tx}
	case *structs.Event_TxConfirm:
		onlyAdmin = true
		wsMsg.Action = "Event_TxConfirm"
		tx := event.GetTxConfirm().Tx
		wsMsg.Data = &structs.WsMessage_Tx{Tx: tx}
	case *structs.Event_NewWithdrawalOrder:
		wsMsg.Action = "Event_NewWithdrawalOrder"
		wo := event.GetNewWithdrawalOrder().WithdrawalOrder
		wsMsg.Data = &structs.WsMessage_WithdrawalOrder{WithdrawalOrder: wo}
	case *structs.Event_WithdrawalPerformed:
		wsMsg.Action = "Event_WithdrawalPerformed"
		wo := event.GetWithdrawalPerformed().WithdrawalOrder
		wsMsg.Data = &structs.WsMessage_WithdrawalOrder{WithdrawalOrder: wo}
	case *structs.Event_TxProccessUpdate:
		onlyAdmin = true
		wsMsg.Action = "Event_TxProccessUpdate"
		tx := event.GetTxProccessUpdate().Tx
		wsMsg.Data = &structs.WsMessage_Tx{Tx: tx}
	case *structs.Event_PaySystemAdded:
		onlyAdmin = true
		wsMsg.Action = "Event_PaySystemAdded"
		wsMsg.Data = &structs.WsMessage_Currency{Currency: event.GetPaySystemAdded().Currency}
	case *structs.Event_CurrencyActivated:
		onlyAdmin = true
		wsMsg.Action = "Event_CurrencyActivated"
		wsMsg.Data = &structs.WsMessage_Currency{Currency: event.GetCurrencyActivated().Currency}
	case *structs.Event_CurrencyDeactivated:
		onlyAdmin = true
		wsMsg.Action = "Event_CurrencyDeactivated"
		wsMsg.Data = &structs.WsMessage_Currency{Currency: event.GetCurrencyDeactivated().Currency}
	}

	// Присваивать на основе интерфейсов OneOf из protobuf
	byteMsg, _ := json.Marshal(wsMsg)

	if !onlyAdmin {
		if broadcast {
			ws.m.Broadcast(byteMsg)
		}

		if s, ok := ws.userSessions[uid]; ok {
			s.Write(byteMsg)
		}
	}

	if len(ws.adminSessions) > 0 {
		ws.mAdmin.Broadcast(byteMsg)
	}

}

var (
	srvc_auth   structs.ServiceAuthClient
	srvc_c      structs.ServiceCurrencyClient
	srvc_user   structs.ServiceUserClient
	srvc_acc    structs.ServiceAccountClient
	srvc_accbal structs.ServiceAccountBalanceClient
	srvc_tx     structs.ServiceTransactionClient
	srvc_ord    structs.ServiceOrderClient
)

func InitWS(grpcConn *grpc.ClientConn, mux *http.ServeMux, bus *events.Bus) {
	m := melody.New()
	mAdmin := melody.New()
	lock := new(sync.Mutex)
	ws := WS{
		m:             m,
		mAdmin:        mAdmin,
		userSessions:  make(map[userid]*melody.Session),
		adminSessions: make(map[userid]*melody.Session),
	}

	ws.AddEventBus(bus)

	srvc_auth = structs.NewServiceAuthClient(grpcConn)
	srvc_c = structs.NewServiceCurrencyClient(grpcConn)
	srvc_user = structs.NewServiceUserClient(grpcConn)
	srvc_acc = structs.NewServiceAccountClient(grpcConn)
	srvc_accbal = structs.NewServiceAccountBalanceClient(grpcConn)
	srvc_ord = structs.NewServiceOrderClient(grpcConn)
	srvc_tx = structs.NewServiceTransactionClient(grpcConn)

	mux.Handle("/ws", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenStr := helper.GetTokenStringFromReq(r)
		_, err := srvc_auth.DecodeSession(r.Context(), &structs.Session{Token: tokenStr})
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}
		log.Println("WS: connected")
		m.HandleRequest(w, r)
	}))

	m.HandleConnect(func(s *melody.Session) {
		tokenStr := helper.GetTokenStringFromReq(s.Request)
		u, err := srvc_auth.DecodeSession(context.Background(), &structs.Session{Token: tokenStr})
		if err != nil {
			s.CloseWithMsg([]byte(err.Error()))
			return
		}

		go ws.initUserConnection(lock, u, s)
	})

	mux.Handle("/admin/ws", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenStr := helper.GetTokenStringFromReq(r)
		_, err := srvc_auth.DecodeSession(r.Context(), &structs.Session{Token: tokenStr})
		if err != nil {
			log.Printf("Error: /admin/ws %s", err)
			w.WriteHeader(http.StatusConflict)
			w.Write([]byte(err.Error()))
			return
		}

		mAdmin.HandleRequest(w, r)
	}))

	mAdmin.HandleConnect(func(s *melody.Session) {
		tokenStr := helper.GetTokenStringFromReq(s.Request)
		u, err := srvc_auth.DecodeSession(context.Background(), &structs.Session{Token: tokenStr})
		if err != nil {
			log.Printf("Admin ws handle. Error decode session")
			s.CloseWithMsg([]byte(err.Error()))
			return
		}
		if u.Status == structs.UserStatus_ADMINISTRATOR {
			log.Printf("Admin ws handle. Init connection")
			go ws.initAdminConnection(lock, u, s)
		} else {
			s.Write([]byte("you shell not pass"))
			s.Close()
		}

	})

}

func (ws *WS) initUserConnection(mu *sync.Mutex, u *structs.User, s *melody.Session) {

	uidStr := u.Id

	log.Printf("User %s connected to ws", uidStr)

	ctx := context.WithValue(context.Background(), "pid", uidStr)
	ctx = metadata.AppendToOutgoingContext(ctx, "pid", uidStr)

	uid := userid(uidStr)

	mu.Lock()
	ws.userSessions[uid] = s
	mu.Unlock()

	// в отдельную горутину и убрать инициализацию куда-нибудь
	go initUserWsAccounts(ctx, s)
	// Отправляем все ордера
	go initWsOrders(ctx, s)

	go initWithdrawal(ctx, s)
}

func (ws *WS) initAdminConnection(mu *sync.Mutex, u *structs.User, s *melody.Session) {
	uidStr := u.Id

	log.Printf("Admin %s connected to ws", uidStr)

	ctx := context.WithValue(context.Background(), "admin-id", uidStr)
	ctx = metadata.AppendToOutgoingContext(ctx, "admin-id", uidStr)

	uid := userid(uidStr)

	mu.Lock()
	ws.adminSessions[uid] = s
	mu.Unlock()

	go initWsCurrency(ctx, s)

	// Все пользователи
	go initWsUsers(ctx, s)

	// Все ордера
	go initWsOrders(ctx, s)

	// Все контракты
	go initWsContracts(ctx, s)

	// Все счета
	go initWsAccounts(ctx, s, true)

	// Все транзакции
	go initWsTxs(ctx, s)

	go initWithdrawal(ctx, s)
}

func initWsCurrency(ctx context.Context, s *melody.Session) {
	wsCurrResult := new(structs.WsCurrencies)
	wsCurrResult.Action = "InitCurrencies"
	wsCurrResult.Namespace = "objects"

	if resp_cs, err := srvc_c.GetCurrency(ctx, &structs.Query_Currency{}); err == nil {
		wsCurrResult.Currencies = resp_cs.Items

		if ms, err := json.Marshal(wsCurrResult); err == nil {
			s.Write(ms)
		} else {
			fmt.Println("initWsCurrency: err ", err)
		}
	}

}

func initUserWsAccounts(ctx context.Context, s *melody.Session) {
	initWsAccounts(ctx, s, false)
}

func initWsAccounts(ctx context.Context, s *melody.Session, admin bool) {
	wsAccsResult := new(structs.WsAccounts)

	wsAccsResult.Action = "InitAccounts"
	wsAccsResult.Namespace = "objects"

	var ra *structs.Response_Account
	var err error
	if admin {
		ra, err = srvc_acc.GetAllAccount(ctx, &structs.Empty{})
	} else {
		ra, err = srvc_acc.GetAllUserAccount(ctx, &structs.Empty{})
	}

	if err != nil {
		log.Printf("InitWsAccounts: %s. Admin: %b", err, admin)
		return
	}

	for _, a := range ra.Items {
		wsAcc := new(structs.WsAccount)
		wsAcc.Account = a

		rab, _ := srvc_accbal.GetBalance(ctx, &structs.Query_Account{Address: a.Address})
		wsAcc.Balance = rab.Object

		wsAccsResult.Accounts = append(wsAccsResult.Accounts, wsAcc)
	}

	if msgAcc, err := json.Marshal(wsAccsResult); err == nil {
		s.Write([]byte(msgAcc))
	}

}

func initWsUsers(ctx context.Context, s *melody.Session) {
	if rus, err := srvc_user.GetAllUsers(ctx, &structs.Empty{}); err == nil {
		msg := structs.WsUsers{Action: "InitUsers", Namespace: "objects", Users: rus.Items}
		if msgUsr, err := json.Marshal(msg); err == nil {
			s.Write([]byte(msgUsr))
		}
	}
}

func initWsOrders(ctx context.Context, s *melody.Session) {
	ro, err := srvc_ord.GetOrders(ctx, &structs.Query_Order{})
	if err != nil {
		log.Printf("initWsOrders: %s", err)
		return
	}

	wsMessage := structs.WsOrders{Namespace: "objects", Action: "InitOrders", Orders: ro.Items}
	if msgOrd, err := json.Marshal(wsMessage); err == nil {
		s.Write([]byte(msgOrd))
	}
}

func initWsContracts(ctx context.Context, s *melody.Session) {
	rcs, err := srvc_ord.GetContracts(ctx, &structs.Query_Contract{})
	if err != nil {
		log.Printf("initWsContracts: %s", err)
		return
	}

	wsMessage := structs.WsContracts{Namespace: "objects", Action: "InitContracts", Contracts: rcs.Items}
	if msgOrd, err := json.Marshal(wsMessage); err == nil {
		s.Write([]byte(msgOrd))
	}
}

func initWsTxs(ctx context.Context, s *melody.Session) {
	rtxs, err := srvc_tx.GetAllTxs(ctx, &structs.Empty{})
	if err != nil {
		log.Printf("initWsTxs: %s", err)
		return
	}

	wsMessage := structs.WsTxs{Namespace: "objects", Action: "InitTxs", Txs: rtxs.Items}
	if msgOrd, err := json.Marshal(wsMessage); err == nil {
		s.Write([]byte(msgOrd))
	}
}

func initWithdrawal(ctx context.Context, s *melody.Session) {
	rwos, err := srvc_ord.GetWithdrawalOrders(ctx, &structs.Query_Withdrawal{Status: &structs.DealStatus{Status: structs.DealStatus_CREATED}})
	if err != nil {
		log.Printf("initWithdrawal: %s", err)
		return
	}

	wsMessage := structs.WsWithdrawals{Namespace: "objects", Action: "InitWithdrawal", WithdrawalOrders: rwos.Items}
	if msgOrd, err := json.Marshal(wsMessage); err == nil {
		s.Write([]byte(msgOrd))
	}
}

func getUserIdFromToken(tokenStr string) string {
	token, _ := jwt.Parse(tokenStr, nil)
	claims := token.Claims.(jwt.MapClaims)
	if uid, ok := claims["sub"]; ok {
		return uid.(string)
	}
	return ""
}
