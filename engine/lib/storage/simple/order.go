package simple

import (
	"engine/lib/helper"
	pb "engine/lib/structs"
	"fmt"
	"github.com/google/uuid"
	"log"
)

func (repo *Repository) CreateOrder(in *pb.Order) (*pb.Order, error) {
	in.Id = uuid.New().String()
	in.CreatedAt = helper.CurrentTimestamp()
	in.Status = &pb.DealStatus{
		Status:    pb.DealStatus_CREATED,
		CreatedAt: in.CreatedAt,
	}
	update := append(repo.orders, in)
	repo.orders = update
	log.Printf("Repo-CreateOrder: Order created %s ", in.Id)

	return in, nil
}

func (repo *Repository) GetOrder(id string) (*pb.Order, error) {
	if len(id) > 0 {
		for _, o := range repo.orders {
			if o.Id == id {
				return o, nil
			}
		}
	}
	return nil, nil
}

// Применим ли комутативный закон к запросу для выполнения исключений к общей куче
func (repo *Repository) GetOrders(req *pb.Query_Order) ([]*pb.Order, error) {
	heap := repo.orders
	if len(req.Id) > 0 {
		var _heap []*pb.Order
		for _, o := range heap {
			if o.Id == req.Id {
				_heap = append(_heap, o)
				break
			}
		}
		heap = _heap
	}

	if len(req.SendingAddress) > 0 {
		var _heap []*pb.Order
		for _, o := range heap {
			if o.SendingAddress == req.SendingAddress {
				_heap = append(_heap, o)
			}
		}
		heap = _heap
	}

	if len(req.OwnerId) > 0 {
		var _heap []*pb.Order
		for _, o := range heap {
			if o.OwnerId == req.OwnerId {
				_heap = append(_heap, o)
			}
		}
		heap = _heap
	}

	if len(req.BuyCurrencySymbol) > 0 {
		var _heap []*pb.Order
		for _, o := range heap {
			if o.BuyCurrencySymbol == req.BuyCurrencySymbol {
				_heap = append(_heap, o)
			}
		}
		heap = _heap
	}

	if len(req.SellCurrencySymbol) > 0 {
		var _heap []*pb.Order
		for _, o := range heap {
			if o.SellCurrencySymbol == req.SellCurrencySymbol {
				_heap = append(_heap, o)
			}
		}
		heap = _heap
	}

	return heap, nil
}

func (repo *Repository) CreateContract(c *pb.Contract) (*pb.Contract, error) {
	if len(c.Id) > 0 {
		return nil, fmt.Errorf("Repo-CreateContract: Id contract not empty")
	}

	c.Id = uuid.New().String()
	c.CreatedAt = helper.CurrentTimestamp()

	c.Status = &pb.DealStatus{
		Status:    pb.DealStatus_CREATED,
		CreatedAt: c.CreatedAt,
	}

	updated := append(repo.contracts, c)
	repo.contracts = updated
	return c, nil
}

func (repo *Repository) GetContract(id string) (*pb.Contract, error) {
	if len(id) > 0 {
		for _, c := range repo.contracts {
			if c.Id == id {
				return c, nil
			}
		}
	}
	return nil, nil
}

func (repo *Repository) GetContracts(req *pb.Query_Contract) (cs []*pb.Contract, err error) {

	if len(req.Id) > 0 {
		for _, c := range repo.contracts {
			if c.Id == req.Id {
				cs = append(cs, c)
				return cs, err
			}
		}
		return cs, err
	}

	//Пришли tx от seller
	if len(req.SellerSendAddress) > 0 && len(req.BuyerReceiveAddress) > 0 {
		for _, c := range repo.contracts {
			if c.SellerSendAddress == req.SellerSendAddress &&
				c.BuyerReceiveAddress == req.BuyerReceiveAddress {
				if req.Active && c.Available > 0 {
					cs = append(cs, c)
				} else if !req.Active {
					cs = append(cs, c)
				}
			}
		}
		return cs, err
	}

	//пришла tx от buyer
	if len(req.SellerReceiveAddress) > 0 && len(req.BuyerSendAddress) > 0 {
		for _, c := range repo.contracts {
			if c.SellerReceiveAddress == req.SellerReceiveAddress &&
				c.BuyerSendAddress == req.BuyerSendAddress {
				if req.Active && c.Available > 0 {
					cs = append(cs, c)
				} else if !req.Active {
					cs = append(cs, c)
				}
			}
		}
		return cs, err
	}

	//проверка баланса
	if len(req.BuyerSendAddress) > 0 && len(req.SellerSendAddress) > 0 {
		for _, c := range repo.contracts {
			if c.SellerSendAddress == req.SellerSendAddress ||
				c.BuyerSendAddress == req.BuyerSendAddress {
				if req.Active && c.Available > 0 {
					cs = append(cs, c)
				} else if !req.Active {
					cs = append(cs, c)
				}
			}
		}
		return cs, err
	}

	if len(req.UserId) > 0 {
		for _, c := range repo.contracts {
			if c.SellerId == req.UserId ||
				c.BuyerId == req.UserId {
				if req.Active && c.Available > 0 {
					cs = append(cs, c)
				} else if !req.Active {
					cs = append(cs, c)
				}
			}
		}
		return cs, err
	}

	cs = repo.contracts
	return cs, err
}

