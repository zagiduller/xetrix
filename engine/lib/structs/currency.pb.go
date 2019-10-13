// Code generated by protoc-gen-go. DO NOT EDIT.
// source: currency.proto

package structs

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	math "math"
)

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
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

type Currency_CurrencyType int32

const (
	Currency_CRYPTO_CURRENCY    Currency_CurrencyType = 0
	Currency_FIAT_CURRENCY      Currency_CurrencyType = 1
	Currency_ETH_CONTRACT_TOKEN Currency_CurrencyType = 2
)

var Currency_CurrencyType_name = map[int32]string{
	0: "CRYPTO_CURRENCY",
	1: "FIAT_CURRENCY",
	2: "ETH_CONTRACT_TOKEN",
}

var Currency_CurrencyType_value = map[string]int32{
	"CRYPTO_CURRENCY":    0,
	"FIAT_CURRENCY":      1,
	"ETH_CONTRACT_TOKEN": 2,
}

func (x Currency_CurrencyType) String() string {
	return proto.EnumName(Currency_CurrencyType_name, int32(x))
}

func (Currency_CurrencyType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_d3dc60ed002193ea, []int{3, 0}
}

type EventNewCurrency struct {
	Currency             *Currency `protobuf:"bytes,1,opt,name=currency,proto3" json:"currency,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *EventNewCurrency) Reset()         { *m = EventNewCurrency{} }
func (m *EventNewCurrency) String() string { return proto.CompactTextString(m) }
func (*EventNewCurrency) ProtoMessage()    {}
func (*EventNewCurrency) Descriptor() ([]byte, []int) {
	return fileDescriptor_d3dc60ed002193ea, []int{0}
}

func (m *EventNewCurrency) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventNewCurrency.Unmarshal(m, b)
}
func (m *EventNewCurrency) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventNewCurrency.Marshal(b, m, deterministic)
}
func (m *EventNewCurrency) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventNewCurrency.Merge(m, src)
}
func (m *EventNewCurrency) XXX_Size() int {
	return xxx_messageInfo_EventNewCurrency.Size(m)
}
func (m *EventNewCurrency) XXX_DiscardUnknown() {
	xxx_messageInfo_EventNewCurrency.DiscardUnknown(m)
}

var xxx_messageInfo_EventNewCurrency proto.InternalMessageInfo

func (m *EventNewCurrency) GetCurrency() *Currency {
	if m != nil {
		return m.Currency
	}
	return nil
}

type EventCurrencyActivated struct {
	Currency             *Currency `protobuf:"bytes,1,opt,name=currency,proto3" json:"currency,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *EventCurrencyActivated) Reset()         { *m = EventCurrencyActivated{} }
func (m *EventCurrencyActivated) String() string { return proto.CompactTextString(m) }
func (*EventCurrencyActivated) ProtoMessage()    {}
func (*EventCurrencyActivated) Descriptor() ([]byte, []int) {
	return fileDescriptor_d3dc60ed002193ea, []int{1}
}

func (m *EventCurrencyActivated) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventCurrencyActivated.Unmarshal(m, b)
}
func (m *EventCurrencyActivated) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventCurrencyActivated.Marshal(b, m, deterministic)
}
func (m *EventCurrencyActivated) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventCurrencyActivated.Merge(m, src)
}
func (m *EventCurrencyActivated) XXX_Size() int {
	return xxx_messageInfo_EventCurrencyActivated.Size(m)
}
func (m *EventCurrencyActivated) XXX_DiscardUnknown() {
	xxx_messageInfo_EventCurrencyActivated.DiscardUnknown(m)
}

var xxx_messageInfo_EventCurrencyActivated proto.InternalMessageInfo

func (m *EventCurrencyActivated) GetCurrency() *Currency {
	if m != nil {
		return m.Currency
	}
	return nil
}

