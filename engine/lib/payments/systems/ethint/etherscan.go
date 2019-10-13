package ethint

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

type Client struct {
	commands    []Command
	active      bool
	RelTxCh     chan *RelatedTxRespone
	BlockNumber uint64
}

var (
	apikey = "&apikey="
)

func NewClient(ApiKey string) *Client {
	apikey += ApiKey
	client := &Client{
		commands: make([]Command, 0),
		RelTxCh:  make(chan *RelatedTxRespone, 10),
	}

	return client
}

func (c *Client) CurrentBlock() uint64 {
	//fmt.Println("Do balanceCommand ", cmd.query)
	url := "https://api.etherscan.io/api?module=proxy&action=eth_blockNumber" + apikey

	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Etherscan CurrentBlock request error: %s \n URL: %s \n", err, url)
		return 0
	}
	simpleResp := SimpleResponse{}
	if err := json.NewDecoder(resp.Body).Decode(&simpleResp); err != nil {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		log.Printf("Etherscan CurrentBlock decode error: %s \n URL: %s \n Body: %s \n", err, url, string(bodyBytes))
		return 0
	}

	if len(simpleResp.Result) > 2 {
		blockNumber, _ := strconv.ParseUint(simpleResp.Result[2:], 16, 64)
		return blockNumber
	}
	return 0
}

func (c *Client) NewBalancesCmd(ch chan *BalancesResponse, q string) {
	c.addCmd(&balanceCommand{query: q, blsRespCh: ch})
}

func (c *Client) NewAddressTxsCmd(ch chan *AddressTxsResponse, addr string, startblock uint64) {
	c.addCmd(&addressTxsCommand{query: addr, addrsTxsRespCh: ch, startBlock: strconv.FormatUint(startblock, 10)})
}

func (c *Client) NewContractTxsCmd(ch chan *AddressTxsResponse, contractId, addr string, startblock uint64) {
	c.addCmd(&contractTxsCommand{contractId: contractId, query: addr, addrsTxsRespCh: ch, startBlock: strconv.FormatUint(startblock, 10)})
}

func (c *Client) NewRelatedTxSendCmd(relateType, txId, hex string, addr string) {
	c.addCmd(&sendRawTxCommand{relatedTxCh: c.RelTxCh, resp: &RelatedTxRespone{Address: addr, RelateType: relateType, TxId: txId}, hex: hex})
}

func (c *Client) addCmd(cmd Command) {
	//fmt.Println(len(c.commands))
	c.commands = append(c.commands, cmd)
	fmt.Printf("New command added: %T len %d \n", cmd, len(c.commands))
}

func (c *Client) PopCmd() Command {
	if len(c.commands) > 0 {
		var cmd Command
		cmd, c.commands = c.commands[0], c.commands[1:]
		return cmd
	}
	return nil
}

func (c *Client) Run() {
	if c.active {
		fmt.Println("Etherscan client already used")
		return
	}
	c.active = true

	// Запрашиваем номер блока
	go func() {
		bntick := time.Tick(1000 * time.Millisecond)
		for range bntick {
			c.BlockNumber = c.CurrentBlock()
		}
	}()

	// 5 запросов в секунду берем с запасом
	tickMills := time.Tick(300 * time.Millisecond)
	for range tickMills {
		if cmd := c.PopCmd(); cmd != nil {
			go cmd.Do()
			log.Printf("Etherscan Do command %T ", cmd)
		}
	}
}

type Command interface {
	Do()
}

type BlockRequestCommand struct{}

type addressTxsCommand struct {
	startBlock     string
	query          string
	addrsTxsRespCh chan<- *AddressTxsResponse
}

type contractTxsCommand struct {
	startBlock     string
	contractId     string
	query          string
	addrsTxsRespCh chan<- *AddressTxsResponse
}

type sendRawTxCommand struct {
	relatedTxCh chan<- *RelatedTxRespone
	resp        *RelatedTxRespone
	hex         string
}

type balanceCommand struct {
	query     string
	blsRespCh chan<- *BalancesResponse
}

type ScanApiResponse struct {
}

type SimpleResponse struct {
	Result string
	Error  struct {
		Message string
	}
}

type AddressTxsResponse struct {
	Address         string
	Status, Message string
	Result          []struct {
		From, To, ContractAddress                string
		Value, GasPrice                          string
		Hash, Nonce, BlockHash, TransactionIndex string
		BlockNumber, TimeStamp                   string
		IsError, Confirmations                   string
	}
	Error struct {
		Message string
	}
}

type BalancesResponse struct {
	Status, Message string
	Result          []struct {
		Account, Balance string
	}
	Error struct {
		Message string
	}
}

type RelatedTxRespone struct {
	Address      string
	RelateType   string
	TxId, Result string
	Error        struct {
		Message string
	}
}

func (cmd *balanceCommand) Do() {
	//fmt.Println("Do balanceCommand ", cmd.query)
	url := "https://api.etherscan.io/api?module=account&action=balancemulti&tag=latest&address=" + cmd.query + apikey
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Etherscan client request error: %s  \n URL: %s \n", err, url)
		return
	}
	var blresp BalancesResponse
	if err := json.NewDecoder(resp.Body).Decode(&blresp); err != nil {
		fmt.Printf("Etherscan client request error: %s \n URL: %s \n", err, url)
		return
	}

	cmd.blsRespCh <- &blresp

}
func (cmd *addressTxsCommand) Do() {
	url := "https://api.etherscan.io/api?module=account&action=txlist&sort=desc&address=" + cmd.query + apikey
	url += "&startblock=" + cmd.startBlock
	url += "&endblock=latest"
	fmt.Println("Do addressTxsCommand ", url)
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Etherscan addressTxsCommand client request error: %s \n URL: %s \n", err, url)
		return
	}
	var txsResp AddressTxsResponse
	if err := json.NewDecoder(resp.Body).Decode(&txsResp); err != nil {
		fmt.Printf("Etherscan addressTxsCommand client request error: %s  \n URL: %s \n", err, url)
		return
	}
	txsResp.Address = cmd.query
	cmd.addrsTxsRespCh <- &txsResp
}

func (cmd *contractTxsCommand) Do() {
	//fmt.Println("Do contractTxsCommand ", cmd.query)
	url := "https://api.etherscan.io/api?module=account&action=tokentx&sort=desc&contractaddress=" + cmd.contractId + "&address=" + cmd.query + apikey
	url += "&startblock=" + cmd.startBlock
	url += "&endblock=latest"
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Etherscan contractTxsCommand client request error: %s \n URL: %s \n", err, url)
		return
	}
	var txsResp AddressTxsResponse
	if err := json.NewDecoder(resp.Body).Decode(&txsResp); err != nil {
		fmt.Printf("Etherscan contractTxsCommand client request error: %s  \n URL: %s \n", err, url)
		return
	}

	txsResp.Address = cmd.query
	cmd.addrsTxsRespCh <- &txsResp
}

func (cmd *sendRawTxCommand) Do() {
	url := "https://api.etherscan.io/api?module=proxy&action=eth_sendRawTransaction&hex=" + cmd.hex + apikey
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Etherscan sendRawTxCommand client request error: %s \n URL: %s \n", err, url)
		return
	}

	if err := json.NewDecoder(resp.Body).Decode(&cmd.resp); err != nil {
		fmt.Printf("Etherscan sendRawTxCommand client request error: %s \n URL: %s \n", err, url)
		return
	}
	cmd.relatedTxCh <- cmd.resp
}
