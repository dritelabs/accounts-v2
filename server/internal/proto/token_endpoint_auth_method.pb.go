// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.8
// source: token_endpoint_auth_method.proto

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

type TokenEndpointAuthMethod int32

const (
	TokenEndpointAuthMethod_CLIENT_SECRET_BASIC TokenEndpointAuthMethod = 0
	TokenEndpointAuthMethod_PRIVATE_KEY_JWT     TokenEndpointAuthMethod = 1
	TokenEndpointAuthMethod_NONE                TokenEndpointAuthMethod = 2
)

// Enum value maps for TokenEndpointAuthMethod.
var (
	TokenEndpointAuthMethod_name = map[int32]string{
		0: "CLIENT_SECRET_BASIC",
		1: "PRIVATE_KEY_JWT",
		2: "NONE",
	}
	TokenEndpointAuthMethod_value = map[string]int32{
		"CLIENT_SECRET_BASIC": 0,
		"PRIVATE_KEY_JWT":     1,
		"NONE":                2,
	}
)

func (x TokenEndpointAuthMethod) Enum() *TokenEndpointAuthMethod {
	p := new(TokenEndpointAuthMethod)
	*p = x
	return p
}

func (x TokenEndpointAuthMethod) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TokenEndpointAuthMethod) Descriptor() protoreflect.EnumDescriptor {
	return file_token_endpoint_auth_method_proto_enumTypes[0].Descriptor()
}

func (TokenEndpointAuthMethod) Type() protoreflect.EnumType {
	return &file_token_endpoint_auth_method_proto_enumTypes[0]
}

func (x TokenEndpointAuthMethod) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TokenEndpointAuthMethod.Descriptor instead.
func (TokenEndpointAuthMethod) EnumDescriptor() ([]byte, []int) {
	return file_token_endpoint_auth_method_proto_rawDescGZIP(), []int{0}
}

var File_token_endpoint_auth_method_proto protoreflect.FileDescriptor

var file_token_endpoint_auth_method_proto_rawDesc = []byte{
	0x0a, 0x20, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x5f, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x2a, 0x51, 0x0a, 0x17, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x45,
	0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x41, 0x75, 0x74, 0x68, 0x4d, 0x65, 0x74, 0x68, 0x6f,
	0x64, 0x12, 0x17, 0x0a, 0x13, 0x43, 0x4c, 0x49, 0x45, 0x4e, 0x54, 0x5f, 0x53, 0x45, 0x43, 0x52,
	0x45, 0x54, 0x5f, 0x42, 0x41, 0x53, 0x49, 0x43, 0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f, 0x50, 0x52,
	0x49, 0x56, 0x41, 0x54, 0x45, 0x5f, 0x4b, 0x45, 0x59, 0x5f, 0x4a, 0x57, 0x54, 0x10, 0x01, 0x12,
	0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x02, 0x42, 0x22, 0x5a, 0x20, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x72, 0x69, 0x74, 0x65, 0x6c, 0x61, 0x62,
	0x73, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_token_endpoint_auth_method_proto_rawDescOnce sync.Once
	file_token_endpoint_auth_method_proto_rawDescData = file_token_endpoint_auth_method_proto_rawDesc
)

func file_token_endpoint_auth_method_proto_rawDescGZIP() []byte {
	file_token_endpoint_auth_method_proto_rawDescOnce.Do(func() {
		file_token_endpoint_auth_method_proto_rawDescData = protoimpl.X.CompressGZIP(file_token_endpoint_auth_method_proto_rawDescData)
	})
	return file_token_endpoint_auth_method_proto_rawDescData
}

var file_token_endpoint_auth_method_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_token_endpoint_auth_method_proto_goTypes = []interface{}{
	(TokenEndpointAuthMethod)(0), // 0: pb.TokenEndpointAuthMethod
}
var file_token_endpoint_auth_method_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_token_endpoint_auth_method_proto_init() }
func file_token_endpoint_auth_method_proto_init() {
	if File_token_endpoint_auth_method_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_token_endpoint_auth_method_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_token_endpoint_auth_method_proto_goTypes,
		DependencyIndexes: file_token_endpoint_auth_method_proto_depIdxs,
		EnumInfos:         file_token_endpoint_auth_method_proto_enumTypes,
	}.Build()
	File_token_endpoint_auth_method_proto = out.File
	file_token_endpoint_auth_method_proto_rawDesc = nil
	file_token_endpoint_auth_method_proto_goTypes = nil
	file_token_endpoint_auth_method_proto_depIdxs = nil
}
