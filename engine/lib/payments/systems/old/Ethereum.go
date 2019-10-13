package old

import (
	"context"
	"encoding/json"
	"engine/lib/payments/systems/eth"
	"engine/lib/payments/systems/ethapi"
	"engine/lib/payments/systems/ethint"
	"engine/lib/structs"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/params"
	"github.com/gophergala2016/etherapis/etherapis/Godeps/_workspace/src/github.com/ethereum/go-ethereum/rpc"
	"log"
	"math/big"
	"strings"
	"sync"
	"time"
)

//Блоки [0, 1.. N]
//В них есть транзакции на счета
//
//currentBlockNumber записать последний проверенный номер блока
//Таймер на проверку блоков  blockCount = endBlockNumber(ebn) - currentBlockNumber(cbn)
//Проверить транзакции в блоках на предмет соответствия адресов счетам внутри системы
//Инициировать создание транзакции пополнения внутри системы и записать блокчейн хеш транзакции
//Проверить отсутствие транзакции с таким хешем
//Записать и подтвердить транзакцию

type Ethereum struct {
	name, symbol string
	blockChain   *eth.EthBlockchainScanner
	ethint       *ethint.Ethint
	watchers     map[string]*AddressWatcher
	mu           *sync.Mutex
}

func InitEthereumSystem(addresses []string, blockChain *eth.EthBlockchainScanner) *Ethereum {
	log.Printf("Initialize ETH system with len(%d) addresses ", len(addresses))
	watchers := make(map[string]*AddressWatcher)

	for _, addr := range addresses {
		watchers[addr] = NewAddressWatcher(addr)
	}

	e := &Ethereum{
		name:       "Ethereum",
		symbol:     "ETH",
		watchers:   watchers,
		mu:         &sync.Mutex{},
		blockChain: blockChain,
	}
	return e
}

func (e *Ethereum) Name() string {
	return e.name
}

func (e *Ethereum) Symbol() string {
	return e.symbol
}

func (e *Ethereum) AddToWatch(addr string) {
	log.Printf("%s AddToWatch: %s \n", e.Name(), addr)
	e.mu.Lock()
	if _, ok := e.watchers[addr]; !ok {
		e.watchers[addr] = NewAddressWatcher(addr)
	}
	e.mu.Unlock()
}

func (e *Ethereum) GenerateAddress() string {
	return e.blockChain.GenerateAddress()
}

func (e *Ethereum) Run(wg *sync.WaitGroup, txpvd chan<- structs.Query_RawTx) {
	//
}

func (e *Ethereum) Run3(wg *sync.WaitGroup, txpvd chan<- structs.Query_RawTx) {
	defer wg.Done()

	headers := make(chan *types.Header)

	_, err := e.blockChain.Client.SubscribeNewHead(context.Background(), headers)
	if err != nil {
		log.Printf("Ethereum. Run: %s", err)
		return
	}

	for header := range headers {
		block, err := e.blockChain.Client.BlockByHash(context.Background(), header.Hash())
		if err != nil {
			log.Printf("Ethereum. Run: ", err)
			return
		}

		txs := block.Transactions()
		for _, tx := range txs {
			if tx.To() != nil && tx.Value().Cmp(new(big.Int).SetInt64(0)) > 0 {
				addr, _ := tx.To().MarshalText()
				if aw, ok := e.watchers[string(addr)]; ok {
					if aw.isNewPsTx(tx.Hash().String()) {
						ammEth, _ := new(big.Float).Quo(new(big.Float).SetInt(tx.Value()), new(big.Float).SetInt(big.NewInt(params.Ether))).Float64()
						log.Printf("Txid: %s to: %s amount: %f", tx.Hash().String(), tx.To().String(), ammEth)
						txpvd <- structs.Query_RawTx{
							FromAddress:     "",
							ToAddress:       aw.address,
							Amount:          ammEth,
							PaymentSystemID: tx.Hash().String(),
						}
					}
				}
			}
		}
	}

}

func (e *Ethereum) Run2(wg *sync.WaitGroup, txpvd chan<- structs.Query_RawTx) {
	defer wg.Done()
	log.Printf("%s payment system run started!", e.name)

	txCh := make(chan *types.Transaction)

	e.blockChain.AddTxCh(txCh)

	for tx := range txCh {
		addr, _ := tx.To().MarshalText()
		if aw, ok := e.watchers[string(addr)]; ok {
			if aw.isNewPsTx(tx.Hash().String()) {
				ammEth, _ := new(big.Float).Quo(new(big.Float).SetInt(tx.Value()), new(big.Float).SetInt(big.NewInt(params.Ether))).Float64()
				log.Printf("Txid: %s to: %s amount: %f", tx.Hash().String(), tx.To().String(), ammEth)
				txpvd <- structs.Query_RawTx{
					FromAddress:     "",
					ToAddress:       aw.address,
					Amount:          ammEth,
					PaymentSystemID: tx.Hash().String(),
				}
			}
		}
	}
}

