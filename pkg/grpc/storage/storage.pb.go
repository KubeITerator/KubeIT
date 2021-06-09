// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.16.0
// source: pkg/grpc/storage/storage.proto

package storage

import (
	_ "github.com/srikrsna/protoc-gen-gotag/tagger"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	common "kubeIT/pkg/grpc/common"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Status int32

const (
	Status_INIT      Status = 0
	Status_AVAILABLE Status = 1
	Status_ERROR     Status = 2
)

// Enum value maps for Status.
var (
	Status_name = map[int32]string{
		0: "INIT",
		1: "AVAILABLE",
		2: "ERROR",
	}
	Status_value = map[string]int32{
		"INIT":      0,
		"AVAILABLE": 1,
		"ERROR":     2,
	}
)

func (x Status) Enum() *Status {
	p := new(Status)
	*p = x
	return p
}

func (x Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Status) Descriptor() protoreflect.EnumDescriptor {
	return file_pkg_grpc_storage_storage_proto_enumTypes[0].Descriptor()
}

func (Status) Type() protoreflect.EnumType {
	return &file_pkg_grpc_storage_storage_proto_enumTypes[0]
}

func (x Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Status.Descriptor instead.
func (Status) EnumDescriptor() ([]byte, []int) {
	return file_pkg_grpc_storage_storage_proto_rawDescGZIP(), []int{0}
}

type StorageType int32

const (
	StorageType_S3     StorageType = 0
	StorageType_VOLUME StorageType = 1
	StorageType_CORE   StorageType = 2
)

// Enum value maps for StorageType.
var (
	StorageType_name = map[int32]string{
		0: "S3",
		1: "VOLUME",
		2: "CORE",
	}
	StorageType_value = map[string]int32{
		"S3":     0,
		"VOLUME": 1,
		"CORE":   2,
	}
)

func (x StorageType) Enum() *StorageType {
	p := new(StorageType)
	*p = x
	return p
}

func (x StorageType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (StorageType) Descriptor() protoreflect.EnumDescriptor {
	return file_pkg_grpc_storage_storage_proto_enumTypes[1].Descriptor()
}

func (StorageType) Type() protoreflect.EnumType {
	return &file_pkg_grpc_storage_storage_proto_enumTypes[1]
}

func (x StorageType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use StorageType.Descriptor instead.
func (StorageType) EnumDescriptor() ([]byte, []int) {
	return file_pkg_grpc_storage_storage_proto_rawDescGZIP(), []int{1}
}

type InitRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Userid      string       `protobuf:"bytes,1,opt,name=userid,proto3" json:"userid,omitempty"`
	Filename    string       `protobuf:"bytes,2,opt,name=filename,proto3" json:"filename,omitempty"`
	StorageInfo *StorageInfo `protobuf:"bytes,3,opt,name=storage_info,json=storageInfo,proto3" json:"storage_info,omitempty"`
}

func (x *InitRequest) Reset() {
	*x = InitRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_grpc_storage_storage_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InitRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InitRequest) ProtoMessage() {}

func (x *InitRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_grpc_storage_storage_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InitRequest.ProtoReflect.Descriptor instead.
func (*InitRequest) Descriptor() ([]byte, []int) {
	return file_pkg_grpc_storage_storage_proto_rawDescGZIP(), []int{0}
}

func (x *InitRequest) GetUserid() string {
	if x != nil {
		return x.Userid
	}
	return ""
}

func (x *InitRequest) GetFilename() string {
	if x != nil {
		return x.Filename
	}
	return ""
}

func (x *InitRequest) GetStorageInfo() *StorageInfo {
	if x != nil {
		return x.StorageInfo
	}
	return nil
}

type UrlResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TransferId string `protobuf:"bytes,1,opt,name=transfer_id,json=transferId,proto3" json:"transfer_id,omitempty"`
	Url        string `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *UrlResponse) Reset() {
	*x = UrlResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_grpc_storage_storage_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UrlResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UrlResponse) ProtoMessage() {}

func (x *UrlResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_grpc_storage_storage_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UrlResponse.ProtoReflect.Descriptor instead.
func (*UrlResponse) Descriptor() ([]byte, []int) {
	return file_pkg_grpc_storage_storage_proto_rawDescGZIP(), []int{1}
}

func (x *UrlResponse) GetTransferId() string {
	if x != nil {
		return x.TransferId
	}
	return ""
}

func (x *UrlResponse) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

type IDmessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TransferId string `protobuf:"bytes,1,opt,name=transfer_id,json=transferId,proto3" json:"transfer_id,omitempty"`
}

func (x *IDmessage) Reset() {
	*x = IDmessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_grpc_storage_storage_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IDmessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IDmessage) ProtoMessage() {}

func (x *IDmessage) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_grpc_storage_storage_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IDmessage.ProtoReflect.Descriptor instead.
func (*IDmessage) Descriptor() ([]byte, []int) {
	return file_pkg_grpc_storage_storage_proto_rawDescGZIP(), []int{2}
}

func (x *IDmessage) GetTransferId() string {
	if x != nil {
		return x.TransferId
	}
	return ""
}

type File struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string       `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" bson:"_id,omitempty"`
	UserId      string       `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Name        string       `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	StorageInfo *StorageInfo `protobuf:"bytes,4,opt,name=storage_info,json=storageInfo,proto3" json:"storage_info,omitempty"`
	Location    string       `protobuf:"bytes,5,opt,name=location,proto3" json:"location,omitempty"`
	Status      Status       `protobuf:"varint,6,opt,name=status,proto3,enum=v1alpha2.storage.Status" json:"status,omitempty"`
}

func (x *File) Reset() {
	*x = File{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_grpc_storage_storage_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *File) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*File) ProtoMessage() {}

func (x *File) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_grpc_storage_storage_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use File.ProtoReflect.Descriptor instead.
func (*File) Descriptor() ([]byte, []int) {
	return file_pkg_grpc_storage_storage_proto_rawDescGZIP(), []int{3}
}

func (x *File) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *File) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *File) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *File) GetStorageInfo() *StorageInfo {
	if x != nil {
		return x.StorageInfo
	}
	return nil
}

