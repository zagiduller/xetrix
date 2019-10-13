package payments

import (
	"engine/lib/payments/systems"
	"engine/lib/payments/systems/ethint"
	"engine/lib/services/events"
	"engine/lib/structs"
	"fmt"
	"github.com/gin-gonic/gin/json"
	"github.com/olebedev/config"
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"
)

type System interface {
	Name() string
	Symbol() string
	Currency() *structs.Currency
	AddToWatch(acc *structs.Account)
	GenerateAddress(user *structs.User) (string, uint64) // Нужны отдельные генераторы
	Run(wg *sync.WaitGroup, txpvd chan<- structs.Query_RawTx)
}

// Можно сделать что бы сама валюта отдавала объект System
type CurrencyWithPaySystem interface {
	GetName() string
	GetSymbol() string
	GetPaySystem() System
}

type Registry struct {
	systems    map[string]System
	bus        *events.Bus
	addSysChan chan System
	rawTxCh    chan structs.Query_RawTx
	outPsTxCh  chan structs.Query_RawTx
	ethInOutCh chan ethint.AddressTxsResponse

	HooksMux *http.ServeMux

	et2h *ethint.Ethint
}

func (r *Registry) AddEventBus(bus *events.Bus) {
	r.bus = bus
	r.bus.Subscribe(r,
		&structs.Event{Type: &structs.Event_NewCurrency{}},
		&structs.Event{Type: &structs.Event_NewAccount{}},
		&structs.Event{Type: &structs.Event_ETHtxNeedPrepare{}},
		&structs.Event{Type: &structs.Event_ETHtxPrepared{}},
	)
}

func (r *Registry) Notify(event *structs.Event) {
	r.bus.NewEvent(event)
}

func (r *Registry) Update(event *structs.Event) {
	switch event.Type.(type) {
	case *structs.Event_NewCurrency:
		curr := event.GetNewCurrency().Currency
		if curr.Type == structs.Currency_ETH_CONTRACT_TOKEN {
			r.AddSystem(systems.InitContractSystem(curr, []*structs.Account{}, r.et2h))
		}
	case *structs.Event_NewAccount:
		acc := event.GetNewAccount().Account
		if ps, ok := r.systems[acc.Currency.Symbol]; ok {
			ps.AddToWatch(acc)
		} else {
			log.Printf("Payment registry by %s not found! Account given before PS initialized", acc.Currency.Name)
		}
		//case *structs.Event_ETHtxNeedPrepare: // нужно скинуть на оплату комиссии
		//	needPrepareProccess := event.GetETHtxNeedPrepare()
		//	if needPrepareProccess.Tx.Reason.Status == structs.TxReason_FUND_UNPERFORMED_TX {
		//		// Транзакция придет в whaitRelatedTx() тип RelateTypePrepare
		//		r.et2h.RelatedPrepareTxFromAdmin(needPrepareProccess.Currency, needPrepareProccess.OwnerTo, needPrepareProccess.To, needPrepareProccess.Tx)
		//	}
		//case *structs.Event_ETHtxPrepared: // В сестему пришли eth или токены, нужно переслать их на мастер счет
		//	prepared := event.GetETHtxPrepared()
		//	if prepared.Tx.Reason.Status == structs.TxReason_FUND_PREPARED_TX {
		//		r.et2h.RelatedTxToMaster(prepared.Currency, prepared.OwnerTo, prepared.To, prepared.Tx)
		//	}
	}
}

func InitRegistry(cfg *config.Config, currencies []*structs.Currency, addrMap map[string][]*structs.Account) *Registry {
	r := &Registry{
		systems:    make(map[string]System),
		rawTxCh:    make(chan structs.Query_RawTx, 1),
		addSysChan: make(chan System, 50),
		HooksMux:   http.NewServeMux(),
	}

	findCurr := func(symbol string) *structs.Currency {
		for _, curr := range currencies {
			if curr.Symbol == symbol {
				return curr
			}
		}
		return nil
	}

	//if btcsys, err := systems.InitBitcoinSystem(findCurr("BTC"), addrMap["BTC"]); err == nil {
	//	r.AddSystem(btcsys)
	//}

	r.AddSystem(systems.InitYandexSystem(findCurr("RUB"), addrMap["RUB"], cfg.UString("yandexmoney.secret"), r.HooksMux))

	var mnemonic, derivationTemplate, masterAddress, apikey string

	if cfg != nil {
		mnemonic = cfg.UString("ethereum.mnemonic")
		derivationTemplate = cfg.UString("ethereum.derivationTemplate")
		masterAddress = cfg.UString("ethereum.masterAddress")
		apikey = cfg.UString("ethereum.apikey")
	}

	r.et2h = ethint.NewHD(mnemonic, derivationTemplate, masterAddress, apikey)
	r.AddSystem(systems.InitEthereumSystem(findCurr("ETH"), addrMap["ETH"], r.et2h))
	for _, curr := range currencies {
		if curr.Type == structs.Currency_ETH_CONTRACT_TOKEN {
			r.AddSystem(systems.InitContractSystem(curr, addrMap[curr.Symbol], r.et2h))
		}
	}

	r.registerSystemsInfoHandler()
	return r
}

