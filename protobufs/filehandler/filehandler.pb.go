// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: filehandler.proto

package filehandler

import (
	common "github.com/thomas-osgood/ofs/protobufs/common"
	pingpong "github.com/thomas-osgood/ofs/protobufs/pingpong"
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

// message designed to transfer chunks of a
// file from one machine to another.
type FileChunk struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Chunk []byte `protobuf:"bytes,1,opt,name=chunk,proto3" json:"chunk,omitempty"`
}

func (x *FileChunk) Reset() {
	*x = FileChunk{}
	if protoimpl.UnsafeEnabled {
		mi := &file_filehandler_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileChunk) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileChunk) ProtoMessage() {}

func (x *FileChunk) ProtoReflect() protoreflect.Message {
	mi := &file_filehandler_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileChunk.ProtoReflect.Descriptor instead.
func (*FileChunk) Descriptor() ([]byte, []int) {
	return file_filehandler_proto_rawDescGZIP(), []int{0}
}

func (x *FileChunk) GetChunk() []byte {
	if x != nil {
		return x.Chunk
	}
	return nil
}

// message designed to hold the information
// about a specific file.
type FileInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name         string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Sizebytes    int64                  `protobuf:"varint,2,opt,name=sizebytes,proto3" json:"sizebytes,omitempty"`
	Isdir        bool                   `protobuf:"varint,3,opt,name=isdir,proto3" json:"isdir,omitempty"`
	Lastmodified *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=lastmodified,proto3" json:"lastmodified,omitempty"`
}

func (x *FileInfo) Reset() {
	*x = FileInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_filehandler_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileInfo) ProtoMessage() {}

func (x *FileInfo) ProtoReflect() protoreflect.Message {
	mi := &file_filehandler_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileInfo.ProtoReflect.Descriptor instead.
func (*FileInfo) Descriptor() ([]byte, []int) {
	return file_filehandler_proto_rawDescGZIP(), []int{1}
}

func (x *FileInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *FileInfo) GetSizebytes() int64 {
	if x != nil {
		return x.Sizebytes
	}
	return 0
}

func (x *FileInfo) GetIsdir() bool {
	if x != nil {
		return x.Isdir
	}
	return false
}

func (x *FileInfo) GetLastmodified() *timestamppb.Timestamp {
	if x != nil {
		return x.Lastmodified
	}
	return nil
}

// message designed to request a file.
type FileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Filename string `protobuf:"bytes,1,opt,name=filename,proto3" json:"filename,omitempty"`
}

func (x *FileRequest) Reset() {
	*x = FileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_filehandler_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileRequest) ProtoMessage() {}

func (x *FileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_filehandler_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileRequest.ProtoReflect.Descriptor instead.
func (*FileRequest) Descriptor() ([]byte, []int) {
	return file_filehandler_proto_rawDescGZIP(), []int{2}
}

func (x *FileRequest) GetFilename() string {
	if x != nil {
		return x.Filename
	}
	return ""
}

// message designed to hold the information that will
// be passed when a directory request is initiated.
type MakeDirectoryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Dirname string `protobuf:"bytes,1,opt,name=dirname,proto3" json:"dirname,omitempty"`
}

func (x *MakeDirectoryRequest) Reset() {
	*x = MakeDirectoryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_filehandler_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MakeDirectoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MakeDirectoryRequest) ProtoMessage() {}

func (x *MakeDirectoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_filehandler_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MakeDirectoryRequest.ProtoReflect.Descriptor instead.
func (*MakeDirectoryRequest) Descriptor() ([]byte, []int) {
	return file_filehandler_proto_rawDescGZIP(), []int{3}
}

func (x *MakeDirectoryRequest) GetDirname() string {
	if x != nil {
		return x.Dirname
	}
	return ""
}

// message designed to request a file on the server
// be renamed.
type RenameFileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Oldfilename string `protobuf:"bytes,1,opt,name=oldfilename,proto3" json:"oldfilename,omitempty"`
	Newfilename string `protobuf:"bytes,2,opt,name=newfilename,proto3" json:"newfilename,omitempty"`
}

