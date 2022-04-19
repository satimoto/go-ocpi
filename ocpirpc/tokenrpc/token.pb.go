// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/token.proto

package tokenrpc

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
	Type                 string   `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	Allowed              string   `protobuf:"bytes,3,opt,name=allowed,proto3" json:"allowed,omitempty"`
	Whitelist            string   `protobuf:"bytes,4,opt,name=whitelist,proto3" json:"whitelist,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateTokenRequest) Reset()         { *m = CreateTokenRequest{} }
func (m *CreateTokenRequest) String() string { return proto.CompactTextString(m) }
func (*CreateTokenRequest) ProtoMessage()    {}
func (*CreateTokenRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0498de655d1970f6, []int{0}
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

type TokenResponse struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Type                 string   `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	AuthId               string   `protobuf:"bytes,3,opt,name=auth_id,json=authId,proto3" json:"auth_id,omitempty"`
	Allowed              string   `protobuf:"bytes,4,opt,name=allowed,proto3" json:"allowed,omitempty"`
	Whitelist            string   `protobuf:"bytes,5,opt,name=whitelist,proto3" json:"whitelist,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TokenResponse) Reset()         { *m = TokenResponse{} }
func (m *TokenResponse) String() string { return proto.CompactTextString(m) }
func (*TokenResponse) ProtoMessage()    {}
func (*TokenResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0498de655d1970f6, []int{1}
}

func (m *TokenResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TokenResponse.Unmarshal(m, b)
}
func (m *TokenResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TokenResponse.Marshal(b, m, deterministic)
}
func (m *TokenResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TokenResponse.Merge(m, src)
}
func (m *TokenResponse) XXX_Size() int {
	return xxx_messageInfo_TokenResponse.Size(m)
}
func (m *TokenResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TokenResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TokenResponse proto.InternalMessageInfo

func (m *TokenResponse) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *TokenResponse) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *TokenResponse) GetAuthId() string {
	if m != nil {
		return m.AuthId
	}
	return ""
}

func (m *TokenResponse) GetAllowed() string {
	if m != nil {
		return m.Allowed
	}
	return ""
}

func (m *TokenResponse) GetWhitelist() string {
	if m != nil {
		return m.Whitelist
	}
	return ""
}

func init() {
	proto.RegisterType((*CreateTokenRequest)(nil), "token.CreateTokenRequest")
	proto.RegisterType((*TokenResponse)(nil), "token.TokenResponse")
}

func init() { proto.RegisterFile("proto/token.proto", fileDescriptor_0498de655d1970f6) }

