// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.15.8
// source: comment/rpc/comment.proto

package comment

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

// --------------------------------Comment--------------------------------
type Comment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id"`               // id
	CreatedAt int64  `protobuf:"varint,2,opt,name=createdAt,proto3" json:"createdAt"` // createdAt
	Content   string `protobuf:"bytes,3,opt,name=content,proto3" json:"content"`      // content
	PostID    int64  `protobuf:"varint,4,opt,name=postID,proto3" json:"postID"`       // postID
}

func (x *Comment) Reset() {
	*x = Comment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comment_rpc_comment_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Comment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Comment) ProtoMessage() {}

func (x *Comment) ProtoReflect() protoreflect.Message {
	mi := &file_comment_rpc_comment_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Comment.ProtoReflect.Descriptor instead.
func (*Comment) Descriptor() ([]byte, []int) {
	return file_comment_rpc_comment_proto_rawDescGZIP(), []int{0}
}

func (x *Comment) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Comment) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *Comment) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *Comment) GetPostID() int64 {
	if x != nil {
		return x.PostID
	}
	return 0
}

type AddCommentReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CreatedAt int64  `protobuf:"varint,1,opt,name=createdAt,proto3" json:"createdAt"` // createdAt
	Content   string `protobuf:"bytes,2,opt,name=content,proto3" json:"content"`      // content
	PostID    int64  `protobuf:"varint,3,opt,name=postID,proto3" json:"postID"`       // postID
}

func (x *AddCommentReq) Reset() {
	*x = AddCommentReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comment_rpc_comment_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddCommentReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddCommentReq) ProtoMessage() {}

func (x *AddCommentReq) ProtoReflect() protoreflect.Message {
	mi := &file_comment_rpc_comment_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddCommentReq.ProtoReflect.Descriptor instead.
func (*AddCommentReq) Descriptor() ([]byte, []int) {
	return file_comment_rpc_comment_proto_rawDescGZIP(), []int{1}
}

func (x *AddCommentReq) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *AddCommentReq) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *AddCommentReq) GetPostID() int64 {
	if x != nil {
		return x.PostID
	}
	return 0
}

type AddCommentResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AddCommentResp) Reset() {
	*x = AddCommentResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comment_rpc_comment_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddCommentResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddCommentResp) ProtoMessage() {}

func (x *AddCommentResp) ProtoReflect() protoreflect.Message {
	mi := &file_comment_rpc_comment_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddCommentResp.ProtoReflect.Descriptor instead.
func (*AddCommentResp) Descriptor() ([]byte, []int) {
	return file_comment_rpc_comment_proto_rawDescGZIP(), []int{2}
}

type UpdateCommentReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id"`               // id
	CreatedAt int64  `protobuf:"varint,2,opt,name=createdAt,proto3" json:"createdAt"` // createdAt
	Content   string `protobuf:"bytes,3,opt,name=content,proto3" json:"content"`      // content
	PostID    int64  `protobuf:"varint,4,opt,name=postID,proto3" json:"postID"`       // postID
}

func (x *UpdateCommentReq) Reset() {
	*x = UpdateCommentReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comment_rpc_comment_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateCommentReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCommentReq) ProtoMessage() {}

func (x *UpdateCommentReq) ProtoReflect() protoreflect.Message {
	mi := &file_comment_rpc_comment_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCommentReq.ProtoReflect.Descriptor instead.
func (*UpdateCommentReq) Descriptor() ([]byte, []int) {
	return file_comment_rpc_comment_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateCommentReq) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateCommentReq) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *UpdateCommentReq) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *UpdateCommentReq) GetPostID() int64 {
	if x != nil {
		return x.PostID
	}
	return 0
}

type UpdateCommentResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateCommentResp) Reset() {
	*x = UpdateCommentResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comment_rpc_comment_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateCommentResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCommentResp) ProtoMessage() {}