func (x *RenameFileRequest) Reset() {
	*x = RenameFileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_filehandler_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RenameFileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RenameFileRequest) ProtoMessage() {}

func (x *RenameFileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_filehandler_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RenameFileRequest.ProtoReflect.Descriptor instead.
func (*RenameFileRequest) Descriptor() ([]byte, []int) {
	return file_filehandler_proto_rawDescGZIP(), []int{4}
}

func (x *RenameFileRequest) GetOldfilename() string {
	if x != nil {
		return x.Oldfilename
	}
	return ""
}

func (x *RenameFileRequest) GetNewfilename() string {
	if x != nil {
		return x.Newfilename
	}
	return ""
}

// message designe to hold storage information breakdown.
type StorageInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// total space consumed.
	Total uint64 `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	// space consumed by the downloads directory.
	Downloads uint64 `protobuf:"varint,2,opt,name=downloads,proto3" json:"downloads,omitempty"`
	// space consumed by the uploads directory.
	Uploads uint64 `protobuf:"varint,3,opt,name=uploads,proto3" json:"uploads,omitempty"`
}

func (x *StorageInfo) Reset() {
	*x = StorageInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_filehandler_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StorageInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StorageInfo) ProtoMessage() {}

func (x *StorageInfo) ProtoReflect() protoreflect.Message {
	mi := &file_filehandler_proto_msgTypes[5]
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
	return file_filehandler_proto_rawDescGZIP(), []int{5}
}

func (x *StorageInfo) GetTotal() uint64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *StorageInfo) GetDownloads() uint64 {
	if x != nil {
		return x.Downloads
	}
	return 0
}

func (x *StorageInfo) GetUploads() uint64 {
	if x != nil {
		return x.Uploads
	}
	return 0
}

var File_filehandler_proto protoreflect.FileDescriptor

var file_filehandler_proto_rawDesc = []byte{
	0x0a, 0x11, 0x66, 0x69, 0x6c, 0x65, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x66, 0x69, 0x6c, 0x65, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72,
	0x1a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0e,
	0x70, 0x69, 0x6e, 0x67, 0x70, 0x6f, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x21, 0x0a, 0x09, 0x46, 0x69, 0x6c, 0x65, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x12, 0x14, 0x0a, 0x05,
	0x63, 0x68, 0x75, 0x6e, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x63, 0x68, 0x75,
	0x6e, 0x6b, 0x22, 0x92, 0x01, 0x0a, 0x08, 0x46, 0x69, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x7a, 0x65, 0x62, 0x79, 0x74, 0x65, 0x73,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x73, 0x69, 0x7a, 0x65, 0x62, 0x79, 0x74, 0x65,
	0x73, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x73, 0x64, 0x69, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x05, 0x69, 0x73, 0x64, 0x69, 0x72, 0x12, 0x3e, 0x0a, 0x0c, 0x6c, 0x61, 0x73, 0x74, 0x6d,
	0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0c, 0x6c, 0x61, 0x73, 0x74, 0x6d,
	0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x22, 0x29, 0x0a, 0x0b, 0x46, 0x69, 0x6c, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61,
	0x6d, 0x65, 0x22, 0x30, 0x0a, 0x14, 0x4d, 0x61, 0x6b, 0x65, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74,
	0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x69,
	0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x64, 0x69, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x22, 0x57, 0x0a, 0x11, 0x52, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x46, 0x69,
	0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x6f, 0x6c, 0x64,
	0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x6f, 0x6c, 0x64, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x6e,
	0x65, 0x77, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x6e, 0x65, 0x77, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x5b, 0x0a,
	0x0b, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x12, 0x1c, 0x0a, 0x09, 0x64, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x73, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x64, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x73,
	0x12, 0x18, 0x0a, 0x07, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x07, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x73, 0x32, 0x8d, 0x05, 0x0a, 0x0b, 0x46,
	0x69, 0x6c, 0x65, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3f, 0x0a, 0x0a, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x18, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x68,
	0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x15, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00, 0x12, 0x41, 0x0a, 0x0c, 0x44,
	0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x16, 0x2e, 0x66, 0x69,
	0x6c, 0x65, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x43, 0x68,
	0x75, 0x6e, 0x6b, 0x1a, 0x15, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00, 0x28, 0x01, 0x12, 0x40,
	0x0a, 0x0b, 0x44, 0x65, 0x63, 0x72, 0x79, 0x70, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x18, 0x2e,
	0x66, 0x69, 0x6c, 0x65, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x2e, 0x46, 0x69, 0x6c, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00,
	0x12, 0x40, 0x0a, 0x0b, 0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x12,
	0x18, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x2e, 0x46, 0x69,
	0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x22, 0x00, 0x12, 0x35, 0x0a, 0x09, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x12,
	0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x15,
	0x2e, 0x66, 0x69, 0x6c, 0x65, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x2e, 0x46, 0x69, 0x6c,
	0x65, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x00, 0x30, 0x01, 0x12, 0x4b, 0x0a, 0x0d, 0x4d, 0x61, 0x6b,
	0x65, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x21, 0x2e, 0x66, 0x69, 0x6c,
	0x65, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x2e, 0x4d, 0x61, 0x6b, 0x65, 0x44, 0x69, 0x72,
	0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x22, 0x00, 0x12, 0x28, 0x0a, 0x04, 0x50, 0x69, 0x6e, 0x67, 0x12, 0x0e,
	0x2e, 0x70, 0x69, 0x6e, 0x67, 0x70, 0x6f, 0x6e, 0x67, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x1a, 0x0e,
	0x2e, 0x70, 0x69, 0x6e, 0x67, 0x70, 0x6f, 0x6e, 0x67, 0x2e, 0x50, 0x6f, 0x6e, 0x67, 0x22, 0x00,
	0x12, 0x45, 0x0a, 0x0a, 0x52, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x1e,
	0x2e, 0x66, 0x69, 0x6c, 0x65, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x6e,
	0x61, 0x6d, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00, 0x12, 0x3d, 0x0a, 0x10, 0x53, 0x74, 0x6f, 0x72, 0x61,
	0x67, 0x65, 0x42, 0x72, 0x65, 0x61, 0x6b, 0x64, 0x6f, 0x77, 0x6e, 0x12, 0x0d, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x18, 0x2e, 0x66, 0x69, 0x6c,
	0x65, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65,
	0x49, 0x6e, 0x66, 0x6f, 0x22, 0x00, 0x12, 0x42, 0x0a, 0x0a, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64,
	0x46, 0x69, 0x6c, 0x65, 0x12, 0x18, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x68, 0x61, 0x6e, 0x64, 0x6c,
	0x65, 0x72, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16,
	0x2e, 0x66, 0x69, 0x6c, 0x65, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x2e, 0x46, 0x69, 0x6c,
	0x65, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x22, 0x00, 0x30, 0x01, 0x42, 0x34, 0x5a, 0x32, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x68, 0x6f, 0x6d, 0x61, 0x73, 0x2d,
	0x6f, 0x73, 0x67, 0x6f, 0x6f, 0x64, 0x2f, 0x6f, 0x66, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x73, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_filehandler_proto_rawDescOnce sync.Once
	file_filehandler_proto_rawDescData = file_filehandler_proto_rawDesc
)

func file_filehandler_proto_rawDescGZIP() []byte {
	file_filehandler_proto_rawDescOnce.Do(func() {
		file_filehandler_proto_rawDescData = protoimpl.X.CompressGZIP(file_filehandler_proto_rawDescData)
	})
	return file_filehandler_proto_rawDescData
}

var file_filehandler_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_filehandler_proto_goTypes = []interface{}{
	(*FileChunk)(nil),             // 0: filehandler.FileChunk
	(*FileInfo)(nil),              // 1: filehandler.FileInfo
	(*FileRequest)(nil),           // 2: filehandler.FileRequest
	(*MakeDirectoryRequest)(nil),  // 3: filehandler.MakeDirectoryRequest
	(*RenameFileRequest)(nil),     // 4: filehandler.RenameFileRequest
	(*StorageInfo)(nil),           // 5: filehandler.StorageInfo
	(*timestamppb.Timestamp)(nil), // 6: google.protobuf.Timestamp
	(*common.Empty)(nil),          // 7: common.Empty
	(*pingpong.Ping)(nil),         // 8: pingpong.Ping
	(*common.StatusMessage)(nil),  // 9: common.StatusMessage
	(*pingpong.Pong)(nil),         // 10: pingpong.Pong
}
var file_filehandler_proto_depIdxs = []int32{
	6,  // 0: filehandler.FileInfo.lastmodified:type_name -> google.protobuf.Timestamp
	2,  // 1: filehandler.Fileservice.DeleteFile:input_type -> filehandler.FileRequest
	0,  // 2: filehandler.Fileservice.DownloadFile:input_type -> filehandler.FileChunk
	2,  // 3: filehandler.Fileservice.DecryptFile:input_type -> filehandler.FileRequest
	2,  // 4: filehandler.Fileservice.EncryptFile:input_type -> filehandler.FileRequest
	7,  // 5: filehandler.Fileservice.ListFiles:input_type -> common.Empty
	3,  // 6: filehandler.Fileservice.MakeDirectory:input_type -> filehandler.MakeDirectoryRequest
	8,  // 7: filehandler.Fileservice.Ping:input_type -> pingpong.Ping
	4,  // 8: filehandler.Fileservice.RenameFile:input_type -> filehandler.RenameFileRequest
	7,  // 9: filehandler.Fileservice.StorageBreakdown:input_type -> common.Empty
	2,  // 10: filehandler.Fileservice.UploadFile:input_type -> filehandler.FileRequest
	9,  // 11: filehandler.Fileservice.DeleteFile:output_type -> common.StatusMessage
	9,  // 12: filehandler.Fileservice.DownloadFile:output_type -> common.StatusMessage
	9,  // 13: filehandler.Fileservice.DecryptFile:output_type -> common.StatusMessage
	9,  // 14: filehandler.Fileservice.EncryptFile:output_type -> common.StatusMessage
	1,  // 15: filehandler.Fileservice.ListFiles:output_type -> filehandler.FileInfo
	9,  // 16: filehandler.Fileservice.MakeDirectory:output_type -> common.StatusMessage
	10, // 17: filehandler.Fileservice.Ping:output_type -> pingpong.Pong
	9,  // 18: filehandler.Fileservice.RenameFile:output_type -> common.StatusMessage
	5,  // 19: filehandler.Fileservice.StorageBreakdown:output_type -> filehandler.StorageInfo
	0,  // 20: filehandler.Fileservice.UploadFile:output_type -> filehandler.FileChunk
	11, // [11:21] is the sub-list for method output_type
	1,  // [1:11] is the sub-list for method input_type
	1,  // [1:1] is the sub-list for extension type_name
	1,  // [1:1] is the sub-list for extension extendee
	0,  // [0:1] is the sub-list for field type_name
}

func init() { file_filehandler_proto_init() }
func file_filehandler_proto_init() {
	if File_filehandler_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_filehandler_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileChunk); i {
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
		file_filehandler_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileInfo); i {
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
		file_filehandler_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileRequest); i {
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
		file_filehandler_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MakeDirectoryRequest); i {
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
		file_filehandler_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RenameFileRequest); i {
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
		file_filehandler_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_filehandler_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_filehandler_proto_goTypes,
		DependencyIndexes: file_filehandler_proto_depIdxs,
		MessageInfos:      file_filehandler_proto_msgTypes,
	}.Build()
	File_filehandler_proto = out.File
	file_filehandler_proto_rawDesc = nil
	file_filehandler_proto_goTypes = nil
	file_filehandler_proto_depIdxs = nil
}
