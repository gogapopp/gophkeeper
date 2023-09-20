// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.2
// source: grpc.proto

package proto

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

type RegisterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Login    string `protobuf:"bytes,1,opt,name=login,proto3" json:"login,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *RegisterRequest) Reset() {
	*x = RegisterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterRequest) ProtoMessage() {}

func (x *RegisterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterRequest.ProtoReflect.Descriptor instead.
func (*RegisterRequest) Descriptor() ([]byte, []int) {
	return file_grpc_proto_rawDescGZIP(), []int{0}
}

func (x *RegisterRequest) GetLogin() string {
	if x != nil {
		return x.Login
	}
	return ""
}

func (x *RegisterRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type RegisterResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result string `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *RegisterResponse) Reset() {
	*x = RegisterResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterResponse) ProtoMessage() {}

func (x *RegisterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterResponse.ProtoReflect.Descriptor instead.
func (*RegisterResponse) Descriptor() ([]byte, []int) {
	return file_grpc_proto_rawDescGZIP(), []int{1}
}

func (x *RegisterResponse) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

type LoginRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Login    string `protobuf:"bytes,1,opt,name=login,proto3" json:"login,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *LoginRequest) Reset() {
	*x = LoginRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginRequest) ProtoMessage() {}

func (x *LoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginRequest.ProtoReflect.Descriptor instead.
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return file_grpc_proto_rawDescGZIP(), []int{2}
}

func (x *LoginRequest) GetLogin() string {
	if x != nil {
		return x.Login
	}
	return ""
}

func (x *LoginRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type LoginResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result string  `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	Jwt    *string `protobuf:"bytes,2,opt,name=jwt,proto3,oneof" json:"jwt,omitempty"`
}

func (x *LoginResponse) Reset() {
	*x = LoginResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginResponse) ProtoMessage() {}

func (x *LoginResponse) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginResponse.ProtoReflect.Descriptor instead.
func (*LoginResponse) Descriptor() ([]byte, []int) {
	return file_grpc_proto_rawDescGZIP(), []int{3}
}

func (x *LoginResponse) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

func (x *LoginResponse) GetJwt() string {
	if x != nil && x.Jwt != nil {
		return *x.Jwt
	}
	return ""
}

type TextDataRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID     int64  `protobuf:"varint,1,opt,name=userID,proto3" json:"userID,omitempty"`
	UniqueKey  string `protobuf:"bytes,2,opt,name=uniqueKey,proto3" json:"uniqueKey,omitempty"`
	TextData   []byte `protobuf:"bytes,3,opt,name=textData,proto3" json:"textData,omitempty"`
	UploadedAt string `protobuf:"bytes,4,opt,name=uploadedAt,proto3" json:"uploadedAt,omitempty"`
	Metainfo   []byte `protobuf:"bytes,5,opt,name=metainfo,proto3" json:"metainfo,omitempty"`
}

func (x *TextDataRequest) Reset() {
	*x = TextDataRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TextDataRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TextDataRequest) ProtoMessage() {}

func (x *TextDataRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TextDataRequest.ProtoReflect.Descriptor instead.
func (*TextDataRequest) Descriptor() ([]byte, []int) {
	return file_grpc_proto_rawDescGZIP(), []int{4}
}

func (x *TextDataRequest) GetUserID() int64 {
	if x != nil {
		return x.UserID
	}
	return 0
}

func (x *TextDataRequest) GetUniqueKey() string {
	if x != nil {
		return x.UniqueKey
	}
	return ""
}

func (x *TextDataRequest) GetTextData() []byte {
	if x != nil {
		return x.TextData
	}
	return nil
}

func (x *TextDataRequest) GetUploadedAt() string {
	if x != nil {
		return x.UploadedAt
	}
	return ""
}

func (x *TextDataRequest) GetMetainfo() []byte {
	if x != nil {
		return x.Metainfo
	}
	return nil
}

type BinaryDataRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID     int64  `protobuf:"varint,1,opt,name=userID,proto3" json:"userID,omitempty"`
	UniqueKey  string `protobuf:"bytes,2,opt,name=uniqueKey,proto3" json:"uniqueKey,omitempty"`
	BinaryData []byte `protobuf:"bytes,3,opt,name=binaryData,proto3" json:"binaryData,omitempty"`
	UploadedAt string `protobuf:"bytes,4,opt,name=uploadedAt,proto3" json:"uploadedAt,omitempty"`
	Metainfo   []byte `protobuf:"bytes,5,opt,name=metainfo,proto3" json:"metainfo,omitempty"`
}

