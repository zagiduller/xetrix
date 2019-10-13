package leveldb

import (
	"engine/lib/helper"
	pb "engine/lib/structs"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/google/uuid"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/util"
)

func (repo *Repository) SaveTx(tx *pb.Tx) (*pb.Tx, error) {
	if len(tx.Id) > 0 {
		return nil, fmt.Errorf("Storage-SaveTx: Id required")
	}

	tx.Id = uuid.New().String()
	tx.CreatedAt = helper.CurrentTimestamp()

	byted, err := proto.Marshal(tx)
	if err != nil {
		return nil, fmt.Errorf("Repo-CreateUser: %s", err)
	}
	batch := new(leveldb.Batch)

	batch.Put([]byte("object-"+tx.Id), byted)
	batch.Put([]byte("currencySymbol-"+tx.CurrencySymbol+"-"+tx.Id), []byte(tx.Id))
	batch.Put([]byte("address-"+tx.ToAddress+"-to-"+tx.Id), []byte(tx.Id))
	batch.Put([]byte("address-"+tx.FromAddress+"-from-"+tx.Id), []byte(tx.Id))
	batch.Put([]byte("owner-"+tx.FromAddressOwnerId+"-from-"+tx.Id), []byte(tx.Id))
	batch.Put([]byte("owner-"+tx.ToAddressOwnerId+"-to-"+tx.Id), []byte(tx.Id))
	batch.Put([]byte("reasonStatus-"+tx.Reason.Status.String()+"-"+tx.Id), []byte(tx.Id))

	if tx.Reason.Status == pb.TxReason_SELLER_CONTRACT_TX || tx.Reason.Status == pb.TxReason_BUYER_CONTRACT_TX {
		batch.Put([]byte("contractId-"+tx.Reason.ContractId+"-"+tx.Id), []byte(tx.Id))
	}

	if len(tx.Reason.InPStxId) > 0 {
		batch.Put([]byte("inPStxId-"+tx.Reason.InPStxId), []byte(tx.Id))
	}

	if err := repo.txs.Write(batch, nil); err != nil {
		return nil, fmt.Errorf("Repo-SaveTx: %s", err)
	}

	return tx, nil
}

func (repo *Repository) UpdateTx(tx *pb.Tx) (*pb.Tx, error) {
	prevTx, err := repo.GetTx(&pb.Query_Tx{TxId: tx.Id})
	if err != nil {
		return nil, fmt.Errorf("Repo-UpdateTx: %s", err)
	}

	byted, err := proto.Marshal(tx)
	if err != nil {
		return nil, fmt.Errorf("Repo-UpdateTx: %s", err)
	}
	if err := repo.txs.Put([]byte("object-"+tx.Id), byted, nil); err != nil {
		return nil, fmt.Errorf("Repo-UpdateTx: %s", err)
	}
	if prevTx.Reason.Status != tx.Reason.Status {
		repo.txs.Delete([]byte("reasonStatus-"+prevTx.Reason.Status.String()+"-"+tx.Id), nil)
		repo.txs.Put([]byte("reasonStatus-"+tx.Reason.Status.String()+"-"+tx.Id), []byte(tx.Id), nil)
	}

	return tx, nil
}

func (repo *Repository) GetAllTxs() ([]*pb.Tx, error) {
	var result []*pb.Tx
	iter := repo.txs.NewIterator(util.BytesPrefix([]byte("object-")), nil)
	for iter.Next() {
		obj := new(pb.Tx)
		if err := proto.Unmarshal(iter.Value(), obj); err == nil {
			result = append(result, obj)
		} else {
			fmt.Printf("Repo-GetAllTxs: Error. %s \n", err)
		}
	}
	return result, nil
}

func (repo *Repository) GetTxs(req *pb.Query_Tx) ([]*pb.Tx, error) {
	var result []*pb.Tx
	if len(req.ContractId) > 0 {
		iter := repo.txs.NewIterator(util.BytesPrefix([]byte("contractId-"+req.ContractId)), nil)
		for iter.Next() {
			if obj, err := repo.GetTx(&pb.Query_Tx{TxId: string(iter.Value())}); err == nil {
				result = append(result, obj)
			} else {
				fmt.Printf("Repo-GetTxs: Error. %s\n", err)
			}
		}
		return result, nil
	}
	//
	if req.Reason != nil {
		iter := repo.txs.NewIterator(util.BytesPrefix([]byte("reasonStatus-"+req.Reason.Status.String())), nil)
		for iter.Next() {
			if obj, err := repo.GetTx(&pb.Query_Tx{TxId: string(iter.Value())}); err == nil {
				result = append(result, obj)
			} else {
				fmt.Printf("Repo-GetTxs: Error. %s\n", err)
			}
		}
		return result, nil
	}

	if len(req.InPStxId) > 0 {
		bytedId, err := repo.txs.Get([]byte("inPStxId-"+req.InPStxId), nil)
		if err != nil {
			return nil, fmt.Errorf("Repo-GetTxs: %s", err)
		}
		obj := new(pb.Tx)
		byted, err := repo.txs.Get([]byte("object-"+string(bytedId)), nil)
		if err != nil {
			return nil, fmt.Errorf("Repo-GetTxs: %s", err)
		}
		if err := proto.Unmarshal(byted, obj); err != nil {
			return nil, fmt.Errorf("Repo-GetTxs: %s", err)
		}
		result = append(result, obj)
		return result, nil
	}

	if len(req.Address) > 0 {
		iter := repo.txs.NewIterator(util.BytesPrefix([]byte("address-"+req.Address)), nil)
		for iter.Next() {
			if obj, err := repo.GetTx(&pb.Query_Tx{TxId: string(iter.Value())}); err == nil {
				result = append(result, obj)
			} else {
				fmt.Printf("Repo-GetTxs: Error. %s\n", err)
			}
		}
		return result, nil
	}

	if len(req.TxId) > 0 {
		tx, err := repo.GetTx(req)
		if err != nil {
			fmt.Printf("Repo-GetTxs: Error. %s\n", err)
		}
		result = append(result, tx)
		return result, nil
	}

	return nil, nil
}

func (repo *Repository) GetTx(req *pb.Query_Tx) (*pb.Tx, error) {
	if len(req.TxId) > 0 {
		obj := new(pb.Tx)
		byted, err := repo.txs.Get([]byte("object-"+req.TxId), nil)
		if err != nil {
			return nil, fmt.Errorf("Repo-GetTx: %s", err)
		}
		if err := proto.Unmarshal(byted, obj); err != nil {
			return nil, fmt.Errorf("Repo-GetTx: %s", err)
		}
		return obj, nil
	}

	if len(req.InPStxId) > 0 {
		bytedId, err := repo.txs.Get([]byte("inPStxId-"+req.InPStxId), nil)
		if err != nil {
			return nil, fmt.Errorf("Repo-GetTx: %s", err)
		}
		obj := new(pb.Tx)
		byted, err := repo.txs.Get([]byte("object-"+string(bytedId)), nil)
		if err != nil {
			return nil, fmt.Errorf("Repo-GetTx: %s", err)
		}
		if err := proto.Unmarshal(byted, obj); err != nil {
			return nil, fmt.Errorf("Repo-GetTx: %s", err)
		}
		return obj, nil
	}

	return nil, nil
}

func (r *Repository) ConfirmTx(tx *pb.Tx) (bool, error) {
	tx.Status = pb.TxStatus_CONFIRMED
	if _, err := r.UpdateTx(tx); err != nil {
		return false, err
	}
	return true, nil
}
