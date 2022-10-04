// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ocpirpc/credential.proto

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

type CreateCredentialRequest struct {
	ClientToken          string                       `protobuf:"bytes,1,opt,name=client_token,json=clientToken,proto3" json:"client_token,omitempty"`
	Url                  string                       `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	BusinessDetail       *CreateBusinessDetailRequest `protobuf:"bytes,3,opt,name=business_detail,json=businessDetail,proto3" json:"business_detail,omitempty"`
	CountryCode          string                       `protobuf:"bytes,4,opt,name=country_code,json=countryCode,proto3" json:"country_code,omitempty"`
	PartyId              string                       `protobuf:"bytes,5,opt,name=party_id,json=partyId,proto3" json:"party_id,omitempty"`
	IsHub                bool                         `protobuf:"varint,6,opt,name=is_hub,json=isHub,proto3" json:"is_hub,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *CreateCredentialRequest) Reset()         { *m = CreateCredentialRequest{} }
func (m *CreateCredentialRequest) String() string { return proto.CompactTextString(m) }
func (*CreateCredentialRequest) ProtoMessage()    {}
func (*CreateCredentialRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_19dfbe8a464b7ce7, []int{0}
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

func (m *CreateCredentialRequest) GetBusinessDetail() *CreateBusinessDetailRequest {
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

type CreateCredentialResponse struct {
	Id                   int64                   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	ClientToken          string                  `protobuf:"bytes,2,opt,name=client_token,json=clientToken,proto3" json:"client_token,omitempty"`
	Url                  string                  `protobuf:"bytes,3,opt,name=url,proto3" json:"url,omitempty"`
	BusinessDetail       *BusinessDetailResponse `protobuf:"bytes,4,opt,name=business_detail,json=businessDetail,proto3" json:"business_detail,omitempty"`
	CountryCode          string                  `protobuf:"bytes,5,opt,name=country_code,json=countryCode,proto3" json:"country_code,omitempty"`
	PartyId              string                  `protobuf:"bytes,6,opt,name=party_id,json=partyId,proto3" json:"party_id,omitempty"`
	IsHub                bool                    `protobuf:"varint,7,opt,name=is_hub,json=isHub,proto3" json:"is_hub,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *CreateCredentialResponse) Reset()         { *m = CreateCredentialResponse{} }
func (m *CreateCredentialResponse) String() string { return proto.CompactTextString(m) }
func (*CreateCredentialResponse) ProtoMessage()    {}
func (*CreateCredentialResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_19dfbe8a464b7ce7, []int{1}
}

func (m *CreateCredentialResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateCredentialResponse.Unmarshal(m, b)
}
func (m *CreateCredentialResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateCredentialResponse.Marshal(b, m, deterministic)
}
func (m *CreateCredentialResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateCredentialResponse.Merge(m, src)
}
func (m *CreateCredentialResponse) XXX_Size() int {
	return xxx_messageInfo_CreateCredentialResponse.Size(m)
}
func (m *CreateCredentialResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateCredentialResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateCredentialResponse proto.InternalMessageInfo

func (m *CreateCredentialResponse) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *CreateCredentialResponse) GetClientToken() string {
	if m != nil {
		return m.ClientToken
	}
	return ""
}

func (m *CreateCredentialResponse) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *CreateCredentialResponse) GetBusinessDetail() *BusinessDetailResponse {
	if m != nil {
		return m.BusinessDetail
	}
	return nil
}

func (m *CreateCredentialResponse) GetCountryCode() string {
	if m != nil {
		return m.CountryCode
	}
	return ""
}

func (m *CreateCredentialResponse) GetPartyId() string {
	if m != nil {
		return m.PartyId
	}
	return ""
}

func (m *CreateCredentialResponse) GetIsHub() bool {
	if m != nil {
		return m.IsHub
	}
	return false
}

type RegisterCredentialRequest struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	ClientToken          string   `protobuf:"bytes,2,opt,name=client_token,json=clientToken,proto3" json:"client_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterCredentialRequest) Reset()         { *m = RegisterCredentialRequest{} }
func (m *RegisterCredentialRequest) String() string { return proto.CompactTextString(m) }
func (*RegisterCredentialRequest) ProtoMessage()    {}
func (*RegisterCredentialRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_19dfbe8a464b7ce7, []int{2}
}

func (m *RegisterCredentialRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterCredentialRequest.Unmarshal(m, b)
}
func (m *RegisterCredentialRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterCredentialRequest.Marshal(b, m, deterministic)
}
func (m *RegisterCredentialRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterCredentialRequest.Merge(m, src)
}
func (m *RegisterCredentialRequest) XXX_Size() int {
	return xxx_messageInfo_RegisterCredentialRequest.Size(m)
}
func (m *RegisterCredentialRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterCredentialRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterCredentialRequest proto.InternalMessageInfo

func (m *RegisterCredentialRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *RegisterCredentialRequest) GetClientToken() string {
	if m != nil {
		return m.ClientToken
	}
	return ""
}

type RegisterCredentialResponse struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterCredentialResponse) Reset()         { *m = RegisterCredentialResponse{} }
func (m *RegisterCredentialResponse) String() string { return proto.CompactTextString(m) }
func (*RegisterCredentialResponse) ProtoMessage()    {}
func (*RegisterCredentialResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_19dfbe8a464b7ce7, []int{3}
}

func (m *RegisterCredentialResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterCredentialResponse.Unmarshal(m, b)
}
func (m *RegisterCredentialResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterCredentialResponse.Marshal(b, m, deterministic)
}
func (m *RegisterCredentialResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterCredentialResponse.Merge(m, src)
}
func (m *RegisterCredentialResponse) XXX_Size() int {
	return xxx_messageInfo_RegisterCredentialResponse.Size(m)
}
func (m *RegisterCredentialResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterCredentialResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterCredentialResponse proto.InternalMessageInfo

func (m *RegisterCredentialResponse) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type SyncCredentialRequest struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	FromDate             string   `protobuf:"bytes,2,opt,name=from_date,json=fromDate,proto3" json:"from_date,omitempty"`
	CountryCode          string   `protobuf:"bytes,3,opt,name=country_code,json=countryCode,proto3" json:"country_code,omitempty"`
	PartyId              string   `protobuf:"bytes,4,opt,name=party_id,json=partyId,proto3" json:"party_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SyncCredentialRequest) Reset()         { *m = SyncCredentialRequest{} }
func (m *SyncCredentialRequest) String() string { return proto.CompactTextString(m) }
func (*SyncCredentialRequest) ProtoMessage()    {}
func (*SyncCredentialRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_19dfbe8a464b7ce7, []int{4}
}

func (m *SyncCredentialRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SyncCredentialRequest.Unmarshal(m, b)
}
func (m *SyncCredentialRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SyncCredentialRequest.Marshal(b, m, deterministic)
}
func (m *SyncCredentialRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SyncCredentialRequest.Merge(m, src)
}
func (m *SyncCredentialRequest) XXX_Size() int {
	return xxx_messageInfo_SyncCredentialRequest.Size(m)
}
func (m *SyncCredentialRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SyncCredentialRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SyncCredentialRequest proto.InternalMessageInfo

func (m *SyncCredentialRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *SyncCredentialRequest) GetFromDate() string {
	if m != nil {
		return m.FromDate
	}
	return ""
}

func (m *SyncCredentialRequest) GetCountryCode() string {
	if m != nil {
		return m.CountryCode
	}
	return ""
}

func (m *SyncCredentialRequest) GetPartyId() string {
	if m != nil {
		return m.PartyId
	}
	return ""
}

type SyncCredentialResponse struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SyncCredentialResponse) Reset()         { *m = SyncCredentialResponse{} }
func (m *SyncCredentialResponse) String() string { return proto.CompactTextString(m) }
func (*SyncCredentialResponse) ProtoMessage()    {}
func (*SyncCredentialResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_19dfbe8a464b7ce7, []int{5}
}

func (m *SyncCredentialResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SyncCredentialResponse.Unmarshal(m, b)
}
func (m *SyncCredentialResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SyncCredentialResponse.Marshal(b, m, deterministic)
}
func (m *SyncCredentialResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SyncCredentialResponse.Merge(m, src)
}
func (m *SyncCredentialResponse) XXX_Size() int {
	return xxx_messageInfo_SyncCredentialResponse.Size(m)
}
func (m *SyncCredentialResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SyncCredentialResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SyncCredentialResponse proto.InternalMessageInfo

func (m *SyncCredentialResponse) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type UnregisterCredentialRequest struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UnregisterCredentialRequest) Reset()         { *m = UnregisterCredentialRequest{} }
func (m *UnregisterCredentialRequest) String() string { return proto.CompactTextString(m) }
func (*UnregisterCredentialRequest) ProtoMessage()    {}
func (*UnregisterCredentialRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_19dfbe8a464b7ce7, []int{6}
}

func (m *UnregisterCredentialRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UnregisterCredentialRequest.Unmarshal(m, b)
}
func (m *UnregisterCredentialRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UnregisterCredentialRequest.Marshal(b, m, deterministic)
}
func (m *UnregisterCredentialRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UnregisterCredentialRequest.Merge(m, src)
}
func (m *UnregisterCredentialRequest) XXX_Size() int {
	return xxx_messageInfo_UnregisterCredentialRequest.Size(m)
}
func (m *UnregisterCredentialRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UnregisterCredentialRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UnregisterCredentialRequest proto.InternalMessageInfo

func (m *UnregisterCredentialRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type UnregisterCredentialResponse struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UnregisterCredentialResponse) Reset()         { *m = UnregisterCredentialResponse{} }
func (m *UnregisterCredentialResponse) String() string { return proto.CompactTextString(m) }
func (*UnregisterCredentialResponse) ProtoMessage()    {}
func (*UnregisterCredentialResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_19dfbe8a464b7ce7, []int{7}
}

func (m *UnregisterCredentialResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UnregisterCredentialResponse.Unmarshal(m, b)
}
func (m *UnregisterCredentialResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UnregisterCredentialResponse.Marshal(b, m, deterministic)
}
func (m *UnregisterCredentialResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UnregisterCredentialResponse.Merge(m, src)
}
func (m *UnregisterCredentialResponse) XXX_Size() int {
	return xxx_messageInfo_UnregisterCredentialResponse.Size(m)
}
func (m *UnregisterCredentialResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UnregisterCredentialResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UnregisterCredentialResponse proto.InternalMessageInfo

func (m *UnregisterCredentialResponse) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func init() {
	proto.RegisterType((*CreateCredentialRequest)(nil), "crediential.CreateCredentialRequest")
	proto.RegisterType((*CreateCredentialResponse)(nil), "crediential.CreateCredentialResponse")
	proto.RegisterType((*RegisterCredentialRequest)(nil), "crediential.RegisterCredentialRequest")
	proto.RegisterType((*RegisterCredentialResponse)(nil), "crediential.RegisterCredentialResponse")
	proto.RegisterType((*SyncCredentialRequest)(nil), "crediential.SyncCredentialRequest")
	proto.RegisterType((*SyncCredentialResponse)(nil), "crediential.SyncCredentialResponse")
	proto.RegisterType((*UnregisterCredentialRequest)(nil), "crediential.UnregisterCredentialRequest")
	proto.RegisterType((*UnregisterCredentialResponse)(nil), "crediential.UnregisterCredentialResponse")
}

func init() { proto.RegisterFile("ocpirpc/credential.proto", fileDescriptor_19dfbe8a464b7ce7) }

var fileDescriptor_19dfbe8a464b7ce7 = []byte{
	// 498 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x94, 0xcd, 0x6e, 0xd3, 0x40,
	0x14, 0x85, 0x65, 0x3b, 0x49, 0xd3, 0x1b, 0x14, 0xca, 0x88, 0x82, 0xeb, 0x76, 0x11, 0x5c, 0x5a,
	0x8c, 0xa0, 0x8e, 0x54, 0xde, 0xa0, 0xe9, 0x02, 0x36, 0x20, 0xb9, 0x65, 0x01, 0x1b, 0xcb, 0x3f,
	0x97, 0x74, 0xd4, 0xc4, 0x63, 0x66, 0xc6, 0x48, 0xd9, 0x22, 0xf1, 0x24, 0x3c, 0x25, 0x3b, 0x64,
	0x7b, 0x4a, 0xeb, 0xbf, 0x24, 0xec, 0x92, 0xe3, 0x33, 0xf7, 0xdc, 0xf9, 0xe6, 0xce, 0x80, 0xc9,
	0xa2, 0x94, 0xf2, 0x34, 0x9a, 0x46, 0x1c, 0x63, 0x4c, 0x24, 0x0d, 0x16, 0x6e, 0xca, 0x99, 0x64,
	0x64, 0x94, 0x2b, 0xb4, 0x94, 0xac, 0xa3, 0x3b, 0x5b, 0x98, 0x09, 0x9a, 0xa0, 0x10, 0x31, 0xca,
	0x80, 0x2a, 0xab, 0xfd, 0x47, 0x83, 0xe7, 0x33, 0x8e, 0x81, 0xc4, 0xd9, 0xbf, 0x2a, 0x1e, 0x7e,
	0xcf, 0x50, 0x48, 0xf2, 0x02, 0x1e, 0x45, 0x8b, 0xbc, 0x8c, 0x2f, 0xd9, 0x2d, 0x26, 0xa6, 0x36,
	0xd1, 0x9c, 0x5d, 0x6f, 0x54, 0x6a, 0xd7, 0xb9, 0x44, 0xf6, 0xc0, 0xc8, 0xf8, 0xc2, 0xd4, 0x8b,
	0x2f, 0xf9, 0x4f, 0x72, 0x0d, 0x8f, 0xef, 0x82, 0xfc, 0x32, 0xc9, 0x34, 0x26, 0x9a, 0x33, 0x3a,
	0x7f, 0xe3, 0xd6, 0x1a, 0x28, 0x63, 0x2f, 0x94, 0x78, 0x59, 0x88, 0x2a, 0xda, 0x1b, 0x87, 0x15,
	0xb9, 0x68, 0x85, 0x65, 0x89, 0xe4, 0x2b, 0x3f, 0x62, 0x31, 0x9a, 0x3d, 0xd5, 0x4a, 0xa9, 0xcd,
	0x58, 0x8c, 0xe4, 0x00, 0x86, 0x69, 0xc0, 0xe5, 0xca, 0xa7, 0xb1, 0xd9, 0x2f, 0x3e, 0xef, 0x14,
	0xff, 0x3f, 0xc4, 0x64, 0x1f, 0x06, 0x54, 0xf8, 0x37, 0x59, 0x68, 0x0e, 0x26, 0x9a, 0x33, 0xf4,
	0xfa, 0x54, 0xbc, 0xcf, 0x42, 0xfb, 0x97, 0x0e, 0x66, 0x73, 0xef, 0x22, 0x65, 0x89, 0x40, 0x32,
	0x06, 0x9d, 0xc6, 0xc5, 0x96, 0x0d, 0x4f, 0xa7, 0x71, 0x03, 0x86, 0xde, 0x09, 0xc3, 0xb8, 0x87,
	0xf1, 0xa9, 0x09, 0xa3, 0x57, 0xc0, 0x38, 0xad, 0xc3, 0xa8, 0x63, 0x28, 0xbb, 0xd8, 0xc8, 0xa1,
	0xbf, 0x9e, 0xc3, 0xa0, 0x8b, 0xc3, 0xce, 0x43, 0x0e, 0x1f, 0xe1, 0xc0, 0xc3, 0x39, 0x15, 0x12,
	0x79, 0x73, 0x08, 0xfe, 0x9f, 0x83, 0xfd, 0x16, 0xac, 0xb6, 0x7a, 0xed, 0x60, 0xed, 0x9f, 0x1a,
	0xec, 0x5f, 0xad, 0x92, 0x68, 0x73, 0xf4, 0x21, 0xec, 0x7e, 0xe3, 0x6c, 0xe9, 0xc7, 0x81, 0x44,
	0x95, 0x3b, 0xcc, 0x85, 0xcb, 0x40, 0x62, 0x83, 0x8c, 0xb1, 0x9e, 0x4c, 0xaf, 0x42, 0xc6, 0x76,
	0xe0, 0x59, 0xbd, 0x87, 0x8e, 0x76, 0xcf, 0xe0, 0xf0, 0x73, 0xc2, 0xb7, 0xc5, 0x65, 0xbb, 0x70,
	0xd4, 0x6e, 0x6f, 0x2f, 0x7f, 0xfe, 0xdb, 0x80, 0x27, 0xf7, 0xb6, 0x2b, 0xe4, 0x3f, 0x68, 0x84,
	0xc4, 0x87, 0xbd, 0xfa, 0xa0, 0x92, 0x97, 0xee, 0x83, 0x5b, 0xee, 0x76, 0xdc, 0x61, 0xeb, 0x64,
	0x83, 0x4b, 0xb5, 0x81, 0x40, 0x9a, 0x47, 0x46, 0x4e, 0x2b, 0x8b, 0x3b, 0x67, 0xc4, 0x7a, 0xb5,
	0xd1, 0xa7, 0x62, 0xbe, 0xc0, 0xb8, 0x8a, 0x99, 0xd8, 0x95, 0xa5, 0xad, 0x73, 0x60, 0x1d, 0xaf,
	0xf5, 0xa8, 0xd2, 0xb7, 0xf0, 0xb4, 0x0d, 0x34, 0x71, 0x2a, 0x8b, 0xd7, 0x1c, 0x9d, 0xf5, 0x7a,
	0x0b, 0x67, 0x19, 0x76, 0x71, 0xf2, 0xf5, 0x78, 0x4e, 0xe5, 0x4d, 0x16, 0xba, 0x11, 0x5b, 0x4e,
	0x45, 0x20, 0xe9, 0x92, 0x49, 0x36, 0x9d, 0xb3, 0xb3, 0xfc, 0xb1, 0x9d, 0xaa, 0x17, 0x37, 0x1c,
	0x14, 0x6f, 0xec, 0xbb, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xc1, 0x6c, 0x0d, 0xfe, 0xaa, 0x05,
	0x00, 0x00,
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
	CreateCredential(ctx context.Context, in *CreateCredentialRequest, opts ...grpc.CallOption) (*CreateCredentialResponse, error)
	RegisterCredential(ctx context.Context, in *RegisterCredentialRequest, opts ...grpc.CallOption) (*RegisterCredentialResponse, error)
	SyncCredential(ctx context.Context, in *SyncCredentialRequest, opts ...grpc.CallOption) (*SyncCredentialResponse, error)
	UnregisterCredential(ctx context.Context, in *UnregisterCredentialRequest, opts ...grpc.CallOption) (*UnregisterCredentialResponse, error)
}

type credentialServiceClient struct {
	cc *grpc.ClientConn
}

func NewCredentialServiceClient(cc *grpc.ClientConn) CredentialServiceClient {
	return &credentialServiceClient{cc}
}

func (c *credentialServiceClient) CreateCredential(ctx context.Context, in *CreateCredentialRequest, opts ...grpc.CallOption) (*CreateCredentialResponse, error) {
	out := new(CreateCredentialResponse)
	err := c.cc.Invoke(ctx, "/crediential.CredentialService/CreateCredential", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *credentialServiceClient) RegisterCredential(ctx context.Context, in *RegisterCredentialRequest, opts ...grpc.CallOption) (*RegisterCredentialResponse, error) {
	out := new(RegisterCredentialResponse)
	err := c.cc.Invoke(ctx, "/crediential.CredentialService/RegisterCredential", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *credentialServiceClient) SyncCredential(ctx context.Context, in *SyncCredentialRequest, opts ...grpc.CallOption) (*SyncCredentialResponse, error) {
	out := new(SyncCredentialResponse)
	err := c.cc.Invoke(ctx, "/crediential.CredentialService/SyncCredential", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *credentialServiceClient) UnregisterCredential(ctx context.Context, in *UnregisterCredentialRequest, opts ...grpc.CallOption) (*UnregisterCredentialResponse, error) {
	out := new(UnregisterCredentialResponse)
	err := c.cc.Invoke(ctx, "/crediential.CredentialService/UnregisterCredential", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CredentialServiceServer is the server API for CredentialService service.
type CredentialServiceServer interface {
	CreateCredential(context.Context, *CreateCredentialRequest) (*CreateCredentialResponse, error)
	RegisterCredential(context.Context, *RegisterCredentialRequest) (*RegisterCredentialResponse, error)
	SyncCredential(context.Context, *SyncCredentialRequest) (*SyncCredentialResponse, error)
	UnregisterCredential(context.Context, *UnregisterCredentialRequest) (*UnregisterCredentialResponse, error)
}

// UnimplementedCredentialServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCredentialServiceServer struct {
}

func (*UnimplementedCredentialServiceServer) CreateCredential(ctx context.Context, req *CreateCredentialRequest) (*CreateCredentialResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCredential not implemented")
}
func (*UnimplementedCredentialServiceServer) RegisterCredential(ctx context.Context, req *RegisterCredentialRequest) (*RegisterCredentialResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterCredential not implemented")
}
func (*UnimplementedCredentialServiceServer) SyncCredential(ctx context.Context, req *SyncCredentialRequest) (*SyncCredentialResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SyncCredential not implemented")
}
func (*UnimplementedCredentialServiceServer) UnregisterCredential(ctx context.Context, req *UnregisterCredentialRequest) (*UnregisterCredentialResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnregisterCredential not implemented")
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

func _CredentialService_RegisterCredential_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterCredentialRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CredentialServiceServer).RegisterCredential(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/crediential.CredentialService/RegisterCredential",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CredentialServiceServer).RegisterCredential(ctx, req.(*RegisterCredentialRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CredentialService_SyncCredential_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SyncCredentialRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CredentialServiceServer).SyncCredential(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/crediential.CredentialService/SyncCredential",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CredentialServiceServer).SyncCredential(ctx, req.(*SyncCredentialRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CredentialService_UnregisterCredential_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnregisterCredentialRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CredentialServiceServer).UnregisterCredential(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/crediential.CredentialService/UnregisterCredential",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CredentialServiceServer).UnregisterCredential(ctx, req.(*UnregisterCredentialRequest))
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
		{
			MethodName: "RegisterCredential",
			Handler:    _CredentialService_RegisterCredential_Handler,
		},
		{
			MethodName: "SyncCredential",
			Handler:    _CredentialService_SyncCredential_Handler,
		},
		{
			MethodName: "UnregisterCredential",
			Handler:    _CredentialService_UnregisterCredential_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ocpirpc/credential.proto",
}
