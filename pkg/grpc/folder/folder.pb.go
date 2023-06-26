// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.23.3
// source: folder.proto

package folder

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

type FolderDTO struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId     int64  `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	FolderName string `protobuf:"bytes,3,opt,name=folder_name,json=folderName,proto3" json:"folder_name,omitempty"`
}

func (x *FolderDTO) Reset() {
	*x = FolderDTO{}
	if protoimpl.UnsafeEnabled {
		mi := &file_folder_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FolderDTO) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FolderDTO) ProtoMessage() {}

func (x *FolderDTO) ProtoReflect() protoreflect.Message {
	mi := &file_folder_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FolderDTO.ProtoReflect.Descriptor instead.
func (*FolderDTO) Descriptor() ([]byte, []int) {
	return file_folder_proto_rawDescGZIP(), []int{0}
}

func (x *FolderDTO) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *FolderDTO) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *FolderDTO) GetFolderName() string {
	if x != nil {
		return x.FolderName
	}
	return ""
}

type FolderAddDTO struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId     int64  `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	FolderName string `protobuf:"bytes,2,opt,name=folder_name,json=folderName,proto3" json:"folder_name,omitempty"`
}

func (x *FolderAddDTO) Reset() {
	*x = FolderAddDTO{}
	if protoimpl.UnsafeEnabled {
		mi := &file_folder_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FolderAddDTO) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FolderAddDTO) ProtoMessage() {}

func (x *FolderAddDTO) ProtoReflect() protoreflect.Message {
	mi := &file_folder_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FolderAddDTO.ProtoReflect.Descriptor instead.
func (*FolderAddDTO) Descriptor() ([]byte, []int) {
	return file_folder_proto_rawDescGZIP(), []int{1}
}

func (x *FolderAddDTO) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *FolderAddDTO) GetFolderName() string {
	if x != nil {
		return x.FolderName
	}
	return ""
}

type FolderDeleteDTO struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId int64 `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *FolderDeleteDTO) Reset() {
	*x = FolderDeleteDTO{}
	if protoimpl.UnsafeEnabled {
		mi := &file_folder_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FolderDeleteDTO) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FolderDeleteDTO) ProtoMessage() {}

func (x *FolderDeleteDTO) ProtoReflect() protoreflect.Message {
	mi := &file_folder_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FolderDeleteDTO.ProtoReflect.Descriptor instead.
func (*FolderDeleteDTO) Descriptor() ([]byte, []int) {
	return file_folder_proto_rawDescGZIP(), []int{2}
}

func (x *FolderDeleteDTO) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *FolderDeleteDTO) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type FolderDeleteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Mess    string `protobuf:"bytes,2,opt,name=mess,proto3" json:"mess,omitempty"`
}

func (x *FolderDeleteResponse) Reset() {
	*x = FolderDeleteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_folder_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FolderDeleteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FolderDeleteResponse) ProtoMessage() {}

func (x *FolderDeleteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_folder_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FolderDeleteResponse.ProtoReflect.Descriptor instead.
func (*FolderDeleteResponse) Descriptor() ([]byte, []int) {
	return file_folder_proto_rawDescGZIP(), []int{3}
}

func (x *FolderDeleteResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *FolderDeleteResponse) GetMess() string {
	if x != nil {
		return x.Mess
	}
	return ""
}

type FolderRenameDTO struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	NewFolderName string `protobuf:"bytes,2,opt,name=new_folder_name,json=newFolderName,proto3" json:"new_folder_name,omitempty"`
	UserId        int64  `protobuf:"varint,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *FolderRenameDTO) Reset() {
	*x = FolderRenameDTO{}
	if protoimpl.UnsafeEnabled {
		mi := &file_folder_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FolderRenameDTO) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FolderRenameDTO) ProtoMessage() {}

func (x *FolderRenameDTO) ProtoReflect() protoreflect.Message {
	mi := &file_folder_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FolderRenameDTO.ProtoReflect.Descriptor instead.
func (*FolderRenameDTO) Descriptor() ([]byte, []int) {
	return file_folder_proto_rawDescGZIP(), []int{4}
}

func (x *FolderRenameDTO) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *FolderRenameDTO) GetNewFolderName() string {
	if x != nil {
		return x.NewFolderName
	}
	return ""
}