var fileDescriptor_0498de655d1970f6 = []byte{
	// 269 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0x3f, 0x4f, 0xc3, 0x30,
	0x14, 0xc4, 0x95, 0x34, 0x4d, 0xd4, 0x87, 0x40, 0xaa, 0x85, 0xd4, 0x80, 0x18, 0xaa, 0x4e, 0x5d,
	0x9a, 0xa0, 0xb2, 0x33, 0xd0, 0xa9, 0x6b, 0xca, 0xc4, 0x82, 0xdc, 0xf8, 0xa9, 0xb1, 0x48, 0x63,
	0x63, 0xbf, 0x50, 0x75, 0xe7, 0x83, 0x23, 0x3b, 0x05, 0xca, 0x1f, 0x75, 0xbb, 0x9f, 0xdf, 0x70,
	0xe7, 0x3b, 0x18, 0x6a, 0xa3, 0x48, 0xe5, 0xa4, 0x5e, 0xb0, 0xc9, 0xbc, 0x66, 0x7d, 0x0f, 0x93,
	0x3d, 0xb0, 0x85, 0x41, 0x4e, 0xf8, 0xe8, 0xb0, 0xc0, 0xd7, 0x16, 0x2d, 0xb1, 0x11, 0x24, 0xad,
	0x45, 0xf3, 0x2c, 0x45, 0x1a, 0x8c, 0x83, 0x69, 0xaf, 0x88, 0x1d, 0x2e, 0x05, 0x63, 0x10, 0xd1,
	0x5e, 0x63, 0x1a, 0x8e, 0x83, 0xe9, 0xa0, 0xf0, 0x9a, 0xa5, 0x90, 0xf0, 0xba, 0x56, 0x3b, 0x14,
	0x69, 0xcf, 0x3f, 0x7f, 0x22, 0xbb, 0x81, 0xc1, 0xae, 0x92, 0x84, 0xb5, 0xb4, 0x94, 0x46, 0xfe,
	0xf6, 0xfd, 0x30, 0x79, 0x0f, 0xe0, 0xfc, 0xe0, 0x6a, 0xb5, 0x6a, 0x2c, 0xb2, 0x0b, 0x08, 0xbf,
	0x1c, 0x43, 0xf9, 0xbf, 0xdb, 0x08, 0x12, 0xde, 0x52, 0xe5, 0xa2, 0x75, 0x6e, 0xb1, 0xc3, 0xa5,
	0x38, 0x8e, 0x11, 0x9d, 0x88, 0xd1, 0xff, 0x15, 0x63, 0xbe, 0x82, 0xe1, 0xc2, 0xa0, 0xc0, 0x86,
	0x24, 0xaf, 0x57, 0x68, 0xde, 0x64, 0x89, 0xec, 0x1e, 0xce, 0x8e, 0x6a, 0x61, 0x57, 0x59, 0x57,
	0xdd, 0xdf, 0xaa, 0xae, 0x2f, 0x0f, 0xa7, 0x1f, 0x3f, 0x79, 0x98, 0x3f, 0xdd, 0x6e, 0x24, 0x55,
	0xed, 0x3a, 0x2b, 0xd5, 0x36, 0xb7, 0x9c, 0xe4, 0xd6, 0x0d, 0xb0, 0x51, 0x33, 0x55, 0x6a, 0x39,
	0xe3, 0x5a, 0xe6, 0x4e, 0x18, 0x5d, 0x76, 0xa3, 0x18, 0x5d, 0xae, 0x63, 0x3f, 0xcc, 0xdd, 0x47,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x47, 0xc7, 0x8a, 0x4b, 0xad, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CredentialServiceClient is the client API for CredentialService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CredentialServiceClient interface {
	CreateToken(ctx context.Context, in *CreateTokenRequest, opts ...grpc.CallOption) (*TokenResponse, error)
}

type credentialServiceClient struct {
	cc *grpc.ClientConn
}

func NewCredentialServiceClient(cc *grpc.ClientConn) CredentialServiceClient {
	return &credentialServiceClient{cc}
}

func (c *credentialServiceClient) CreateToken(ctx context.Context, in *CreateTokenRequest, opts ...grpc.CallOption) (*TokenResponse, error) {
	out := new(TokenResponse)
	err := c.cc.Invoke(ctx, "/token.CredentialService/CreateToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CredentialServiceServer is the server API for CredentialService service.
type CredentialServiceServer interface {
	CreateToken(context.Context, *CreateTokenRequest) (*TokenResponse, error)
}

// UnimplementedCredentialServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCredentialServiceServer struct {
}

func (*UnimplementedCredentialServiceServer) CreateToken(ctx context.Context, req *CreateTokenRequest) (*TokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateToken not implemented")
}

func RegisterCredentialServiceServer(s *grpc.Server, srv CredentialServiceServer) {
	s.RegisterService(&_CredentialService_serviceDesc, srv)
}

func _CredentialService_CreateToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CredentialServiceServer).CreateToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/token.CredentialService/CreateToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CredentialServiceServer).CreateToken(ctx, req.(*CreateTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CredentialService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "token.CredentialService",
	HandlerType: (*CredentialServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateToken",
			Handler:    _CredentialService_CreateToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/token.proto",
}
