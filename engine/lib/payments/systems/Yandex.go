package systems

import (
	"crypto/sha1"
	"encoding/hex"
	"engine/lib/structs"
	"go/build"
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"
)

func InitYandexSystem(curr *structs.Currency, accounts []*structs.Account, secret string, mux *http.ServeMux) *YandexMoney {
	ym := &YandexMoney{
		name:     "YandexMoney",
		symbol:   "RUB",
		secret:   secret,
		currency: curr,
		watchers: make(map[string]*AddressWatcher),
		mu:       &sync.Mutex{},
	}

	for _, acc := range accounts {
		ym.watchers[acc.Address] = NewAddressWatcher(acc)
	}

	ym.registerIncomingTxHook(mux)

	return ym
}

type YandexMoney struct {
	name, symbol string
	watchers     map[string]*AddressWatcher
	build.MultiplePackageError
	mu       *sync.Mutex
	secret   string
	currency *structs.Currency
	rawTxCh  chan<- structs.Query_RawTx
}

func (ym *YandexMoney) Name() string {
	return ym.name
}

func (ym *YandexMoney) Symbol() string {
	return ym.symbol
}

func (ym *YandexMoney) Currency() *structs.Currency {
	return ym.currency
}

func (ym *YandexMoney) AddToWatch(acc *structs.Account) {
	log.Printf("%s AddToWatch: %s \n", ym.Name(), acc.Address)
	ym.mu.Lock()
	if _, ok := ym.watchers[acc.Address]; !ok {
		ym.watchers[acc.Address] = NewAddressWatcher(acc)
	}
	ym.mu.Unlock()
}

func (ym *YandexMoney) GenerateAddress(user *structs.User) (string, uint64) {
	return "", 0
}

func (ym *YandexMoney) Run(wg *sync.WaitGroup, rawTxCh chan<- structs.Query_RawTx) {
	defer wg.Done()
	log.Printf("%s payment system run started!", ym.name)

	ym.rawTxCh = rawTxCh
}

func (ym *YandexMoney) registerIncomingTxHook(mux *http.ServeMux) {
	mux.Handle("/yandex", http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if req.Method != "POST" {
			w.Write([]byte("Method not allowed"))
			w.WriteHeader(http.StatusNotFound)
			return
		}

		req.ParseForm()

		var sha1_hash, paramString string
		var amount float64

		operation_id := req.Form.Get("operation_id")
		label := req.Form.Get("label")

		if "true" == req.Form.Get("codepro") {
			log.Printf("YandexMoney: Tx operation_id(%s) label(%s)  protected with code. ", operation_id, label)
			w.Write([]byte("YandexMoney: Tx operation protected with code."))
			return
		}

		sha1_hash = req.Form.Get("sha1_hash")
		log.Printf("YandexMoney: sha1_hash: %s CHECK: ", sha1_hash)

		paramString += req.Form.Get("notification_type")
		paramString += "&" + operation_id
		paramString += "&" + req.Form.Get("amount")
		paramString += "&" + req.Form.Get("currency")
		paramString += "&" + strings.Replace(req.Form.Get("datetime"), " ", "+", 1)
		paramString += "&" + req.Form.Get("sender")
		paramString += "&" + req.Form.Get("codepro")
		paramString += "&" + ym.secret
		paramString += "&" + label

		//log.Printf("YandexMoney: paramString: %s ", paramString)

		sha1str := sha1.Sum([]byte(paramString))
		hexstr := hex.EncodeToString(sha1str[:])

		log.Printf("YandexMoney: given sum: (%s) calculated sum (%s) , compare %t", sha1_hash, hexstr, hexstr == sha1_hash)

		if hexstr != sha1_hash {
			log.Printf("YandexMoney: corrupt sha1_hash operation_id %s", operation_id)
			w.Write([]byte("YandexMoney: corrupt sha1_hash operation_id " + operation_id))
			return
		}

		amount, _ = strconv.ParseFloat(req.Form.Get("amount"), 64)

		if aw, ok := ym.watchers[label]; ok && aw.isNewPsTx(operation_id) {
			ym.rawTxCh <- structs.Query_RawTx{
				FromAddress: "",
				ToAddress:   label,
				Amount:      amount,
				InPStxId:    operation_id,
			}
		}

	}))
}
