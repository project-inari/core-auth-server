// Version

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v4.25.2
// source: protobuf/authPb/auth.proto

package authPb

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

// Structure
type VerifyTokenReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *VerifyTokenReq) Reset() {
	*x = VerifyTokenReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_authPb_auth_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VerifyTokenReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyTokenReq) ProtoMessage() {}

func (x *VerifyTokenReq) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_authPb_auth_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifyTokenReq.ProtoReflect.Descriptor instead.
func (*VerifyTokenReq) Descriptor() ([]byte, []int) {
	return file_protobuf_authPb_auth_proto_rawDescGZIP(), []int{0}
}

func (x *VerifyTokenReq) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type VerifyTokenRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Uid      string `protobuf:"bytes,2,opt,name=uid,proto3" json:"uid,omitempty"`
	Success  bool   `protobuf:"varint,3,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *VerifyTokenRes) Reset() {
	*x = VerifyTokenRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_authPb_auth_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VerifyTokenRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyTokenRes) ProtoMessage() {}

func (x *VerifyTokenRes) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_authPb_auth_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifyTokenRes.ProtoReflect.Descriptor instead.
func (*VerifyTokenRes) Descriptor() ([]byte, []int) {
	return file_protobuf_authPb_auth_proto_rawDescGZIP(), []int{1}
}

func (x *VerifyTokenRes) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *VerifyTokenRes) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *VerifyTokenRes) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_protobuf_authPb_auth_proto protoreflect.FileDescriptor

var file_protobuf_authPb_auth_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x50,
	0x62, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x26, 0x0a, 0x0e,
	0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x58, 0x0a, 0x0e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x75, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x32, 0x42,
	0x0a, 0x0f, 0x41, 0x75, 0x74, 0x68, 0x47, 0x72, 0x70, 0x63, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x2f, 0x0a, 0x0b, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x12, 0x0f, 0x2e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65,
	0x71, 0x1a, 0x0f, 0x2e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52,
	0x65, 0x73, 0x42, 0x3b, 0x5a, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2d, 0x69, 0x6e, 0x61, 0x72, 0x69, 0x2f, 0x63,
	0x6f, 0x72, 0x65, 0x2d, 0x61, 0x75, 0x74, 0x68, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x50, 0x62, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protobuf_authPb_auth_proto_rawDescOnce sync.Once
	file_protobuf_authPb_auth_proto_rawDescData = file_protobuf_authPb_auth_proto_rawDesc
)

func file_protobuf_authPb_auth_proto_rawDescGZIP() []byte {
	file_protobuf_authPb_auth_proto_rawDescOnce.Do(func() {
		file_protobuf_authPb_auth_proto_rawDescData = protoimpl.X.CompressGZIP(file_protobuf_authPb_auth_proto_rawDescData)
	})
	return file_protobuf_authPb_auth_proto_rawDescData
}

var file_protobuf_authPb_auth_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_protobuf_authPb_auth_proto_goTypes = []any{
	(*VerifyTokenReq)(nil), // 0: VerifyTokenReq
	(*VerifyTokenRes)(nil), // 1: VerifyTokenRes
}
var file_protobuf_authPb_auth_proto_depIdxs = []int32{
	0, // 0: AuthGrpcService.VerifyToken:input_type -> VerifyTokenReq
	1, // 1: AuthGrpcService.VerifyToken:output_type -> VerifyTokenRes
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_protobuf_authPb_auth_proto_init() }
func file_protobuf_authPb_auth_proto_init() {
	if File_protobuf_authPb_auth_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protobuf_authPb_auth_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*VerifyTokenReq); i {
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
		file_protobuf_authPb_auth_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*VerifyTokenRes); i {
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
			RawDescriptor: file_protobuf_authPb_auth_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protobuf_authPb_auth_proto_goTypes,
		DependencyIndexes: file_protobuf_authPb_auth_proto_depIdxs,
		MessageInfos:      file_protobuf_authPb_auth_proto_msgTypes,
	}.Build()
	File_protobuf_authPb_auth_proto = out.File
	file_protobuf_authPb_auth_proto_rawDesc = nil
	file_protobuf_authPb_auth_proto_goTypes = nil
	file_protobuf_authPb_auth_proto_depIdxs = nil
}
