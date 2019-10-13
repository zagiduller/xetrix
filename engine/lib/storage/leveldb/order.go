package leveldb

import (
	"engine/lib/helper"
	pb "engine/lib/structs"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/google/uuid"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/util"
	"log"
)

func (repo *Repository) CreateOrder(in *pb.Order) (*pb.Order, error) {
	in.Id = uuid.New().String()
	in.CreatedAt = helper.CurrentTimestamp()
	in.Status = &pb.DealStatus{
		Status:    pb.DealStatus_CREATED,
		CreatedAt: in.CreatedAt,
	}

	byted, err := proto.Marshal(in)
	if err != nil {
		return nil, fmt.Errorf("Repo-CreateOrder: %s", err)
	}
	batch := new(leveldb.Batch)

	batch.Put([]byte("object-"+in.Id), byted)
	batch.Put([]byte(("sendingAddress-" + in.SendingAddress + "-" + in.Id)), []byte(in.Id))
	batch.Put([]byte(("ownerId-" + in.OwnerId + "-" + in.Id)), []byte(in.Id))
	batch.Put([]byte(("currency-" + in.BuyCurrencySymbol + "-buy-" + in.Id)), []byte(in.Id))
	batch.Put([]byte(("currency-" + in.SellCurrencySymbol + "-sell-" + in.Id)), []byte(in.Id))

	if err := repo.orders.Write(batch, nil); err != nil {
		return nil, fmt.Errorf("Repo-CreateOrder: %s", err)
	}

	log.Printf("Repo-CreateOrder: Order created %s ", in.Id)

	return in, nil
}

func (repo *Repository) GetOrder(id string) (*pb.Order, error) {
	if len(id) > 0 {
		obj := new(pb.Order)
		byted, err := repo.orders.Get([]byte("object-"+id), nil)
		if err != nil {
			return nil, fmt.Errorf("Repo-GetOrder: %s", err)
		}
		if err := proto.Unmarshal(byted, obj); err != nil {
			return nil, fmt.Errorf("Repo-GetOrder: %s", err)
		}
		return obj, nil
	}
	return nil, nil
}

func (repo *Repository) GetOrders(req *pb.Query_Order) ([]*pb.Order, error) {
	var heap []*pb.Order

	iter := repo.orders.NewIterator(util.BytesPrefix([]byte("object-")), nil)
	for iter.Next() {
		obj := new(pb.Order)
		if err := proto.Unmarshal(iter.Value(), obj); err == nil {
			heap = append(heap, obj)
		} else {
			fmt.Printf("Repo-GetOrders: Error. %s", err)
		}
	}

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

	byted, err := proto.Marshal(c)
	if err != nil {
		return nil, fmt.Errorf("Repo-CreateContract: %s", err)
	}
	batch := new(leveldb.Batch)

	batch.Put([]byte("object-"+c.Id), byted)
	batch.Put([]byte(("orderId-" + c.OrderId + "-" + c.Id)), []byte(c.Id))
	batch.Put([]byte(("userId-" + c.SellerId + "-seller-" + c.Id)), []byte(c.Id))
	batch.Put([]byte(("userId-" + c.BuyerId + "-buyer-" + c.Id)), []byte(c.Id))
	batch.Put([]byte(("address-" + c.SellerSendAddress + "-sellerSend-" + c.Id)), []byte(c.Id))
	batch.Put([]byte(("address-" + c.SellerReceiveAddress + "-sellerReceive-" + c.Id)), []byte(c.Id))
	batch.Put([]byte(("address-" + c.BuyerSendAddress + "-buyerSend-" + c.Id)), []byte(c.Id))
	batch.Put([]byte(("address-" + c.BuyerReceiveAddress + "-buyerReceive-" + c.Id)), []byte(c.Id))

	// Индексы для простоты GetContracts
	batch.Put([]byte(("sellerSendBuyerReceive-" + c.SellerSendAddress + "-" + c.BuyerReceiveAddress + "-" + c.Id)), []byte(c.Id))
	batch.Put([]byte(("sellerReceiveBuyerSend-" + c.SellerReceiveAddress + "-" + c.BuyerSendAddress + "-" + c.Id)), []byte(c.Id))

	if err := repo.contracts.Write(batch, nil); err != nil {
		return nil, fmt.Errorf("Repo-CreateContract: %s", err)
	}

	return c, nil
}

