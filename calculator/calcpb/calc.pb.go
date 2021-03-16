// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.15.6
// source: calculator/calcpb/calc.proto

package calcpb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Addition struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Params []int32 `protobuf:"varint,1,rep,packed,name=params,proto3" json:"params,omitempty"`
}

func (x *Addition) Reset() {
	*x = Addition{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calculator_calcpb_calc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Addition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Addition) ProtoMessage() {}

func (x *Addition) ProtoReflect() protoreflect.Message {
	mi := &file_calculator_calcpb_calc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Addition.ProtoReflect.Descriptor instead.
func (*Addition) Descriptor() ([]byte, []int) {
	return file_calculator_calcpb_calc_proto_rawDescGZIP(), []int{0}
}

func (x *Addition) GetParams() []int32 {
	if x != nil {
		return x.Params
	}
	return nil
}

type AddRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Addition *Addition `protobuf:"bytes,1,opt,name=addition,proto3" json:"addition,omitempty"`
}

func (x *AddRequest) Reset() {
	*x = AddRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calculator_calcpb_calc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddRequest) ProtoMessage() {}

func (x *AddRequest) ProtoReflect() protoreflect.Message {
	mi := &file_calculator_calcpb_calc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddRequest.ProtoReflect.Descriptor instead.
func (*AddRequest) Descriptor() ([]byte, []int) {
	return file_calculator_calcpb_calc_proto_rawDescGZIP(), []int{1}
}

func (x *AddRequest) GetAddition() *Addition {
	if x != nil {
		return x.Addition
	}
	return nil
}

type AddResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result int32 `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *AddResponse) Reset() {
	*x = AddResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calculator_calcpb_calc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddResponse) ProtoMessage() {}

func (x *AddResponse) ProtoReflect() protoreflect.Message {
	mi := &file_calculator_calcpb_calc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddResponse.ProtoReflect.Descriptor instead.
func (*AddResponse) Descriptor() ([]byte, []int) {
	return file_calculator_calcpb_calc_proto_rawDescGZIP(), []int{2}
}

func (x *AddResponse) GetResult() int32 {
	if x != nil {
		return x.Result
	}
	return 0
}

type PNDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Number int32 `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
}

func (x *PNDRequest) Reset() {
	*x = PNDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calculator_calcpb_calc_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PNDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PNDRequest) ProtoMessage() {}

func (x *PNDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_calculator_calcpb_calc_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PNDRequest.ProtoReflect.Descriptor instead.
func (*PNDRequest) Descriptor() ([]byte, []int) {
	return file_calculator_calcpb_calc_proto_rawDescGZIP(), []int{3}
}

func (x *PNDRequest) GetNumber() int32 {
	if x != nil {
		return x.Number
	}
	return 0
}

type PNDResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Number int32 `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
}

func (x *PNDResponse) Reset() {
	*x = PNDResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calculator_calcpb_calc_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PNDResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PNDResponse) ProtoMessage() {}

func (x *PNDResponse) ProtoReflect() protoreflect.Message {
	mi := &file_calculator_calcpb_calc_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PNDResponse.ProtoReflect.Descriptor instead.
func (*PNDResponse) Descriptor() ([]byte, []int) {
	return file_calculator_calcpb_calc_proto_rawDescGZIP(), []int{4}
}

func (x *PNDResponse) GetNumber() int32 {
	if x != nil {
		return x.Number
	}
	return 0
}

type ComputeAvgResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result float64 `protobuf:"fixed64,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *ComputeAvgResponse) Reset() {
	*x = ComputeAvgResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calculator_calcpb_calc_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ComputeAvgResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ComputeAvgResponse) ProtoMessage() {}

func (x *ComputeAvgResponse) ProtoReflect() protoreflect.Message {
	mi := &file_calculator_calcpb_calc_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ComputeAvgResponse.ProtoReflect.Descriptor instead.
func (*ComputeAvgResponse) Descriptor() ([]byte, []int) {
	return file_calculator_calcpb_calc_proto_rawDescGZIP(), []int{5}
}

func (x *ComputeAvgResponse) GetResult() float64 {
	if x != nil {
		return x.Result
	}
	return 0
}

type ComputeAvgRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Number int32 `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
}

func (x *ComputeAvgRequest) Reset() {
	*x = ComputeAvgRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calculator_calcpb_calc_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ComputeAvgRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ComputeAvgRequest) ProtoMessage() {}

func (x *ComputeAvgRequest) ProtoReflect() protoreflect.Message {
	mi := &file_calculator_calcpb_calc_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ComputeAvgRequest.ProtoReflect.Descriptor instead.
func (*ComputeAvgRequest) Descriptor() ([]byte, []int) {
	return file_calculator_calcpb_calc_proto_rawDescGZIP(), []int{6}
}

func (x *ComputeAvgRequest) GetNumber() int32 {
	if x != nil {
		return x.Number
	}
	return 0
}

type FindMaxRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Number int32 `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
}

func (x *FindMaxRequest) Reset() {
	*x = FindMaxRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calculator_calcpb_calc_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindMaxRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindMaxRequest) ProtoMessage() {}

func (x *FindMaxRequest) ProtoReflect() protoreflect.Message {
	mi := &file_calculator_calcpb_calc_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindMaxRequest.ProtoReflect.Descriptor instead.
func (*FindMaxRequest) Descriptor() ([]byte, []int) {
	return file_calculator_calcpb_calc_proto_rawDescGZIP(), []int{7}
}

func (x *FindMaxRequest) GetNumber() int32 {
	if x != nil {
		return x.Number
	}
	return 0
}

type FindMaxResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MaxNumber int32 `protobuf:"varint,1,opt,name=maxNumber,proto3" json:"maxNumber,omitempty"`
}

func (x *FindMaxResponse) Reset() {
	*x = FindMaxResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calculator_calcpb_calc_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindMaxResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindMaxResponse) ProtoMessage() {}

func (x *FindMaxResponse) ProtoReflect() protoreflect.Message {
	mi := &file_calculator_calcpb_calc_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindMaxResponse.ProtoReflect.Descriptor instead.
func (*FindMaxResponse) Descriptor() ([]byte, []int) {
	return file_calculator_calcpb_calc_proto_rawDescGZIP(), []int{8}
}

func (x *FindMaxResponse) GetMaxNumber() int32 {
	if x != nil {
		return x.MaxNumber
	}
	return 0
}

type SqrtRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Number float64 `protobuf:"fixed64,1,opt,name=number,proto3" json:"number,omitempty"`
}

func (x *SqrtRequest) Reset() {
	*x = SqrtRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calculator_calcpb_calc_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SqrtRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SqrtRequest) ProtoMessage() {}

func (x *SqrtRequest) ProtoReflect() protoreflect.Message {
	mi := &file_calculator_calcpb_calc_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SqrtRequest.ProtoReflect.Descriptor instead.
func (*SqrtRequest) Descriptor() ([]byte, []int) {
	return file_calculator_calcpb_calc_proto_rawDescGZIP(), []int{9}
}

func (x *SqrtRequest) GetNumber() float64 {
	if x != nil {
		return x.Number
	}
	return 0
}

type SqrtResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result float64 `protobuf:"fixed64,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *SqrtResponse) Reset() {
	*x = SqrtResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calculator_calcpb_calc_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SqrtResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SqrtResponse) ProtoMessage() {}

func (x *SqrtResponse) ProtoReflect() protoreflect.Message {
	mi := &file_calculator_calcpb_calc_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SqrtResponse.ProtoReflect.Descriptor instead.
func (*SqrtResponse) Descriptor() ([]byte, []int) {
	return file_calculator_calcpb_calc_proto_rawDescGZIP(), []int{10}
}

func (x *SqrtResponse) GetResult() float64 {
	if x != nil {
		return x.Result
	}
	return 0
}

var File_calculator_calcpb_calc_proto protoreflect.FileDescriptor

var file_calculator_calcpb_calc_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2f, 0x63, 0x61, 0x6c,
	0x63, 0x70, 0x62, 0x2f, 0x63, 0x61, 0x6c, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a,
	0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x22, 0x22, 0x0a, 0x08, 0x41, 0x64,
	0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x05, 0x52, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x22, 0x3e,
	0x0a, 0x0a, 0x41, 0x64, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x30, 0x0a, 0x08,
	0x61, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14,
	0x2e, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x41, 0x64, 0x64, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x61, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x25,
	0x0a, 0x0b, 0x41, 0x64, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x72,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x24, 0x0a, 0x0a, 0x50, 0x4e, 0x44, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x25, 0x0a, 0x0b, 0x50,
	0x4e, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x22, 0x2c, 0x0a, 0x12, 0x43, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x41, 0x76, 0x67,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x22, 0x2b, 0x0a, 0x11, 0x43, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x41, 0x76, 0x67, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x28, 0x0a,
	0x0e, 0x46, 0x69, 0x6e, 0x64, 0x4d, 0x61, 0x78, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x2f, 0x0a, 0x0f, 0x46, 0x69, 0x6e, 0x64, 0x4d,
	0x61, 0x78, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6d, 0x61,
	0x78, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x6d,
	0x61, 0x78, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x25, 0x0a, 0x0b, 0x53, 0x71, 0x72, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22,
	0x26, 0x0a, 0x0c, 0x53, 0x71, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x32, 0xe7, 0x02, 0x0a, 0x0a, 0x43, 0x61, 0x6c, 0x63,
	0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x36, 0x0a, 0x03, 0x41, 0x64, 0x64, 0x12, 0x16, 0x2e,
	0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x41, 0x64, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74,
	0x6f, 0x72, 0x2e, 0x41, 0x64, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4f,
	0x0a, 0x1a, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x44,
	0x65, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x2e, 0x63,
	0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x50, 0x4e, 0x44, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f,
	0x72, 0x2e, 0x50, 0x4e, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01, 0x12,
	0x4d, 0x0a, 0x0a, 0x43, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x41, 0x76, 0x67, 0x12, 0x1d, 0x2e,
	0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x75,
	0x74, 0x65, 0x41, 0x76, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x63,
	0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x75, 0x74,
	0x65, 0x41, 0x76, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x28, 0x01, 0x12, 0x46,
	0x0a, 0x07, 0x46, 0x69, 0x6e, 0x64, 0x4d, 0x61, 0x78, 0x12, 0x1a, 0x2e, 0x63, 0x61, 0x6c, 0x63,
	0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x4d, 0x61, 0x78, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74,
	0x6f, 0x72, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x4d, 0x61, 0x78, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x28, 0x01, 0x30, 0x01, 0x12, 0x39, 0x0a, 0x04, 0x53, 0x71, 0x72, 0x74, 0x12, 0x17,
	0x2e, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x53, 0x71, 0x72, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c,
	0x61, 0x74, 0x6f, 0x72, 0x2e, 0x53, 0x71, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x42, 0x15, 0x5a, 0x13, 0x2e, 0x2f, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f,
	0x72, 0x2f, 0x63, 0x61, 0x6c, 0x63, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_calculator_calcpb_calc_proto_rawDescOnce sync.Once
	file_calculator_calcpb_calc_proto_rawDescData = file_calculator_calcpb_calc_proto_rawDesc
)

func file_calculator_calcpb_calc_proto_rawDescGZIP() []byte {
	file_calculator_calcpb_calc_proto_rawDescOnce.Do(func() {
		file_calculator_calcpb_calc_proto_rawDescData = protoimpl.X.CompressGZIP(file_calculator_calcpb_calc_proto_rawDescData)
	})
	return file_calculator_calcpb_calc_proto_rawDescData
}

var file_calculator_calcpb_calc_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_calculator_calcpb_calc_proto_goTypes = []interface{}{
	(*Addition)(nil),           // 0: calculator.Addition
	(*AddRequest)(nil),         // 1: calculator.AddRequest
	(*AddResponse)(nil),        // 2: calculator.AddResponse
	(*PNDRequest)(nil),         // 3: calculator.PNDRequest
	(*PNDResponse)(nil),        // 4: calculator.PNDResponse
	(*ComputeAvgResponse)(nil), // 5: calculator.ComputeAvgResponse
	(*ComputeAvgRequest)(nil),  // 6: calculator.ComputeAvgRequest
	(*FindMaxRequest)(nil),     // 7: calculator.FindMaxRequest
	(*FindMaxResponse)(nil),    // 8: calculator.FindMaxResponse
	(*SqrtRequest)(nil),        // 9: calculator.SqrtRequest
	(*SqrtResponse)(nil),       // 10: calculator.SqrtResponse
}
var file_calculator_calcpb_calc_proto_depIdxs = []int32{
	0,  // 0: calculator.AddRequest.addition:type_name -> calculator.Addition
	1,  // 1: calculator.Calculator.Add:input_type -> calculator.AddRequest
	3,  // 2: calculator.Calculator.PrimaryNumberDecomposition:input_type -> calculator.PNDRequest
	6,  // 3: calculator.Calculator.ComputeAvg:input_type -> calculator.ComputeAvgRequest
	7,  // 4: calculator.Calculator.FindMax:input_type -> calculator.FindMaxRequest
	9,  // 5: calculator.Calculator.Sqrt:input_type -> calculator.SqrtRequest
	2,  // 6: calculator.Calculator.Add:output_type -> calculator.AddResponse
	4,  // 7: calculator.Calculator.PrimaryNumberDecomposition:output_type -> calculator.PNDResponse
	5,  // 8: calculator.Calculator.ComputeAvg:output_type -> calculator.ComputeAvgResponse
	8,  // 9: calculator.Calculator.FindMax:output_type -> calculator.FindMaxResponse
	10, // 10: calculator.Calculator.Sqrt:output_type -> calculator.SqrtResponse
	6,  // [6:11] is the sub-list for method output_type
	1,  // [1:6] is the sub-list for method input_type
	1,  // [1:1] is the sub-list for extension type_name
	1,  // [1:1] is the sub-list for extension extendee
	0,  // [0:1] is the sub-list for field type_name
}

func init() { file_calculator_calcpb_calc_proto_init() }
func file_calculator_calcpb_calc_proto_init() {
	if File_calculator_calcpb_calc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_calculator_calcpb_calc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Addition); i {
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
		file_calculator_calcpb_calc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddRequest); i {
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
		file_calculator_calcpb_calc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddResponse); i {
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
		file_calculator_calcpb_calc_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PNDRequest); i {
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
		file_calculator_calcpb_calc_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PNDResponse); i {
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
		file_calculator_calcpb_calc_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ComputeAvgResponse); i {
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
		file_calculator_calcpb_calc_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ComputeAvgRequest); i {
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
		file_calculator_calcpb_calc_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindMaxRequest); i {
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
		file_calculator_calcpb_calc_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindMaxResponse); i {
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
		file_calculator_calcpb_calc_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SqrtRequest); i {
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
		file_calculator_calcpb_calc_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SqrtResponse); i {
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
			RawDescriptor: file_calculator_calcpb_calc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_calculator_calcpb_calc_proto_goTypes,
		DependencyIndexes: file_calculator_calcpb_calc_proto_depIdxs,
		MessageInfos:      file_calculator_calcpb_calc_proto_msgTypes,
	}.Build()
	File_calculator_calcpb_calc_proto = out.File
	file_calculator_calcpb_calc_proto_rawDesc = nil
	file_calculator_calcpb_calc_proto_goTypes = nil
	file_calculator_calcpb_calc_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CalculatorClient is the client API for Calculator service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CalculatorClient interface {
	// Uniray RPC
	Add(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*AddResponse, error)
	// Streaming Server RPC
	PrimaryNumberDecomposition(ctx context.Context, in *PNDRequest, opts ...grpc.CallOption) (Calculator_PrimaryNumberDecompositionClient, error)
	// Streaming Client RPC
	ComputeAvg(ctx context.Context, opts ...grpc.CallOption) (Calculator_ComputeAvgClient, error)
	// Bi-Direcctional Streaming RPC
	FindMax(ctx context.Context, opts ...grpc.CallOption) (Calculator_FindMaxClient, error)
	// Unary RPC
	Sqrt(ctx context.Context, in *SqrtRequest, opts ...grpc.CallOption) (*SqrtResponse, error)
}

type calculatorClient struct {
	cc grpc.ClientConnInterface
}

func NewCalculatorClient(cc grpc.ClientConnInterface) CalculatorClient {
	return &calculatorClient{cc}
}

func (c *calculatorClient) Add(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*AddResponse, error) {
	out := new(AddResponse)
	err := c.cc.Invoke(ctx, "/calculator.Calculator/Add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calculatorClient) PrimaryNumberDecomposition(ctx context.Context, in *PNDRequest, opts ...grpc.CallOption) (Calculator_PrimaryNumberDecompositionClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Calculator_serviceDesc.Streams[0], "/calculator.Calculator/PrimaryNumberDecomposition", opts...)
	if err != nil {
		return nil, err
	}
	x := &calculatorPrimaryNumberDecompositionClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Calculator_PrimaryNumberDecompositionClient interface {
	Recv() (*PNDResponse, error)
	grpc.ClientStream
}

type calculatorPrimaryNumberDecompositionClient struct {
	grpc.ClientStream
}

func (x *calculatorPrimaryNumberDecompositionClient) Recv() (*PNDResponse, error) {
	m := new(PNDResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *calculatorClient) ComputeAvg(ctx context.Context, opts ...grpc.CallOption) (Calculator_ComputeAvgClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Calculator_serviceDesc.Streams[1], "/calculator.Calculator/ComputeAvg", opts...)
	if err != nil {
		return nil, err
	}
	x := &calculatorComputeAvgClient{stream}
	return x, nil
}

type Calculator_ComputeAvgClient interface {
	Send(*ComputeAvgRequest) error
	CloseAndRecv() (*ComputeAvgResponse, error)
	grpc.ClientStream
}

type calculatorComputeAvgClient struct {
	grpc.ClientStream
}

func (x *calculatorComputeAvgClient) Send(m *ComputeAvgRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *calculatorComputeAvgClient) CloseAndRecv() (*ComputeAvgResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(ComputeAvgResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *calculatorClient) FindMax(ctx context.Context, opts ...grpc.CallOption) (Calculator_FindMaxClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Calculator_serviceDesc.Streams[2], "/calculator.Calculator/FindMax", opts...)
	if err != nil {
		return nil, err
	}
	x := &calculatorFindMaxClient{stream}
	return x, nil
}

type Calculator_FindMaxClient interface {
	Send(*FindMaxRequest) error
	Recv() (*FindMaxResponse, error)
	grpc.ClientStream
}

type calculatorFindMaxClient struct {
	grpc.ClientStream
}

func (x *calculatorFindMaxClient) Send(m *FindMaxRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *calculatorFindMaxClient) Recv() (*FindMaxResponse, error) {
	m := new(FindMaxResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *calculatorClient) Sqrt(ctx context.Context, in *SqrtRequest, opts ...grpc.CallOption) (*SqrtResponse, error) {
	out := new(SqrtResponse)
	err := c.cc.Invoke(ctx, "/calculator.Calculator/Sqrt", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CalculatorServer is the server API for Calculator service.
type CalculatorServer interface {
	// Uniray RPC
	Add(context.Context, *AddRequest) (*AddResponse, error)
	// Streaming Server RPC
	PrimaryNumberDecomposition(*PNDRequest, Calculator_PrimaryNumberDecompositionServer) error
	// Streaming Client RPC
	ComputeAvg(Calculator_ComputeAvgServer) error
	// Bi-Direcctional Streaming RPC
	FindMax(Calculator_FindMaxServer) error
	// Unary RPC
	Sqrt(context.Context, *SqrtRequest) (*SqrtResponse, error)
}

// UnimplementedCalculatorServer can be embedded to have forward compatible implementations.
type UnimplementedCalculatorServer struct {
}

func (*UnimplementedCalculatorServer) Add(context.Context, *AddRequest) (*AddResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}
func (*UnimplementedCalculatorServer) PrimaryNumberDecomposition(*PNDRequest, Calculator_PrimaryNumberDecompositionServer) error {
	return status.Errorf(codes.Unimplemented, "method PrimaryNumberDecomposition not implemented")
}
func (*UnimplementedCalculatorServer) ComputeAvg(Calculator_ComputeAvgServer) error {
	return status.Errorf(codes.Unimplemented, "method ComputeAvg not implemented")
}
func (*UnimplementedCalculatorServer) FindMax(Calculator_FindMaxServer) error {
	return status.Errorf(codes.Unimplemented, "method FindMax not implemented")
}
func (*UnimplementedCalculatorServer) Sqrt(context.Context, *SqrtRequest) (*SqrtResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Sqrt not implemented")
}

func RegisterCalculatorServer(s *grpc.Server, srv CalculatorServer) {
	s.RegisterService(&_Calculator_serviceDesc, srv)
}

func _Calculator_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculatorServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calculator.Calculator/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculatorServer).Add(ctx, req.(*AddRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Calculator_PrimaryNumberDecomposition_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(PNDRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CalculatorServer).PrimaryNumberDecomposition(m, &calculatorPrimaryNumberDecompositionServer{stream})
}

type Calculator_PrimaryNumberDecompositionServer interface {
	Send(*PNDResponse) error
	grpc.ServerStream
}

type calculatorPrimaryNumberDecompositionServer struct {
	grpc.ServerStream
}

func (x *calculatorPrimaryNumberDecompositionServer) Send(m *PNDResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _Calculator_ComputeAvg_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CalculatorServer).ComputeAvg(&calculatorComputeAvgServer{stream})
}

type Calculator_ComputeAvgServer interface {
	SendAndClose(*ComputeAvgResponse) error
	Recv() (*ComputeAvgRequest, error)
	grpc.ServerStream
}

type calculatorComputeAvgServer struct {
	grpc.ServerStream
}

func (x *calculatorComputeAvgServer) SendAndClose(m *ComputeAvgResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *calculatorComputeAvgServer) Recv() (*ComputeAvgRequest, error) {
	m := new(ComputeAvgRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Calculator_FindMax_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CalculatorServer).FindMax(&calculatorFindMaxServer{stream})
}

type Calculator_FindMaxServer interface {
	Send(*FindMaxResponse) error
	Recv() (*FindMaxRequest, error)
	grpc.ServerStream
}

type calculatorFindMaxServer struct {
	grpc.ServerStream
}

func (x *calculatorFindMaxServer) Send(m *FindMaxResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *calculatorFindMaxServer) Recv() (*FindMaxRequest, error) {
	m := new(FindMaxRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Calculator_Sqrt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SqrtRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculatorServer).Sqrt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calculator.Calculator/Sqrt",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculatorServer).Sqrt(ctx, req.(*SqrtRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Calculator_serviceDesc = grpc.ServiceDesc{
	ServiceName: "calculator.Calculator",
	HandlerType: (*CalculatorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Add",
			Handler:    _Calculator_Add_Handler,
		},
		{
			MethodName: "Sqrt",
			Handler:    _Calculator_Sqrt_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "PrimaryNumberDecomposition",
			Handler:       _Calculator_PrimaryNumberDecomposition_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ComputeAvg",
			Handler:       _Calculator_ComputeAvg_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "FindMax",
			Handler:       _Calculator_FindMax_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "calculator/calcpb/calc.proto",
}
