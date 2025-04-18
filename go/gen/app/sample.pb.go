// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: app/sample.proto

package appv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Sample struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Sample) Reset() {
	*x = Sample{}
	mi := &file_app_sample_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Sample) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Sample) ProtoMessage() {}

func (x *Sample) ProtoReflect() protoreflect.Message {
	mi := &file_app_sample_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Sample.ProtoReflect.Descriptor instead.
func (*Sample) Descriptor() ([]byte, []int) {
	return file_app_sample_proto_rawDescGZIP(), []int{0}
}

func (x *Sample) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Sample) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_app_sample_proto protoreflect.FileDescriptor

const file_app_sample_proto_rawDesc = "" +
	"\n" +
	"\x10app/sample.proto\x12\x06app.v1\",\n" +
	"\x06Sample\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x12\n" +
	"\x04name\x18\x02 \x01(\tR\x04nameB\x7f\n" +
	"\n" +
	"com.app.v1B\vSampleProtoP\x01Z+github.com/hiroky1983/talk/go/gen/app;appv1\xa2\x02\x03AXX\xaa\x02\x06App.V1\xca\x02\x06App\\V1\xe2\x02\x12App\\V1\\GPBMetadata\xea\x02\aApp::V1b\x06proto3"

var (
	file_app_sample_proto_rawDescOnce sync.Once
	file_app_sample_proto_rawDescData []byte
)

func file_app_sample_proto_rawDescGZIP() []byte {
	file_app_sample_proto_rawDescOnce.Do(func() {
		file_app_sample_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_app_sample_proto_rawDesc), len(file_app_sample_proto_rawDesc)))
	})
	return file_app_sample_proto_rawDescData
}

var file_app_sample_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_app_sample_proto_goTypes = []any{
	(*Sample)(nil), // 0: app.v1.Sample
}
var file_app_sample_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_app_sample_proto_init() }
func file_app_sample_proto_init() {
	if File_app_sample_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_app_sample_proto_rawDesc), len(file_app_sample_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_app_sample_proto_goTypes,
		DependencyIndexes: file_app_sample_proto_depIdxs,
		MessageInfos:      file_app_sample_proto_msgTypes,
	}.Build()
	File_app_sample_proto = out.File
	file_app_sample_proto_goTypes = nil
	file_app_sample_proto_depIdxs = nil
}
