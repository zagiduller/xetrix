package main

import (
	"engine/lib/payments/systems"
	eth2 "engine/lib/payments/systems/eth"
	"sync"
)

//geth --dev --ipcpath ~/.ethereum/dev.ipc --minerthreads=5 --rpc --rpcaddr 'localhost' --rpcport '8545' --rpcapi 'admin,debug,eth,miner,net,personal,rpc,txpool,web3' --rpccorsdomain '*' console
//ethereumwallet --mode wallet --rpc ~/.ethereum/dev.ipc

func main() {
	eth := eth2.NewEthBlockchainScanner()

	extok := systems.InitExampleToken([]string{}, eth)
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go extok.Run(wg, nil)
	wg.Wait()
	//eth.BlockBetweenTxs(0,10)
	//acc := eth.GenerateAddress()
	//eth.AccountInTransactions("")

}
