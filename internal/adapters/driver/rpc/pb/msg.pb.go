// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.1
// source: internal/adapters/driver/rpc/proto/msg.proto

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

type TextFieldReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ts        string `protobuf:"bytes,1,opt,name=ts,proto3" json:"ts,omitempty"`
	Name      string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Filename  string `protobuf:"bytes,3,opt,name=filename,proto3" json:"filename,omitempty"`
	ByteChunk []byte `protobuf:"bytes,4,opt,name=byteChunk,proto3" json:"byteChunk,omitempty"`
	IsFirst   bool   `protobuf:"varint,5,opt,name=isFirst,proto3" json:"isFirst,omitempty"`
	IsLast    bool   `protobuf:"varint,6,opt,name=isLast,proto3" json:"isLast,omitempty"`
}

func (x *TextFieldReq) Reset() {
	*x = TextFieldReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_adapters_driver_rpc_proto_msg_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TextFieldReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TextFieldReq) ProtoMessage() {}

func (x *TextFieldReq) ProtoReflect() protoreflect.Message {
	mi := &file_internal_adapters_driver_rpc_proto_msg_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TextFieldReq.ProtoReflect.Descriptor instead.
func (*TextFieldReq) Descriptor() ([]byte, []int) {
	return file_internal_adapters_driver_rpc_proto_msg_proto_rawDescGZIP(), []int{0}
}

func (x *TextFieldReq) GetTs() string {
	if x != nil {
		return x.Ts
	}
	return ""
}

func (x *TextFieldReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *TextFieldReq) GetFilename() string {
	if x != nil {
		return x.Filename
	}
	return ""
}

func (x *TextFieldReq) GetByteChunk() []byte {
	if x != nil {
		return x.ByteChunk
	}
	return nil
}

func (x *TextFieldReq) GetIsFirst() bool {
	if x != nil {
		return x.IsFirst
	}
	return false
}

func (x *TextFieldReq) GetIsLast() bool {
	if x != nil {
		return x.IsLast
	}
	return false
}

type TextFieldRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result bool `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *TextFieldRes) Reset() {
	*x = TextFieldRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_adapters_driver_rpc_proto_msg_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TextFieldRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TextFieldRes) ProtoMessage() {}

func (x *TextFieldRes) ProtoReflect() protoreflect.Message {
	mi := &file_internal_adapters_driver_rpc_proto_msg_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TextFieldRes.ProtoReflect.Descriptor instead.
func (*TextFieldRes) Descriptor() ([]byte, []int) {
	return file_internal_adapters_driver_rpc_proto_msg_proto_rawDescGZIP(), []int{1}
}

func (x *TextFieldRes) GetResult() bool {
	if x != nil {
		return x.Result
	}
	return false
}

type FileInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ts        string `protobuf:"bytes,1,opt,name=ts,proto3" json:"ts,omitempty"`
	FieldName string `protobuf:"bytes,2,opt,name=fieldName,proto3" json:"fieldName,omitempty"`
	FileName  string `protobuf:"bytes,3,opt,name=fileName,proto3" json:"fileName,omitempty"`
	IsFirst   bool   `protobuf:"varint,4,opt,name=isFirst,proto3" json:"isFirst,omitempty"`
}

func (x *FileInfo) Reset() {
	*x = FileInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_adapters_driver_rpc_proto_msg_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileInfo) ProtoMessage() {}

func (x *FileInfo) ProtoReflect() protoreflect.Message {
	mi := &file_internal_adapters_driver_rpc_proto_msg_proto_msgTypes[2]
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
	return file_internal_adapters_driver_rpc_proto_msg_proto_rawDescGZIP(), []int{2}
}

func (x *FileInfo) GetTs() string {
	if x != nil {
		return x.Ts
	}
	return ""
}

func (x *FileInfo) GetFieldName() string {
	if x != nil {
		return x.FieldName
	}
	return ""
}

func (x *FileInfo) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

func (x *FileInfo) GetIsFirst() bool {
	if x != nil {
		return x.IsFirst
	}
	return false
}

type FileData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ts        string `protobuf:"bytes,1,opt,name=ts,proto3" json:"ts,omitempty"`
	FieldName string `protobuf:"bytes,2,opt,name=fieldName,proto3" json:"fieldName,omitempty"`
	Number    uint32 `protobuf:"varint,3,opt,name=number,proto3" json:"number,omitempty"`
	ByteChunk []byte `protobuf:"bytes,4,opt,name=byteChunk,proto3" json:"byteChunk,omitempty"`
	IsLast    bool   `protobuf:"varint,5,opt,name=isLast,proto3" json:"isLast,omitempty"`
}

func (x *FileData) Reset() {
	*x = FileData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_adapters_driver_rpc_proto_msg_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileData) ProtoMessage() {}

func (x *FileData) ProtoReflect() protoreflect.Message {
	mi := &file_internal_adapters_driver_rpc_proto_msg_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileData.ProtoReflect.Descriptor instead.
func (*FileData) Descriptor() ([]byte, []int) {
	return file_internal_adapters_driver_rpc_proto_msg_proto_rawDescGZIP(), []int{3}
}

func (x *FileData) GetTs() string {
	if x != nil {
		return x.Ts
	}
	return ""
}

func (x *FileData) GetFieldName() string {
	if x != nil {
		return x.FieldName
	}
	return ""
}

func (x *FileData) GetNumber() uint32 {
	if x != nil {
		return x.Number
	}
	return 0
}

func (x *FileData) GetByteChunk() []byte {
	if x != nil {
		return x.ByteChunk
	}
	return nil
}

func (x *FileData) GetIsLast() bool {
	if x != nil {
		return x.IsLast
	}
	return false
}

type FileUploadReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Info:
	//	*FileUploadReq_FileInfo
	//	*FileUploadReq_FileData
	Info isFileUploadReq_Info `protobuf_oneof:"info"`
}

func (x *FileUploadReq) Reset() {
	*x = FileUploadReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_adapters_driver_rpc_proto_msg_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileUploadReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileUploadReq) ProtoMessage() {}

func (x *FileUploadReq) ProtoReflect() protoreflect.Message {
	mi := &file_internal_adapters_driver_rpc_proto_msg_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileUploadReq.ProtoReflect.Descriptor instead.
func (*FileUploadReq) Descriptor() ([]byte, []int) {
	return file_internal_adapters_driver_rpc_proto_msg_proto_rawDescGZIP(), []int{4}
}

func (m *FileUploadReq) GetInfo() isFileUploadReq_Info {
	if m != nil {
		return m.Info
	}
	return nil
}

func (x *FileUploadReq) GetFileInfo() *FileInfo {
	if x, ok := x.GetInfo().(*FileUploadReq_FileInfo); ok {
		return x.FileInfo
	}
	return nil
}

func (x *FileUploadReq) GetFileData() *FileData {
	if x, ok := x.GetInfo().(*FileUploadReq_FileData); ok {
		return x.FileData
	}
	return nil
}

type isFileUploadReq_Info interface {
	isFileUploadReq_Info()
}

type FileUploadReq_FileInfo struct {
	FileInfo *FileInfo `protobuf:"bytes,1,opt,name=fileInfo,proto3,oneof"`
}

type FileUploadReq_FileData struct {
	FileData *FileData `protobuf:"bytes,2,opt,name=fileData,proto3,oneof"`
}

func (*FileUploadReq_FileInfo) isFileUploadReq_Info() {}

func (*FileUploadReq_FileData) isFileUploadReq_Info() {}

type FileUploadRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileName string `protobuf:"bytes,1,opt,name=fileName,proto3" json:"fileName,omitempty"`
	FileSize uint32 `protobuf:"varint,2,opt,name=fileSize,proto3" json:"fileSize,omitempty"`
}

func (x *FileUploadRes) Reset() {
	*x = FileUploadRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_adapters_driver_rpc_proto_msg_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileUploadRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileUploadRes) ProtoMessage() {}

func (x *FileUploadRes) ProtoReflect() protoreflect.Message {
	mi := &file_internal_adapters_driver_rpc_proto_msg_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileUploadRes.ProtoReflect.Descriptor instead.
func (*FileUploadRes) Descriptor() ([]byte, []int) {
	return file_internal_adapters_driver_rpc_proto_msg_proto_rawDescGZIP(), []int{5}
}

func (x *FileUploadRes) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

func (x *FileUploadRes) GetFileSize() uint32 {
	if x != nil {
		return x.FileSize
	}
	return 0
}

var File_internal_adapters_driver_rpc_proto_msg_proto protoreflect.FileDescriptor

var file_internal_adapters_driver_rpc_proto_msg_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x61, 0x64, 0x61, 0x70, 0x74,
	0x65, 0x72, 0x73, 0x2f, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x73, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03,
	0x72, 0x70, 0x63, 0x22, 0x9e, 0x01, 0x0a, 0x0c, 0x54, 0x65, 0x78, 0x74, 0x46, 0x69, 0x65, 0x6c,
	0x64, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x74, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x62, 0x79, 0x74, 0x65, 0x43, 0x68, 0x75, 0x6e,
	0x6b, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x62, 0x79, 0x74, 0x65, 0x43, 0x68, 0x75,
	0x6e, 0x6b, 0x12, 0x18, 0x0a, 0x07, 0x69, 0x73, 0x46, 0x69, 0x72, 0x73, 0x74, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x07, 0x69, 0x73, 0x46, 0x69, 0x72, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06,
	0x69, 0x73, 0x4c, 0x61, 0x73, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x69, 0x73,
	0x4c, 0x61, 0x73, 0x74, 0x22, 0x26, 0x0a, 0x0c, 0x54, 0x65, 0x78, 0x74, 0x46, 0x69, 0x65, 0x6c,
	0x64, 0x52, 0x65, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x6e, 0x0a, 0x08,
	0x46, 0x69, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x74, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x4e, 0x61,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x69, 0x73, 0x46, 0x69, 0x72, 0x73, 0x74, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x07, 0x69, 0x73, 0x46, 0x69, 0x72, 0x73, 0x74, 0x22, 0x86, 0x01, 0x0a,
	0x08, 0x46, 0x69, 0x6c, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x74, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12,
	0x1c, 0x0a, 0x09, 0x62, 0x79, 0x74, 0x65, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x09, 0x62, 0x79, 0x74, 0x65, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x12, 0x16, 0x0a,
	0x06, 0x69, 0x73, 0x4c, 0x61, 0x73, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x69,
	0x73, 0x4c, 0x61, 0x73, 0x74, 0x22, 0x71, 0x0a, 0x0d, 0x46, 0x69, 0x6c, 0x65, 0x55, 0x70, 0x6c,
	0x6f, 0x61, 0x64, 0x52, 0x65, 0x71, 0x12, 0x2b, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x49, 0x6e,
	0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x46,
	0x69, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x48, 0x00, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x2b, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x44, 0x61, 0x74, 0x61, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x46, 0x69, 0x6c, 0x65,
	0x44, 0x61, 0x74, 0x61, 0x48, 0x00, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x44, 0x61, 0x74, 0x61,
	0x42, 0x06, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x22, 0x47, 0x0a, 0x0d, 0x46, 0x69, 0x6c, 0x65,
	0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c,
	0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c,
	0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x53, 0x69, 0x7a,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x53, 0x69, 0x7a,
	0x65, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_internal_adapters_driver_rpc_proto_msg_proto_rawDescOnce sync.Once
	file_internal_adapters_driver_rpc_proto_msg_proto_rawDescData = file_internal_adapters_driver_rpc_proto_msg_proto_rawDesc
)

func file_internal_adapters_driver_rpc_proto_msg_proto_rawDescGZIP() []byte {
	file_internal_adapters_driver_rpc_proto_msg_proto_rawDescOnce.Do(func() {
		file_internal_adapters_driver_rpc_proto_msg_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_adapters_driver_rpc_proto_msg_proto_rawDescData)
	})
	return file_internal_adapters_driver_rpc_proto_msg_proto_rawDescData
}

var file_internal_adapters_driver_rpc_proto_msg_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_internal_adapters_driver_rpc_proto_msg_proto_goTypes = []interface{}{
	(*TextFieldReq)(nil),  // 0: rpc.TextFieldReq
	(*TextFieldRes)(nil),  // 1: rpc.TextFieldRes
	(*FileInfo)(nil),      // 2: rpc.FileInfo
	(*FileData)(nil),      // 3: rpc.FileData
	(*FileUploadReq)(nil), // 4: rpc.FileUploadReq
	(*FileUploadRes)(nil), // 5: rpc.FileUploadRes
}
var file_internal_adapters_driver_rpc_proto_msg_proto_depIdxs = []int32{
	2, // 0: rpc.FileUploadReq.fileInfo:type_name -> rpc.FileInfo
	3, // 1: rpc.FileUploadReq.fileData:type_name -> rpc.FileData
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_internal_adapters_driver_rpc_proto_msg_proto_init() }
func file_internal_adapters_driver_rpc_proto_msg_proto_init() {
	if File_internal_adapters_driver_rpc_proto_msg_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_adapters_driver_rpc_proto_msg_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TextFieldReq); i {
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
		file_internal_adapters_driver_rpc_proto_msg_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TextFieldRes); i {
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
		file_internal_adapters_driver_rpc_proto_msg_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_internal_adapters_driver_rpc_proto_msg_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileData); i {
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
		file_internal_adapters_driver_rpc_proto_msg_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileUploadReq); i {
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
		file_internal_adapters_driver_rpc_proto_msg_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileUploadRes); i {
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
	file_internal_adapters_driver_rpc_proto_msg_proto_msgTypes[4].OneofWrappers = []interface{}{
		(*FileUploadReq_FileInfo)(nil),
		(*FileUploadReq_FileData)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_internal_adapters_driver_rpc_proto_msg_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_internal_adapters_driver_rpc_proto_msg_proto_goTypes,
		DependencyIndexes: file_internal_adapters_driver_rpc_proto_msg_proto_depIdxs,
		MessageInfos:      file_internal_adapters_driver_rpc_proto_msg_proto_msgTypes,
	}.Build()
	File_internal_adapters_driver_rpc_proto_msg_proto = out.File
	file_internal_adapters_driver_rpc_proto_msg_proto_rawDesc = nil
	file_internal_adapters_driver_rpc_proto_msg_proto_goTypes = nil
	file_internal_adapters_driver_rpc_proto_msg_proto_depIdxs = nil
}
