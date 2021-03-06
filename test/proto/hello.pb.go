// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: test/proto/hello.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type HelloRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      string                  `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	StrVal    *wrapperspb.StringValue `protobuf:"bytes,2,opt,name=strVal,proto3" json:"strVal,omitempty"`
	FloatVal  *wrapperspb.FloatValue  `protobuf:"bytes,3,opt,name=floatVal,proto3" json:"floatVal,omitempty"`
	DoubleVal *wrapperspb.DoubleValue `protobuf:"bytes,4,opt,name=doubleVal,proto3" json:"doubleVal,omitempty"`
	BoolVal   *wrapperspb.BoolValue   `protobuf:"bytes,5,opt,name=boolVal,proto3" json:"boolVal,omitempty"`
	BytesVal  *wrapperspb.BytesValue  `protobuf:"bytes,6,opt,name=bytesVal,proto3" json:"bytesVal,omitempty"`
	Int32Val  *wrapperspb.Int32Value  `protobuf:"bytes,7,opt,name=int32Val,proto3" json:"int32Val,omitempty"`
	Uint32Val *wrapperspb.UInt32Value `protobuf:"bytes,8,opt,name=uint32Val,proto3" json:"uint32Val,omitempty"`
	Int64Val  *wrapperspb.Int64Value  `protobuf:"bytes,9,opt,name=int64Val,proto3" json:"int64Val,omitempty"`
	Uint64Val *wrapperspb.UInt64Value `protobuf:"bytes,10,opt,name=uint64Val,proto3" json:"uint64Val,omitempty"`
}

func (x *HelloRequest) Reset() {
	*x = HelloRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_proto_hello_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloRequest) ProtoMessage() {}

func (x *HelloRequest) ProtoReflect() protoreflect.Message {
	mi := &file_test_proto_hello_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloRequest.ProtoReflect.Descriptor instead.
func (*HelloRequest) Descriptor() ([]byte, []int) {
	return file_test_proto_hello_proto_rawDescGZIP(), []int{0}
}

func (x *HelloRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *HelloRequest) GetStrVal() *wrapperspb.StringValue {
	if x != nil {
		return x.StrVal
	}
	return nil
}

func (x *HelloRequest) GetFloatVal() *wrapperspb.FloatValue {
	if x != nil {
		return x.FloatVal
	}
	return nil
}

func (x *HelloRequest) GetDoubleVal() *wrapperspb.DoubleValue {
	if x != nil {
		return x.DoubleVal
	}
	return nil
}

func (x *HelloRequest) GetBoolVal() *wrapperspb.BoolValue {
	if x != nil {
		return x.BoolVal
	}
	return nil
}

func (x *HelloRequest) GetBytesVal() *wrapperspb.BytesValue {
	if x != nil {
		return x.BytesVal
	}
	return nil
}

func (x *HelloRequest) GetInt32Val() *wrapperspb.Int32Value {
	if x != nil {
		return x.Int32Val
	}
	return nil
}

func (x *HelloRequest) GetUint32Val() *wrapperspb.UInt32Value {
	if x != nil {
		return x.Uint32Val
	}
	return nil
}

func (x *HelloRequest) GetInt64Val() *wrapperspb.Int64Value {
	if x != nil {
		return x.Int64Val
	}
	return nil
}

func (x *HelloRequest) GetUint64Val() *wrapperspb.UInt64Value {
	if x != nil {
		return x.Uint64Val
	}
	return nil
}

type HelloResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *HelloResponse) Reset() {
	*x = HelloResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_proto_hello_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloResponse) ProtoMessage() {}

