package leveldb

import (
	"engine/lib/helper"
	pb "engine/lib/structs"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/google/uuid"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/util"
	"log"
)

func (r *Repository) CreateAccount(obj *pb.Account) (*pb.Account, error) {

	if len(obj.Id) == 0 {
		obj.Id = uuid.New().String()
	}
	obj.CreatedAt = helper.CurrentTimestamp()
	byted, err := proto.Marshal(obj)
	if err != nil {
		return nil, fmt.Errorf("Repo-CreateAccount: %s", err)
	}
	batch := new(leveldb.Batch)

	batch.Put([]byte("object-"+obj.Id), byted)
	batch.Put([]byte("ownerId-"+obj.OwnerId+"-"+obj.Id), []byte(obj.Id))
	batch.Put([]byte("address-"+obj.Address), []byte(obj.Id))
	if err := r.accounts.Write(batch, nil); err != nil {
		return nil, fmt.Errorf("Repo-CreateAccount: %s", err)
	}
	return obj, nil
}

func (r *Repository) FindAccountById(id string) (*pb.Account, error) {
	if len(id) > 0 {
		obj := new(pb.Account)
		byted, err := r.accounts.Get([]byte("object-"+id), nil)
		if err != nil {
			return nil, fmt.Errorf("Repo-FindAccountById: %s", err)
		}
		if err := proto.Unmarshal(byted, obj); err != nil {
			return nil, fmt.Errorf("Repo-FindAccountById: %s", err)
		}
		return obj, nil
	}

	return nil, nil
}

func (r *Repository) FindAccountByAddress(address string) (*pb.Account, error) {
	if len(address) > 0 {
		obj := new(pb.Account)

		byteId, err := r.accounts.Get([]byte("address-"+address), nil)
		if err != nil {
			return nil, fmt.Errorf("Repo-FindAccountByAddress: %s", err)
		}

		byted, err := r.accounts.Get([]byte("object-"+string(byteId)), nil)
		if err != nil {
			return nil, fmt.Errorf("Repo-FindAccountByAddress: %s", err)
		}
		if err := proto.Unmarshal(byted, obj); err != nil {
			return nil, fmt.Errorf("Repo-FindAccountByAddress: %s", err)
		}
		return obj, nil
	}
	return nil, nil
}

func (r *Repository) FindAccountByOwnerId(id string) ([]*pb.Account, error) {
	var result []*pb.Account
	iter := r.accounts.NewIterator(util.BytesPrefix([]byte("ownerId-"+id)), nil)
	for iter.Next() {
		if obj, err := r.FindAccountById(string(iter.Value())); err == nil {
			result = append(result, obj)
		} else {
			log.Printf("Repo-FindAccountByOwnerId: Error. %s", err)
		}
	}
	return result, nil
}

func (r *Repository) FindAccountByCurrencyAndOwnerId(id, currencySymbol string) (*pb.Account, error) {
	accs, err := r.FindAccountByOwnerId(id)
	if err == nil {
		for _, a := range accs {
			if currencySymbol == a.Currency.Symbol {
				return a, nil
			}
		}
	}
	return nil, fmt.Errorf("Repo-FindAccountByCurrencyAndOwnerId: %s", err)
}

func (r *Repository) GetAllAccount() ([]*pb.Account, error) {
	var result []*pb.Account
	iter := r.accounts.NewIterator(util.BytesPrefix([]byte("object-")), nil)
	for iter.Next() {
		obj := new(pb.Account)
		if err := proto.Unmarshal(iter.Value(), obj); err != nil {
			return nil, fmt.Errorf("Repo-GetAllAccount: %s", err)
		}
		result = append(result, obj)
	}
	return result, nil
}

func (r *Repository) UpdateAccount(obj *pb.Account) (*pb.Account, error) {
	byted, err := proto.Marshal(obj)
	if err != nil {
		return nil, fmt.Errorf("Repo-UpdateAccount: %s", err)
	}
	if err := r.accounts.Put([]byte("object-"+obj.Id), byted, nil); err != nil {
		return nil, fmt.Errorf("Repo-UpdateAccount: %s", err)
	}
	return obj, nil
}