func (x *File) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

func (x *File) GetStatus() Status {
	if x != nil {
		return x.Status
	}
	return Status_INIT
}

type S3Info struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Multi  bool   `protobuf:"varint,1,opt,name=multi,proto3" json:"multi,omitempty"`
	Part   int32  `protobuf:"varint,2,opt,name=part,proto3" json:"part,omitempty"`
	Bucket string `protobuf:"bytes,3,opt,name=bucket,proto3" json:"bucket,omitempty"`
	Key    string `protobuf:"bytes,4,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *S3Info) Reset() {
	*x = S3Info{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_grpc_storage_storage_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *S3Info) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*S3Info) ProtoMessage() {}

func (x *S3Info) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_grpc_storage_storage_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use S3Info.ProtoReflect.Descriptor instead.
func (*S3Info) Descriptor() ([]byte, []int) {
	return file_pkg_grpc_storage_storage_proto_rawDescGZIP(), []int{4}
}

func (x *S3Info) GetMulti() bool {
	if x != nil {
		return x.Multi
	}
	return false
}

func (x *S3Info) GetPart() int32 {
	if x != nil {
		return x.Part
	}
	return 0
}

func (x *S3Info) GetBucket() string {
	if x != nil {
		return x.Bucket
	}
	return ""
}

func (x *S3Info) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type StorageInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StorageType StorageType `protobuf:"varint,1,opt,name=storage_type,json=storageType,proto3,enum=v1alpha2.storage.StorageType" json:"storage_type,omitempty"`
	// Types that are assignable to TypeInfo:
	//	*StorageInfo_S3Info
	TypeInfo isStorageInfo_TypeInfo `protobuf_oneof:"type_info"`
}

func (x *StorageInfo) Reset() {
	*x = StorageInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_grpc_storage_storage_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StorageInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StorageInfo) ProtoMessage() {}

func (x *StorageInfo) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_grpc_storage_storage_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StorageInfo.ProtoReflect.Descriptor instead.
func (*StorageInfo) Descriptor() ([]byte, []int) {
	return file_pkg_grpc_storage_storage_proto_rawDescGZIP(), []int{5}
}

func (x *StorageInfo) GetStorageType() StorageType {
	if x != nil {
		return x.StorageType
	}
	return StorageType_S3
}

func (m *StorageInfo) GetTypeInfo() isStorageInfo_TypeInfo {
	if m != nil {
		return m.TypeInfo
	}
	return nil
}

func (x *StorageInfo) GetS3Info() *S3Info {
	if x, ok := x.GetTypeInfo().(*StorageInfo_S3Info); ok {
		return x.S3Info
	}
	return nil
}

type isStorageInfo_TypeInfo interface {
	isStorageInfo_TypeInfo()
}

type StorageInfo_S3Info struct {
	S3Info *S3Info `protobuf:"bytes,2,opt,name=s3info,proto3,oneof"`
}

func (*StorageInfo_S3Info) isStorageInfo_TypeInfo() {}

var File_pkg_grpc_storage_storage_proto protoreflect.FileDescriptor

var file_pkg_grpc_storage_storage_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x70, 0x6b, 0x67, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x61,
	0x67, 0x65, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x10, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61,
	0x67, 0x65, 0x1a, 0x2d, 0x70, 0x6b, 0x67, 0x2f, 0x64, 0x65, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x6e,
	0x63, 0x69, 0x65, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1c, 0x70, 0x6b, 0x67, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x24, 0x70, 0x6b, 0x67, 0x2f, 0x64, 0x65, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x6e, 0x63, 0x69, 0x65,
	0x73, 0x2f, 0x74, 0x61, 0x67, 0x67, 0x65, 0x72, 0x2f, 0x74, 0x61, 0x67, 0x67, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x83, 0x01, 0x0a, 0x0b, 0x49, 0x6e, 0x69, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x69, 0x64, 0x12, 0x1a, 0x0a,
	0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x40, 0x0a, 0x0c, 0x73, 0x74, 0x6f,
	0x72, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1d, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61,
	0x67, 0x65, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0b,
	0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x40, 0x0a, 0x0b, 0x55,
	0x72, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x75,
	0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x22, 0x2c, 0x0a,
	0x09, 0x49, 0x44, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x49, 0x64, 0x22, 0xee, 0x01, 0x0a, 0x04,
	0x46, 0x69, 0x6c, 0x65, 0x12, 0x29, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x19, 0x9a, 0x84, 0x9e, 0x03, 0x14, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x5f, 0x69, 0x64,
	0x2c, 0x6f, 0x6d, 0x69, 0x74, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x40, 0x0a, 0x0c,
	0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x73, 0x74,
	0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x0b, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1a,
	0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x30, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x18, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x5c, 0x0a, 0x06,
	0x53, 0x33, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x12, 0x12, 0x0a, 0x04,
	0x70, 0x61, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x72, 0x74,
	0x12, 0x16, 0x0a, 0x06, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x22, 0x90, 0x01, 0x0a, 0x0b, 0x53,
	0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x40, 0x0a, 0x0c, 0x73, 0x74,
	0x6f, 0x72, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x1d, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x73, 0x74, 0x6f, 0x72,
	0x61, 0x67, 0x65, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x0b, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x32, 0x0a, 0x06,
	0x73, 0x33, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e,
	0x53, 0x33, 0x49, 0x6e, 0x66, 0x6f, 0x48, 0x00, 0x52, 0x06, 0x73, 0x33, 0x69, 0x6e, 0x66, 0x6f,
	0x42, 0x0b, 0x0a, 0x09, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x2a, 0x2c, 0x0a,
	0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x08, 0x0a, 0x04, 0x49, 0x4e, 0x49, 0x54, 0x10,
	0x00, 0x12, 0x0d, 0x0a, 0x09, 0x41, 0x56, 0x41, 0x49, 0x4c, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x01,
	0x12, 0x09, 0x0a, 0x05, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x02, 0x2a, 0x2b, 0x0a, 0x0b, 0x53,
	0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x06, 0x0a, 0x02, 0x53, 0x33,
	0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x56, 0x4f, 0x4c, 0x55, 0x4d, 0x45, 0x10, 0x01, 0x12, 0x08,
	0x0a, 0x04, 0x43, 0x4f, 0x52, 0x45, 0x10, 0x02, 0x32, 0xb2, 0x04, 0x0a, 0x18, 0x53, 0x74, 0x6f,
	0x72, 0x61, 0x67, 0x65, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5f, 0x0a, 0x0a, 0x49, 0x6e, 0x69, 0x74, 0x55, 0x70, 0x6c,
	0x6f, 0x61, 0x64, 0x12, 0x1d, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x73,
	0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x49, 0x6e, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x73, 0x74,
	0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x49, 0x44, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22,
	0x15, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0f, 0x22, 0x0a, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x70, 0x6c,
	0x6f, 0x61, 0x64, 0x3a, 0x01, 0x2a, 0x12, 0x6c, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x55, 0x70, 0x6c,
	0x6f, 0x61, 0x64, 0x55, 0x72, 0x6c, 0x12, 0x1b, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x32, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x49, 0x44, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x1a, 0x1d, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x73,
	0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x55, 0x72, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x20, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1a, 0x22, 0x18, 0x2f, 0x76, 0x31, 0x2f,
	0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x2f, 0x7b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x7d, 0x12, 0x6c, 0x0a, 0x0c, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x55, 0x70,
	0x6c, 0x6f, 0x61, 0x64, 0x12, 0x1b, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e,
	0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x49, 0x44, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x1a, 0x1d, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74,
	0x22, 0x20, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1a, 0x22, 0x18, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x70,
	0x6c, 0x6f, 0x61, 0x64, 0x2f, 0x7b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x7d, 0x12, 0x6c, 0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x70, 0x6c, 0x6f,
	0x61, 0x64, 0x12, 0x1b, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x73, 0x74,
	0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x49, 0x44, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a,
	0x1d, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x22, 0x20,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1a, 0x22, 0x18, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x70, 0x6c, 0x6f,
	0x61, 0x64, 0x2f, 0x7b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x7d,
	0x12, 0x6b, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x12,
	0x1b, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61,
	0x67, 0x65, 0x2e, 0x49, 0x44, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x1d, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e,
	0x55, 0x72, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x20, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x1a, 0x12, 0x18, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x2f,
	0x7b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x7d, 0x42, 0x19, 0x5a,
	0x17, 0x6b, 0x75, 0x62, 0x65, 0x49, 0x54, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x67, 0x72, 0x70, 0x63,
	0x2f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_grpc_storage_storage_proto_rawDescOnce sync.Once
	file_pkg_grpc_storage_storage_proto_rawDescData = file_pkg_grpc_storage_storage_proto_rawDesc
)

func file_pkg_grpc_storage_storage_proto_rawDescGZIP() []byte {
	file_pkg_grpc_storage_storage_proto_rawDescOnce.Do(func() {
		file_pkg_grpc_storage_storage_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_grpc_storage_storage_proto_rawDescData)
	})
	return file_pkg_grpc_storage_storage_proto_rawDescData
}

var file_pkg_grpc_storage_storage_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_pkg_grpc_storage_storage_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_pkg_grpc_storage_storage_proto_goTypes = []interface{}{
	(Status)(0),                 // 0: v1alpha2.storage.Status
	(StorageType)(0),            // 1: v1alpha2.storage.StorageType
	(*InitRequest)(nil),         // 2: v1alpha2.storage.InitRequest
	(*UrlResponse)(nil),         // 3: v1alpha2.storage.UrlResponse
	(*IDmessage)(nil),           // 4: v1alpha2.storage.IDmessage
	(*File)(nil),                // 5: v1alpha2.storage.File
	(*S3Info)(nil),              // 6: v1alpha2.storage.S3Info
	(*StorageInfo)(nil),         // 7: v1alpha2.storage.StorageInfo
	(*common.StatusReport)(nil), // 8: v1alpha2.common.StatusReport
}
var file_pkg_grpc_storage_storage_proto_depIdxs = []int32{
	7,  // 0: v1alpha2.storage.InitRequest.storage_info:type_name -> v1alpha2.storage.StorageInfo
	7,  // 1: v1alpha2.storage.File.storage_info:type_name -> v1alpha2.storage.StorageInfo
	0,  // 2: v1alpha2.storage.File.status:type_name -> v1alpha2.storage.Status
	1,  // 3: v1alpha2.storage.StorageInfo.storage_type:type_name -> v1alpha2.storage.StorageType
	6,  // 4: v1alpha2.storage.StorageInfo.s3info:type_name -> v1alpha2.storage.S3Info
	2,  // 5: v1alpha2.storage.StorageManagementService.InitUpload:input_type -> v1alpha2.storage.InitRequest
	4,  // 6: v1alpha2.storage.StorageManagementService.GetUploadUrl:input_type -> v1alpha2.storage.IDmessage
	4,  // 7: v1alpha2.storage.StorageManagementService.FinishUpload:input_type -> v1alpha2.storage.IDmessage
	4,  // 8: v1alpha2.storage.StorageManagementService.DeleteUpload:input_type -> v1alpha2.storage.IDmessage
	4,  // 9: v1alpha2.storage.StorageManagementService.GetDownload:input_type -> v1alpha2.storage.IDmessage
	4,  // 10: v1alpha2.storage.StorageManagementService.InitUpload:output_type -> v1alpha2.storage.IDmessage
	3,  // 11: v1alpha2.storage.StorageManagementService.GetUploadUrl:output_type -> v1alpha2.storage.UrlResponse
	8,  // 12: v1alpha2.storage.StorageManagementService.FinishUpload:output_type -> v1alpha2.common.StatusReport
	8,  // 13: v1alpha2.storage.StorageManagementService.DeleteUpload:output_type -> v1alpha2.common.StatusReport
	3,  // 14: v1alpha2.storage.StorageManagementService.GetDownload:output_type -> v1alpha2.storage.UrlResponse
	10, // [10:15] is the sub-list for method output_type
	5,  // [5:10] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_pkg_grpc_storage_storage_proto_init() }
func file_pkg_grpc_storage_storage_proto_init() {
	if File_pkg_grpc_storage_storage_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_grpc_storage_storage_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InitRequest); i {
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
		file_pkg_grpc_storage_storage_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UrlResponse); i {
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
		file_pkg_grpc_storage_storage_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IDmessage); i {
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
		file_pkg_grpc_storage_storage_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*File); i {
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
		file_pkg_grpc_storage_storage_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*S3Info); i {
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
		file_pkg_grpc_storage_storage_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StorageInfo); i {
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
	file_pkg_grpc_storage_storage_proto_msgTypes[5].OneofWrappers = []interface{}{
		(*StorageInfo_S3Info)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pkg_grpc_storage_storage_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_grpc_storage_storage_proto_goTypes,
		DependencyIndexes: file_pkg_grpc_storage_storage_proto_depIdxs,
		EnumInfos:         file_pkg_grpc_storage_storage_proto_enumTypes,
		MessageInfos:      file_pkg_grpc_storage_storage_proto_msgTypes,
	}.Build()
	File_pkg_grpc_storage_storage_proto = out.File
	file_pkg_grpc_storage_storage_proto_rawDesc = nil
	file_pkg_grpc_storage_storage_proto_goTypes = nil
	file_pkg_grpc_storage_storage_proto_depIdxs = nil
}
