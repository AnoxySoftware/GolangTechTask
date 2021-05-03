// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.8
// source: api/service.proto

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

type Voteable struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid     string   `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Question string   `protobuf:"bytes,2,opt,name=question,proto3" json:"question,omitempty"`
	Answers  []string `protobuf:"bytes,3,rep,name=answers,proto3" json:"answers,omitempty"`
}

func (x *Voteable) Reset() {
	*x = Voteable{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Voteable) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Voteable) ProtoMessage() {}

func (x *Voteable) ProtoReflect() protoreflect.Message {
	mi := &file_api_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Voteable.ProtoReflect.Descriptor instead.
func (*Voteable) Descriptor() ([]byte, []int) {
	return file_api_service_proto_rawDescGZIP(), []int{0}
}

func (x *Voteable) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *Voteable) GetQuestion() string {
	if x != nil {
		return x.Question
	}
	return ""
}

func (x *Voteable) GetAnswers() []string {
	if x != nil {
		return x.Answers
	}
	return nil
}

type CreateVoteableRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Question string   `protobuf:"bytes,1,opt,name=question,proto3" json:"question,omitempty"`
	Answers  []string `protobuf:"bytes,2,rep,name=answers,proto3" json:"answers,omitempty"`
}

func (x *CreateVoteableRequest) Reset() {
	*x = CreateVoteableRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateVoteableRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateVoteableRequest) ProtoMessage() {}

func (x *CreateVoteableRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateVoteableRequest.ProtoReflect.Descriptor instead.
func (*CreateVoteableRequest) Descriptor() ([]byte, []int) {
	return file_api_service_proto_rawDescGZIP(), []int{1}
}

func (x *CreateVoteableRequest) GetQuestion() string {
	if x != nil {
		return x.Question
	}
	return ""
}

func (x *CreateVoteableRequest) GetAnswers() []string {
	if x != nil {
		return x.Answers
	}
	return nil
}

type CreateVoteableResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
}

func (x *CreateVoteableResponse) Reset() {
	*x = CreateVoteableResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateVoteableResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateVoteableResponse) ProtoMessage() {}

func (x *CreateVoteableResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateVoteableResponse.ProtoReflect.Descriptor instead.
func (*CreateVoteableResponse) Descriptor() ([]byte, []int) {
	return file_api_service_proto_rawDescGZIP(), []int{2}
}

func (x *CreateVoteableResponse) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

type ListVoteablesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PageSize  int64  `protobuf:"varint,1,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	PageToken string `protobuf:"bytes,2,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
}

func (x *ListVoteablesRequest) Reset() {
	*x = ListVoteablesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListVoteablesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListVoteablesRequest) ProtoMessage() {}

func (x *ListVoteablesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListVoteablesRequest.ProtoReflect.Descriptor instead.
func (*ListVoteablesRequest) Descriptor() ([]byte, []int) {
	return file_api_service_proto_rawDescGZIP(), []int{3}
}

func (x *ListVoteablesRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListVoteablesRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

type ListVoteablesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Votables      []*Voteable `protobuf:"bytes,1,rep,name=votables,proto3" json:"votables,omitempty"`
	NextPageToken *string     `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3,oneof" json:"next_page_token,omitempty"`
}

func (x *ListVoteablesResponse) Reset() {
	*x = ListVoteablesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListVoteablesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListVoteablesResponse) ProtoMessage() {}

func (x *ListVoteablesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListVoteablesResponse.ProtoReflect.Descriptor instead.
func (*ListVoteablesResponse) Descriptor() ([]byte, []int) {
	return file_api_service_proto_rawDescGZIP(), []int{4}
}

func (x *ListVoteablesResponse) GetVotables() []*Voteable {
	if x != nil {
		return x.Votables
	}
	return nil
}

func (x *ListVoteablesResponse) GetNextPageToken() string {
	if x != nil && x.NextPageToken != nil {
		return *x.NextPageToken
	}
	return ""
}

type CastVoteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid        string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	AnswerIndex int64  `protobuf:"varint,2,opt,name=answer_index,json=answerIndex,proto3" json:"answer_index,omitempty"`
}

func (x *CastVoteRequest) Reset() {
	*x = CastVoteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CastVoteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CastVoteRequest) ProtoMessage() {}

func (x *CastVoteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CastVoteRequest.ProtoReflect.Descriptor instead.
func (*CastVoteRequest) Descriptor() ([]byte, []int) {
	return file_api_service_proto_rawDescGZIP(), []int{5}
}

func (x *CastVoteRequest) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *CastVoteRequest) GetAnswerIndex() int64 {
	if x != nil {
		return x.AnswerIndex
	}
	return 0
}

type CastVoteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CastVoteResponse) Reset() {
	*x = CastVoteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CastVoteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CastVoteResponse) ProtoMessage() {}

