package ethint

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"engine/lib/helper"
	"engine/lib/structs"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/params"
	"github.com/gophergala2016/etherapis/etherapis/Godeps/_workspace/src/github.com/ethereum/go-ethereum/crypto/sha3"
	"github.com/miguelmota/go-ethereum-hdwallet"
	"log"
	"math/big"
	"net/http"
	"strconv"
	"sync"
	"text/template"
)

var (
	mu = &sync.Mutex{}
)

type Ethint struct {
	hd              bool
	master          string
	derivationPath  accounts.DerivationPath
	path            string
	Sc              *Client
	AdminAddress    string
	wallet          *hdwallet.Wallet
	drvTmpl         *template.Template
	MainTxsResultCh chan *AddressTxsResponse
}

type forHd struct {
	CurInc, UsrInc uint32
}

func New(apikey, path string) *Ethint {
	if len(path) == 0 {
		path = "../.ethereum/keystore"
	}
	log.Printf("Initialize Ethint with keystore path: %s", path)

	ethint := new(Ethint)
	ethint.path = path
	ethint.Sc = NewClient(apikey)

	return ethint
}

func NewHD(mnemonic, derivationTemplate, master, apikey string) *Ethint {
	wallet, err := hdwallet.NewFromMnemonic(mnemonic)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Initialize Ethint with mnemonic HD")

	ethint := new(Ethint)
	ethint.hd = true
	ethint.wallet = wallet
	ethint.Sc = NewClient(apikey)

	ethint.MainTxsResultCh = make(chan *AddressTxsResponse, 10)

	if len(derivationTemplate) == 0 {
		derivationTemplate = "m/44'/60'/0'/{{.CurInc}}/{{.UsrInc}}"
	}

	ethint.drvTmpl = template.Must(template.New("derivationTempalate").Parse(derivationTemplate))

	ethint.master = master

	ethint.AdminAddress = ethint.GenerateHD(0, 0)
	log.Printf("Admin Address: %s", ethint.AdminAddress)

	return ethint
}

func (e *Ethint) GenerateAddress(curr *structs.Currency, user *structs.User) (string, uint64) {
	// Можно проверочку сделать
	//if curr.Type != structs.Currency_ETH_CONTRACT_TOKEN

	if !e.hd {
		acc, err := keystore.NewKeyStore(e.path+"/"+user.Id, keystore.StandardScryptN, keystore.StandardScryptP).NewAccount(user.Id)
		if err != nil {
			log.Printf("Ethint: GenerateAddress: Error %s", err)
			return "", 0
		}

		return acc.Address.String(), e.Sc.BlockNumber
	}

	return e.GenerateHD(curr.Inc, user.Inc), e.Sc.BlockNumber

}

func (e *Ethint) GenerateHD(cinc, uinc uint32) string {
	buf := bytes.Buffer{}
	if err := e.drvTmpl.Execute(&buf, forHd{CurInc: cinc, UsrInc: uinc}); err != nil {
		log.Printf("NewHD: Error %s", err)
		return ""
	}
	spath := buf.String()

	log.Printf("NewHD: %s", spath)
	account, err := e.wallet.Derive(hdwallet.MustParseDerivationPath(spath), true)
	if err != nil {
		log.Printf("NewHD: Error %s", err)
		return ""
	}

	addr := account.Address.String()

	return addr
}

func (e *Ethint) GetDerivationPath(cinc, uinc uint32) string {
	buf := bytes.Buffer{}
	if err := e.drvTmpl.Execute(&buf, forHd{CurInc: cinc, UsrInc: uinc}); err != nil {
		log.Printf("GetDerivationPath: Error %s", err)
		return ""
	}
	return buf.String()
}

func (e *Ethint) Run() {

}

const RelateTypeOut = "out"
const RelateTypePrepare = "prepare"