func (r *Registry) AddSystem(ps System) {
	r.addSysChan <- ps
	log.Printf("System %s added to registry ", ps.Name())
}

func (r *Registry) GenerateAddress(user *structs.User, symbol string) (string, uint64) {
	if gen, ok := r.systems[symbol]; ok {
		return gen.GenerateAddress(user)
	}
	return "", 0
}

func (r *Registry) whaitRawTx() {
	for rawtx := range r.rawTxCh {
		r.Notify(&structs.Event{Type: &structs.Event_PaySystemRawTx{
			PaySystemRawTx: &structs.EventPaySystemRawTx{Raw: &rawtx},
		}})
	}
}

func (r *Registry) waitEthInOutTx() {
	for result := range r.et2h.MainTxsResultCh {
		go r.analyseEthInOutTxs(result)
	}
}

func (r *Registry) analyseEthInOutTxs(result *ethint.AddressTxsResponse) {
	toAddress := strings.ToLower(result.Address)

	for _, tx := range result.Result {
		if conf, err := strconv.ParseUint(tx.Confirmations, 10, 32); err != nil || conf < 6 {
			continue
		}

		inout := &structs.EventEthInOut{}
		inout.Address = result.Address
		inout.TxId = tx.Hash
		inout.Amount = tx.Value
		inout.ContractAddress = tx.ContractAddress
		if toAddress == strings.ToLower(tx.To) {
			inout.Type = structs.EventEthInOut_IN
		} else {
			inout.Type = structs.EventEthInOut_OUT
		}
		r.Notify(&structs.Event{Type: &structs.Event_EthInOut{EthInOut: inout}})
	}
}

func (r *Registry) whaitRelatedTx() {
	for resp := range r.et2h.Sc.RelTxCh {
		fmt.Println("whaitRelatedTx resp getting:", resp)
		if len(resp.Result) > 0 {
			// Ждем хеши отправленных транзакци и обьявляем о получении
			// в зависимости от цели отправленной транзакции. Обрабатывается в ServiceTransactionProcessing
			if resp.RelateType == ethint.RelateTypePrepare {
				r.Notify(&structs.Event{Type: &structs.Event_RelatedPaySystemPrepareTx{
					RelatedPaySystemPrepareTx: &structs.EventRelatedPaySystemTx{TxId: resp.TxId, RelatedId: resp.Result},
				}})
			} else if resp.RelateType == ethint.RelateTypeOut {
				r.Notify(&structs.Event{Type: &structs.Event_RelatedPaySystemOutTx{
					RelatedPaySystemOutTx: &structs.EventRelatedPaySystemTx{TxId: resp.TxId, RelatedId: resp.Result},
				}})
			}
		} else {
			log.Printf("whaitRelatedTx: Receive error from send api: %s ", resp.Error.Message)
		}

	}
}

func (r *Registry) Run() {
	wg := &sync.WaitGroup{}

	go r.whaitRawTx()
	go r.whaitRelatedTx()

	for sys := range r.addSysChan {
		wg.Add(1)

		r.systems[sys.Symbol()] = sys
		r.Notify(&structs.Event{Type: &structs.Event_PaySystemAdded{
			PaySystemAdded: &structs.EventPaySystemAdded{Currency: sys.Currency()},
		}})

		go sys.Run(wg, r.rawTxCh)
	}

	wg.Wait()
}

func (r *Registry) registerSystemsInfoHandler() {
	r.HooksMux.Handle("/systems", http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if req.Method != "GET" {
			w.Write([]byte("Method not allowed"))
			w.WriteHeader(http.StatusNotFound)
			return
		}

		type psInfoRes struct{ Name, Symbol string }

		var psInfo []*psInfoRes

		for _, ps := range r.systems {
			psInfo = append(psInfo, &psInfoRes{ps.Name(), ps.Symbol()})
		}

		result, _ := json.Marshal(psInfo)

		w.Write([]byte(result))
	}))
}

func (r *Registry) GetEthint() *ethint.Ethint {
	return r.et2h
}
