// Code generated by protoc-gen-gogo.
// source: enumcustomname.proto
// DO NOT EDIT!

/*
	Package enumcustomname is a generated protocol buffer package.

	Package enumcustomname tests the behavior of enum_customname and
	enumvalue_customname extensions.

	It is generated from these files:
		enumcustomname.proto

	It has these top-level messages:
		OnlyEnums
*/
package enumcustomname

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import test "github.com/gogo/protobuf/test"

import strconv "strconv"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type MyCustomEnum int32

const (
	// The following field will take on the custom name and the prefix, joined
	// by an underscore.
	MyCustomEnum_MyBetterNameA MyCustomEnum = 0
	MyCustomEnum_B             MyCustomEnum = 1
)

var MyCustomEnum_name = map[int32]string{
	0: "A",
	1: "B",
}
var MyCustomEnum_value = map[string]int32{
	"A": 0,
	"B": 1,
}

func (x MyCustomEnum) Enum() *MyCustomEnum {
	p := new(MyCustomEnum)
	*p = x
	return p
}
func (x MyCustomEnum) String() string {
	return proto.EnumName(MyCustomEnum_name, int32(x))
}
func (x *MyCustomEnum) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(MyCustomEnum_value, data, "MyCustomEnum")
	if err != nil {
		return err
	}
	*x = MyCustomEnum(value)
	return nil
}
func (MyCustomEnum) EnumDescriptor() ([]byte, []int) { return fileDescriptorEnumcustomname, []int{0} }

type MyCustomUnprefixedEnum int32

const (
	MyBetterNameUnprefixedA MyCustomUnprefixedEnum = 0
	UNPREFIXED_B            MyCustomUnprefixedEnum = 1
)

var MyCustomUnprefixedEnum_name = map[int32]string{
	0: "UNPREFIXED_A",
	1: "UNPREFIXED_B",
}
var MyCustomUnprefixedEnum_value = map[string]int32{
	"UNPREFIXED_A": 0,
	"UNPREFIXED_B": 1,
}

func (x MyCustomUnprefixedEnum) Enum() *MyCustomUnprefixedEnum {
	p := new(MyCustomUnprefixedEnum)
	*p = x
	return p
}
func (x MyCustomUnprefixedEnum) MarshalJSON() ([]byte, error) {
	return proto.MarshalJSONEnum(MyCustomUnprefixedEnum_name, int32(x))
}
func (x *MyCustomUnprefixedEnum) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(MyCustomUnprefixedEnum_value, data, "MyCustomUnprefixedEnum")
	if err != nil {
		return err
	}
	*x = MyCustomUnprefixedEnum(value)
	return nil
}
func (MyCustomUnprefixedEnum) EnumDescriptor() ([]byte, []int) {
	return fileDescriptorEnumcustomname, []int{1}
}

type MyEnumWithEnumStringer int32

const (
	MyEnumWithEnumStringer_EnumValueStringerA MyEnumWithEnumStringer = 0
	MyEnumWithEnumStringer_STRINGER_B         MyEnumWithEnumStringer = 1
)

var MyEnumWithEnumStringer_name = map[int32]string{
	0: "STRINGER_A",
	1: "STRINGER_B",
}
var MyEnumWithEnumStringer_value = map[string]int32{
	"STRINGER_A": 0,
	"STRINGER_B": 1,
}

func (x MyEnumWithEnumStringer) Enum() *MyEnumWithEnumStringer {
	p := new(MyEnumWithEnumStringer)
	*p = x
	return p
}
func (x MyEnumWithEnumStringer) MarshalJSON() ([]byte, error) {
	return proto.MarshalJSONEnum(MyEnumWithEnumStringer_name, int32(x))
}
func (x *MyEnumWithEnumStringer) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(MyEnumWithEnumStringer_value, data, "MyEnumWithEnumStringer")
	if err != nil {
		return err
	}
	*x = MyEnumWithEnumStringer(value)
	return nil
}
func (MyEnumWithEnumStringer) EnumDescriptor() ([]byte, []int) {
	return fileDescriptorEnumcustomname, []int{2}
}

