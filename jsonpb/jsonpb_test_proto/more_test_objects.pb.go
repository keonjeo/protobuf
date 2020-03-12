// Code generated by protoc-gen-go. DO NOT EDIT.
// source: more_test_objects.proto

package jsonpb

import (
	fmt "fmt"
	proto "github.com/keonjeo/protobuf/proto"
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

type Numeral int32

const (
	Numeral_UNKNOWN Numeral = 0
	Numeral_ARABIC  Numeral = 1
	Numeral_ROMAN   Numeral = 2
)

var Numeral_name = map[int32]string{
	0: "UNKNOWN",
	1: "ARABIC",
	2: "ROMAN",
}

var Numeral_value = map[string]int32{
	"UNKNOWN": 0,
	"ARABIC":  1,
	"ROMAN":   2,
}

func (x Numeral) String() string {
	return proto.EnumName(Numeral_name, int32(x))
}

func (Numeral) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e6c135db3023e377, []int{0}
}

type Simple3 struct {
	Dub                  float64  `protobuf:"fixed64,1,opt,name=dub,proto3" json:"dub,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Simple3) Reset()         { *m = Simple3{} }
func (m *Simple3) String() string { return proto.CompactTextString(m) }
func (*Simple3) ProtoMessage()    {}
func (*Simple3) Descriptor() ([]byte, []int) {
	return fileDescriptor_e6c135db3023e377, []int{0}
}

func (m *Simple3) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Simple3.Unmarshal(m, b)
}
func (m *Simple3) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Simple3.Marshal(b, m, deterministic)
}
func (m *Simple3) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Simple3.Merge(m, src)
}
func (m *Simple3) XXX_Size() int {
	return xxx_messageInfo_Simple3.Size(m)
}
func (m *Simple3) XXX_DiscardUnknown() {
	xxx_messageInfo_Simple3.DiscardUnknown(m)
}

var xxx_messageInfo_Simple3 proto.InternalMessageInfo

func (m *Simple3) GetDub() float64 {
	if m != nil {
		return m.Dub
	}
	return 0
}

type SimpleSlice3 struct {
	Slices               []string `protobuf:"bytes,1,rep,name=slices,proto3" json:"slices,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SimpleSlice3) Reset()         { *m = SimpleSlice3{} }
func (m *SimpleSlice3) String() string { return proto.CompactTextString(m) }
func (*SimpleSlice3) ProtoMessage()    {}
func (*SimpleSlice3) Descriptor() ([]byte, []int) {
	return fileDescriptor_e6c135db3023e377, []int{1}
}

func (m *SimpleSlice3) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SimpleSlice3.Unmarshal(m, b)
}
func (m *SimpleSlice3) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SimpleSlice3.Marshal(b, m, deterministic)
}
func (m *SimpleSlice3) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SimpleSlice3.Merge(m, src)
}
func (m *SimpleSlice3) XXX_Size() int {
	return xxx_messageInfo_SimpleSlice3.Size(m)
}
func (m *SimpleSlice3) XXX_DiscardUnknown() {
	xxx_messageInfo_SimpleSlice3.DiscardUnknown(m)
}

var xxx_messageInfo_SimpleSlice3 proto.InternalMessageInfo

func (m *SimpleSlice3) GetSlices() []string {
	if m != nil {
		return m.Slices
	}
	return nil
}

type SimpleMap3 struct {
	Stringy              map[string]string `protobuf:"bytes,1,rep,name=stringy,proto3" json:"stringy,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *SimpleMap3) Reset()         { *m = SimpleMap3{} }
func (m *SimpleMap3) String() string { return proto.CompactTextString(m) }
func (*SimpleMap3) ProtoMessage()    {}
func (*SimpleMap3) Descriptor() ([]byte, []int) {
	return fileDescriptor_e6c135db3023e377, []int{2}
}

func (m *SimpleMap3) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SimpleMap3.Unmarshal(m, b)
}
func (m *SimpleMap3) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SimpleMap3.Marshal(b, m, deterministic)
}
func (m *SimpleMap3) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SimpleMap3.Merge(m, src)
}
func (m *SimpleMap3) XXX_Size() int {
	return xxx_messageInfo_SimpleMap3.Size(m)
}
func (m *SimpleMap3) XXX_DiscardUnknown() {
	xxx_messageInfo_SimpleMap3.DiscardUnknown(m)
}

var xxx_messageInfo_SimpleMap3 proto.InternalMessageInfo

func (m *SimpleMap3) GetStringy() map[string]string {
	if m != nil {
		return m.Stringy
	}
	return nil
}

type SimpleNull3 struct {
	Simple               *Simple3 `protobuf:"bytes,1,opt,name=simple,proto3" json:"simple,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SimpleNull3) Reset()         { *m = SimpleNull3{} }
func (m *SimpleNull3) String() string { return proto.CompactTextString(m) }
func (*SimpleNull3) ProtoMessage()    {}
func (*SimpleNull3) Descriptor() ([]byte, []int) {
	return fileDescriptor_e6c135db3023e377, []int{3}
}

func (m *SimpleNull3) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SimpleNull3.Unmarshal(m, b)
}
func (m *SimpleNull3) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SimpleNull3.Marshal(b, m, deterministic)
}
func (m *SimpleNull3) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SimpleNull3.Merge(m, src)
}
func (m *SimpleNull3) XXX_Size() int {
	return xxx_messageInfo_SimpleNull3.Size(m)
}
func (m *SimpleNull3) XXX_DiscardUnknown() {
	xxx_messageInfo_SimpleNull3.DiscardUnknown(m)
}

