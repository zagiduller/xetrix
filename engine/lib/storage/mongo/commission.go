package mongo

import (
	pb "engine/lib/structs"
)

func (r *Repository) FindAddressCommissions(address string, active bool) ([]*pb.Commission, error) {
	var result []*pb.Commission
	//var orders []*pb.Order
	//var contracts []*pb.Contract
	//
	//if err := r.GetOrders(orders); err != nil {
	//	return nil, fmt.Errorf("Storage-Commission: FindAddressCommissions %s",err)
	//}
	//
	//for _, o := range orders {
	//	if active && o.Available == 0 {
	//		continue
	//	}
	//	if o.Commission != nil && (o.Commission.SendingAddress == address || o.Commission.ReceiveAddress == address) {
	//		result = append(result, o.Commission)
	//	}
	//}
	//for _, c := range contracts{
	//	if active && c.Available == 0 {
	//		continue
	//	}
	//	if c.SellerCommission != nil && (c.SellerCommission.SendingAddress == address || c.SellerCommission.ReceiveAddress == address) {
	//		result = append(result, c.SellerCommission)
	//	}
	//	if c.BuyerCommission != nil && (c.BuyerCommission.SendingAddress == address || c.BuyerCommission.ReceiveAddress == address)  {
	//		result = append(result, c.BuyerCommission)
	//	}
	//
	//}

	return result, nil
}
