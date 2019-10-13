package app

import (
	pb "engine/lib/structs"
	"fmt"
	"log"
)

//--from --to --amount
func (app *App) CreateTransaction(from, to string, amount float64) error {
	_tx_resp, err := app.srvc_tp.UnderstandingRawTx(app.ctx, &pb.Query_RawTx{
		FromAddress: from, ToAddress: to, Amount: amount,
	})
	if err != nil {
		return fmt.Errorf("CreateTransaction: %s", err)
	}

	log.Printf("Транзакция опознана. Получен статус: %s", _tx_resp.Object.Reason.Status)

	tx_resp, err := app.srvc_t.CreateTx(app.ctx, _tx_resp.Object)
	if err != nil {
		return fmt.Errorf("CreateTransaction: %s", err)
	}
	log.Printf("Транзакция сохранена. id: %s", tx_resp.Object.Id)

	return nil
}

func (app *App) ConfirmTransaction(id string) error {
	if _, err := app.srvc_tp.ConfirmTx(app.ctx, &pb.Query_Tx{TxId: id}); err != nil {
		return fmt.Errorf("ConfirmTransaction: %s", err)
	}
	fmt.Printf("Транзакция подтверждена")
	return nil
}

func (app *App) GetTransaction(id string) ([]*pb.Tx, error) {
	tx_resp, err := app.srvc_t.GetTx(app.ctx, &pb.Query_Tx{TxId: id})
	if err != nil {
		return nil, fmt.Errorf("GetTransaction: %s", err)
	}
	log.Printf("Найдено (%d) транзакций", tx_resp.ItemsCount)

	return tx_resp.Items, nil
}
