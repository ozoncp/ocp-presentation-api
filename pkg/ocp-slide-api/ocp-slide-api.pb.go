// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.2
// source: api/ocp-slide-api/ocp-slide-api.proto

package ocp_slide_api

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type SlideType int32

const (
	SlideType_Question SlideType = 0
	SlideType_Video    SlideType = 1
	SlideType_Document SlideType = 2
	SlideType_Task     SlideType = 3
)

// Enum value maps for SlideType.
var (
	SlideType_name = map[int32]string{
		0: "Question",
		1: "Video",
		2: "Document",
		3: "Task",
	}
	SlideType_value = map[string]int32{
		"Question": 0,
		"Video":    1,
		"Document": 2,
		"Task":     3,
	}
)

func (x SlideType) Enum() *SlideType {
	p := new(SlideType)
	*p = x
	return p
}

func (x SlideType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SlideType) Descriptor() protoreflect.EnumDescriptor {
	return file_api_ocp_slide_api_ocp_slide_api_proto_enumTypes[0].Descriptor()
}

func (SlideType) Type() protoreflect.EnumType {
	return &file_api_ocp_slide_api_ocp_slide_api_proto_enumTypes[0]
}

func (x SlideType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SlideType.Descriptor instead.
func (SlideType) EnumDescriptor() ([]byte, []int) {
	return file_api_ocp_slide_api_ocp_slide_api_proto_rawDescGZIP(), []int{0}
}

type CreateSlideV1Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PresentationId uint64    `protobuf:"varint,1,opt,name=presentation_id,json=presentationId,proto3" json:"presentation_id,omitempty"`
	Number         uint64    `protobuf:"varint,2,opt,name=number,proto3" json:"number,omitempty"`
	Type           SlideType `protobuf:"varint,3,opt,name=type,proto3,enum=ocp.slide.api.SlideType" json:"type,omitempty"`
}

func (x *CreateSlideV1Request) Reset() {
	*x = CreateSlideV1Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ocp_slide_api_ocp_slide_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSlideV1Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSlideV1Request) ProtoMessage() {}

func (x *CreateSlideV1Request) ProtoReflect() protoreflect.Message {
	mi := &file_api_ocp_slide_api_ocp_slide_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSlideV1Request.ProtoReflect.Descriptor instead.
func (*CreateSlideV1Request) Descriptor() ([]byte, []int) {
	return file_api_ocp_slide_api_ocp_slide_api_proto_rawDescGZIP(), []int{0}
}

func (x *CreateSlideV1Request) GetPresentationId() uint64 {
	if x != nil {
		return x.PresentationId
	}
	return 0
}

func (x *CreateSlideV1Request) GetNumber() uint64 {
	if x != nil {
		return x.Number
	}
	return 0
}

func (x *CreateSlideV1Request) GetType() SlideType {
	if x != nil {
		return x.Type
	}
	return SlideType_Question
}

type CreateSlideV1Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SlideId uint64 `protobuf:"varint,1,opt,name=slide_id,json=slideId,proto3" json:"slide_id,omitempty"`
}

func (x *CreateSlideV1Response) Reset() {
	*x = CreateSlideV1Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ocp_slide_api_ocp_slide_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSlideV1Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSlideV1Response) ProtoMessage() {}

func (x *CreateSlideV1Response) ProtoReflect() protoreflect.Message {
	mi := &file_api_ocp_slide_api_ocp_slide_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSlideV1Response.ProtoReflect.Descriptor instead.
func (*CreateSlideV1Response) Descriptor() ([]byte, []int) {
	return file_api_ocp_slide_api_ocp_slide_api_proto_rawDescGZIP(), []int{1}
}

func (x *CreateSlideV1Response) GetSlideId() uint64 {
	if x != nil {
		return x.SlideId
	}
	return 0
}

type DescribeSlideV1Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SlideId uint64 `protobuf:"varint,1,opt,name=slide_id,json=slideId,proto3" json:"slide_id,omitempty"`
}

func (x *DescribeSlideV1Request) Reset() {
	*x = DescribeSlideV1Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ocp_slide_api_ocp_slide_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescribeSlideV1Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeSlideV1Request) ProtoMessage() {}

