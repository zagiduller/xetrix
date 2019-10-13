package mongo

import (
	"engine/lib/helper"
	pb "engine/lib/structs"
	"fmt"
	"github.com/google/uuid"
	"gopkg.in/mgo.v2/bson"
	"log"
)

func (r *Repository) CreateAccount(acc *pb.Account) (*pb.Account, error) {
	acc.Id = uuid.New().String()
	acc.CreatedAt = helper.CurrentTimestamp()
	if err := r.accounts.Insert(&acc); err != nil {
		return nil, fmt.Errorf("Storage-CreateAccount: %s ", err)
	}
	log.Printf("Storage-CreateAccount: Account added %+v \n", acc)

	return acc, nil
}

func (r *Repository) FindAccountById(id string) (*pb.Account, error) {
	var a *pb.Account
	if err := r.accounts.Find(bson.M{"id": id}).One(&a); err != nil {
		return nil, err
	}

	return a, nil
}

func (r *Repository) FindAccountByAddress(address string) (*pb.Account, error) {
	var a *pb.Account
	if err := r.accounts.Find(bson.M{"address": address}).One(&a); err != nil {
		return nil, err
	}

	return a, nil
}

func (r *Repository) FindAccountByOwnerId(id string) ([]*pb.Account, error) {

	var result []*pb.Account

	if err := r.accounts.Find(bson.M{"ownerid": id}).All(&result); err != nil {
		return nil, err
	}
	return result, nil
}

func (r *Repository) FindAccountByCurrencyAndOwnerId(id, currencySymbol string) (*pb.Account, error) {
	var result pb.Account

	if err := r.accounts.Find(bson.M{"ownerid": id, "currency.symbol": currencySymbol}).One(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (r *Repository) GetAllAccount() ([]*pb.Account, error) {
	var result []*pb.Account
	if err := r.accounts.Find(nil).All(&result); err != nil {
		return nil, err
	}
	return result, nil
}