type EventCurrencyDeactivated struct {
	Currency             *Currency `protobuf:"bytes,1,opt,name=currency,proto3" json:"currency,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *EventCurrencyDeactivated) Reset()         { *m = EventCurrencyDeactivated{} }
func (m *EventCurrencyDeactivated) String() string { return proto.CompactTextString(m) }
func (*EventCurrencyDeactivated) ProtoMessage()    {}
func (*EventCurrencyDeactivated) Descriptor() ([]byte, []int) {
	return fileDescriptor_d3dc60ed002193ea, []int{2}
}

func (m *EventCurrencyDeactivated) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventCurrencyDeactivated.Unmarshal(m, b)
}
func (m *EventCurrencyDeactivated) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventCurrencyDeactivated.Marshal(b, m, deterministic)
}
func (m *EventCurrencyDeactivated) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventCurrencyDeactivated.Merge(m, src)
}
func (m *EventCurrencyDeactivated) XXX_Size() int {
	return xxx_messageInfo_EventCurrencyDeactivated.Size(m)
}
func (m *EventCurrencyDeactivated) XXX_DiscardUnknown() {
	xxx_messageInfo_EventCurrencyDeactivated.DiscardUnknown(m)
}

var xxx_messageInfo_EventCurrencyDeactivated proto.InternalMessageInfo

func (m *EventCurrencyDeactivated) GetCurrency() *Currency {
	if m != nil {
		return m.Currency
	}
	return nil
}

type Currency struct {
	// @inject_tag: bson:"_id,omitempty"
	BsonId               []byte                `protobuf:"bytes,6,opt,name=bsonId,proto3" json:"bsonId,omitempty" bson:"_id,omitempty"`
	ContractId           string                `protobuf:"bytes,7,opt,name=contractId,proto3" json:"contractId,omitempty"`
	Id                   string                `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string                `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Symbol               string                `protobuf:"bytes,3,opt,name=symbol,proto3" json:"symbol,omitempty"`
	Decimal              uint32                `protobuf:"varint,4,opt,name=decimal,proto3" json:"decimal,omitempty"`
	Type                 Currency_CurrencyType `protobuf:"varint,5,opt,name=type,proto3,enum=structs.Currency_CurrencyType" json:"type,omitempty"`
	Active               bool                  `protobuf:"varint,8,opt,name=active,proto3" json:"active,omitempty"`
	Inc                  uint32                `protobuf:"varint,9,opt,name=inc,proto3" json:"inc,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *Currency) Reset()         { *m = Currency{} }
func (m *Currency) String() string { return proto.CompactTextString(m) }
func (*Currency) ProtoMessage()    {}
func (*Currency) Descriptor() ([]byte, []int) {
	return fileDescriptor_d3dc60ed002193ea, []int{3}
}

func (m *Currency) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Currency.Unmarshal(m, b)
}
func (m *Currency) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Currency.Marshal(b, m, deterministic)
}
func (m *Currency) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Currency.Merge(m, src)
}
func (m *Currency) XXX_Size() int {
	return xxx_messageInfo_Currency.Size(m)
}
func (m *Currency) XXX_DiscardUnknown() {
	xxx_messageInfo_Currency.DiscardUnknown(m)
}

var xxx_messageInfo_Currency proto.InternalMessageInfo

func (m *Currency) GetBsonId() []byte {
	if m != nil {
		return m.BsonId
	}
	return nil
}

func (m *Currency) GetContractId() string {
	if m != nil {
		return m.ContractId
	}
	return ""
}

func (m *Currency) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Currency) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Currency) GetSymbol() string {
	if m != nil {
		return m.Symbol
	}
	return ""
}

func (m *Currency) GetDecimal() uint32 {
	if m != nil {
		return m.Decimal
	}
	return 0
}

func (m *Currency) GetType() Currency_CurrencyType {
	if m != nil {
		return m.Type
	}
	return Currency_CRYPTO_CURRENCY
}

