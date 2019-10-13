package leveldb

import (
	"fmt"
	"github.com/syndtr/goleveldb/leveldb"
)

type Repository struct {
	currencies  *leveldb.DB
	users       *leveldb.DB
	accounts    *leveldb.DB
	txs         *leveldb.DB
	orders      *leveldb.DB
	contracts   *leveldb.DB
	withdrawals *leveldb.DB
}

func InitLevelDB(dbpath string) *Repository {
	var err error
	repo := new(Repository)

	repo.currencies, err = leveldb.OpenFile(dbpath+"/currencies", nil)
	if err != nil {
		panic(err)
	}

	repo.users, err = leveldb.OpenFile(dbpath+"/users", nil)
	if err != nil {
		repo.Close()
		panic(err)
	}

	repo.accounts, err = leveldb.OpenFile(dbpath+"/accounts", nil)
	if err != nil {
		repo.Close()
		panic(err)
	}

	repo.txs, err = leveldb.OpenFile(dbpath+"/txs", nil)
	if err != nil {
		repo.Close()
		panic(err)
	}

	repo.orders, err = leveldb.OpenFile(dbpath+"/orders", nil)
	if err != nil {
		repo.Close()
		panic(err)
	}

	repo.contracts, err = leveldb.OpenFile(dbpath+"/contracts", nil)
	if err != nil {
		repo.Close()
		panic(err)
	}

	repo.withdrawals, err = leveldb.OpenFile(dbpath+"/withdrawals", nil)
	if err != nil {
		repo.Close()
		panic(err)
	}

	return repo
}

func (r *Repository) Close() {
	if r.currencies != nil {
		r.currencies.Close()
	}
	if r.users != nil {
		r.users.Close()
	}
	if r.accounts != nil {
		r.accounts.Close()
	}
	if r.txs != nil {
		r.txs.Close()
	}
	if r.orders != nil {
		r.orders.Close()
	}
	if r.contracts != nil {
		r.contracts.Close()
	}
	if r.withdrawals != nil {
		r.withdrawals.Close()
	}
	fmt.Println("Leveldb Repo closed")
}