func (repo *Repository) GetContract(id string) (*pb.Contract, error) {
	if len(id) > 0 {
		obj := new(pb.Contract)
		byted, err := repo.contracts.Get([]byte("object-"+id), nil)
		if err != nil {
			return nil, fmt.Errorf("Repo-GetContract: %s", err)
		}
		if err := proto.Unmarshal(byted, obj); err != nil {
			return nil, fmt.Errorf("Repo-GetContract: %s", err)
		}
		return obj, nil
	}
	return nil, nil
}

func (repo *Repository) GetContracts(req *pb.Query_Contract) (cs []*pb.Contract, err error) {
	if len(req.Id) > 0 {
		if c, err := repo.GetContract(req.Id); err != nil {
			cs = append(cs, c)
			return cs, nil
		}
		return cs, fmt.Errorf("Repo-GetContracts: req.Id: %s ", err)
	}

	//Такой запрос характерен для поиска контрактов по которому проходит транзакция.
	//Ищется контракт по адресу отправки и адресу получения
	if len(req.SellerSendAddress) > 0 && len(req.BuyerReceiveAddress) > 0 {
		prefix := []byte(("sellerSendBuyerReceive-" + req.SellerSendAddress + "-" + req.BuyerReceiveAddress))
		iter := repo.contracts.NewIterator(util.BytesPrefix(prefix), nil)
		for iter.Next() {
			if c, err := repo.GetContract(string(iter.Value())); err == nil {
				if req.Active && c.Available > 0 {
					cs = append(cs, c)
				} else if !req.Active {
					cs = append(cs, c)
				}
			} else {
				log.Printf("Repo-GetContracts: SellerSendAddress || BuyerReceiveAddress. %s", err)
			}
		}

		return cs, err
	}

	//пришла tx от buyer
	if len(req.SellerReceiveAddress) > 0 && len(req.BuyerSendAddress) > 0 {
		prefix := []byte(("sellerReceiveBuyerSend-" + req.SellerReceiveAddress + "-" + req.BuyerSendAddress))
		iter := repo.contracts.NewIterator(util.BytesPrefix(prefix), nil)
		for iter.Next() {
			if c, err := repo.GetContract(string(iter.Value())); err == nil {
				if req.Active && c.Available > 0 {
					cs = append(cs, c)
				} else if !req.Active {
					cs = append(cs, c)
				}
			} else {
				log.Printf("Repo-GetContracts: SellerReceiveAddress || BuyerSendAddress. %s", err)
			}
		}
		return cs, err
	}

	//проверка баланса
	if len(req.BuyerSendAddress) > 0 && len(req.SellerSendAddress) > 0 {
		prefix := []byte(("address-" + req.BuyerSendAddress + "-buyerSend"))
		iter := repo.contracts.NewIterator(util.BytesPrefix(prefix), nil)
		for iter.Next() {
			if c, err := repo.GetContract(string(iter.Value())); err == nil {
				if req.Active && c.Available > 0 {
					cs = append(cs, c)
				} else if !req.Active {
					cs = append(cs, c)
				}
			} else {
				log.Printf("Repo-GetContracts: BuyerSendAddress || SellerSendAddress. %s", err)
			}
		}

		prefix = []byte(("address-" + req.SellerSendAddress + "-sellerSend"))
		iter = repo.contracts.NewIterator(util.BytesPrefix(prefix), nil)
		for iter.Next() {
			if c, err := repo.GetContract(string(iter.Value())); err == nil {
				if req.Active && c.Available > 0 {
					cs = append(cs, c)
				} else if !req.Active {
					cs = append(cs, c)
				}
			} else {
				log.Printf("Repo-GetContracts: All. %s", err)
			}
		}

		return cs, err
	}

	if len(req.UserId) > 0 {
		prefix := []byte(("userId-" + req.UserId))
		iter := repo.contracts.NewIterator(util.BytesPrefix(prefix), nil)
		for iter.Next() {
			if c, err := repo.GetContract(string(iter.Value())); err == nil {
				if req.Active && c.Available > 0 {
					cs = append(cs, c)
				} else if !req.Active {
					cs = append(cs, c)
				}
			} else {
				log.Printf("Repo-GetContracts: UserId. %s", err)
			}
		}
		return cs, err
	}

	/// AAAAAAAAAa
	prefix := []byte(("object-"))
	iter := repo.contracts.NewIterator(util.BytesPrefix(prefix), nil)
	for iter.Next() {
		obj := new(pb.Contract)
		if err := proto.Unmarshal(iter.Value(), obj); err != nil {
			return nil, fmt.Errorf("Repo-GetContract: %s", err)
		}
		if req.Active && obj.Available > 0 {
			cs = append(cs, obj)
		} else if !req.Active {
			cs = append(cs, obj)
		}

	}

	return cs, err
}

