package systems

import (
	"engine/lib/payments/systems/ethint"
	"engine/lib/structs"
	"github.com/ethereum/go-ethereum/params"
	"log"
	"math/big"
	"strconv"
	"strings"
	"sync"
	"time"
)

type EthereumContract struct {
	contractId   string
	name, symbol string
	ethint       *ethint.Ethint
	currency     *structs.Currency
	watchers     map[string]*AddressWatcher
	mu           *sync.Mutex
	txCh         chan *ethint.AddressTxsResponse
}

func (c *EthereumContract) Name() string {
	return c.name
}

func (c *EthereumContract) Symbol() string {
	return c.symbol
}

func (c *EthereumContract) Currency() *structs.Currency {
	return c.currency
}

func (c *EthereumContract) GenerateAddress(user *structs.User) (string, uint64) {
	return c.ethint.GenerateAddress(c.currency, user)
}

func (c *EthereumContract) AddToWatch(acc *structs.Account) {
	log.Printf("%s: AddToWatch: %s \n", c.Name(), acc.Address)
	c.mu.Lock()
	if _, ok := c.watchers[acc.Address]; !ok {
		c.watchers[acc.Address] = NewAddressWatcher(acc)
	}
	c.mu.Unlock()
}

func InitContractSystem(curr *structs.Currency, accounts []*structs.Account, ethh *ethint.Ethint) *EthereumContract {
	log.Printf("Initialize %s system with len(%d) accounts ", curr.Name, len(accounts))
	watchers := make(map[string]*AddressWatcher)

	var blockNumber uint64
	if ethh.Sc.BlockNumber == 0 {
		blockNumber = ethh.Sc.CurrentBlock()
	}

	for _, acc := range accounts {
		watchers[acc.Address] = NewEthAddressWatcher(acc, blockNumber)
	}

	c := &EthereumContract{
		contractId: curr.ContractId,
		name:       curr.Name,
		symbol:     curr.Symbol,
		currency:   curr,
		watchers:   watchers,
		mu:         &sync.Mutex{},
		ethint:     ethh,
		txCh:       make(chan *ethint.AddressTxsResponse, 1),
	}
	return c
}

func (c *EthereumContract) Run(wg *sync.WaitGroup, txpvd chan<- structs.Query_RawTx) {
	defer wg.Done()

	go c.ethint.Sc.Run() // Запускается так, потому что лимит на запросы. 5 в секунду
	go c.ScanAddresses()

	//Слушаем канал транзакций
	for txResp := range c.txCh {
		go c.analyseIncomingTxs(txResp, txpvd)
		c.ethint.MainTxsResultCh <- txResp
	}
}

func (c *EthereumContract) analyseIncomingTxs(resps *ethint.AddressTxsResponse, txpvd chan<- structs.Query_RawTx) {
	if aw, ok := c.watchers[resps.Address]; ok && len(resps.Result) > 0 {
		awAddress, adminAddress := strings.ToLower(resps.Address), strings.ToLower(c.ethint.AdminAddress)
		var maxBlock uint64 = 0

		for _, tx := range resps.Result {
			if conf, err := strconv.ParseUint(tx.Confirmations, 10, 32); err != nil || conf < 5 {
				continue
			}

			// Старший номер блока в выборке
			if maxBlock == 0 {
				if bn, err := strconv.ParseUint(tx.BlockNumber, 10, 64); err == nil {
					maxBlock = bn
				}
			}

			// Адреса эфира не чувствительны к регистру
			if awAddress == strings.ToLower(tx.To) && //чего? Пришло на адрес
				strings.ToLower(tx.From) != adminAddress { // Не от Администратора
				if ok && aw.isNewPsTx(tx.Hash) {
					if balance, ok := new(big.Int).SetString(tx.Value, 10); ok {
						ammBc, _ := new(big.Float).Quo(new(big.Float).SetInt(balance), new(big.Float).SetInt(big.NewInt(params.Ether))).Float64()
						log.Printf("Contract %s: Txid: %s to: %s amount: %f", c.Name(), tx.Hash, tx.To, ammBc)
						blockNumber, _ := strconv.ParseUint(tx.BlockNumber, 10, 32)

						txpvd <- structs.Query_RawTx{
							FromAddress: "",
							ToAddress:   aw.address,
							Amount:      ammBc,
							InPStxId:    tx.Hash,
							BlockNumber: blockNumber,
						}
					}

				}
			}
		}
		// Запоминаем номер старшего блока
		aw.SetCurrentBlock(maxBlock)

	}
}

//ScanAddresses Раз в 10 секунд формируем команду на получение транзакций
//по адресу и отправляем команду в очередь
func (c *EthereumContract) ScanAddresses() {
	sec10 := time.Tick(60 * time.Second)
	for range sec10 {
		for addr, w := range c.watchers {
			c.ethint.Sc.NewContractTxsCmd(c.txCh, c.contractId, addr, w.GetCurrentBlock())
		}
	}
}
