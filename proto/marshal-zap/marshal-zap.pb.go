// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/marshal-zap/marshal-zap.proto

package marshal_zap

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	descriptor "github.com/golang/protobuf/protoc-gen-go/descriptor"
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

var E_Mask = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         50001,
	Name:          "github.com.kei2100.playground_protobuf.marshal_zap.mask",
	Tag:           "varint,50001,opt,name=mask",
	Filename:      "proto/marshal-zap/marshal-zap.proto",
}

func init() {
	proto.RegisterExtension(E_Mask)
}

func init() {
	proto.RegisterFile("proto/marshal-zap/marshal-zap.proto", fileDescriptor_ed0eae37d6ff6637)
}

var fileDescriptor_ed0eae37d6ff6637 = []byte{
	// 178 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x2e, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0xcf, 0x4d, 0x2c, 0x2a, 0xce, 0x48, 0xcc, 0xd1, 0xad, 0x4a, 0x2c, 0x40, 0x66, 0xeb,
	0x81, 0x65, 0x85, 0x8c, 0xd2, 0x33, 0x4b, 0x32, 0x4a, 0x93, 0xf4, 0x92, 0xf3, 0x73, 0xf5, 0xb2,
	0x53, 0x33, 0x8d, 0x0c, 0x0d, 0x0c, 0xf4, 0x0a, 0x72, 0x12, 0x2b, 0xd3, 0x8b, 0xf2, 0x4b, 0xf3,
	0x52, 0xe2, 0xc1, 0x8a, 0x92, 0x4a, 0xd3, 0xf4, 0xa0, 0x3a, 0xe3, 0xab, 0x12, 0x0b, 0xa4, 0x14,
	0xd2, 0xf3, 0xf3, 0xd3, 0x73, 0x52, 0xf5, 0x61, 0x92, 0xfa, 0x29, 0xa9, 0xc5, 0xc9, 0x45, 0x99,
	0x05, 0x25, 0xf9, 0x45, 0x10, 0x53, 0xad, 0x8c, 0xb9, 0x58, 0x72, 0x13, 0x8b, 0xb3, 0x85, 0x64,
	0xf5, 0x20, 0x4a, 0xf5, 0xe0, 0xe6, 0xb8, 0x65, 0xa6, 0xe6, 0xa4, 0xf8, 0x17, 0x94, 0x64, 0xe6,
	0xe7, 0x15, 0x4b, 0x5c, 0x6c, 0x63, 0x56, 0x60, 0xd4, 0xe0, 0x08, 0x02, 0x2b, 0x76, 0xb2, 0x8a,
	0xb2, 0x40, 0x38, 0x46, 0x1f, 0xea, 0x18, 0x7d, 0x84, 0x63, 0x74, 0xe1, 0xf6, 0x61, 0x78, 0x2c,
	0x89, 0x0d, 0x2c, 0x64, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x8e, 0x46, 0xc2, 0x28, 0xf4, 0x00,
	0x00, 0x00,
}