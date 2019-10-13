package services

import (
	"context"
	"engine/lib/services/events"
	"engine/lib/structs"
)

type ServiceWithdrawal struct {
}

func NewServiceWithdrawal(bus *events.Bus) {

}

func (s *ServiceWithdrawal) CreateWithdrawal(ctx context.Context, q *structs.Query_Withdrawal) (error, *structs.WithdrawalOrder) {

	return nil, nil
}