func (r *Repository) UpdateContractAvailable(c *pb.Contract) (*pb.Contract, error) {
	c.Cost = c.Available * c.Price
	log.Println("Repo-UpdateContractAvailable: Contract available updated ")

	byted, err := proto.Marshal(c)
	if err != nil {
		return nil, fmt.Errorf("Repo-UpdateContractAvailable: %s", err)
	}
	if err := r.contracts.Put([]byte("object-"+c.Id), byted, nil); err != nil {
		return nil, fmt.Errorf("Repo-UpdateContractAvailable: %s", err)
	}

	return c, nil
}

func (r *Repository) UpdateOrderAvailable(o *pb.Order) (*pb.Order, error) {
	byted, err := proto.Marshal(o)
	if err != nil {
		return nil, fmt.Errorf("Repo-UpdateOrderAvailable: %s", err)
	}
	if err := r.orders.Put([]byte("object-"+o.Id), byted, nil); err != nil {
		return nil, fmt.Errorf("Repo-UpdateOrderAvailable: %s", err)
	}
	log.Println("Repo-UpdateOrderAvailable: Order available updated ")
	return o, nil
}

func (r *Repository) UpdateOrderStatus(o *pb.Order) (*pb.Order, error) {
	o.Status.CreatedAt = helper.CurrentTimestamp()

	byted, err := proto.Marshal(o)
	if err != nil {
		return nil, fmt.Errorf("Repo-UpdateOrderStatus: %s", err)
	}
	if err := r.orders.Put([]byte("object-"+o.Id), byted, nil); err != nil {
		return nil, fmt.Errorf("Repo-UpdateOrderStatus: %s", err)
	}

	log.Println("Repo-UpdateOrderStatus: Order status updated ")
	return o, nil
}

func (r *Repository) UpdateContractStatus(c *pb.Contract) (*pb.Contract, error) {
	c.Status.CreatedAt = helper.CurrentTimestamp()

	byted, err := proto.Marshal(c)
	if err != nil {
		return nil, fmt.Errorf("Repo-UpdateContractAvailable: %s", err)
	}
	if err := r.contracts.Put([]byte("object-"+c.Id), byted, nil); err != nil {
		return nil, fmt.Errorf("Repo-UpdateContractAvailable: %s", err)
	}

	log.Println("Repo-UpdateContractStatus: Contract status updated ")

	return c, nil
}

