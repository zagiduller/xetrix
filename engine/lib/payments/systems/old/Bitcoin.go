package old

import (
	"engine/lib/structs"
	"fmt"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/rpcclient"
	"log"
	"sync"
	"time"
)

func InitBitcoinSystem(curr *structs.Currency, accounts []*structs.Account) (*Bitcoin, error) {
	client, err := rpcclient.New(&rpcclient.ConnConfig{
		HTTPPostMode: true,
		DisableTLS:   true,
		Host:         "127.0.0.1:18443",
		User:         "alice",
		Pass:         "123456",
	}, nil)
	if err != nil {
		return nil, fmt.Errorf("InitBitcoinSystem: Error creating new btc client: %v", err)
	}
	log.Printf("Initialize BTC system with len(%d) addresses ", len(accounts))
	watchers := make(map[string]*AddressWatcher)

	for _, acc := range accounts {
		watchers[acc.Address] = NewAddressWatcher(acc)
	}

	btcs := &Bitcoin{
		name:     "Bitcoin",
		conn:     client,
		symbol:   "BTC",
		currency: curr,
		watchers: watchers,
		mu:       &sync.Mutex{},
	}
	return btcs, nil
}

type Bitcoin struct {
	name, symbol string
	currency     *structs.Currency
	conn         *rpcclient.Client
	watchers     map[string]*AddressWatcher
	mu           *sync.Mutex
}

func (b *Bitcoin) Name() string {
	return b.name
}

func (b *Bitcoin) Symbol() string {
	return b.symbol
}

func (b *Bitcoin) Currency() *structs.Currency {
	return b.currency
}

func (b *Bitcoin) GenerateAddress(user *structs.User) string {
	if addr, err := b.conn.GetNewAddress("aaa"); err == nil {
		//b.AddToWatch(addr.String())
		return addr.String()
	}
	return ""
}

func (b *Bitcoin) AddToWatch(acc *structs.Account) {
	log.Printf("%s AddToWatch: %s \n", b.Name(), acc.Address)
	b.mu.Lock()
	if _, ok := b.watchers[acc.Address]; !ok {
		b.watchers[acc.Address] = NewAddressWatcher(acc)
	}
	b.mu.Unlock()
}

func (b *Bitcoin) Run(wg *sync.WaitGroup, txpvd chan<- structs.Query_RawTx) {
	defer wg.Done()
	log.Printf("%s payment system run started!", b.name)

	conn := b.conn

	tick := time.NewTicker(10 * time.Second)
	for range tick.C {
		// List addresses with balances confirmed by at least six blocks, including watch-only addresses:
		receivedTxs, err := conn.ListReceivedByAddress()
		if err == nil {
			for _, recd := range receivedTxs {
				b.mu.Lock()
				if wt, ok := b.watchers[recd.Address]; ok {
					// транзакции адреса
					for _, txId := range recd.TxIDs {

						// детализируем транзакцию что бы получить сумму
						hash, _ := chainhash.NewHashFromStr(txId)
						tx, _ := conn.GetTransaction(hash)

						for _, v := range tx.Details {
							if v.Category == "receive" {
								//Проверяем в wather и отправляем транзакцию в канал
								if wt.isNewPsTx(txId) {
									txpvd <- structs.Query_RawTx{
										FromAddress: "",
										ToAddress:   v.Address,
										Amount:      v.Amount,
										InPStxId:    txId,
									}
								}
							}
						}

					}
				}
				b.mu.Unlock()
			}

		}
	}
}
