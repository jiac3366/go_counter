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

type LikeNumRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ArticleId int64 `protobuf:"varint,1,opt,name=article_id,json=articleId,proto3" json:"article_id,omitempty"`
}

func (x *LikeNumRequest) Reset() {
	*x = LikeNumRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_counter_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LikeNumRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LikeNumRequest) ProtoMessage() {}

func (x *LikeNumRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use LikeNumRequest.ProtoReflect.Descriptor instead.
func (*LikeNumRequest) Descriptor() ([]byte, []int) {
	return file_go_counter_proto_rawDescGZIP(), []int{2}
}

func (x *LikeNumRequest) GetArticleId() int64 {
	if x != nil {
		return x.ArticleId
	}
	return 0
}

type LikeNumResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count bool `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *LikeNumResponse) Reset() {
	*x = LikeNumResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_counter_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LikeNumResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LikeNumResponse) ProtoMessage() {}

func (x *LikeNumResponse) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use LikeNumResponse.ProtoReflect.Descriptor instead.
func (*LikeNumResponse) Descriptor() ([]byte, []int) {
	return file_go_counter_proto_rawDescGZIP(), []int{3}
}

func (x *LikeNumResponse) GetCount() bool {
	if x != nil {
		return x.Count
	}
	return false
}

type LikeRecordRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId    int64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	ArticleId int64 `protobuf:"varint,2,opt,name=article_id,json=articleId,proto3" json:"article_id,omitempty"`
}

func (x *LikeRecordRequest) Reset() {
	*x = LikeRecordRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_counter_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LikeRecordRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LikeRecordRequest) ProtoMessage() {}

func (x *LikeRecordRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use LikeRecordRequest.ProtoReflect.Descriptor instead.
func (*LikeRecordRequest) Descriptor() ([]byte, []int) {
	return file_go_counter_proto_rawDescGZIP(), []int{4}
}

func (x *LikeRecordRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *LikeRecordRequest) GetArticleId() int64 {
	if x != nil {
		return x.ArticleId
	}
	return 0
}

type LikeRecordResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *LikeRecordResponse) Reset() {
	*x = LikeRecordResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_counter_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LikeRecordResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LikeRecordResponse) ProtoMessage() {}

func (x *LikeRecordResponse) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use LikeRecordResponse.ProtoReflect.Descriptor instead.
func (*LikeRecordResponse) Descriptor() ([]byte, []int) {
	return file_go_counter_proto_rawDescGZIP(), []int{5}
}

func (x *LikeRecordResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

// view the count of favorite
type FavoriteNumRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ArticleId int64 `protobuf:"varint,1,opt,name=article_id,json=articleId,proto3" json:"article_id,omitempty"`
}

func (x *FavoriteNumRequest) Reset() {
	*x = FavoriteNumRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_counter_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FavoriteNumRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FavoriteNumRequest) ProtoMessage() {}

func (x *FavoriteNumRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use FavoriteNumRequest.ProtoReflect.Descriptor instead.
func (*FavoriteNumRequest) Descriptor() ([]byte, []int) {
	return file_go_counter_proto_rawDescGZIP(), []int{6}
}

func (x *FavoriteNumRequest) GetArticleId() int64 {
	if x != nil {
		return x.ArticleId
	}
	return 0
}

type FavoriteNumResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count int64 `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *FavoriteNumResponse) Reset() {
	*x = FavoriteNumResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_counter_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FavoriteNumResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FavoriteNumResponse) ProtoMessage() {}

func (x *FavoriteNumResponse) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use FavoriteNumResponse.ProtoReflect.Descriptor instead.
func (*FavoriteNumResponse) Descriptor() ([]byte, []int) {
	return file_go_counter_proto_rawDescGZIP(), []int{7}
}

func (x *FavoriteNumResponse) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

// favorite record operations
type FavoriteRecordRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId    int64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	ArticleId int64 `protobuf:"varint,2,opt,name=article_id,json=articleId,proto3" json:"article_id,omitempty"`
}

func (x *FavoriteRecordRequest) Reset() {
	*x = FavoriteRecordRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_counter_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FavoriteRecordRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FavoriteRecordRequest) ProtoMessage() {}

