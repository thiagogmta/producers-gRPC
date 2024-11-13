// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v3.12.4
// source: producer.proto

package producer

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

type ResourceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Quantity int32 `protobuf:"varint,1,opt,name=quantity,proto3" json:"quantity,omitempty"`
}

func (x *ResourceRequest) Reset() {
	*x = ResourceRequest{}
	mi := &file_producer_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ResourceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResourceRequest) ProtoMessage() {}

func (x *ResourceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_producer_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResourceRequest.ProtoReflect.Descriptor instead.
func (*ResourceRequest) Descriptor() ([]byte, []int) {
	return file_producer_proto_rawDescGZIP(), []int{0}
}

func (x *ResourceRequest) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

type ResourceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResourceName string `protobuf:"bytes,1,opt,name=resourceName,proto3" json:"resourceName,omitempty"`
	Quantity     int32  `protobuf:"varint,2,opt,name=quantity,proto3" json:"quantity,omitempty"`
}

func (x *ResourceResponse) Reset() {
	*x = ResourceResponse{}
	mi := &file_producer_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ResourceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResourceResponse) ProtoMessage() {}

func (x *ResourceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_producer_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResourceResponse.ProtoReflect.Descriptor instead.
func (*ResourceResponse) Descriptor() ([]byte, []int) {
	return file_producer_proto_rawDescGZIP(), []int{1}
}

func (x *ResourceResponse) GetResourceName() string {
	if x != nil {
		return x.ResourceName
	}
	return ""
}

func (x *ResourceResponse) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

var File_producer_proto protoreflect.FileDescriptor

var file_producer_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x65, 0x72, 0x22, 0x2d, 0x0a, 0x0f, 0x52, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a,
	0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x22, 0x52, 0x0a, 0x10, 0x52, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x22, 0x0a,
	0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x32, 0x50, 0x0a,
	0x08, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x65, 0x72, 0x12, 0x44, 0x0a, 0x0b, 0x47, 0x65, 0x74,
	0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x65, 0x72, 0x2e, 0x52,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42,
	0x0b, 0x5a, 0x09, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_producer_proto_rawDescOnce sync.Once
	file_producer_proto_rawDescData = file_producer_proto_rawDesc
)

func file_producer_proto_rawDescGZIP() []byte {
	file_producer_proto_rawDescOnce.Do(func() {
		file_producer_proto_rawDescData = protoimpl.X.CompressGZIP(file_producer_proto_rawDescData)
	})
	return file_producer_proto_rawDescData
}

var file_producer_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_producer_proto_goTypes = []any{
	(*ResourceRequest)(nil),  // 0: producer.ResourceRequest
	(*ResourceResponse)(nil), // 1: producer.ResourceResponse
}
var file_producer_proto_depIdxs = []int32{
	0, // 0: producer.Producer.GetResource:input_type -> producer.ResourceRequest
	1, // 1: producer.Producer.GetResource:output_type -> producer.ResourceResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_producer_proto_init() }
func file_producer_proto_init() {
	if File_producer_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_producer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_producer_proto_goTypes,
		DependencyIndexes: file_producer_proto_depIdxs,
		MessageInfos:      file_producer_proto_msgTypes,
	}.Build()
	File_producer_proto = out.File
	file_producer_proto_rawDesc = nil
	file_producer_proto_goTypes = nil
	file_producer_proto_depIdxs = nil
}
