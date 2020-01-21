// Code generated by protoc-gen-go. DO NOT EDIT.
// source: briscola/pb/points.proto

package pb

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

type CardPointsRequest struct {
	CardNumber           uint32   `protobuf:"varint,1,opt,name=card_number,json=cardNumber,proto3" json:"card_number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CardPointsRequest) Reset()         { *m = CardPointsRequest{} }
func (m *CardPointsRequest) String() string { return proto.CompactTextString(m) }
func (*CardPointsRequest) ProtoMessage()    {}
func (*CardPointsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5fe2b4e9a2747153, []int{0}
}

func (m *CardPointsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CardPointsRequest.Unmarshal(m, b)
}
func (m *CardPointsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CardPointsRequest.Marshal(b, m, deterministic)
}
func (m *CardPointsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CardPointsRequest.Merge(m, src)
}
func (m *CardPointsRequest) XXX_Size() int {
	return xxx_messageInfo_CardPointsRequest.Size(m)
}
func (m *CardPointsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CardPointsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CardPointsRequest proto.InternalMessageInfo

func (m *CardPointsRequest) GetCardNumber() uint32 {
	if m != nil {
		return m.CardNumber
	}
	return 0
}

type CardPointsResponse struct {
	Points               uint32   `protobuf:"varint,1,opt,name=points,proto3" json:"points,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CardPointsResponse) Reset()         { *m = CardPointsResponse{} }
func (m *CardPointsResponse) String() string { return proto.CompactTextString(m) }
func (*CardPointsResponse) ProtoMessage()    {}
func (*CardPointsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5fe2b4e9a2747153, []int{1}
}

func (m *CardPointsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CardPointsResponse.Unmarshal(m, b)
}
func (m *CardPointsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CardPointsResponse.Marshal(b, m, deterministic)
}
func (m *CardPointsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CardPointsResponse.Merge(m, src)
}
func (m *CardPointsResponse) XXX_Size() int {
	return xxx_messageInfo_CardPointsResponse.Size(m)
}
func (m *CardPointsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CardPointsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CardPointsResponse proto.InternalMessageInfo

func (m *CardPointsResponse) GetPoints() uint32 {
	if m != nil {
		return m.Points
	}
	return 0
}

func init() {
	proto.RegisterType((*CardPointsRequest)(nil), "pb.CardPointsRequest")
	proto.RegisterType((*CardPointsResponse)(nil), "pb.CardPointsResponse")
}

func init() { proto.RegisterFile("briscola/pb/points.proto", fileDescriptor_5fe2b4e9a2747153) }

var fileDescriptor_5fe2b4e9a2747153 = []byte{
	// 127 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x48, 0x2a, 0xca, 0x2c,
	0x4e, 0xce, 0xcf, 0x49, 0xd4, 0x2f, 0x48, 0xd2, 0x2f, 0xc8, 0xcf, 0xcc, 0x2b, 0x29, 0xd6, 0x2b,
	0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x48, 0x52, 0x32, 0xe1, 0x12, 0x74, 0x4e, 0x2c, 0x4a,
	0x09, 0x00, 0x8b, 0x07, 0xa5, 0x16, 0x96, 0xa6, 0x16, 0x97, 0x08, 0xc9, 0x73, 0x71, 0x27, 0x27,
	0x16, 0xa5, 0xc4, 0xe7, 0x95, 0xe6, 0x26, 0xa5, 0x16, 0x49, 0x30, 0x2a, 0x30, 0x6a, 0xf0, 0x06,
	0x71, 0x81, 0x84, 0xfc, 0xc0, 0x22, 0x4a, 0x3a, 0x5c, 0x42, 0xc8, 0xba, 0x8a, 0x0b, 0xf2, 0xf3,
	0x8a, 0x53, 0x85, 0xc4, 0xb8, 0xd8, 0x20, 0xe6, 0x43, 0x75, 0x40, 0x79, 0x49, 0x6c, 0x60, 0xeb,
	0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0xa6, 0xed, 0x68, 0xd4, 0x8a, 0x00, 0x00, 0x00,
}