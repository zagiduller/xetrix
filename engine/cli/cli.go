package main

import (
	xapp "engine/cli/app"
	"engine/lib/helper"
	"fmt"
	"github.com/urfave/cli"
	"log"
	"os"
	"strings"
	"time"
)

//rpc server
var rpc string

// build version
var version string

func createAction(c *cli.Context) error {
	app := xapp.AppInit(rpc)
	defer app.Conn.Close()

	if app.U == nil {
		pname := fmt.Sprintf("John %s", helper.RandStringRunes(5))
		if err := app.SignUp(pname, helper.RandStringRunes(8)); err != nil {
			return err
		}
	}

	time.Sleep(1 * time.Second)

	comPath := strings.Split(c.Command.FullName(), " ")

	entity_name := comPath[0]

	switch entity_name {
	case "order":
		return app.CreateOrder(c.String("from"), c.String("to"), c.Float64("amount"), c.Float64("price"))
		break
	case "contract":
		return app.CreateContract(c.String("orderId"), c.Float64("amount"))
		break
	case "transaction":
		return app.CreateTransaction(c.String("fromAddress"), c.String("toAddress"), c.Float64("amount"))
		break
	}
	return nil
}

func entityAction(c *cli.Context) (err error) {
	app := xapp.AppInit(rpc)
	defer app.Conn.Close()
	var dataView interface{}
	var id string
	if c.NArg() > 0 {
		id = c.Args().Get(0)
	} else {
		id = c.String("id")
	}
	comPath := strings.Split(c.Command.FullName(), " ")

	entity_name := comPath[0]
	switch entity_name {
	case "order":
		dataView, err = app.GetOrder(id)
		if err != nil {
			return fmt.Errorf("ViewContract: %s ", err)
		}
		break
	case "contract":
		dataView, err = app.GetContract(id)
		if err != nil {
			return fmt.Errorf("ViewContract: %s ", err)
		}
		break
	case "account":
		switch c.Command.Name {
		case "balance":
			dataView, err = app.Balance(id)
			if err != nil {
				return fmt.Errorf("ViewContract: %s ", err)
			}
			break
		case "fund":
			err = app.Fund(id, c.Float64("amount"))
			if err != nil {
				return fmt.Errorf("ViewContract: %s ", err)
			}
			break
		}
		if c.Command.Name == "balance" {
			dataView, err = app.Balance(id)
			if err != nil {
				return fmt.Errorf("ViewContract: %s ", err)
			}
		}
		break
	case "transaction":
		switch c.Command.Name {
		case "view":
			dataView, err = app.GetTransaction(id)
			if err != nil {
				return fmt.Errorf("ViewTransaction: %s ", err)
			}
		}
		break
	}
	format, err := helper.ToFormatedJson(dataView)
	if err != nil {
		return fmt.Errorf("View(%s): %s ", c.Command.Name, err)
	}
	fmt.Printf("%s: %s \n", c.Command.Name, format)
	return nil
}