var xxx_messageInfo_SimpleNull3 proto.InternalMessageInfo

func (m *SimpleNull3) GetSimple() *Simple3 {
	if m != nil {
		return m.Simple
	}
	return nil
}

type Mappy struct {
	Nummy                map[int64]int32    `protobuf:"bytes,1,rep,name=nummy,proto3" json:"nummy,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	Strry                map[string]string  `protobuf:"bytes,2,rep,name=strry,proto3" json:"strry,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Objjy                map[int32]*Simple3 `protobuf:"bytes,3,rep,name=objjy,proto3" json:"objjy,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Buggy                map[int64]string   `protobuf:"bytes,4,rep,name=buggy,proto3" json:"buggy,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Booly                map[bool]bool      `protobuf:"bytes,5,rep,name=booly,proto3" json:"booly,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	Enumy                map[string]Numeral `protobuf:"bytes,6,rep,name=enumy,proto3" json:"enumy,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3,enum=jsonpb.Numeral"`
	S32Booly             map[int32]bool     `protobuf:"bytes,7,rep,name=s32booly,proto3" json:"s32booly,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	S64Booly             map[int64]bool     `protobuf:"bytes,8,rep,name=s64booly,proto3" json:"s64booly,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	U32Booly             map[uint32]bool    `protobuf:"bytes,9,rep,name=u32booly,proto3" json:"u32booly,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	U64Booly             map[uint64]bool    `protobuf:"bytes,10,rep,name=u64booly,proto3" json:"u64booly,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *Mappy) Reset()         { *m = Mappy{} }
func (m *Mappy) String() string { return proto.CompactTextString(m) }
func (*Mappy) ProtoMessage()    {}
func (*Mappy) Descriptor() ([]byte, []int) {
	return fileDescriptor_e6c135db3023e377, []int{4}
}

func (m *Mappy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Mappy.Unmarshal(m, b)
}
func (m *Mappy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Mappy.Marshal(b, m, deterministic)
}
func (m *Mappy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Mappy.Merge(m, src)
}
func (m *Mappy) XXX_Size() int {
	return xxx_messageInfo_Mappy.Size(m)
}
func (m *Mappy) XXX_DiscardUnknown() {
	xxx_messageInfo_Mappy.DiscardUnknown(m)
}

var xxx_messageInfo_Mappy proto.InternalMessageInfo

func (m *Mappy) GetNummy() map[int64]int32 {
	if m != nil {
		return m.Nummy
	}
	return nil
}

func (m *Mappy) GetStrry() map[string]string {
	if m != nil {
		return m.Strry
	}
	return nil
}

func (m *Mappy) GetObjjy() map[int32]*Simple3 {
	if m != nil {
		return m.Objjy
	}
	return nil
}

func (m *Mappy) GetBuggy() map[int64]string {
	if m != nil {
		return m.Buggy
	}
	return nil
}

func (m *Mappy) GetBooly() map[bool]bool {
	if m != nil {
		return m.Booly
	}
	return nil
}

func (m *Mappy) GetEnumy() map[string]Numeral {
	if m != nil {
		return m.Enumy
	}
	return nil
}

func (m *Mappy) GetS32Booly() map[int32]bool {
	if m != nil {
		return m.S32Booly
	}
	return nil
}

func (m *Mappy) GetS64Booly() map[int64]bool {
	if m != nil {
		return m.S64Booly
	}
	return nil
}

func (m *Mappy) GetU32Booly() map[uint32]bool {
	if m != nil {
		return m.U32Booly
	}
	return nil
}

func (m *Mappy) GetU64Booly() map[uint64]bool {
	if m != nil {
		return m.U64Booly
	}
	return nil
}

func init() {
	proto.RegisterEnum("jsonpb.Numeral", Numeral_name, Numeral_value)
	proto.RegisterType((*Simple3)(nil), "jsonpb.Simple3")
	proto.RegisterType((*SimpleSlice3)(nil), "jsonpb.SimpleSlice3")
	proto.RegisterType((*SimpleMap3)(nil), "jsonpb.SimpleMap3")
	proto.RegisterMapType((map[string]string)(nil), "jsonpb.SimpleMap3.StringyEntry")
	proto.RegisterType((*SimpleNull3)(nil), "jsonpb.SimpleNull3")
	proto.RegisterType((*Mappy)(nil), "jsonpb.Mappy")
	proto.RegisterMapType((map[bool]bool)(nil), "jsonpb.Mappy.BoolyEntry")
	proto.RegisterMapType((map[int64]string)(nil), "jsonpb.Mappy.BuggyEntry")
	proto.RegisterMapType((map[string]Numeral)(nil), "jsonpb.Mappy.EnumyEntry")
	proto.RegisterMapType((map[int64]int32)(nil), "jsonpb.Mappy.NummyEntry")
	proto.RegisterMapType((map[int32]*Simple3)(nil), "jsonpb.Mappy.ObjjyEntry")
	proto.RegisterMapType((map[int32]bool)(nil), "jsonpb.Mappy.S32boolyEntry")
	proto.RegisterMapType((map[int64]bool)(nil), "jsonpb.Mappy.S64boolyEntry")
	proto.RegisterMapType((map[string]string)(nil), "jsonpb.Mappy.StrryEntry")
	proto.RegisterMapType((map[uint32]bool)(nil), "jsonpb.Mappy.U32boolyEntry")
	proto.RegisterMapType((map[uint64]bool)(nil), "jsonpb.Mappy.U64boolyEntry")
}

func init() { proto.RegisterFile("more_test_objects.proto", fileDescriptor_e6c135db3023e377) }

var fileDescriptor_e6c135db3023e377 = []byte{
	// 526 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x94, 0xdd, 0x6b, 0xdb, 0x3c,
	0x14, 0x87, 0x5f, 0x27, 0xf5, 0xd7, 0x49, 0xfb, 0x2e, 0x88, 0xb1, 0x99, 0xf4, 0x62, 0xc5, 0xb0,
	0xad, 0x0c, 0xe6, 0x8b, 0x78, 0x74, 0x5d, 0x77, 0x95, 0x8e, 0x5e, 0x94, 0x11, 0x07, 0x1c, 0xc2,
	0x2e, 0x4b, 0xdc, 0x99, 0x90, 0xcc, 0x5f, 0xd8, 0xd6, 0xc0, 0xd7, 0xfb, 0xbb, 0x07, 0xe3, 0x48,
	0x72, 0x2d, 0x07, 0x85, 0x6c, 0x77, 0x52, 0x7e, 0xcf, 0xe3, 0x73, 0x24, 0x1d, 0x02, 0x2f, 0xd3,
	0xbc, 0x8c, 0x1f, 0xea, 0xb8, 0xaa, 0x1f, 0xf2, 0x68, 0x17, 0x3f, 0xd6, 0x95, 0x57, 0x94, 0x79,
	0x9d, 0x13, 0x63, 0x57, 0xe5, 0x59, 0x11, 0xb9, 0xe7, 0x60, 0x2e, 0xb7, 0x69, 0x91, 0xc4, 0x3e,
	0x19, 0xc3, 0xf0, 0x3b, 0x8d, 0x1c, 0xed, 0x42, 0xbb, 0xd4, 0x42, 0x5c, 0xba, 0x6f, 0xe0, 0x94,
	0x87, 0xcb, 0x64, 0xfb, 0x18, 0xfb, 0xe4, 0x05, 0x18, 0x15, 0xae, 0x2a, 0x47, 0xbb, 0x18, 0x5e,
	0xda, 0xa1, 0xd8, 0xb9, 0xbf, 0x34, 0x00, 0x0e, 0xce, 0xd7, 0x85, 0x4f, 0x3e, 0x81, 0x59, 0xd5,
	0xe5, 0x36, 0xdb, 0x34, 0x8c, 0x1b, 0x4d, 0x5f, 0x79, 0xbc, 0x9a, 0xd7, 0x41, 0xde, 0x92, 0x13,
	0x77, 0x59, 0x5d, 0x36, 0x61, 0xcb, 0x4f, 0x6e, 0xe0, 0x54, 0x0e, 0xb0, 0xa7, 0x1f, 0x71, 0xc3,
	0x7a, 0xb2, 0x43, 0x5c, 0x92, 0xe7, 0xa0, 0xff, 0x5c, 0x27, 0x34, 0x76, 0x06, 0xec, 0x37, 0xbe,
	0xb9, 0x19, 0x5c, 0x6b, 0xee, 0x15, 0x8c, 0xf8, 0xf7, 0x03, 0x9a, 0x24, 0x3e, 0x79, 0x0b, 0x46,
	0xc5, 0xb6, 0xcc, 0x1e, 0x4d, 0x9f, 0xf5, 0x9b, 0xf0, 0x43, 0x11, 0xbb, 0xbf, 0x2d, 0xd0, 0xe7,
	0xeb, 0xa2, 0x68, 0x88, 0x07, 0x7a, 0x46, 0xd3, 0xb4, 0x6d, 0xdb, 0x69, 0x0d, 0x96, 0x7a, 0x01,
	0x46, 0xbc, 0x5f, 0x8e, 0x21, 0x5f, 0xd5, 0x65, 0xd9, 0x38, 0x03, 0x15, 0xbf, 0xc4, 0x48, 0xf0,
	0x0c, 0x43, 0x3e, 0x8f, 0x76, 0xbb, 0xc6, 0x19, 0xaa, 0xf8, 0x05, 0x46, 0x82, 0x67, 0x18, 0xf2,
	0x11, 0xdd, 0x6c, 0x1a, 0xe7, 0x44, 0xc5, 0xdf, 0x62, 0x24, 0x78, 0x86, 0x31, 0x3e, 0xcf, 0x93,
	0xc6, 0xd1, 0x95, 0x3c, 0x46, 0x2d, 0x8f, 0x6b, 0xe4, 0xe3, 0x8c, 0xa6, 0x8d, 0x63, 0xa8, 0xf8,
	0x3b, 0x8c, 0x04, 0xcf, 0x30, 0xf2, 0x11, 0xac, 0xca, 0x9f, 0xf2, 0x12, 0x26, 0x53, 0xce, 0xf7,
	0x8e, 0x2c, 0x52, 0x6e, 0x3d, 0xc1, 0x4c, 0xbc, 0xfa, 0xc0, 0x45, 0x4b, 0x29, 0x8a, 0xb4, 0x15,
	0xc5, 0x16, 0x45, 0xda, 0x56, 0xb4, 0x55, 0xe2, 0xaa, 0x5f, 0x91, 0x4a, 0x15, 0x69, 0x5b, 0x11,
	0x94, 0x62, 0xbf, 0x62, 0x0b, 0x4f, 0xae, 0x01, 0xba, 0x87, 0x96, 0xe7, 0x6f, 0xa8, 0x98, 0x3f,
	0x5d, 0x9a, 0x3f, 0x34, 0xbb, 0x27, 0xff, 0x97, 0xc9, 0x9d, 0xdc, 0x03, 0x74, 0x8f, 0x2f, 0x9b,
	0x3a, 0x37, 0x5f, 0xcb, 0xa6, 0x62, 0x92, 0xfb, 0x4d, 0x74, 0x73, 0x71, 0xac, 0x7d, 0x7b, 0xdf,
	0x7c, 0xba, 0x10, 0xd9, 0xb4, 0x14, 0xa6, 0xb5, 0xd7, 0x7e, 0x37, 0x2b, 0x8a, 0x83, 0xf7, 0xda,
	0xff, 0xbf, 0x6b, 0x3f, 0xa0, 0x69, 0x5c, 0xae, 0x13, 0xf9, 0x53, 0x9f, 0xe1, 0xac, 0x37, 0x43,
	0x8a, 0xcb, 0x38, 0xdc, 0x07, 0xca, 0xf2, 0xab, 0x1e, 0x3b, 0xfe, 0xbe, 0xbc, 0x3a, 0x54, 0xf9,
	0xec, 0x6f, 0xe4, 0x43, 0x95, 0x4f, 0x8e, 0xc8, 0xef, 0xde, 0x83, 0x29, 0x6e, 0x82, 0x8c, 0xc0,
	0x5c, 0x05, 0x5f, 0x83, 0xc5, 0xb7, 0x60, 0xfc, 0x1f, 0x01, 0x30, 0x66, 0xe1, 0xec, 0xf6, 0xfe,
	0xcb, 0x58, 0x23, 0x36, 0xe8, 0xe1, 0x62, 0x3e, 0x0b, 0xc6, 0x83, 0xc8, 0x60, 0x7f, 0xe0, 0xfe,
	0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xdc, 0x84, 0x34, 0xaf, 0xdb, 0x05, 0x00, 0x00,
}
