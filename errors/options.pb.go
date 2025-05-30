// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v4.25.3
// source: options.proto

package errorspb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var file_options_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.EnumValueOptions)(nil),
		ExtensionType: (*int32)(nil),
		Field:         1109,
		Name:          "errors.code",
		Tag:           "varint,1109,opt,name=code",
		Filename:      "options.proto",
	},
	{
		ExtendedType:  (*descriptorpb.EnumOptions)(nil),
		ExtensionType: (*int32)(nil),
		Field:         1108,
		Name:          "errors.default_code",
		Tag:           "varint,1108,opt,name=default_code",
		Filename:      "options.proto",
	},
}

// Extension fields to descriptorpb.EnumValueOptions.
var (
	// optional int32 code = 1109;
	E_Code = &file_options_proto_extTypes[0]
)

// Extension fields to descriptorpb.EnumOptions.
var (
	// optional int32 default_code = 1108;
	E_DefaultCode = &file_options_proto_extTypes[1]
)

var File_options_proto protoreflect.FileDescriptor

const file_options_proto_rawDesc = "" +
	"\n" +
	"\roptions.proto\x12\x06errors\x1a google/protobuf/descriptor.proto:6\n" +
	"\x04code\x12!.google.protobuf.EnumValueOptions\x18\xd5\b \x01(\x05R\x04code:@\n" +
	"\fdefault_code\x12\x1c.google.protobuf.EnumOptions\x18\xd4\b \x01(\x05R\vdefaultCodeB:Z8github.com/honeybbq/go-zero-errors-proto/errors;errorspbb\x06proto3"

var file_options_proto_goTypes = []any{
	(*descriptorpb.EnumValueOptions)(nil), // 0: google.protobuf.EnumValueOptions
	(*descriptorpb.EnumOptions)(nil),      // 1: google.protobuf.EnumOptions
}
var file_options_proto_depIdxs = []int32{
	0, // 0: errors.code:extendee -> google.protobuf.EnumValueOptions
	1, // 1: errors.default_code:extendee -> google.protobuf.EnumOptions
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	0, // [0:2] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_options_proto_init() }
func file_options_proto_init() {
	if File_options_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_options_proto_rawDesc), len(file_options_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 2,
			NumServices:   0,
		},
		GoTypes:           file_options_proto_goTypes,
		DependencyIndexes: file_options_proto_depIdxs,
		ExtensionInfos:    file_options_proto_extTypes,
	}.Build()
	File_options_proto = out.File
	file_options_proto_goTypes = nil
	file_options_proto_depIdxs = nil
}
