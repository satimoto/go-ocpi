// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ocpirpc/cdr.proto

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

type CdrCreatedRequest struct {
	CdrUid               string   `protobuf:"bytes,1,opt,name=cdr_uid,json=cdrUid,proto3" json:"cdr_uid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CdrCreatedRequest) Reset()         { *m = CdrCreatedRequest{} }
func (m *CdrCreatedRequest) String() string { return proto.CompactTextString(m) }
func (*CdrCreatedRequest) ProtoMessage()    {}
func (*CdrCreatedRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6ae38dcf97755a6b, []int{0}
}

func (m *CdrCreatedRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CdrCreatedRequest.Unmarshal(m, b)
}
func (m *CdrCreatedRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CdrCreatedRequest.Marshal(b, m, deterministic)
}
func (m *CdrCreatedRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CdrCreatedRequest.Merge(m, src)
}
func (m *CdrCreatedRequest) XXX_Size() int {
	return xxx_messageInfo_CdrCreatedRequest.Size(m)
}
func (m *CdrCreatedRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CdrCreatedRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CdrCreatedRequest proto.InternalMessageInfo

func (m *CdrCreatedRequest) GetCdrUid() string {
	if m != nil {
		return m.CdrUid
	}
	return ""
}

type CdrCreatedResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CdrCreatedResponse) Reset()         { *m = CdrCreatedResponse{} }
func (m *CdrCreatedResponse) String() string { return proto.CompactTextString(m) }
func (*CdrCreatedResponse) ProtoMessage()    {}
func (*CdrCreatedResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6ae38dcf97755a6b, []int{1}
}

func (m *CdrCreatedResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CdrCreatedResponse.Unmarshal(m, b)
}
func (m *CdrCreatedResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CdrCreatedResponse.Marshal(b, m, deterministic)
}
func (m *CdrCreatedResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CdrCreatedResponse.Merge(m, src)
}
func (m *CdrCreatedResponse) XXX_Size() int {
	return xxx_messageInfo_CdrCreatedResponse.Size(m)
}
func (m *CdrCreatedResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CdrCreatedResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CdrCreatedResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*CdrCreatedRequest)(nil), "cdr.CdrCreatedRequest")
	proto.RegisterType((*CdrCreatedResponse)(nil), "cdr.CdrCreatedResponse")
}

func init() { proto.RegisterFile("ocpirpc/cdr.proto", fileDescriptor_6ae38dcf97755a6b) }

var fileDescriptor_6ae38dcf97755a6b = []byte{
	// 177 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xcc, 0x4f, 0x2e, 0xc8,
	0x2c, 0x2a, 0x48, 0xd6, 0x4f, 0x4e, 0x29, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4e,
	0x4e, 0x29, 0x52, 0xd2, 0xe1, 0x12, 0x74, 0x4e, 0x29, 0x72, 0x2e, 0x4a, 0x4d, 0x2c, 0x49, 0x4d,
	0x09, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0x12, 0xe7, 0x62, 0x4f, 0x4e, 0x29, 0x8a, 0x2f,
	0xcd, 0x4c, 0x91, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x62, 0x4b, 0x4e, 0x29, 0x0a, 0xcd, 0x4c,
	0x51, 0x12, 0xe1, 0x12, 0x42, 0x56, 0x5d, 0x5c, 0x90, 0x9f, 0x57, 0x9c, 0x6a, 0xe4, 0xcd, 0xc5,
	0xe5, 0x9c, 0x52, 0x14, 0x9c, 0x5a, 0x54, 0x96, 0x99, 0x9c, 0x2a, 0x64, 0x0b, 0xe6, 0x41, 0xd5,
	0x08, 0x89, 0xe9, 0x81, 0x2c, 0xc4, 0xb0, 0x42, 0x4a, 0x1c, 0x43, 0x1c, 0x62, 0x98, 0x93, 0x6a,
	0x94, 0x72, 0x7a, 0x66, 0x49, 0x46, 0x69, 0x92, 0x5e, 0x72, 0x7e, 0xae, 0x7e, 0x71, 0x62, 0x49,
	0x66, 0x6e, 0x7e, 0x49, 0xbe, 0x7e, 0x7a, 0xbe, 0x2e, 0xc8, 0x07, 0xfa, 0x50, 0x6f, 0x24, 0xb1,
	0x81, 0xfd, 0x60, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x75, 0xc7, 0x3b, 0x54, 0xd8, 0x00, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CdrServiceClient is the client API for CdrService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CdrServiceClient interface {
	CdrCreated(ctx context.Context, in *CdrCreatedRequest, opts ...grpc.CallOption) (*CdrCreatedResponse, error)
}

type cdrServiceClient struct {
	cc *grpc.ClientConn
}

func NewCdrServiceClient(cc *grpc.ClientConn) CdrServiceClient {
	return &cdrServiceClient{cc}
}

func (c *cdrServiceClient) CdrCreated(ctx context.Context, in *CdrCreatedRequest, opts ...grpc.CallOption) (*CdrCreatedResponse, error) {
	out := new(CdrCreatedResponse)
	err := c.cc.Invoke(ctx, "/cdr.CdrService/CdrCreated", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CdrServiceServer is the server API for CdrService service.
type CdrServiceServer interface {
	CdrCreated(context.Context, *CdrCreatedRequest) (*CdrCreatedResponse, error)
}

// UnimplementedCdrServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCdrServiceServer struct {
}

func (*UnimplementedCdrServiceServer) CdrCreated(ctx context.Context, req *CdrCreatedRequest) (*CdrCreatedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CdrCreated not implemented")
}

func RegisterCdrServiceServer(s *grpc.Server, srv CdrServiceServer) {
	s.RegisterService(&_CdrService_serviceDesc, srv)
}

func _CdrService_CdrCreated_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CdrCreatedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CdrServiceServer).CdrCreated(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cdr.CdrService/CdrCreated",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CdrServiceServer).CdrCreated(ctx, req.(*CdrCreatedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CdrService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cdr.CdrService",
	HandlerType: (*CdrServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CdrCreated",
			Handler:    _CdrService_CdrCreated_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ocpirpc/cdr.proto",
}
