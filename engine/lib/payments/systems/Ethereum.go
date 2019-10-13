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

type Ethereum struct {
	name, symbol string
	ethint       *ethint.Ethint
	watchers     map[string]*AddressWatcher
	mu           *sync.Mutex
	currency     *structs.Currency
	blnCh        chan *ethint.BalancesResponse
	txCh         chan *ethint.AddressTxsResponse
}

func (e *Ethereum) Name() string {
	return e.name
}

func (e *Ethereum) Symbol() string {
	return e.symbol
}

func (e *Ethereum) Currency() *structs.Currency {
	return e.currency
}

func (e *Ethereum) GenerateAddress(user *structs.User) (string, uint64) {
	return e.ethint.GenerateAddress(e.currency, user)
}

func (e *Ethereum) AddToWatch(acc *structs.Account) {
	log.Printf("Ethereum: %s AddToWatch: %s \n", e.Name(), acc.Address)
	e.mu.Lock()
	if _, ok := e.watchers[acc.Address]; !ok {
		e.watchers[acc.Address] = NewAddressWatcher(acc)
	}
	e.mu.Unlock()
}

func InitEthereumSystem(curr *structs.Currency, accounts []*structs.Account, ethh *ethint.Ethint) *Ethereum {
	log.Printf("Initialize ETH system with len(%d) addresses ", len(accounts))
	watchers := make(map[string]*AddressWatcher)

	var blockNumber uint64
	if ethh.Sc.BlockNumber == 0 {
		blockNumber = ethh.Sc.CurrentBlock()
	}

	for _, acc := range accounts {
		watchers[acc.Address] = NewEthAddressWatcher(acc, blockNumber)
	}

	e := &Ethereum{
		name:     "Ethereum",
		symbol:   "ETH",
		watchers: watchers,
		mu:       &sync.Mutex{},
		ethint:   ethh,
		currency: curr,
		blnCh:    make(chan *ethint.BalancesResponse, 1),
		txCh:     make(chan *ethint.AddressTxsResponse, 1),
	}
	return e
}

func (e *Ethereum) Run(wg *sync.WaitGroup, txpvd chan<- structs.Query_RawTx) {
	defer wg.Done()

	go e.ethint.Sc.Run() // Запускается так, потому что лимит на запросы. 5 в секунду
	go e.ethScanBalances()

	// Слушаем канал проверяем изменения балансов кошельков
	go func() {
		for blResp := range e.blnCh {
			e.checkBalancesChange(blResp)
		}
	}()

	//Слушаем канал транзакций
	for txResp := range e.txCh {
		go e.analyseIncomingTxs(txResp, txpvd)
		e.ethint.MainTxsResultCh <- txResp
	}
}

func (e *Ethereum) analyseIncomingTxs(resps *ethint.AddressTxsResponse, txpvd chan<- structs.Query_RawTx) {
	// Пришел ответ на наш запрос
	if aw, ok := e.watchers[resps.Address]; ok && len(resps.Result) > 0 {
		awAddress, adminAddress := strings.ToLower(resps.Address), strings.ToLower(e.ethint.AdminAddress)
		var maxBlock uint64 = 0
		//// в Result транзакции отсортированы от большей к меньшей
		for _, tx := range resps.Result {
			if conf, err := strconv.ParseUint(tx.Confirmations, 10, 32); err != nil || conf < 6 {
				continue
			}

			// Старший номер блока в выборке
			if maxBlock == 0 {
				if bn, err := strconv.ParseUint(tx.BlockNumber, 10, 64); err == nil {
					maxBlock = bn
				}
			}

			// Адреса эфира не чувствительны к регистру
			if awAddress == strings.ToLower(tx.To) && // Пришло на адрес
				strings.ToLower(tx.From) != adminAddress { // Не от Администратора
				if ok && aw.isNewPsTx(tx.Hash) {
					if balance, ok := new(big.Int).SetString(tx.Value, 10); ok {
						ammEth, _ := new(big.Float).Quo(new(big.Float).SetInt(balance), new(big.Float).SetInt(big.NewInt(params.Ether))).Float64()
						log.Printf("Ethereum: analyseIncomingTxs Txid: %s to: %s amount: %f", tx.Hash, tx.To, ammEth)
						blockNumber, _ := strconv.ParseUint(tx.BlockNumber, 10, 32)
						txpvd <- structs.Query_RawTx{
							FromAddress: "",
							ToAddress:   aw.address,
							Amount:      ammEth,
							InPStxId:    tx.Hash,
							BlockNumber: blockNumber,
						}
					}
				}
			} else {
				//TODO можно проверять исходящие транзакции и привязывать к транзакциям в системе
				//c Reason.Status == FUND_WAIT_PREPARE_TX но для ETH это не актуально, актуально для токенов
				//но адреса токенов тут не рассматриваются. Ещё один аргумент в пользу способа "Один счет -- Один человек"

			}
		}

		aw.SetCurrentBlock(maxBlock)
	}

}

func (e *Ethereum) checkBalancesChange(resps *ethint.BalancesResponse) {
	for _, res := range resps.Result {
		if aw, ok := e.watchers[res.Account]; ok {
			if aw.balanceChanged(res.Balance) {
				// изменился баланс. заправшиваем транзакции
				// они вернутся в канал e.ethint.Sc.AddrsTxsRespCh
				e.ethint.Sc.NewAddressTxsCmd(e.txCh, res.Account, aw.GetCurrentBlock())
			}
		}
	}
}

//EthScanBalances Раз в 10 секунд формируем пакеты по 20 штук из всех адресов
//и отправляем команду в очередь
func (e *Ethereum) ethScanBalances() {
	var wtchPackPool = sync.Pool{}
	// получить буфер
	getPackPool := func() (b []string) {
		ifc := wtchPackPool.Get()
		if ifc != nil {
			b = ifc.([]string)
		}
		return
	}
	// вернуть буфер
	putPackPool := func(b []string) {
		b = b[:0] // сброс
		wtchPackPool.Put(b)
	}

	sec10 := time.Tick(30 * time.Second)
	for range sec10 {
		wtchs := make([]string, 0, len(e.watchers))
		for addr := range e.watchers {
			if addr == e.ethint.AdminAddress { // Не проверяем баланс счета администратора
				continue
			}
			wtchs = append(wtchs, addr)
		}
		for len(wtchs) > 0 {
			pack := getPackPool()
			if len(wtchs) > 20 {
				pack, wtchs = wtchs[:20], wtchs[20:]
			} else {
				pack, wtchs = wtchs, wtchs[:0]
			}

			q := strings.Join(pack, ",")
			e.ethint.Sc.NewBalancesCmd(e.blnCh, q)

			putPackPool(pack)
		}
	}
}