func (x *BinaryDataRequest) Reset() {
	*x = BinaryDataRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BinaryDataRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BinaryDataRequest) ProtoMessage() {}

func (x *BinaryDataRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BinaryDataRequest.ProtoReflect.Descriptor instead.
func (*BinaryDataRequest) Descriptor() ([]byte, []int) {
	return file_grpc_proto_rawDescGZIP(), []int{5}
}

func (x *BinaryDataRequest) GetUserID() int64 {
	if x != nil {
		return x.UserID
	}
	return 0
}

func (x *BinaryDataRequest) GetUniqueKey() string {
	if x != nil {
		return x.UniqueKey
	}
	return ""
}

func (x *BinaryDataRequest) GetBinaryData() []byte {
	if x != nil {
		return x.BinaryData
	}
	return nil
}

func (x *BinaryDataRequest) GetUploadedAt() string {
	if x != nil {
		return x.UploadedAt
	}
	return ""
}

func (x *BinaryDataRequest) GetMetainfo() []byte {
	if x != nil {
		return x.Metainfo
	}
	return nil
}

type CardDataRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID         int64  `protobuf:"varint,1,opt,name=userID,proto3" json:"userID,omitempty"`
	UniqueKey      string `protobuf:"bytes,2,opt,name=uniqueKey,proto3" json:"uniqueKey,omitempty"`
	CardNumberData []byte `protobuf:"bytes,3,opt,name=cardNumberData,proto3" json:"cardNumberData,omitempty"`
	CardNameData   []byte `protobuf:"bytes,4,opt,name=cardNameData,proto3" json:"cardNameData,omitempty"`
	CardDateData   []byte `protobuf:"bytes,5,opt,name=cardDateData,proto3" json:"cardDateData,omitempty"`
	CvvData        []byte `protobuf:"bytes,6,opt,name=cvvData,proto3" json:"cvvData,omitempty"`
	UploadedAt     string `protobuf:"bytes,7,opt,name=uploadedAt,proto3" json:"uploadedAt,omitempty"`
	Metainfo       []byte `protobuf:"bytes,8,opt,name=metainfo,proto3" json:"metainfo,omitempty"`
}

func (x *CardDataRequest) Reset() {
	*x = CardDataRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CardDataRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CardDataRequest) ProtoMessage() {}

