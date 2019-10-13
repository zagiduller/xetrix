package leveldb

import (
	"bytes"
	"encoding/binary"
	pb "engine/lib/structs"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/google/uuid"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/util"
)

func (repo *Repository) NewCurrInc() uint32 {
	var inc uint32
	v, err := repo.currencies.Get([]byte("inc-"), nil)
	if err != nil {
		inc = 0
	} else {
		inc = binary.LittleEndian.Uint32(v)
		inc += 1
	}

	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, inc)
	repo.currencies.Put([]byte("inc-"), buf.Bytes(), nil)

	return inc
}

func (repo *Repository) CreateCurrency(currency *pb.Currency) (*pb.Currency, error) {
	if exist, _ := repo.FindCurrency(&pb.Query_Currency{Symbol: currency.Symbol}); exist != nil {
		return nil, fmt.Errorf("Storage-CreateCurrency: Currency as %s exist", currency.Symbol)
	}

	if len(currency.Id) == 0 {
		currency.Id = uuid.New().String()
		currency.Inc = repo.NewCurrInc()
	}
	batch := new(leveldb.Batch)

	byted, err := proto.Marshal(currency)
	if err != nil {
		return nil, fmt.Errorf("Repo-CreateCurrency: %s", err)
	}

	batch.Put([]byte("object-"+currency.Id), byted)
	batch.Put([]byte("symbol-"+currency.Symbol), []byte(currency.Id))
	batch.Put([]byte("name-"+currency.Name), []byte(currency.Id))

	if err := repo.currencies.Write(batch, nil); err != nil {
		return nil, fmt.Errorf("Repo-CreateCurrency: %s", err)
	}

	return currency, nil
}

func (repo *Repository) GetAllCurrency() ([]*pb.Currency, error) {
	var result []*pb.Currency
	iter := repo.currencies.NewIterator(util.BytesPrefix([]byte("object-")), nil)
	for iter.Next() {
		c := new(pb.Currency)
		if err := proto.Unmarshal(iter.Value(), c); err == nil {
			result = append(result, c)
		} else {
			fmt.Println(err)
		}
	}

	return result, nil
}

func (repo *Repository) FindCurrency(req *pb.Query_Currency) (*pb.Currency, error) {
	curr := new(pb.Currency)
	if len(req.Id) > 0 {
		byteCurr, err := repo.currencies.Get([]byte("object-"+req.Id), nil)
		if err != nil {
			return nil, fmt.Errorf("Repo-FindCurrency: %s", err)
		}
		if err := proto.Unmarshal(byteCurr, curr); err != nil {
			return nil, fmt.Errorf("Repo-FindCurrency: %s", err)
		}

		return curr, nil
	}

	if len(req.Symbol) > 0 {
		byteId, err := repo.currencies.Get([]byte("symbol-"+req.Symbol), nil)
		if err != nil {
			return nil, fmt.Errorf("Repo-FindCurrency: %s", err)
		}
		byteCurr, err := repo.currencies.Get([]byte("object-"+string(byteId)), nil)
		if err != nil {
			return nil, fmt.Errorf("Repo-FindCurrency: %s", err)
		}
		if err := proto.Unmarshal(byteCurr, curr); err != nil {
			return nil, fmt.Errorf("Repo-FindCurrency: %s", err)
		}
		return curr, nil
	}

	if len(req.Name) > 0 {
		byteId, err := repo.currencies.Get([]byte("name-"+req.Name), nil)
		if err != nil {
			return nil, fmt.Errorf("Repo-FindCurrency: %s", err)
		}
		byteCurr, err := repo.currencies.Get([]byte("object-"+string(byteId)), nil)
		if err != nil {
			return nil, fmt.Errorf("Repo-FindCurrency: %s", err)
		}
		if err := proto.Unmarshal(byteCurr, curr); err != nil {
			return nil, fmt.Errorf("Repo-FindCurrency: %s", err)
		}

		return curr, nil
	}

	return nil, nil
}

func (r *Repository) UpdateCurrency(obj *pb.Currency) (*pb.Currency, error) {
	byted, err := proto.Marshal(obj)
	if err != nil {
		return nil, fmt.Errorf("Repo-UpdateCurrency: %s", err)
	}
	if err := r.currencies.Put([]byte("object-"+obj.Id), byted, nil); err != nil {
		return nil, fmt.Errorf("Repo-UpdateCurrency: %s", err)
	}
	return obj, nil
}
