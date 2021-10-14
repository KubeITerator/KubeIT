// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: grpc/common.proto

package api

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

type StatusReport struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Action  string `protobuf:"bytes,1,opt,name=action,proto3" json:"action,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *StatusReport) Reset() {
	*x = StatusReport{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatusReport) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatusReport) ProtoMessage() {}

func (x *StatusReport) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatusReport.ProtoReflect.Descriptor instead.
func (*StatusReport) Descriptor() ([]byte, []int) {
	return file_grpc_common_proto_rawDescGZIP(), []int{0}
}

func (x *StatusReport) GetAction() string {
	if x != nil {
		return x.Action
	}
	return ""
}

func (x *StatusReport) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type IDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *IDRequest) Reset() {
	*x = IDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_common_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IDRequest) ProtoMessage() {}

func (x *IDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_common_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IDRequest.ProtoReflect.Descriptor instead.
func (*IDRequest) Descriptor() ([]byte, []int) {
	return file_grpc_common_proto_rawDescGZIP(), []int{1}
}

func (x *IDRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_common_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_common_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_grpc_common_proto_rawDescGZIP(), []int{2}
}

var File_grpc_common_proto protoreflect.FileDescriptor

var file_grpc_common_proto_rawDesc = []byte{
	0x0a, 0x11, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x6b, 0x75, 0x62, 0x65, 0x69, 0x74, 0x2e, 0x76, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x31, 0x22, 0x40, 0x0a, 0x0c, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x70,
	0x6f, 0x72, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x1b, 0x0a, 0x09, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x2d, 0x5a, 0x2b, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4b, 0x75, 0x62, 0x65, 0x49, 0x54,
	0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x2f, 0x4b, 0x75, 0x62, 0x65, 0x49, 0x54, 0x2f, 0x70, 0x6b,
	0x67, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_grpc_common_proto_rawDescOnce sync.Once
	file_grpc_common_proto_rawDescData = file_grpc_common_proto_rawDesc
)

func file_grpc_common_proto_rawDescGZIP() []byte {
	file_grpc_common_proto_rawDescOnce.Do(func() {
		file_grpc_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_grpc_common_proto_rawDescData)
	})
	return file_grpc_common_proto_rawDescData
}

var file_grpc_common_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_grpc_common_proto_goTypes = []interface{}{
	(*StatusReport)(nil), // 0: kubeit.v1beta1.StatusReport
	(*IDRequest)(nil),    // 1: kubeit.v1beta1.IDRequest
	(*Empty)(nil),        // 2: kubeit.v1beta1.Empty
}
var file_grpc_common_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_grpc_common_proto_init() }
func file_grpc_common_proto_init() {
	if File_grpc_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_grpc_common_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatusReport); i {
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
		file_grpc_common_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IDRequest); i {
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
		file_grpc_common_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
			RawDescriptor: file_grpc_common_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_grpc_common_proto_goTypes,
		DependencyIndexes: file_grpc_common_proto_depIdxs,
		MessageInfos:      file_grpc_common_proto_msgTypes,
	}.Build()
	File_grpc_common_proto = out.File
	file_grpc_common_proto_rawDesc = nil
	file_grpc_common_proto_goTypes = nil
	file_grpc_common_proto_depIdxs = nil
}
