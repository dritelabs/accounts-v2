// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.8
// source: public_keys_configuration.proto

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

type PublicKeysConfiguration int32

const (
	PublicKeysConfiguration_LOCAL  PublicKeysConfiguration = 0
	PublicKeysConfiguration_REMOTE PublicKeysConfiguration = 1
)

// Enum value maps for PublicKeysConfiguration.
var (
	PublicKeysConfiguration_name = map[int32]string{
		0: "LOCAL",
		1: "REMOTE",
	}
	PublicKeysConfiguration_value = map[string]int32{
		"LOCAL":  0,
		"REMOTE": 1,
	}
)

func (x PublicKeysConfiguration) Enum() *PublicKeysConfiguration {
	p := new(PublicKeysConfiguration)
	*p = x
	return p
}

func (x PublicKeysConfiguration) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PublicKeysConfiguration) Descriptor() protoreflect.EnumDescriptor {
	return file_public_keys_configuration_proto_enumTypes[0].Descriptor()
}

func (PublicKeysConfiguration) Type() protoreflect.EnumType {
	return &file_public_keys_configuration_proto_enumTypes[0]
}

func (x PublicKeysConfiguration) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PublicKeysConfiguration.Descriptor instead.
func (PublicKeysConfiguration) EnumDescriptor() ([]byte, []int) {
	return file_public_keys_configuration_proto_rawDescGZIP(), []int{0}
}

var File_public_keys_configuration_proto protoreflect.FileDescriptor

var file_public_keys_configuration_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x6b, 0x65, 0x79, 0x73, 0x5f, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x02, 0x70, 0x62, 0x2a, 0x30, 0x0a, 0x17, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b,
	0x65, 0x79, 0x73, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x09, 0x0a, 0x05, 0x4c, 0x4f, 0x43, 0x41, 0x4c, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x52,
	0x45, 0x4d, 0x4f, 0x54, 0x45, 0x10, 0x01, 0x42, 0x22, 0x5a, 0x20, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x72, 0x69, 0x74, 0x65, 0x6c, 0x61, 0x62, 0x73, 0x2f,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_public_keys_configuration_proto_rawDescOnce sync.Once
	file_public_keys_configuration_proto_rawDescData = file_public_keys_configuration_proto_rawDesc
)

func file_public_keys_configuration_proto_rawDescGZIP() []byte {
	file_public_keys_configuration_proto_rawDescOnce.Do(func() {
		file_public_keys_configuration_proto_rawDescData = protoimpl.X.CompressGZIP(file_public_keys_configuration_proto_rawDescData)
	})
	return file_public_keys_configuration_proto_rawDescData
}

var file_public_keys_configuration_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_public_keys_configuration_proto_goTypes = []interface{}{
	(PublicKeysConfiguration)(0), // 0: pb.PublicKeysConfiguration
}
var file_public_keys_configuration_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_public_keys_configuration_proto_init() }
func file_public_keys_configuration_proto_init() {
	if File_public_keys_configuration_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_public_keys_configuration_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_public_keys_configuration_proto_goTypes,
		DependencyIndexes: file_public_keys_configuration_proto_depIdxs,
		EnumInfos:         file_public_keys_configuration_proto_enumTypes,
	}.Build()
	File_public_keys_configuration_proto = out.File
	file_public_keys_configuration_proto_rawDesc = nil
	file_public_keys_configuration_proto_goTypes = nil
	file_public_keys_configuration_proto_depIdxs = nil
}