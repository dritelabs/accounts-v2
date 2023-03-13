// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.29.0
// 	protoc        (unknown)
// source: drite/account/v1/client_approval.proto

package accountv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

type ClientApproval struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId    string                 `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	ClientId  string                 `protobuf:"bytes,3,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	Scopes    []string               `protobuf:"bytes,4,rep,name=scopes,proto3" json:"scopes,omitempty"`
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *ClientApproval) Reset() {
	*x = ClientApproval{}
	if protoimpl.UnsafeEnabled {
		mi := &file_drite_account_v1_client_approval_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientApproval) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientApproval) ProtoMessage() {}

func (x *ClientApproval) ProtoReflect() protoreflect.Message {
	mi := &file_drite_account_v1_client_approval_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientApproval.ProtoReflect.Descriptor instead.
func (*ClientApproval) Descriptor() ([]byte, []int) {
	return file_drite_account_v1_client_approval_proto_rawDescGZIP(), []int{0}
}

func (x *ClientApproval) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ClientApproval) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *ClientApproval) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (x *ClientApproval) GetScopes() []string {
	if x != nil {
		return x.Scopes
	}
	return nil
}

func (x *ClientApproval) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *ClientApproval) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

var File_drite_account_v1_client_approval_proto protoreflect.FileDescriptor

var file_drite_account_v1_client_approval_proto_rawDesc = []byte{
	0x0a, 0x26, 0x64, 0x72, 0x69, 0x74, 0x65, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f,
	0x76, 0x31, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x76,
	0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x64, 0x72, 0x69, 0x74, 0x65, 0x2e,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe4, 0x01, 0x0a, 0x0e,
	0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x17,
	0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x73, 0x18, 0x04,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x73, 0x12, 0x39, 0x0a, 0x0a,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x42, 0xca, 0x01, 0x0a, 0x14, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x72, 0x69, 0x74, 0x65,
	0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x42, 0x13, 0x43, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x3b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64,
	0x72, 0x69, 0x74, 0x65, 0x6c, 0x61, 0x62, 0x73, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x73, 0x2f, 0x70, 0x62, 0x2f, 0x64, 0x72, 0x69, 0x74, 0x65, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x2f, 0x76, 0x31, 0x3b, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x76, 0x31, 0xa2,
	0x02, 0x03, 0x44, 0x41, 0x58, 0xaa, 0x02, 0x10, 0x44, 0x72, 0x69, 0x74, 0x65, 0x2e, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x10, 0x44, 0x72, 0x69, 0x74, 0x65,
	0x5c, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x1c, 0x44, 0x72,
	0x69, 0x74, 0x65, 0x5c, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5c, 0x56, 0x31, 0x5c, 0x47,
	0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x12, 0x44, 0x72, 0x69,
	0x74, 0x65, 0x3a, 0x3a, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x3a, 0x3a, 0x56, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_drite_account_v1_client_approval_proto_rawDescOnce sync.Once
	file_drite_account_v1_client_approval_proto_rawDescData = file_drite_account_v1_client_approval_proto_rawDesc
)

func file_drite_account_v1_client_approval_proto_rawDescGZIP() []byte {
	file_drite_account_v1_client_approval_proto_rawDescOnce.Do(func() {
		file_drite_account_v1_client_approval_proto_rawDescData = protoimpl.X.CompressGZIP(file_drite_account_v1_client_approval_proto_rawDescData)
	})
	return file_drite_account_v1_client_approval_proto_rawDescData
}

var file_drite_account_v1_client_approval_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_drite_account_v1_client_approval_proto_goTypes = []interface{}{
	(*ClientApproval)(nil),        // 0: drite.account.v1.ClientApproval
	(*timestamppb.Timestamp)(nil), // 1: google.protobuf.Timestamp
}
var file_drite_account_v1_client_approval_proto_depIdxs = []int32{
	1, // 0: drite.account.v1.ClientApproval.created_at:type_name -> google.protobuf.Timestamp
	1, // 1: drite.account.v1.ClientApproval.updated_at:type_name -> google.protobuf.Timestamp
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_drite_account_v1_client_approval_proto_init() }
func file_drite_account_v1_client_approval_proto_init() {
	if File_drite_account_v1_client_approval_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_drite_account_v1_client_approval_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientApproval); i {
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
			RawDescriptor: file_drite_account_v1_client_approval_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_drite_account_v1_client_approval_proto_goTypes,
		DependencyIndexes: file_drite_account_v1_client_approval_proto_depIdxs,
		MessageInfos:      file_drite_account_v1_client_approval_proto_msgTypes,
	}.Build()
	File_drite_account_v1_client_approval_proto = out.File
	file_drite_account_v1_client_approval_proto_rawDesc = nil
	file_drite_account_v1_client_approval_proto_goTypes = nil
	file_drite_account_v1_client_approval_proto_depIdxs = nil
}
