// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/credential.proto

package credentialrpc

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	ocpirpc "github.com/satimoto/go-ocpi-api/ocpirpc"
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

type CreateCredentialRequest struct {
	ClientToken          string                               `protobuf:"bytes,1,opt,name=client_token,json=clientToken,proto3" json:"client_token,omitempty"`
	Url                  string                               `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	BusinessDetail       *ocpirpc.CreateBusinessDetailRequest `protobuf:"bytes,3,opt,name=business_detail,json=businessDetail,proto3" json:"business_detail,omitempty"`
	CountryCode          string                               `protobuf:"bytes,4,opt,name=country_code,json=countryCode,proto3" json:"country_code,omitempty"`
	PartyId              string                               `protobuf:"bytes,5,opt,name=party_id,json=partyId,proto3" json:"party_id,omitempty"`
	IsHub                bool                                 `protobuf:"varint,6,opt,name=is_hub,json=isHub,proto3" json:"is_hub,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                             `json:"-"`
	XXX_unrecognized     []byte                               `json:"-"`
	XXX_sizecache        int32                                `json:"-"`
}

func (m *CreateCredentialRequest) Reset()         { *m = CreateCredentialRequest{} }
func (m *CreateCredentialRequest) String() string { return proto.CompactTextString(m) }
func (*CreateCredentialRequest) ProtoMessage()    {}
func (*CreateCredentialRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1dd6b9da2fa6db73, []int{0}
}

func (m *CreateCredentialRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateCredentialRequest.Unmarshal(m, b)
}
func (m *CreateCredentialRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateCredentialRequest.Marshal(b, m, deterministic)
}
func (m *CreateCredentialRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateCredentialRequest.Merge(m, src)
}
func (m *CreateCredentialRequest) XXX_Size() int {
	return xxx_messageInfo_CreateCredentialRequest.Size(m)
}
func (m *CreateCredentialRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateCredentialRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateCredentialRequest proto.InternalMessageInfo

func (m *CreateCredentialRequest) GetClientToken() string {
	if m != nil {
		return m.ClientToken
	}
	return ""
}

func (m *CreateCredentialRequest) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *CreateCredentialRequest) GetBusinessDetail() *ocpirpc.CreateBusinessDetailRequest {
	if m != nil {
		return m.BusinessDetail
	}
	return nil
}

func (m *CreateCredentialRequest) GetCountryCode() string {
	if m != nil {
		return m.CountryCode
	}
	return ""
}

func (m *CreateCredentialRequest) GetPartyId() string {
	if m != nil {
		return m.PartyId
	}
	return ""
}

func (m *CreateCredentialRequest) GetIsHub() bool {
	if m != nil {
		return m.IsHub
	}
	return false
}

type CredentialResponse struct {
	Id                   int64                           `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	ClientToken          string                          `protobuf:"bytes,2,opt,name=client_token,json=clientToken,proto3" json:"client_token,omitempty"`
	Url                  string                          `protobuf:"bytes,3,opt,name=url,proto3" json:"url,omitempty"`
	BusinessDetail       *ocpirpc.BusinessDetailResponse `protobuf:"bytes,4,opt,name=business_detail,json=businessDetail,proto3" json:"business_detail,omitempty"`
	CountryCode          string                          `protobuf:"bytes,5,opt,name=country_code,json=countryCode,proto3" json:"country_code,omitempty"`
	PartyId              string                          `protobuf:"bytes,6,opt,name=party_id,json=partyId,proto3" json:"party_id,omitempty"`
	IsHub                bool                            `protobuf:"varint,7,opt,name=is_hub,json=isHub,proto3" json:"is_hub,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *CredentialResponse) Reset()         { *m = CredentialResponse{} }
func (m *CredentialResponse) String() string { return proto.CompactTextString(m) }
func (*CredentialResponse) ProtoMessage()    {}
func (*CredentialResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1dd6b9da2fa6db73, []int{1}
}

func (m *CredentialResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CredentialResponse.Unmarshal(m, b)
}
func (m *CredentialResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CredentialResponse.Marshal(b, m, deterministic)
}
func (m *CredentialResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CredentialResponse.Merge(m, src)
}
func (m *CredentialResponse) XXX_Size() int {
	return xxx_messageInfo_CredentialResponse.Size(m)
}
func (m *CredentialResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CredentialResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CredentialResponse proto.InternalMessageInfo

func (m *CredentialResponse) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *CredentialResponse) GetClientToken() string {
	if m != nil {
		return m.ClientToken
	}
	return ""
}

func (m *CredentialResponse) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *CredentialResponse) GetBusinessDetail() *ocpirpc.BusinessDetailResponse {
	if m != nil {
		return m.BusinessDetail
	}
	return nil
}

