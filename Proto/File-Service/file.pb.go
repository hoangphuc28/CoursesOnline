// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: File-Service/file.proto

package file

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

type File struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileName string `protobuf:"bytes,1,opt,name=fileName,proto3" json:"fileName,omitempty"`
	Size     int64  `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
	Content  []byte `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	Folder   string `protobuf:"bytes,4,opt,name=folder,proto3" json:"folder,omitempty"`
}

func (x *File) Reset() {
	*x = File{}
	if protoimpl.UnsafeEnabled {
		mi := &file_File_Service_file_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *File) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*File) ProtoMessage() {}

func (x *File) ProtoReflect() protoreflect.Message {
	mi := &file_File_Service_file_proto_msgTypes[0]
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
	return file_File_Service_file_proto_rawDescGZIP(), []int{0}
}

func (x *File) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

func (x *File) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *File) GetContent() []byte {
	if x != nil {
		return x.Content
	}
	return nil
}

func (x *File) GetFolder() string {
	if x != nil {
		return x.Folder
	}
	return ""
}

type UploadAvatarRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	File   *File  `protobuf:"bytes,1,opt,name=file,proto3" json:"file,omitempty"`
	OldUrl string `protobuf:"bytes,2,opt,name=oldUrl,proto3" json:"oldUrl,omitempty"`
}

func (x *UploadAvatarRequest) Reset() {
	*x = UploadAvatarRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_File_Service_file_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadAvatarRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadAvatarRequest) ProtoMessage() {}

func (x *UploadAvatarRequest) ProtoReflect() protoreflect.Message {
	mi := &file_File_Service_file_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadAvatarRequest.ProtoReflect.Descriptor instead.
func (*UploadAvatarRequest) Descriptor() ([]byte, []int) {
	return file_File_Service_file_proto_rawDescGZIP(), []int{1}
}

func (x *UploadAvatarRequest) GetFile() *File {
	if x != nil {
		return x.File
	}
	return nil
}

func (x *UploadAvatarRequest) GetOldUrl() string {
	if x != nil {
		return x.OldUrl
	}
	return ""
}

type UploadAvatarResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url   string         `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	Error *ErrorResponse `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *UploadAvatarResponse) Reset() {
	*x = UploadAvatarResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_File_Service_file_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadAvatarResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadAvatarResponse) ProtoMessage() {}

func (x *UploadAvatarResponse) ProtoReflect() protoreflect.Message {
	mi := &file_File_Service_file_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadAvatarResponse.ProtoReflect.Descriptor instead.
func (*UploadAvatarResponse) Descriptor() ([]byte, []int) {
	return file_File_Service_file_proto_rawDescGZIP(), []int{2}
}

func (x *UploadAvatarResponse) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *UploadAvatarResponse) GetError() *ErrorResponse {
	if x != nil {
		return x.Error
	}
	return nil
}

type UploadAssetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	File *File `protobuf:"bytes,1,opt,name=file,proto3" json:"file,omitempty"`
}

func (x *UploadAssetRequest) Reset() {
	*x = UploadAssetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_File_Service_file_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadAssetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadAssetRequest) ProtoMessage() {}

func (x *UploadAssetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_File_Service_file_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadAssetRequest.ProtoReflect.Descriptor instead.
func (*UploadAssetRequest) Descriptor() ([]byte, []int) {
	return file_File_Service_file_proto_rawDescGZIP(), []int{3}
}

func (x *UploadAssetRequest) GetFile() *File {
	if x != nil {
		return x.File
	}
	return nil
}

type Video struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Duration int64  `protobuf:"varint,1,opt,name=duration,proto3" json:"duration,omitempty"`
	Unit     string `protobuf:"bytes,2,opt,name=unit,proto3" json:"unit,omitempty"`
}

func (x *Video) Reset() {
	*x = Video{}
	if protoimpl.UnsafeEnabled {
		mi := &file_File_Service_file_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Video) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Video) ProtoMessage() {}

func (x *Video) ProtoReflect() protoreflect.Message {
	mi := &file_File_Service_file_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Video.ProtoReflect.Descriptor instead.
func (*Video) Descriptor() ([]byte, []int) {
	return file_File_Service_file_proto_rawDescGZIP(), []int{4}
}

func (x *Video) GetDuration() int64 {
	if x != nil {
		return x.Duration
	}
	return 0
}

func (x *Video) GetUnit() string {
	if x != nil {
		return x.Unit
	}
	return ""
}

type UploadAssetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url   string         `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	Video *Video         `protobuf:"bytes,2,opt,name=video,proto3" json:"video,omitempty"`
	Error *ErrorResponse `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *UploadAssetResponse) Reset() {
	*x = UploadAssetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_File_Service_file_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadAssetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadAssetResponse) ProtoMessage() {}

func (x *UploadAssetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_File_Service_file_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadAssetResponse.ProtoReflect.Descriptor instead.
func (*UploadAssetResponse) Descriptor() ([]byte, []int) {
	return file_File_Service_file_proto_rawDescGZIP(), []int{5}
}

func (x *UploadAssetResponse) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *UploadAssetResponse) GetVideo() *Video {
	if x != nil {
		return x.Video
	}
	return nil
}

func (x *UploadAssetResponse) GetError() *ErrorResponse {
	if x != nil {
		return x.Error
	}
	return nil
}

type ErrorResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int64  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *ErrorResponse) Reset() {
	*x = ErrorResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_File_Service_file_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ErrorResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ErrorResponse) ProtoMessage() {}

