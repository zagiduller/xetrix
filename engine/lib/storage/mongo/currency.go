package mongo

import (
	pb "engine/lib/structs"
	"fmt"
	"github.com/google/uuid"
	"gopkg.in/mgo.v2/bson"
	"log"
)

func (repo *Repository) CreateCurrency(currency *pb.Currency) (*pb.Currency, error) {
	//currency.Id = bson.NewObjectId().Hex()
	if c, _ := repo.FindCurrency(&pb.Query_Currency{Name: currency.Name, Symbol: currency.Symbol}); c != nil {
		return nil, fmt.Errorf("Storage-Currency: currency already exist (%s)", c.Name)
	}

	currency.Id = uuid.New().String()
	if err := repo.currencies.Insert(&currency); err != nil {
		return nil, fmt.Errorf("Storage-CreateCurrency: %s ", err)
	}
	log.Printf("Storage-Currency added %+v \n", currency)

	return currency, nil
}

func (repo *Repository) GetAllCurrency() ([]*pb.Currency, error) {
	var res []*pb.Currency
	if err := repo.currencies.Find(nil).All(&res); err != nil {
		return nil, fmt.Errorf("Storage-GetAllCurrency: %s ", err)
	}

	return res, nil
}

func (repo *Repository) FindCurrency(req *pb.Query_Currency) (*pb.Currency, error) {
	var c pb.Currency
	qr := bson.M{}

	empty := true
	if len(req.Id) > 0 {
		qr["id"] = req.Id
		empty = false
	}
	if len(req.Symbol) > 0 {
		qr["symbol"] = req.Symbol
		empty = false
	}
	if len(req.Name) > 0 {
		qr["name"] = req.Name
		empty = false
	}

	if empty {
		return nil, nil
	}

	err := repo.currencies.Find(qr).One(&c)
	if err != nil {
		return nil, fmt.Errorf("Storage-FindCurrency: %s ", err)
	}

	return &c, nil
}
