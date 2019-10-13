package app

import (
	"engine/lib/helper"
	pb "engine/lib/structs"
	"fmt"
	"log"
)

func (app *App) Fund(addr string, amount float64) error {
	resp_raw_tx, err := app.srvc_tp.UnderstandingRawTx(app.ctx, &pb.Query_RawTx{
		FromAddress: helper.RandStringRunes(10),
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

func (app *App) Balance(addr string) (*pb.AccountBalance, error) {
	b_resp, err := app.srvc_ab.GetBalance(app.ctx, &pb.Query_Account{
		Address: addr,
	})
	if err != nil {
		return nil, fmt.Errorf("Balance: %s", err)
	}
	return b_resp.Object, nil
}

func (app *App) GetAccounts(addr string) (*pb.AccountBalance, error) {
	b_resp, err := app.srvc_ab.GetBalance(app.ctx, &pb.Query_Account{
		Address: addr,
	})
	if err != nil {
		return nil, fmt.Errorf("Balance: %s", err)
	}
	return b_resp.Object, nil

}

func (app *App) AddAccount(name, pass, addr, symbol string) (*pb.Account, error) {
	if err := app.StartSession(name, pass, false); err != nil {
		return nil, fmt.Errorf("AddAccoutn: %s", err)
	}

	respAcc, err := app.srvc_a.CreateAccount(app.ctx, &pb.Query_CreateAccount{
		OwnerId:        app.U.Id,
		CurrencySymbol: symbol,
		Address:        addr,
	})
	if err != nil {
		return nil, fmt.Errorf("AddAccount: %s", err)
	}
	return respAcc.Object, nil
}
