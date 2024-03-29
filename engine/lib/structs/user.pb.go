// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user.proto

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

type UserStatus int32

const (
	UserStatus_UNVERIFIED    UserStatus = 0
	UserStatus_VERIFIED      UserStatus = 1
	UserStatus_ADMINISTRATOR UserStatus = 2
)

var UserStatus_name = map[int32]string{
	0: "UNVERIFIED",
	1: "VERIFIED",
	2: "ADMINISTRATOR",
}

var UserStatus_value = map[string]int32{
	"UNVERIFIED":    0,
	"VERIFIED":      1,
	"ADMINISTRATOR": 2,
}

func (x UserStatus) String() string {
	return proto.EnumName(UserStatus_name, int32(x))
}

func (UserStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{0}
}

type User struct {
	// @inject_tag: bson:"_id,omitempty"
	BsonId   []byte     `protobuf:"bytes,1,opt,name=bsonId,proto3" json:"bsonId,omitempty" bson:"_id,omitempty"`
	Id       string     `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Name     string     `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Email    string     `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	Password string     `protobuf:"bytes,5,opt,name=password,proto3" json:"password,omitempty"`
	Status   UserStatus `protobuf:"varint,6,opt,name=status,proto3,enum=structs.UserStatus" json:"status,omitempty"`
	// Может быть много полей
	CreatedAt string `protobuf:"bytes,7,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	// Инкремент
	Inc                  uint32   `protobuf:"varint,8,opt,name=inc,proto3" json:"inc,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{0}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetBsonId() []byte {
	if m != nil {
		return m.BsonId
	}
	return nil
}

func (m *User) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *User) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *User) GetStatus() UserStatus {
	if m != nil {
		return m.Status
	}
	return UserStatus_UNVERIFIED
}

func (m *User) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *User) GetInc() uint32 {
	if m != nil {
		return m.Inc
	}
	return 0
}

type Query_User struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Query_User) Reset()         { *m = Query_User{} }
func (m *Query_User) String() string { return proto.CompactTextString(m) }
func (*Query_User) ProtoMessage()    {}
func (*Query_User) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{1}
}

func (m *Query_User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Query_User.Unmarshal(m, b)
}
func (m *Query_User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Query_User.Marshal(b, m, deterministic)
}
func (m *Query_User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Query_User.Merge(m, src)
}
func (m *Query_User) XXX_Size() int {
	return xxx_messageInfo_Query_User.Size(m)
}
func (m *Query_User) XXX_DiscardUnknown() {
	xxx_messageInfo_Query_User.DiscardUnknown(m)
}

var xxx_messageInfo_Query_User proto.InternalMessageInfo

func (m *Query_User) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type Response_User struct {
	Created              bool        `protobuf:"varint,1,opt,name=created,proto3" json:"created,omitempty"`
	Object               *User       `protobuf:"bytes,2,opt,name=object,proto3" json:"object,omitempty"`
	Items                []*User     `protobuf:"bytes,3,rep,name=items,proto3" json:"items,omitempty"`
	ItemsCount           uint32      `protobuf:"varint,4,opt,name=itemsCount,proto3" json:"itemsCount,omitempty"`
	QueryStatus          QueryStatus `protobuf:"varint,5,opt,name=queryStatus,proto3,enum=structs.QueryStatus" json:"queryStatus,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Response_User) Reset()         { *m = Response_User{} }
func (m *Response_User) String() string { return proto.CompactTextString(m) }
func (*Response_User) ProtoMessage()    {}
func (*Response_User) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{2}
}

func (m *Response_User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response_User.Unmarshal(m, b)
}
func (m *Response_User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response_User.Marshal(b, m, deterministic)
}
func (m *Response_User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response_User.Merge(m, src)
}
func (m *Response_User) XXX_Size() int {
	return xxx_messageInfo_Response_User.Size(m)
}
func (m *Response_User) XXX_DiscardUnknown() {
	xxx_messageInfo_Response_User.DiscardUnknown(m)
}

var xxx_messageInfo_Response_User proto.InternalMessageInfo

func (m *Response_User) GetCreated() bool {
	if m != nil {
		return m.Created
	}
	return false
}

