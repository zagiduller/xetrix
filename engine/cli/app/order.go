package app

import (
	"engine/lib/helper"
	pb "engine/lib/structs"
	"fmt"
	"google.golang.org/grpc/metadata"
	"log"

	"strconv"
)

func (app *App) CreateOrder(from, to string, amount, price float64) error {
	if app.U == nil {
		return fmt.Errorf("CreateOrder:  participant null ")
	}
	q_ord := &pb.Query_Order{}
	q_ord.OwnerId = app.U.Id

	q_ord.SellCurrencySymbol = from
	q_ord.BuyCurrencySymbol = to
	q_ord.Amount = amount
	q_ord.Price = price
	fmt.Printf("ctx: %#v, %#v\n", app.ctx.Value("pid"), app.ctx)

	resp_accs, err := app.srvc_a.GetAccount(app.ctx, &pb.Query_Account{ParticipantId: app.U.Id})
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

	//if resp_cms_acc, err := app.srvc_a.GetAccount(app.ctx, &pb.Query_Account{ParticipantId: q_ord.OwnerId, CurrencySymbol: "MXPC"}); err == nil {
	//	cm_query := &pb.Query_CalculateCommission{
	//		Participant: &pb.Participant{Id: app.P.Id},
	//		Order: &pb.Order{
	//			Amount: amount,
	//		},
	//	}
	//
	//	if cm_resp, err := app.srvc_cm.Calc(app.ctx, cm_query); err != nil {
	//		log.Printf("CreateOrder: %s", err)
	//	}
	//	else {
	//		if err := app.Fund(resp_cms_acc.Object.Address, cm_resp.Amount); err != nil {
	//			return err
	//		}
	//	}
	//}

	q_ord.FrontMetaData = &pb.FrontMetaData{UserPriceInput: strconv.FormatFloat(q_ord.Price, 'f', 6, 64)}

	resp_o, err := app.srvc_o.CreateOrder(app.ctx, q_ord)
	if err != nil {
		return fmt.Errorf("Create Order: %s ", err)
	}
	log.Printf("Создан ордер: %s \n", resp_o.Object.Id)
	return nil
}

func (app *App) GetOrder(id string) (*pb.Order, error) {
	resp, err := app.srvc_o.GetOrder(app.ctx, &pb.Query_Order{Id: id})
	if err != nil {
		return nil, fmt.Errorf("GetOrder: %s ", err)
	}
	return resp.Object, nil
}

func (app *App) CancelOrder(id string) (*pb.Order, error) {
	resp, err := app.srvc_o.CancelOrder(app.ctx, &pb.Query_Order{Id: id})
	if err != nil {
		return nil, fmt.Errorf("CancelOrder: %s", err)
	}
	fmt.Printf("%+v", resp)
	return resp.Object, nil
}

func (app *App) ListOrder() ([]*pb.Order, error) {
	resp, err := app.srvc_o.GetOrders(app.ctx, &pb.Query_Order{})
	if err != nil {
		return nil, fmt.Errorf("ListOrder: %s ", err)
	}
	return resp.Items, nil
}

func (app *App) OrderGenerate(n int, a, p float64) {
	log.Printf("Генерация %n ордеров. Amount: %f, p: %f", n, a, p)
	rcurrs, _ := app.srvc_c.GetCurrency(app.ctx, &pb.Query_Currency{})
	var currs []string
	for _, c := range rcurrs.Items {
		currs = append(currs, c.Symbol)
	}

	currLen := len(currs)
	for i := 0; i < n; i++ {
		var from, to string
		fKey := i % currLen
		from = currs[fKey]
		if (currLen - 1) == fKey {
			to = currs[0]
		} else {
			to = currs[fKey+1]
		}
		log.Printf("Создается ордер #%d. %s>%s ", i+1, from, to)

		if err := app.SignUp(fmt.Sprintf("Generator %d", i+1), helper.RandStringRunes(8)); err != nil {
			log.Printf("Generate: %s", err)
			return
		}
		if err := app.CreateOrder(from, to, a, p); err != nil {
			log.Printf("Generate: %s", err)
			return
		}
	}

}

func (app *App) CreateWithdrawal(addr string, amount float64) (*pb.WithdrawalOrder, error) {
	ab, err := app.GetAccounts(addr)
	if err != nil {
		return nil, err
	}

	app.ctx = metadata.AppendToOutgoingContext(app.ctx, "pid", ab.AccountOwnerId)

	attrs := []*pb.KeyValueAttribute{
		{Key: "Test1", Value: "Test2"},
		{Key: "Test2", Value: "Test3"},
	}

	wo, err := app.srvc_o.CreateWithdrawal(app.ctx, &pb.Query_Withdrawal{
		SendingAddress: addr,
		Attributes:     attrs,
		Amount:         amount,
	})

	return wo, err
}