func (x *HelloResponse) ProtoReflect() protoreflect.Message {
	mi := &file_test_proto_hello_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloResponse.ProtoReflect.Descriptor instead.
func (*HelloResponse) Descriptor() ([]byte, []int) {
	return file_test_proto_hello_proto_rawDescGZIP(), []int{1}
}

func (x *HelloResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type SendMyGiftRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GiftId   int32  `protobuf:"varint,1,opt,name=gift_id,json=giftId,proto3" json:"gift_id,omitempty"`
	GiftName string `protobuf:"bytes,2,opt,name=gift_name,json=giftName,proto3" json:"gift_name,omitempty"`
}

func (x *SendMyGiftRequest) Reset() {
	*x = SendMyGiftRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_proto_hello_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendMyGiftRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendMyGiftRequest) ProtoMessage() {}

func (x *SendMyGiftRequest) ProtoReflect() protoreflect.Message {
	mi := &file_test_proto_hello_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendMyGiftRequest.ProtoReflect.Descriptor instead.
func (*SendMyGiftRequest) Descriptor() ([]byte, []int) {
	return file_test_proto_hello_proto_rawDescGZIP(), []int{2}
}

func (x *SendMyGiftRequest) GetGiftId() int32 {
	if x != nil {
		return x.GiftId
	}
	return 0
}

func (x *SendMyGiftRequest) GetGiftName() string {
	if x != nil {
		return x.GiftName
	}
	return ""
}

type SendMyGiftResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SendMyGiftResponse) Reset() {
	*x = SendMyGiftResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_proto_hello_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendMyGiftResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendMyGiftResponse) ProtoMessage() {}

func (x *SendMyGiftResponse) ProtoReflect() protoreflect.Message {
	mi := &file_test_proto_hello_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendMyGiftResponse.ProtoReflect.Descriptor instead.
func (*SendMyGiftResponse) Descriptor() ([]byte, []int) {
	return file_test_proto_hello_proto_rawDescGZIP(), []int{3}
}

var File_test_proto_hello_proto protoreflect.FileDescriptor

var file_test_proto_hello_proto_rawDesc = []byte{
	0x0a, 0x16, 0x74, 0x65, 0x73, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x68, 0x65, 0x6c,
	0x6c, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa6, 0x04, 0x0a,
	0x0c, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x34, 0x0a, 0x06, 0x73, 0x74, 0x72, 0x56, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52,
	0x06, 0x73, 0x74, 0x72, 0x56, 0x61, 0x6c, 0x12, 0x37, 0x0a, 0x08, 0x66, 0x6c, 0x6f, 0x61, 0x74,
	0x56, 0x61, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x6c, 0x6f, 0x61,
	0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x08, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x56, 0x61, 0x6c,
	0x12, 0x3a, 0x0a, 0x09, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x56, 0x61, 0x6c, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x52, 0x09, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x56, 0x61, 0x6c, 0x12, 0x34, 0x0a, 0x07,
	0x62, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x07, 0x62, 0x6f, 0x6f, 0x6c, 0x56,
	0x61, 0x6c, 0x12, 0x37, 0x0a, 0x08, 0x62, 0x79, 0x74, 0x65, 0x73, 0x56, 0x61, 0x6c, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x79, 0x74, 0x65, 0x73, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x52, 0x08, 0x62, 0x79, 0x74, 0x65, 0x73, 0x56, 0x61, 0x6c, 0x12, 0x37, 0x0a, 0x08, 0x69,
	0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x08, 0x69, 0x6e, 0x74, 0x33,
	0x32, 0x56, 0x61, 0x6c, 0x12, 0x3a, 0x0a, 0x09, 0x75, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61,
	0x6c, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e, 0x74, 0x33, 0x32,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x09, 0x75, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c,
	0x12, 0x37, 0x0a, 0x08, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52,
	0x08, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x12, 0x3a, 0x0a, 0x09, 0x75, 0x69, 0x6e,
	0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55,
	0x49, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x09, 0x75, 0x69, 0x6e, 0x74,
	0x36, 0x34, 0x56, 0x61, 0x6c, 0x22, 0x29, 0x0a, 0x0d, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x22, 0x49, 0x0a, 0x11, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x79, 0x47, 0x69, 0x66, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x67, 0x69, 0x66, 0x74, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x67, 0x69, 0x66, 0x74, 0x49, 0x64, 0x12, 0x1b,
	0x0a, 0x09, 0x67, 0x69, 0x66, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x67, 0x69, 0x66, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x14, 0x0a, 0x12, 0x53,
	0x65, 0x6e, 0x64, 0x4d, 0x79, 0x47, 0x69, 0x66, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x32, 0xb9, 0x01, 0x0a, 0x05, 0x47, 0x72, 0x65, 0x65, 0x74, 0x12, 0x34, 0x0a, 0x05, 0x48,
	0x65, 0x6c, 0x6c, 0x6f, 0x12, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x48, 0x65, 0x6c,
	0x6c, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x43, 0x0a, 0x0a, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x79, 0x47, 0x69, 0x66, 0x74, 0x12,
	0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x79, 0x47, 0x69,
	0x66, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x79, 0x47, 0x69, 0x66, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x35, 0x0a, 0x06, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x32,
	0x12, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x48, 0x65,
	0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x32, 0x5e, 0x0a,
	0x1c, 0x41, 0x6e, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x57,
	0x69, 0x74, 0x68, 0x4e, 0x6f, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x3e, 0x0a,
	0x0a, 0x4e, 0x6f, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x31, 0x5a,
	0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x79, 0x78, 0x6c, 0x69,
	0x6d, 0x6f, 0x2f, 0x67, 0x6f, 0x2d, 0x6a, 0x73, 0x6f, 0x6e, 0x72, 0x70, 0x63, 0x2d, 0x67, 0x61,
	0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_test_proto_hello_proto_rawDescOnce sync.Once
	file_test_proto_hello_proto_rawDescData = file_test_proto_hello_proto_rawDesc
)

func file_test_proto_hello_proto_rawDescGZIP() []byte {
	file_test_proto_hello_proto_rawDescOnce.Do(func() {
		file_test_proto_hello_proto_rawDescData = protoimpl.X.CompressGZIP(file_test_proto_hello_proto_rawDescData)
	})
	return file_test_proto_hello_proto_rawDescData
}

var file_test_proto_hello_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_test_proto_hello_proto_goTypes = []interface{}{
	(*HelloRequest)(nil),           // 0: proto.HelloRequest
	(*HelloResponse)(nil),          // 1: proto.HelloResponse
	(*SendMyGiftRequest)(nil),      // 2: proto.SendMyGiftRequest
	(*SendMyGiftResponse)(nil),     // 3: proto.SendMyGiftResponse
	(*wrapperspb.StringValue)(nil), // 4: google.protobuf.StringValue
	(*wrapperspb.FloatValue)(nil),  // 5: google.protobuf.FloatValue
	(*wrapperspb.DoubleValue)(nil), // 6: google.protobuf.DoubleValue
	(*wrapperspb.BoolValue)(nil),   // 7: google.protobuf.BoolValue
	(*wrapperspb.BytesValue)(nil),  // 8: google.protobuf.BytesValue
	(*wrapperspb.Int32Value)(nil),  // 9: google.protobuf.Int32Value
	(*wrapperspb.UInt32Value)(nil), // 10: google.protobuf.UInt32Value
	(*wrapperspb.Int64Value)(nil),  // 11: google.protobuf.Int64Value
	(*wrapperspb.UInt64Value)(nil), // 12: google.protobuf.UInt64Value
	(*emptypb.Empty)(nil),          // 13: google.protobuf.Empty
}
var file_test_proto_hello_proto_depIdxs = []int32{
	4,  // 0: proto.HelloRequest.strVal:type_name -> google.protobuf.StringValue
	5,  // 1: proto.HelloRequest.floatVal:type_name -> google.protobuf.FloatValue
	6,  // 2: proto.HelloRequest.doubleVal:type_name -> google.protobuf.DoubleValue
	7,  // 3: proto.HelloRequest.boolVal:type_name -> google.protobuf.BoolValue
	8,  // 4: proto.HelloRequest.bytesVal:type_name -> google.protobuf.BytesValue
	9,  // 5: proto.HelloRequest.int32Val:type_name -> google.protobuf.Int32Value
	10, // 6: proto.HelloRequest.uint32Val:type_name -> google.protobuf.UInt32Value
	11, // 7: proto.HelloRequest.int64Val:type_name -> google.protobuf.Int64Value
	12, // 8: proto.HelloRequest.uint64Val:type_name -> google.protobuf.UInt64Value
	0,  // 9: proto.Greet.Hello:input_type -> proto.HelloRequest
	2,  // 10: proto.Greet.SendMyGift:input_type -> proto.SendMyGiftRequest
	0,  // 11: proto.Greet.Hello2:input_type -> proto.HelloRequest
	13, // 12: proto.AnotherServiceWithNoBindings.NoBindings:input_type -> google.protobuf.Empty
	1,  // 13: proto.Greet.Hello:output_type -> proto.HelloResponse
	3,  // 14: proto.Greet.SendMyGift:output_type -> proto.SendMyGiftResponse
	1,  // 15: proto.Greet.Hello2:output_type -> proto.HelloResponse
	13, // 16: proto.AnotherServiceWithNoBindings.NoBindings:output_type -> google.protobuf.Empty
	13, // [13:17] is the sub-list for method output_type
	9,  // [9:13] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_test_proto_hello_proto_init() }
func file_test_proto_hello_proto_init() {
	if File_test_proto_hello_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_test_proto_hello_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloRequest); i {
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
		file_test_proto_hello_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloResponse); i {
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
		file_test_proto_hello_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendMyGiftRequest); i {
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
		file_test_proto_hello_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendMyGiftResponse); i {
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
			RawDescriptor: file_test_proto_hello_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_test_proto_hello_proto_goTypes,
		DependencyIndexes: file_test_proto_hello_proto_depIdxs,
		MessageInfos:      file_test_proto_hello_proto_msgTypes,
	}.Build()
	File_test_proto_hello_proto = out.File
	file_test_proto_hello_proto_rawDesc = nil
	file_test_proto_hello_proto_goTypes = nil
	file_test_proto_hello_proto_depIdxs = nil
}
