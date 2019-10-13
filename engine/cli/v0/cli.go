package v0

import (
	"encoding/json"
	"fmt"
	"github.com/urfave/cli"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"io/ioutil"
	"log"
	"math/rand"
	pb "mxp/mxp-protobuf/pkg"
	"os"
	"reflect"
	"time"
)

var version string

var rpc string
var conn *grpc.ClientConn

func parseFile(file string) error {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}

	type jobStruct struct {
		Method, Payload string
	}
	type serviceJobs struct {
		Service string
		Jobs    []jobStruct
	}
	var file_session []serviceJobs

	if err := json.Unmarshal(data, &file_session); err != nil {
		return fmt.Errorf("Unmarshaling error: \n\t%s", err)
	}

	fmt.Printf("%v", file_session)
	for _, service_jobs := range file_session {
		var client interface{}
		switch service_jobs.Service {
		case "currency":
			client = pb.NewServiceCurrencyClient(conn)
			break
		case "participant":
			client = pb.NewServiceParticipantClient(conn)
			break
		case "account":
			client = pb.NewServiceAccountClient(conn)
			break
		case "transaction":
			client = pb.NewServiceTransactionClient(conn)
			break
		case "order":
			client = pb.NewServiceOrderClient(conn)
			break
		default:
			fmt.Printf("Внимание! Сервис %s не найден\n", service_jobs.Service)
			continue
		}

		fmt.Println("----------------------")
		fmt.Printf("Инициализация запросов к сервису %s\n", service_jobs.Service)
		fmt.Println("----------------------")

		for _, j := range service_jobs.Jobs {
			service_call(client, j.Method, j.Payload)
		}
	}

	return nil
}

//@author wizardry
//payload - аргументы
func service_call(client interface{}, proc, payload string) {

	fmt.Printf("Client Type: %s \n", reflect.TypeOf(client))

	method := reflect.ValueOf(client).MethodByName(proc)

	//Аргументы
	in := make([]reflect.Value, method.Type().NumIn())

	ctx := context.Background()
	in[0] = reflect.New(reflect.TypeOf(ctx)).Elem()

	for i := 0; i < method.Type().NumIn(); i++ {
		t := method.Type().In(i)
		if t.String() != "context.Context" && t.String() != "[]grpc.CallOption" {
			objType := t.Elem()
			obj := reflect.New(objType).Interface()
			if len(payload) > 0 {
				if err := json.Unmarshal([]byte(payload), &obj); err != nil {
					fmt.Printf("Unmarshaling error %s \n", err)
					return
				}
			}

			in[i] = reflect.ValueOf(obj)
		}
	}

	var opt grpc.EmptyCallOption
	optType := reflect.TypeOf(opt)

	in[2] = reflect.New(optType).Elem()

	response := method.Call(in)

	fmt.Printf("\n------\n")
	fmt.Printf("Request Payload: { %+v }\n", in[1])
	for i, v := range response {
		fmt.Printf("%d: \n", i+1)
		fmt.Printf(" - Response Type: %T \n", v)
		fmt.Printf(" - Response Value: %v \n", v)

	}

}

