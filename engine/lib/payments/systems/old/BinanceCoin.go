package old

import (
	"engine/lib/payments/systems/ethint"
	"engine/lib/structs"
	"github.com/ethereum/go-ethereum/params"
	"log"
	"math/big"
	"strings"
	"sync"
	"time"
)

type BinanceCoin struct {
	contractId   string
	name, symbol string
	ethint       *ethint.Ethint
	watchers     map[string]*AddressWatcher
	mu           *sync.Mutex
	txCh         chan *ethint.AddressTxsResponse
}

func (bc *BinanceCoin) Name() string {
	return bc.name
}

func (bc *BinanceCoin) Symbol() string {
	return bc.symbol
}

func (bc *BinanceCoin) GenerateAddress(user *structs.User) string {
	//return bc.ethint.GenerateAddress(bc.user)
	return ""
}

func (bc *BinanceCoin) AddToWatch(acc *structs.Account) {
	log.Printf("BinanceCoin: %s AddToWatch: %s \n", bc.Name(), acc.Address)
	bc.mu.Lock()
	if _, ok := bc.watchers[acc.Address]; !ok {
		bc.watchers[acc.Address] = NewAddressWatcher(acc)
	}
	bc.mu.Unlock()
}

func InitBinanceystem(accounts []*structs.Account, ethh *ethint.Ethint) *BinanceCoin {
	log.Printf("Initialize BinanceCoin system with len(%d) addresses ", len(accounts))
	watchers := make(map[string]*AddressWatcher)

	for _, acc := range accounts {
		watchers[acc.Address] = NewAddressWatcher(acc)
	}

	bc := &BinanceCoin{
		contractId: "0xB8c77482e45F1F44dE1745F52C74426C631bDD52",
		name:       "BinanceCoin",
		symbol:     "BNB",
		watchers:   watchers,
		mu:         &sync.Mutex{},
		ethint:     ethh,
		txCh:       make(chan *ethint.AddressTxsResponse, 1),
	}
	return bc
}

func (bc *BinanceCoin) Run(wg *sync.WaitGroup, txpvd chan<- structs.Query_RawTx) {
	go bc.ethint.Sc.Run() // Запускается так, потому что лимит на запросы. 5 в секунду
	go bc.ScanAddresses()

	//Слушаем канал транзакций
	for txResp := range bc.txCh {
		go bc.analyseIncomingTxs(txResp, txpvd)
	}
}

// TODO общая для криптовалют - сложность с разницей регистров между адресами из блокчейн-транзакций и теми что в системе и у ватчеров
func (bc *BinanceCoin) analyseIncomingTxs(resps *ethint.AddressTxsResponse, txpvd chan<- structs.Query_RawTx) {
	for _, tx := range resps.Result {
		// Адреса эфира не чувствительны к регистру
		if strings.ToLower(resps.Address) == strings.ToLower(tx.To) { // Пришло на адрес
			aw, ok := bc.watchers[resps.Address]
			if ok && aw.isNewPsTx(tx.Hash) {
				if balance, ok := new(big.Int).SetString(tx.Value, 10); ok {
					ammBc, _ := new(big.Float).Quo(new(big.Float).SetInt(balance), new(big.Float).SetInt(big.NewInt(params.Ether))).Float64()
					log.Printf("BinanceCoin: Txid: %s to: %s amount: %f", tx.Hash, tx.To, ammBc)
					txpvd <- structs.Query_RawTx{
						FromAddress: "",
						ToAddress:   aw.address,
						Amount:      ammBc,
						InPStxId:    tx.Hash,
					}
				}

			}
		}
	}
}

//ScanAddresses Раз в 10 секунд формируем команду на получение транзакций
//по адресу и отправляем команду в очередь
func (bc *BinanceCoin) ScanAddresses() {

	sec10 := time.Tick(10 * time.Second)
	for range sec10 {
		for addr := range bc.watchers {
			bc.ethint.Sc.NewContractTxsCmd(bc.txCh, bc.contractId, addr)
		}
	}
}
