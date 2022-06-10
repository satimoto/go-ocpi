// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ocpirpc/rpc.proto

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

type TestConnectionRequest struct {
	Addr                 string   `protobuf:"bytes,1,opt,name=addr,proto3" json:"addr,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TestConnectionRequest) Reset()         { *m = TestConnectionRequest{} }
func (m *TestConnectionRequest) String() string { return proto.CompactTextString(m) }
func (*TestConnectionRequest) ProtoMessage()    {}
func (*TestConnectionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_98b68261e5047fb9, []int{0}
}

func (m *TestConnectionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TestConnectionRequest.Unmarshal(m, b)
}
func (m *TestConnectionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TestConnectionRequest.Marshal(b, m, deterministic)
}
func (m *TestConnectionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestConnectionRequest.Merge(m, src)
}
func (m *TestConnectionRequest) XXX_Size() int {
	return xxx_messageInfo_TestConnectionRequest.Size(m)
}
func (m *TestConnectionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TestConnectionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TestConnectionRequest proto.InternalMessageInfo

func (m *TestConnectionRequest) GetAddr() string {
	if m != nil {
		return m.Addr
	}
	return ""
}

type TestConnectionResponse struct {
	Result               string   `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TestConnectionResponse) Reset()         { *m = TestConnectionResponse{} }
func (m *TestConnectionResponse) String() string { return proto.CompactTextString(m) }
func (*TestConnectionResponse) ProtoMessage()    {}
func (*TestConnectionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_98b68261e5047fb9, []int{1}
}

func (m *TestConnectionResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TestConnectionResponse.Unmarshal(m, b)
}
func (m *TestConnectionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TestConnectionResponse.Marshal(b, m, deterministic)
}
func (m *TestConnectionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestConnectionResponse.Merge(m, src)
}
func (m *TestConnectionResponse) XXX_Size() int {
	return xxx_messageInfo_TestConnectionResponse.Size(m)
}
func (m *TestConnectionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TestConnectionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TestConnectionResponse proto.InternalMessageInfo

func (m *TestConnectionResponse) GetResult() string {
	if m != nil {
		return m.Result
	}
	return ""
}

type TestMessageRequest struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TestMessageRequest) Reset()         { *m = TestMessageRequest{} }
func (m *TestMessageRequest) String() string { return proto.CompactTextString(m) }
func (*TestMessageRequest) ProtoMessage()    {}
func (*TestMessageRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_98b68261e5047fb9, []int{2}
}

func (m *TestMessageRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TestMessageRequest.Unmarshal(m, b)
}
func (m *TestMessageRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TestMessageRequest.Marshal(b, m, deterministic)
}
func (m *TestMessageRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestMessageRequest.Merge(m, src)
}
func (m *TestMessageRequest) XXX_Size() int {
	return xxx_messageInfo_TestMessageRequest.Size(m)
}
func (m *TestMessageRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TestMessageRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TestMessageRequest proto.InternalMessageInfo

func (m *TestMessageRequest) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type TestMessageResponse struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TestMessageResponse) Reset()         { *m = TestMessageResponse{} }
func (m *TestMessageResponse) String() string { return proto.CompactTextString(m) }
func (*TestMessageResponse) ProtoMessage()    {}
func (*TestMessageResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_98b68261e5047fb9, []int{3}
}

func (m *TestMessageResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TestMessageResponse.Unmarshal(m, b)
}
func (m *TestMessageResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TestMessageResponse.Marshal(b, m, deterministic)
}
func (m *TestMessageResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestMessageResponse.Merge(m, src)
}
func (m *TestMessageResponse) XXX_Size() int {
	return xxx_messageInfo_TestMessageResponse.Size(m)
}
func (m *TestMessageResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TestMessageResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TestMessageResponse proto.InternalMessageInfo

func (m *TestMessageResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*TestConnectionRequest)(nil), "ping.TestConnectionRequest")
	proto.RegisterType((*TestConnectionResponse)(nil), "ping.TestConnectionResponse")
	proto.RegisterType((*TestMessageRequest)(nil), "ping.TestMessageRequest")
	proto.RegisterType((*TestMessageResponse)(nil), "ping.TestMessageResponse")
}

func init() { proto.RegisterFile("ocpirpc/rpc.proto", fileDescriptor_98b68261e5047fb9) }

