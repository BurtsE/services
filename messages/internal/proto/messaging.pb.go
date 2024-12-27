// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.1
// 	protoc        v3.12.4
// source: messaging.proto

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

type CreateMessage struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Text          string                 `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	Number        string                 `protobuf:"bytes,2,opt,name=number,proto3" json:"number,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateMessage) Reset() {
	*x = CreateMessage{}
	mi := &file_messaging_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateMessage) ProtoMessage() {}

func (x *CreateMessage) ProtoReflect() protoreflect.Message {
	mi := &file_messaging_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateMessage.ProtoReflect.Descriptor instead.
func (*CreateMessage) Descriptor() ([]byte, []int) {
	return file_messaging_proto_rawDescGZIP(), []int{0}
}

func (x *CreateMessage) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *CreateMessage) GetNumber() string {
	if x != nil {
		return x.Number
	}
	return ""
}

type MessageResponce struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Ok            bool                   `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MessageResponce) Reset() {
	*x = MessageResponce{}
	mi := &file_messaging_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MessageResponce) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageResponce) ProtoMessage() {}

func (x *MessageResponce) ProtoReflect() protoreflect.Message {
	mi := &file_messaging_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageResponce.ProtoReflect.Descriptor instead.
func (*MessageResponce) Descriptor() ([]byte, []int) {
	return file_messaging_proto_rawDescGZIP(), []int{1}
}

func (x *MessageResponce) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

var File_messaging_proto protoreflect.FileDescriptor

var file_messaging_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x3b, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x21,
	0x0a, 0x0f, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x63,
	0x65, 0x12, 0x0e, 0x0a, 0x02, 0x6f, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x02, 0x6f,
	0x6b, 0x32, 0x41, 0x0a, 0x0e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x2f, 0x0a, 0x0b, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x0e, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x1a, 0x10, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x63, 0x65, 0x42, 0x1a, 0x5a, 0x18, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_messaging_proto_rawDescOnce sync.Once
	file_messaging_proto_rawDescData = file_messaging_proto_rawDesc
)

func file_messaging_proto_rawDescGZIP() []byte {
	file_messaging_proto_rawDescOnce.Do(func() {
		file_messaging_proto_rawDescData = protoimpl.X.CompressGZIP(file_messaging_proto_rawDescData)
	})
	return file_messaging_proto_rawDescData
}

var file_messaging_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_messaging_proto_goTypes = []any{
	(*CreateMessage)(nil),   // 0: CreateMessage
	(*MessageResponce)(nil), // 1: MessageResponce
}
var file_messaging_proto_depIdxs = []int32{
	0, // 0: MessageService.SendMessage:input_type -> CreateMessage
	1, // 1: MessageService.SendMessage:output_type -> MessageResponce
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_messaging_proto_init() }
func file_messaging_proto_init() {
	if File_messaging_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_messaging_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_messaging_proto_goTypes,
		DependencyIndexes: file_messaging_proto_depIdxs,
		MessageInfos:      file_messaging_proto_msgTypes,
	}.Build()
	File_messaging_proto = out.File
	file_messaging_proto_rawDesc = nil
	file_messaging_proto_goTypes = nil
	file_messaging_proto_depIdxs = nil
}