func initialize(N int) {
	srvc_c := pb.NewServiceCurrencyClient(conn)
	srvc_p := pb.NewServiceParticipantClient(conn)
	srvc_a := pb.NewServiceAccountClient(conn)
	//srvc_ab := pb.NewServiceAccountBalanceClient(conn)
	srvc_t := pb.NewServiceTransactionClient(conn)
	srvc_tp := pb.NewServiceTransactionProcessingClient(conn)
	srvc_o := pb.NewServiceOrderClient(conn)

	ctx := context.Background()

	fmt.Printf("#Получение валют.")
	currResp, err := srvc_c.GetCurrency(ctx, &pb.Query_Currency{})
	if err != nil {
		fmt.Printf("- ошибка инициализации. Не удалось получить валюты: %s\n", err)
		return
	}

	curLen := currResp.ItemsCount

	fmt.Printf("- получено (%d) валют\n", curLen)

	fmt.Printf("#Инициализация (%d) участников:\n", N)

	type partInit struct {
		Participant  *pb.Participant
		Accounts     []*pb.Account
		Orders       []*pb.Order
		BuyContracts []*pb.Contract
	}

	var participants []*partInit
	for i := 0; i < N; i++ {
		name := fmt.Sprintf("test%d", i)
		fmt.Printf(" Имя: %s\n", name)
		fmt.Printf(" #Добавляем в базу: ")
		pResp, err := srvc_p.CreateParticipant(ctx, &pb.Participant{Name: name, Password: "123123"})
		if err != nil {
			fmt.Printf("Ошибка! %s\n", err)
			return
		}
		fmt.Printf("Добавлен. ID: %s\n", pResp.Object.Id)
		pi := partInit{Participant: pResp.Object}

		fmt.Printf(" #Создание и пополнение внутренних валютных счетов: ")
		for _, c := range currResp.Items {
			fmt.Printf(" Валюта (%s): ", c.Symbol)
			accResp, err := srvc_a.CreateAccount(ctx, &pb.Query_CreateAccount{CurrencyId: c.Id, OwnerId: pResp.Object.Id, Type: pb.AccountType_INTERNAL})
			if err != nil {
				fmt.Printf("Ошибка! %s\n %+v", err)
				return
			}
			fmt.Printf("Счет создан. ID: %s \n", accResp.Object.Id)
			pi.Accounts = append(pi.Accounts, accResp.Object)
		}
		fmt.Println("Все счета созданы.")
		fmt.Println("#Создание FUND транзакций для пополения счетов:")
		for _, a := range pi.Accounts {
			fmt.Printf(" #Пополнение счета (%s): \n", a.Id)
			_tx_resp, err := srvc_tp.UnderstandingRawTx(ctx, &pb.Query_RawTx{FromAddress: "test", ToAddress: a.Address, Amount: 100 * float64(curLen)})
			if err != nil {
				fmt.Printf("Ошибка! %s\n", err)
				return
			}
			if _tx_resp.Object.Reason.Status != pb.TxReason_FUND_TX {
				fmt.Printf("Ошибка! Причина созданной транзакции не пополнение\n", err)
				return
			}
			tx_resp, err := srvc_t.CreateTx(ctx, _tx_resp.Object)
			if err != nil {
				fmt.Printf("Ошибка создания транзакции")
				return
			}
			fmt.Printf(" Создана транзакция. ID: %s\n #Подтверждение: \n", tx_resp.Object.Id)

			if _, err := srvc_tp.ConfirmTx(ctx, &pb.Query_Tx{TxId: tx_resp.Object.Id}); err != nil {
				fmt.Printf("Ошибка! %s\n", err)
				return
			}

			fmt.Printf(" Подтверждено. Счет пополнен.\n")
		}

		fmt.Println(" Все счета участника пополнены.")
		fmt.Println("#Создание ордера о продаже валюты: ")

		var receiveAcc *pb.Account

		for _, a := range pi.Accounts {
			if a.Currency.Symbol == "USD" {
				receiveAcc = a
				break
			}
		}

		for _, a := range pi.Accounts {
			if a.Currency.Symbol == "USD" {
				continue
			}
			fmt.Printf(" #Валюта (%s): ", a.Currency.Symbol)
			var o_query pb.Query_Order
			o_query.OwnerId = pi.Participant.Id
			o_query.Amount = 100
			o_query.Price = 1
			o_query.SellCurrencySymbol = a.Currency.Symbol
			o_query.SendingAddress = a.Address
			o_query.ReceiveAddress = receiveAcc.Address
			o_query.BuyCurrencySymbol = receiveAcc.Currency.Symbol

			o_resp, err := srvc_o.CreateOrder(ctx, &o_query)
			if err != nil {
				fmt.Printf("Ошибка создания ордера! %s\n", err)
				return
			}
			fmt.Printf("Ордер создан. ID %s\n", o_resp.Object.Id)
			pi.Orders = append(pi.Orders, o_resp.Object)
		}

		participants = append(participants, &pi)
		fmt.Printf(" Инициализирован %d-й участник\n", i)
	}
	fmt.Printf("Инициализация участников(%d) завершена\n", len(participants))

	fmt.Printf("#Инициализация торговых контрактов участниками \n")

	for _, pi1 := range participants {
		fmt.Printf(" #Участник %s торгует... \n", pi1.Participant.Name)
		sendAccountAddress := pi1.Orders[0].ReceiveAddress
		for _, pi2 := range participants {
			if pi1.Participant.Id == pi2.Participant.Id {
				continue
			}
			fmt.Printf("с участником %s: \n", pi2.Participant.Name)
			for _, o := range pi2.Orders {
				fmt.Printf(" Рассматривается ордер %d(%s) по цене %d(%s): \n", o.Amount, o.SellCurrencySymbol, o.Price, o.BuyCurrencySymbol)
				amount := o.Amount / float64(curLen)
				fmt.Printf(" %s готов переобрести %d(%s).\n #Создается контракт. ", pi1.Participant.Name, amount, o.SellCurrencySymbol)
				var c_query pb.Query_CreateContract
				c_query.BuyerId = pi1.Participant.Id
				c_query.Amount = amount
				c_query.OrderId = o.Id
				c_query.SendingAddress = sendAccountAddress
				for _, a := range pi1.Accounts {
					if a.Currency.Symbol == o.SellCurrencySymbol {
						c_query.ReceiveAddress = a.Address
						break
					}
				}
				if len(c_query.ReceiveAddress) == 0 {
					fmt.Printf("Ошибка! Не удалось задать адрес получения\n")
					return
				}
				c_resp, err := srvc_o.CreateContract(ctx, &c_query)
				if err != nil {
					fmt.Printf("Ошибка! %s\n", err)
					return
				}
				pi1.BuyContracts = append(pi1.BuyContracts, c_resp.Object)
				fmt.Printf("Создано: ID (%s).\n Количество контрактов у %s изменено (%d)\n", c_resp.Object.Id, pi1.Participant.Name, len(pi1.BuyContracts))
			}
		}
	}
	fmt.Printf("#Инициализация торговых контрактов завершена \n")
	fmt.Printf("#Инициализация транзакций на покупку \n")

	for _, buyer := range participants {
		fmt.Printf("#Инициализация транзакций для ордеров покупки для участника %s. Ордеров (%d) \n", buyer.Participant.Name, len(buyer.BuyContracts))

		for _, c := range buyer.BuyContracts {
			fmt.Printf("-Продавец отправляет средства на счет покупателя:")
			_t_resp, err := srvc_tp.UnderstandingRawTx(ctx, &pb.Query_RawTx{
				FromAddress: c.SellerSendAddress,
				ToAddress:   c.BuyerReceiveAddress,
				Amount:      c.Amount,
			})
			if err != nil {
				fmt.Printf(" Ошибка! %s", err)
				return
			}
			fmt.Printf("Отправлено. TxID: %s.\n", _t_resp.Object.Id)
			if _t_resp.Object.Reason.Status == pb.TxReason_BUYER_CONTRACT_TX || _t_resp.Object.Reason.ContractId != c.Id {
				fmt.Printf("Ошибка! Контракт транзакции определен неверено")
				return
			}
			t_resp, err := srvc_t.CreateTx(ctx, _t_resp.Object)
			if err != nil {
				fmt.Printf("Ошибка сохранения транзакции %s ", err)
			}

			fmt.Printf("#Подтверждение транзакции по ордеру: ")
			if _, err := srvc_tp.ConfirmTx(ctx, &pb.Query_Tx{TxId: t_resp.Object.Id}); err != nil {
				fmt.Printf("Ошибка! %s\n", err)
				return
			}

			fmt.Printf("Подтверждено. Когнтракт исполнен.\n")
		}
	}
	fmt.Println("Все ордера исполнены. Цикл полной инициалзиации завершен. ")

}

