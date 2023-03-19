// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: drite/account/v1/jwk.proto

package accountv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Jwk struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ClientId  string                 `protobuf:"bytes,3,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	Jwk       *anypb.Any             `protobuf:"bytes,8,opt,name=jwk,proto3" json:"jwk,omitempty"`
	Scopes    []string               `protobuf:"bytes,4,rep,name=scopes,proto3" json:"scopes,omitempty"`
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *Jwk) Reset() {
	*x = Jwk{}
	if protoimpl.UnsafeEnabled {
		mi := &file_drite_account_v1_jwk_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Jwk) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Jwk) ProtoMessage() {}

func (x *Jwk) ProtoReflect() protoreflect.Message {
	mi := &file_drite_account_v1_jwk_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Jwk.ProtoReflect.Descriptor instead.
func (*Jwk) Descriptor() ([]byte, []int) {
	return file_drite_account_v1_jwk_proto_rawDescGZIP(), []int{0}
}

func (x *Jwk) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Jwk) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (x *Jwk) GetJwk() *anypb.Any {
	if x != nil {
		return x.Jwk
	}
	return nil
}

func (x *Jwk) GetScopes() []string {
	if x != nil {
		return x.Scopes
	}
	return nil
}

func (x *Jwk) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Jwk) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type GetJwksRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetJwksRequest) Reset() {
	*x = GetJwksRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_drite_account_v1_jwk_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetJwksRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetJwksRequest) ProtoMessage() {}

func (x *GetJwksRequest) ProtoReflect() protoreflect.Message {
	mi := &file_drite_account_v1_jwk_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetJwksRequest.ProtoReflect.Descriptor instead.
func (*GetJwksRequest) Descriptor() ([]byte, []int) {
	return file_drite_account_v1_jwk_proto_rawDescGZIP(), []int{1}
}

type GetJwksResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Jwks *anypb.Any `protobuf:"bytes,1,opt,name=jwks,proto3" json:"jwks,omitempty"`
}

func (x *GetJwksResponse) Reset() {
	*x = GetJwksResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_drite_account_v1_jwk_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetJwksResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetJwksResponse) ProtoMessage() {}

func (x *GetJwksResponse) ProtoReflect() protoreflect.Message {
	mi := &file_drite_account_v1_jwk_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetJwksResponse.ProtoReflect.Descriptor instead.
func (*GetJwksResponse) Descriptor() ([]byte, []int) {
	return file_drite_account_v1_jwk_proto_rawDescGZIP(), []int{2}
}

func (x *GetJwksResponse) GetJwks() *anypb.Any {
	if x != nil {
		return x.Jwks
	}
	return nil
}

var File_drite_account_v1_jwk_proto protoreflect.FileDescriptor

var file_drite_account_v1_jwk_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x64, 0x72, 0x69, 0x74, 0x65, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f,
	0x76, 0x31, 0x2f, 0x6a, 0x77, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x64, 0x72,
	0x69, 0x74, 0x65, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x1a, 0x19,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe8, 0x01, 0x0a, 0x03, 0x4a,
	0x77, 0x6b, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12,
	0x26, 0x0a, 0x03, 0x6a, 0x77, 0x6b, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41,
	0x6e, 0x79, 0x52, 0x03, 0x6a, 0x77, 0x6b, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x63, 0x6f, 0x70, 0x65,
	0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x73, 0x12,
	0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x10, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x4a, 0x77, 0x6b, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x3b, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x4a, 0x77,
	0x6b, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x04, 0x6a, 0x77,
	0x6b, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x04,
	0x6a, 0x77, 0x6b, 0x73, 0x42, 0xbf, 0x01, 0x0a, 0x14, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x72, 0x69,
	0x74, 0x65, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x42, 0x08, 0x4a,
	0x77, 0x6b, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3b, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x72, 0x69, 0x74, 0x65, 0x6c, 0x61, 0x62, 0x73, 0x2f,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x2f, 0x70, 0x62, 0x2f, 0x64, 0x72, 0x69, 0x74,
	0x65, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x3b, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x44, 0x41, 0x58, 0xaa, 0x02, 0x10, 0x44,
	0x72, 0x69, 0x74, 0x65, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x56, 0x31, 0xca,
	0x02, 0x10, 0x44, 0x72, 0x69, 0x74, 0x65, 0x5c, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5c,
	0x56, 0x31, 0xe2, 0x02, 0x1c, 0x44, 0x72, 0x69, 0x74, 0x65, 0x5c, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0xea, 0x02, 0x12, 0x44, 0x72, 0x69, 0x74, 0x65, 0x3a, 0x3a, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_drite_account_v1_jwk_proto_rawDescOnce sync.Once
	file_drite_account_v1_jwk_proto_rawDescData = file_drite_account_v1_jwk_proto_rawDesc
)

func file_drite_account_v1_jwk_proto_rawDescGZIP() []byte {
	file_drite_account_v1_jwk_proto_rawDescOnce.Do(func() {
		file_drite_account_v1_jwk_proto_rawDescData = protoimpl.X.CompressGZIP(file_drite_account_v1_jwk_proto_rawDescData)
	})
	return file_drite_account_v1_jwk_proto_rawDescData
}

var file_drite_account_v1_jwk_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_drite_account_v1_jwk_proto_goTypes = []interface{}{
	(*Jwk)(nil),                   // 0: drite.account.v1.Jwk
	(*GetJwksRequest)(nil),        // 1: drite.account.v1.GetJwksRequest
	(*GetJwksResponse)(nil),       // 2: drite.account.v1.GetJwksResponse
	(*anypb.Any)(nil),             // 3: google.protobuf.Any
	(*timestamppb.Timestamp)(nil), // 4: google.protobuf.Timestamp
}
var file_drite_account_v1_jwk_proto_depIdxs = []int32{
	3, // 0: drite.account.v1.Jwk.jwk:type_name -> google.protobuf.Any
	4, // 1: drite.account.v1.Jwk.created_at:type_name -> google.protobuf.Timestamp
	4, // 2: drite.account.v1.Jwk.updated_at:type_name -> google.protobuf.Timestamp
	3, // 3: drite.account.v1.GetJwksResponse.jwks:type_name -> google.protobuf.Any
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_drite_account_v1_jwk_proto_init() }
func file_drite_account_v1_jwk_proto_init() {
	if File_drite_account_v1_jwk_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_drite_account_v1_jwk_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Jwk); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_drite_account_v1_jwk_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetJwksRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_drite_account_v1_jwk_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetJwksResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_drite_account_v1_jwk_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_drite_account_v1_jwk_proto_goTypes,
		DependencyIndexes: file_drite_account_v1_jwk_proto_depIdxs,
		MessageInfos:      file_drite_account_v1_jwk_proto_msgTypes,
	}.Build()
	File_drite_account_v1_jwk_proto = out.File
	file_drite_account_v1_jwk_proto_rawDesc = nil
	file_drite_account_v1_jwk_proto_goTypes = nil
	file_drite_account_v1_jwk_proto_depIdxs = nil
}