func (r *Repository) CreateWithdrawal(wo *pb.WithdrawalOrder) (*pb.WithdrawalOrder, error) {
	wo.Id = uuid.New().String()
	wo.CreatedAt = helper.CurrentTimestamp()
	wo.Status = &pb.DealStatus{
		Status:    pb.DealStatus_CREATED,
		CreatedAt: wo.CreatedAt,
	}

	byted, err := proto.Marshal(wo)
	if err != nil {
		return nil, fmt.Errorf("Repo-CreateContract: %s", err)
	}
	batch := new(leveldb.Batch)

	batch.Put([]byte("object-"+wo.Id), byted)

	if err := r.withdrawals.Write(batch, nil); err != nil {
		return nil, fmt.Errorf("Repo-CreateWithdrawal: %s", err)
	}

	log.Printf("Repo-CreateWithdrawal: Withdrawal order created %s ", wo.Id)
	return wo, nil
}

func (r *Repository) GetWithdrawal(q *pb.Query_Withdrawal) ([]*pb.WithdrawalOrder, error) {
	heap, err := r.GetAllWithdrawals()
	if err != nil {
		return nil, fmt.Errorf("Repo-GetWithdrawal: %s", err)
	}

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

	if len(q.PaymentSystem) > 0 {
		var result []*pb.WithdrawalOrder
		for _, wo := range heap {
			if wo.PaymentSystem == q.PaymentSystem {
				result = append(result, wo)
			}
		}
		heap = result
	}

	if q.Status != nil {
		var result []*pb.WithdrawalOrder
		for _, wo := range heap {
			if wo.Status.Status == q.Status.Status {
				result = append(result, wo)
			}
		}
		heap = result
	}

	return heap, nil
}

func (r *Repository) GetAllWithdrawals() ([]*pb.WithdrawalOrder, error) {
	var result []*pb.WithdrawalOrder
	prefix := []byte(("object-"))
	iter := r.withdrawals.NewIterator(util.BytesPrefix(prefix), nil)
	for iter.Next() {
		obj := new(pb.WithdrawalOrder)
		if err := proto.Unmarshal(iter.Value(), obj); err != nil {
			return nil, fmt.Errorf("Repo-GetAllWithdrawals: %s", err)
		}

		result = append(result, obj)
	}

	return result, nil
}

func (r *Repository) UpdateWithdrawal(wo *pb.WithdrawalOrder) (*pb.WithdrawalOrder, error) {
	byted, err := proto.Marshal(wo)
	if err != nil {
		return nil, fmt.Errorf("Repo-UpdateWithdrawal: %s", err)
	}
	if err := r.withdrawals.Put([]byte("object-"+wo.Id), byted, nil); err != nil {
		return nil, fmt.Errorf("Repo-UpdateWithdrawal: %s", err)
	}

	return wo, nil
}

/// DNKNW

func (r *Repository) FindAddressCommissions(address string, active bool) ([]*pb.Commission, error) {
	var result []*pb.Commission

	orders, err := r.GetOrders(&pb.Query_Order{})
	if err != nil {
		return nil, fmt.Errorf("Repo-FindAddressCommissions: %s", err)
	}

	for _, o := range orders {
		if active && o.Available == 0 {
			continue
		}
		if o.Commission != nil && (o.Commission.SendingAddress == address || o.Commission.ReceiveAddress == address) {
			result = append(result, o.Commission)
		}
	}

	contracts, err := r.GetContracts(&pb.Query_Contract{})
	if err != nil {
		return nil, fmt.Errorf("Repo-FindAddressCommissions: %s", err)
	}

	for _, c := range contracts {
		if active && c.Available == 0 {
			continue
		}
		if c.SellerCommission != nil && (c.SellerCommission.SendingAddress == address || c.SellerCommission.ReceiveAddress == address) {
			result = append(result, c.SellerCommission)
		}
		if c.BuyerCommission != nil && (c.BuyerCommission.SendingAddress == address || c.BuyerCommission.ReceiveAddress == address) {
			result = append(result, c.BuyerCommission)
		}

	}
	return result, nil
}
