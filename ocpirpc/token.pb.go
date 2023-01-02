// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ocpirpc/token.proto

package ocpirpc

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

type CreateTokenRequest struct {
	UserId               int64    `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Uid                  string   `protobuf:"bytes,2,opt,name=uid,proto3" json:"uid,omitempty"`
	Type                 string   `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	Allowed              string   `protobuf:"bytes,4,opt,name=allowed,proto3" json:"allowed,omitempty"`
	Whitelist            string   `protobuf:"bytes,5,opt,name=whitelist,proto3" json:"whitelist,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateTokenRequest) Reset()         { *m = CreateTokenRequest{} }
func (m *CreateTokenRequest) String() string { return proto.CompactTextString(m) }
func (*CreateTokenRequest) ProtoMessage()    {}
func (*CreateTokenRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_68c231a783882be6, []int{0}
}

func (m *CreateTokenRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateTokenRequest.Unmarshal(m, b)
}
func (m *CreateTokenRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateTokenRequest.Marshal(b, m, deterministic)
}
func (m *CreateTokenRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateTokenRequest.Merge(m, src)
}
func (m *CreateTokenRequest) XXX_Size() int {
	return xxx_messageInfo_CreateTokenRequest.Size(m)
}
func (m *CreateTokenRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateTokenRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateTokenRequest proto.InternalMessageInfo

func (m *CreateTokenRequest) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *CreateTokenRequest) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *CreateTokenRequest) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *CreateTokenRequest) GetAllowed() string {
	if m != nil {
		return m.Allowed
	}
	return ""
}

func (m *CreateTokenRequest) GetWhitelist() string {
	if m != nil {
		return m.Whitelist
	}
	return ""
}

type CreateTokenResponse struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Uid                  string   `protobuf:"bytes,2,opt,name=uid,proto3" json:"uid,omitempty"`
	Type                 string   `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	AuthId               string   `protobuf:"bytes,4,opt,name=auth_id,json=authId,proto3" json:"auth_id,omitempty"`
	VisualNumber         string   `protobuf:"bytes,5,opt,name=visual_number,json=visualNumber,proto3" json:"visual_number,omitempty"`
	Allowed              string   `protobuf:"bytes,6,opt,name=allowed,proto3" json:"allowed,omitempty"`
	Whitelist            string   `protobuf:"bytes,7,opt,name=whitelist,proto3" json:"whitelist,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateTokenResponse) Reset()         { *m = CreateTokenResponse{} }
func (m *CreateTokenResponse) String() string { return proto.CompactTextString(m) }
func (*CreateTokenResponse) ProtoMessage()    {}
func (*CreateTokenResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_68c231a783882be6, []int{1}
}

func (m *CreateTokenResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateTokenResponse.Unmarshal(m, b)
}
func (m *CreateTokenResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateTokenResponse.Marshal(b, m, deterministic)
}
func (m *CreateTokenResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateTokenResponse.Merge(m, src)
}
func (m *CreateTokenResponse) XXX_Size() int {
	return xxx_messageInfo_CreateTokenResponse.Size(m)
}
func (m *CreateTokenResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateTokenResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateTokenResponse proto.InternalMessageInfo

func (m *CreateTokenResponse) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *CreateTokenResponse) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *CreateTokenResponse) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *CreateTokenResponse) GetAuthId() string {
	if m != nil {
		return m.AuthId
	}
	return ""
}

func (m *CreateTokenResponse) GetVisualNumber() string {
	if m != nil {
		return m.VisualNumber
	}
	return ""
}

func (m *CreateTokenResponse) GetAllowed() string {
	if m != nil {
		return m.Allowed
	}
	return ""
}

func (m *CreateTokenResponse) GetWhitelist() string {
	if m != nil {
		return m.Whitelist
	}
	return ""
}

type UpdateTokensRequest struct {
	UserId               int64    `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Uid                  string   `protobuf:"bytes,2,opt,name=uid,proto3" json:"uid,omitempty"`
	Allowed              string   `protobuf:"bytes,3,opt,name=allowed,proto3" json:"allowed,omitempty"`
	Whitelist            string   `protobuf:"bytes,4,opt,name=whitelist,proto3" json:"whitelist,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateTokensRequest) Reset()         { *m = UpdateTokensRequest{} }
func (m *UpdateTokensRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateTokensRequest) ProtoMessage()    {}
func (*UpdateTokensRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_68c231a783882be6, []int{2}
}

func (m *UpdateTokensRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateTokensRequest.Unmarshal(m, b)
}
func (m *UpdateTokensRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateTokensRequest.Marshal(b, m, deterministic)
}
func (m *UpdateTokensRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateTokensRequest.Merge(m, src)
}
func (m *UpdateTokensRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateTokensRequest.Size(m)
}
func (m *UpdateTokensRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateTokensRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateTokensRequest proto.InternalMessageInfo

func (m *UpdateTokensRequest) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *UpdateTokensRequest) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *UpdateTokensRequest) GetAllowed() string {
	if m != nil {
		return m.Allowed
	}
	return ""
}

func (m *UpdateTokensRequest) GetWhitelist() string {
	if m != nil {
		return m.Whitelist
	}
	return ""
}

type UpdateTokensResponse struct {
	UserId               int64    `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Uid                  string   `protobuf:"bytes,2,opt,name=uid,proto3" json:"uid,omitempty"`
	Allowed              string   `protobuf:"bytes,3,opt,name=allowed,proto3" json:"allowed,omitempty"`
	Whitelist            string   `protobuf:"bytes,4,opt,name=whitelist,proto3" json:"whitelist,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateTokensResponse) Reset()         { *m = UpdateTokensResponse{} }
func (m *UpdateTokensResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateTokensResponse) ProtoMessage()    {}
func (*UpdateTokensResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_68c231a783882be6, []int{3}
}

func (m *UpdateTokensResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateTokensResponse.Unmarshal(m, b)
}
func (m *UpdateTokensResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateTokensResponse.Marshal(b, m, deterministic)
}
func (m *UpdateTokensResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateTokensResponse.Merge(m, src)
}
func (m *UpdateTokensResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateTokensResponse.Size(m)
}
func (m *UpdateTokensResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateTokensResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateTokensResponse proto.InternalMessageInfo

func (m *UpdateTokensResponse) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *UpdateTokensResponse) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *UpdateTokensResponse) GetAllowed() string {
	if m != nil {
		return m.Allowed
	}
	return ""
}

func (m *UpdateTokensResponse) GetWhitelist() string {
	if m != nil {
		return m.Whitelist
	}
	return ""
}

func init() {
	proto.RegisterType((*CreateTokenRequest)(nil), "token.CreateTokenRequest")
	proto.RegisterType((*CreateTokenResponse)(nil), "token.CreateTokenResponse")
	proto.RegisterType((*UpdateTokensRequest)(nil), "token.UpdateTokensRequest")
	proto.RegisterType((*UpdateTokensResponse)(nil), "token.UpdateTokensResponse")
}

func init() { proto.RegisterFile("ocpirpc/token.proto", fileDescriptor_68c231a783882be6) }

var fileDescriptor_68c231a783882be6 = []byte{
	// 344 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x53, 0xc1, 0x4e, 0xc2, 0x40,
	0x10, 0x4d, 0x29, 0xb4, 0x61, 0x44, 0x63, 0x06, 0x13, 0x2a, 0x7a, 0x20, 0x10, 0x13, 0x2e, 0xd2,
	0x44, 0xff, 0x40, 0x4d, 0x0c, 0x17, 0x0f, 0xa8, 0x17, 0x2f, 0xa4, 0xb4, 0x13, 0xd8, 0x58, 0xd8,
	0xda, 0xdd, 0x05, 0xf9, 0x06, 0xbf, 0xc1, 0x8f, 0xf1, 0xcf, 0xcc, 0x6e, 0x8b, 0x16, 0x24, 0xc4,
	0x98, 0x78, 0x9b, 0xf7, 0xa6, 0xe9, 0x7b, 0x6f, 0x76, 0x06, 0xea, 0x3c, 0x4c, 0x58, 0x9a, 0x84,
	0xbe, 0xe4, 0xcf, 0x34, 0xeb, 0x25, 0x29, 0x97, 0x1c, 0x2b, 0x06, 0xb4, 0xdf, 0x2c, 0xc0, 0xeb,
	0x94, 0x02, 0x49, 0x0f, 0x1a, 0x0f, 0xe8, 0x45, 0x91, 0x90, 0xd8, 0x00, 0x57, 0x09, 0x4a, 0x87,
	0x2c, 0xf2, 0xac, 0x96, 0xd5, 0xb5, 0x07, 0x8e, 0x86, 0xfd, 0x08, 0x0f, 0xc1, 0x56, 0x2c, 0xf2,
	0x4a, 0x2d, 0xab, 0x5b, 0x1d, 0xe8, 0x12, 0x11, 0xca, 0x72, 0x99, 0x90, 0x67, 0x1b, 0xca, 0xd4,
	0xe8, 0x81, 0x1b, 0xc4, 0x31, 0x5f, 0x50, 0xe4, 0x95, 0x0d, 0xbd, 0x82, 0x78, 0x0a, 0xd5, 0xc5,
	0x84, 0x49, 0x8a, 0x99, 0x90, 0x5e, 0xc5, 0xf4, 0xbe, 0x89, 0xf6, 0x87, 0x05, 0xf5, 0x35, 0x37,
	0x22, 0xe1, 0x33, 0x41, 0x78, 0x00, 0xa5, 0x2f, 0x27, 0x25, 0xf6, 0x5b, 0x17, 0x0d, 0x70, 0x03,
	0x25, 0x27, 0x3a, 0x44, 0xe6, 0xc2, 0xd1, 0xb0, 0x1f, 0x61, 0x07, 0xf6, 0xe7, 0x4c, 0xa8, 0x20,
	0x1e, 0xce, 0xd4, 0x74, 0x44, 0x69, 0x6e, 0xa4, 0x96, 0x91, 0x77, 0x86, 0x2b, 0x66, 0x70, 0x76,
	0x64, 0x70, 0x37, 0x33, 0xbc, 0x42, 0xfd, 0x31, 0x89, 0x56, 0x11, 0xc4, 0x1f, 0x26, 0x5a, 0x50,
	0xb6, 0x77, 0x28, 0x97, 0x37, 0x95, 0x97, 0x70, 0xb4, 0xae, 0x9c, 0x4f, 0xef, 0xff, 0xa5, 0x2f,
	0xde, 0x2d, 0xa8, 0x19, 0xd5, 0x7b, 0x4a, 0xe7, 0x2c, 0x24, 0xbc, 0x81, 0xbd, 0xc2, 0x43, 0xe2,
	0x71, 0x2f, 0xdb, 0xbd, 0x9f, 0xab, 0xd6, 0x6c, 0x6e, 0x6b, 0xe5, 0xce, 0x6f, 0xa1, 0x56, 0x4c,
	0x84, 0xab, 0x6f, 0xb7, 0x0c, 0xb8, 0x79, 0xb2, 0xb5, 0x97, 0xfd, 0xe8, 0xea, 0xec, 0xa9, 0x33,
	0x66, 0x72, 0xa2, 0x46, 0xbd, 0x90, 0x4f, 0x7d, 0x11, 0x48, 0x36, 0xe5, 0x92, 0xfb, 0x63, 0x7e,
	0xae, 0x6f, 0xc3, 0xcf, 0x0f, 0x64, 0xe4, 0x98, 0xdb, 0xb8, 0xfc, 0x0c, 0x00, 0x00, 0xff, 0xff,
	0x3f, 0xac, 0xba, 0x04, 0x32, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TokenServiceClient is the client API for TokenService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TokenServiceClient interface {
	CreateToken(ctx context.Context, in *CreateTokenRequest, opts ...grpc.CallOption) (*CreateTokenResponse, error)
	UpdateTokens(ctx context.Context, in *UpdateTokensRequest, opts ...grpc.CallOption) (*UpdateTokensResponse, error)
}

type tokenServiceClient struct {
	cc *grpc.ClientConn
}

func NewTokenServiceClient(cc *grpc.ClientConn) TokenServiceClient {
	return &tokenServiceClient{cc}
}

func (c *tokenServiceClient) CreateToken(ctx context.Context, in *CreateTokenRequest, opts ...grpc.CallOption) (*CreateTokenResponse, error) {
	out := new(CreateTokenResponse)
	err := c.cc.Invoke(ctx, "/token.TokenService/CreateToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tokenServiceClient) UpdateTokens(ctx context.Context, in *UpdateTokensRequest, opts ...grpc.CallOption) (*UpdateTokensResponse, error) {
	out := new(UpdateTokensResponse)
	err := c.cc.Invoke(ctx, "/token.TokenService/UpdateTokens", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TokenServiceServer is the server API for TokenService service.
type TokenServiceServer interface {
	CreateToken(context.Context, *CreateTokenRequest) (*CreateTokenResponse, error)
	UpdateTokens(context.Context, *UpdateTokensRequest) (*UpdateTokensResponse, error)
}

// UnimplementedTokenServiceServer can be embedded to have forward compatible implementations.
type UnimplementedTokenServiceServer struct {
}

func (*UnimplementedTokenServiceServer) CreateToken(ctx context.Context, req *CreateTokenRequest) (*CreateTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateToken not implemented")
}
func (*UnimplementedTokenServiceServer) UpdateTokens(ctx context.Context, req *UpdateTokensRequest) (*UpdateTokensResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTokens not implemented")
}

func RegisterTokenServiceServer(s *grpc.Server, srv TokenServiceServer) {
	s.RegisterService(&_TokenService_serviceDesc, srv)
}

func _TokenService_CreateToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TokenServiceServer).CreateToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/token.TokenService/CreateToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TokenServiceServer).CreateToken(ctx, req.(*CreateTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TokenService_UpdateTokens_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTokensRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TokenServiceServer).UpdateTokens(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/token.TokenService/UpdateTokens",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TokenServiceServer).UpdateTokens(ctx, req.(*UpdateTokensRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TokenService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "token.TokenService",
	HandlerType: (*TokenServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateToken",
			Handler:    _TokenService_CreateToken_Handler,
		},
		{
			MethodName: "UpdateTokens",
			Handler:    _TokenService_UpdateTokens_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ocpirpc/token.proto",
}
