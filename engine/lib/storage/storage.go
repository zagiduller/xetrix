package storage

import (
	serv "engine/lib/services"
	"engine/lib/storage/leveldb"
	"engine/lib/storage/simple"
	pb "engine/lib/structs"
	"log"
)

var (
	supported = []*pb.Currency{
		//{Name: "Bitcoin", Symbol: "BTC", Type: pb.Currency_CRYPTO_CURRENCY, Decimal: 8},
		{Name: "Ethereum", Symbol: "ETH", Type: pb.Currency_CRYPTO_CURRENCY, Decimal: 8, Active: true},
		{Name: "BinanceCoin", Symbol: "BNB", Type: pb.Currency_ETH_CONTRACT_TOKEN, Decimal: 8, ContractId: "0xB8c77482e45F1F44dE1745F52C74426C631bDD52", Active: true},
		//{Name: "Dollar", Symbol: "USD", Type: pb.Currency_FIAT_CURRENCY, Decimal: 2},
		{Name: "Ruble", Symbol: "RUB", Type: pb.Currency_FIAT_CURRENCY, Decimal: 2, Active: true},
	}
)

func InitSimpleRepo() serv.IRepo {
	repo := &simple.Repository{}

	for _, sc := range supported {
		if c, _ := repo.FindCurrency(&pb.Query_Currency{Symbol: sc.Symbol}); c == nil {
			repo.CreateCurrency(sc)
		}
	}

	log.Println("INITIALIZE SIMPLE StORAGE")

	return repo
}

func InitLevelDBRepo(dbpath string) serv.IRepo {
	repo := leveldb.InitLevelDB(dbpath)

	for _, sc := range supported {
		if c, _ := repo.FindCurrency(&pb.Query_Currency{Symbol: sc.Symbol}); c == nil {
			repo.CreateCurrency(sc)
		}
	}

	log.Println("INITIALIZE LEVELDB StORAGE")

	return repo
}

//func InitMongoRepo(dsn string) serv.IRepo {
//	if len(dsn) == 0 {
//		log.Println("-dsn key is empty. Exit")
//		os.Exit(2)
//	}
//	info, err := mgo.ParseURL(dsn)
//	if err != nil {
//		panic("failed db connect: " + dsn)
//	}
//	log.Printf("Connecting to %v as a %s", info.Addrs, info.Username)
//	repo := mongo.Connect(info)
//
//	return repo
//}
