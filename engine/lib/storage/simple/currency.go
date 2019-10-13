package simple

import (
	pb "engine/lib/structs"
	"github.com/google/uuid"
)

func (repo *Repository) NewCurrInc() uint32 {
	return uint32(len(repo.currencies))
}

func (repo *Repository) CreateCurrency(currency *pb.Currency) (*pb.Currency, error) {
	if len(currency.Id) == 0 {
		currency.Id = uuid.New().String()
		currency.Inc = repo.NewCurrInc()
	}
	updated := append(repo.currencies, currency)
	repo.currencies = updated
	return currency, nil
}

func (repo *Repository) GetAllCurrency() ([]*pb.Currency, error) {
	return repo.currencies, nil
}

func (repo *Repository) FindCurrency(req *pb.Query_Currency) (*pb.Currency, error) {
	if len(req.Id) > 0 {
		for _, v := range repo.currencies {
			if v.Id == req.Id {
				return v, nil
			}
		}
	}

	if len(req.Symbol) > 0 {
		for _, v := range repo.currencies {
			if v.Symbol == req.Symbol {
				return v, nil
			}
		}
	}

	if len(req.Name) > 0 {
		for _, v := range repo.currencies {
			if v.Name == req.Name {
				return v, nil
			}
		}
	}

	return nil, nil
}

func (r *Repository) UpdateCurrency(obj *pb.Currency) (*pb.Currency, error) {
	return obj, nil
}
