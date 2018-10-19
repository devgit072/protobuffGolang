// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protobuffGolang/test1/enumProto.proto

package test1

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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type DayOfWeek int32

const (
	DayOfWeek_UNKNOWN   DayOfWeek = 0
	DayOfWeek_MONDAY    DayOfWeek = 1
	DayOfWeek_TUESDAY   DayOfWeek = 2
	DayOfWeek_WEDNESDAY DayOfWeek = 3
	DayOfWeek_THRUSTDAY DayOfWeek = 4
	DayOfWeek_FRIDAY    DayOfWeek = 5
	DayOfWeek_SATURDAY  DayOfWeek = 6
	DayOfWeek_SUNDAY    DayOfWeek = 7
)

var DayOfWeek_name = map[int32]string{
	0: "UNKNOWN",
	1: "MONDAY",
	2: "TUESDAY",
	3: "WEDNESDAY",
	4: "THRUSTDAY",
	5: "FRIDAY",
	6: "SATURDAY",
	7: "SUNDAY",
}

var DayOfWeek_value = map[string]int32{
	"UNKNOWN":   0,
	"MONDAY":    1,
	"TUESDAY":   2,
	"WEDNESDAY": 3,
	"THRUSTDAY": 4,
	"FRIDAY":    5,
	"SATURDAY":  6,
	"SUNDAY":    7,
}

func (x DayOfWeek) String() string {
	return proto.EnumName(DayOfWeek_name, int32(x))
}

func (DayOfWeek) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b01e59999dab7046, []int{0}
}

type EnumProto struct {
	Id                   int32     `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	WeekOfday            DayOfWeek `protobuf:"varint,2,opt,name=weekOfday,proto3,enum=test1.DayOfWeek" json:"weekOfday,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *EnumProto) Reset()         { *m = EnumProto{} }
func (m *EnumProto) String() string { return proto.CompactTextString(m) }
func (*EnumProto) ProtoMessage()    {}
func (*EnumProto) Descriptor() ([]byte, []int) {
	return fileDescriptor_b01e59999dab7046, []int{0}
}

func (m *EnumProto) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EnumProto.Unmarshal(m, b)
}
func (m *EnumProto) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EnumProto.Marshal(b, m, deterministic)
}
func (m *EnumProto) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EnumProto.Merge(m, src)
}
func (m *EnumProto) XXX_Size() int {
	return xxx_messageInfo_EnumProto.Size(m)
}
func (m *EnumProto) XXX_DiscardUnknown() {
	xxx_messageInfo_EnumProto.DiscardUnknown(m)
}

var xxx_messageInfo_EnumProto proto.InternalMessageInfo

func (m *EnumProto) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *EnumProto) GetWeekOfday() DayOfWeek {
	if m != nil {
		return m.WeekOfday
	}
	return DayOfWeek_UNKNOWN
}

func init() {
	proto.RegisterEnum("test1.DayOfWeek", DayOfWeek_name, DayOfWeek_value)
	proto.RegisterType((*EnumProto)(nil), "test1.EnumProto")
}

func init() {
	proto.RegisterFile("protobuffGolang/test1/enumProto.proto", fileDescriptor_b01e59999dab7046)
}

var fileDescriptor_b01e59999dab7046 = []byte{
	// 214 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x2d, 0x28, 0xca, 0x2f,
	0xc9, 0x4f, 0x2a, 0x4d, 0x4b, 0x73, 0xcf, 0xcf, 0x49, 0xcc, 0x4b, 0xd7, 0x2f, 0x49, 0x2d, 0x2e,
	0x31, 0xd4, 0x4f, 0xcd, 0x2b, 0xcd, 0x0d, 0x00, 0xc9, 0xe8, 0x81, 0xe5, 0x85, 0x58, 0xc1, 0xc2,
	0x4a, 0xde, 0x5c, 0x9c, 0xae, 0x30, 0x19, 0x21, 0x3e, 0x2e, 0xa6, 0xcc, 0x14, 0x09, 0x46, 0x05,
	0x46, 0x0d, 0xd6, 0x20, 0xa6, 0xcc, 0x14, 0x21, 0x3d, 0x2e, 0xce, 0xf2, 0xd4, 0xd4, 0x6c, 0xff,
	0xb4, 0x94, 0xc4, 0x4a, 0x09, 0x26, 0x05, 0x46, 0x0d, 0x3e, 0x23, 0x01, 0x3d, 0xb0, 0x3e, 0x3d,
	0x97, 0xc4, 0x4a, 0xff, 0xb4, 0xf0, 0xd4, 0xd4, 0xec, 0x20, 0x84, 0x12, 0xad, 0x52, 0x2e, 0x4e,
	0xb8, 0xb8, 0x10, 0x37, 0x17, 0x7b, 0xa8, 0x9f, 0xb7, 0x9f, 0x7f, 0xb8, 0x9f, 0x00, 0x83, 0x10,
	0x17, 0x17, 0x9b, 0xaf, 0xbf, 0x9f, 0x8b, 0x63, 0xa4, 0x00, 0x23, 0x48, 0x22, 0x24, 0xd4, 0x35,
	0x18, 0xc4, 0x61, 0x12, 0xe2, 0xe5, 0xe2, 0x0c, 0x77, 0x75, 0xf1, 0x83, 0x70, 0x99, 0x41, 0xdc,
	0x10, 0x8f, 0xa0, 0xd0, 0xe0, 0x10, 0x10, 0x97, 0x05, 0xa4, 0xcd, 0x2d, 0xc8, 0x13, 0xc4, 0x66,
	0x15, 0xe2, 0xe1, 0xe2, 0x08, 0x76, 0x0c, 0x09, 0x0d, 0x02, 0xf1, 0xd8, 0x40, 0x32, 0xc1, 0xa1,
	0x60, 0x03, 0xd9, 0x93, 0xd8, 0xc0, 0x3e, 0x32, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0xa4, 0x52,
	0x71, 0x2c, 0xfa, 0x00, 0x00, 0x00,
}