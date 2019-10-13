package simple

import (
	pb "engine/lib/structs"
)

//FindAddressCommissions поиск комиссий в ордерах и контрактах к оплате которых привязан адрес
func (r *Repository) FindAddressCommissions(address string, active bool) ([]*pb.Commission, error) {
	var result []*pb.Commission
	for _, o := range r.orders {
		if active && o.Available == 0 {
			continue
		}
		if o.Commission != nil && (o.Commission.SendingAddress == address || o.Commission.ReceiveAddress == address) {
			result = append(result, o.Commission)
		}
	}
	for _, c := range r.contracts {
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
