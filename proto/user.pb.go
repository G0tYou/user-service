// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user.proto

package user

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type User struct {
	Id                     int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	FirstName              string   `protobuf:"bytes,2,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName               string   `protobuf:"bytes,3,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	Username               string   `protobuf:"bytes,4,opt,name=username,proto3" json:"username,omitempty"`
	Password               string   `protobuf:"bytes,5,opt,name=password,proto3" json:"password,omitempty"`
	EmailAddress           string   `protobuf:"bytes,6,opt,name=email_address,json=emailAddress,proto3" json:"email_address,omitempty"`
	PhoneNumber            string   `protobuf:"bytes,7,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"`
	DateOfBirth            string   `protobuf:"bytes,8,opt,name=date_of_birth,json=dateOfBirth,proto3" json:"date_of_birth,omitempty"`
	Address                string   `protobuf:"bytes,9,opt,name=address,proto3" json:"address,omitempty"`
	Role                   string   `protobuf:"bytes,10,opt,name=role,proto3" json:"role,omitempty"`
	CreditCardNumber       string   `protobuf:"bytes,11,opt,name=credit_card_number,json=creditCardNumber,proto3" json:"credit_card_number,omitempty"`
	CreditCardType         string   `protobuf:"bytes,12,opt,name=credit_card_type,json=creditCardType,proto3" json:"credit_card_type,omitempty"`
	CreditCardExpiredMonth string   `protobuf:"bytes,13,opt,name=credit_card_expired_month,json=creditCardExpiredMonth,proto3" json:"credit_card_expired_month,omitempty"`
	CreditCardExpiredYear  string   `protobuf:"bytes,14,opt,name=credit_card_expired_year,json=creditCardExpiredYear,proto3" json:"credit_card_expired_year,omitempty"`
	CreditCardCvv          string   `protobuf:"bytes,15,opt,name=credit_card_cvv,json=creditCardCvv,proto3" json:"credit_card_cvv,omitempty"`
	Status                 bool     `protobuf:"varint,16,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral   struct{} `json:"-"`
	XXX_unrecognized       []byte   `json:"-"`
	XXX_sizecache          int32    `json:"-"`
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

func (m *User) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *User) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *User) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

func (m *User) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *User) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *User) GetEmailAddress() string {
	if m != nil {
		return m.EmailAddress
	}
	return ""
}

func (m *User) GetPhoneNumber() string {
	if m != nil {
		return m.PhoneNumber
	}
	return ""
}

func (m *User) GetDateOfBirth() string {
	if m != nil {
		return m.DateOfBirth
	}
	return ""
}

func (m *User) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *User) GetRole() string {
	if m != nil {
		return m.Role
	}
	return ""
}

func (m *User) GetCreditCardNumber() string {
	if m != nil {
		return m.CreditCardNumber
	}
	return ""
}

func (m *User) GetCreditCardType() string {
	if m != nil {
		return m.CreditCardType
	}
	return ""
}

func (m *User) GetCreditCardExpiredMonth() string {
	if m != nil {
		return m.CreditCardExpiredMonth
	}
	return ""
}

func (m *User) GetCreditCardExpiredYear() string {
	if m != nil {
		return m.CreditCardExpiredYear
	}
	return ""
}

func (m *User) GetCreditCardCvv() string {
	if m != nil {
		return m.CreditCardCvv
	}
	return ""
}

func (m *User) GetStatus() bool {
	if m != nil {
		return m.Status
	}
	return false
}

//IndexUsersRequest is a message for requesting ll user
type IndexUsersRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IndexUsersRequest) Reset()         { *m = IndexUsersRequest{} }
func (m *IndexUsersRequest) String() string { return proto.CompactTextString(m) }
func (*IndexUsersRequest) ProtoMessage()    {}
func (*IndexUsersRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{1}
}

func (m *IndexUsersRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IndexUsersRequest.Unmarshal(m, b)
}
func (m *IndexUsersRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IndexUsersRequest.Marshal(b, m, deterministic)
}
func (m *IndexUsersRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IndexUsersRequest.Merge(m, src)
}
func (m *IndexUsersRequest) XXX_Size() int {
	return xxx_messageInfo_IndexUsersRequest.Size(m)
}
func (m *IndexUsersRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_IndexUsersRequest.DiscardUnknown(m)
}

var xxx_messageInfo_IndexUsersRequest proto.InternalMessageInfo

//Response is a message for a response
type Response struct {
	User                 *User    `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Error                *Error   `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	Users                []*User  `protobuf:"bytes,3,rep,name=users,proto3" json:"users,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{2}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *Response) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *Response) GetUsers() []*User {
	if m != nil {
		return m.Users
	}
	return nil
}

