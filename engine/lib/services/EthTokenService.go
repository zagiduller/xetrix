package services

import (
	"context"
	"engine/lib/payments/systems/ethint"
	"engine/lib/services/events"
	"engine/lib/structs"
	"github.com/olebedev/config"
	"log"
	"math/big"
)

type IEthToken interface {
	//UpdateContractStatus(c *pb.Contract) (*pb.Contract, error)
}

type ServiceEthToken struct {
	addresses map[string]*structs.EthAddress
	srvu      *ServiceUser
	bus       *events.Bus
	et2h      *ethint.Ethint
	admDrvt   string
}

func NewEthTokenService(cfg *config.Config, et2h *ethint.Ethint, srvu *ServiceUser) *ServiceEthToken {
	return &ServiceEthToken{
		srvu:      srvu,
		addresses: make(map[string]*structs.EthAddress),
		et2h:      et2h,
		admDrvt:   et2h.GetDerivationPath(0, 0),
	}
}

func (s *ServiceEthToken) AddEventBus(bus *events.Bus) {
	s.bus = bus
	bus.Subscribe(s,
		&structs.Event{Type: &structs.Event_EthInOut{}},
	)
}

func (s *ServiceEthToken) addAddress(acc *structs.Account) {
	if acc.Currency.Type == structs.Currency_ETH_CONTRACT_TOKEN || acc.Currency.Symbol == "ETH" {
		if uresp, err := s.srvu.GetUser(context.Background(), &structs.Query_User{Id: acc.OwnerId}); err == nil && uresp != nil {
			addr := &structs.EthAddress{
				ContractAddress: acc.Currency.ContractId,
				Address:         acc.Address,
				DerivationPath:  s.et2h.GetDerivationPath(acc.Currency.Inc, uresp.Object.Inc),
			}
			s.addresses[acc.Address] = addr
			log.Printf("ServiceEthToken: address added: %s ", addr.Address)
		}
	}
}

func (s *ServiceEthToken) Update(event *structs.Event) {
	switch event.Type.(type) {
	case *structs.Event_NewAccount:
		s.addAddress(event.GetNewAccount().Account)
	case *structs.Event_EthInOut:
		s.calcBalance(event.GetEthInOut())
	}
}

func (s *ServiceEthToken) calcBalance(inout *structs.EventEthInOut) {
	if addr, ok := s.addresses[inout.Address]; ok {
		in, out := new(big.Int), new(big.Int)
		var txs []*structs.EventEthInOut

		if len(inout.ContractAddress) > 0 {
			txs = addr.TokenTxs
		} else {
			txs = addr.EthTxs
		}

		for _, tx := range txs {
			if tx.TxId == inout.TxId {
				return
			}
			if amount, ok := new(big.Int).SetString(tx.Amount, 10); ok {
				if tx.Type == structs.EventEthInOut_IN {
					in = new(big.Int).Add(in, amount)
				} else if tx.Type == structs.EventEthInOut_OUT {
					out = new(big.Int).Add(out, amount)
				}
			}
		}

		if inoutamount, ok := new(big.Int).SetString(inout.Amount, 10); ok {
			if inout.Type == structs.EventEthInOut_IN {
				in = new(big.Int).Add(in, inoutamount)
			} else if inout.Type == structs.EventEthInOut_OUT {
				out = new(big.Int).Add(out, inoutamount)
			}
		}

		amount := new(big.Int).Sub(in, out)

		if len(inout.Token) > 0 {
			addr.TokenBalance = amount.String()
			addr.TokenTxs = append(addr.TokenTxs, inout)
		} else {
			addr.EthBalance = amount.String()
			addr.EthTxs = append(addr.EthTxs, inout)
		}
	}
}

func (s *ServiceEthToken) GetAll(ctx context.Context, empty *structs.Empty) (*structs.Response_EthAddress, error) {
	result := &structs.Response_EthAddress{}

	for _, a := range s.addresses {
		result.Addresses = append(result.Addresses, a)
	}
	return result, nil
}

func (s *ServiceEthToken) FundFee(ctx context.Context, q *structs.Query_FundWithdrawEth) (*structs.Bool, error) {
	if addr, ok := s.addresses[q.Address]; ok {
		log.Printf("ServiceEthToken: FromMaster: master path: %s", s.admDrvt)

	}

	return nil, nil
}

func (s *ServiceEthToken) Withdraw(ctx context.Context, q *structs.Query_FundWithdrawEth) (*structs.Bool, error) {
	return nil, nil
}
