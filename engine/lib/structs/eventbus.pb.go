// Code generated by protoc-gen-go. DO NOT EDIT.
// source: eventbus.proto

package structs

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Event struct {
	// Types that are valid to be assigned to Type:
	//	*Event_NewUser
	//	*Event_NewAccount
	//	*Event_BalanceChange
	//	*Event_NewTransaction
	//	*Event_TxConfirm
	//	*Event_NewOrder
	//	*Event_NewContract
	//	*Event_OrderChange
	//	*Event_OrderPerformed
	//	*Event_ContractChange
	//	*Event_ContractPerformed
	//	*Event_OrderCanceled
	//	*Event_PaySystemRawTx
	//	*Event_NewWithdrawalOrder
	//	*Event_WithdrawalPerformed
	//	*Event_NewCurrency
	//	*Event_CurrencyActivated
	//	*Event_CurrencyDeactivated
	//	*Event_AccountUpdate
	//	*Event_PaySystemAdded
	//	*Event_ETHtxPrepared
	//	*Event_ETHtxNeedPrepare
	//	*Event_RelatedPaySystemOutTx
	//	*Event_RelatedPaySystemPrepareTx
	//	*Event_TxProccessUpdate
	//	*Event_BlockhainAccountUpdate
	//	*Event_EthInOut
	Type                 isEvent_Type `protobuf_oneof:"Type"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Event) Reset()         { *m = Event{} }
func (m *Event) String() string { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()    {}
func (*Event) Descriptor() ([]byte, []int) {
	return fileDescriptor_cf5d638f5c8c3cc4, []int{0}
}

func (m *Event) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Event.Unmarshal(m, b)
}
func (m *Event) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Event.Marshal(b, m, deterministic)
}
func (m *Event) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Event.Merge(m, src)
}
func (m *Event) XXX_Size() int {
	return xxx_messageInfo_Event.Size(m)
}
func (m *Event) XXX_DiscardUnknown() {
	xxx_messageInfo_Event.DiscardUnknown(m)
}

var xxx_messageInfo_Event proto.InternalMessageInfo

type isEvent_Type interface {
	isEvent_Type()
}

type Event_NewUser struct {
	NewUser *EventNewUser `protobuf:"bytes,1,opt,name=NewUser,proto3,oneof"`
}

type Event_NewAccount struct {
	NewAccount *EventNewAccount `protobuf:"bytes,2,opt,name=NewAccount,proto3,oneof"`
}

type Event_BalanceChange struct {
	BalanceChange *EventAccountBalanceChange `protobuf:"bytes,3,opt,name=BalanceChange,proto3,oneof"`
}

type Event_NewTransaction struct {
	NewTransaction *EventNewTx `protobuf:"bytes,4,opt,name=NewTransaction,proto3,oneof"`
}

type Event_TxConfirm struct {
	TxConfirm *EventTxConfirm `protobuf:"bytes,5,opt,name=TxConfirm,proto3,oneof"`
}

type Event_NewOrder struct {
	NewOrder *EventNewOrder `protobuf:"bytes,6,opt,name=NewOrder,proto3,oneof"`
}

type Event_NewContract struct {
	NewContract *EventNewContract `protobuf:"bytes,7,opt,name=NewContract,proto3,oneof"`
}

type Event_OrderChange struct {
	OrderChange *EventOrderChange `protobuf:"bytes,8,opt,name=OrderChange,proto3,oneof"`
}

type Event_OrderPerformed struct {
	OrderPerformed *EventOrderPerformed `protobuf:"bytes,9,opt,name=OrderPerformed,proto3,oneof"`
}

type Event_ContractChange struct {
	ContractChange *EventContractChange `protobuf:"bytes,10,opt,name=ContractChange,proto3,oneof"`
}

type Event_ContractPerformed struct {
	ContractPerformed *EventContractPerformed `protobuf:"bytes,11,opt,name=ContractPerformed,proto3,oneof"`
}

type Event_OrderCanceled struct {
	OrderCanceled *EventOrderCanceled `protobuf:"bytes,12,opt,name=OrderCanceled,proto3,oneof"`
}

type Event_PaySystemRawTx struct {
	PaySystemRawTx *EventPaySystemRawTx `protobuf:"bytes,13,opt,name=PaySystemRawTx,proto3,oneof"`
}

type Event_NewWithdrawalOrder struct {
	NewWithdrawalOrder *EventNewWithdrawal `protobuf:"bytes,14,opt,name=NewWithdrawalOrder,proto3,oneof"`
}

type Event_WithdrawalPerformed struct {
	WithdrawalPerformed *EventWithdrawalPerformed `protobuf:"bytes,15,opt,name=WithdrawalPerformed,proto3,oneof"`
}

type Event_NewCurrency struct {
	NewCurrency *EventNewCurrency `protobuf:"bytes,16,opt,name=NewCurrency,proto3,oneof"`
}

type Event_CurrencyActivated struct {
	CurrencyActivated *EventCurrencyActivated `protobuf:"bytes,17,opt,name=CurrencyActivated,proto3,oneof"`
}

type Event_CurrencyDeactivated struct {
	CurrencyDeactivated *EventCurrencyDeactivated `protobuf:"bytes,18,opt,name=CurrencyDeactivated,proto3,oneof"`
}

type Event_AccountUpdate struct {
	AccountUpdate *EventAccountUpdate `protobuf:"bytes,19,opt,name=AccountUpdate,proto3,oneof"`
}

type Event_PaySystemAdded struct {
	PaySystemAdded *EventPaySystemAdded `protobuf:"bytes,20,opt,name=PaySystemAdded,proto3,oneof"`
}

type Event_ETHtxPrepared struct {
	ETHtxPrepared *EventETHtxProcces `protobuf:"bytes,21,opt,name=ETHtxPrepared,proto3,oneof"`
}

type Event_ETHtxNeedPrepare struct {
	ETHtxNeedPrepare *EventETHtxProcces `protobuf:"bytes,22,opt,name=ETHtxNeedPrepare,proto3,oneof"`
}

type Event_RelatedPaySystemOutTx struct {
	RelatedPaySystemOutTx *EventRelatedPaySystemTx `protobuf:"bytes,23,opt,name=RelatedPaySystemOutTx,proto3,oneof"`
}

type Event_RelatedPaySystemPrepareTx struct {
	RelatedPaySystemPrepareTx *EventRelatedPaySystemTx `protobuf:"bytes,24,opt,name=RelatedPaySystemPrepareTx,proto3,oneof"`
}

type Event_TxProccessUpdate struct {
	TxProccessUpdate *EventTxProccessUpdate `protobuf:"bytes,25,opt,name=TxProccessUpdate,proto3,oneof"`
}

type Event_BlockhainAccountUpdate struct {
	BlockhainAccountUpdate *EventBlockhainAccountUpdate `protobuf:"bytes,26,opt,name=BlockhainAccountUpdate,proto3,oneof"`
}

type Event_EthInOut struct {
	EthInOut *EventEthInOut `protobuf:"bytes,27,opt,name=EthInOut,proto3,oneof"`
}

func (*Event_NewUser) isEvent_Type() {}

func (*Event_NewAccount) isEvent_Type() {}

func (*Event_BalanceChange) isEvent_Type() {}

func (*Event_NewTransaction) isEvent_Type() {}

func (*Event_TxConfirm) isEvent_Type() {}

func (*Event_NewOrder) isEvent_Type() {}

func (*Event_NewContract) isEvent_Type() {}

func (*Event_OrderChange) isEvent_Type() {}

func (*Event_OrderPerformed) isEvent_Type() {}

func (*Event_ContractChange) isEvent_Type() {}

func (*Event_ContractPerformed) isEvent_Type() {}

func (*Event_OrderCanceled) isEvent_Type() {}

func (*Event_PaySystemRawTx) isEvent_Type() {}

func (*Event_NewWithdrawalOrder) isEvent_Type() {}

func (*Event_WithdrawalPerformed) isEvent_Type() {}

func (*Event_NewCurrency) isEvent_Type() {}

func (*Event_CurrencyActivated) isEvent_Type() {}

func (*Event_CurrencyDeactivated) isEvent_Type() {}

func (*Event_AccountUpdate) isEvent_Type() {}

func (*Event_PaySystemAdded) isEvent_Type() {}

func (*Event_ETHtxPrepared) isEvent_Type() {}

func (*Event_ETHtxNeedPrepare) isEvent_Type() {}

func (*Event_RelatedPaySystemOutTx) isEvent_Type() {}

func (*Event_RelatedPaySystemPrepareTx) isEvent_Type() {}

func (*Event_TxProccessUpdate) isEvent_Type() {}

func (*Event_BlockhainAccountUpdate) isEvent_Type() {}

func (*Event_EthInOut) isEvent_Type() {}

func (m *Event) GetType() isEvent_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (m *Event) GetNewUser() *EventNewUser {
	if x, ok := m.GetType().(*Event_NewUser); ok {
		return x.NewUser
	}
	return nil
}

func (m *Event) GetNewAccount() *EventNewAccount {
	if x, ok := m.GetType().(*Event_NewAccount); ok {
		return x.NewAccount
	}
	return nil
}

func (m *Event) GetBalanceChange() *EventAccountBalanceChange {
	if x, ok := m.GetType().(*Event_BalanceChange); ok {
		return x.BalanceChange
	}
	return nil
}

func (m *Event) GetNewTransaction() *EventNewTx {
	if x, ok := m.GetType().(*Event_NewTransaction); ok {
		return x.NewTransaction
	}
	return nil
}

func (m *Event) GetTxConfirm() *EventTxConfirm {
	if x, ok := m.GetType().(*Event_TxConfirm); ok {
		return x.TxConfirm
	}
	return nil
}

func (m *Event) GetNewOrder() *EventNewOrder {
	if x, ok := m.GetType().(*Event_NewOrder); ok {
		return x.NewOrder
	}
	return nil
}

func (m *Event) GetNewContract() *EventNewContract {
	if x, ok := m.GetType().(*Event_NewContract); ok {
		return x.NewContract
	}
	return nil
}

func (m *Event) GetOrderChange() *EventOrderChange {
	if x, ok := m.GetType().(*Event_OrderChange); ok {
		return x.OrderChange
	}
	return nil
}

func (m *Event) GetOrderPerformed() *EventOrderPerformed {
	if x, ok := m.GetType().(*Event_OrderPerformed); ok {
		return x.OrderPerformed
	}
	return nil
}

func (m *Event) GetContractChange() *EventContractChange {
	if x, ok := m.GetType().(*Event_ContractChange); ok {
		return x.ContractChange
	}
	return nil
}

func (m *Event) GetContractPerformed() *EventContractPerformed {
	if x, ok := m.GetType().(*Event_ContractPerformed); ok {
		return x.ContractPerformed
	}
	return nil
}

func (m *Event) GetOrderCanceled() *EventOrderCanceled {
	if x, ok := m.GetType().(*Event_OrderCanceled); ok {
		return x.OrderCanceled
	}
	return nil
}

func (m *Event) GetPaySystemRawTx() *EventPaySystemRawTx {
	if x, ok := m.GetType().(*Event_PaySystemRawTx); ok {
		return x.PaySystemRawTx
	}
	return nil
}

func (m *Event) GetNewWithdrawalOrder() *EventNewWithdrawal {
	if x, ok := m.GetType().(*Event_NewWithdrawalOrder); ok {
		return x.NewWithdrawalOrder
	}
	return nil
}

func (m *Event) GetWithdrawalPerformed() *EventWithdrawalPerformed {
	if x, ok := m.GetType().(*Event_WithdrawalPerformed); ok {
		return x.WithdrawalPerformed
	}
	return nil
}

func (m *Event) GetNewCurrency() *EventNewCurrency {
	if x, ok := m.GetType().(*Event_NewCurrency); ok {
		return x.NewCurrency
	}
	return nil
}

func (m *Event) GetCurrencyActivated() *EventCurrencyActivated {
	if x, ok := m.GetType().(*Event_CurrencyActivated); ok {
		return x.CurrencyActivated
	}
	return nil
}

func (m *Event) GetCurrencyDeactivated() *EventCurrencyDeactivated {
	if x, ok := m.GetType().(*Event_CurrencyDeactivated); ok {
		return x.CurrencyDeactivated
	}
	return nil
}

func (m *Event) GetAccountUpdate() *EventAccountUpdate {
	if x, ok := m.GetType().(*Event_AccountUpdate); ok {
		return x.AccountUpdate
	}
	return nil
}

func (m *Event) GetPaySystemAdded() *EventPaySystemAdded {
	if x, ok := m.GetType().(*Event_PaySystemAdded); ok {
		return x.PaySystemAdded
	}
	return nil
}

func (m *Event) GetETHtxPrepared() *EventETHtxProcces {
	if x, ok := m.GetType().(*Event_ETHtxPrepared); ok {
		return x.ETHtxPrepared
	}
	return nil
}

func (m *Event) GetETHtxNeedPrepare() *EventETHtxProcces {
	if x, ok := m.GetType().(*Event_ETHtxNeedPrepare); ok {
		return x.ETHtxNeedPrepare
	}
	return nil
}

func (m *Event) GetRelatedPaySystemOutTx() *EventRelatedPaySystemTx {
	if x, ok := m.GetType().(*Event_RelatedPaySystemOutTx); ok {
		return x.RelatedPaySystemOutTx
	}
	return nil
}

func (m *Event) GetRelatedPaySystemPrepareTx() *EventRelatedPaySystemTx {
	if x, ok := m.GetType().(*Event_RelatedPaySystemPrepareTx); ok {
		return x.RelatedPaySystemPrepareTx
	}
	return nil
}

func (m *Event) GetTxProccessUpdate() *EventTxProccessUpdate {
	if x, ok := m.GetType().(*Event_TxProccessUpdate); ok {
		return x.TxProccessUpdate
	}
	return nil
}

func (m *Event) GetBlockhainAccountUpdate() *EventBlockhainAccountUpdate {
	if x, ok := m.GetType().(*Event_BlockhainAccountUpdate); ok {
		return x.BlockhainAccountUpdate
	}
	return nil
}

func (m *Event) GetEthInOut() *EventEthInOut {
	if x, ok := m.GetType().(*Event_EthInOut); ok {
		return x.EthInOut
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Event) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Event_OneofMarshaler, _Event_OneofUnmarshaler, _Event_OneofSizer, []interface{}{
		(*Event_NewUser)(nil),
		(*Event_NewAccount)(nil),
		(*Event_BalanceChange)(nil),
		(*Event_NewTransaction)(nil),
		(*Event_TxConfirm)(nil),
		(*Event_NewOrder)(nil),
		(*Event_NewContract)(nil),
		(*Event_OrderChange)(nil),
		(*Event_OrderPerformed)(nil),
		(*Event_ContractChange)(nil),
		(*Event_ContractPerformed)(nil),
		(*Event_OrderCanceled)(nil),
		(*Event_PaySystemRawTx)(nil),
		(*Event_NewWithdrawalOrder)(nil),
		(*Event_WithdrawalPerformed)(nil),
		(*Event_NewCurrency)(nil),
		(*Event_CurrencyActivated)(nil),
		(*Event_CurrencyDeactivated)(nil),
		(*Event_AccountUpdate)(nil),
		(*Event_PaySystemAdded)(nil),
		(*Event_ETHtxPrepared)(nil),
		(*Event_ETHtxNeedPrepare)(nil),
		(*Event_RelatedPaySystemOutTx)(nil),
		(*Event_RelatedPaySystemPrepareTx)(nil),
		(*Event_TxProccessUpdate)(nil),
		(*Event_BlockhainAccountUpdate)(nil),
		(*Event_EthInOut)(nil),
	}
}

func _Event_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Event)
	// Type
	switch x := m.Type.(type) {
	case *Event_NewUser:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.NewUser); err != nil {
			return err
		}
	case *Event_NewAccount:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.NewAccount); err != nil {
			return err
		}
	case *Event_BalanceChange:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.BalanceChange); err != nil {
			return err
		}
	case *Event_NewTransaction:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.NewTransaction); err != nil {
			return err
		}
	case *Event_TxConfirm:
		b.EncodeVarint(5<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.TxConfirm); err != nil {
			return err
		}
	case *Event_NewOrder:
		b.EncodeVarint(6<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.NewOrder); err != nil {
			return err
		}
	case *Event_NewContract:
		b.EncodeVarint(7<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.NewContract); err != nil {
			return err
		}
	case *Event_OrderChange:
		b.EncodeVarint(8<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.OrderChange); err != nil {
			return err
		}
	case *Event_OrderPerformed:
		b.EncodeVarint(9<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.OrderPerformed); err != nil {
			return err
		}
	case *Event_ContractChange:
		b.EncodeVarint(10<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ContractChange); err != nil {
			return err
		}
	case *Event_ContractPerformed:
		b.EncodeVarint(11<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ContractPerformed); err != nil {
			return err
		}
	case *Event_OrderCanceled:
		b.EncodeVarint(12<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.OrderCanceled); err != nil {
			return err
		}
	case *Event_PaySystemRawTx:
		b.EncodeVarint(13<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.PaySystemRawTx); err != nil {
			return err
		}
	case *Event_NewWithdrawalOrder:
		b.EncodeVarint(14<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.NewWithdrawalOrder); err != nil {
			return err
		}
	case *Event_WithdrawalPerformed:
		b.EncodeVarint(15<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.WithdrawalPerformed); err != nil {
			return err
		}
	case *Event_NewCurrency:
		b.EncodeVarint(16<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.NewCurrency); err != nil {
			return err
		}
	case *Event_CurrencyActivated:
		b.EncodeVarint(17<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.CurrencyActivated); err != nil {
			return err
		}
	case *Event_CurrencyDeactivated:
		b.EncodeVarint(18<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.CurrencyDeactivated); err != nil {
			return err
		}
	case *Event_AccountUpdate:
		b.EncodeVarint(19<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.AccountUpdate); err != nil {
			return err
		}
	case *Event_PaySystemAdded:
		b.EncodeVarint(20<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.PaySystemAdded); err != nil {
			return err
		}
	case *Event_ETHtxPrepared:
		b.EncodeVarint(21<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ETHtxPrepared); err != nil {
			return err
		}
	case *Event_ETHtxNeedPrepare:
		b.EncodeVarint(22<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ETHtxNeedPrepare); err != nil {
			return err
		}
	case *Event_RelatedPaySystemOutTx:
		b.EncodeVarint(23<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.RelatedPaySystemOutTx); err != nil {
			return err
		}
	case *Event_RelatedPaySystemPrepareTx:
		b.EncodeVarint(24<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.RelatedPaySystemPrepareTx); err != nil {
			return err
		}
	case *Event_TxProccessUpdate:
		b.EncodeVarint(25<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.TxProccessUpdate); err != nil {
			return err
		}
	case *Event_BlockhainAccountUpdate:
		b.EncodeVarint(26<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.BlockhainAccountUpdate); err != nil {
			return err
		}
	case *Event_EthInOut:
		b.EncodeVarint(27<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.EthInOut); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Event.Type has unexpected type %T", x)
	}
	return nil
}

func _Event_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Event)
	switch tag {
	case 1: // Type.NewUser
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(EventNewUser)
		err := b.DecodeMessage(msg)
		m.Type = &Event_NewUser{msg}
		return true, err
	case 2: // Type.NewAccount
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(EventNewAccount)
		err := b.DecodeMessage(msg)
		m.Type = &Event_NewAccount{msg}
		return true, err
	case 3: // Type.BalanceChange
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(EventAccountBalanceChange)
		err := b.DecodeMessage(msg)
		m.Type = &Event_BalanceChange{msg}
		return true, err
	case 4: // Type.NewTransaction
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(EventNewTx)
		err := b.DecodeMessage(msg)
		m.Type = &Event_NewTransaction{msg}
		return true, err
	case 5: // Type.TxConfirm
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(EventTxConfirm)
		err := b.DecodeMessage(msg)
		m.Type = &Event_TxConfirm{msg}
		return true, err
	case 6: // Type.NewOrder
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(EventNewOrder)
		err := b.DecodeMessage(msg)
		m.Type = &Event_NewOrder{msg}
		return true, err
	case 7: // Type.NewContract
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(EventNewContract)
		err := b.DecodeMessage(msg)
		m.Type = &Event_NewContract{msg}
		return true, err
	case 8: // Type.OrderChange
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(EventOrderChange)
		err := b.DecodeMessage(msg)
		m.Type = &Event_OrderChange{msg}
		return true, err
	case 9: // Type.OrderPerformed
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(EventOrderPerformed)
		err := b.DecodeMessage(msg)
		m.Type = &Event_OrderPerformed{msg}
		return true, err
	case 10: // Type.ContractChange
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(EventContractChange)
		err := b.DecodeMessage(msg)
		m.Type = &Event_ContractChange{msg}
		return true, err
	case 11: // Type.ContractPerformed
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(EventContractPerformed)
		err := b.DecodeMessage(msg)
		m.Type = &Event_ContractPerformed{msg}
		return true, err
	case 12: // Type.OrderCanceled
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(EventOrderCanceled)
		err := b.DecodeMessage(msg)
		m.Type = &Event_OrderCanceled{msg}
		return true, err
	case 13: // Type.PaySystemRawTx
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(EventPaySystemRawTx)
		err := b.DecodeMessage(msg)
		m.Type = &Event_PaySystemRawTx{msg}
		return true, err
	case 14: // Type.NewWithdrawalOrder
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(EventNewWithdrawal)
		err := b.DecodeMessage(msg)
		m.Type = &Event_NewWithdrawalOrder{msg}
		return true, err
	case 15: // Type.WithdrawalPerformed
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(EventWithdrawalPerformed)
		err := b.DecodeMessage(msg)
		m.Type = &Event_WithdrawalPerformed{msg}
		return true, err
	case 16: // Type.NewCurrency
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(EventNewCurrency)
		err := b.DecodeMessage(msg)
		m.Type = &Event_NewCurrency{msg}
		return true, err
	case 17: // Type.CurrencyActivated
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(EventCurrencyActivated)
		err := b.DecodeMessage(msg)
		m.Type = &Event_CurrencyActivated{msg}
		return true, err
	case 18: // Type.CurrencyDeactivated
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(EventCurrencyDeactivated)
		err := b.DecodeMessage(msg)
		m.Type = &Event_CurrencyDeactivated{msg}
		return true, err
	case 19: // Type.AccountUpdate
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(EventAccountUpdate)
		err := b.DecodeMessage(msg)
		m.Type = &Event_AccountUpdate{msg}
		return true, err
	case 20: // Type.PaySystemAdded
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(EventPaySystemAdded)
		err := b.DecodeMessage(msg)
		m.Type = &Event_PaySystemAdded{msg}
		return true, err
	case 21: // Type.ETHtxPrepared
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(EventETHtxProcces)
		err := b.DecodeMessage(msg)
		m.Type = &Event_ETHtxPrepared{msg}
		return true, err
	case 22: // Type.ETHtxNeedPrepare
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(EventETHtxProcces)
		err := b.DecodeMessage(msg)
		m.Type = &Event_ETHtxNeedPrepare{msg}
		return true, err
	case 23: // Type.RelatedPaySystemOutTx
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(EventRelatedPaySystemTx)
		err := b.DecodeMessage(msg)
		m.Type = &Event_RelatedPaySystemOutTx{msg}
		return true, err
	case 24: // Type.RelatedPaySystemPrepareTx
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(EventRelatedPaySystemTx)
		err := b.DecodeMessage(msg)
		m.Type = &Event_RelatedPaySystemPrepareTx{msg}
		return true, err
	case 25: // Type.TxProccessUpdate
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(EventTxProccessUpdate)
		err := b.DecodeMessage(msg)
		m.Type = &Event_TxProccessUpdate{msg}
		return true, err
	case 26: // Type.BlockhainAccountUpdate
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(EventBlockhainAccountUpdate)
		err := b.DecodeMessage(msg)
		m.Type = &Event_BlockhainAccountUpdate{msg}
		return true, err
	case 27: // Type.EthInOut
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(EventEthInOut)
		err := b.DecodeMessage(msg)
		m.Type = &Event_EthInOut{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Event_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Event)
	// Type
	switch x := m.Type.(type) {
	case *Event_NewUser:
		s := proto.Size(x.NewUser)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Event_NewAccount:
		s := proto.Size(x.NewAccount)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Event_BalanceChange:
		s := proto.Size(x.BalanceChange)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Event_NewTransaction:
		s := proto.Size(x.NewTransaction)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Event_TxConfirm:
		s := proto.Size(x.TxConfirm)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Event_NewOrder:
		s := proto.Size(x.NewOrder)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Event_NewContract:
		s := proto.Size(x.NewContract)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Event_OrderChange:
		s := proto.Size(x.OrderChange)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Event_OrderPerformed:
		s := proto.Size(x.OrderPerformed)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Event_ContractChange:
		s := proto.Size(x.ContractChange)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Event_ContractPerformed:
		s := proto.Size(x.ContractPerformed)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Event_OrderCanceled:
		s := proto.Size(x.OrderCanceled)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Event_PaySystemRawTx:
		s := proto.Size(x.PaySystemRawTx)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Event_NewWithdrawalOrder:
		s := proto.Size(x.NewWithdrawalOrder)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Event_WithdrawalPerformed:
		s := proto.Size(x.WithdrawalPerformed)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Event_NewCurrency:
		s := proto.Size(x.NewCurrency)
		n += 2 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Event_CurrencyActivated:
		s := proto.Size(x.CurrencyActivated)
		n += 2 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Event_CurrencyDeactivated:
		s := proto.Size(x.CurrencyDeactivated)
		n += 2 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Event_AccountUpdate:
		s := proto.Size(x.AccountUpdate)
		n += 2 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Event_PaySystemAdded:
		s := proto.Size(x.PaySystemAdded)
		n += 2 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Event_ETHtxPrepared:
		s := proto.Size(x.ETHtxPrepared)
		n += 2 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Event_ETHtxNeedPrepare:
		s := proto.Size(x.ETHtxNeedPrepare)
		n += 2 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Event_RelatedPaySystemOutTx:
		s := proto.Size(x.RelatedPaySystemOutTx)
		n += 2 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Event_RelatedPaySystemPrepareTx:
		s := proto.Size(x.RelatedPaySystemPrepareTx)
		n += 2 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Event_TxProccessUpdate:
		s := proto.Size(x.TxProccessUpdate)
		n += 2 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Event_BlockhainAccountUpdate:
		s := proto.Size(x.BlockhainAccountUpdate)
		n += 2 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Event_EthInOut:
		s := proto.Size(x.EthInOut)
		n += 2 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*Event)(nil), "structs.Event")
}

func init() { proto.RegisterFile("eventbus.proto", fileDescriptor_cf5d638f5c8c3cc4) }

var fileDescriptor_cf5d638f5c8c3cc4 = []byte{
	// 724 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x55, 0x5d, 0x4f, 0xdb, 0x48,
	0x14, 0x35, 0xbb, 0x7c, 0x4e, 0x36, 0x01, 0x86, 0xaf, 0x21, 0xa0, 0x5d, 0x16, 0xf5, 0xa1, 0x4f,
	0x48, 0xfd, 0x90, 0x2a, 0x55, 0xe2, 0x81, 0xa4, 0x54, 0x6e, 0xd5, 0x3a, 0xc8, 0x35, 0x6a, 0x9f,
	0x4a, 0x07, 0xfb, 0x42, 0xa2, 0x26, 0xe3, 0x68, 0x3c, 0xc6, 0xc9, 0x7f, 0xea, 0x8f, 0xac, 0xe6,
	0x66, 0x9c, 0x64, 0x9c, 0xa1, 0xe2, 0xcd, 0x3e, 0xf7, 0x9c, 0x93, 0x7b, 0xef, 0x9c, 0x78, 0x48,
	0x03, 0x1e, 0x40, 0xa8, 0xdb, 0x3c, 0x3b, 0x1b, 0xca, 0x54, 0xa5, 0x74, 0x2d, 0x53, 0x32, 0x8f,
	0x55, 0xd6, 0x6c, 0xc4, 0xb9, 0x94, 0x20, 0xe2, 0xf1, 0xa4, 0xd0, 0x24, 0x79, 0x06, 0xd2, 0x3c,
	0xd7, 0x79, 0x1c, 0xa7, 0xb9, 0x50, 0xe6, 0x75, 0x5b, 0x49, 0x2e, 0x32, 0x1e, 0xab, 0x5e, 0x2a,
	0x0c, 0x74, 0x3c, 0x07, 0xdd, 0x0c, 0x65, 0x1a, 0x43, 0x96, 0xf5, 0xc4, 0xbd, 0xa9, 0xd6, 0x52,
	0x99, 0x4c, 0xcd, 0x76, 0x87, 0x7c, 0x3c, 0x00, 0xa1, 0x6e, 0xb2, 0x71, 0xa6, 0x60, 0x30, 0x41,
	0x4f, 0x7f, 0x6d, 0x92, 0x95, 0x4b, 0xdd, 0x1a, 0x7d, 0x41, 0xd6, 0x02, 0x28, 0xae, 0x33, 0x90,
	0x6c, 0xe9, 0x64, 0xe9, 0x79, 0xed, 0xe5, 0xde, 0x99, 0xe9, 0xf1, 0x0c, 0x09, 0xa6, 0xe8, 0x7b,
	0x61, 0xc9, 0xa3, 0x6f, 0x09, 0x09, 0xa0, 0xb8, 0x98, 0x34, 0xc9, 0xfe, 0x42, 0x15, 0x5b, 0x50,
	0x99, 0xba, 0xef, 0x85, 0x73, 0x6c, 0xfa, 0x91, 0xd4, 0x5b, 0xbc, 0xcf, 0x45, 0x0c, 0xed, 0x2e,
	0x17, 0xf7, 0xc0, 0xfe, 0x46, 0xf9, 0xa9, 0x2d, 0x37, 0x6c, 0x8b, 0xe9, 0x7b, 0xa1, 0x2d, 0xa5,
	0xe7, 0xa4, 0x11, 0x40, 0x11, 0xcd, 0x56, 0xc1, 0x96, 0xd1, 0x6c, 0x67, 0xa1, 0x97, 0x68, 0xe4,
	0x7b, 0x61, 0x85, 0x4c, 0xdf, 0x90, 0x8d, 0x68, 0xd4, 0x4e, 0xc5, 0x5d, 0x4f, 0x0e, 0xd8, 0x0a,
	0x2a, 0x0f, 0x6c, 0xe5, 0xb4, 0xec, 0x7b, 0xe1, 0x8c, 0x4b, 0x5f, 0x93, 0xf5, 0x00, 0x8a, 0x8e,
	0x5e, 0x32, 0x5b, 0x45, 0xdd, 0xfe, 0xc2, 0x2f, 0x62, 0xd5, 0xf7, 0xc2, 0x29, 0x93, 0x9e, 0x93,
	0x5a, 0x00, 0x45, 0x3b, 0x15, 0x4a, 0xf2, 0x58, 0xb1, 0x35, 0x14, 0x1e, 0x2e, 0x08, 0x4b, 0x82,
	0xef, 0x85, 0xf3, 0x7c, 0x2d, 0x47, 0x1f, 0xb3, 0xb6, 0x75, 0x97, 0x7c, 0x8e, 0xa0, 0xe5, 0x73,
	0xaf, 0xf4, 0x3d, 0x69, 0xe0, 0xeb, 0x15, 0xc8, 0xbb, 0x54, 0x0e, 0x20, 0x61, 0x1b, 0xe8, 0x70,
	0xec, 0x70, 0x98, 0x72, 0xf4, 0xd2, 0x6c, 0x44, 0xfb, 0x94, 0x2d, 0x99, 0x4e, 0x88, 0xcb, 0xc7,
	0xe6, 0x68, 0x1f, 0x1b, 0xa1, 0x1d, 0xb2, 0x5d, 0x22, 0xb3, 0x96, 0x6a, 0x68, 0xf5, 0x9f, 0xdb,
	0x6a, 0xbe, 0xab, 0x45, 0x2d, 0x6d, 0x93, 0xfa, 0x64, 0x5e, 0x1d, 0x90, 0x3e, 0x24, 0xec, 0x1f,
	0x34, 0x3b, 0x72, 0x6d, 0xc8, 0x50, 0x74, 0xa2, 0x2c, 0x40, 0x4f, 0x77, 0xc5, 0xc7, 0x5f, 0xf0,
	0x9f, 0x12, 0xf2, 0x22, 0x1a, 0xb1, 0xba, 0x6b, 0x3a, 0x9b, 0xa3, 0xa7, 0xb3, 0x11, 0xfa, 0x99,
	0xd0, 0x00, 0x8a, 0xaf, 0x3d, 0xd5, 0x4d, 0x24, 0x2f, 0x78, 0x7f, 0x92, 0x95, 0x86, 0xab, 0x23,
	0x8b, 0xe7, 0x7b, 0xa1, 0x43, 0x48, 0xaf, 0xc9, 0xce, 0x0c, 0x9a, 0xad, 0x6b, 0x13, 0xfd, 0xfe,
	0xb7, 0xfd, 0x1c, 0x44, 0xdf, 0x0b, 0x5d, 0xfa, 0x32, 0x91, 0xe6, 0x43, 0xc4, 0xb6, 0x1e, 0x4b,
	0xa4, 0x21, 0x94, 0x89, 0x34, 0xaf, 0x78, 0x84, 0xe6, 0xf9, 0x22, 0x56, 0xbd, 0x07, 0xae, 0x20,
	0x61, 0xdb, 0xce, 0x23, 0xac, 0xd2, 0xf0, 0x08, 0xab, 0xa0, 0x1e, 0xb3, 0x04, 0xdf, 0x01, 0x9f,
	0x5a, 0x52, 0xd7, 0x98, 0x0e, 0xa2, 0x1e, 0xd3, 0x01, 0xeb, 0x64, 0x98, 0xef, 0xc9, 0xf5, 0x30,
	0xe1, 0x0a, 0xd8, 0x8e, 0xeb, 0x1c, 0x2c, 0x8a, 0x4e, 0x86, 0x05, 0x58, 0xc9, 0xb8, 0x48, 0x12,
	0x48, 0xd8, 0xee, 0x1f, 0x93, 0x81, 0x1c, 0x2b, 0x19, 0x88, 0xd0, 0x16, 0xa9, 0x5f, 0x46, 0xbe,
	0x1a, 0x5d, 0x49, 0x18, 0x72, 0x09, 0x09, 0xdb, 0x43, 0x9b, 0xa6, 0x6d, 0x63, 0x28, 0x69, 0x1c,
	0x43, 0xa6, 0x7b, 0xb1, 0x24, 0xd4, 0x27, 0x5b, 0x08, 0x04, 0x00, 0x89, 0x01, 0xd9, 0xfe, 0x13,
	0x6c, 0x16, 0x54, 0xf4, 0x1b, 0xd9, 0x0b, 0xa1, 0xaf, 0xb7, 0x34, 0x6d, 0xb3, 0x93, 0xab, 0x68,
	0xc4, 0x0e, 0xd0, 0xee, 0xc4, 0xb6, 0xab, 0x52, 0x31, 0xfa, 0x6e, 0x03, 0xfa, 0x83, 0x1c, 0x56,
	0x0b, 0xe6, 0x47, 0xa3, 0x11, 0x63, 0x4f, 0x76, 0x7f, 0xdc, 0x84, 0x7e, 0x22, 0x5b, 0x51, 0x39,
	0x5c, 0x66, 0x4e, 0xf6, 0x10, 0x8d, 0xff, 0xad, 0x7e, 0xc5, 0x6d, 0x96, 0xde, 0x44, 0x15, 0xa3,
	0xdf, 0xc9, 0x7e, 0xab, 0x9f, 0xc6, 0x3f, 0xbb, 0xbc, 0x27, 0xec, 0xb4, 0x34, 0xd1, 0xf3, 0x99,
	0xed, 0xe9, 0xe6, 0xfa, 0x5e, 0xf8, 0x88, 0x8b, 0xbe, 0x33, 0x2e, 0x55, 0xf7, 0x83, 0xe8, 0xe4,
	0x8a, 0x1d, 0xb9, 0xee, 0x8c, 0xb2, 0xaa, 0xef, 0x8c, 0xf2, 0xb9, 0xb5, 0x4a, 0x96, 0xa3, 0xf1,
	0x10, 0x6e, 0x57, 0xf1, 0xd6, 0x7e, 0xf5, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x23, 0xa8, 0xf2, 0x09,
	0x4f, 0x08, 0x00, 0x00,
}