func (m *Response_User) GetObject() *User {
	if m != nil {
		return m.Object
	}
	return nil
}

func (m *Response_User) GetItems() []*User {
	if m != nil {
		return m.Items
	}
	return nil
}

func (m *Response_User) GetItemsCount() uint32 {
	if m != nil {
		return m.ItemsCount
	}
	return 0
}

func (m *Response_User) GetQueryStatus() QueryStatus {
	if m != nil {
		return m.QueryStatus
	}
	return QueryStatus_Query_Missed
}

type EventNewUser struct {
	User                 *User    `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EventNewUser) Reset()         { *m = EventNewUser{} }
func (m *EventNewUser) String() string { return proto.CompactTextString(m) }
func (*EventNewUser) ProtoMessage()    {}
func (*EventNewUser) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{3}
}

func (m *EventNewUser) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventNewUser.Unmarshal(m, b)
}
func (m *EventNewUser) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventNewUser.Marshal(b, m, deterministic)
}
func (m *EventNewUser) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventNewUser.Merge(m, src)
}
func (m *EventNewUser) XXX_Size() int {
	return xxx_messageInfo_EventNewUser.Size(m)
}
func (m *EventNewUser) XXX_DiscardUnknown() {
	xxx_messageInfo_EventNewUser.DiscardUnknown(m)
}

var xxx_messageInfo_EventNewUser proto.InternalMessageInfo

func (m *EventNewUser) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

type Authenticate struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	Email                string   `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Authenticate) Reset()         { *m = Authenticate{} }
func (m *Authenticate) String() string { return proto.CompactTextString(m) }
func (*Authenticate) ProtoMessage()    {}
func (*Authenticate) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{4}
}

func (m *Authenticate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Authenticate.Unmarshal(m, b)
}
func (m *Authenticate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Authenticate.Marshal(b, m, deterministic)
}
func (m *Authenticate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Authenticate.Merge(m, src)
}
func (m *Authenticate) XXX_Size() int {
	return xxx_messageInfo_Authenticate.Size(m)
}
func (m *Authenticate) XXX_DiscardUnknown() {
	xxx_messageInfo_Authenticate.DiscardUnknown(m)
}

var xxx_messageInfo_Authenticate proto.InternalMessageInfo

func (m *Authenticate) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Authenticate) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *Authenticate) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

type Session struct {
	SessionId            string   `protobuf:"bytes,1,opt,name=SessionId,proto3" json:"SessionId,omitempty"`
	User                 *User    `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	Token                string   `protobuf:"bytes,3,opt,name=Token,proto3" json:"Token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Session) Reset()         { *m = Session{} }
func (m *Session) String() string { return proto.CompactTextString(m) }
func (*Session) ProtoMessage()    {}
func (*Session) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{5}
}

func (m *Session) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Session.Unmarshal(m, b)
}
func (m *Session) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Session.Marshal(b, m, deterministic)
}
func (m *Session) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Session.Merge(m, src)
}
func (m *Session) XXX_Size() int {
	return xxx_messageInfo_Session.Size(m)
}
func (m *Session) XXX_DiscardUnknown() {
	xxx_messageInfo_Session.DiscardUnknown(m)
}

var xxx_messageInfo_Session proto.InternalMessageInfo

func (m *Session) GetSessionId() string {
	if m != nil {
		return m.SessionId
	}
	return ""
}

func (m *Session) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *Session) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func init() {
	proto.RegisterEnum("structs.UserStatus", UserStatus_name, UserStatus_value)
	proto.RegisterType((*User)(nil), "structs.User")
	proto.RegisterType((*Query_User)(nil), "structs.Query_User")
	proto.RegisterType((*Response_User)(nil), "structs.Response_User")
	proto.RegisterType((*EventNewUser)(nil), "structs.EventNewUser")
	proto.RegisterType((*Authenticate)(nil), "structs.Authenticate")
	proto.RegisterType((*Session)(nil), "structs.Session")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ServiceUserClient is the client API for ServiceUser service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ServiceUserClient interface {
	CreateUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*Response_User, error)
	GetUser(ctx context.Context, in *Query_User, opts ...grpc.CallOption) (*Response_User, error)
	FindByNamePassword(ctx context.Context, in *Authenticate, opts ...grpc.CallOption) (*Response_User, error)
	GetInfo(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*User, error)
	GetAllUsers(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Response_User, error)
}