func (x *ErrorResponse) ProtoReflect() protoreflect.Message {
	mi := &file_File_Service_file_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ErrorResponse.ProtoReflect.Descriptor instead.
func (*ErrorResponse) Descriptor() ([]byte, []int) {
	return file_File_Service_file_proto_rawDescGZIP(), []int{6}
}

func (x *ErrorResponse) GetCode() int64 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *ErrorResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_File_Service_file_proto protoreflect.FileDescriptor

var file_File_Service_file_proto_rawDesc = []byte{
	0x0a, 0x17, 0x46, 0x69, 0x6c, 0x65, 0x2d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x66,
	0x69, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x22,
	0x68, 0x0a, 0x04, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x4e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x22, 0x4d, 0x0a, 0x13, 0x55, 0x70, 0x6c,
	0x6f, 0x61, 0x64, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1e, 0x0a, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a,
	0x2e, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x04, 0x66, 0x69, 0x6c, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x6f, 0x6c, 0x64, 0x55, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x6f, 0x6c, 0x64, 0x55, 0x72, 0x6c, 0x22, 0x53, 0x0a, 0x14, 0x55, 0x70, 0x6c, 0x6f,
	0x61, 0x64, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75,
	0x72, 0x6c, 0x12, 0x29, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x13, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x34, 0x0a,
	0x12, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x41, 0x73, 0x73, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0a, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x04, 0x66,
	0x69, 0x6c, 0x65, 0x22, 0x37, 0x0a, 0x05, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x12, 0x1a, 0x0a, 0x08,
	0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08,
	0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x6e, 0x69, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x6e, 0x69, 0x74, 0x22, 0x75, 0x0a, 0x13,
	0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x41, 0x73, 0x73, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x21, 0x0a, 0x05, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x56, 0x69, 0x64, 0x65,
	0x6f, 0x52, 0x05, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x12, 0x29, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x45,
	0x72, 0x72, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x05, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x22, 0x3d, 0x0a, 0x0d, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x32, 0x9c, 0x01, 0x0a, 0x0b, 0x46, 0x69, 0x6c, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x45, 0x0a, 0x0c, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x41, 0x76, 0x61, 0x74,
	0x61, 0x72, 0x12, 0x19, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64,
	0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e,
	0x66, 0x69, 0x6c, 0x65, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x41, 0x76, 0x61, 0x74, 0x61,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x46, 0x0a, 0x0b, 0x55, 0x70, 0x6c,
	0x6f, 0x61, 0x64, 0x41, 0x73, 0x73, 0x65, 0x74, 0x12, 0x18, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x2e,
	0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x41, 0x73, 0x73, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x19, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64,
	0x41, 0x73, 0x73, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x28, 0x01, 0x30,
	0x01, 0x42, 0x26, 0x5a, 0x24, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x5a, 0x68, 0x6f, 0x61, 0x6e, 0x67, 0x70, 0x2f, 0x55, 0x73, 0x65, 0x72, 0x2d, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_File_Service_file_proto_rawDescOnce sync.Once
	file_File_Service_file_proto_rawDescData = file_File_Service_file_proto_rawDesc
)

func file_File_Service_file_proto_rawDescGZIP() []byte {
	file_File_Service_file_proto_rawDescOnce.Do(func() {
		file_File_Service_file_proto_rawDescData = protoimpl.X.CompressGZIP(file_File_Service_file_proto_rawDescData)
	})
	return file_File_Service_file_proto_rawDescData
}

var file_File_Service_file_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_File_Service_file_proto_goTypes = []interface{}{
	(*File)(nil),                 // 0: file.File
	(*UploadAvatarRequest)(nil),  // 1: file.UploadAvatarRequest
	(*UploadAvatarResponse)(nil), // 2: file.UploadAvatarResponse
	(*UploadAssetRequest)(nil),   // 3: file.UploadAssetRequest
	(*Video)(nil),                // 4: file.Video
	(*UploadAssetResponse)(nil),  // 5: file.UploadAssetResponse
	(*ErrorResponse)(nil),        // 6: file.ErrorResponse
}
var file_File_Service_file_proto_depIdxs = []int32{
	0, // 0: file.UploadAvatarRequest.file:type_name -> file.File
	6, // 1: file.UploadAvatarResponse.error:type_name -> file.ErrorResponse
	0, // 2: file.UploadAssetRequest.file:type_name -> file.File
	4, // 3: file.UploadAssetResponse.video:type_name -> file.Video
	6, // 4: file.UploadAssetResponse.error:type_name -> file.ErrorResponse
	1, // 5: file.FileService.UploadAvatar:input_type -> file.UploadAvatarRequest
	3, // 6: file.FileService.UploadAsset:input_type -> file.UploadAssetRequest
	2, // 7: file.FileService.UploadAvatar:output_type -> file.UploadAvatarResponse
	5, // 8: file.FileService.UploadAsset:output_type -> file.UploadAssetResponse
	7, // [7:9] is the sub-list for method output_type
	5, // [5:7] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_File_Service_file_proto_init() }
func file_File_Service_file_proto_init() {
	if File_File_Service_file_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_File_Service_file_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_File_Service_file_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadAvatarRequest); i {
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
		file_File_Service_file_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadAvatarResponse); i {
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
		file_File_Service_file_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadAssetRequest); i {
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
		file_File_Service_file_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Video); i {
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
		file_File_Service_file_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadAssetResponse); i {
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
		file_File_Service_file_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ErrorResponse); i {
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
			RawDescriptor: file_File_Service_file_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_File_Service_file_proto_goTypes,
		DependencyIndexes: file_File_Service_file_proto_depIdxs,
		MessageInfos:      file_File_Service_file_proto_msgTypes,
	}.Build()
	File_File_Service_file_proto = out.File
	file_File_Service_file_proto_rawDesc = nil
	file_File_Service_file_proto_goTypes = nil
	file_File_Service_file_proto_depIdxs = nil
}