type OnlyEnums struct {
	MyEnum                         *MyCustomEnum               `protobuf:"varint,1,opt,name=my_enum,json=myEnum,enum=enumcustomname.MyCustomEnum" json:"my_enum,omitempty"`
	MyEnumDefaultA                 *MyCustomEnum               `protobuf:"varint,2,opt,name=my_enum_default_a,json=myEnumDefaultA,enum=enumcustomname.MyCustomEnum,def=0" json:"my_enum_default_a,omitempty"`
	MyEnumDefaultB                 *MyCustomEnum               `protobuf:"varint,3,opt,name=my_enum_default_b,json=myEnumDefaultB,enum=enumcustomname.MyCustomEnum,def=1" json:"my_enum_default_b,omitempty"`
	MyUnprefixedEnum               *MyCustomUnprefixedEnum     `protobuf:"varint,4,opt,name=my_unprefixed_enum,json=myUnprefixedEnum,enum=enumcustomname.MyCustomUnprefixedEnum" json:"my_unprefixed_enum,omitempty"`
	MyUnprefixedEnumDefaultA       *MyCustomUnprefixedEnum     `protobuf:"varint,5,opt,name=my_unprefixed_enum_default_a,json=myUnprefixedEnumDefaultA,enum=enumcustomname.MyCustomUnprefixedEnum,def=0" json:"my_unprefixed_enum_default_a,omitempty"`
	MyUnprefixedEnumDefaultB       *MyCustomUnprefixedEnum     `protobuf:"varint,6,opt,name=my_unprefixed_enum_default_b,json=myUnprefixedEnumDefaultB,enum=enumcustomname.MyCustomUnprefixedEnum,def=1" json:"my_unprefixed_enum_default_b,omitempty"`
	YetAnotherTestEnum             *test.YetAnotherTestEnum    `protobuf:"varint,7,opt,name=yet_another_test_enum,json=yetAnotherTestEnum,enum=test.YetAnotherTestEnum" json:"yet_another_test_enum,omitempty"`
	YetAnotherTestEnumDefaultAa    *test.YetAnotherTestEnum    `protobuf:"varint,8,opt,name=yet_another_test_enum_default_aa,json=yetAnotherTestEnumDefaultAa,enum=test.YetAnotherTestEnum,def=0" json:"yet_another_test_enum_default_aa,omitempty"`
	YetAnotherTestEnumDefaultBb    *test.YetAnotherTestEnum    `protobuf:"varint,9,opt,name=yet_another_test_enum_default_bb,json=yetAnotherTestEnumDefaultBb,enum=test.YetAnotherTestEnum,def=1" json:"yet_another_test_enum_default_bb,omitempty"`
	YetYetAnotherTestEnum          *test.YetYetAnotherTestEnum `protobuf:"varint,10,opt,name=yet_yet_another_test_enum,json=yetYetAnotherTestEnum,enum=test.YetYetAnotherTestEnum" json:"yet_yet_another_test_enum,omitempty"`
	YetYetAnotherTestEnumDefaultCc *test.YetYetAnotherTestEnum `protobuf:"varint,11,opt,name=yet_yet_another_test_enum_default_cc,json=yetYetAnotherTestEnumDefaultCc,enum=test.YetYetAnotherTestEnum,def=0" json:"yet_yet_another_test_enum_default_cc,omitempty"`
	YetYetAnotherTestEnumDefaultDd *test.YetYetAnotherTestEnum `protobuf:"varint,12,opt,name=yet_yet_another_test_enum_default_dd,json=yetYetAnotherTestEnumDefaultDd,enum=test.YetYetAnotherTestEnum,def=1" json:"yet_yet_another_test_enum_default_dd,omitempty"`
	XXX_unrecognized               []byte                      `json:"-"`
}

func (m *OnlyEnums) Reset()                    { *m = OnlyEnums{} }
func (m *OnlyEnums) String() string            { return proto.CompactTextString(m) }
func (*OnlyEnums) ProtoMessage()               {}
func (*OnlyEnums) Descriptor() ([]byte, []int) { return fileDescriptorEnumcustomname, []int{0} }