func (e *Ethint) CreateTokenTransferData(to string, amount float64) []byte {
	var data []byte

	transferFnSignature := []byte("transfer(address,uint256)")
	toAddress := common.HexToAddress(to)

	hash := sha3.NewKeccak256()
	hash.Write(transferFnSignature)
	methodID := hash.Sum(nil)[:4]

	paddedAddress := common.LeftPadBytes(toAddress.Bytes(), 32)
	fmt.Println(hexutil.Encode(paddedAddress))

	amountValue := floatToBigInt(amount)

	paddedAmount := common.LeftPadBytes(amountValue.Bytes(), 32)

	data = append(data, methodID...)
	data = append(data, paddedAddress...)
	data = append(data, paddedAmount...)
	return data
}

// func (e *Ethint) CreateTx(drvPathFrom, addressTo string, amount string, data []byte) {
// 	//log.Printf("Ethint: CreateTx from %s ")
// 	log.Printf("Ethint: CreateTx: from path: %s to %s", drvPathFrom, addressTo)
// 	fromAccount, err := e.wallet.Derive(hdwallet.MustParseDerivationPath(drvPathFrom), true)
// 	if err != nil {
// 		log.Printf("Ethint: CreateTx: %s", err)
// 		return
// 	}

// }

// Пополнение токен счета клиента что бы вывести с него средства
func (e *Ethint) RelatedPrepareTxFromAdmin(c *structs.Currency, u *structs.User, a *structs.Account, tx *structs.Tx) {
	log.Printf("Ethint: RelatedPrepareTxFromAdmin start deploying %s transaction from admin to address %s", c.Symbol, a.Address)
	buf := bytes.Buffer{}
	// Мастер счет типа получаю
	if err := e.drvTmpl.Execute(&buf, forHd{CurInc: 0, UsrInc: 0}); err != nil {
		log.Printf("Ethint: FromMaster: Error %s", err)
		return
	}
	spath := buf.String()
	log.Printf("Ethint: FromMaster: master path: %s", spath)
	adminAccount, err := e.wallet.Derive(hdwallet.MustParseDerivationPath(spath), true)
	if err != nil {
		log.Printf("Ethint: FromMaster: %s", err)
		return
	}

	nonce := getNonce(adminAccount.Address.String())

	var data []byte

	tokenGasLimit := uint64(100000)
	tokenGasPrice := big.NewInt(10000000000) // Рассчет газа

	gasLimit := uint64(25000)
	gasPrice := big.NewInt(4000000000)

	value := calcEstimate(tokenGasLimit, tokenGasPrice)

	_tx := types.NewTransaction(nonce, common.HexToAddress(a.Address), value, gasLimit, gasPrice, data)

	signedTx, err := e.wallet.SignTx(adminAccount, _tx, nil)
	if err != nil {
		log.Printf("Ethint: FromMaster: SignTx %s", err)
		return
	}

	ts := types.Transactions{signedTx}
	rawTxBytes := ts.GetRlp(0)
	rawTxHex := hex.EncodeToString(rawTxBytes)

	amm := new(big.Int).Add(value, value)
	ammEth, _ := new(big.Float).Quo(new(big.Float).SetInt(amm), new(big.Float).SetInt(big.NewInt(params.Ether))).Float64()
	fmt.Printf("Value and gas cost: %s WEI \t %f ETH \n", amm.String(), ammEth)

	fmt.Println("rawTxHex:")
	fmt.Println(rawTxHex)

	e.Sc.NewRelatedTxSendCmd(RelateTypePrepare, tx.Id, rawTxHex, a.Address)
}