func (x *CardDataRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CardDataRequest.ProtoReflect.Descriptor instead.
func (*CardDataRequest) Descriptor() ([]byte, []int) {
	return file_grpc_proto_rawDescGZIP(), []int{6}
}

func (x *CardDataRequest) GetUserID() int64 {
	if x != nil {
		return x.UserID
	}
	return 0
}

func (x *CardDataRequest) GetUniqueKey() string {
	if x != nil {
		return x.UniqueKey
	}
	return ""
}

func (x *CardDataRequest) GetCardNumberData() []byte {
	if x != nil {
		return x.CardNumberData
	}
	return nil
}

func (x *CardDataRequest) GetCardNameData() []byte {
	if x != nil {
		return x.CardNameData
	}
	return nil
}

func (x *CardDataRequest) GetCardDateData() []byte {
	if x != nil {
		return x.CardDateData
	}
	return nil
}

func (x *CardDataRequest) GetCvvData() []byte {
	if x != nil {
		return x.CvvData
	}
	return nil
}

func (x *CardDataRequest) GetUploadedAt() string {
	if x != nil {
		return x.UploadedAt
	}
	return ""
}

func (x *CardDataRequest) GetMetainfo() []byte {
	if x != nil {
		return x.Metainfo
	}
	return nil
}

type Result struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result string `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *Result) Reset() {
	*x = Result{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Result) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Result) ProtoMessage() {}

func (x *Result) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Result.ProtoReflect.Descriptor instead.
func (*Result) Descriptor() ([]byte, []int) {
	return file_grpc_proto_rawDescGZIP(), []int{7}
}

func (x *Result) GetResult() string {
	if x != nil {
		return x.Result
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
		mi := &file_grpc_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_msgTypes[8]
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
	return file_grpc_proto_rawDescGZIP(), []int{8}
}

var File_grpc_proto protoreflect.FileDescriptor

var file_grpc_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x22, 0x43, 0x0a, 0x0f, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x6f, 0x67,
	0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x12,
	0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x2a, 0x0a, 0x10, 0x52,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x40, 0x0a, 0x0c, 0x4c, 0x6f, 0x67, 0x69, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x1a, 0x0a,
	0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x46, 0x0a, 0x0d, 0x4c, 0x6f, 0x67,
	0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x12, 0x15, 0x0a, 0x03, 0x6a, 0x77, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x00, 0x52, 0x03, 0x6a, 0x77, 0x74, 0x88, 0x01, 0x01, 0x42, 0x06, 0x0a, 0x04, 0x5f, 0x6a, 0x77,
	0x74, 0x22, 0x9f, 0x01, 0x0a, 0x0f, 0x54, 0x65, 0x78, 0x74, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x1c, 0x0a,
	0x09, 0x75, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x4b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x75, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x4b, 0x65, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x74,
	0x65, 0x78, 0x74, 0x44, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x74,
	0x65, 0x78, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1e, 0x0a, 0x0a, 0x75, 0x70, 0x6c, 0x6f, 0x61,
	0x64, 0x65, 0x64, 0x41, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x75, 0x70, 0x6c,
	0x6f, 0x61, 0x64, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x69,
	0x6e, 0x66, 0x6f, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x69,
	0x6e, 0x66, 0x6f, 0x22, 0xa5, 0x01, 0x0a, 0x11, 0x42, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x44, 0x61,
	0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x44, 0x12, 0x1c, 0x0a, 0x09, 0x75, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x4b, 0x65, 0x79, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x4b, 0x65, 0x79, 0x12,
	0x1e, 0x0a, 0x0a, 0x62, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x44, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x0a, 0x62, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x44, 0x61, 0x74, 0x61, 0x12,
	0x1e, 0x0a, 0x0a, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x65, 0x64, 0x41, 0x74, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x65, 0x64, 0x41, 0x74, 0x12,
	0x1a, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x69, 0x6e, 0x66, 0x6f, 0x22, 0x8d, 0x02, 0x0a, 0x0f,
	0x43, 0x61, 0x72, 0x64, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x75, 0x6e, 0x69, 0x71, 0x75,
	0x65, 0x4b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x6e, 0x69, 0x71,
	0x75, 0x65, 0x4b, 0x65, 0x79, 0x12, 0x26, 0x0a, 0x0e, 0x63, 0x61, 0x72, 0x64, 0x4e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0e, 0x63,
	0x61, 0x72, 0x64, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x12, 0x22, 0x0a,
	0x0c, 0x63, 0x61, 0x72, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x44, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x0c, 0x63, 0x61, 0x72, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x44, 0x61, 0x74,
	0x61, 0x12, 0x22, 0x0a, 0x0c, 0x63, 0x61, 0x72, 0x64, 0x44, 0x61, 0x74, 0x65, 0x44, 0x61, 0x74,
	0x61, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0c, 0x63, 0x61, 0x72, 0x64, 0x44, 0x61, 0x74,
	0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x76, 0x76, 0x44, 0x61, 0x74, 0x61,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x63, 0x76, 0x76, 0x44, 0x61, 0x74, 0x61, 0x12,
	0x1e, 0x0a, 0x0a, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x65, 0x64, 0x41, 0x74, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x65, 0x64, 0x41, 0x74, 0x12,
	0x1a, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x69, 0x6e, 0x66, 0x6f, 0x22, 0x20, 0x0a, 0x06, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x07, 0x0a,
	0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x32, 0xc3, 0x02, 0x0a, 0x0c, 0x4d, 0x75, 0x6c, 0x74, 0x69,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x41, 0x0a, 0x08, 0x52, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x65, 0x72, 0x12, 0x19, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x52,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a,
	0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x05, 0x4c, 0x6f,
	0x67, 0x69, 0x6e, 0x12, 0x16, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x4c,
	0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a, 0x0b, 0x41, 0x64, 0x64, 0x54, 0x65, 0x78, 0x74, 0x44,
	0x61, 0x74, 0x61, 0x12, 0x19, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x54,
	0x65, 0x78, 0x74, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10,
	0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x12, 0x3e, 0x0a, 0x0d, 0x41, 0x64, 0x64, 0x42, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x44, 0x61, 0x74,
	0x61, 0x12, 0x1b, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x42, 0x69, 0x6e,
	0x61, 0x72, 0x79, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10,
	0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x12, 0x3a, 0x0a, 0x0b, 0x41, 0x64, 0x64, 0x43, 0x61, 0x72, 0x64, 0x44, 0x61, 0x74, 0x61, 0x12,
	0x19, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x43, 0x61, 0x72, 0x64, 0x44,
	0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x42, 0x10, 0x5a, 0x0e,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_grpc_proto_rawDescOnce sync.Once
	file_grpc_proto_rawDescData = file_grpc_proto_rawDesc
)

func file_grpc_proto_rawDescGZIP() []byte {
	file_grpc_proto_rawDescOnce.Do(func() {
		file_grpc_proto_rawDescData = protoimpl.X.CompressGZIP(file_grpc_proto_rawDescData)
	})
	return file_grpc_proto_rawDescData
}

var file_grpc_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_grpc_proto_goTypes = []interface{}{
	(*RegisterRequest)(nil),   // 0: internal.RegisterRequest
	(*RegisterResponse)(nil),  // 1: internal.RegisterResponse
	(*LoginRequest)(nil),      // 2: internal.LoginRequest
	(*LoginResponse)(nil),     // 3: internal.LoginResponse
	(*TextDataRequest)(nil),   // 4: internal.TextDataRequest
	(*BinaryDataRequest)(nil), // 5: internal.BinaryDataRequest
	(*CardDataRequest)(nil),   // 6: internal.CardDataRequest
	(*Result)(nil),            // 7: internal.Result
	(*Empty)(nil),             // 8: internal.Empty
}
var file_grpc_proto_depIdxs = []int32{
	0, // 0: internal.MultiService.Register:input_type -> internal.RegisterRequest
	2, // 1: internal.MultiService.Login:input_type -> internal.LoginRequest
	4, // 2: internal.MultiService.AddTextData:input_type -> internal.TextDataRequest
	5, // 3: internal.MultiService.AddBinaryData:input_type -> internal.BinaryDataRequest
	6, // 4: internal.MultiService.AddCardData:input_type -> internal.CardDataRequest
	1, // 5: internal.MultiService.Register:output_type -> internal.RegisterResponse
	3, // 6: internal.MultiService.Login:output_type -> internal.LoginResponse
	7, // 7: internal.MultiService.AddTextData:output_type -> internal.Result
	7, // 8: internal.MultiService.AddBinaryData:output_type -> internal.Result
	7, // 9: internal.MultiService.AddCardData:output_type -> internal.Result
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_grpc_proto_init() }
func file_grpc_proto_init() {
	if File_grpc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_grpc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterRequest); i {
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
		file_grpc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterResponse); i {
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
		file_grpc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginRequest); i {
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
		file_grpc_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginResponse); i {
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
		file_grpc_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TextDataRequest); i {
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
		file_grpc_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BinaryDataRequest); i {
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
		file_grpc_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CardDataRequest); i {
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
		file_grpc_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Result); i {
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
		file_grpc_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
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
	file_grpc_proto_msgTypes[3].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_grpc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_grpc_proto_goTypes,
		DependencyIndexes: file_grpc_proto_depIdxs,
		MessageInfos:      file_grpc_proto_msgTypes,
	}.Build()
	File_grpc_proto = out.File
	file_grpc_proto_rawDesc = nil
	file_grpc_proto_goTypes = nil
	file_grpc_proto_depIdxs = nil
}