func (m *CredentialResponse) GetCountryCode() string {
	if m != nil {
		return m.CountryCode
	}
	return ""
}

func (m *CredentialResponse) GetPartyId() string {
	if m != nil {
		return m.PartyId
	}
	return ""
}

func (m *CredentialResponse) GetIsHub() bool {
	if m != nil {
		return m.IsHub
	}
	return false
}

func init() {
	proto.RegisterType((*CreateCredentialRequest)(nil), "crediential.CreateCredentialRequest")
	proto.RegisterType((*CredentialResponse)(nil), "crediential.CredentialResponse")
}

func init() { proto.RegisterFile("proto/credential.proto", fileDescriptor_1dd6b9da2fa6db73) }

var fileDescriptor_1dd6b9da2fa6db73 = []byte{
	// 361 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0xc1, 0x4f, 0xc2, 0x30,
	0x14, 0x87, 0xb3, 0x0d, 0x06, 0x16, 0x83, 0xd8, 0x44, 0x9d, 0x5c, 0x44, 0x62, 0x0c, 0x89, 0x61,
	0x24, 0x18, 0xe3, 0x1d, 0x3c, 0xe8, 0xc9, 0x64, 0x72, 0xd1, 0xcb, 0xb2, 0xb5, 0x0d, 0xbc, 0x08,
	0xeb, 0x6c, 0x3b, 0x13, 0xfe, 0x74, 0x4f, 0x9a, 0xb5, 0x53, 0xe6, 0x08, 0x7a, 0x6b, 0xbf, 0xbe,
	0xf4, 0xbd, 0xf7, 0xe5, 0x87, 0x8e, 0x53, 0xc1, 0x15, 0x1f, 0x11, 0xc1, 0x28, 0x4b, 0x14, 0x44,
	0x4b, 0x5f, 0x03, 0xdc, 0xca, 0x09, 0x18, 0xd4, 0xed, 0x9a, 0xa2, 0x38, 0x93, 0x90, 0x30, 0x29,
	0x29, 0x53, 0x11, 0x14, 0x85, 0xfd, 0x0f, 0x0b, 0x9d, 0x4c, 0x05, 0x8b, 0x14, 0x9b, 0xfe, 0xfc,
	0x11, 0xb0, 0xb7, 0x8c, 0x49, 0x85, 0xcf, 0xd1, 0x3e, 0x59, 0xe6, 0x9f, 0x84, 0x8a, 0xbf, 0xb2,
	0xc4, 0xb3, 0x7a, 0xd6, 0x60, 0x2f, 0x68, 0x19, 0x36, 0xcb, 0x11, 0xee, 0x20, 0x27, 0x13, 0x4b,
	0xcf, 0xd6, 0x2f, 0xf9, 0x11, 0xcf, 0xd0, 0xc1, 0x77, 0xa3, 0xd0, 0x74, 0xf2, 0x9c, 0x9e, 0x35,
	0x68, 0x8d, 0xaf, 0xfc, 0xca, 0x00, 0xa6, 0xed, 0xa4, 0x80, 0x77, 0x1a, 0x16, 0xad, 0x83, 0x76,
	0xfc, 0x0b, 0xeb, 0x51, 0x78, 0x96, 0x28, 0xb1, 0x0e, 0x09, 0xa7, 0xcc, 0xab, 0x15, 0xa3, 0x18,
	0x36, 0xe5, 0x94, 0xe1, 0x53, 0xd4, 0x4c, 0x23, 0xa1, 0xd6, 0x21, 0x50, 0xaf, 0xae, 0x9f, 0x1b,
	0xfa, 0xfe, 0x40, 0xf1, 0x11, 0x72, 0x41, 0x86, 0x8b, 0x2c, 0xf6, 0xdc, 0x9e, 0x35, 0x68, 0x06,
	0x75, 0x90, 0xf7, 0x59, 0xdc, 0xff, 0xb4, 0x10, 0x2e, 0x6f, 0x2d, 0x53, 0x9e, 0x48, 0x86, 0xdb,
	0xc8, 0x06, 0xaa, 0x97, 0x75, 0x02, 0x1b, 0xe8, 0x96, 0x06, 0x7b, 0xa7, 0x06, 0x67, 0xa3, 0xe1,
	0x71, 0x5b, 0x43, 0x4d, 0x6b, 0xb8, 0xac, 0x6a, 0xa8, 0x0a, 0x30, 0x53, 0xfc, 0x6b, 0xa0, 0xfe,
	0xb7, 0x01, 0x77, 0x97, 0x81, 0x46, 0xc9, 0xc0, 0x38, 0x41, 0x87, 0x1b, 0x01, 0x4f, 0x4c, 0xbc,
	0x03, 0x61, 0xf8, 0x19, 0x75, 0xaa, 0x89, 0xc0, 0x17, 0x7e, 0x29, 0x50, 0xfe, 0x8e, 0xc0, 0x74,
	0xcf, 0xaa, 0x55, 0x15, 0xb5, 0x93, 0xdb, 0x97, 0x9b, 0x39, 0xa8, 0x45, 0x16, 0xfb, 0x84, 0xaf,
	0x46, 0x32, 0x52, 0xb0, 0xca, 0x93, 0x39, 0xe7, 0x43, 0x4e, 0x52, 0x18, 0x46, 0x29, 0x8c, 0xf2,
	0x83, 0x48, 0x49, 0x29, 0xd2, 0x22, 0x25, 0xb1, 0xab, 0xd3, 0x7a, 0xfd, 0x15, 0x00, 0x00, 0xff,
	0xff, 0xf3, 0x9c, 0x1b, 0xe5, 0xf0, 0x02, 0x00, 0x00,
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
	CreateCredential(ctx context.Context, in *CreateCredentialRequest, opts ...grpc.CallOption) (*CredentialResponse, error)
}