func (x *CastVoteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CastVoteResponse.ProtoReflect.Descriptor instead.
func (*CastVoteResponse) Descriptor() ([]byte, []int) {
	return file_api_service_proto_rawDescGZIP(), []int{6}
}

var File_api_service_proto protoreflect.FileDescriptor

var file_api_service_proto_rawDesc = []byte{
	0x0a, 0x11, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x54, 0x0a, 0x08, 0x56, 0x6f, 0x74, 0x65, 0x61, 0x62, 0x6c, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75,
	0x75, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x18, 0x0a, 0x07, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x07, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x73, 0x22, 0x4d, 0x0a, 0x15, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x56, 0x6f, 0x74, 0x65, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18,
	0x0a, 0x07, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x07, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x73, 0x22, 0x2c, 0x0a, 0x16, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x56, 0x6f, 0x74, 0x65, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x22, 0x52, 0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74, 0x56, 0x6f,
	0x74, 0x65, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b,
	0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70,
	0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x70, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x7f, 0x0a, 0x15, 0x4c, 0x69,
	0x73, 0x74, 0x56, 0x6f, 0x74, 0x65, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x08, 0x76, 0x6f, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x56, 0x6f, 0x74, 0x65, 0x61, 0x62, 0x6c, 0x65,
	0x52, 0x08, 0x76, 0x6f, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x12, 0x2b, 0x0a, 0x0f, 0x6e, 0x65,
	0x78, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0d, 0x6e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x88, 0x01, 0x01, 0x42, 0x12, 0x0a, 0x10, 0x5f, 0x6e, 0x65, 0x78, 0x74,
	0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x48, 0x0a, 0x0f, 0x43,
	0x61, 0x73, 0x74, 0x56, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75,
	0x69, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x5f, 0x69, 0x6e, 0x64,
	0x65, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72,
	0x49, 0x6e, 0x64, 0x65, 0x78, 0x22, 0x12, 0x0a, 0x10, 0x43, 0x61, 0x73, 0x74, 0x56, 0x6f, 0x74,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xc9, 0x01, 0x0a, 0x0d, 0x56, 0x6f,
	0x74, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x43, 0x0a, 0x0e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x56, 0x6f, 0x74, 0x65, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x16, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x56, 0x6f, 0x74, 0x65, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x56, 0x6f,
	0x74, 0x65, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x40, 0x0a, 0x0d, 0x4c, 0x69, 0x73, 0x74, 0x56, 0x6f, 0x74, 0x65, 0x61, 0x62, 0x6c, 0x65,
	0x73, 0x12, 0x15, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x56, 0x6f, 0x74, 0x65, 0x61, 0x62, 0x6c, 0x65,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x56,
	0x6f, 0x74, 0x65, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x31, 0x0a, 0x08, 0x43, 0x61, 0x73, 0x74, 0x56, 0x6f, 0x74, 0x65, 0x12, 0x10,
	0x2e, 0x43, 0x61, 0x73, 0x74, 0x56, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x11, 0x2e, 0x43, 0x61, 0x73, 0x74, 0x56, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x2a, 0x5a, 0x28, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x75, 0x66, 0x66, 0x75, 0x70, 0x2f, 0x47, 0x6f, 0x6c, 0x61, 0x6e,
	0x67, 0x54, 0x65, 0x63, 0x68, 0x54, 0x61, 0x73, 0x6b, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70,
	0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_service_proto_rawDescOnce sync.Once
	file_api_service_proto_rawDescData = file_api_service_proto_rawDesc
)

func file_api_service_proto_rawDescGZIP() []byte {
	file_api_service_proto_rawDescOnce.Do(func() {
		file_api_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_service_proto_rawDescData)
	})
	return file_api_service_proto_rawDescData
}

var file_api_service_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_api_service_proto_goTypes = []interface{}{
	(*Voteable)(nil),               // 0: Voteable
	(*CreateVoteableRequest)(nil),  // 1: CreateVoteableRequest
	(*CreateVoteableResponse)(nil), // 2: CreateVoteableResponse
	(*ListVoteablesRequest)(nil),   // 3: ListVoteablesRequest
	(*ListVoteablesResponse)(nil),  // 4: ListVoteablesResponse
	(*CastVoteRequest)(nil),        // 5: CastVoteRequest
	(*CastVoteResponse)(nil),       // 6: CastVoteResponse
}
var file_api_service_proto_depIdxs = []int32{
	0, // 0: ListVoteablesResponse.votables:type_name -> Voteable
	1, // 1: VotingService.CreateVoteable:input_type -> CreateVoteableRequest
	3, // 2: VotingService.ListVoteables:input_type -> ListVoteablesRequest
	5, // 3: VotingService.CastVote:input_type -> CastVoteRequest
	2, // 4: VotingService.CreateVoteable:output_type -> CreateVoteableResponse
	4, // 5: VotingService.ListVoteables:output_type -> ListVoteablesResponse
	6, // 6: VotingService.CastVote:output_type -> CastVoteResponse
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_api_service_proto_init() }
func file_api_service_proto_init() {
	if File_api_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Voteable); i {
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
		file_api_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateVoteableRequest); i {
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
		file_api_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateVoteableResponse); i {
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
		file_api_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListVoteablesRequest); i {
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
		file_api_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListVoteablesResponse); i {
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
		file_api_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CastVoteRequest); i {
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
		file_api_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CastVoteResponse); i {
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
	file_api_service_proto_msgTypes[4].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_service_proto_goTypes,
		DependencyIndexes: file_api_service_proto_depIdxs,
		MessageInfos:      file_api_service_proto_msgTypes,
	}.Build()
	File_api_service_proto = out.File
	file_api_service_proto_rawDesc = nil
	file_api_service_proto_goTypes = nil
	file_api_service_proto_depIdxs = nil
}
