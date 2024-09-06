// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: dummy/v1/message.proto

package dummyv1

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

type Dummy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string        `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Value    int32         `protobuf:"varint,2,opt,name=value,proto3" json:"value,omitempty"`
	Values   []string      `protobuf:"bytes,3,rep,name=values,proto3" json:"values,omitempty"`
	TestType TestEnumType  `protobuf:"varint,4,opt,name=test_type,json=testType,proto3,enum=dummy.v1.TestEnumType" json:"test_type,omitempty"`
	ConfigA  *DummyConfigA `protobuf:"bytes,11,opt,name=config_a,json=configA,proto3" json:"config_a,omitempty"`
	ConfigB  *DummyConfigB `protobuf:"bytes,12,opt,name=config_b,json=configB,proto3" json:"config_b,omitempty"`
}

func (x *Dummy) Reset() {
	*x = Dummy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dummy_v1_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Dummy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Dummy) ProtoMessage() {}

func (x *Dummy) ProtoReflect() protoreflect.Message {
	mi := &file_dummy_v1_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Dummy.ProtoReflect.Descriptor instead.
func (*Dummy) Descriptor() ([]byte, []int) {
	return file_dummy_v1_message_proto_rawDescGZIP(), []int{0}
}

func (x *Dummy) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Dummy) GetValue() int32 {
	if x != nil {
		return x.Value
	}
	return 0
}

func (x *Dummy) GetValues() []string {
	if x != nil {
		return x.Values
	}
	return nil
}

func (x *Dummy) GetTestType() TestEnumType {
	if x != nil {
		return x.TestType
	}
	return TestEnumType_TEST_ENUM_TYPE_UNSPECIFIED
}

func (x *Dummy) GetConfigA() *DummyConfigA {
	if x != nil {
		return x.ConfigA
	}
	return nil
}

func (x *Dummy) GetConfigB() *DummyConfigB {
	if x != nil {
		return x.ConfigB
	}
	return nil
}

type DummyConfigA struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *DummyConfigA) Reset() {
	*x = DummyConfigA{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dummy_v1_message_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DummyConfigA) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DummyConfigA) ProtoMessage() {}

func (x *DummyConfigA) ProtoReflect() protoreflect.Message {
	mi := &file_dummy_v1_message_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DummyConfigA.ProtoReflect.Descriptor instead.
func (*DummyConfigA) Descriptor() ([]byte, []int) {
	return file_dummy_v1_message_proto_rawDescGZIP(), []int{1}
}

func (x *DummyConfigA) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DummyConfigA) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type DummyConfigB struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Value int64  `protobuf:"varint,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *DummyConfigB) Reset() {
	*x = DummyConfigB{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dummy_v1_message_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DummyConfigB) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DummyConfigB) ProtoMessage() {}

func (x *DummyConfigB) ProtoReflect() protoreflect.Message {
	mi := &file_dummy_v1_message_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DummyConfigB.ProtoReflect.Descriptor instead.
func (*DummyConfigB) Descriptor() ([]byte, []int) {
	return file_dummy_v1_message_proto_rawDescGZIP(), []int{2}
}

func (x *DummyConfigB) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DummyConfigB) GetValue() int64 {
	if x != nil {
		return x.Value
	}
	return 0
}

var File_dummy_v1_message_proto protoreflect.FileDescriptor

