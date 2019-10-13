package simple

import (
	pb "engine/lib/structs"
	"sync"
)

var (
	Mu = &sync.Mutex{}
)

// Repository - Dummy repository, this simulates the use of a datastore
// of some kind. We'll replace this with a real implementation later on.
type Repository struct {
	currencies  []*pb.Currency
	users       []*pb.User
	accounts    []*pb.Account
	txs         []*pb.Tx
	orders      []*pb.Order
	contracts   []*pb.Contract
	withdrawals []*pb.WithdrawalOrder
}
