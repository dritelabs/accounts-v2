// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: grant_type.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GrantType int32

const (
	GrantType_AUTHORIZATION_CODE GrantType = 0
	GrantType_CLIENT_CREDENTIALS GrantType = 1
	GrantType_REFRESH_TOKEN      GrantType = 2
)

// Enum value maps for GrantType.
var (
	GrantType_name = map[int32]string{
		0: "AUTHORIZATION_CODE",
		1: "CLIENT_CREDENTIALS",
		2: "REFRESH_TOKEN",
	}
	GrantType_value = map[string]int32{
		"AUTHORIZATION_CODE": 0,
		"CLIENT_CREDENTIALS": 1,
		"REFRESH_TOKEN":      2,
	}
)

func (x GrantType) Enum() *GrantType {
	p := new(GrantType)
	*p = x
	return p
}

func (x GrantType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GrantType) Descriptor() protoreflect.EnumDescriptor {
	return file_grant_type_proto_enumTypes[0].Descriptor()
}

func (GrantType) Type() protoreflect.EnumType {
	return &file_grant_type_proto_enumTypes[0]
}

func (x GrantType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GrantType.Descriptor instead.
func (GrantType) EnumDescriptor() ([]byte, []int) {
	return file_grant_type_proto_rawDescGZIP(), []int{0}
}

var File_grant_type_proto protoreflect.FileDescriptor

var file_grant_type_proto_rawDesc = []byte{
	0x0a, 0x10, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x2a, 0x4e, 0x0a, 0x09, 0x47, 0x72, 0x61, 0x6e, 0x74, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x12, 0x41, 0x55, 0x54, 0x48, 0x4f, 0x52, 0x49, 0x5a, 0x41,
	0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x10, 0x00, 0x12, 0x16, 0x0a, 0x12, 0x43,
	0x4c, 0x49, 0x45, 0x4e, 0x54, 0x5f, 0x43, 0x52, 0x45, 0x44, 0x45, 0x4e, 0x54, 0x49, 0x41, 0x4c,
	0x53, 0x10, 0x01, 0x12, 0x11, 0x0a, 0x0d, 0x52, 0x45, 0x46, 0x52, 0x45, 0x53, 0x48, 0x5f, 0x54,
	0x4f, 0x4b, 0x45, 0x4e, 0x10, 0x02, 0x42, 0x62, 0x0a, 0x06, 0x63, 0x6f, 0x6d, 0x2e, 0x70, 0x62,
	0x42, 0x0e, 0x47, 0x72, 0x61, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x20, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64,
	0x72, 0x69, 0x74, 0x65, 0x6c, 0x61, 0x62, 0x73, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x73, 0x2f, 0x70, 0x62, 0xa2, 0x02, 0x03, 0x50, 0x58, 0x58, 0xaa, 0x02, 0x02, 0x50, 0x62, 0xca,
	0x02, 0x02, 0x50, 0x62, 0xe2, 0x02, 0x0e, 0x50, 0x62, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x02, 0x50, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_grant_type_proto_rawDescOnce sync.Once
	file_grant_type_proto_rawDescData = file_grant_type_proto_rawDesc
)

func file_grant_type_proto_rawDescGZIP() []byte {
	file_grant_type_proto_rawDescOnce.Do(func() {
		file_grant_type_proto_rawDescData = protoimpl.X.CompressGZIP(file_grant_type_proto_rawDescData)
	})
	return file_grant_type_proto_rawDescData
}

var file_grant_type_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_grant_type_proto_goTypes = []interface{}{
	(GrantType)(0), // 0: pb.GrantType
}
var file_grant_type_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_grant_type_proto_init() }
func file_grant_type_proto_init() {
	if File_grant_type_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_grant_type_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_grant_type_proto_goTypes,
		DependencyIndexes: file_grant_type_proto_depIdxs,
		EnumInfos:         file_grant_type_proto_enumTypes,
	}.Build()
	File_grant_type_proto = out.File
	file_grant_type_proto_rawDesc = nil
	file_grant_type_proto_goTypes = nil
	file_grant_type_proto_depIdxs = nil
}
