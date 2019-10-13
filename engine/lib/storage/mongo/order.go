package mongo

import (
	"engine/lib/helper"
	pb "engine/lib/structs"
	"fmt"
	"github.com/google/uuid"
	"gopkg.in/mgo.v2/bson"
	"log"
)

func (r *Repository) CreateOrder(o *pb.Order) (*pb.Order, error) {
	o.Id = uuid.New().String()
	o.CreatedAt = helper.CurrentTimestamp()
	o.Status = &pb.DealStatus{
		Status:    pb.DealStatus_CREATED,
		CreatedAt: o.CreatedAt,
	}
	if err := r.orders.Insert(&o); err != nil {
		return nil, fmt.Errorf("Storage-CreateOrder: %s ", err)
	}
	log.Printf("Storage-CreateOrder: Order added %s \n", o.Id)

	return o, nil
}

func (r *Repository) GetOrder(id string) (*pb.Order, error) {
	var o pb.Order
	if len(id) > 0 {
		if err := r.orders.Find(bson.M{"id": id}).One(&o); err != nil {
			return nil, fmt.Errorf("Storage-GetOrder: %s ", err)
		}
	}
	return &o, nil
}

func (r *Repository) GetOrders(req *pb.Query_Order) ([]*pb.Order, error) {
	var orders []*pb.Order
	qbson := make(bson.M)

	if len(req.Id) > 0 {
		qbson["id"] = req.Id
	}

	if len(req.SendingAddress) > 0 {
		qbson["sendingaddress"] = req.SendingAddress
	}

	if len(req.OwnerId) > 0 {
		qbson["ownerid"] = req.OwnerId
	}

	if len(req.BuyCurrencySymbol) > 0 {
		qbson["buycurrencysymbol"] = req.BuyCurrencySymbol
	}

	if len(req.SellCurrencySymbol) > 0 {
		qbson["sellcurrencysymbol"] = req.SellCurrencySymbol
	}

	if err := r.orders.Find(qbson).All(&orders); err != nil {
		return nil, fmt.Errorf("Storage-GetOrders: %s ", err)
	}

	return orders, nil
}

func (r *Repository) CreateContract(c *pb.Contract) (*pb.Contract, error) {
	c.Id = uuid.New().String()
	c.CreatedAt = helper.CurrentTimestamp()

	c.Status = &pb.DealStatus{
		Status:    pb.DealStatus_CREATED,
		CreatedAt: c.CreatedAt,
	}

	if err := r.contracts.Insert(&c); err != nil {
		return nil, fmt.Errorf("Storage-CreateContract: %s ", err)
	}
	log.Printf("Storage-CreateContract: Contract added %+v \n", c)

	return c, nil
}

func (r *Repository) GetContract(id string) (*pb.Contract, error) {
	var c pb.Contract
	if len(id) > 0 {
		if err := r.contracts.Find(bson.M{"id": id}).One(&c); err != nil {
			return nil, fmt.Errorf("Storage-GetContract: %s ", err)
		}
	}
	return &c, nil
}

func (r *Repository) GetContracts(req *pb.Query_Contract) (cs []*pb.Contract, err error) {
	qbson := make(bson.M)
	if len(req.Id) > 0 {
		qbson["id"] = req.Id
	}

	//Пришли tx от seller
	if len(req.SellerSendAddress) > 0 && len(req.BuyerReceiveAddress) > 0 {
		qbson["sellersendaddress"], qbson["buyerreceiveaddress"] = req.SellerSendAddress, req.BuyerReceiveAddress
	}

	//пришла tx от buyer
	if len(req.SellerReceiveAddress) > 0 && len(req.BuyerSendAddress) > 0 {
		qbson["sellerreceiveaddress"], qbson["buyersendaddress"] = req.SellerReceiveAddress, req.BuyerSendAddress
	}

	//проверка баланса
	if len(req.BuyerSendAddress) > 0 && len(req.SellerSendAddress) > 0 {
		qbson["$or"] = []bson.M{
			{"buyersendaddress": req.BuyerSendAddress},
			{"sellersendaddress": req.SellerSendAddress},
		}
	}

	if len(req.UserId) > 0 {
		qbson["$or"] = []bson.M{
			{"buyerid": req.UserId},
			{"sellerid": req.UserId},
		}
	}

	if req.Active {
		qbson["available"] = bson.M{"$gt": 0}
	}

	if err = r.contracts.Find(qbson).All(&cs); err != nil {
		return nil, fmt.Errorf("Storage-GetContracts: %s ", err)
	}

	return cs, nil
}

func (r *Repository) UpdateContractAvailable(c *pb.Contract) (*pb.Contract, error) {
	c.Cost = c.Available * c.Price
	if err := r.contracts.Update(bson.M{"id": c.Id}, bson.M{"$set": bson.M{
		"available":        c.Available,
		"cost":             c.Cost,
		"status":           c.Status,
		"sellercommission": c.SellerCommission,
		"buyercommission":  c.BuyerCommission,
	},
	}); err != nil {
		return nil, fmt.Errorf("Storage-UpdateContractAvailable: %s ", err)
	}
	log.Println("Storage-UpdateContractAvailable: Contract available and status updated ")

	return c, nil
}

func (r *Repository) UpdateOrderAvailable(o *pb.Order) (*pb.Order, error) {
	if err := r.orders.Update(bson.M{"id": o.Id}, bson.M{"$set": bson.M{
		"available":  o.Available,
		"status":     o.Status,
		"commission": o.Commission,
	},
	}); err != nil {
		return nil, fmt.Errorf("Storage-UpdateOrderAvailable: %s ", err)
	}
	log.Println("Storage-UpdateOrderAvailable: Order available and status updated ")

	return o, nil
}

func (r *Repository) UpdateOrderStatus(o *pb.Order) (*pb.Order, error) {
	o.Status.CreatedAt = helper.CurrentTimestamp()
	if err := r.orders.Update(bson.M{"id": o.Id}, bson.M{"$set": bson.M{"status": o.Status}}); err != nil {
		return nil, fmt.Errorf("Storage-UpdateOrderStatus: %s ", err)
	}
	log.Println("Storage-UpdateOrderStatus: Order status updated ")

	return o, nil
}

func (r *Repository) UpdateContractStatus(c *pb.Contract) (*pb.Contract, error) {
	c.Status.CreatedAt = helper.CurrentTimestamp()
	if err := r.orders.Update(bson.M{"id": c.Id}, bson.M{"$set": bson.M{"status": c.Status}}); err != nil {
		return nil, fmt.Errorf("Storage-UpdateContractStatus: %s ", err)
	}
	log.Println("Storage-UpdateContractStatus: Contract status updated ")

	return c, nil
}