func (x *FavoriteRecordRequest) ProtoReflect() protoreflect.Message {
	mi := &file_go_counter_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FavoriteRecordRequest.ProtoReflect.Descriptor instead.
func (*FavoriteRecordRequest) Descriptor() ([]byte, []int) {
	return file_go_counter_proto_rawDescGZIP(), []int{8}
}

func (x *FavoriteRecordRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *FavoriteRecordRequest) GetArticleId() int64 {
	if x != nil {
		return x.ArticleId
	}
	return 0
}

type FavoriteRecordResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *FavoriteRecordResponse) Reset() {
	*x = FavoriteRecordResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_counter_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FavoriteRecordResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FavoriteRecordResponse) ProtoMessage() {}

func (x *FavoriteRecordResponse) ProtoReflect() protoreflect.Message {
	mi := &file_go_counter_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FavoriteRecordResponse.ProtoReflect.Descriptor instead.
func (*FavoriteRecordResponse) Descriptor() ([]byte, []int) {
	return file_go_counter_proto_rawDescGZIP(), []int{9}
}

func (x *FavoriteRecordResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type ViewNumRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ArticleId int64 `protobuf:"varint,1,opt,name=article_id,json=articleId,proto3" json:"article_id,omitempty"`
}

func (x *ViewNumRequest) Reset() {
	*x = ViewNumRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_counter_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ViewNumRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ViewNumRequest) ProtoMessage() {}

func (x *ViewNumRequest) ProtoReflect() protoreflect.Message {
	mi := &file_go_counter_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ViewNumRequest.ProtoReflect.Descriptor instead.
func (*ViewNumRequest) Descriptor() ([]byte, []int) {
	return file_go_counter_proto_rawDescGZIP(), []int{10}
}

func (x *ViewNumRequest) GetArticleId() int64 {
	if x != nil {
		return x.ArticleId
	}
	return 0
}

type ViewNumResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *ViewNumResponse) Reset() {
	*x = ViewNumResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_counter_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ViewNumResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ViewNumResponse) ProtoMessage() {}

func (x *ViewNumResponse) ProtoReflect() protoreflect.Message {
	mi := &file_go_counter_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ViewNumResponse.ProtoReflect.Descriptor instead.
func (*ViewNumResponse) Descriptor() ([]byte, []int) {
	return file_go_counter_proto_rawDescGZIP(), []int{11}
}

func (x *ViewNumResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type ViewRecordRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId    int64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	ArticleId int64 `protobuf:"varint,2,opt,name=article_id,json=articleId,proto3" json:"article_id,omitempty"`
}

func (x *ViewRecordRequest) Reset() {
	*x = ViewRecordRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_counter_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ViewRecordRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ViewRecordRequest) ProtoMessage() {}

func (x *ViewRecordRequest) ProtoReflect() protoreflect.Message {
	mi := &file_go_counter_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ViewRecordRequest.ProtoReflect.Descriptor instead.
func (*ViewRecordRequest) Descriptor() ([]byte, []int) {
	return file_go_counter_proto_rawDescGZIP(), []int{12}
}

func (x *ViewRecordRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *ViewRecordRequest) GetArticleId() int64 {
	if x != nil {
		return x.ArticleId
	}
	return 0
}

type ViewRecordResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *ViewRecordResponse) Reset() {
	*x = ViewRecordResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_counter_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ViewRecordResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ViewRecordResponse) ProtoMessage() {}

func (x *ViewRecordResponse) ProtoReflect() protoreflect.Message {
	mi := &file_go_counter_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ViewRecordResponse.ProtoReflect.Descriptor instead.
func (*ViewRecordResponse) Descriptor() ([]byte, []int) {
	return file_go_counter_proto_rawDescGZIP(), []int{13}
}

func (x *ViewRecordResponse) GetSuccess() bool {
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
	0x04, 0x70, 0x6f, 0x6e, 0x67, 0x22, 0x2f, 0x0a, 0x0e, 0x4c, 0x69, 0x6b, 0x65, 0x4e, 0x75, 0x6d,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x72, 0x74, 0x69, 0x63,
	0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x61, 0x72, 0x74,
	0x69, 0x63, 0x6c, 0x65, 0x49, 0x64, 0x22, 0x27, 0x0a, 0x0f, 0x4c, 0x69, 0x6b, 0x65, 0x4e, 0x75,
	0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22,
	0x4b, 0x0a, 0x11, 0x4c, 0x69, 0x6b, 0x65, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a,
	0x0a, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x09, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x49, 0x64, 0x22, 0x2e, 0x0a, 0x12,
	0x4c, 0x69, 0x6b, 0x65, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x33, 0x0a, 0x12,
	0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x4e, 0x75, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x49,
	0x64, 0x22, 0x2b, 0x0a, 0x13, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x4e, 0x75, 0x6d,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x4f,
	0x0a, 0x15, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x49, 0x64, 0x22,
	0x32, 0x0a, 0x16, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x52, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x22, 0x2f, 0x0a, 0x0e, 0x56, 0x69, 0x65, 0x77, 0x4e, 0x75, 0x6d, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x61, 0x72, 0x74, 0x69, 0x63,
	0x6c, 0x65, 0x49, 0x64, 0x22, 0x2b, 0x0a, 0x0f, 0x56, 0x69, 0x65, 0x77, 0x4e, 0x75, 0x6d, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x22, 0x4b, 0x0a, 0x11, 0x56, 0x69, 0x65, 0x77, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x1d, 0x0a, 0x0a, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x09, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x49, 0x64, 0x22, 0x2e,
	0x0a, 0x12, 0x56, 0x69, 0x65, 0x77, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x32, 0x92,
	0x04, 0x0a, 0x0a, 0x47, 0x6f, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x12, 0x39, 0x0a,
	0x04, 0x50, 0x69, 0x6e, 0x67, 0x12, 0x17, 0x2e, 0x67, 0x6f, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x65, 0x72, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18,
	0x2e, 0x67, 0x6f, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x50, 0x69, 0x6e, 0x67,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x42, 0x0a, 0x07, 0x4c, 0x69, 0x6b, 0x65,
	0x4e, 0x75, 0x6d, 0x12, 0x1a, 0x2e, 0x67, 0x6f, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72,
	0x2e, 0x4c, 0x69, 0x6b, 0x65, 0x4e, 0x75, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1b, 0x2e, 0x67, 0x6f, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x4c, 0x69, 0x6b,
	0x65, 0x4e, 0x75, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4b, 0x0a, 0x0a,
	0x4c, 0x69, 0x6b, 0x65, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x5f,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x4c, 0x69, 0x6b, 0x65, 0x52, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x67, 0x6f, 0x5f, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x4c, 0x69, 0x6b, 0x65, 0x52, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4e, 0x0a, 0x0b, 0x46, 0x61, 0x76,
	0x6f, 0x72, 0x69, 0x74, 0x65, 0x4e, 0x75, 0x6d, 0x12, 0x1e, 0x2e, 0x67, 0x6f, 0x5f, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x4e, 0x75,
	0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x67, 0x6f, 0x5f, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x4e, 0x75,
	0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x57, 0x0a, 0x0e, 0x46, 0x61, 0x76,
	0x6f, 0x72, 0x69, 0x74, 0x65, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x21, 0x2e, 0x67, 0x6f,
	0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74,
	0x65, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22,
	0x2e, 0x67, 0x6f, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x46, 0x61, 0x76, 0x6f,
	0x72, 0x69, 0x74, 0x65, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x42, 0x0a, 0x07, 0x56, 0x69, 0x65, 0x77, 0x4e, 0x75, 0x6d, 0x12, 0x1a, 0x2e,
	0x67, 0x6f, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x56, 0x69, 0x65, 0x77, 0x4e,
	0x75, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x67, 0x6f, 0x5f, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x56, 0x69, 0x65, 0x77, 0x4e, 0x75, 0x6d, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4b, 0x0a, 0x0a, 0x56, 0x69, 0x65, 0x77, 0x52, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65,
	0x72, 0x2e, 0x56, 0x69, 0x65, 0x77, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x67, 0x6f, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72,
	0x2e, 0x56, 0x69, 0x65, 0x77, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x42, 0x0e, 0x5a, 0x0c, 0x2e, 0x2f, 0x67, 0x6f, 0x5f, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_go_counter_proto_msgTypes = make([]protoimpl.MessageInfo, 14)
var file_go_counter_proto_goTypes = []interface{}{
	(*PingRequest)(nil),            // 0: go_counter.PingRequest
	(*PingResponse)(nil),           // 1: go_counter.PingResponse
	(*LikeNumRequest)(nil),         // 2: go_counter.LikeNumRequest
	(*LikeNumResponse)(nil),        // 3: go_counter.LikeNumResponse
	(*LikeRecordRequest)(nil),      // 4: go_counter.LikeRecordRequest
	(*LikeRecordResponse)(nil),     // 5: go_counter.LikeRecordResponse
	(*FavoriteNumRequest)(nil),     // 6: go_counter.FavoriteNumRequest
	(*FavoriteNumResponse)(nil),    // 7: go_counter.FavoriteNumResponse
	(*FavoriteRecordRequest)(nil),  // 8: go_counter.FavoriteRecordRequest
	(*FavoriteRecordResponse)(nil), // 9: go_counter.FavoriteRecordResponse
	(*ViewNumRequest)(nil),         // 10: go_counter.ViewNumRequest
	(*ViewNumResponse)(nil),        // 11: go_counter.ViewNumResponse
	(*ViewRecordRequest)(nil),      // 12: go_counter.ViewRecordRequest
	(*ViewRecordResponse)(nil),     // 13: go_counter.ViewRecordResponse
}
var file_go_counter_proto_depIdxs = []int32{
	0,  // 0: go_counter.Go_counter.Ping:input_type -> go_counter.PingRequest
	2,  // 1: go_counter.Go_counter.LikeNum:input_type -> go_counter.LikeNumRequest
	4,  // 2: go_counter.Go_counter.LikeRecord:input_type -> go_counter.LikeRecordRequest
	6,  // 3: go_counter.Go_counter.FavoriteNum:input_type -> go_counter.FavoriteNumRequest
	8,  // 4: go_counter.Go_counter.FavoriteRecord:input_type -> go_counter.FavoriteRecordRequest
	10, // 5: go_counter.Go_counter.ViewNum:input_type -> go_counter.ViewNumRequest
	12, // 6: go_counter.Go_counter.ViewRecord:input_type -> go_counter.ViewRecordRequest
	1,  // 7: go_counter.Go_counter.Ping:output_type -> go_counter.PingResponse
	3,  // 8: go_counter.Go_counter.LikeNum:output_type -> go_counter.LikeNumResponse
	5,  // 9: go_counter.Go_counter.LikeRecord:output_type -> go_counter.LikeRecordResponse
	7,  // 10: go_counter.Go_counter.FavoriteNum:output_type -> go_counter.FavoriteNumResponse
	9,  // 11: go_counter.Go_counter.FavoriteRecord:output_type -> go_counter.FavoriteRecordResponse
	11, // 12: go_counter.Go_counter.ViewNum:output_type -> go_counter.ViewNumResponse
	13, // 13: go_counter.Go_counter.ViewRecord:output_type -> go_counter.ViewRecordResponse
	7,  // [7:14] is the sub-list for method output_type
	0,  // [0:7] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
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
			switch v := v.(*LikeNumRequest); i {
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
			switch v := v.(*LikeNumResponse); i {
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
			switch v := v.(*LikeRecordRequest); i {
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
			switch v := v.(*LikeRecordResponse); i {
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
			switch v := v.(*FavoriteNumRequest); i {
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
			switch v := v.(*FavoriteNumResponse); i {
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
		file_go_counter_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FavoriteRecordRequest); i {
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
		file_go_counter_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FavoriteRecordResponse); i {
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
		file_go_counter_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ViewNumRequest); i {
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
		file_go_counter_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ViewNumResponse); i {
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
		file_go_counter_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ViewRecordRequest); i {
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
		file_go_counter_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ViewRecordResponse); i {
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
			NumMessages:   14,
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
