// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v3.20.3
// source: go_counter.proto

package go_counter

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

type PingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ping string `protobuf:"bytes,1,opt,name=ping,proto3" json:"ping,omitempty"`
}

func (x *PingRequest) Reset() {
	*x = PingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_counter_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingRequest) ProtoMessage() {}

func (x *PingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_go_counter_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingRequest.ProtoReflect.Descriptor instead.
func (*PingRequest) Descriptor() ([]byte, []int) {
	return file_go_counter_proto_rawDescGZIP(), []int{0}
}

func (x *PingRequest) GetPing() string {
	if x != nil {
		return x.Ping
	}
	return ""
}

type PingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pong string `protobuf:"bytes,1,opt,name=pong,proto3" json:"pong,omitempty"`
}

func (x *PingResponse) Reset() {
	*x = PingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_counter_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingResponse) ProtoMessage() {}

func (x *PingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_go_counter_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingResponse.ProtoReflect.Descriptor instead.
func (*PingResponse) Descriptor() ([]byte, []int) {
	return file_go_counter_proto_rawDescGZIP(), []int{1}
}

func (x *PingResponse) GetPong() string {
	if x != nil {
		return x.Pong
	}
	return ""
}

type LikeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId    int64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	ArticleId int64 `protobuf:"varint,2,opt,name=article_id,json=articleId,proto3" json:"article_id,omitempty"`
}

func (x *LikeRequest) Reset() {
	*x = LikeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_counter_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LikeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LikeRequest) ProtoMessage() {}

func (x *LikeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_go_counter_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LikeRequest.ProtoReflect.Descriptor instead.
func (*LikeRequest) Descriptor() ([]byte, []int) {
	return file_go_counter_proto_rawDescGZIP(), []int{2}
}

func (x *LikeRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *LikeRequest) GetArticleId() int64 {
	if x != nil {
		return x.ArticleId
	}
	return 0
}

type LikeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *LikeResponse) Reset() {
	*x = LikeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_counter_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LikeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LikeResponse) ProtoMessage() {}

