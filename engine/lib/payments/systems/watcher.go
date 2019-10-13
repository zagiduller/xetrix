package systems

import (
	"engine/lib/structs"
	"sync"
)

type AddressWatcher struct {
	mu      *sync.Mutex
	account *structs.Account
	address string
	prevTxs map[string]bool
	balance string

	lastBalance    float64
	checkedBalance float64

	currentBlockNumber uint64
}

func NewEthAddressWatcher(acc *structs.Account, currentBlockNumber uint64) *AddressWatcher {
	w := &AddressWatcher{
		account:            acc,
		address:            acc.Address,
		mu:                 &sync.Mutex{},
		prevTxs:            make(map[string]bool),
		currentBlockNumber: currentBlockNumber,
		balance:            "0",
	}
	return w
}

func NewAddressWatcher(acc *structs.Account) *AddressWatcher {
	w := &AddressWatcher{
		account: acc,
		address: acc.Address,
		mu:      &sync.Mutex{},
		prevTxs: make(map[string]bool),
	}
	return w
}

//checkReceivedTx проверяет не создавалась ли транзакция с psTxId ранее
func (aw *AddressWatcher) isNewPsTx(psTxId string) bool {
	aw.mu.Lock()
	defer aw.mu.Unlock()

	_, ok := aw.prevTxs[psTxId]
	if !ok {
		aw.prevTxs[psTxId] = true
	}
	return !ok
}

func (aw *AddressWatcher) balanceChanged(newBal string) bool {
	aw.mu.Lock()
	defer aw.mu.Unlock()

	if newBal == aw.balance {
		return false
	}
	aw.balance = newBal
	return true
}

func (aw *AddressWatcher) SetCurrentBlock(cbn uint64) {
	aw.mu.Lock()
	defer aw.mu.Unlock()
	aw.currentBlockNumber = cbn
}

func (aw *AddressWatcher) GetCurrentBlock() uint64 {
	aw.mu.Lock()
	defer aw.mu.Unlock()
	return aw.currentBlockNumber
}
