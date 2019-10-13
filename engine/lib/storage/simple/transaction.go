package simple

import (
	"engine/lib/helper"
	pb "engine/lib/structs"
	"fmt"
	"github.com/google/uuid"
)

func (repo *Repository) SaveTx(tx *pb.Tx) (*pb.Tx, error) {
	if len(tx.Id) > 0 {
		return nil, fmt.Errorf("Storage-SaveTx: Id required")
	}

	tx.Id = uuid.New().String()
	tx.CreatedAt = helper.CurrentTimestamp()
	updated := append(repo.txs, tx)
	repo.txs = updated
	return tx, nil
}

func (repo *Repository) UpdateTx(tx *pb.Tx) (*pb.Tx, error) {
	return tx, nil
}

func (repo *Repository) GetAllTxs() ([]*pb.Tx, error) {
	return repo.txs, nil
}

func (repo *Repository) GetTxs(req *pb.Query_Tx) ([]*pb.Tx, error) {
	//pb.Query
	if len(req.ContractId) > 0 {
		var txs []*pb.Tx
		for _, t := range repo.txs {
			if (t.Reason.Status == pb.TxReason_SELLER_CONTRACT_TX || t.Reason.Status == pb.TxReason_BUYER_CONTRACT_TX) && t.Reason.ContractId == req.ContractId {
				txs = append(txs, t)
			}
		}
		return txs, nil
	}

	if req.Reason != nil {
		var txs []*pb.Tx
		for _, t := range repo.txs {
			if t.Reason.Status == req.Reason.Status {
				txs = append(txs, t)
			}
		}
		return txs, nil
	}

	if len(req.InPStxId) > 0 {
		var txs []*pb.Tx
		for _, t := range repo.txs {
			if t.Reason.InPStxId == req.InPStxId {
				txs = append(txs, t)
			}
		}
		return txs, nil
	}

	if len(req.Address) > 0 {
		var txs []*pb.Tx
		for _, t := range repo.txs {
			if len(t.FromAddress) > 0 && len(t.ToAddress) > 0 {
				// TODO кажется, можно оставить только две строчки для всего этого блока
				if (t.FromAddress == req.Address) || (t.ToAddress == req.Address) {
					txs = append(txs, t)
				}
			} else if len(t.FromAddress) > 0 {
				if t.FromAddress == req.Address {
					txs = append(txs, t)
				}
			} else if len(t.ToAddress) > 0 {
				if t.ToAddress == req.Address {
					txs = append(txs, t)
				}
			}
		}
		return txs, nil
	}

	if len(req.TxId) > 0 {
		var txs []*pb.Tx
		for _, t := range repo.txs {
			if t.Id == req.TxId {
				txs = append(txs, t)
			}
		}
		return txs, nil
	}

	return nil, nil
}

func (repo *Repository) GetTx(req *pb.Query_Tx) (*pb.Tx, error) {
	if len(req.TxId) > 0 {
		for _, t := range repo.txs {
			if t.Id == req.TxId {
				return t, nil
			}
		}
	}
	if len(req.InPStxId) > 0 {
		for _, t := range repo.txs {
			if t.Reason.InPStxId == req.InPStxId {
				return t, nil
			}
		}
	}
	return nil, nil
}

func (r *Repository) ConfirmTx(tx *pb.Tx) (bool, error) {
	tx.Status = pb.TxStatus_CONFIRMED
	return true, nil
}
