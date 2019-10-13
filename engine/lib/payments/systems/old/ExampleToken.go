package old

import (
	"context"
	"engine/lib/payments/systems/eth"
	"engine/lib/payments/systems/eth/contracts/fixed"
	"engine/lib/structs"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/params"
	"log"
	"math/big"
	"strings"
	"sync"
)

type ExampleToken struct {
	contractId   string
	name, symbol string
	blockChain   *eth.EthBlockchainScanner
	mu           *sync.Mutex
	watchers     map[string]*AddressWatcher
}

func InitExampleToken(addresses []string, blockChain *eth.EthBlockchainScanner) *ExampleToken {
	log.Printf("Initialize ExampleToken system with len(%d) addresses ", len(addresses))
	watchers := make(map[string]*AddressWatcher)

	for _, addr := range addresses {
		watchers[addr] = NewAddressWatcher(addr)
	}

	e := &ExampleToken{
		name:       "ExampleToken",
		symbol:     "FIXED",
		contractId: "0x00A357CEb81f314F2ad9A3830CCBe116C65Ffca5",
		watchers:   watchers,
		mu:         &sync.Mutex{},
		blockChain: blockChain,
	}

	return e
}

func (e *ExampleToken) Name() string {
	return e.name
}

func (e *ExampleToken) Symbol() string {
	return e.symbol
}

func (e *ExampleToken) AddToWatch(addr string) {
	log.Printf("%s AddToWatch: %s \n", e.Name(), addr)
	e.mu.Lock()
	if _, ok := e.watchers[addr]; !ok {
		e.watchers[addr] = NewAddressWatcher(addr)
	}
	e.mu.Unlock()
}

func (e *ExampleToken) GenerateAddress() string {
	return e.blockChain.GenerateAddress()
}

func (e *ExampleToken) Run(wg *sync.WaitGroup, txpvd chan<- structs.Query_RawTx) {
	defer wg.Done()
	contractAddress := common.HexToAddress(e.contractId)
	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
	}

	logs := make(chan types.Log)

	sub, err := e.blockChain.Client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatal(err)
	}

	contractAbi, err := abi.JSON(strings.NewReader(string(fixed.FixedABI)))
	if err != nil {
		log.Fatal(err)
	}

	logTransferSig := []byte("Transfer(address,address,uint256)")
	logTransferSigHash := crypto.Keccak256Hash(logTransferSig)

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case vLog := <-logs:
			fmt.Printf("Log Block Number: %d\n", vLog.BlockNumber)
			fmt.Printf("Log Index: %d\n", vLog.Index)

			switch vLog.Topics[0].Hex() {
			case logTransferSigHash.Hex():
				fmt.Printf("Log Name: Transfer\n")

				var transferEvent LogTransfer

				err := contractAbi.Unpack(&transferEvent, "Transfer", vLog.Data)
				if err != nil {
					log.Fatal(err)
				}

				transferEvent.From = common.HexToAddress(vLog.Topics[1].Hex())
				transferEvent.To = common.HexToAddress(vLog.Topics[2].Hex())
				to, _ := transferEvent.To.MarshalText()

				fmt.Printf("From: %s\n", transferEvent.From.Hex())
				fmt.Printf("To: %s\n", to)

				if aw, ok := e.watchers[string(to)]; ok {
					if aw.isNewPsTx(vLog.TxHash.String()) {
						ammEth, _ := new(big.Float).Quo(new(big.Float).SetInt(transferEvent.Tokens), new(big.Float).SetInt(big.NewInt(params.Ether))).Float64()
						fmt.Printf("Tokens: %f\n", ammEth)
						txpvd <- structs.Query_RawTx{
							FromAddress:     "",
							ToAddress:       aw.address,
							Amount:          ammEth,
							PaymentSystemID: vLog.TxHash.String(),
						}
					}
				}
			}

			fmt.Printf("\n\n")
		}
	}
}

// LogTransfer ..
type LogTransfer struct {
	From   common.Address
	To     common.Address
	Tokens *big.Int
}

// LogApproval ..
type LogApproval struct {
	TokenOwner common.Address
	Spender    common.Address
	Tokens     *big.Int
}