func (r *Repository) UpdateContractAvailable(c *pb.Contract) (*pb.Contract, error) {
	Mu.Lock()
	defer Mu.Unlock()

	c.Cost = c.Available * c.Price
	log.Println("Repo-UpdateContractAvailable: Contract available updated ")
	return c, nil
}

func (r *Repository) UpdateOrderAvailable(o *pb.Order) (*pb.Order, error) {
	log.Println("Repo-UpdateOrderAvailable: Order available updated ")
	return o, nil
}

func (r *Repository) UpdateOrderStatus(o *pb.Order) (*pb.Order, error) {
	o.Status.CreatedAt = helper.CurrentTimestamp()
	log.Println("Repo-UpdateOrderStatus: Order status updated ")

	return o, nil
}

func (r *Repository) UpdateContractStatus(c *pb.Contract) (*pb.Contract, error) {
	c.Status.CreatedAt = helper.CurrentTimestamp()
	log.Println("Repo-UpdateContractStatus: Contract status updated ")

	return c, nil
}

//CreateWithdrawal(o *pb.WithdrawalOrder) (*pb.WithdrawalOrder, error)
//GetWithdrawal(o *pb.Query_Withdrawal) ([]*pb.WithdrawalOrder, error)
//UpdateWithdrawal(o *pb.WithdrawalOrder) (*pb.WithdrawalOrder, error)

func (r *Repository) CreateWithdrawal(wo *pb.WithdrawalOrder) (*pb.WithdrawalOrder, error) {
	wo.Id = uuid.New().String()
	wo.CreatedAt = helper.CurrentTimestamp()
	wo.Status = &pb.DealStatus{
		Status:    pb.DealStatus_CREATED,
		CreatedAt: wo.CreatedAt,
	}
	update := append(r.withdrawals, wo)
	r.withdrawals = update
	log.Printf("Repo-CreateWithdrawal: Withdrawal order created %s ", wo.Id)
	return wo, nil
}

func (r *Repository) GetWithdrawal(q *pb.Query_Withdrawal) ([]*pb.WithdrawalOrder, error) {
	heap := r.withdrawals

	if len(q.Id) > 0 {
		var result []*pb.WithdrawalOrder
		for _, wo := range heap {
			if wo.Id == q.Id {
				result = append(result, wo)
				break
			}
		}
		heap = result
	}

	if len(q.OwnerId) > 0 {
		var result []*pb.WithdrawalOrder
		for _, wo := range heap {
			if wo.OwnerId == q.OwnerId {
				result = append(result, wo)
			}
		}
		heap = result
	}

	if len(q.OwnerId) > 0 {
		var result []*pb.WithdrawalOrder
		for _, wo := range r.withdrawals {
			if wo.OwnerId == q.OwnerId {
				result = append(result, wo)
			}
		}
		heap = result
	}

	if len(q.PaymentSystem) > 0 {
		var result []*pb.WithdrawalOrder
		for _, wo := range r.withdrawals {
			if wo.PaymentSystem == q.PaymentSystem {
				result = append(result, wo)
			}
		}
		heap = result
	}

	if q.Status != nil {
		var result []*pb.WithdrawalOrder
		for _, wo := range r.withdrawals {
			if wo.Status.Status == q.Status.Status {
				result = append(result, wo)
			}
		}
		heap = result
	}

	return heap, nil
}

func (r *Repository) GetAllWithdrawals() ([]*pb.WithdrawalOrder, error) {
	return r.withdrawals, nil
}

func (r *Repository) UpdateWithdrawal(wo *pb.WithdrawalOrder) (*pb.WithdrawalOrder, error) {
	return wo, nil
}