type serviceUserClient struct {
	cc *grpc.ClientConn
}

func NewServiceUserClient(cc *grpc.ClientConn) ServiceUserClient {
	return &serviceUserClient{cc}
}

func (c *serviceUserClient) CreateUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*Response_User, error) {
	out := new(Response_User)
	err := c.cc.Invoke(ctx, "/structs.ServiceUser/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceUserClient) GetUser(ctx context.Context, in *Query_User, opts ...grpc.CallOption) (*Response_User, error) {
	out := new(Response_User)
	err := c.cc.Invoke(ctx, "/structs.ServiceUser/GetUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceUserClient) FindByNamePassword(ctx context.Context, in *Authenticate, opts ...grpc.CallOption) (*Response_User, error) {
	out := new(Response_User)
	err := c.cc.Invoke(ctx, "/structs.ServiceUser/FindByNamePassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceUserClient) GetInfo(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/structs.ServiceUser/GetInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceUserClient) GetAllUsers(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Response_User, error) {
	out := new(Response_User)
	err := c.cc.Invoke(ctx, "/structs.ServiceUser/GetAllUsers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServiceUserServer is the server API for ServiceUser service.
type ServiceUserServer interface {
	CreateUser(context.Context, *User) (*Response_User, error)
	GetUser(context.Context, *Query_User) (*Response_User, error)
	FindByNamePassword(context.Context, *Authenticate) (*Response_User, error)
	GetInfo(context.Context, *Empty) (*User, error)
	GetAllUsers(context.Context, *Empty) (*Response_User, error)
}

func RegisterServiceUserServer(s *grpc.Server, srv ServiceUserServer) {
	s.RegisterService(&_ServiceUser_serviceDesc, srv)
}

func _ServiceUser_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceUserServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/structs.ServiceUser/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceUserServer).CreateUser(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceUser_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Query_User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceUserServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/structs.ServiceUser/GetUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceUserServer).GetUser(ctx, req.(*Query_User))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceUser_FindByNamePassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Authenticate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceUserServer).FindByNamePassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/structs.ServiceUser/FindByNamePassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceUserServer).FindByNamePassword(ctx, req.(*Authenticate))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceUser_GetInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceUserServer).GetInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/structs.ServiceUser/GetInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceUserServer).GetInfo(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceUser_GetAllUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceUserServer).GetAllUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/structs.ServiceUser/GetAllUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceUserServer).GetAllUsers(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _ServiceUser_serviceDesc = grpc.ServiceDesc{
	ServiceName: "structs.ServiceUser",
	HandlerType: (*ServiceUserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUser",
			Handler:    _ServiceUser_CreateUser_Handler,
		},
		{
			MethodName: "GetUser",
			Handler:    _ServiceUser_GetUser_Handler,
		},
		{
			MethodName: "FindByNamePassword",
			Handler:    _ServiceUser_FindByNamePassword_Handler,
		},
		{
			MethodName: "GetInfo",
			Handler:    _ServiceUser_GetInfo_Handler,
		},
		{
			MethodName: "GetAllUsers",
			Handler:    _ServiceUser_GetAllUsers_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}

// ServiceAuthClient is the client API for ServiceAuth service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ServiceAuthClient interface {
	SignUp(ctx context.Context, in *Authenticate, opts ...grpc.CallOption) (*Session, error)
	StartSession(ctx context.Context, in *Authenticate, opts ...grpc.CallOption) (*Session, error)
	CloseSession(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Bool, error)
	DecodeSession(ctx context.Context, in *Session, opts ...grpc.CallOption) (*User, error)
}

type serviceAuthClient struct {
	cc *grpc.ClientConn
}

func NewServiceAuthClient(cc *grpc.ClientConn) ServiceAuthClient {
	return &serviceAuthClient{cc}
}

func (c *serviceAuthClient) SignUp(ctx context.Context, in *Authenticate, opts ...grpc.CallOption) (*Session, error) {
	out := new(Session)
	err := c.cc.Invoke(ctx, "/structs.ServiceAuth/SignUp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceAuthClient) StartSession(ctx context.Context, in *Authenticate, opts ...grpc.CallOption) (*Session, error) {
	out := new(Session)
	err := c.cc.Invoke(ctx, "/structs.ServiceAuth/StartSession", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceAuthClient) CloseSession(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Bool, error) {
	out := new(Bool)
	err := c.cc.Invoke(ctx, "/structs.ServiceAuth/CloseSession", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceAuthClient) DecodeSession(ctx context.Context, in *Session, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/structs.ServiceAuth/DecodeSession", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServiceAuthServer is the server API for ServiceAuth service.
type ServiceAuthServer interface {
	SignUp(context.Context, *Authenticate) (*Session, error)
	StartSession(context.Context, *Authenticate) (*Session, error)
	CloseSession(context.Context, *Empty) (*Bool, error)
	DecodeSession(context.Context, *Session) (*User, error)
}

func RegisterServiceAuthServer(s *grpc.Server, srv ServiceAuthServer) {
	s.RegisterService(&_ServiceAuth_serviceDesc, srv)
}

func _ServiceAuth_SignUp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Authenticate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceAuthServer).SignUp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/structs.ServiceAuth/SignUp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceAuthServer).SignUp(ctx, req.(*Authenticate))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceAuth_StartSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Authenticate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceAuthServer).StartSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/structs.ServiceAuth/StartSession",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceAuthServer).StartSession(ctx, req.(*Authenticate))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceAuth_CloseSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceAuthServer).CloseSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/structs.ServiceAuth/CloseSession",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceAuthServer).CloseSession(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceAuth_DecodeSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Session)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceAuthServer).DecodeSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/structs.ServiceAuth/DecodeSession",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceAuthServer).DecodeSession(ctx, req.(*Session))
	}
	return interceptor(ctx, in, info, handler)
}