const Default_OnlyEnums_MyEnumDefaultA MyCustomEnum = MyCustomEnum_MyBetterNameA
const Default_OnlyEnums_MyEnumDefaultB MyCustomEnum = MyCustomEnum_B
const Default_OnlyEnums_MyUnprefixedEnumDefaultA MyCustomUnprefixedEnum = MyBetterNameUnprefixedA
const Default_OnlyEnums_MyUnprefixedEnumDefaultB MyCustomUnprefixedEnum = UNPREFIXED_B
const Default_OnlyEnums_YetAnotherTestEnumDefaultAa test.YetAnotherTestEnum = test.AA
const Default_OnlyEnums_YetAnotherTestEnumDefaultBb test.YetAnotherTestEnum = test.BetterYetBB
const Default_OnlyEnums_YetYetAnotherTestEnumDefaultCc test.YetYetAnotherTestEnum = test.YetYetAnotherTestEnum_CC
const Default_OnlyEnums_YetYetAnotherTestEnumDefaultDd test.YetYetAnotherTestEnum = test.YetYetAnotherTestEnum_BetterYetDD

func (m *OnlyEnums) GetMyEnum() MyCustomEnum {
	if m != nil && m.MyEnum != nil {
		return *m.MyEnum
	}
	return MyCustomEnum_MyBetterNameA
}

func (m *OnlyEnums) GetMyEnumDefaultA() MyCustomEnum {
	if m != nil && m.MyEnumDefaultA != nil {
		return *m.MyEnumDefaultA
	}
	return Default_OnlyEnums_MyEnumDefaultA
}

func (m *OnlyEnums) GetMyEnumDefaultB() MyCustomEnum {
	if m != nil && m.MyEnumDefaultB != nil {
		return *m.MyEnumDefaultB
	}
	return Default_OnlyEnums_MyEnumDefaultB
}

func (m *OnlyEnums) GetMyUnprefixedEnum() MyCustomUnprefixedEnum {
	if m != nil && m.MyUnprefixedEnum != nil {
		return *m.MyUnprefixedEnum
	}
	return MyBetterNameUnprefixedA
}

func (m *OnlyEnums) GetMyUnprefixedEnumDefaultA() MyCustomUnprefixedEnum {
	if m != nil && m.MyUnprefixedEnumDefaultA != nil {
		return *m.MyUnprefixedEnumDefaultA
	}
	return Default_OnlyEnums_MyUnprefixedEnumDefaultA
}

func (m *OnlyEnums) GetMyUnprefixedEnumDefaultB() MyCustomUnprefixedEnum {
	if m != nil && m.MyUnprefixedEnumDefaultB != nil {
		return *m.MyUnprefixedEnumDefaultB
	}
	return Default_OnlyEnums_MyUnprefixedEnumDefaultB
}

func (m *OnlyEnums) GetYetAnotherTestEnum() test.YetAnotherTestEnum {
	if m != nil && m.YetAnotherTestEnum != nil {
		return *m.YetAnotherTestEnum
	}
	return test.AA
}

func (m *OnlyEnums) GetYetAnotherTestEnumDefaultAa() test.YetAnotherTestEnum {
	if m != nil && m.YetAnotherTestEnumDefaultAa != nil {
		return *m.YetAnotherTestEnumDefaultAa
	}
	return Default_OnlyEnums_YetAnotherTestEnumDefaultAa
}

func (m *OnlyEnums) GetYetAnotherTestEnumDefaultBb() test.YetAnotherTestEnum {
	if m != nil && m.YetAnotherTestEnumDefaultBb != nil {
		return *m.YetAnotherTestEnumDefaultBb
	}
	return Default_OnlyEnums_YetAnotherTestEnumDefaultBb
}

func (m *OnlyEnums) GetYetYetAnotherTestEnum() test.YetYetAnotherTestEnum {
	if m != nil && m.YetYetAnotherTestEnum != nil {
		return *m.YetYetAnotherTestEnum
	}
	return test.YetYetAnotherTestEnum_CC
}

func (m *OnlyEnums) GetYetYetAnotherTestEnumDefaultCc() test.YetYetAnotherTestEnum {
	if m != nil && m.YetYetAnotherTestEnumDefaultCc != nil {
		return *m.YetYetAnotherTestEnumDefaultCc
	}
	return Default_OnlyEnums_YetYetAnotherTestEnumDefaultCc
}

func (m *OnlyEnums) GetYetYetAnotherTestEnumDefaultDd() test.YetYetAnotherTestEnum {
	if m != nil && m.YetYetAnotherTestEnumDefaultDd != nil {
		return *m.YetYetAnotherTestEnumDefaultDd
	}
	return Default_OnlyEnums_YetYetAnotherTestEnumDefaultDd
}