func (x *DescribeSlideV1Request) ProtoReflect() protoreflect.Message {
	mi := &file_api_ocp_slide_api_ocp_slide_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeSlideV1Request.ProtoReflect.Descriptor instead.
func (*DescribeSlideV1Request) Descriptor() ([]byte, []int) {
	return file_api_ocp_slide_api_ocp_slide_api_proto_rawDescGZIP(), []int{2}
}

func (x *DescribeSlideV1Request) GetSlideId() uint64 {
	if x != nil {
		return x.SlideId
	}
	return 0
}

type DescribeSlideV1Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Slide *Slide `protobuf:"bytes,1,opt,name=slide,proto3" json:"slide,omitempty"`
}

func (x *DescribeSlideV1Response) Reset() {
	*x = DescribeSlideV1Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ocp_slide_api_ocp_slide_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescribeSlideV1Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeSlideV1Response) ProtoMessage() {}

func (x *DescribeSlideV1Response) ProtoReflect() protoreflect.Message {
	mi := &file_api_ocp_slide_api_ocp_slide_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeSlideV1Response.ProtoReflect.Descriptor instead.
func (*DescribeSlideV1Response) Descriptor() ([]byte, []int) {
	return file_api_ocp_slide_api_ocp_slide_api_proto_rawDescGZIP(), []int{3}
}

func (x *DescribeSlideV1Response) GetSlide() *Slide {
	if x != nil {
		return x.Slide
	}
	return nil
}

type ListSlidesV1Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Limit  uint64 `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset uint64 `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
}

func (x *ListSlidesV1Request) Reset() {
	*x = ListSlidesV1Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ocp_slide_api_ocp_slide_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListSlidesV1Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListSlidesV1Request) ProtoMessage() {}

func (x *ListSlidesV1Request) ProtoReflect() protoreflect.Message {
	mi := &file_api_ocp_slide_api_ocp_slide_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListSlidesV1Request.ProtoReflect.Descriptor instead.
func (*ListSlidesV1Request) Descriptor() ([]byte, []int) {
	return file_api_ocp_slide_api_ocp_slide_api_proto_rawDescGZIP(), []int{4}
}

func (x *ListSlidesV1Request) GetLimit() uint64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *ListSlidesV1Request) GetOffset() uint64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

type ListSlidesV1Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Slides []*Slide `protobuf:"bytes,1,rep,name=slides,proto3" json:"slides,omitempty"`
}

func (x *ListSlidesV1Response) Reset() {
	*x = ListSlidesV1Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ocp_slide_api_ocp_slide_api_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListSlidesV1Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListSlidesV1Response) ProtoMessage() {}

func (x *ListSlidesV1Response) ProtoReflect() protoreflect.Message {
	mi := &file_api_ocp_slide_api_ocp_slide_api_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListSlidesV1Response.ProtoReflect.Descriptor instead.
func (*ListSlidesV1Response) Descriptor() ([]byte, []int) {
	return file_api_ocp_slide_api_ocp_slide_api_proto_rawDescGZIP(), []int{5}
}

func (x *ListSlidesV1Response) GetSlides() []*Slide {
	if x != nil {
		return x.Slides
	}
	return nil
}

type RemoveSlideV1Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SlideId uint64 `protobuf:"varint,1,opt,name=slide_id,json=slideId,proto3" json:"slide_id,omitempty"`
}

func (x *RemoveSlideV1Request) Reset() {
	*x = RemoveSlideV1Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ocp_slide_api_ocp_slide_api_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveSlideV1Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveSlideV1Request) ProtoMessage() {}

func (x *RemoveSlideV1Request) ProtoReflect() protoreflect.Message {
	mi := &file_api_ocp_slide_api_ocp_slide_api_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveSlideV1Request.ProtoReflect.Descriptor instead.
func (*RemoveSlideV1Request) Descriptor() ([]byte, []int) {
	return file_api_ocp_slide_api_ocp_slide_api_proto_rawDescGZIP(), []int{6}
}

func (x *RemoveSlideV1Request) GetSlideId() uint64 {
	if x != nil {
		return x.SlideId
	}
	return 0
}

type RemoveSlideV1Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Found bool `protobuf:"varint,1,opt,name=found,proto3" json:"found,omitempty"`
}

func (x *RemoveSlideV1Response) Reset() {
	*x = RemoveSlideV1Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ocp_slide_api_ocp_slide_api_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveSlideV1Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveSlideV1Response) ProtoMessage() {}

