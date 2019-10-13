package events

import (
	"engine/lib/structs"
	"fmt"
	"sync"
)

var (
	mu = &sync.Mutex{}
)

type SubscribePublisher interface {
	AddEventBus(bus *Bus)
}

type Subscriber interface {
	SubscribePublisher
	Update(event *structs.Event)
}

type Publisher interface {
	SubscribePublisher
	Notify(event *structs.Event)
}

func NewBus() *Bus {
	return &Bus{
		subscribers: make(map[string][]Subscriber),
	}
}

type Bus struct {
	subscribers map[string][]Subscriber
}

func EventTypeString(e *structs.Event) string {
	var t string
	switch e.Type.(type) {
	case *structs.Event_NewUser:
		t = "Event_NewUser"
	case *structs.Event_NewAccount:
		t = "Event_NewAccount"
	case *structs.Event_NewTransaction:
		t = "Event_NewTransaction"
	case *structs.Event_TxConfirm:
		t = "Event_TxConfirm"
	case *structs.Event_NewOrder:
		t = "Event_NewOrder"
	case *structs.Event_NewContract:
		t = "Event_NewContract"
	case *structs.Event_OrderChange:
		t = "Event_OrderChange"
	case *structs.Event_OrderPerformed:
		t = "Event_OrderPerformed"
	case *structs.Event_ContractChange:
		t = "Event_ContractChange"
	case *structs.Event_ContractPerformed:
		t = "Event_ContractPerformed"
	case *structs.Event_OrderCanceled:
		t = "Event_OrderCanceled"
	case *structs.Event_BalanceChange:
		t = "Event_BalanceChange"
	case *structs.Event_NewWithdrawalOrder:
		t = "Event_NewWithdrawalOrder"
	case *structs.Event_WithdrawalPerformed:
		t = "Event_WithdrawalPerformed"
	case *structs.Event_PaySystemRawTx:
		t = "Event_PaySystemRawTx"
	case *structs.Event_NewCurrency:
		t = "Event_NewCurrency"
	case *structs.Event_CurrencyActivated:
		t = "Event_CurrencyActivated"
	case *structs.Event_CurrencyDeactivated:
		t = "Event_CurrencyDeactivated"
	case *structs.Event_AccountUpdate:
		t = "Event_AccountUpdate"
	case *structs.Event_PaySystemAdded:
		t = "Event_PaySystemAdded"
	case *structs.Event_ETHtxNeedPrepare:
		t = "Event_ETHtxNeedPrepare"
	case *structs.Event_ETHtxPrepared:
		t = "Event_ETHtxPrepared"
	case *structs.Event_RelatedPaySystemOutTx:
		t = "Event_RelatedPaySystemOutTx"
	case *structs.Event_RelatedPaySystemPrepareTx:
		t = "Event_RelatedPaySystemPrepareTx"
	case *structs.Event_TxProccessUpdate:
		t = "Event_TxProccessUpdate"
	case *structs.Event_BlockhainAccountUpdate:
		t = "Event_BlockhainAccountUpdate"
	default:
		fmt.Printf("EventBus: EventTypeString.Error type assertation. Event: %v \n", e)
	}
	return t
}

func (b *Bus) InitSubscribePublishers(sps ...SubscribePublisher) {
	for _, s := range sps {
		s.AddEventBus(b)
	}
}

func (b *Bus) Subscribe(sub Subscriber, events ...*structs.Event) {
	mu.Lock()
	defer mu.Unlock()

	for _, event := range events {
		eventType := EventTypeString(event)

		if len(eventType) > 0 {
			b.subscribers[eventType] = append(b.subscribers[eventType], sub)
		}
	}

}

func (b *Bus) NewEvent(event *structs.Event) {
	mu.Lock()
	defer mu.Unlock()
	eventType := EventTypeString(event)

	fmt.Printf("NewEvent: %s \t Subscribers: - ", eventType)

	if subs, ok := b.subscribers[eventType]; ok {
		fmt.Printf(" %d \n", len(subs))
		for _, s := range subs {
			go s.Update(event)
		}
	} else {
		fmt.Printf("Unsubscrabed event. \n")
	}

}