func (x *UpdateCommentResp) ProtoReflect() protoreflect.Message {
	mi := &file_comment_rpc_comment_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCommentResp.ProtoReflect.Descriptor instead.
func (*UpdateCommentResp) Descriptor() ([]byte, []int) {
	return file_comment_rpc_comment_proto_rawDescGZIP(), []int{4}
}

type DelCommentReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id"` // id
}

func (x *DelCommentReq) Reset() {
	*x = DelCommentReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comment_rpc_comment_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DelCommentReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DelCommentReq) ProtoMessage() {}

func (x *DelCommentReq) ProtoReflect() protoreflect.Message {
	mi := &file_comment_rpc_comment_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DelCommentReq.ProtoReflect.Descriptor instead.
func (*DelCommentReq) Descriptor() ([]byte, []int) {
	return file_comment_rpc_comment_proto_rawDescGZIP(), []int{5}
}

func (x *DelCommentReq) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DelCommentResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DelCommentResp) Reset() {
	*x = DelCommentResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comment_rpc_comment_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DelCommentResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DelCommentResp) ProtoMessage() {}

func (x *DelCommentResp) ProtoReflect() protoreflect.Message {
	mi := &file_comment_rpc_comment_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DelCommentResp.ProtoReflect.Descriptor instead.
func (*DelCommentResp) Descriptor() ([]byte, []int) {
	return file_comment_rpc_comment_proto_rawDescGZIP(), []int{6}
}

type GetCommentByIdReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id"` // id
}

func (x *GetCommentByIdReq) Reset() {
	*x = GetCommentByIdReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comment_rpc_comment_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCommentByIdReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCommentByIdReq) ProtoMessage() {}

func (x *GetCommentByIdReq) ProtoReflect() protoreflect.Message {
	mi := &file_comment_rpc_comment_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCommentByIdReq.ProtoReflect.Descriptor instead.
func (*GetCommentByIdReq) Descriptor() ([]byte, []int) {
	return file_comment_rpc_comment_proto_rawDescGZIP(), []int{7}
}

func (x *GetCommentByIdReq) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetCommentByIdResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Comment *Comment `protobuf:"bytes,1,opt,name=comment,proto3" json:"comment"` // comment
}

func (x *GetCommentByIdResp) Reset() {
	*x = GetCommentByIdResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comment_rpc_comment_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCommentByIdResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCommentByIdResp) ProtoMessage() {}

func (x *GetCommentByIdResp) ProtoReflect() protoreflect.Message {
	mi := &file_comment_rpc_comment_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCommentByIdResp.ProtoReflect.Descriptor instead.
func (*GetCommentByIdResp) Descriptor() ([]byte, []int) {
	return file_comment_rpc_comment_proto_rawDescGZIP(), []int{8}
}

func (x *GetCommentByIdResp) GetComment() *Comment {
	if x != nil {
		return x.Comment
	}
	return nil
}

type SearchCommentReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page      int64  `protobuf:"varint,1,opt,name=page,proto3" json:"page"`           // page
	PageSize  int64  `protobuf:"varint,2,opt,name=pageSize,proto3" json:"pageSize"`   // pageSize
	Id        int64  `protobuf:"varint,3,opt,name=id,proto3" json:"id"`               // id
	CreatedAt int64  `protobuf:"varint,4,opt,name=createdAt,proto3" json:"createdAt"` // createdAt
	Content   string `protobuf:"bytes,5,opt,name=content,proto3" json:"content"`      // content
	PostID    int64  `protobuf:"varint,6,opt,name=postID,proto3" json:"postID"`       // postID
}

func (x *SearchCommentReq) Reset() {
	*x = SearchCommentReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comment_rpc_comment_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchCommentReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchCommentReq) ProtoMessage() {}

func (x *SearchCommentReq) ProtoReflect() protoreflect.Message {
	mi := &file_comment_rpc_comment_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchCommentReq.ProtoReflect.Descriptor instead.
func (*SearchCommentReq) Descriptor() ([]byte, []int) {
	return file_comment_rpc_comment_proto_rawDescGZIP(), []int{9}
}

func (x *SearchCommentReq) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *SearchCommentReq) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *SearchCommentReq) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *SearchCommentReq) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *SearchCommentReq) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *SearchCommentReq) GetPostID() int64 {
	if x != nil {
		return x.PostID
	}
	return 0
}

type SearchCommentResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Comment []*Comment `protobuf:"bytes,1,rep,name=comment,proto3" json:"comment"` // comment
}

func (x *SearchCommentResp) Reset() {
	*x = SearchCommentResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comment_rpc_comment_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchCommentResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchCommentResp) ProtoMessage() {}

func (x *SearchCommentResp) ProtoReflect() protoreflect.Message {
	mi := &file_comment_rpc_comment_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchCommentResp.ProtoReflect.Descriptor instead.
func (*SearchCommentResp) Descriptor() ([]byte, []int) {
	return file_comment_rpc_comment_proto_rawDescGZIP(), []int{10}
}

func (x *SearchCommentResp) GetComment() []*Comment {
	if x != nil {
		return x.Comment
	}
	return nil
}

var File_comment_rpc_comment_proto protoreflect.FileDescriptor

var file_comment_rpc_comment_proto_rawDesc = []byte{
	0x0a, 0x19, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x63, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x22, 0x69, 0x0a, 0x07, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x1c, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x18, 0x0a,
	0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x6f, 0x73, 0x74, 0x49,
	0x44, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x70, 0x6f, 0x73, 0x74, 0x49, 0x44, 0x22,
	0x5f, 0x0a, 0x0d, 0x41, 0x64, 0x64, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71,
	0x12, 0x1c, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x18,
	0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x6f, 0x73, 0x74,
	0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x70, 0x6f, 0x73, 0x74, 0x49, 0x44,
	0x22, 0x10, 0x0a, 0x0e, 0x41, 0x64, 0x64, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x22, 0x72, 0x0a, 0x10, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x16,
	0x0a, 0x06, 0x70, 0x6f, 0x73, 0x74, 0x49, 0x44, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06,
	0x70, 0x6f, 0x73, 0x74, 0x49, 0x44, 0x22, 0x13, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x22, 0x1f, 0x0a, 0x0d, 0x44,
	0x65, 0x6c, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x10, 0x0a, 0x0e,
	0x44, 0x65, 0x6c, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x22, 0x23,
	0x0a, 0x11, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x42, 0x79, 0x49, 0x64,
	0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x02, 0x69, 0x64, 0x22, 0x40, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x74, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x12, 0x2a, 0x0a, 0x07, 0x63, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x07, 0x63, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0xa2, 0x01, 0x0a, 0x10, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61,
	0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x6f, 0x73, 0x74, 0x49, 0x44, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x06, 0x70, 0x6f, 0x73, 0x74, 0x49, 0x44, 0x22, 0x3f, 0x0a, 0x11, 0x53, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12,
	0x2a, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x32, 0xe2, 0x02, 0x0a, 0x07,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3d, 0x0a, 0x0a, 0x41, 0x64, 0x64, 0x43, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x16, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e,
	0x41, 0x64, 0x64, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x17, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x41, 0x64, 0x64, 0x43, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x46, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x19, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x74, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52,
	0x65, 0x71, 0x1a, 0x1a, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x3d,
	0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x16, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x44, 0x65, 0x6c, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x74, 0x52, 0x65, 0x71, 0x1a, 0x17, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x44,
	0x65, 0x6c, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x49, 0x0a,
	0x0e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x42, 0x79, 0x49, 0x64, 0x12,
	0x1a, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x1a, 0x1b, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74,
	0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x12, 0x46, 0x0a, 0x0d, 0x53, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x19, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x74, 0x52, 0x65, 0x71, 0x1a, 0x1a, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x53,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x42, 0x0b, 0x5a, 0x09, 0x2e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_comment_rpc_comment_proto_rawDescOnce sync.Once
	file_comment_rpc_comment_proto_rawDescData = file_comment_rpc_comment_proto_rawDesc
)

func file_comment_rpc_comment_proto_rawDescGZIP() []byte {
	file_comment_rpc_comment_proto_rawDescOnce.Do(func() {
		file_comment_rpc_comment_proto_rawDescData = protoimpl.X.CompressGZIP(file_comment_rpc_comment_proto_rawDescData)
	})
	return file_comment_rpc_comment_proto_rawDescData
}

var file_comment_rpc_comment_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_comment_rpc_comment_proto_goTypes = []interface{}{
	(*Comment)(nil),            // 0: comment.Comment
	(*AddCommentReq)(nil),      // 1: comment.AddCommentReq
	(*AddCommentResp)(nil),     // 2: comment.AddCommentResp
	(*UpdateCommentReq)(nil),   // 3: comment.UpdateCommentReq
	(*UpdateCommentResp)(nil),  // 4: comment.UpdateCommentResp
	(*DelCommentReq)(nil),      // 5: comment.DelCommentReq
	(*DelCommentResp)(nil),     // 6: comment.DelCommentResp
	(*GetCommentByIdReq)(nil),  // 7: comment.GetCommentByIdReq
	(*GetCommentByIdResp)(nil), // 8: comment.GetCommentByIdResp
	(*SearchCommentReq)(nil),   // 9: comment.SearchCommentReq
	(*SearchCommentResp)(nil),  // 10: comment.SearchCommentResp
}
var file_comment_rpc_comment_proto_depIdxs = []int32{
	0,  // 0: comment.GetCommentByIdResp.comment:type_name -> comment.Comment
	0,  // 1: comment.SearchCommentResp.comment:type_name -> comment.Comment
	1,  // 2: comment.service.AddComment:input_type -> comment.AddCommentReq
	3,  // 3: comment.service.UpdateComment:input_type -> comment.UpdateCommentReq
	5,  // 4: comment.service.DelComment:input_type -> comment.DelCommentReq
	7,  // 5: comment.service.GetCommentById:input_type -> comment.GetCommentByIdReq
	9,  // 6: comment.service.SearchComment:input_type -> comment.SearchCommentReq
	2,  // 7: comment.service.AddComment:output_type -> comment.AddCommentResp
	4,  // 8: comment.service.UpdateComment:output_type -> comment.UpdateCommentResp
	6,  // 9: comment.service.DelComment:output_type -> comment.DelCommentResp
	8,  // 10: comment.service.GetCommentById:output_type -> comment.GetCommentByIdResp
	10, // 11: comment.service.SearchComment:output_type -> comment.SearchCommentResp
	7,  // [7:12] is the sub-list for method output_type
	2,  // [2:7] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_comment_rpc_comment_proto_init() }
func file_comment_rpc_comment_proto_init() {
	if File_comment_rpc_comment_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_comment_rpc_comment_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Comment); i {
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
		file_comment_rpc_comment_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddCommentReq); i {
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
		file_comment_rpc_comment_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddCommentResp); i {
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
		file_comment_rpc_comment_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateCommentReq); i {
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
		file_comment_rpc_comment_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateCommentResp); i {
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
		file_comment_rpc_comment_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DelCommentReq); i {
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
		file_comment_rpc_comment_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DelCommentResp); i {
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
		file_comment_rpc_comment_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCommentByIdReq); i {
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
		file_comment_rpc_comment_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCommentByIdResp); i {
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
		file_comment_rpc_comment_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchCommentReq); i {
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
		file_comment_rpc_comment_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchCommentResp); i {
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
			RawDescriptor: file_comment_rpc_comment_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_comment_rpc_comment_proto_goTypes,
		DependencyIndexes: file_comment_rpc_comment_proto_depIdxs,
		MessageInfos:      file_comment_rpc_comment_proto_msgTypes,
	}.Build()
	File_comment_rpc_comment_proto = out.File
	file_comment_rpc_comment_proto_rawDesc = nil
	file_comment_rpc_comment_proto_goTypes = nil
	file_comment_rpc_comment_proto_depIdxs = nil
}