func (m *Currency) GetActive() bool {
	if m != nil {
		return m.Active
	}
	return false
}

func (m *Currency) GetInc() uint32 {
	if m != nil {
		return m.Inc
	}
	return 0
}

type Query_CreateCurrency struct {
	Object               *Currency   `protobuf:"bytes,1,opt,name=object,proto3" json:"object,omitempty"`
	Items                []*Currency `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Query_CreateCurrency) Reset()         { *m = Query_CreateCurrency{} }
func (m *Query_CreateCurrency) String() string { return proto.CompactTextString(m) }
func (*Query_CreateCurrency) ProtoMessage()    {}
func (*Query_CreateCurrency) Descriptor() ([]byte, []int) {
	return fileDescriptor_d3dc60ed002193ea, []int{4}
}

func (m *Query_CreateCurrency) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Query_CreateCurrency.Unmarshal(m, b)
}
func (m *Query_CreateCurrency) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Query_CreateCurrency.Marshal(b, m, deterministic)
}
func (m *Query_CreateCurrency) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Query_CreateCurrency.Merge(m, src)
}
func (m *Query_CreateCurrency) XXX_Size() int {
	return xxx_messageInfo_Query_CreateCurrency.Size(m)
}
func (m *Query_CreateCurrency) XXX_DiscardUnknown() {
	xxx_messageInfo_Query_CreateCurrency.DiscardUnknown(m)
}

var xxx_messageInfo_Query_CreateCurrency proto.InternalMessageInfo

func (m *Query_CreateCurrency) GetObject() *Currency {
	if m != nil {
		return m.Object
	}
	return nil
}

func (m *Query_CreateCurrency) GetItems() []*Currency {
	if m != nil {
		return m.Items
	}
	return nil
}

type Query_Currency struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Symbol               string   `protobuf:"bytes,4,opt,name=symbol,proto3" json:"symbol,omitempty"`
	Active               bool     `protobuf:"varint,5,opt,name=active,proto3" json:"active,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Query_Currency) Reset()         { *m = Query_Currency{} }
func (m *Query_Currency) String() string { return proto.CompactTextString(m) }
func (*Query_Currency) ProtoMessage()    {}
func (*Query_Currency) Descriptor() ([]byte, []int) {
	return fileDescriptor_d3dc60ed002193ea, []int{5}
}

func (m *Query_Currency) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Query_Currency.Unmarshal(m, b)
}
func (m *Query_Currency) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Query_Currency.Marshal(b, m, deterministic)
}
func (m *Query_Currency) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Query_Currency.Merge(m, src)
}
func (m *Query_Currency) XXX_Size() int {
	return xxx_messageInfo_Query_Currency.Size(m)
}
func (m *Query_Currency) XXX_DiscardUnknown() {
	xxx_messageInfo_Query_Currency.DiscardUnknown(m)
}

var xxx_messageInfo_Query_Currency proto.InternalMessageInfo

func (m *Query_Currency) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Query_Currency) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Query_Currency) GetSymbol() string {
	if m != nil {
		return m.Symbol
	}
	return ""
}

func (m *Query_Currency) GetActive() bool {
	if m != nil {
		return m.Active
	}
	return false
}