func create_participants(n int32) {

}

func fund_acc(address string, amount float64) {

}

func init() {
	rand.Seed(time.Now().UnixNano())
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func AppInit() *App {
	app := App{
		ctx:     context.Background(),
		srvc_c:  pb.NewServiceCurrencyClient(conn),
		srvc_p:  pb.NewServiceParticipantClient(conn),
		srvc_a:  pb.NewServiceAccountClient(conn),
		srvc_ab: pb.NewServiceAccountBalanceClient(conn),
		srvc_t:  pb.NewServiceTransactionClient(conn),
		srvc_o:  pb.NewServiceOrderClient(conn),
		srvc_tp: pb.NewServiceTransactionProcessingClient(conn),
	}
	return &app
}

type App struct {
	ctx     context.Context
	srvc_c  pb.ServiceCurrencyClient
	srvc_p  pb.ServiceParticipantClient
	srvc_a  pb.ServiceAccountClient
	srvc_ab pb.ServiceAccountBalanceClient
	srvc_t  pb.ServiceTransactionClient
	srvc_tp pb.ServiceTransactionProcessingClient
	srvc_o  pb.ServiceOrderClient
	p       *pb.Participant
}

func (app *App) AddCurrency(name, symbol string, t int32, d uint32) error {
	if len(name) == 0 || len(symbol) == 0 {
		return fmt.Errorf("AddCurrency: not enought arguments ")
	}

	_c := &pb.Currency{
		Name: name, Symbol: symbol, Decimal: d,
	}

	if pb.Currency_CurrencyType_value[pb.Currency_FIAT_CURRENCY.String()] == t {
		_c.Type = pb.Currency_FIAT_CURRENCY
	} else {
		_c.Type = pb.Currency_CRYPTO_CURRENCY
	}

	c, err := app.srvc_c.CreateCurrency(app.ctx, &pb.Query_CreateCurrency{Object: _c})
	if err != nil {
		return fmt.Errorf("AddCurrency: %s ", err)
	}
	log.Printf("Currency added: %+v", c.Object.Id)
	return nil
}

func (app *App) ViewCurrency(symbol string) error {
	resp, err := app.srvc_c.GetCurrency(app.ctx, &pb.Query_Currency{Symbol: symbol})
	if err != nil {
		return fmt.Errorf("ViewCurrency: %s ", err)
	}
	fmt.Printf("Response: %+v \n", resp)
	return nil
}

func (app *App) ViewOrder(oid string) error {
	resp, err := app.srvc_o.GetOrder(app.ctx, &pb.Query_Order{Id: oid})
	if err != nil {
		return fmt.Errorf("ViewOrder: %s ", err)
	}
	format, err := json.MarshalIndent(resp.Object, "", "		")
	if err != nil {
		return fmt.Errorf("ViewOrder: %s ", err)
	}
	fmt.Printf("Ордер: %s \n", format)
	return nil
}

func (app *App) ViewContract(cid string) error {
	resp, err := app.srvc_o.GetContract(app.ctx, &pb.Query_Contract{ContractId: cid})
	if err != nil {
		return fmt.Errorf("ViewContract: %s ", err)
	}
	format, err := json.MarshalIndent(resp.Items, "", "		")
	if err != nil {
		return fmt.Errorf("ViewContract: %s ", err)
	}
	fmt.Printf("Контракт: %s \n", format)
	return nil
}

func (app *App) AccountBalance(address string) error {
	resp, err := app.srvc_ab.GetBalance(app.ctx, &pb.Query_Account{Address: address})
	if err != nil {
		return fmt.Errorf("AccountBalance: %s ", err)
	}
	data, _ := json.MarshalIndent(resp, "", "    ")
	fmt.Printf("%s \n", data)
	return nil
}

func (app *App) SignUp(name, pass string) error {
	p, err := app.srvc_p.CreateParticipant(app.ctx, &pb.Participant{Name: name, Password: pass})
	if err != nil {
		return err
	}
	app.p = p.Object
	log.Printf("Зарегистрирован новый участние: Name: %s, Pass: %s \n", p.Object.Name, pass)
	return nil

}

func (app *App) Fund(addr string, amount float64) error {
	resp_raw_tx, err := app.srvc_tp.UnderstandingRawTx(app.ctx, &pb.Query_RawTx{
		FromAddress: RandStringRunes(10),
		ToAddress:   addr,
		Amount:      amount,
	})
	if err != nil {
		return err
	}

	resp_tx, err := app.srvc_t.CreateTx(app.ctx, resp_raw_tx.Object)
	if err != nil {
		return err
	}

	_, err = app.srvc_tp.ConfirmTx(app.ctx, &pb.Query_Tx{TxId: resp_tx.Object.Id})
	if err != nil {
		return err
	}
	log.Printf("Адрес %s пополнен(%f)\n", addr, amount)
	return nil
}

func (app *App) CreateOrder(from, to string, amount, price float64) error {
	if app.p == nil {
		if err := app.SignUp("John "+RandStringRunes(5), RandStringRunes(8)); err != nil {
			return err
		}
	}
	q_ord := &pb.Query_Order{}
	q_ord.OwnerId = app.p.Id

	q_ord.SellCurrencySymbol = from
	q_ord.BuyCurrencySymbol = to
	q_ord.Amount = amount
	q_ord.Price = price

	resp_accs, err := app.srvc_a.GetAccount(app.ctx, &pb.Query_Account{ParticipantId: app.p.Id})
	if err != nil {
		return err
	}
	if resp_accs.ItemsCount == 0 {
		return fmt.Errorf("Не удалось получить счета ")
	}
	for _, a := range resp_accs.Items {
		if a.Currency.Symbol == from {
			q_ord.SendingAddress = a.Address
		}
		if a.Currency.Symbol == to {
			q_ord.ReceiveAddress = a.Address
		}
	}

	//Balance
	resp_sab, err := app.srvc_ab.GetBalance(app.ctx, &pb.Query_Account{Address: q_ord.SendingAddress})
	if err != nil {
		return err
	}
	if resp_sab.Object.Available < amount {
		if err := app.Fund(q_ord.SendingAddress, amount); err != nil {
			return err
		}
	}
	resp_o, err := app.srvc_o.CreateOrder(app.ctx, q_ord)
	if err != nil {
		return fmt.Errorf("Create Order: %s ", err)
	}
	log.Printf("Создан ордер: %s \n", resp_o.Object.Id)
	return nil
}

func (app *App) CreateContract(oid string, amount float64) error {
	resp_o, err := app.srvc_o.GetOrder(app.ctx, &pb.Query_Order{Id: oid})
	if err != nil {
		return fmt.Errorf("BuyOrder: %s ", err)
	}
	if app.p == nil {
		if err := app.SignUp("Patrick "+RandStringRunes(3), RandStringRunes(6)); err != nil {
			return fmt.Errorf("BuyOrder: %s ", err)
		}
	}
	resp_accs, err := app.srvc_a.GetAccount(app.ctx, &pb.Query_Account{ParticipantId: app.p.Id})
	if err != nil {
		return fmt.Errorf("BuyOrder: %s ", err)
	}
	var sa, rs *pb.Account
	for _, a := range resp_accs.Items {
		// Счет для отправки
		if a.Currency.Symbol == resp_o.Object.BuyCurrencySymbol {
			sa = a
		}
		// Счет для получения
		if a.Currency.Symbol == resp_o.Object.SellCurrencySymbol {
			rs = a
		}
	}

	if err := app.Fund(sa.Address, amount*resp_o.Object.Price); err != nil {
		return fmt.Errorf("BuyOrder: %s ", err)
	}

	resp_c, err := app.srvc_o.CreateContract(app.ctx, &pb.Query_CreateContract{
		OrderId:        resp_o.Object.Id,
		BuyerId:        app.p.Id,
		SendingAddress: sa.Address,
		ReceiveAddress: rs.Address,
		Amount:         amount,
	})
	if err != nil {
		return fmt.Errorf("BuyOrder: %s ", err)
	}

	log.Printf("Создан контракт: %s \n", resp_c.Object.Id)

	return nil
}

func (app *App) PayContract(cid string, amount float64) error {
	resp_c, err := app.srvc_o.GetContract(app.ctx, &pb.Query_Contract{ContractId: cid})
	if err != nil {
		return fmt.Errorf("PayContract: %s ", err)
	}
	if resp_c.ItemsCount == 0 {
		return fmt.Errorf("PayContract: Contract not find ")
	}
	c := resp_c.Items[0]

	//if err := app.Fund(c.SellerSendAddress, amount); err != nil {
	//	return fmt.Errorf("PayContract: %s ", err)
	//}

	_resp_tx, err := app.srvc_tp.UnderstandingRawTx(app.ctx, &pb.Query_RawTx{
		FromAddress: c.SellerSendAddress,
		ToAddress:   c.BuyerReceiveAddress,
		Amount:      amount,
	})
	if err != nil {
		return fmt.Errorf("PayContract: %s ", err)
	}

	resp_tx, err := app.srvc_t.CreateTx(app.ctx, _resp_tx.Object)
	if err != nil {
		return fmt.Errorf("PayContract: %s ", err)
	}

	if _, err := app.srvc_tp.ConfirmTx(app.ctx, &pb.Query_Tx{TxId: resp_tx.Object.Id}); err != nil {
		return fmt.Errorf("PayContract: %s ", err)
	}

	log.Printf("Транзакция создана и подтверждена: %s \n", resp_tx.Object.Id)

	return nil
}

func main() {
	fmt.Printf("cli version: %s \n", version)
	// Set up a connection to the server.
	app := cli.NewApp()

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "rpc",
			Value:       "localhost:50051",
			Usage:       "rpc сервер",
			Destination: &rpc,
		},
		cli.StringFlag{
			Name:  "payload",
			Value: "",
			Usage: "Программа возьмет аргументы из JSON",
		},
		cli.StringFlag{
			Name:  "service",
			Value: "currency",
			Usage: "Сервис",
		},
		cli.StringFlag{
			Name:  "proc",
			Value: "GetCurrency",
			Usage: "Процедура",
		},
		cli.StringFlag{
			Name:  "file",
			Value: "",
			Usage: "Загрузить payload из файла",
		},
		cli.StringFlag{
			Name:  "initialize",
			Usage: "Сгенерировать юзеров и провести обменные ордера",
		},

		cli.BoolFlag{
			Name:  "createOrder, co",
			Usage: "Создать обменные ордера",
		},
		cli.StringFlag{
			Name:  "from, f",
			Value: "",
			Usage: "Адрес для отправки",
		},
		cli.StringFlag{
			Name:  "to, t",
			Usage: "Адрес для получения",
		},
		cli.Float64Flag{
			Name:  "amount, a",
			Usage: "Сумма",
		},
		cli.Float64Flag{
			Name:  "price, p",
			Usage: "Цена",
		},

		cli.BoolFlag{
			Name:  "createContract, cc",
			Usage: "Создать контракт",
		},
		cli.StringFlag{
			Name:  "orderId, oid",
			Value: "",
			Usage: "Идентификатор ордера",
		},

		cli.BoolFlag{
			Name:  "createCurrency",
			Usage: "Добавить валюту",
		},
		cli.StringFlag{
			Name:  "name, n",
			Usage: "Название валюты",
		},
		cli.StringFlag{
			Name:  "symbol, s",
			Usage: "Символ валюты",
		},
		cli.UintFlag{
			Name:  "decimal, d",
			Value: 8,
			Usage: "Разрядность валюты",
		},
		cli.IntFlag{
			Name:  "type",
			Value: 0,
			Usage: "Тип валюты. 0 - криптовалюта, 1 - фиат",
		},

		cli.BoolFlag{
			Name:  "viewCurrency, vcur",
			Usage: "Глянуть валюту",
		},

		cli.BoolFlag{
			Name:  "accountBalance, ab",
			Usage: "Вывести баланс счета",
		},
		cli.StringFlag{
			Name:  "address, addr",
			Usage: "Адрес счета",
		},

		cli.BoolFlag{
			Name:  "viewOrder, vo",
			Usage: "Глянуть ордер",
		},
		cli.BoolFlag{
			Name:  "viewContract, vc",
			Usage: "Глянуть контракт",
		},
		cli.StringFlag{
			Name:  "contractId, cid",
			Value: "",
			Usage: "Идентификатор контракта",
		},

		cli.BoolFlag{
			Name:  "payContract, pc",
			Usage: "Заплатить по контракту",
		},
	}

	app.Action = func(c *cli.Context) error {
		if len(rpc) > 0 {
			var connErr error
			conn, connErr = grpc.Dial(rpc, grpc.WithInsecure())
			if connErr != nil {
				log.Fatalf("Did not connect: %v", connErr)
			}
			log.Printf("Connected to: %s ", rpc)
			defer conn.Close()
		} else {
			panic("rpc address not set")
		}

		if filename := c.String("file"); len(filename) > 0 {
			if err := parseFile(filename); err != nil {
				return fmt.Errorf("\n------\nОшибка загрузки файла: \n- filename: %s \n- error:%s\n\n", filename, err)
			}
		}
		if n := c.Int("initialize"); n > 0 {
			initialize(n)
		}

		if c.Bool("createOrder") {
			app := AppInit()
			err := app.CreateOrder(c.String("from"), c.String("to"), c.Float64("amount"), c.Float64("price"))
			if err != nil {
				log.Fatalf("%s ", err)
			}
		}

		if c.Bool("createCurrency") {
			app := AppInit()

			err := app.AddCurrency(c.String("name"),
				c.String("symbol"),
				int32(c.Int("type")),
				uint32(c.Uint("decimal")))
			if err != nil {
				log.Fatalf("Application: %s ", err)
			}
		}

		if c.Bool("viewCurrency") {
			app := AppInit()
			err := app.ViewCurrency(c.String("symbol"))
			if err != nil {
				log.Fatalf("Application: %s ", err)
			}
		}

		if c.Bool("accountBalance") {
			app := AppInit()
			err := app.AccountBalance(c.String("address"))
			if err != nil {
				log.Fatalf("Application: %s ", err)
			}
		}

		if c.Bool("createContract") {
			app := AppInit()
			err := app.CreateContract(c.String("orderId"), c.Float64("amount"))
			if err != nil {
				log.Fatalf("Application: %s ", err)
			}
		}

		if c.Bool("viewOrder") {
			app := AppInit()
			err := app.ViewOrder(c.String("orderId"))
			if err != nil {
				log.Fatalf("Application: %s ", err)
			}
		}

		if c.Bool("viewContract") {
			app := AppInit()
			err := app.ViewContract(c.String("contractId"))
			if err != nil {
				log.Fatalf("Application: %s ", err)
			}
		}

		if c.Bool("payContract") {
			app := AppInit()
			err := app.PayContract(c.String("contractId"), c.Float64("amount"))
			if err != nil {
				log.Fatalf("Application: %s ", err)
			}
		}
		return nil
	}

	app.Run(os.Args)

}
