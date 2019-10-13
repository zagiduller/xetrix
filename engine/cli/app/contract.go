package app

import (
	pb "engine/lib/structs"
	"fmt"
	"log"
)

func (app *App) PayContract(cid string, amount float64) error {
	resp_c, err := app.srvc_o.GetContract(app.ctx, &pb.Query_Contract{Id: cid})
	if err != nil {
		return fmt.Errorf("PayContract: %s ", err)
	}
	if resp_c.Object == nil {
		return fmt.Errorf("PayContract: Contract not find ")
	}
	c := resp_c.Object

	_resp_tx, err := app.srvc_tp.UnderstandingRawTx(app.ctx, &pb.Query_RawTx{
		FromAddress: c.SellerSendAddress,
		ToAddress:   c.BuyerReceiveAddress,
		Amount:      amount,
	})
	if err != nil {
		return fmt.Errorf("PayContract: %s ", err)
	}

	log.Printf("UnderstandingRawTx: получен reason %s", _resp_tx.Object.Reason.Status.String())

	resp_tx, err := app.srvc_t.CreateTx(app.ctx, _resp_tx.Object)
	if err != nil {
		return fmt.Errorf("PayContract: %s ", err)
	}
	log.Printf("CreateTx: транзакция добавлена %s", resp_tx.Object.Id)

	if _, err := app.srvc_tp.ConfirmTx(app.ctx, &pb.Query_Tx{TxId: resp_tx.Object.Id}); err != nil {
		return fmt.Errorf("PayContract: %s ", err)
	}
	log.Println("ConfirmTx: транзакция подтверждена")

	log.Printf("Завершение: Транзакция создана и подтверждена\n")

	return nil
}

func (app *App) CreateContract(oid string, amount float64) error {
	if app.U == nil {
		return fmt.Errorf("CreateContract: user null ")
	}
	resp_o, err := app.srvc_o.GetOrder(app.ctx, &pb.Query_Order{Id: oid})
	if err != nil {
		return fmt.Errorf("BuyOrder: %s ", err)
	}
	resp_accs, err := app.srvc_a.GetAccount(app.ctx, &pb.Query_Account{ParticipantId: app.U.Id})
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

	resp_c, err := app.srvc_tp.CreateInternalContract(app.ctx, &pb.Query_CreateContract{
		OrderId:        resp_o.Object.Id,
		BuyerId:        app.U.Id,
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

func (app *App) GetContract(id string) (*pb.Contract, error) {
	resp, err := app.srvc_o.GetContract(app.ctx, &pb.Query_Contract{Id: id})
	if err != nil {
		return nil, fmt.Errorf("GetContract: %s ", err)
	}

	return resp.Object, nil
}

func (app *App) ListContract() ([]*pb.Contract, error) {
	resp, err := app.srvc_o.GetContracts(app.ctx, &pb.Query_Contract{})
	if err != nil {
		return nil, fmt.Errorf("GetContract: %s ", err)
	}

	return resp.Items, nil
}
