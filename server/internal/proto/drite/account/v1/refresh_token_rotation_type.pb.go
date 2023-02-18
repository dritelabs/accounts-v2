// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: drite/account/v1/refresh_token_rotation_type.proto

package accountv1

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

type RefreshTokenRotationType int32

const (
	RefreshTokenRotationType_REFRESH_TOKEN_ROTATION_TYPE_UNSPECIFIED RefreshTokenRotationType = 0
	RefreshTokenRotationType_REFRESH_TOKEN_ROTATION_TYPE_ROTATE      RefreshTokenRotationType = 1
	RefreshTokenRotationType_REFRESH_TOKEN_ROTATION_TYPE_STATIC      RefreshTokenRotationType = 2
)

// Enum value maps for RefreshTokenRotationType.
var (
	RefreshTokenRotationType_name = map[int32]string{
		0: "REFRESH_TOKEN_ROTATION_TYPE_UNSPECIFIED",
		1: "REFRESH_TOKEN_ROTATION_TYPE_ROTATE",
		2: "REFRESH_TOKEN_ROTATION_TYPE_STATIC",
	}
	RefreshTokenRotationType_value = map[string]int32{
		"REFRESH_TOKEN_ROTATION_TYPE_UNSPECIFIED": 0,
		"REFRESH_TOKEN_ROTATION_TYPE_ROTATE":      1,
		"REFRESH_TOKEN_ROTATION_TYPE_STATIC":      2,
	}
)

func (x RefreshTokenRotationType) Enum() *RefreshTokenRotationType {
	p := new(RefreshTokenRotationType)
	*p = x
	return p
}

func (x RefreshTokenRotationType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RefreshTokenRotationType) Descriptor() protoreflect.EnumDescriptor {
	return file_drite_account_v1_refresh_token_rotation_type_proto_enumTypes[0].Descriptor()
}

func (RefreshTokenRotationType) Type() protoreflect.EnumType {
	return &file_drite_account_v1_refresh_token_rotation_type_proto_enumTypes[0]
}

func (x RefreshTokenRotationType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RefreshTokenRotationType.Descriptor instead.
func (RefreshTokenRotationType) EnumDescriptor() ([]byte, []int) {
	return file_drite_account_v1_refresh_token_rotation_type_proto_rawDescGZIP(), []int{0}
}

var File_drite_account_v1_refresh_token_rotation_type_proto protoreflect.FileDescriptor

var file_drite_account_v1_refresh_token_rotation_type_proto_rawDesc = []byte{
	0x0a, 0x32, 0x64, 0x72, 0x69, 0x74, 0x65, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f,
	0x76, 0x31, 0x2f, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x5f, 0x72, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x64, 0x72, 0x69, 0x74, 0x65, 0x2e, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2a, 0x97, 0x01, 0x0a, 0x18, 0x52, 0x65, 0x66, 0x72, 0x65,
	0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x2b, 0x0a, 0x27, 0x52, 0x45, 0x46, 0x52, 0x45, 0x53, 0x48, 0x5f, 0x54,
	0x4f, 0x4b, 0x45, 0x4e, 0x5f, 0x52, 0x4f, 0x54, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00,
	0x12, 0x26, 0x0a, 0x22, 0x52, 0x45, 0x46, 0x52, 0x45, 0x53, 0x48, 0x5f, 0x54, 0x4f, 0x4b, 0x45,
	0x4e, 0x5f, 0x52, 0x4f, 0x54, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x52, 0x4f, 0x54, 0x41, 0x54, 0x45, 0x10, 0x01, 0x12, 0x26, 0x0a, 0x22, 0x52, 0x45, 0x46, 0x52,
	0x45, 0x53, 0x48, 0x5f, 0x54, 0x4f, 0x4b, 0x45, 0x4e, 0x5f, 0x52, 0x4f, 0x54, 0x41, 0x54, 0x49,
	0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x49, 0x43, 0x10, 0x02,
	0x42, 0xd4, 0x01, 0x0a, 0x14, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x72, 0x69, 0x74, 0x65, 0x2e, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x42, 0x1d, 0x52, 0x65, 0x66, 0x72, 0x65,
	0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54,
	0x79, 0x70, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3b, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x72, 0x69, 0x74, 0x65, 0x6c, 0x61, 0x62, 0x73,
	0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x2f, 0x70, 0x62, 0x2f, 0x64, 0x72, 0x69,
	0x74, 0x65, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x3b, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x44, 0x41, 0x58, 0xaa, 0x02, 0x10,
	0x44, 0x72, 0x69, 0x74, 0x65, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x56, 0x31,
	0xca, 0x02, 0x10, 0x44, 0x72, 0x69, 0x74, 0x65, 0x5c, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x5c, 0x56, 0x31, 0xe2, 0x02, 0x1c, 0x44, 0x72, 0x69, 0x74, 0x65, 0x5c, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0xea, 0x02, 0x12, 0x44, 0x72, 0x69, 0x74, 0x65, 0x3a, 0x3a, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_drite_account_v1_refresh_token_rotation_type_proto_rawDescOnce sync.Once
	file_drite_account_v1_refresh_token_rotation_type_proto_rawDescData = file_drite_account_v1_refresh_token_rotation_type_proto_rawDesc
)

func file_drite_account_v1_refresh_token_rotation_type_proto_rawDescGZIP() []byte {
	file_drite_account_v1_refresh_token_rotation_type_proto_rawDescOnce.Do(func() {
		file_drite_account_v1_refresh_token_rotation_type_proto_rawDescData = protoimpl.X.CompressGZIP(file_drite_account_v1_refresh_token_rotation_type_proto_rawDescData)
	})
	return file_drite_account_v1_refresh_token_rotation_type_proto_rawDescData
}

var file_drite_account_v1_refresh_token_rotation_type_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_drite_account_v1_refresh_token_rotation_type_proto_goTypes = []interface{}{
	(RefreshTokenRotationType)(0), // 0: drite.account.v1.RefreshTokenRotationType
}
var file_drite_account_v1_refresh_token_rotation_type_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_drite_account_v1_refresh_token_rotation_type_proto_init() }
func file_drite_account_v1_refresh_token_rotation_type_proto_init() {
	if File_drite_account_v1_refresh_token_rotation_type_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_drite_account_v1_refresh_token_rotation_type_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_drite_account_v1_refresh_token_rotation_type_proto_goTypes,
		DependencyIndexes: file_drite_account_v1_refresh_token_rotation_type_proto_depIdxs,
		EnumInfos:         file_drite_account_v1_refresh_token_rotation_type_proto_enumTypes,
	}.Build()
	File_drite_account_v1_refresh_token_rotation_type_proto = out.File
	file_drite_account_v1_refresh_token_rotation_type_proto_rawDesc = nil
	file_drite_account_v1_refresh_token_rotation_type_proto_goTypes = nil
	file_drite_account_v1_refresh_token_rotation_type_proto_depIdxs = nil
}