var _ServiceAuth_serviceDesc = grpc.ServiceDesc{
	ServiceName: "structs.ServiceAuth",
	HandlerType: (*ServiceAuthServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SignUp",
			Handler:    _ServiceAuth_SignUp_Handler,
		},
		{
			MethodName: "StartSession",
			Handler:    _ServiceAuth_StartSession_Handler,
		},
		{
			MethodName: "CloseSession",
			Handler:    _ServiceAuth_CloseSession_Handler,
		},
		{
			MethodName: "DecodeSession",
			Handler:    _ServiceAuth_DecodeSession_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}

func init() { proto.RegisterFile("user.proto", fileDescriptor_116e343673f7ffaf) }

var fileDescriptor_116e343673f7ffaf = []byte{
	// 691 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0x6d, 0x6f, 0xd2, 0x40,
	0x1c, 0x5f, 0x0b, 0x03, 0xf6, 0x87, 0x22, 0xbb, 0x3d, 0xa4, 0x21, 0xc4, 0x60, 0x8d, 0x09, 0x99,
	0xc9, 0xc8, 0x30, 0x6a, 0x62, 0x62, 0x0c, 0xdb, 0xd8, 0xec, 0x8b, 0xa1, 0x2b, 0xe0, 0x5b, 0xec,
	0xda, 0x13, 0x4f, 0xe9, 0x5d, 0xed, 0x5d, 0x59, 0xf6, 0xd6, 0xaf, 0xe0, 0xe7, 0xf0, 0xa3, 0xf8,
	0xc2, 0xe8, 0x47, 0xf0, 0x83, 0x98, 0x5e, 0x1f, 0x80, 0x21, 0x8b, 0xef, 0xee, 0xff, 0x70, 0xbf,
	0xfe, 0x1e, 0x2e, 0x05, 0x08, 0x39, 0x0e, 0x0e, 0xfd, 0x80, 0x09, 0x86, 0x8a, 0x5c, 0x04, 0xa1,
	0x23, 0x78, 0xbd, 0xe2, 0x30, 0xcf, 0x63, 0x34, 0x6e, 0xd7, 0x1b, 0x13, 0xc6, 0x26, 0x53, 0xdc,
	0xb6, 0x7d, 0xd2, 0xb6, 0x29, 0x65, 0xc2, 0x16, 0x84, 0x51, 0x1e, 0x4f, 0x8d, 0x9f, 0x0a, 0xe4,
	0x47, 0x1c, 0x07, 0x68, 0x1f, 0x0a, 0x57, 0x9c, 0x51, 0xd3, 0xd5, 0x95, 0xa6, 0xd2, 0xaa, 0x58,
	0x49, 0x85, 0xaa, 0xa0, 0x12, 0x57, 0x57, 0x9b, 0x4a, 0x6b, 0xcb, 0x52, 0x89, 0x8b, 0x10, 0xe4,
	0xa9, 0xed, 0x61, 0x3d, 0x27, 0x3b, 0xf2, 0x8c, 0x76, 0x61, 0x13, 0x7b, 0x36, 0x99, 0xea, 0x79,
	0xd9, 0x8c, 0x0b, 0x54, 0x87, 0x92, 0x6f, 0x73, 0x7e, 0xcd, 0x02, 0x57, 0xdf, 0x94, 0x83, 0xac,
	0x46, 0x8f, 0xa1, 0xc0, 0x85, 0x2d, 0x42, 0xae, 0x17, 0x9a, 0x4a, 0xab, 0xda, 0xd9, 0x39, 0x4c,
	0xc8, 0x1f, 0x46, 0x64, 0x06, 0x72, 0x64, 0x25, 0x2b, 0xa8, 0x01, 0x5b, 0x4e, 0x80, 0x6d, 0x81,
	0xdd, 0xae, 0xd0, 0x8b, 0x12, 0x69, 0xde, 0x40, 0x35, 0xc8, 0x11, 0xea, 0xe8, 0xa5, 0xa6, 0xd2,
	0xd2, 0xac, 0xe8, 0x68, 0x34, 0x00, 0x2e, 0x43, 0x1c, 0xdc, 0x8c, 0xa5, 0xb0, 0x58, 0x80, 0x92,
	0x0a, 0x30, 0x7e, 0x28, 0xa0, 0x59, 0x98, 0xfb, 0x8c, 0x72, 0x1c, 0x6f, 0xe8, 0x50, 0x4c, 0xe0,
	0xe4, 0x5a, 0xc9, 0x4a, 0x4b, 0xf4, 0x08, 0x0a, 0xec, 0xea, 0x13, 0x76, 0x84, 0x34, 0xa0, 0xdc,
	0xd1, 0x96, 0x68, 0x5a, 0xc9, 0x10, 0x3d, 0x84, 0x4d, 0x22, 0xb0, 0xc7, 0xf5, 0x5c, 0x33, 0xb7,
	0xba, 0x15, 0xcf, 0xd0, 0x7d, 0x00, 0x79, 0x38, 0x61, 0x21, 0x15, 0xd2, 0x29, 0xcd, 0x5a, 0xe8,
	0xa0, 0x67, 0x50, 0xfe, 0x12, 0xb1, 0x8e, 0xc5, 0x4b, 0xc7, 0xaa, 0x9d, 0xdd, 0x0c, 0xea, 0x72,
	0x3e, 0xb3, 0x16, 0x17, 0x8d, 0x23, 0xa8, 0xf4, 0x66, 0x98, 0x8a, 0x3e, 0xbe, 0x96, 0x6a, 0x1e,
	0x40, 0x3e, 0x7a, 0x14, 0x52, 0xca, 0x0a, 0x17, 0x39, 0x32, 0x86, 0x50, 0xe9, 0x86, 0xe2, 0x23,
	0xa6, 0x82, 0x38, 0xb6, 0xc0, 0x59, 0xa6, 0xca, 0x42, 0xa6, 0x8b, 0xe9, 0xa9, 0xb7, 0xd2, 0xcb,
	0xf2, 0xce, 0x2d, 0xe4, 0x6d, 0xbc, 0x87, 0xe2, 0x00, 0x73, 0x4e, 0x18, 0x8d, 0x12, 0x4b, 0x8e,
	0x66, 0x6a, 0xfd, 0xbc, 0x91, 0x31, 0x54, 0xd7, 0x32, 0x8c, 0xbe, 0x30, 0x64, 0x9f, 0x31, 0x4d,
	0xbf, 0x20, 0x8b, 0x83, 0x97, 0x00, 0xf3, 0xe7, 0x81, 0xaa, 0x00, 0xa3, 0xfe, 0xbb, 0x9e, 0x65,
	0x9e, 0x99, 0xbd, 0xd3, 0xda, 0x06, 0xaa, 0x40, 0x29, 0xab, 0x14, 0xb4, 0x0d, 0x5a, 0xf7, 0xf4,
	0xc2, 0xec, 0x9b, 0x83, 0xa1, 0xd5, 0x1d, 0xbe, 0xb1, 0x6a, 0x6a, 0xe7, 0xb7, 0x0a, 0xe5, 0x01,
	0x0e, 0x66, 0xc4, 0xc1, 0xd2, 0xa9, 0xa7, 0x00, 0x27, 0x32, 0x68, 0x59, 0x2d, 0xf3, 0xa8, 0xef,
	0x67, 0xe5, 0xd2, 0x63, 0x31, 0x36, 0xd0, 0x05, 0x14, 0xcf, 0xb1, 0x90, 0x77, 0x76, 0x96, 0xe3,
	0x19, 0xdf, 0x79, 0x73, 0xef, 0xeb, 0xaf, 0x3f, 0xdf, 0xd4, 0x7b, 0x48, 0x6b, 0x8f, 0x67, 0x47,
	0xed, 0x09, 0x16, 0x63, 0x29, 0xb5, 0x07, 0xe8, 0x8c, 0x50, 0xf7, 0xf8, 0xa6, 0x6f, 0x7b, 0xf8,
	0x6d, 0x6a, 0xf1, 0x5e, 0x06, 0xb2, 0x98, 0xd4, 0x1d, 0xac, 0x5e, 0x49, 0x56, 0x26, 0xfd, 0xc0,
	0x50, 0x35, 0x5b, 0xea, 0x79, 0xbe, 0xb8, 0xa9, 0x2f, 0x2b, 0xfb, 0x07, 0x0f, 0x12, 0xdd, 0x7a,
	0x0e, 0xe5, 0x73, 0x2c, 0xba, 0xd3, 0x69, 0xb4, 0xc4, 0x57, 0x40, 0xd6, 0x7e, 0xb9, 0xf3, 0x7d,
	0x6e, 0x6b, 0xc4, 0x15, 0xbd, 0x86, 0xc2, 0x80, 0x4c, 0xe8, 0xc8, 0x5f, 0x27, 0xa2, 0x96, 0xb5,
	0x93, 0x37, 0x91, 0x52, 0x32, 0xa0, 0x3d, 0x3b, 0x6a, 0x73, 0x32, 0xa1, 0xa1, 0xff, 0x42, 0x39,
	0x40, 0x23, 0xa8, 0x0c, 0x84, 0x1d, 0x88, 0xf4, 0x59, 0xfd, 0x37, 0x5e, 0x43, 0xe2, 0xed, 0x1b,
	0xdb, 0x12, 0x2f, 0x82, 0x18, 0xf3, 0x78, 0x14, 0xc1, 0x9a, 0x50, 0x39, 0x99, 0x32, 0x8e, 0x53,
	0xd8, 0xf5, 0x7e, 0x1d, 0x33, 0x36, 0x35, 0xea, 0x12, 0x6c, 0x17, 0x21, 0xe9, 0x97, 0x13, 0xdd,
	0x4c, 0xd1, 0x50, 0x07, 0xb4, 0x53, 0xec, 0x30, 0x37, 0xc3, 0x5a, 0xe1, 0x72, 0xdb, 0xfd, 0x8d,
	0xab, 0x82, 0xfc, 0xf3, 0x3e, 0xf9, 0x1b, 0x00, 0x00, 0xff, 0xff, 0xfe, 0x8d, 0xb6, 0xaa, 0xbc,
	0x05, 0x00, 0x00,
}