type Response_CurrencyQuery struct {
	Created              bool        `protobuf:"varint,1,opt,name=created,proto3" json:"created,omitempty"`
	Object               *Currency   `protobuf:"bytes,2,opt,name=object,proto3" json:"object,omitempty"`
	Items                []*Currency `protobuf:"bytes,3,rep,name=items,proto3" json:"items,omitempty"`
	ItemsCount           uint32      `protobuf:"varint,4,opt,name=itemsCount,proto3" json:"itemsCount,omitempty"`
	QueryStatus          QueryStatus `protobuf:"varint,5,opt,name=queryStatus,proto3,enum=structs.QueryStatus" json:"queryStatus,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Response_CurrencyQuery) Reset()         { *m = Response_CurrencyQuery{} }
func (m *Response_CurrencyQuery) String() string { return proto.CompactTextString(m) }
func (*Response_CurrencyQuery) ProtoMessage()    {}
func (*Response_CurrencyQuery) Descriptor() ([]byte, []int) {
	return fileDescriptor_d3dc60ed002193ea, []int{6}
}

func (m *Response_CurrencyQuery) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response_CurrencyQuery.Unmarshal(m, b)
}
func (m *Response_CurrencyQuery) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response_CurrencyQuery.Marshal(b, m, deterministic)
}
func (m *Response_CurrencyQuery) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response_CurrencyQuery.Merge(m, src)
}
func (m *Response_CurrencyQuery) XXX_Size() int {
	return xxx_messageInfo_Response_CurrencyQuery.Size(m)
}
func (m *Response_CurrencyQuery) XXX_DiscardUnknown() {
	xxx_messageInfo_Response_CurrencyQuery.DiscardUnknown(m)
}

var xxx_messageInfo_Response_CurrencyQuery proto.InternalMessageInfo

func (m *Response_CurrencyQuery) GetCreated() bool {
	if m != nil {
		return m.Created
	}
	return false
}

func (m *Response_CurrencyQuery) GetObject() *Currency {
	if m != nil {
		return m.Object
	}
	return nil
}

func (m *Response_CurrencyQuery) GetItems() []*Currency {
	if m != nil {
		return m.Items
	}
	return nil
}

func (m *Response_CurrencyQuery) GetItemsCount() uint32 {
	if m != nil {
		return m.ItemsCount
	}
	return 0
}

func (m *Response_CurrencyQuery) GetQueryStatus() QueryStatus {
	if m != nil {
		return m.QueryStatus
	}
	return QueryStatus_Query_Missed
}

func init() {
	proto.RegisterEnum("structs.Currency_CurrencyType", Currency_CurrencyType_name, Currency_CurrencyType_value)
	proto.RegisterType((*EventNewCurrency)(nil), "structs.EventNewCurrency")
	proto.RegisterType((*EventCurrencyActivated)(nil), "structs.EventCurrencyActivated")
	proto.RegisterType((*EventCurrencyDeactivated)(nil), "structs.EventCurrencyDeactivated")
	proto.RegisterType((*Currency)(nil), "structs.Currency")
	proto.RegisterType((*Query_CreateCurrency)(nil), "structs.Query_CreateCurrency")
	proto.RegisterType((*Query_Currency)(nil), "structs.Query_Currency")
	proto.RegisterType((*Response_CurrencyQuery)(nil), "structs.Response_CurrencyQuery")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ServiceCurrencyClient is the client API for ServiceCurrency service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ServiceCurrencyClient interface {
	CreateCurrency(ctx context.Context, in *Query_CreateCurrency, opts ...grpc.CallOption) (*Response_CurrencyQuery, error)
	GetCurrency(ctx context.Context, in *Query_Currency, opts ...grpc.CallOption) (*Response_CurrencyQuery, error)
	Activate(ctx context.Context, in *Query_Currency, opts ...grpc.CallOption) (*Currency, error)
	Deactivate(ctx context.Context, in *Query_Currency, opts ...grpc.CallOption) (*Currency, error)
}

type serviceCurrencyClient struct {
	cc *grpc.ClientConn
}

func NewServiceCurrencyClient(cc *grpc.ClientConn) ServiceCurrencyClient {
	return &serviceCurrencyClient{cc}
}

func (c *serviceCurrencyClient) CreateCurrency(ctx context.Context, in *Query_CreateCurrency, opts ...grpc.CallOption) (*Response_CurrencyQuery, error) {
	out := new(Response_CurrencyQuery)
	err := c.cc.Invoke(ctx, "/structs.ServiceCurrency/CreateCurrency", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceCurrencyClient) GetCurrency(ctx context.Context, in *Query_Currency, opts ...grpc.CallOption) (*Response_CurrencyQuery, error) {
	out := new(Response_CurrencyQuery)
	err := c.cc.Invoke(ctx, "/structs.ServiceCurrency/GetCurrency", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceCurrencyClient) Activate(ctx context.Context, in *Query_Currency, opts ...grpc.CallOption) (*Currency, error) {
	out := new(Currency)
	err := c.cc.Invoke(ctx, "/structs.ServiceCurrency/Activate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceCurrencyClient) Deactivate(ctx context.Context, in *Query_Currency, opts ...grpc.CallOption) (*Currency, error) {
	out := new(Currency)
	err := c.cc.Invoke(ctx, "/structs.ServiceCurrency/Deactivate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServiceCurrencyServer is the server API for ServiceCurrency service.
type ServiceCurrencyServer interface {
	CreateCurrency(context.Context, *Query_CreateCurrency) (*Response_CurrencyQuery, error)
	GetCurrency(context.Context, *Query_Currency) (*Response_CurrencyQuery, error)
	Activate(context.Context, *Query_Currency) (*Currency, error)
	Deactivate(context.Context, *Query_Currency) (*Currency, error)
}

func RegisterServiceCurrencyServer(s *grpc.Server, srv ServiceCurrencyServer) {
	s.RegisterService(&_ServiceCurrency_serviceDesc, srv)
}

func _ServiceCurrency_CreateCurrency_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Query_CreateCurrency)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceCurrencyServer).CreateCurrency(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/structs.ServiceCurrency/CreateCurrency",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceCurrencyServer).CreateCurrency(ctx, req.(*Query_CreateCurrency))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceCurrency_GetCurrency_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Query_Currency)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceCurrencyServer).GetCurrency(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/structs.ServiceCurrency/GetCurrency",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceCurrencyServer).GetCurrency(ctx, req.(*Query_Currency))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceCurrency_Activate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Query_Currency)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceCurrencyServer).Activate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/structs.ServiceCurrency/Activate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceCurrencyServer).Activate(ctx, req.(*Query_Currency))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceCurrency_Deactivate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Query_Currency)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceCurrencyServer).Deactivate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/structs.ServiceCurrency/Deactivate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceCurrencyServer).Deactivate(ctx, req.(*Query_Currency))
	}
	return interceptor(ctx, in, info, handler)
}

var _ServiceCurrency_serviceDesc = grpc.ServiceDesc{
	ServiceName: "structs.ServiceCurrency",
	HandlerType: (*ServiceCurrencyServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCurrency",
			Handler:    _ServiceCurrency_CreateCurrency_Handler,
		},
		{
			MethodName: "GetCurrency",
			Handler:    _ServiceCurrency_GetCurrency_Handler,
		},
		{
			MethodName: "Activate",
			Handler:    _ServiceCurrency_Activate_Handler,
		},
		{
			MethodName: "Deactivate",
			Handler:    _ServiceCurrency_Deactivate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "currency.proto",
}

func init() { proto.RegisterFile("currency.proto", fileDescriptor_d3dc60ed002193ea) }

var fileDescriptor_d3dc60ed002193ea = []byte{
	// 614 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0xae, 0x9d, 0x34, 0x4d, 0xa7, 0x6d, 0x9a, 0x2e, 0x55, 0x58, 0x15, 0x4a, 0x5d, 0x4b, 0xa8,
	0x01, 0xa9, 0x89, 0x68, 0x25, 0x0e, 0xdc, 0x22, 0x13, 0x4a, 0x84, 0x94, 0xd0, 0xad, 0x39, 0xf4,
	0x80, 0xa2, 0xcd, 0x7a, 0x55, 0xb9, 0xaa, 0x77, 0x83, 0xbd, 0x09, 0x8a, 0x10, 0x17, 0x5e, 0xa1,
	0x2f, 0xc5, 0x9d, 0x3b, 0x27, 0x1e, 0x04, 0x79, 0xe3, 0x9f, 0xa4, 0x3f, 0xa8, 0xe1, 0xb6, 0xf3,
	0xa3, 0xef, 0x9b, 0x99, 0x6f, 0x76, 0xa0, 0xc2, 0x46, 0x61, 0xc8, 0x05, 0x9b, 0x34, 0x86, 0xa1,
	0x54, 0x12, 0xad, 0x44, 0x2a, 0x1c, 0x31, 0x15, 0xed, 0xac, 0x33, 0x19, 0x04, 0x52, 0x4c, 0xdd,
	0x3b, 0x4f, 0x2f, 0xa4, 0xbc, 0xb8, 0xe2, 0x4d, 0x3a, 0xf4, 0x9b, 0x54, 0x08, 0xa9, 0xa8, 0xf2,
	0xa5, 0x88, 0xa6, 0x51, 0xbb, 0x05, 0xd5, 0xf6, 0x98, 0x0b, 0xd5, 0xe5, 0x5f, 0x9d, 0x04, 0x0e,
	0x1d, 0x42, 0x39, 0x85, 0xc6, 0x86, 0x65, 0xd4, 0xd7, 0x8e, 0xb6, 0x1a, 0x09, 0x76, 0x23, 0x4d,
	0x22, 0x59, 0x8a, 0x7d, 0x02, 0x35, 0x0d, 0x91, 0x86, 0x5a, 0x4c, 0xf9, 0x63, 0xaa, 0xb8, 0xb7,
	0x28, 0x50, 0x07, 0xf0, 0x1c, 0xd0, 0x5b, 0x4e, 0xff, 0x17, 0xea, 0xa7, 0x09, 0xe5, 0xac, 0x9f,
	0x1a, 0x94, 0x06, 0x91, 0x14, 0x1d, 0x0f, 0x97, 0x2c, 0xa3, 0xbe, 0x4e, 0x12, 0x0b, 0x3d, 0x03,
	0x60, 0x52, 0xa8, 0x90, 0x32, 0xd5, 0xf1, 0xf0, 0x8a, 0x65, 0xd4, 0x57, 0xc9, 0x8c, 0x07, 0x55,
	0xc0, 0xf4, 0x3d, 0xcd, 0xb6, 0x4a, 0x4c, 0xdf, 0x43, 0x08, 0x8a, 0x82, 0x06, 0x1c, 0x9b, 0xda,
	0xa3, 0xdf, 0x31, 0x76, 0x34, 0x09, 0x06, 0xf2, 0x0a, 0x17, 0xb4, 0x37, 0xb1, 0x10, 0x86, 0x15,
	0x8f, 0x33, 0x3f, 0xa0, 0x57, 0xb8, 0x68, 0x19, 0xf5, 0x0d, 0x92, 0x9a, 0xe8, 0x18, 0x8a, 0x6a,
	0x32, 0xe4, 0x78, 0xd9, 0x32, 0xea, 0x95, 0xa3, 0xbd, 0x5b, 0x5d, 0x34, 0xfa, 0x69, 0x17, 0xee,
	0x64, 0xc8, 0x89, 0x4e, 0x8e, 0x69, 0xf4, 0x2c, 0x38, 0x2e, 0x5b, 0x46, 0xbd, 0x4c, 0x12, 0x0b,
	0x55, 0xa1, 0xe0, 0x0b, 0x86, 0x57, 0x35, 0x45, 0xfc, 0xb4, 0x7b, 0xb0, 0x31, 0x07, 0x80, 0x1e,
	0xc1, 0xa6, 0x43, 0xce, 0x3f, 0xba, 0xbd, 0xbe, 0xf3, 0x89, 0x90, 0x76, 0xd7, 0x39, 0xaf, 0x2e,
	0xa1, 0x2d, 0xd8, 0x78, 0xd7, 0x69, 0xb9, 0xb9, 0xcb, 0x40, 0x35, 0x40, 0x6d, 0xf7, 0x7d, 0xdf,
	0xe9, 0x75, 0x5d, 0xd2, 0x72, 0xdc, 0xbe, 0xdb, 0xfb, 0xd0, 0xee, 0x56, 0x4d, 0xfb, 0x12, 0xb6,
	0x4f, 0x47, 0x3c, 0x9c, 0xf4, 0x9d, 0x90, 0x53, 0xc5, 0xb3, 0xa9, 0xbe, 0x80, 0x92, 0x1c, 0x5c,
	0x72, 0xa6, 0xee, 0xd7, 0x23, 0x49, 0x40, 0x07, 0xb0, 0xec, 0x2b, 0x1e, 0x44, 0xd8, 0xb4, 0x0a,
	0x77, 0x67, 0x4e, 0xe3, 0xb6, 0x07, 0x95, 0x84, 0x2b, 0x65, 0x59, 0x4c, 0x83, 0xe2, 0x9c, 0x06,
	0xf9, 0xd0, 0x96, 0x67, 0x87, 0x66, 0xff, 0x36, 0xa0, 0x46, 0x78, 0x34, 0x94, 0x22, 0xe2, 0x19,
	0x93, 0xe6, 0x8d, 0x65, 0x63, 0xba, 0xcd, 0x29, 0x67, 0x99, 0xa4, 0xe6, 0x4c, 0xbb, 0xe6, 0x83,
	0xdb, 0x2d, 0xfc, 0xbb, 0xdd, 0x78, 0x01, 0xf5, 0xc3, 0x91, 0x23, 0xa1, 0x92, 0x3d, 0x99, 0xf1,
	0xa0, 0xd7, 0xb0, 0xf6, 0x25, 0x2e, 0xeb, 0x4c, 0x51, 0x35, 0x8a, 0x92, 0x8d, 0xd9, 0xce, 0xe0,
	0x4e, 0xf3, 0x18, 0x99, 0x4d, 0x3c, 0xba, 0x2e, 0xc0, 0xe6, 0x19, 0x0f, 0xc7, 0x3e, 0xcb, 0xe5,
	0x8a, 0xa0, 0x72, 0x43, 0xc0, 0xdd, 0x79, 0xa0, 0x1b, 0xfa, 0xee, 0xe4, 0x9b, 0x79, 0xf7, 0xac,
	0xec, 0xfd, 0x1f, 0xbf, 0xfe, 0x5c, 0x9b, 0x4f, 0xec, 0x5a, 0xb3, 0x3f, 0x7e, 0x75, 0x48, 0x9b,
	0xd3, 0x49, 0x65, 0xcb, 0xf7, 0xc6, 0x78, 0x89, 0x3a, 0xb0, 0x76, 0xc2, 0xb3, 0xff, 0x8c, 0x1e,
	0xdf, 0x64, 0x7c, 0x30, 0xd7, 0x12, 0xfa, 0x0c, 0xe5, 0xf4, 0xb0, 0xdc, 0x8f, 0x73, 0x7b, 0xd4,
	0xf6, 0x73, 0x5d, 0xe5, 0x1e, 0xda, 0x4d, 0xaa, 0x4c, 0x4f, 0x4a, 0x56, 0x67, 0xf3, 0x9b, 0xef,
	0x7d, 0x47, 0x14, 0x20, 0x3f, 0x37, 0x0b, 0x11, 0x1c, 0x68, 0x82, 0x7d, 0xb4, 0x97, 0x10, 0x78,
	0xfc, 0x6e, 0x8a, 0x41, 0x49, 0x5f, 0xdc, 0xe3, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x3a, 0xc2,
	0xe1, 0xea, 0xb8, 0x05, 0x00, 0x00,
}