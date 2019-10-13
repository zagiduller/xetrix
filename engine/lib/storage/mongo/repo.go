package mongo

import (
	//pb "mxp-protobuf/pkg"
	"context"
	"gopkg.in/mgo.v2"
	"log"
)

type Repository struct {
	ctx        context.Context
	currencies *mgo.Collection
	users      *mgo.Collection
	accounts   *mgo.Collection
	txs        *mgo.Collection
	orders     *mgo.Collection
	contracts  *mgo.Collection
}

func Connect(info *mgo.DialInfo) *Repository {
	repo := Repository{}
	sess, err := mgo.DialWithInfo(info)
	if err != nil {
		panic("Storage: Db session create fail: " + err.Error())
	}
	if err := sess.Ping(); err != nil {
		panic("Storage: Ping session fail: " + err.Error())
	}
	repo.ctx = context.Background()

	db := sess.DB("mxp")

	repo.currencies = db.C("currencies")
	repo.users = db.C("users")
	repo.accounts = db.C("accounts")
	repo.txs = db.C("txs")
	repo.orders = db.C("orders")
	repo.contracts = db.C("contracts")

	log.Printf("Storage: Mongo session start")
	return &repo
}