func (e *Ethint) RelatedTxToMaster(c *structs.Currency, u *structs.User, a *structs.Account, tx *structs.Tx) {
	log.Printf("Ethint: RelatedTxToMaster start deploying %s transaction from %s to master address", c.Symbol, a.Address)

	if u.Status == structs.UserStatus_ADMINISTRATOR && a.Currency.Symbol == "ETH" {
		log.Println("RelatedTxToMaster: Incoming FUNd to admin ETH account. Skip.")
		e.Sc.RelTxCh <- &RelatedTxRespone{
			Address:    a.Address,
			RelateType: RelateTypeOut,
			TxId:       tx.Id,
			Result:     "AdminFundTx-" + helper.RandStringRunes(6),
		}
		return
	}

	buf := bytes.Buffer{}
	if err := e.drvTmpl.Execute(&buf, forHd{CurInc: c.Inc, UsrInc: u.Inc}); err != nil {
		log.Printf("Ethint: TxToMaster: Error %s", err)
		return
	}
	spath := buf.String()
	log.Printf("Ethint: TxToMaster: path: %s", spath)
	account, err := e.wallet.Derive(hdwallet.MustParseDerivationPath(spath), true)
	if err != nil {
		log.Printf("Ethint: TxToMaster: %s", err)
		return
	}

	value := big.NewInt(0)

	toAddress := common.HexToAddress(e.master)

	nonce := getNonce(a.Address)

	var data []byte
	var _tx *types.Transaction

	if c.Symbol == "ETH" {
		amount := floatToBigInt(tx.Amount)

		gasLimit := uint64(25000)
		gasPrice := big.NewInt(4000000000)

		value = new(big.Int).Sub(amount, calcEstimate(gasLimit, gasPrice))

		_tx = types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)

	} else {
		// Все остальные контракты
		contractAddress := common.HexToAddress(c.ContractId)
		transferFnSignature := []byte("transfer(address,uint256)")

		tokenGasLimit := uint64(100000)
		tokenGasPrice := big.NewInt(10000000000) // Рассчет газа

		hash := sha3.NewKeccak256()
		hash.Write(transferFnSignature)
		methodID := hash.Sum(nil)[:4]
		fmt.Println(hexutil.Encode(methodID)) // 0xa9059cbb

		paddedAddress := common.LeftPadBytes(toAddress.Bytes(), 32)
		fmt.Println(hexutil.Encode(paddedAddress))

		amount := floatToBigInt(tx.Amount)
		paddedAmount := common.LeftPadBytes(amount.Bytes(), 32)

		data = append(data, methodID...)
		data = append(data, paddedAddress...)
		data = append(data, paddedAmount...)

		_tx = types.NewTransaction(nonce, contractAddress, value, tokenGasLimit, tokenGasPrice, data)
	}

	signedTx, err := e.wallet.SignTx(account, _tx, nil)
	if err != nil {
		log.Printf("Ethint: TxToMaster: SignTx %s", err)
		return
	}

	ts := types.Transactions{signedTx}
	rawTxBytes := ts.GetRlp(0)
	rawTxHex := hex.EncodeToString(rawTxBytes)
	fmt.Println("rawTxHex:")
	fmt.Println(rawTxHex)

	e.Sc.NewRelatedTxSendCmd(RelateTypeOut, tx.Id, rawTxHex, a.Address)

}

func TransferEth() {

}

func TransferToken() {

}

func calcEstimate(gasLimit uint64, gasPrice *big.Int) *big.Int {
	return new(big.Int).Mul(new(big.Int).SetUint64(gasLimit), gasPrice)
}

func floatToBigInt(val float64) *big.Int {
	bigval := new(big.Float)
	bigval.SetFloat64(val)
	// Set precision if required.
	// bigval.SetPrec(64)

	coin := new(big.Float)
	coin.SetInt(big.NewInt(1000000000000000000))

	bigval.Mul(bigval, coin)

	result := new(big.Int)
	bigval.Int(result) // store converted number in result

	return result
}

func getNonce(addr string) uint64 {
	resp, err := http.Get("https://api.etherscan.io/api?module=proxy&action=eth_getTransactionCount&tag=latest&apikey=YourApiKeyToken&address=" + addr)
	if err != nil {
		fmt.Printf("Ethint: error nonce getting  %s ", err)
		return 0
	}
	var data struct {
		Result string
		Error  struct {
			Message string
		}
	}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		fmt.Printf("Ethint: error nonce getting: %s ", err)
		return 0
	}

	if len(data.Result) > 2 {
		nonce, _ := strconv.ParseUint(data.Result[2:], 16, 64)
		return nonce
	}

	return 0

}