var fileDescriptor_98b68261e5047fb9 = []byte{
	// 243 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0xc1, 0x4a, 0x03, 0x31,
	0x10, 0x86, 0x59, 0x28, 0x15, 0x47, 0x10, 0x8c, 0x58, 0xd6, 0xea, 0x41, 0xf6, 0xa2, 0x22, 0x4d,
	0x44, 0xdf, 0xa0, 0x1e, 0xc5, 0xcb, 0xea, 0xc9, 0xdb, 0x36, 0x1d, 0x62, 0xc0, 0xcd, 0xc4, 0xcc,
	0xac, 0x6f, 0xe3, 0xbb, 0xca, 0xb6, 0x29, 0xba, 0xb2, 0xde, 0x32, 0x99, 0x6f, 0xfe, 0x99, 0x9f,
	0x1f, 0x8e, 0xc8, 0x46, 0x9f, 0xa2, 0x35, 0x29, 0x5a, 0x1d, 0x13, 0x09, 0xa9, 0x49, 0xf4, 0xc1,
	0x55, 0x37, 0x70, 0xf2, 0x82, 0x2c, 0x0f, 0x14, 0x02, 0x5a, 0xf1, 0x14, 0x6a, 0xfc, 0xe8, 0x90,
	0x45, 0x29, 0x98, 0x34, 0xeb, 0x75, 0x2a, 0x8b, 0x8b, 0xe2, 0x6a, 0xbf, 0xde, 0xbc, 0xab, 0x5b,
	0x98, 0xfd, 0x85, 0x39, 0x52, 0x60, 0x54, 0x33, 0x98, 0x26, 0xe4, 0xee, 0x5d, 0x32, 0x9f, 0xab,
	0x4a, 0x83, 0xea, 0x27, 0x9e, 0x90, 0xb9, 0x71, 0xb8, 0xd3, 0x2e, 0x61, 0xaf, 0xdd, 0xfe, 0x64,
	0x7c, 0x57, 0x56, 0x06, 0x8e, 0x07, 0x7c, 0x96, 0xff, 0x77, 0xe0, 0xee, 0xab, 0x00, 0xa8, 0xa3,
	0x7d, 0xc6, 0xf4, 0xe9, 0x2d, 0xaa, 0x47, 0x38, 0x1c, 0x5e, 0xa8, 0xce, 0x74, 0xef, 0x53, 0x8f,
	0x9a, 0x9c, 0x9f, 0x8f, 0x37, 0xf3, 0xd6, 0x25, 0x1c, 0xfc, 0x3a, 0x46, 0x95, 0x3f, 0xf0, 0xd0,
	0xcf, 0xfc, 0x74, 0xa4, 0xb3, 0xd5, 0x58, 0x5e, 0xbf, 0x5e, 0x3a, 0x2f, 0x6f, 0xdd, 0x4a, 0x5b,
	0x6a, 0x0d, 0x37, 0xe2, 0x5b, 0x12, 0x32, 0x8e, 0x16, 0x7d, 0x22, 0x8b, 0x26, 0x7a, 0x93, 0xa3,
	0x59, 0x4d, 0x37, 0xb9, 0xdc, 0x7f, 0x07, 0x00, 0x00, 0xff, 0xff, 0x13, 0xd3, 0xca, 0xcd, 0xac,
	0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RpcServiceClient is the client API for RpcService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RpcServiceClient interface {
	TestConnection(ctx context.Context, in *TestConnectionRequest, opts ...grpc.CallOption) (*TestConnectionResponse, error)
	TestMessage(ctx context.Context, in *TestMessageRequest, opts ...grpc.CallOption) (*TestMessageResponse, error)
}

type rpcServiceClient struct {
	cc *grpc.ClientConn
}

func NewRpcServiceClient(cc *grpc.ClientConn) RpcServiceClient {
	return &rpcServiceClient{cc}
}

func (c *rpcServiceClient) TestConnection(ctx context.Context, in *TestConnectionRequest, opts ...grpc.CallOption) (*TestConnectionResponse, error) {
	out := new(TestConnectionResponse)
	err := c.cc.Invoke(ctx, "/ping.RpcService/TestConnection", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rpcServiceClient) TestMessage(ctx context.Context, in *TestMessageRequest, opts ...grpc.CallOption) (*TestMessageResponse, error) {
	out := new(TestMessageResponse)
	err := c.cc.Invoke(ctx, "/ping.RpcService/TestMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RpcServiceServer is the server API for RpcService service.
type RpcServiceServer interface {
	TestConnection(context.Context, *TestConnectionRequest) (*TestConnectionResponse, error)
	TestMessage(context.Context, *TestMessageRequest) (*TestMessageResponse, error)
}

// UnimplementedRpcServiceServer can be embedded to have forward compatible implementations.
type UnimplementedRpcServiceServer struct {
}

func (*UnimplementedRpcServiceServer) TestConnection(ctx context.Context, req *TestConnectionRequest) (*TestConnectionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TestConnection not implemented")
}
func (*UnimplementedRpcServiceServer) TestMessage(ctx context.Context, req *TestMessageRequest) (*TestMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TestMessage not implemented")
}

func RegisterRpcServiceServer(s *grpc.Server, srv RpcServiceServer) {
	s.RegisterService(&_RpcService_serviceDesc, srv)
}

func _RpcService_TestConnection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TestConnectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpcServiceServer).TestConnection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ping.RpcService/TestConnection",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpcServiceServer).TestConnection(ctx, req.(*TestConnectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RpcService_TestMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TestMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpcServiceServer).TestMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ping.RpcService/TestMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpcServiceServer).TestMessage(ctx, req.(*TestMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RpcService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ping.RpcService",
	HandlerType: (*RpcServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "TestConnection",
			Handler:    _RpcService_TestConnection_Handler,
		},
		{
			MethodName: "TestMessage",
			Handler:    _RpcService_TestMessage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ocpirpc/rpc.proto",
}