var file_dummy_v1_message_proto_rawDesc = []byte{
	0x0a, 0x16, 0x64, 0x75, 0x6d, 0x6d, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x64, 0x75, 0x6d, 0x6d, 0x79, 0x2e,
	0x76, 0x31, 0x1a, 0x13, 0x64, 0x75, 0x6d, 0x6d, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6e, 0x75,
	0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe4, 0x01, 0x0a, 0x05, 0x44, 0x75, 0x6d, 0x6d,
	0x79, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x73, 0x12, 0x33, 0x0a, 0x09, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x64, 0x75, 0x6d, 0x6d, 0x79, 0x2e, 0x76,
	0x31, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x45, 0x6e, 0x75, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x52, 0x08,
	0x74, 0x65, 0x73, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x31, 0x0a, 0x08, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x5f, 0x61, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x64, 0x75, 0x6d,
	0x6d, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x75, 0x6d, 0x6d, 0x79, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x41, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x41, 0x12, 0x31, 0x0a, 0x08, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x62, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e,
	0x64, 0x75, 0x6d, 0x6d, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x75, 0x6d, 0x6d, 0x79, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x42, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x42, 0x22, 0x38,
	0x0a, 0x0c, 0x44, 0x75, 0x6d, 0x6d, 0x79, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x41, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x38, 0x0a, 0x0c, 0x44, 0x75, 0x6d, 0x6d,
	0x79, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x42, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x42, 0x9a, 0x01, 0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x75, 0x6d, 0x6d, 0x79,
	0x2e, 0x76, 0x31, 0x42, 0x0c, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x3b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x67, 0x6f, 0x63, 0x68, 0x6f, 0x72, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x6d, 0x61, 0x67,
	0x69, 0x63, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x67, 0x65, 0x6e, 0x2f,
	0x64, 0x75, 0x6d, 0x6d, 0x79, 0x2f, 0x76, 0x31, 0x3b, 0x64, 0x75, 0x6d, 0x6d, 0x79, 0x76, 0x31,
	0xa2, 0x02, 0x03, 0x44, 0x58, 0x58, 0xaa, 0x02, 0x08, 0x44, 0x75, 0x6d, 0x6d, 0x79, 0x2e, 0x56,
	0x31, 0xca, 0x02, 0x08, 0x44, 0x75, 0x6d, 0x6d, 0x79, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x14, 0x44,
	0x75, 0x6d, 0x6d, 0x79, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0xea, 0x02, 0x09, 0x44, 0x75, 0x6d, 0x6d, 0x79, 0x3a, 0x3a, 0x56, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_dummy_v1_message_proto_rawDescOnce sync.Once
	file_dummy_v1_message_proto_rawDescData = file_dummy_v1_message_proto_rawDesc
)

func file_dummy_v1_message_proto_rawDescGZIP() []byte {
	file_dummy_v1_message_proto_rawDescOnce.Do(func() {
		file_dummy_v1_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_dummy_v1_message_proto_rawDescData)
	})
	return file_dummy_v1_message_proto_rawDescData
}

var file_dummy_v1_message_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_dummy_v1_message_proto_goTypes = []any{
	(*Dummy)(nil),        // 0: dummy.v1.Dummy
	(*DummyConfigA)(nil), // 1: dummy.v1.DummyConfigA
	(*DummyConfigB)(nil), // 2: dummy.v1.DummyConfigB
	(TestEnumType)(0),    // 3: dummy.v1.TestEnumType
}
var file_dummy_v1_message_proto_depIdxs = []int32{
	3, // 0: dummy.v1.Dummy.test_type:type_name -> dummy.v1.TestEnumType
	1, // 1: dummy.v1.Dummy.config_a:type_name -> dummy.v1.DummyConfigA
	2, // 2: dummy.v1.Dummy.config_b:type_name -> dummy.v1.DummyConfigB
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_dummy_v1_message_proto_init() }
func file_dummy_v1_message_proto_init() {
	if File_dummy_v1_message_proto != nil {
		return
	}
	file_dummy_v1_enum_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_dummy_v1_message_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Dummy); i {
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
		file_dummy_v1_message_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*DummyConfigA); i {
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
		file_dummy_v1_message_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*DummyConfigB); i {
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
			RawDescriptor: file_dummy_v1_message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_dummy_v1_message_proto_goTypes,
		DependencyIndexes: file_dummy_v1_message_proto_depIdxs,
		MessageInfos:      file_dummy_v1_message_proto_msgTypes,
	}.Build()
	File_dummy_v1_message_proto = out.File
	file_dummy_v1_message_proto_rawDesc = nil
	file_dummy_v1_message_proto_goTypes = nil
	file_dummy_v1_message_proto_depIdxs = nil
}