func (x *FolderRenameDTO) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type FolderGetDTO struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	FolderName string `protobuf:"bytes,2,opt,name=folder_name,json=folderName,proto3" json:"folder_name,omitempty"`
	UserId     int64  `protobuf:"varint,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *FolderGetDTO) Reset() {
	*x = FolderGetDTO{}
	if protoimpl.UnsafeEnabled {
		mi := &file_folder_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FolderGetDTO) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FolderGetDTO) ProtoMessage() {}

func (x *FolderGetDTO) ProtoReflect() protoreflect.Message {
	mi := &file_folder_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FolderGetDTO.ProtoReflect.Descriptor instead.
func (*FolderGetDTO) Descriptor() ([]byte, []int) {
	return file_folder_proto_rawDescGZIP(), []int{5}
}

func (x *FolderGetDTO) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *FolderGetDTO) GetFolderName() string {
	if x != nil {
		return x.FolderName
	}
	return ""
}

func (x *FolderGetDTO) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type Details struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Mess    string `protobuf:"bytes,2,opt,name=mess,proto3" json:"mess,omitempty"`
}

func (x *Details) Reset() {
	*x = Details{}
	if protoimpl.UnsafeEnabled {
		mi := &file_folder_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Details) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Details) ProtoMessage() {}

func (x *Details) ProtoReflect() protoreflect.Message {
	mi := &file_folder_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Details.ProtoReflect.Descriptor instead.
func (*Details) Descriptor() ([]byte, []int) {
	return file_folder_proto_rawDescGZIP(), []int{6}
}

func (x *Details) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *Details) GetMess() string {
	if x != nil {
		return x.Mess
	}
	return ""
}

type FolderResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Details *Details   `protobuf:"bytes,1,opt,name=details,proto3" json:"details,omitempty"`
	Folder  *FolderDTO `protobuf:"bytes,2,opt,name=folder,proto3" json:"folder,omitempty"`
}

func (x *FolderResponse) Reset() {
	*x = FolderResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_folder_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FolderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FolderResponse) ProtoMessage() {}

func (x *FolderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_folder_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FolderResponse.ProtoReflect.Descriptor instead.
func (*FolderResponse) Descriptor() ([]byte, []int) {
	return file_folder_proto_rawDescGZIP(), []int{7}
}

func (x *FolderResponse) GetDetails() *Details {
	if x != nil {
		return x.Details
	}
	return nil
}

func (x *FolderResponse) GetFolder() *FolderDTO {
	if x != nil {
		return x.Folder
	}
	return nil
}

var File_folder_proto protoreflect.FileDescriptor

var file_folder_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d,
	0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0x55, 0x0a,
	0x09, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x44, 0x54, 0x4f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72,
	0x4e, 0x61, 0x6d, 0x65, 0x22, 0x48, 0x0a, 0x0c, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x41, 0x64,
	0x64, 0x44, 0x54, 0x4f, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1f, 0x0a,
	0x0b, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x3a,
	0x0a, 0x0f, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x54,
	0x4f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x44, 0x0a, 0x14, 0x46, 0x6f,
	0x6c, 0x64, 0x65, 0x72, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x12, 0x0a, 0x04,
	0x6d, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6d, 0x65, 0x73, 0x73,
	0x22, 0x62, 0x0a, 0x0f, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x52, 0x65, 0x6e, 0x61, 0x6d, 0x65,
	0x44, 0x54, 0x4f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x26, 0x0a, 0x0f, 0x6e, 0x65, 0x77, 0x5f, 0x66, 0x6f, 0x6c, 0x64, 0x65,
	0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6e, 0x65,
	0x77, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x22, 0x58, 0x0a, 0x0c, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x47, 0x65,
	0x74, 0x44, 0x54, 0x4f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x66, 0x6f, 0x6c, 0x64, 0x65,
	0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x37,
	0x0a, 0x07, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6d, 0x65, 0x73, 0x73, 0x22, 0x74, 0x0a, 0x0e, 0x46, 0x6f, 0x6c, 0x64, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a, 0x07, 0x64, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x66, 0x6f, 0x6c,
	0x64, 0x65, 0x72, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x44, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x73, 0x52, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x30, 0x0a, 0x06, 0x66,
	0x6f, 0x6c, 0x64, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x66, 0x6f,
	0x6c, 0x64, 0x65, 0x72, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x46, 0x6f, 0x6c, 0x64,
	0x65, 0x72, 0x44, 0x54, 0x4f, 0x52, 0x06, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x32, 0xb8, 0x02,
	0x0a, 0x0d, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x47, 0x0a, 0x09, 0x41, 0x64, 0x64, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x12, 0x1b, 0x2e, 0x66,
	0x6f, 0x6c, 0x64, 0x65, 0x72, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x46, 0x6f, 0x6c,
	0x64, 0x65, 0x72, 0x41, 0x64, 0x64, 0x44, 0x54, 0x4f, 0x1a, 0x1d, 0x2e, 0x66, 0x6f, 0x6c, 0x64,
	0x65, 0x72, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x46, 0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x12, 0x1e, 0x2e, 0x66, 0x6f, 0x6c, 0x64, 0x65,
	0x72, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x54, 0x4f, 0x1a, 0x16, 0x2e, 0x66, 0x6f, 0x6c, 0x64, 0x65,
	0x72, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73,
	0x12, 0x47, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x12, 0x1b, 0x2e,
	0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x46, 0x6f,
	0x6c, 0x64, 0x65, 0x72, 0x47, 0x65, 0x74, 0x44, 0x54, 0x4f, 0x1a, 0x1d, 0x2e, 0x66, 0x6f, 0x6c,
	0x64, 0x65, 0x72, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x46, 0x6f, 0x6c, 0x64, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4d, 0x0a, 0x0c, 0x52, 0x65, 0x6e,
	0x61, 0x6d, 0x65, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x12, 0x1e, 0x2e, 0x66, 0x6f, 0x6c, 0x64,
	0x65, 0x72, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72,
	0x52, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x44, 0x54, 0x4f, 0x1a, 0x1d, 0x2e, 0x66, 0x6f, 0x6c, 0x64,
	0x65, 0x72, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x21, 0x5a, 0x1f, 0x73, 0x74, 0x6f, 0x72,
	0x61, 0x67, 0x65, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x6b, 0x67, 0x2f,
	0x67, 0x72, 0x70, 0x63, 0x2f, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_folder_proto_rawDescOnce sync.Once
	file_folder_proto_rawDescData = file_folder_proto_rawDesc
)

func file_folder_proto_rawDescGZIP() []byte {
	file_folder_proto_rawDescOnce.Do(func() {
		file_folder_proto_rawDescData = protoimpl.X.CompressGZIP(file_folder_proto_rawDescData)
	})
	return file_folder_proto_rawDescData
}

var file_folder_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_folder_proto_goTypes = []interface{}{
	(*FolderDTO)(nil),            // 0: folderservice.FolderDTO
	(*FolderAddDTO)(nil),         // 1: folderservice.FolderAddDTO
	(*FolderDeleteDTO)(nil),      // 2: folderservice.FolderDeleteDTO
	(*FolderDeleteResponse)(nil), // 3: folderservice.FolderDeleteResponse
	(*FolderRenameDTO)(nil),      // 4: folderservice.FolderRenameDTO
	(*FolderGetDTO)(nil),         // 5: folderservice.FolderGetDTO
	(*Details)(nil),              // 6: folderservice.Details
	(*FolderResponse)(nil),       // 7: folderservice.FolderResponse
}
var file_folder_proto_depIdxs = []int32{
	6, // 0: folderservice.FolderResponse.details:type_name -> folderservice.Details
	0, // 1: folderservice.FolderResponse.folder:type_name -> folderservice.FolderDTO
	1, // 2: folderservice.FolderService.AddFolder:input_type -> folderservice.FolderAddDTO
	2, // 3: folderservice.FolderService.DeleteFolder:input_type -> folderservice.FolderDeleteDTO
	5, // 4: folderservice.FolderService.GetFolder:input_type -> folderservice.FolderGetDTO
	4, // 5: folderservice.FolderService.RenameFolder:input_type -> folderservice.FolderRenameDTO
	7, // 6: folderservice.FolderService.AddFolder:output_type -> folderservice.FolderResponse
	6, // 7: folderservice.FolderService.DeleteFolder:output_type -> folderservice.Details
	7, // 8: folderservice.FolderService.GetFolder:output_type -> folderservice.FolderResponse
	7, // 9: folderservice.FolderService.RenameFolder:output_type -> folderservice.FolderResponse
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_folder_proto_init() }
func file_folder_proto_init() {
	if File_folder_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_folder_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FolderDTO); i {
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
		file_folder_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FolderAddDTO); i {
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
		file_folder_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FolderDeleteDTO); i {
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
		file_folder_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FolderDeleteResponse); i {
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
		file_folder_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FolderRenameDTO); i {
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
		file_folder_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FolderGetDTO); i {
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
		file_folder_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Details); i {
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
		file_folder_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FolderResponse); i {
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
			RawDescriptor: file_folder_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_folder_proto_goTypes,
		DependencyIndexes: file_folder_proto_depIdxs,
		MessageInfos:      file_folder_proto_msgTypes,
	}.Build()
	File_folder_proto = out.File
	file_folder_proto_rawDesc = nil
	file_folder_proto_goTypes = nil
	file_folder_proto_depIdxs = nil
}