type credentialServiceClient struct {
	cc *grpc.ClientConn
}

func NewCredentialServiceClient(cc *grpc.ClientConn) CredentialServiceClient {
	return &credentialServiceClient{cc}
}

func (c *credentialServiceClient) CreateCredential(ctx context.Context, in *CreateCredentialRequest, opts ...grpc.CallOption) (*CredentialResponse, error) {
	out := new(CredentialResponse)
	err := c.cc.Invoke(ctx, "/crediential.CredentialService/CreateCredential", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CredentialServiceServer is the server API for CredentialService service.
type CredentialServiceServer interface {
	CreateCredential(context.Context, *CreateCredentialRequest) (*CredentialResponse, error)
}

// UnimplementedCredentialServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCredentialServiceServer struct {
}

func (*UnimplementedCredentialServiceServer) CreateCredential(ctx context.Context, req *CreateCredentialRequest) (*CredentialResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCredential not implemented")
}

func RegisterCredentialServiceServer(s *grpc.Server, srv CredentialServiceServer) {
	s.RegisterService(&_CredentialService_serviceDesc, srv)
}

func _CredentialService_CreateCredential_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCredentialRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CredentialServiceServer).CreateCredential(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/crediential.CredentialService/CreateCredential",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CredentialServiceServer).CreateCredential(ctx, req.(*CreateCredentialRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CredentialService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "crediential.CredentialService",
	HandlerType: (*CredentialServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCredential",
			Handler:    _CredentialService_CreateCredential_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/credential.proto",
}