type txdata struct {
	AccountNonce hexutil.Uint64  `json:"nonce"    gencodec:"required"`
	Price        *hexutil.Big    `json:"gasPrice" gencodec:"required"`
	GasLimit     hexutil.Uint64  `json:"gas"      gencodec:"required"`
	Recipient    *common.Address `json:"to"       rlp:"nil"`
	Amount       *hexutil.Big    `json:"value"    gencodec:"required"`
	Payload      hexutil.Bytes   `json:"input"    gencodec:"required"`
	V            *hexutil.Big    `json:"v" gencodec:"required"`
	R            *hexutil.Big    `json:"r" gencodec:"required"`
	S            *hexutil.Big    `json:"s" gencodec:"required"`
	Hash         *common.Hash    `json:"hash" rlp:"-"`
}

type EthBlockchainScanner struct {
	conn     *ethapi.API
	TxCh     chan<- *txdata
	CtrEvChs []chan<- *txdata
	mu       *sync.Mutex
}

func NewEthBlockchainScanner() *EthBlockchainScanner {

	client, err := rpc.NewIPCClient("/home/arthur/.ethereum/dev.ipc")
	if err != nil {
		log.Println("InitEthereumSystem")
		panic(err)
	}
	api := ethapi.NewApi(client)

	return &EthBlockchainScanner{
		conn: api,
		mu:   &sync.Mutex{},
	}
}

func (bs *EthBlockchainScanner) GenerateAddress() string {
	acc, err := bs.conn.Request("personal_newAccount", []interface{}{""})
	if err != nil {
		log.Printf("err: %s ", err)
		return ""
	}
	//e.AddToWatch(string(acc))
	//Убираем ковычки из ответа
	return strings.Replace(string(acc), "\"", "", 2)
}

func (bs *EthBlockchainScanner) AddTxCh(ch chan<- *txdata) {
	bs.TxCh = ch
}

func (bs *EthBlockchainScanner) AddCtrEvChs(ch chan<- *txdata) {
	bs.mu.Lock()
	bs.CtrEvChs = append(bs.CtrEvChs, ch)
	bs.mu.Unlock()
}

func (bs *EthBlockchainScanner) Scan() {
	log.Printf("EthBlockchainScanner scan started!")

	//current block number
	currentBN := uint64(0)
	tick := time.NewTicker(10 * time.Second)
	for range tick.C {
		endBN, err := bs.conn.BlockNumber()
		if err != nil {
			log.Printf("Ethereum Run: %s ", err)
			continue
		}
		if (endBN - currentBN) <= 0 {
			continue
		}

		startBN := currentBN
		if currentBN > 0 {
			startBN = endBN - currentBN
		}

		if psTxs := bs.BlockBetweenTxs(startBN, endBN); len(psTxs) > 0 {
			for _, tx := range psTxs {
				if tx.Recipient == nil {
					if rec, err := bs.conn.Request("eth_getTransactionReceipt", []interface{}{tx.Hash.String()}); err == nil {
						str, _ := rec.MarshalJSON()
						log.Printf("Recipient: %s", str)
					}

					if rec, err := bs.conn.Request("eth_getLogs", []interface{}{}); err == nil {
						str, _ := rec.MarshalJSON()
						log.Printf("Logs: %s", str)
					} else {
						fmt.Println(err)
					}

				} else {
					bs.TxCh <- tx
				}
			}
		}
		currentBN = endBN

	}
}

func (bch *EthBlockchainScanner) BlockBetweenTxs(startBN uint64, endBN uint64) []*txdata {
	var btwBlTxs []*txdata
	if (endBN - startBN) <= 0 {
		log.Printf("BlockBetweenTxs: FAIL! StartBlock: %d EndBlock %d", startBN, endBN)
		return btwBlTxs
	}
	log.Printf("BlockBetweenTxs: StartBlock: %d EndBlock %d", startBN, endBN)

	for i := startBN; i <= endBN; i++ {
		bn := rpc.NewHexNumber(i)

		blockRaw, err := bch.conn.Request("eth_getBlockByNumber", []interface{}{bn, true})
		if err != nil {
			log.Printf("BlockBetweenTxs Error: %s", err)
		}

		blHeader := types.Header{}
		if err := json.Unmarshal(blockRaw, &blHeader); err != nil {
			log.Printf("BlockBetweenTxs Error: %s", err)
		}

		blBody := new(struct{ Transactions []*txdata })
		if err := json.Unmarshal(blockRaw, &blBody); err != nil {
			log.Printf("BlockBetweenTxs Error: %s", err)
			continue
		}

		btwBlTxs = append(btwBlTxs, blBody.Transactions...)

	}
	log.Printf("Ethereum BlockBetweenTxs: %d txs finded ", len(btwBlTxs))
	return btwBlTxs
}