func (x *RemoveSlideV1Response) ProtoReflect() protoreflect.Message {
	mi := &file_api_ocp_slide_api_ocp_slide_api_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveSlideV1Response.ProtoReflect.Descriptor instead.
func (*RemoveSlideV1Response) Descriptor() ([]byte, []int) {
	return file_api_ocp_slide_api_ocp_slide_api_proto_rawDescGZIP(), []int{7}
}

func (x *RemoveSlideV1Response) GetFound() bool {
	if x != nil {
		return x.Found
	}
	return false
}

type Slide struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SlideId        uint64    `protobuf:"varint,1,opt,name=slide_id,json=slideId,proto3" json:"slide_id,omitempty"`
	PresentationId uint64    `protobuf:"varint,2,opt,name=presentation_id,json=presentationId,proto3" json:"presentation_id,omitempty"`
	Number         uint64    `protobuf:"varint,3,opt,name=number,proto3" json:"number,omitempty"`
	Type           SlideType `protobuf:"varint,4,opt,name=type,proto3,enum=ocp.slide.api.SlideType" json:"type,omitempty"`
}

func (x *Slide) Reset() {
	*x = Slide{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ocp_slide_api_ocp_slide_api_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Slide) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Slide) ProtoMessage() {}

func (x *Slide) ProtoReflect() protoreflect.Message {
	mi := &file_api_ocp_slide_api_ocp_slide_api_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Slide.ProtoReflect.Descriptor instead.
func (*Slide) Descriptor() ([]byte, []int) {
	return file_api_ocp_slide_api_ocp_slide_api_proto_rawDescGZIP(), []int{8}
}

func (x *Slide) GetSlideId() uint64 {
	if x != nil {
		return x.SlideId
	}
	return 0
}

func (x *Slide) GetPresentationId() uint64 {
	if x != nil {
		return x.PresentationId
	}
	return 0
}

func (x *Slide) GetNumber() uint64 {
	if x != nil {
		return x.Number
	}
	return 0
}

func (x *Slide) GetType() SlideType {
	if x != nil {
		return x.Type
	}
	return SlideType_Question
}

var File_api_ocp_slide_api_ocp_slide_api_proto protoreflect.FileDescriptor

var file_api_ocp_slide_api_ocp_slide_api_proto_rawDesc = []byte{
	0x0a, 0x25, 0x61, 0x70, 0x69, 0x2f, 0x6f, 0x63, 0x70, 0x2d, 0x73, 0x6c, 0x69, 0x64, 0x65, 0x2d,
	0x61, 0x70, 0x69, 0x2f, 0x6f, 0x63, 0x70, 0x2d, 0x73, 0x6c, 0x69, 0x64, 0x65, 0x2d, 0x61, 0x70,
	0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x6f, 0x63, 0x70, 0x2e, 0x73, 0x6c, 0x69,
	0x64, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x41, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x97, 0x01, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x53, 0x6c, 0x69, 0x64, 0x65, 0x56, 0x31, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x30, 0x0a, 0x0f, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x32, 0x02,
	0x20, 0x00, 0x52, 0x0e, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x49, 0x64, 0x12, 0x1f, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x04, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x32, 0x02, 0x20, 0x00, 0x52, 0x06, 0x6e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x12, 0x2c, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x18, 0x2e, 0x6f, 0x63, 0x70, 0x2e, 0x73, 0x6c, 0x69, 0x64, 0x65, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x53, 0x6c, 0x69, 0x64, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x22, 0x32, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x6c, 0x69, 0x64, 0x65,
	0x56, 0x31, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x6c,
	0x69, 0x64, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x73, 0x6c,
	0x69, 0x64, 0x65, 0x49, 0x64, 0x22, 0x3c, 0x0a, 0x16, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62,
	0x65, 0x53, 0x6c, 0x69, 0x64, 0x65, 0x56, 0x31, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x22, 0x0a, 0x08, 0x73, 0x6c, 0x69, 0x64, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x04, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x32, 0x02, 0x20, 0x00, 0x52, 0x07, 0x73, 0x6c, 0x69, 0x64,
	0x65, 0x49, 0x64, 0x22, 0x45, 0x0a, 0x17, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x53,
	0x6c, 0x69, 0x64, 0x65, 0x56, 0x31, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a,
	0x0a, 0x05, 0x73, 0x6c, 0x69, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e,
	0x6f, 0x63, 0x70, 0x2e, 0x73, 0x6c, 0x69, 0x64, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x6c,
	0x69, 0x64, 0x65, 0x52, 0x05, 0x73, 0x6c, 0x69, 0x64, 0x65, 0x22, 0x43, 0x0a, 0x13, 0x4c, 0x69,
	0x73, 0x74, 0x53, 0x6c, 0x69, 0x64, 0x65, 0x73, 0x56, 0x31, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x22,
	0x44, 0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x6c, 0x69, 0x64, 0x65, 0x73, 0x56, 0x31, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a, 0x06, 0x73, 0x6c, 0x69, 0x64, 0x65,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6f, 0x63, 0x70, 0x2e, 0x73, 0x6c,
	0x69, 0x64, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x6c, 0x69, 0x64, 0x65, 0x52, 0x06, 0x73,
	0x6c, 0x69, 0x64, 0x65, 0x73, 0x22, 0x3a, 0x0a, 0x14, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x53,
	0x6c, 0x69, 0x64, 0x65, 0x56, 0x31, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a,
	0x08, 0x73, 0x6c, 0x69, 0x64, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x42,
	0x07, 0xfa, 0x42, 0x04, 0x32, 0x02, 0x20, 0x00, 0x52, 0x07, 0x73, 0x6c, 0x69, 0x64, 0x65, 0x49,
	0x64, 0x22, 0x2d, 0x0a, 0x15, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x53, 0x6c, 0x69, 0x64, 0x65,
	0x56, 0x31, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x6f,
	0x75, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x66, 0x6f, 0x75, 0x6e, 0x64,
	0x22, 0xac, 0x01, 0x0a, 0x05, 0x53, 0x6c, 0x69, 0x64, 0x65, 0x12, 0x22, 0x0a, 0x08, 0x73, 0x6c,
	0x69, 0x64, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x42, 0x07, 0xfa, 0x42,
	0x04, 0x32, 0x02, 0x20, 0x00, 0x52, 0x07, 0x73, 0x6c, 0x69, 0x64, 0x65, 0x49, 0x64, 0x12, 0x30,
	0x0a, 0x0f, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x32, 0x02, 0x20, 0x00,
	0x52, 0x0e, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64,
	0x12, 0x1f, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04,
	0x42, 0x07, 0xfa, 0x42, 0x04, 0x32, 0x02, 0x20, 0x00, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x12, 0x2c, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x18, 0x2e, 0x6f, 0x63, 0x70, 0x2e, 0x73, 0x6c, 0x69, 0x64, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x53, 0x6c, 0x69, 0x64, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x2a,
	0x3c, 0x0a, 0x09, 0x53, 0x6c, 0x69, 0x64, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0c, 0x0a, 0x08,
	0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x56, 0x69,
	0x64, 0x65, 0x6f, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e,
	0x74, 0x10, 0x02, 0x12, 0x08, 0x0a, 0x04, 0x54, 0x61, 0x73, 0x6b, 0x10, 0x03, 0x32, 0xe3, 0x03,
	0x0a, 0x08, 0x53, 0x6c, 0x69, 0x64, 0x65, 0x41, 0x50, 0x49, 0x12, 0x6e, 0x0a, 0x0d, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x53, 0x6c, 0x69, 0x64, 0x65, 0x56, 0x31, 0x12, 0x23, 0x2e, 0x6f, 0x63,
	0x70, 0x2e, 0x73, 0x6c, 0x69, 0x64, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x53, 0x6c, 0x69, 0x64, 0x65, 0x56, 0x31, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x24, 0x2e, 0x6f, 0x63, 0x70, 0x2e, 0x73, 0x6c, 0x69, 0x64, 0x65, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x6c, 0x69, 0x64, 0x65, 0x56, 0x31, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x12, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0c, 0x22, 0x0a,
	0x2f, 0x76, 0x31, 0x2f, 0x73, 0x6c, 0x69, 0x64, 0x65, 0x73, 0x12, 0x7f, 0x0a, 0x0f, 0x44, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x53, 0x6c, 0x69, 0x64, 0x65, 0x56, 0x31, 0x12, 0x25, 0x2e,
	0x6f, 0x63, 0x70, 0x2e, 0x73, 0x6c, 0x69, 0x64, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x44, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x53, 0x6c, 0x69, 0x64, 0x65, 0x56, 0x31, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x6f, 0x63, 0x70, 0x2e, 0x73, 0x6c, 0x69, 0x64, 0x65,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x53, 0x6c, 0x69,
	0x64, 0x65, 0x56, 0x31, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1d, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x17, 0x12, 0x15, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x6c, 0x69, 0x64, 0x65, 0x73,
	0x2f, 0x7b, 0x73, 0x6c, 0x69, 0x64, 0x65, 0x5f, 0x69, 0x64, 0x7d, 0x12, 0x6b, 0x0a, 0x0c, 0x4c,
	0x69, 0x73, 0x74, 0x53, 0x6c, 0x69, 0x64, 0x65, 0x73, 0x56, 0x31, 0x12, 0x22, 0x2e, 0x6f, 0x63,
	0x70, 0x2e, 0x73, 0x6c, 0x69, 0x64, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x53, 0x6c, 0x69, 0x64, 0x65, 0x73, 0x56, 0x31, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x23, 0x2e, 0x6f, 0x63, 0x70, 0x2e, 0x73, 0x6c, 0x69, 0x64, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x53, 0x6c, 0x69, 0x64, 0x65, 0x73, 0x56, 0x31, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x12, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0c, 0x12, 0x0a, 0x2f, 0x76,
	0x31, 0x2f, 0x73, 0x6c, 0x69, 0x64, 0x65, 0x73, 0x12, 0x79, 0x0a, 0x0d, 0x52, 0x65, 0x6d, 0x6f,
	0x76, 0x65, 0x53, 0x6c, 0x69, 0x64, 0x65, 0x56, 0x31, 0x12, 0x23, 0x2e, 0x6f, 0x63, 0x70, 0x2e,
	0x73, 0x6c, 0x69, 0x64, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65,
	0x53, 0x6c, 0x69, 0x64, 0x65, 0x56, 0x31, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24,
	0x2e, 0x6f, 0x63, 0x70, 0x2e, 0x73, 0x6c, 0x69, 0x64, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x52,
	0x65, 0x6d, 0x6f, 0x76, 0x65, 0x53, 0x6c, 0x69, 0x64, 0x65, 0x56, 0x31, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17, 0x2a, 0x15, 0x2f, 0x76,
	0x31, 0x2f, 0x73, 0x6c, 0x69, 0x64, 0x65, 0x73, 0x2f, 0x7b, 0x73, 0x6c, 0x69, 0x64, 0x65, 0x5f,
	0x69, 0x64, 0x7d, 0x42, 0x3a, 0x5a, 0x38, 0x67, 0x69, 0x68, 0x74, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x6f, 0x7a, 0x6f, 0x6e, 0x63, 0x70, 0x2f, 0x6f, 0x63, 0x70, 0x2d, 0x70, 0x72, 0x65,
	0x73, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x6b,
	0x67, 0x2f, 0x6f, 0x63, 0x70, 0x2d, 0x73, 0x6c, 0x69, 0x64, 0x65, 0x2d, 0x61, 0x70, 0x69, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_ocp_slide_api_ocp_slide_api_proto_rawDescOnce sync.Once
	file_api_ocp_slide_api_ocp_slide_api_proto_rawDescData = file_api_ocp_slide_api_ocp_slide_api_proto_rawDesc
)

func file_api_ocp_slide_api_ocp_slide_api_proto_rawDescGZIP() []byte {
	file_api_ocp_slide_api_ocp_slide_api_proto_rawDescOnce.Do(func() {
		file_api_ocp_slide_api_ocp_slide_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_ocp_slide_api_ocp_slide_api_proto_rawDescData)
	})
	return file_api_ocp_slide_api_ocp_slide_api_proto_rawDescData
}