func (x *LikeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_go_counter_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LikeResponse.ProtoReflect.Descriptor instead.
func (*LikeResponse) Descriptor() ([]byte, []int) {
	return file_go_counter_proto_rawDescGZIP(), []int{3}
}

func (x *LikeResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type FavoriteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId    int64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	ArticleId int64 `protobuf:"varint,2,opt,name=article_id,json=articleId,proto3" json:"article_id,omitempty"`
}

func (x *FavoriteRequest) Reset() {
	*x = FavoriteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_counter_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FavoriteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FavoriteRequest) ProtoMessage() {}

func (x *FavoriteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_go_counter_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FavoriteRequest.ProtoReflect.Descriptor instead.
func (*FavoriteRequest) Descriptor() ([]byte, []int) {
	return file_go_counter_proto_rawDescGZIP(), []int{4}
}

func (x *FavoriteRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *FavoriteRequest) GetArticleId() int64 {
	if x != nil {
		return x.ArticleId
	}
	return 0
}

type FavoriteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *FavoriteResponse) Reset() {
	*x = FavoriteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_counter_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FavoriteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FavoriteResponse) ProtoMessage() {}

func (x *FavoriteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_go_counter_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FavoriteResponse.ProtoReflect.Descriptor instead.
func (*FavoriteResponse) Descriptor() ([]byte, []int) {
	return file_go_counter_proto_rawDescGZIP(), []int{5}
}

func (x *FavoriteResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type ViewRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId    int64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	ArticleId int64 `protobuf:"varint,2,opt,name=article_id,json=articleId,proto3" json:"article_id,omitempty"`
}

func (x *ViewRequest) Reset() {
	*x = ViewRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_counter_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ViewRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ViewRequest) ProtoMessage() {}

func (x *ViewRequest) ProtoReflect() protoreflect.Message {
	mi := &file_go_counter_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ViewRequest.ProtoReflect.Descriptor instead.
func (*ViewRequest) Descriptor() ([]byte, []int) {
	return file_go_counter_proto_rawDescGZIP(), []int{6}
}

func (x *ViewRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *ViewRequest) GetArticleId() int64 {
	if x != nil {
		return x.ArticleId
	}
	return 0
}

type ViewResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *ViewResponse) Reset() {
	*x = ViewResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_counter_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ViewResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ViewResponse) ProtoMessage() {}

func (x *ViewResponse) ProtoReflect() protoreflect.Message {
	mi := &file_go_counter_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ViewResponse.ProtoReflect.Descriptor instead.
func (*ViewResponse) Descriptor() ([]byte, []int) {
	return file_go_counter_proto_rawDescGZIP(), []int{7}
}

func (x *ViewResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_go_counter_proto protoreflect.FileDescriptor

var file_go_counter_proto_rawDesc = []byte{
	0x0a, 0x10, 0x67, 0x6f, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0a, 0x67, 0x6f, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x22, 0x21,
	0x0a, 0x0b, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x70, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x69, 0x6e,
	0x67, 0x22, 0x22, 0x0a, 0x0c, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x70, 0x6f, 0x6e, 0x67, 0x22, 0x45, 0x0a, 0x0b, 0x4c, 0x69, 0x6b, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a,
	0x0a, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x09, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x49, 0x64, 0x22, 0x28, 0x0a, 0x0c,
	0x4c, 0x69, 0x6b, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73,
	0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x49, 0x0a, 0x0f, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x49,
	0x64, 0x22, 0x2c, 0x0a, 0x10, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22,
	0x45, 0x0a, 0x0b, 0x56, 0x69, 0x65, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17,
	0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x72, 0x74, 0x69, 0x63,
	0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x61, 0x72, 0x74,
	0x69, 0x63, 0x6c, 0x65, 0x49, 0x64, 0x22, 0x28, 0x0a, 0x0c, 0x56, 0x69, 0x65, 0x77, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x32, 0x84, 0x02, 0x0a, 0x0a, 0x47, 0x6f, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x12,
	0x39, 0x0a, 0x04, 0x50, 0x69, 0x6e, 0x67, 0x12, 0x17, 0x2e, 0x67, 0x6f, 0x5f, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x65, 0x72, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x18, 0x2e, 0x67, 0x6f, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x50, 0x69,
	0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a, 0x04, 0x4c, 0x69,
	0x6b, 0x65, 0x12, 0x17, 0x2e, 0x67, 0x6f, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x2e,
	0x4c, 0x69, 0x6b, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x67, 0x6f,
	0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x4c, 0x69, 0x6b, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x45, 0x0a, 0x08, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74,
	0x65, 0x12, 0x1b, 0x2e, 0x67, 0x6f, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x46,
	0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c,
	0x2e, 0x67, 0x6f, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x46, 0x61, 0x76, 0x6f,
	0x72, 0x69, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a, 0x04,
	0x56, 0x69, 0x65, 0x77, 0x12, 0x17, 0x2e, 0x67, 0x6f, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65,
	0x72, 0x2e, 0x56, 0x69, 0x65, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e,
	0x67, 0x6f, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x56, 0x69, 0x65, 0x77, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0e, 0x5a, 0x0c, 0x2e, 0x2f, 0x67, 0x6f, 0x5f,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_go_counter_proto_rawDescOnce sync.Once
	file_go_counter_proto_rawDescData = file_go_counter_proto_rawDesc
)

func file_go_counter_proto_rawDescGZIP() []byte {
	file_go_counter_proto_rawDescOnce.Do(func() {
		file_go_counter_proto_rawDescData = protoimpl.X.CompressGZIP(file_go_counter_proto_rawDescData)
	})
	return file_go_counter_proto_rawDescData
}

var file_go_counter_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_go_counter_proto_goTypes = []interface{}{
	(*PingRequest)(nil),      // 0: go_counter.PingRequest
	(*PingResponse)(nil),     // 1: go_counter.PingResponse
	(*LikeRequest)(nil),      // 2: go_counter.LikeRequest
	(*LikeResponse)(nil),     // 3: go_counter.LikeResponse
	(*FavoriteRequest)(nil),  // 4: go_counter.FavoriteRequest
	(*FavoriteResponse)(nil), // 5: go_counter.FavoriteResponse
	(*ViewRequest)(nil),      // 6: go_counter.ViewRequest
	(*ViewResponse)(nil),     // 7: go_counter.ViewResponse
}
var file_go_counter_proto_depIdxs = []int32{
	0, // 0: go_counter.Go_counter.Ping:input_type -> go_counter.PingRequest
	2, // 1: go_counter.Go_counter.Like:input_type -> go_counter.LikeRequest
	4, // 2: go_counter.Go_counter.Favorite:input_type -> go_counter.FavoriteRequest
	6, // 3: go_counter.Go_counter.View:input_type -> go_counter.ViewRequest
	1, // 4: go_counter.Go_counter.Ping:output_type -> go_counter.PingResponse
	3, // 5: go_counter.Go_counter.Like:output_type -> go_counter.LikeResponse
	5, // 6: go_counter.Go_counter.Favorite:output_type -> go_counter.FavoriteResponse
	7, // 7: go_counter.Go_counter.View:output_type -> go_counter.ViewResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_go_counter_proto_init() }
func file_go_counter_proto_init() {
	if File_go_counter_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_go_counter_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PingRequest); i {
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
		file_go_counter_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PingResponse); i {
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
		file_go_counter_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LikeRequest); i {
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
		file_go_counter_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LikeResponse); i {
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
		file_go_counter_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FavoriteRequest); i {
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
		file_go_counter_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FavoriteResponse); i {
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
		file_go_counter_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ViewRequest); i {
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
		file_go_counter_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ViewResponse); i {
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
			RawDescriptor: file_go_counter_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_go_counter_proto_goTypes,
		DependencyIndexes: file_go_counter_proto_depIdxs,
		MessageInfos:      file_go_counter_proto_msgTypes,
	}.Build()
	File_go_counter_proto = out.File
	file_go_counter_proto_rawDesc = nil
	file_go_counter_proto_goTypes = nil
	file_go_counter_proto_depIdxs = nil
}