type Error struct {
	Code                 string   `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Error) Reset()         { *m = Error{} }
func (m *Error) String() string { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()    {}
func (*Error) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{3}
}

func (m *Error) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Error.Unmarshal(m, b)
}
func (m *Error) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Error.Marshal(b, m, deterministic)
}
func (m *Error) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Error.Merge(m, src)
}
func (m *Error) XXX_Size() int {
	return xxx_messageInfo_Error.Size(m)
}
func (m *Error) XXX_DiscardUnknown() {
	xxx_messageInfo_Error.DiscardUnknown(m)
}

var xxx_messageInfo_Error proto.InternalMessageInfo

func (m *Error) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *Error) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*User)(nil), "user.User")
	proto.RegisterType((*IndexUsersRequest)(nil), "user.IndexUsersRequest")
	proto.RegisterType((*Response)(nil), "user.Response")
	proto.RegisterType((*Error)(nil), "user.Error")
}

func init() { proto.RegisterFile("user.proto", fileDescriptor_116e343673f7ffaf) }

var fileDescriptor_116e343673f7ffaf = []byte{
	// 520 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0xcf, 0x6e, 0xd3, 0x40,
	0x10, 0xc6, 0x49, 0x9b, 0xa4, 0xf6, 0x38, 0x49, 0xc3, 0x20, 0xca, 0x52, 0x04, 0x4a, 0x8d, 0x84,
	0xc2, 0x1f, 0xf5, 0x10, 0x04, 0x88, 0x23, 0x94, 0x1e, 0x38, 0x50, 0x24, 0x87, 0x1e, 0x38, 0x59,
	0x9b, 0xec, 0x84, 0x58, 0x8a, 0xbd, 0x66, 0x77, 0x93, 0x36, 0x6f, 0xc5, 0xbb, 0xf1, 0x02, 0x68,
	0xc7, 0x35, 0x89, 0x68, 0x25, 0x7a, 0x9b, 0xf9, 0xbe, 0xdf, 0xb7, 0x23, 0x7b, 0x67, 0x01, 0x96,
	0x96, 0xcc, 0x71, 0x69, 0xb4, 0xd3, 0xd8, 0xf4, 0x75, 0xfc, 0xab, 0x09, 0xcd, 0x73, 0x4b, 0x06,
	0x7b, 0xb0, 0x93, 0x29, 0xd1, 0x18, 0x34, 0x86, 0xad, 0x64, 0x27, 0x53, 0xf8, 0x18, 0x60, 0x96,
	0x19, 0xeb, 0xd2, 0x42, 0xe6, 0x24, 0x76, 0x06, 0x8d, 0x61, 0x98, 0x84, 0xac, 0x9c, 0xc9, 0x9c,
	0xf0, 0x11, 0x84, 0x0b, 0x59, 0xbb, 0xbb, 0xec, 0x06, 0x5e, 0x60, 0xf3, 0x10, 0x02, 0x7f, 0x38,
	0x7b, 0xcd, 0xca, 0xab, 0x7b, 0xef, 0x95, 0xd2, 0xda, 0x0b, 0x6d, 0x94, 0x68, 0x55, 0x5e, 0xdd,
	0xe3, 0x53, 0xe8, 0x52, 0x2e, 0xb3, 0x45, 0x2a, 0x95, 0x32, 0x64, 0xad, 0x68, 0x33, 0xd0, 0x61,
	0xf1, 0x43, 0xa5, 0xe1, 0x11, 0x74, 0xca, 0xb9, 0x2e, 0x28, 0x2d, 0x96, 0xf9, 0x84, 0x8c, 0xd8,
	0x63, 0x26, 0x62, 0xed, 0x8c, 0x25, 0x8c, 0xa1, 0xab, 0xa4, 0xa3, 0x54, 0xcf, 0xd2, 0x49, 0x66,
	0xdc, 0x5c, 0x04, 0x15, 0xe3, 0xc5, 0xaf, 0xb3, 0x8f, 0x5e, 0x42, 0x01, 0x7b, 0xf5, 0x94, 0x90,
	0xdd, 0xba, 0x45, 0x84, 0xa6, 0xd1, 0x0b, 0x12, 0xc0, 0x32, 0xd7, 0xf8, 0x0a, 0x70, 0x6a, 0x48,
	0x65, 0x2e, 0x9d, 0x4a, 0xa3, 0xea, 0xd1, 0x11, 0x13, 0xfd, 0xca, 0x39, 0x91, 0x46, 0x5d, 0xcd,
	0x1f, 0x42, 0x7f, 0x9b, 0x76, 0xeb, 0x92, 0x44, 0x87, 0xd9, 0xde, 0x86, 0xfd, 0xb6, 0x2e, 0x09,
	0xdf, 0xc3, 0xc3, 0x6d, 0x92, 0x2e, 0xcb, 0xcc, 0x90, 0x4a, 0x73, 0x5d, 0xb8, 0xb9, 0xe8, 0x72,
	0xe4, 0x60, 0x13, 0x39, 0xad, 0xec, 0x2f, 0xde, 0xc5, 0x77, 0x20, 0x6e, 0x8a, 0xae, 0x49, 0x1a,
	0xd1, 0xe3, 0xe4, 0xfd, 0x6b, 0xc9, 0xef, 0x24, 0x0d, 0x3e, 0x83, 0xfd, 0xed, 0xe0, 0x74, 0xb5,
	0x12, 0xfb, 0xcc, 0x77, 0x37, 0xfc, 0xc9, 0x6a, 0x85, 0x07, 0xd0, 0xb6, 0x4e, 0xba, 0xa5, 0x15,
	0xfd, 0x41, 0x63, 0x18, 0x24, 0x57, 0x5d, 0x7c, 0x0f, 0xee, 0x7e, 0x2e, 0x14, 0x5d, 0xfa, 0xb5,
	0xb1, 0x09, 0xfd, 0x5c, 0x92, 0x75, 0xb1, 0x86, 0x20, 0x21, 0x5b, 0xea, 0xc2, 0x12, 0x3e, 0x01,
	0xde, 0x2d, 0x5e, 0xa6, 0x68, 0x04, 0xc7, 0xbc, 0x74, 0x9e, 0x4e, 0x58, 0xc7, 0x23, 0x68, 0x91,
	0x31, 0xda, 0xf0, 0x56, 0x45, 0xa3, 0xa8, 0x02, 0x4e, 0xbd, 0x94, 0x54, 0x0e, 0x0e, 0xa0, 0xe5,
	0x45, 0x2b, 0x76, 0x07, 0xbb, 0xff, 0x9c, 0x51, 0x19, 0xf1, 0x1b, 0x68, 0x71, 0xc2, 0x5f, 0xd7,
	0x54, 0x2b, 0xe2, 0x69, 0x61, 0xc2, 0xb5, 0xbf, 0xdc, 0x9c, 0xac, 0x95, 0x3f, 0xea, 0xcd, 0xad,
	0xdb, 0xd1, 0xef, 0x06, 0x44, 0xfe, 0x98, 0x31, 0x99, 0x55, 0x36, 0x25, 0x7c, 0x0e, 0xe1, 0xd8,
	0x69, 0x43, 0xfc, 0x06, 0xb6, 0xc6, 0x1c, 0xf6, 0xaa, 0xba, 0xfe, 0xa8, 0xf8, 0x0e, 0xbe, 0x85,
	0xf0, 0xef, 0x77, 0xe3, 0x83, 0xca, 0xbe, 0xf6, 0x23, 0x6e, 0xc8, 0x0d, 0x21, 0x18, 0xcf, 0xf5,
	0xc5, 0x2d, 0x26, 0xbc, 0x00, 0x38, 0x2f, 0xfd, 0x92, 0xde, 0x82, 0x7d, 0x09, 0xd1, 0x27, 0xb2,
	0xce, 0xe8, 0xf5, 0xff, 0xe1, 0x49, 0x9b, 0x9f, 0xfc, 0xeb, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff,
	0x3f, 0x94, 0x06, 0x11, 0x00, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserServiceClient interface {
	//Method to store a user
	StoreUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*Response, error)
	//Method to get all user
	IndexUser(ctx context.Context, in *IndexUsersRequest, opts ...grpc.CallOption) (*Response, error)
	//Method to show user
	ShowUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*Response, error)
	//Method to update user
	UpdateUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*Response, error)
	//Method to destroy user
	DestroyUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*Response, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) StoreUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/user.UserService/StoreUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) IndexUser(ctx context.Context, in *IndexUsersRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/user.UserService/IndexUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) ShowUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/user.UserService/ShowUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UpdateUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/user.UserService/UpdateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) DestroyUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/user.UserService/DestroyUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
type UserServiceServer interface {
	//Method to store a user
	StoreUser(context.Context, *User) (*Response, error)
	//Method to get all user
	IndexUser(context.Context, *IndexUsersRequest) (*Response, error)
	//Method to show user
	ShowUser(context.Context, *User) (*Response, error)
	//Method to update user
	UpdateUser(context.Context, *User) (*Response, error)
	//Method to destroy user
	DestroyUser(context.Context, *User) (*Response, error)
}

// UnimplementedUserServiceServer can be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (*UnimplementedUserServiceServer) StoreUser(ctx context.Context, req *User) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StoreUser not implemented")
}
func (*UnimplementedUserServiceServer) IndexUser(ctx context.Context, req *IndexUsersRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IndexUser not implemented")
}
func (*UnimplementedUserServiceServer) ShowUser(ctx context.Context, req *User) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ShowUser not implemented")
}
func (*UnimplementedUserServiceServer) UpdateUser(ctx context.Context, req *User) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}
func (*UnimplementedUserServiceServer) DestroyUser(ctx context.Context, req *User) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DestroyUser not implemented")
}

func RegisterUserServiceServer(s *grpc.Server, srv UserServiceServer) {
	s.RegisterService(&_UserService_serviceDesc, srv)
}

func _UserService_StoreUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).StoreUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/StoreUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).StoreUser(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_IndexUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IndexUsersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).IndexUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/IndexUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).IndexUser(ctx, req.(*IndexUsersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_ShowUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).ShowUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/ShowUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).ShowUser(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/UpdateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UpdateUser(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_DestroyUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).DestroyUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/DestroyUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).DestroyUser(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "user.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "StoreUser",
			Handler:    _UserService_StoreUser_Handler,
		},
		{
			MethodName: "IndexUser",
			Handler:    _UserService_IndexUser_Handler,
		},
		{
			MethodName: "ShowUser",
			Handler:    _UserService_ShowUser_Handler,
		},
		{
			MethodName: "UpdateUser",
			Handler:    _UserService_UpdateUser_Handler,
		},
		{
			MethodName: "DestroyUser",
			Handler:    _UserService_DestroyUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}
