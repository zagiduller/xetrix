package leveldb

import (
	"engine/lib/structs"
	"fmt"
	"testing"
)

func TestInitLevelDB(t *testing.T) {
	repo := InitLevelDB("./db")
	if _, err := repo.CreateCurrency(&structs.Currency{Name: "Bitcoin", Symbol: "BTC", Type: structs.Currency_CRYPTO_CURRENCY, Decimal: 8}); err != nil {
		t.Error(err)
	}

	cs, err := repo.GetAllCurrency()
	if err != nil {
		t.Error(err)
	}

	fmt.Println(cs)

	с, err := repo.FindCurrency(&structs.Query_Currency{Id: "92c4f453-9334-40d1-98e6-c1901307a5bb"})
	if err != nil {
		t.Error(err)
	}

	fmt.Println(с)
	repo.Close()

}
