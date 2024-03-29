// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sku/sku.proto

package ccloud

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type Sku int32

const (
	Sku_UNKNOWN          Sku = 0
	Sku_BASIC_LEGACY     Sku = 1
	Sku_BASIC            Sku = 2
	Sku_STANDARD         Sku = 3
	Sku_DEDICATED        Sku = 4
	Sku_DEDICATED_LEGACY Sku = 5
	Sku_ENTERPRISE       Sku = 6
)

var Sku_name = map[int32]string{
	0: "UNKNOWN",
	1: "BASIC_LEGACY",
	2: "BASIC",
	3: "STANDARD",
	4: "DEDICATED",
	5: "DEDICATED_LEGACY",
	6: "ENTERPRISE",
}

var Sku_value = map[string]int32{
	"UNKNOWN":          0,
	"BASIC_LEGACY":     1,
	"BASIC":            2,
	"STANDARD":         3,
	"DEDICATED":        4,
	"DEDICATED_LEGACY": 5,
	"ENTERPRISE":       6,
}

func (x Sku) String() string {
	return proto.EnumName(Sku_name, int32(x))
}

func (Sku) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f71b3be6f63ac651, []int{0}
}

func init() {
	proto.RegisterEnum("ccloud.sku.Sku", Sku_name, Sku_value)
}

func init() { proto.RegisterFile("sku/sku.proto", fileDescriptor_f71b3be6f63ac651) }

var fileDescriptor_f71b3be6f63ac651 = []byte{
	// 252 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x3c, 0x8f, 0xd1, 0x4a, 0xc3, 0x30,
	0x14, 0x86, 0xad, 0x73, 0xd3, 0x1d, 0x37, 0x09, 0xc1, 0xeb, 0x3d, 0x80, 0xd0, 0x94, 0x22, 0x78,
	0xe3, 0x55, 0xd6, 0x04, 0x29, 0x4a, 0x1c, 0x6d, 0x45, 0xf4, 0x46, 0xd6, 0x74, 0x9d, 0x25, 0xb5,
	0x19, 0x6d, 0xcf, 0xc0, 0x37, 0xf4, 0xd2, 0x47, 0x90, 0x3e, 0x89, 0xd8, 0xe2, 0x2e, 0xff, 0x0f,
	0xbe, 0xc3, 0x77, 0x60, 0xde, 0x18, 0xf4, 0x1a, 0x83, 0x6c, 0x57, 0xdb, 0xd6, 0x52, 0xd0, 0xba,
	0xb4, 0x98, 0xb1, 0xc6, 0xe0, 0x55, 0x0d, 0xa3, 0xd8, 0x20, 0x3d, 0x87, 0xd3, 0x27, 0x75, 0xaf,
	0x1e, 0x9f, 0x15, 0x39, 0xa2, 0x04, 0x66, 0x4b, 0x1e, 0x87, 0xc1, 0xdb, 0x83, 0xbc, 0xe3, 0xc1,
	0x0b, 0x71, 0xe8, 0x14, 0xc6, 0x3d, 0x21, 0xc7, 0x74, 0x06, 0x67, 0x71, 0xc2, 0x95, 0xe0, 0x91,
	0x20, 0x23, 0x3a, 0x87, 0xa9, 0x90, 0x22, 0x0c, 0x78, 0x22, 0x05, 0x39, 0xa1, 0x97, 0x40, 0x0e,
	0xf3, 0xdf, 0x1e, 0xd3, 0x0b, 0x00, 0xa9, 0x12, 0x19, 0xad, 0xa2, 0x30, 0x96, 0x64, 0xb2, 0xfc,
	0xfc, 0xea, 0x16, 0xce, 0x77, 0xb7, 0x70, 0x7e, 0xba, 0x85, 0x03, 0x7e, 0x61, 0x99, 0xb6, 0x55,
	0x5e, 0xe2, 0xa6, 0x6a, 0xd9, 0x50, 0xd6, 0x77, 0xa6, 0x98, 0x33, 0xb3, 0xce, 0xcd, 0xfa, 0x6f,
	0x66, 0xa8, 0x5b, 0xa6, 0x6d, 0xbd, 0x61, 0x7b, 0x7f, 0xe5, 0xbc, 0xde, 0x6c, 0x8b, 0xf6, 0x1d,
	0x53, 0xa6, 0xed, 0x87, 0x77, 0x90, 0x8b, 0x4a, 0x7b, 0xc3, 0x6b, 0x6e, 0x93, 0x19, 0x77, 0x6b,
	0xdd, 0xbd, 0xef, 0xee, 0x30, 0x2d, 0x0b, 0x7d, 0x3b, 0xf0, 0x74, 0xd2, 0x5f, 0xbe, 0xfe, 0x0d,
	0x00, 0x00, 0xff, 0xff, 0x18, 0xe7, 0xee, 0x96, 0x12, 0x01, 0x00, 0x00,
}