func main() {
	fmt.Printf("cli build version: %s \n", version)
	var err error
	app := cli.NewApp()
	app.Name = "mxp-cli"
	app.Copyright = "Arthur Zagidullin"
	app.Author = "Arthur Zagidullin"
	if len(version) == 0 {
		app.Description = "!!! Некорректная сборка. \n Используйте команду: \n 'go build -ldflags \"-X main.version=`date -u +%Y.%m.%d-%H:%M.%S`\"'"
	} else {
		app.Description = fmt.Sprintf("Консольный интерфейс mxp. Версия сборки: %s", version)
	}
	app.Email = "design.mgn@gmail.com"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "rpc",
			Value:       "localhost:50051",
			Usage:       "rpc сервер",
			Destination: &rpc,
		},
	}
	app.Version = "0.0.1"
	app.Commands = []cli.Command{
		{
			Name:    "currency",
			Aliases: []string{"curr"},
			Subcommands: cli.Commands{
				{
					Name:    "appendContract",
					Usage:   "cli currency appendContract",
					Aliases: []string{"ac"},
					Flags:   []cli.Flag{},
				},
				{
					Name:  "activate",
					Usage: "cli currency activate",
					Flags: []cli.Flag{cli.StringFlag{Name: "symbol, s", Usage: "Символ валюты"}},
					Action: func(c *cli.Context) error {
						app := xapp.AppInit(rpc)
						if err := app.SignInWithRequest(true); err != nil {
							return err
						}
						return app.ActivateCurrency(c.String("symbol"))

					},
				},
				{
					Name:  "deactivate",
					Usage: "cli currency deactivate",
					Flags: []cli.Flag{
						cli.StringFlag{Name: "symbol, s", Usage: "Символ валюты"},
					},
				},
			},
		},
		{
			Name:    "order",
			Aliases: []string{"o"},
			Subcommands: []cli.Command{
				{
					Name:    "create",
					Usage:   "--from --to --amount --price",
					Aliases: []string{"c"},
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:  "from, f",
							Value: "",
							Usage: "Валюта продажи",
						},
						cli.StringFlag{
							Name:  "to, t",
							Usage: "Валюта покупки",
						},
						cli.Float64Flag{
							Name:  "amount, a",
							Usage: "Сумма продажи",
						},
						cli.Float64Flag{
							Name:  "price, p",
							Usage: "Цена за одну единицу amount валюты from",
						},
					},
					Action: createAction,
					Description: "Команда зарегистрирует нового участника, пополнит счет в валюте from на сумму amount " +
						"и отправит команду на создание нового ордера. Возвращает ID созданого ордера, или ошибку в случае неудачи. " +
						"",
				},
				{
					Name:        "view",
					Aliases:     []string{"v"},
					Usage:       "-id",
					Description: "передайте в качестве аргумента идентификатор ордера",
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:  "orderId, id",
							Usage: "Идентификатор ордера",
						},
					},
					Action: entityAction,
				},
				{
					Name:  "cancel",
					Usage: "-id",
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:  "orderId, id",
							Usage: "Идентификатор ордера",
						},
					},
					Action: func(c *cli.Context) error {
						app := xapp.AppInit(rpc)
						defer app.Conn.Close()
						o, err := app.CancelOrder(c.String("orderId"))
						if err != nil {
							return err
						}

						format, _ := helper.ToFormatedJson(o)
						fmt.Printf("%s: %s\n", c.Command.Name, format)
						return nil
					},
				},
				{
					Name:        "list",
					Aliases:     []string{"l"},
					Description: "Список ордеров в системе",
					Action: func(c *cli.Context) error {
						app := xapp.AppInit(rpc)
						defer app.Conn.Close()
						os, err := app.ListOrder()
						if err != nil {
							return err
						}
						format, _ := helper.ToFormatedJson(os)
						fmt.Printf("%s: %s\n", c.Command.Name, format)
						return nil

					},
				},
				{
					Name:        "generate",
					Aliases:     []string{"g"},
					Description: "Автоматическая генерация ордеров",
					Flags: []cli.Flag{
						cli.IntFlag{
							Name:  "count, c",
							Usage: "Количество генерируемых ордеров",
						},
						cli.Float64Flag{
							Name:  "amount, a",
							Usage: "Сумма продажи",
						},
						cli.Float64Flag{
							Name:  "price, p",
							Usage: "Цена за одну единицу amount валюты from",
						},
					},
					Action: func(c *cli.Context) error {
						app := xapp.AppInit(rpc)
						defer app.Conn.Close()
						app.OrderGenerate(c.Int("count"), c.Float64("amount"), c.Float64("price"))
						return nil

					},
				},
				{
					Name:        "createWithdrawal",
					Aliases:     []string{"cw"},
					Description: "Создать заявку на вывод",
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:  "address, addr",
							Usage: "Адрес вывода",
						},
						cli.Float64Flag{
							Name:  "amount, a",
							Usage: "Сумма вывода",
						},
					},
					Action: func(c *cli.Context) error {
						app := xapp.AppInit(rpc)
						defer app.Conn.Close()
						wo, err := app.CreateWithdrawal(c.String("address"), c.Float64("amount"))
						format, _ := helper.ToFormatedJson(wo)
						fmt.Printf("%s \n", format)
						return err

					},
				},
			},
		},
		{
			Name:    "contract",
			Aliases: []string{"c"},
			Subcommands: cli.Commands{
				cli.Command{
					Name:    "create",
					Aliases: []string{"c"},
					Usage:   "--orderId --amount",
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:  "orderId, id",
							Usage: "Идентификатор ордера",
						},
						cli.Float64Flag{
							Name:  "amount, a",
							Usage: "Сумма контракта",
						},
					},
					Action: createAction,
				},
				cli.Command{
					Name:    "view",
					Aliases: []string{"v"},
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:  "contractId, id",
							Usage: "Идентификатор контракта",
						},
					},
					Usage:  "передайте в качестве аргумента идентификатор контракта",
					Action: entityAction,
				},
				cli.Command{
					Name:    "pay",
					Aliases: []string{"p"},
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:  "contractId, id",
							Usage: "Идентификатор контракта",
						},
						cli.Float64Flag{
							Name:  "amount, a",
							Usage: "сумма оплаты",
						},
					},
					Action: func(c *cli.Context) error {
						app := xapp.AppInit(rpc)
						defer app.Conn.Close()
						return app.PayContract(c.String("contractId"), c.Float64("amount"))
					},
				},
				{
					Name:    "list",
					Aliases: []string{"l"},

					Description: "Список контрактов в системе",
					Action: func(c *cli.Context) error {
						app := xapp.AppInit(rpc)
						defer app.Conn.Close()
						cs, err := app.ListContract()
						if err != nil {
							return err
						}
						format, _ := helper.ToFormatedJson(cs)
						fmt.Printf("%s: %s\n", c.Command.Name, format)
						return nil

					},
				},
			},
		},
		{
			Name:    "account",
			Aliases: []string{"a"},
			Subcommands: cli.Commands{
				cli.Command{
					Name:    "balance",
					Aliases: []string{"b"},
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:  "accountId, id",
							Usage: "Адрес счета",
						},
					},
					Description: "Команда выведет текущий баланс счета",
					Usage:       "передайте в качестве аргумента адрес счета",
					Action:      entityAction,
				},
				cli.Command{
					Name:        "fund",
					Aliases:     []string{"f"},
					Description: "Пополнить баланс счета",
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:  "accountId, id",
							Usage: "Адрес счета",
						},
						cli.Float64Flag{
							Name:  "amount, a",
							Usage: "Сумма пополнения",
						},
					},
					Usage:  "передайте в качестве аргумента адрес счета и сумму пополнения",
					Action: entityAction,
				},
				cli.Command{
					Name:        "add",
					Aliases:     []string{"a"},
					Description: "Добавить новый счет для пользователя",
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:  "name, n",
							Usage: "Логин пользователя",
						},
						cli.StringFlag{
							Name:  "password, p",
							Usage: "Пароль пользователя",
						},
						cli.StringFlag{
							Name:  "address, a",
							Usage: "Адрес нового счета",
						},
						cli.StringFlag{
							Name:  "symbol, s",
							Usage: "Символ валюты счета",
						},
					},
					Usage: "cli account add -n Login -p Pass -a 0x99760bf112ECE3eA2CC4433A79623840ad9A7ef6, -c ETH",
					Action: func(c *cli.Context) error {
						app := xapp.AppInit(rpc)
						defer app.Conn.Close()
						acc, err := app.AddAccount(c.String("name"), c.String("password"), c.String("address"), c.String("symbol"))
						if err != nil {
							return err
						}
						format, _ := helper.ToFormatedJson(acc)
						fmt.Printf("%s: %s\n", c.Command.Name, format)
						return nil

					},
				},
			},
		},
		{
			Name:    "transaction",
			Aliases: []string{"tx"},
			Subcommands: []cli.Command{
				{
					Name:    "create",
					Usage:   "--from --to --amount",
					Aliases: []string{"c"},
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:  "fromAddress, f",
							Value: "",
							Usage: "Адрес отправления",
						},
						cli.StringFlag{
							Name:  "toAddress, t",
							Usage: "Принимающий адрес",
						},
						cli.Float64Flag{
							Name:  "amount, a",
							Usage: "Сумма транзакции",
						},
					},
					Action:      createAction,
					Description: "Создание транзакции",
				},
				{
					Name:        "view",
					Aliases:     []string{"v"},
					Usage:       "-id",
					Description: "передайте в качестве аргумента идентификатор транзакции",
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:  "id",
							Usage: "Идентификатор транзакции",
						},
					},
					Action: entityAction,
				},
				{
					Name:  "confirm",
					Usage: "-id",
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:  "id",
							Usage: "Идентификатор транзакции",
						},
					},
					Action: func(c *cli.Context) error {
						app := xapp.AppInit(rpc)
						defer app.Conn.Close()
						return app.ConfirmTransaction(c.String("id"))
					},
				},
			},
		},
		{
			Name:    "user",
			Aliases: []string{"u"},
			Subcommands: []cli.Command{
				{
					Name:        "signup",
					Aliases:     []string{"su"},
					Description: "Регистрация юзера",
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:  "name, n",
							Value: "",
							Usage: "Имя",
						},
						cli.StringFlag{
							Name:  "password, p",
							Usage: "Пароль",
						},
					},
					Action: func(c *cli.Context) error {
						app := xapp.AppInit(rpc)
						defer app.Conn.Close()
						return app.SignUp(c.String("name"), c.String("password"))
					},
				},
			},
		},
	}

	err = app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
