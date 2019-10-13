package simple

import (
	"engine/lib/helper"
	pb "engine/lib/structs"
	"github.com/google/uuid"
)

func (r *Repository) CreateAccount(acc *pb.Account) (*pb.Account, error) {
	Mu.Lock()
	defer Mu.Unlock()

	if len(acc.Id) == 0 {
		acc.Id = uuid.New().String()
	}
	acc.CreatedAt = helper.CurrentTimestamp()
	updated := append(r.accounts, acc)
	r.accounts = updated

	return acc, nil
}

func (r *Repository) FindAccountById(id string) (*pb.Account, error) {
	Mu.Lock()
	defer Mu.Unlock()

	for _, a := range r.accounts {
		if a.Id == id {
			return a, nil
		}
	}

	return nil, nil
}

func (r *Repository) FindAccountByAddress(address string) (*pb.Account, error) {
	Mu.Lock()
	defer Mu.Unlock()

	for _, a := range r.accounts {
		if a.Address == address {
			return a, nil
		}
	}
	return nil, nil
}

func (r *Repository) FindAccountByOwnerId(id string) ([]*pb.Account, error) {
	Mu.Lock()
	defer Mu.Unlock()

	var result []*pb.Account
	for _, a := range r.accounts {
		if a.OwnerId == id {
			result = append(result, a)
		}
	}
	return result, nil
}

func (r *Repository) FindAccountByCurrencyAndOwnerId(id, currencySymbol string) (*pb.Account, error) {
	Mu.Lock()
	defer Mu.Unlock()

	for _, a := range r.accounts {
		if a.OwnerId == id && currencySymbol == a.Currency.Symbol {
			return a, nil
		}
	}
	return nil, nil
}

func (r *Repository) GetAllAccount() ([]*pb.Account, error) {
	Mu.Lock()
	defer Mu.Unlock()

	return r.accounts, nil
}

func (r *Repository) UpdateAccount(acc *pb.Account) (*pb.Account, error) {
	return acc, nil
}
