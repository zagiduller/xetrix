package eth

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"log"
	"math/big"
	"strings"
	"sync"
	"time"
)

// Сценарий:
// Приложение остановлено - в блокчейне повились новые блоки
// в которых теоретически могут быть транзакции на счета пользователя системы.
// Соответственно --
// При запуске приложения нужно понять где мы остановились в прошлый раз.
// Какой блок был обработан последним, какой блок текущий в блокчейне и на сколько блоков мы отстаем.
// Догрузить блоки и проверить транзакции по каждой валюте работающей с блокчейном.
//
// (Кстати, справедливо ли то же для фиатных платежных?)
//
// Требуется:
// Сохранять указатель на номер последнего обработанного системой блока.
// Для этого хорошо подойдут стуктуры Currency.
// - Реализовать метод получения и записи указателя для крипты.
// - Передавать валюты в обработчики платежных систем
//
// Ссылки:
// Получить события по контракту https://goethereumbook.org/event-read/

/// СТАРОЕ ЧТО БЫ ГДЕ_ТО БЫЛО
// *************************
//ipcpath := ""
//if cfg != nil {
//	ipcpath = cfg.UString("ethereum.ipc")
//}

//if ethscan, err := eth.NewEthBlockchainScanner(ipcpath); err == nil {
//r.AddSystem(systems.InitEthereumSystem(addrMap["ETH"], ethscan))
//r.AddSystem(systems.InitExampleToken(addrMap["FIXED"], ethscan))
//}

// Запуск сканирования блокчейна Ethereum
// -- Кажется уже не надо
// go ethscan.Scan()
//****************************

type EthBlockchainScanner struct {
	Client    *ethclient.Client
	rpcClient *rpc.Client
	ctx       context.Context
	TxCh      chan<- *types.Transaction
	//CtrEvChs [] chan <- *txdata
	mu *sync.Mutex
}

func NewEthBlockchainScanner(ipc string) (*EthBlockchainScanner, error) {
	rpcClient, err := rpc.Dial(ipc)
	if err != nil {
		return nil, fmt.Errorf("NewEthBlockchainScanner: Error %s", err)
	}
	client := ethclient.NewClient(rpcClient)

	return &EthBlockchainScanner{
		Client:    client,
		rpcClient: rpcClient,
		mu:        &sync.Mutex{},
		ctx:       context.Background(),
	}, nil
}

func (bs *EthBlockchainScanner) GenerateAddress() string {

	var result interface{}
	err := bs.rpcClient.Call(&result, "personal_newAccount", "")
	//acc, err := bs.client.Request("personal_newAccount", []interface{}{""})
	if err != nil {
		log.Printf("GenerateAddress err: %s ", err)
		return ""
	}

	//Убираем кавычки из ответа
	return strings.Replace(result.(string), "\"", "", 2)
}

func (bs *EthBlockchainScanner) AddTxCh(ch chan<- *types.Transaction) {
	bs.TxCh = ch
}

func (bs *EthBlockchainScanner) BlockBetweenTxs(startBN, endBN int64) []*types.Transaction {
	var result []*types.Transaction
	for i := startBN; i <= endBN; i++ {
		block, err := bs.Client.BlockByNumber(bs.ctx, new(big.Int).SetInt64(i))
		if err != nil {
			log.Printf("BlockBetweenTxs: Error %s", err)
			continue
		}

		txs := block.Transactions()
		for _, tx := range txs {
			if tx.To() != nil && tx.Value().Cmp(new(big.Int).SetInt64(0)) > 0 {
				result = append(result, tx)
			}
		}
	}

	return result
}

func (bs *EthBlockchainScanner) Scan() {
	log.Printf("EthBlockchainScanner scan started!")

	//current block number
	currentBN := int64(0)
	tick := time.NewTicker(10 * time.Second)
	for range tick.C {
		header, err := bs.Client.HeaderByNumber(context.Background(), nil)
		if err != nil {
			log.Fatal(err)
		}
		endBN := header.Number.Int64()

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
				bs.TxCh <- tx
			}
		}

		currentBN = endBN
		log.Printf("EthBlockchainScanner: current block number: %d", currentBN)
	}
}

func (bs *EthBlockchainScanner) SubscribeNewBlock() {
	headers := make(chan *types.Header)
	sub, err := bs.Client.SubscribeNewHead(context.Background(), headers)
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case header := <-headers:
			fmt.Println(header.Hash().Hex()) // 0xbc10defa8dda384c96a17640d84de5578804945d347072e091b4e5f390ddea7f

			block, err := bs.Client.BlockByHash(context.Background(), header.Hash())
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println(block.Hash().Hex())        // 0xbc10defa8dda384c96a17640d84de5578804945d347072e091b4e5f390ddea7f
			fmt.Println(block.Number().Uint64())   // 3477413
			fmt.Println(block.Time().Uint64())     // 1529525947
			fmt.Println(block.Nonce())             // 130524141876765836
			fmt.Println(len(block.Transactions())) // 7
		}
	}
}