func init() {
	proto.RegisterType((*OnlyEnums)(nil), "enumcustomname.OnlyEnums")
	proto.RegisterEnum("enumcustomname.MyCustomEnum", MyCustomEnum_name, MyCustomEnum_value)
	proto.RegisterEnum("enumcustomname.MyCustomUnprefixedEnum", MyCustomUnprefixedEnum_name, MyCustomUnprefixedEnum_value)
	proto.RegisterEnum("enumcustomname.MyEnumWithEnumStringer", MyEnumWithEnumStringer_name, MyEnumWithEnumStringer_value)
}
func (x MyEnumWithEnumStringer) String() string {
	s, ok := MyEnumWithEnumStringer_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}

func init() { proto.RegisterFile("enumcustomname.proto", fileDescriptorEnumcustomname) }

var fileDescriptorEnumcustomname = []byte{
	// 546 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x94, 0x4d, 0x8f, 0xd2, 0x40,
	0x18, 0xc7, 0x29, 0xba, 0x2c, 0x3b, 0x22, 0xe9, 0x4e, 0x14, 0x47, 0x30, 0x0d, 0xd9, 0x18, 0x63,
	0x30, 0x0b, 0x89, 0x47, 0x3c, 0x4d, 0x29, 0x9a, 0x8d, 0x01, 0x4d, 0x77, 0xf1, 0xed, 0xd2, 0xf4,
	0x65, 0x60, 0x49, 0x68, 0xbb, 0x29, 0xd3, 0x68, 0xbf, 0x81, 0xf1, 0x3b, 0x78, 0x92, 0x83, 0x47,
	0xcf, 0x9e, 0xfd, 0x60, 0xce, 0x4c, 0x17, 0x16, 0xda, 0x52, 0xc8, 0x9e, 0x86, 0x99, 0xfc, 0x9f,
	0xdf, 0x6f, 0xe6, 0xc9, 0x53, 0xc0, 0x03, 0xe2, 0x85, 0xae, 0x1d, 0xce, 0xa9, 0xef, 0x7a, 0xa6,
	0x4b, 0xda, 0x57, 0x81, 0x4f, 0x7d, 0x58, 0xdd, 0x3c, 0xad, 0x9f, 0x4e, 0xa6, 0xf4, 0x32, 0xb4,
	0xda, 0xb6, 0xef, 0x76, 0x26, 0xfe, 0xc4, 0xef, 0x88, 0x98, 0x15, 0x8e, 0xc5, 0x4e, 0x6c, 0xc4,
	0xaf, 0xb8, 0xbc, 0xfe, 0x62, 0x6b, 0x9c, 0x92, 0x39, 0xed, 0xd0, 0x4b, 0xc2, 0xd7, 0x38, 0x7c,
	0xf2, 0xaf, 0x0c, 0x8e, 0xde, 0x79, 0xb3, 0xa8, 0xcf, 0x94, 0x73, 0xd8, 0x01, 0x87, 0x6e, 0x64,
	0x70, 0x3d, 0x92, 0x9a, 0xd2, 0xf3, 0xea, 0xcb, 0x5a, 0x3b, 0x71, 0xc3, 0x81, 0x48, 0xea, 0x25,
	0x57, 0xac, 0x50, 0x03, 0xc7, 0xd7, 0x05, 0x86, 0x43, 0xc6, 0x66, 0x38, 0xa3, 0x86, 0x89, 0x8a,
	0x79, 0xa5, 0x5d, 0x09, 0xeb, 0xd5, 0xb8, 0x5a, 0x8b, 0x2b, 0x70, 0x16, 0xc5, 0x42, 0x77, 0xf2,
	0x29, 0x6a, 0x82, 0xa2, 0xc2, 0x21, 0x80, 0x8c, 0x12, 0x7a, 0x57, 0x01, 0x19, 0x4f, 0xbf, 0x11,
	0x27, 0x7e, 0xc7, 0x5d, 0x81, 0x69, 0xa6, 0x31, 0xa3, 0x55, 0x50, 0xbc, 0x48, 0x76, 0x13, 0x27,
	0xd0, 0x03, 0x4f, 0xd2, 0xbc, 0xb5, 0x67, 0x1e, 0xec, 0x47, 0xee, 0x56, 0x46, 0xc3, 0xf7, 0x7a,
	0xff, 0xf5, 0xd9, 0xa7, 0xbe, 0x66, 0x60, 0x1d, 0x25, 0x3d, 0xab, 0x2e, 0xe4, 0xfb, 0x2c, 0x54,
	0xba, 0x85, 0x4f, 0xdd, 0xea, 0x53, 0xe1, 0x5b, 0xf0, 0x30, 0x22, 0xec, 0x21, 0x9e, 0xcf, 0x46,
	0x22, 0x30, 0xf8, 0x50, 0xc4, 0x2d, 0x3b, 0x14, 0x22, 0xd4, 0x16, 0x63, 0xf2, 0x99, 0x50, 0x1c,
	0x27, 0x2e, 0xd8, 0x56, 0xb4, 0x0a, 0x46, 0xa9, 0x33, 0x68, 0x83, 0x66, 0x26, 0xec, 0xa6, 0x5f,
	0x26, 0x2a, 0xe7, 0x73, 0xbb, 0x45, 0x8c, 0xf5, 0x46, 0x9a, 0xbd, 0x6c, 0x90, 0xb9, 0x5b, 0x62,
	0x59, 0xe8, 0x68, 0x97, 0x44, 0x55, 0x73, 0x24, 0xaa, 0x05, 0x47, 0xe0, 0x31, 0x97, 0x64, 0xb7,
	0x06, 0x08, 0x7a, 0x63, 0x45, 0xcf, 0xe8, 0x0e, 0x6f, 0x6a, 0xfa, 0x18, 0xba, 0xe0, 0xe9, 0x56,
	0xec, 0xea, 0xfe, 0xb6, 0x8d, 0xee, 0xed, 0x34, 0x74, 0x8b, 0xbd, 0x9e, 0xae, 0x64, 0x5a, 0xae,
	0x5f, 0xd1, 0xb3, 0xf7, 0xd3, 0x39, 0x0e, 0xaa, 0xec, 0xa1, 0xd3, 0xb4, 0x7c, 0x9d, 0xe6, 0xb4,
	0x5e, 0x81, 0x52, 0xfc, 0x61, 0x42, 0x04, 0x24, 0x2c, 0x17, 0xea, 0xc7, 0x3f, 0x7e, 0x36, 0xef,
	0x0f, 0x22, 0x95, 0x50, 0x4a, 0x82, 0x21, 0x9b, 0x53, 0x0c, 0x0f, 0x80, 0xa4, 0xca, 0x52, 0x5d,
	0xfe, 0xbb, 0x50, 0x2a, 0x83, 0xa8, 0x27, 0x26, 0x98, 0x97, 0xb4, 0xbe, 0x02, 0x39, 0x39, 0xc4,
	0xf0, 0x14, 0x6c, 0x7c, 0x36, 0x8c, 0xd8, 0x60, 0xc4, 0x47, 0xeb, 0xc4, 0x9b, 0x0a, 0x0c, 0xe5,
	0x8d, 0x38, 0xd7, 0x9c, 0x7c, 0xff, 0xa5, 0x14, 0x7e, 0x2f, 0x94, 0x02, 0xd3, 0xd5, 0x96, 0xba,
	0x4d, 0x49, 0xeb, 0x0b, 0xa8, 0xc5, 0xb7, 0xfe, 0xc8, 0xfe, 0x31, 0xf9, 0x7a, 0x4e, 0x83, 0xa9,
	0x37, 0x21, 0x01, 0x7c, 0x06, 0xc0, 0xf9, 0x85, 0x7e, 0x36, 0x7c, 0xd3, 0xd7, 0x85, 0xbc, 0xc6,
	0xe4, 0x90, 0x27, 0x3e, 0x98, 0xb3, 0x90, 0x2c, 0x63, 0x18, 0x56, 0xd7, 0x72, 0xdc, 0x5a, 0xe6,
	0xc6, 0x3f, 0x0b, 0x45, 0xfa, 0x1f, 0x00, 0x00, 0xff, 0xff, 0xbe, 0x65, 0x55, 0xe7, 0xdb, 0x05,
	0x00, 0x00,
}
