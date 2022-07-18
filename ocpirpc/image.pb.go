// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ocpirpc/image.proto

package ocpirpc

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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type CreateImageRequest struct {
	Url                  string   `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	Thumbnail            string   `protobuf:"bytes,2,opt,name=thumbnail,proto3" json:"thumbnail,omitempty"`
	Category             string   `protobuf:"bytes,3,opt,name=category,proto3" json:"category,omitempty"`
	Type                 string   `protobuf:"bytes,4,opt,name=type,proto3" json:"type,omitempty"`
	Width                int32    `protobuf:"varint,5,opt,name=width,proto3" json:"width,omitempty"`
	Height               int32    `protobuf:"varint,6,opt,name=height,proto3" json:"height,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateImageRequest) Reset()         { *m = CreateImageRequest{} }
func (m *CreateImageRequest) String() string { return proto.CompactTextString(m) }
func (*CreateImageRequest) ProtoMessage()    {}
func (*CreateImageRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d155be59193fd55a, []int{0}
}

func (m *CreateImageRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateImageRequest.Unmarshal(m, b)
}
func (m *CreateImageRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateImageRequest.Marshal(b, m, deterministic)
}
func (m *CreateImageRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateImageRequest.Merge(m, src)
}
func (m *CreateImageRequest) XXX_Size() int {
	return xxx_messageInfo_CreateImageRequest.Size(m)
}
func (m *CreateImageRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateImageRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateImageRequest proto.InternalMessageInfo

func (m *CreateImageRequest) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *CreateImageRequest) GetThumbnail() string {
	if m != nil {
		return m.Thumbnail
	}
	return ""
}

func (m *CreateImageRequest) GetCategory() string {
	if m != nil {
		return m.Category
	}
	return ""
}

func (m *CreateImageRequest) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *CreateImageRequest) GetWidth() int32 {
	if m != nil {
		return m.Width
	}
	return 0
}

func (m *CreateImageRequest) GetHeight() int32 {
	if m != nil {
		return m.Height
	}
	return 0
}

type ImageResponse struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Url                  string   `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	Thumbnail            string   `protobuf:"bytes,3,opt,name=thumbnail,proto3" json:"thumbnail,omitempty"`
	Category             string   `protobuf:"bytes,4,opt,name=category,proto3" json:"category,omitempty"`
	Type                 string   `protobuf:"bytes,5,opt,name=type,proto3" json:"type,omitempty"`
	Width                int32    `protobuf:"varint,6,opt,name=width,proto3" json:"width,omitempty"`
	Height               int32    `protobuf:"varint,7,opt,name=height,proto3" json:"height,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ImageResponse) Reset()         { *m = ImageResponse{} }
func (m *ImageResponse) String() string { return proto.CompactTextString(m) }
func (*ImageResponse) ProtoMessage()    {}
func (*ImageResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d155be59193fd55a, []int{1}
}

func (m *ImageResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ImageResponse.Unmarshal(m, b)
}
func (m *ImageResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ImageResponse.Marshal(b, m, deterministic)
}
func (m *ImageResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ImageResponse.Merge(m, src)
}
func (m *ImageResponse) XXX_Size() int {
	return xxx_messageInfo_ImageResponse.Size(m)
}
func (m *ImageResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ImageResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ImageResponse proto.InternalMessageInfo

func (m *ImageResponse) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ImageResponse) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *ImageResponse) GetThumbnail() string {
	if m != nil {
		return m.Thumbnail
	}
	return ""
}

func (m *ImageResponse) GetCategory() string {
	if m != nil {
		return m.Category
	}
	return ""
}

func (m *ImageResponse) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *ImageResponse) GetWidth() int32 {
	if m != nil {
		return m.Width
	}
	return 0
}

func (m *ImageResponse) GetHeight() int32 {
	if m != nil {
		return m.Height
	}
	return 0
}

func init() {
	proto.RegisterType((*CreateImageRequest)(nil), "image.CreateImageRequest")
	proto.RegisterType((*ImageResponse)(nil), "image.ImageResponse")
}

func init() { proto.RegisterFile("ocpirpc/image.proto", fileDescriptor_d155be59193fd55a) }

var fileDescriptor_d155be59193fd55a = []byte{
	// 249 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0xd1, 0x3f, 0x4b, 0xf4, 0x40,
	0x10, 0x06, 0x70, 0x36, 0xc9, 0xe6, 0x7d, 0x6f, 0x40, 0x91, 0x51, 0x64, 0x11, 0x8b, 0xe3, 0x44,
	0xb8, 0xc6, 0x4b, 0xe1, 0x37, 0xd0, 0xca, 0x36, 0xa5, 0xdd, 0x26, 0x59, 0x76, 0x17, 0x2e, 0x37,
	0xeb, 0x66, 0x82, 0xdc, 0xd7, 0xb1, 0xf7, 0x3b, 0x4a, 0xd6, 0xf8, 0xa7, 0xb8, 0xb3, 0x9b, 0xe7,
	0x99, 0xe6, 0x07, 0x0f, 0x9c, 0x53, 0x1b, 0x7c, 0x0c, 0x6d, 0xe5, 0x7b, 0x6d, 0xcd, 0x26, 0x44,
	0x62, 0x42, 0x99, 0xc2, 0xea, 0x4d, 0x00, 0x3e, 0x46, 0xa3, 0xd9, 0x3c, 0x4d, 0xb9, 0x36, 0x2f,
	0xa3, 0x19, 0x18, 0xcf, 0x20, 0x1f, 0xe3, 0x56, 0x89, 0xa5, 0x58, 0x2f, 0xea, 0xe9, 0xc4, 0x6b,
	0x58, 0xb0, 0x1b, 0xfb, 0x66, 0xa7, 0xfd, 0x56, 0x65, 0xa9, 0xff, 0x29, 0xf0, 0x0a, 0xfe, 0xb7,
	0x9a, 0x8d, 0xa5, 0xb8, 0x57, 0x79, 0x7a, 0x7e, 0x67, 0x44, 0x28, 0x78, 0x1f, 0x8c, 0x2a, 0x52,
	0x9f, 0x6e, 0xbc, 0x00, 0xf9, 0xea, 0x3b, 0x76, 0x4a, 0x2e, 0xc5, 0x5a, 0xd6, 0x9f, 0x01, 0x2f,
	0xa1, 0x74, 0xc6, 0x5b, 0xc7, 0xaa, 0x4c, 0xf5, 0x9c, 0x56, 0xef, 0x02, 0x4e, 0x66, 0xde, 0x10,
	0x68, 0x37, 0x18, 0x3c, 0x85, 0xcc, 0x77, 0x89, 0x97, 0xd7, 0x99, 0xef, 0xbe, 0xbc, 0xd9, 0x11,
	0x6f, 0xfe, 0x97, 0xb7, 0x38, 0xe2, 0x95, 0x87, 0xbc, 0xe5, 0x61, 0xef, 0xbf, 0xdf, 0xde, 0x87,
	0xdb, 0xe7, 0x1b, 0xeb, 0xd9, 0x8d, 0xcd, 0xa6, 0xa5, 0xbe, 0x1a, 0x34, 0xfb, 0x9e, 0x98, 0x2a,
	0x4b, 0x77, 0xd3, 0x12, 0xd5, 0x3c, 0x47, 0x53, 0xa6, 0x25, 0xee, 0x3f, 0x02, 0x00, 0x00, 0xff,
	0xff, 0x2b, 0xfe, 0xd2, 0x36, 0xa0, 0x01, 0x00, 0x00,
}