var file_api_ocp_slide_api_ocp_slide_api_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_ocp_slide_api_ocp_slide_api_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_api_ocp_slide_api_ocp_slide_api_proto_goTypes = []interface{}{
	(SlideType)(0),                  // 0: ocp.slide.api.SlideType
	(*CreateSlideV1Request)(nil),    // 1: ocp.slide.api.CreateSlideV1Request
	(*CreateSlideV1Response)(nil),   // 2: ocp.slide.api.CreateSlideV1Response
	(*DescribeSlideV1Request)(nil),  // 3: ocp.slide.api.DescribeSlideV1Request
	(*DescribeSlideV1Response)(nil), // 4: ocp.slide.api.DescribeSlideV1Response
	(*ListSlidesV1Request)(nil),     // 5: ocp.slide.api.ListSlidesV1Request
	(*ListSlidesV1Response)(nil),    // 6: ocp.slide.api.ListSlidesV1Response
	(*RemoveSlideV1Request)(nil),    // 7: ocp.slide.api.RemoveSlideV1Request
	(*RemoveSlideV1Response)(nil),   // 8: ocp.slide.api.RemoveSlideV1Response
	(*Slide)(nil),                   // 9: ocp.slide.api.Slide
}
var file_api_ocp_slide_api_ocp_slide_api_proto_depIdxs = []int32{
	0, // 0: ocp.slide.api.CreateSlideV1Request.type:type_name -> ocp.slide.api.SlideType
	9, // 1: ocp.slide.api.DescribeSlideV1Response.slide:type_name -> ocp.slide.api.Slide
	9, // 2: ocp.slide.api.ListSlidesV1Response.slides:type_name -> ocp.slide.api.Slide
	0, // 3: ocp.slide.api.Slide.type:type_name -> ocp.slide.api.SlideType
	1, // 4: ocp.slide.api.SlideAPI.CreateSlideV1:input_type -> ocp.slide.api.CreateSlideV1Request
	3, // 5: ocp.slide.api.SlideAPI.DescribeSlideV1:input_type -> ocp.slide.api.DescribeSlideV1Request
	5, // 6: ocp.slide.api.SlideAPI.ListSlidesV1:input_type -> ocp.slide.api.ListSlidesV1Request
	7, // 7: ocp.slide.api.SlideAPI.RemoveSlideV1:input_type -> ocp.slide.api.RemoveSlideV1Request
	2, // 8: ocp.slide.api.SlideAPI.CreateSlideV1:output_type -> ocp.slide.api.CreateSlideV1Response
	4, // 9: ocp.slide.api.SlideAPI.DescribeSlideV1:output_type -> ocp.slide.api.DescribeSlideV1Response
	6, // 10: ocp.slide.api.SlideAPI.ListSlidesV1:output_type -> ocp.slide.api.ListSlidesV1Response
	8, // 11: ocp.slide.api.SlideAPI.RemoveSlideV1:output_type -> ocp.slide.api.RemoveSlideV1Response
	8, // [8:12] is the sub-list for method output_type
	4, // [4:8] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_api_ocp_slide_api_ocp_slide_api_proto_init() }
func file_api_ocp_slide_api_ocp_slide_api_proto_init() {
	if File_api_ocp_slide_api_ocp_slide_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_ocp_slide_api_ocp_slide_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateSlideV1Request); i {
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
		file_api_ocp_slide_api_ocp_slide_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateSlideV1Response); i {
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
		file_api_ocp_slide_api_ocp_slide_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DescribeSlideV1Request); i {
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
		file_api_ocp_slide_api_ocp_slide_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DescribeSlideV1Response); i {
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
		file_api_ocp_slide_api_ocp_slide_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListSlidesV1Request); i {
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
		file_api_ocp_slide_api_ocp_slide_api_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListSlidesV1Response); i {
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
		file_api_ocp_slide_api_ocp_slide_api_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveSlideV1Request); i {
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
		file_api_ocp_slide_api_ocp_slide_api_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveSlideV1Response); i {
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
		file_api_ocp_slide_api_ocp_slide_api_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Slide); i {
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
			RawDescriptor: file_api_ocp_slide_api_ocp_slide_api_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_ocp_slide_api_ocp_slide_api_proto_goTypes,
		DependencyIndexes: file_api_ocp_slide_api_ocp_slide_api_proto_depIdxs,
		EnumInfos:         file_api_ocp_slide_api_ocp_slide_api_proto_enumTypes,
		MessageInfos:      file_api_ocp_slide_api_ocp_slide_api_proto_msgTypes,
	}.Build()
	File_api_ocp_slide_api_ocp_slide_api_proto = out.File
	file_api_ocp_slide_api_ocp_slide_api_proto_rawDesc = nil
	file_api_ocp_slide_api_ocp_slide_api_proto_goTypes = nil
	file_api_ocp_slide_api_ocp_slide_api_proto_depIdxs = nil
}
