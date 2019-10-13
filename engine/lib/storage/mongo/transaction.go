package mongo

import (
	"engine/lib/helper"
	pb "engine/lib/structs"
	"fmt"
	"github.com/google/uuid"
	"gopkg.in/mgo.v2/bson"
	"log"
)

func (r *Repository) SaveTx(tx *pb.Tx) (*pb.Tx, error) {
	tx.Id = uuid.New().String()
	tx.CreatedAt = helper.CurrentTimestamp()
	if err := r.txs.Insert(&tx); err != nil {
		return nil, fmt.Errorf("Storage-SaveTx: %s ", err)
	}
	log.Printf("Storage-SaveTx: Tx saved %+v \n", tx)

	return tx, nil
}

func (r *Repository) UpdateTx(tx *pb.Tx) (*pb.Tx, error) {
	if err := r.txs.Update(bson.M{"id": tx.Id}, tx); err != nil {
		return nil, fmt.Errorf("Storage-UpdateTx: %s ", err)
	}
	log.Printf("Storage-UpdateTx: Tx saved %+v \n", tx)

	return tx, nil
}

func (r *Repository) GetTxs(req *pb.Query_Tx) ([]*pb.Tx, error) {
	qbson := make(bson.M)
	var txs []*pb.Tx
	//pb.Query
	if len(req.ContractId) > 0 {
		qbson["contractid"] = req.ContractId
	}

	if req.Reason != nil {
		qbson["reason"] = req.Reason
	}

	if len(req.Address) > 0 {
		qbson["$or"] = []bson.M{
			{"fromaddress": req.Address},
			{"toaddress": req.Address},
		}
	}

	if err := r.txs.Find(qbson).All(&txs); err != nil {
		return nil, fmt.Errorf("Storage-GetTxs: %s ", err)
	}
	log.Printf("Storage-GetTxs: %d ", len(txs))

	return txs, nil
}

func (r *Repository) GetTx(req *pb.Query_Tx) (*pb.Tx, error) {
	var tx pb.Tx
	if len(req.TxId) > 0 {
		if err := r.txs.Find(bson.M{"id": req.TxId}).One(&tx); err != nil {
			return nil, fmt.Errorf("Storage-GetTx: %s ", err)
		}
	}
	if len(req.PaymentSystemID) > 0 {
		if err := r.txs.Find(bson.M{"reason.paymentSystemID": req.PaymentSystemID}).One(&tx); err != nil {
			return nil, fmt.Errorf("Storage-GetTx: %s ", err)
		}
	}
	return &tx, nil
}

func (r *Repository) GetAllTxs() ([]*pb.Tx, error) {
	var res []*pb.Tx
	if err := r.txs.Find(nil).All(&res); err != nil {
		return nil, fmt.Errorf("Storage-GetAllTxs: %s ", err)
	}
	return res, nil
}

func (r *Repository) ConfirmTx(tx *pb.Tx) (bool, error) {
	tx.Status = pb.TxStatus_CONFIRMED
	if err := r.txs.Update(bson.M{"id": tx.Id}, bson.M{"$set": bson.M{"status": pb.TxStatus_CONFIRMED}}); err != nil {
		return false, err
	}
	log.Printf("Storage-ConfirmTx: Tx confirmed: %s ", tx.Id)
	return true, nil
}
