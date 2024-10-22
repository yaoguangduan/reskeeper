// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.2
// source: resource.options.proto

package resproto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ResMarshalFormat int32

const (
	ResMarshalFormat_Bin  ResMarshalFormat = 0
	ResMarshalFormat_Text ResMarshalFormat = 1
	ResMarshalFormat_Json ResMarshalFormat = 2
)

// Enum value maps for ResMarshalFormat.
var (
	ResMarshalFormat_name = map[int32]string{
		0: "Bin",
		1: "Text",
		2: "Json",
	}
	ResMarshalFormat_value = map[string]int32{
		"Bin":  0,
		"Text": 1,
		"Json": 2,
	}
)

func (x ResMarshalFormat) Enum() *ResMarshalFormat {
	p := new(ResMarshalFormat)
	*p = x
	return p
}

func (x ResMarshalFormat) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ResMarshalFormat) Descriptor() protoreflect.EnumDescriptor {
	return file_resource_options_proto_enumTypes[0].Descriptor()
}

func (ResMarshalFormat) Type() protoreflect.EnumType {
	return &file_resource_options_proto_enumTypes[0]
}

func (x ResMarshalFormat) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ResMarshalFormat.Descriptor instead.
func (ResMarshalFormat) EnumDescriptor() ([]byte, []int) {
	return file_resource_options_proto_rawDescGZIP(), []int{0}
}

type ResourceFileOpt struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ExcelPath      *string            `protobuf:"bytes,1,opt,name=excel_path,json=excelPath,proto3,oneof" json:"excel_path,omitempty"`
	MarshalPath    *string            `protobuf:"bytes,2,opt,name=marshal_path,json=marshalPath,proto3,oneof" json:"marshal_path,omitempty"`
	MarshalFormats []ResMarshalFormat `protobuf:"varint,3,rep,packed,name=marshal_formats,json=marshalFormats,proto3,enum=ResMarshalFormat" json:"marshal_formats,omitempty"`
}

func (x *ResourceFileOpt) Reset() {
	*x = ResourceFileOpt{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resource_options_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResourceFileOpt) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResourceFileOpt) ProtoMessage() {}

func (x *ResourceFileOpt) ProtoReflect() protoreflect.Message {
	mi := &file_resource_options_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResourceFileOpt.ProtoReflect.Descriptor instead.
func (*ResourceFileOpt) Descriptor() ([]byte, []int) {
	return file_resource_options_proto_rawDescGZIP(), []int{0}
}

func (x *ResourceFileOpt) GetExcelPath() string {
	if x != nil && x.ExcelPath != nil {
		return *x.ExcelPath
	}
	return ""
}

func (x *ResourceFileOpt) GetMarshalPath() string {
	if x != nil && x.MarshalPath != nil {
		return *x.MarshalPath
	}
	return ""
}

func (x *ResourceFileOpt) GetMarshalFormats() []ResMarshalFormat {
	if x != nil {
		return x.MarshalFormats
	}
	return nil
}

type ResourceTableOpt struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ExcelAndSheetName  *string  `protobuf:"bytes,1,opt,name=excel_and_sheet_name,json=excelAndSheetName,proto3,oneof" json:"excel_and_sheet_name,omitempty"`
	MarshalPrefix      *string  `protobuf:"bytes,2,opt,name=marshal_prefix,json=marshalPrefix,proto3,oneof" json:"marshal_prefix,omitempty"`
	MarshalTags        []string `protobuf:"bytes,3,rep,name=marshal_tags,json=marshalTags,proto3" json:"marshal_tags,omitempty"`
	ExcelWithFieldType *bool    `protobuf:"varint,4,opt,name=excel_with_field_type,json=excelWithFieldType,proto3,oneof" json:"excel_with_field_type,omitempty"`
}

func (x *ResourceTableOpt) Reset() {
	*x = ResourceTableOpt{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resource_options_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResourceTableOpt) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResourceTableOpt) ProtoMessage() {}

func (x *ResourceTableOpt) ProtoReflect() protoreflect.Message {
	mi := &file_resource_options_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResourceTableOpt.ProtoReflect.Descriptor instead.
func (*ResourceTableOpt) Descriptor() ([]byte, []int) {
	return file_resource_options_proto_rawDescGZIP(), []int{1}
}

func (x *ResourceTableOpt) GetExcelAndSheetName() string {
	if x != nil && x.ExcelAndSheetName != nil {
		return *x.ExcelAndSheetName
	}
	return ""
}

func (x *ResourceTableOpt) GetMarshalPrefix() string {
	if x != nil && x.MarshalPrefix != nil {
		return *x.MarshalPrefix
	}
	return ""
}

func (x *ResourceTableOpt) GetMarshalTags() []string {
	if x != nil {
		return x.MarshalTags
	}
	return nil
}

func (x *ResourceTableOpt) GetExcelWithFieldType() bool {
	if x != nil && x.ExcelWithFieldType != nil {
		return *x.ExcelWithFieldType
	}
	return false
}

type ResourceMsgOpt struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TagIgnoreFields []string `protobuf:"bytes,1,rep,name=tag_ignore_fields,json=tagIgnoreFields,proto3" json:"tag_ignore_fields,omitempty"`
	MsgKey          *int32   `protobuf:"varint,2,opt,name=msg_key,json=msgKey,proto3,oneof" json:"msg_key,omitempty"`
}

func (x *ResourceMsgOpt) Reset() {
	*x = ResourceMsgOpt{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resource_options_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResourceMsgOpt) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResourceMsgOpt) ProtoMessage() {}

func (x *ResourceMsgOpt) ProtoReflect() protoreflect.Message {
	mi := &file_resource_options_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResourceMsgOpt.ProtoReflect.Descriptor instead.
func (*ResourceMsgOpt) Descriptor() ([]byte, []int) {
	return file_resource_options_proto_rawDescGZIP(), []int{2}
}

func (x *ResourceMsgOpt) GetTagIgnoreFields() []string {
	if x != nil {
		return x.TagIgnoreFields
	}
	return nil
}

func (x *ResourceMsgOpt) GetMsgKey() int32 {
	if x != nil && x.MsgKey != nil {
		return *x.MsgKey
	}
	return 0
}

type ResourceFieldValidate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NotNull *bool   `protobuf:"varint,1,opt,name=not_null,json=notNull,proto3,oneof" json:"not_null,omitempty"`
	Uniq    *bool   `protobuf:"varint,2,opt,name=uniq,proto3,oneof" json:"uniq,omitempty"`
	Length  *string `protobuf:"bytes,3,opt,name=length,proto3,oneof" json:"length,omitempty"`
	Range   *string `protobuf:"bytes,4,opt,name=range,proto3,oneof" json:"range,omitempty"`
	Pattern *string `protobuf:"bytes,5,opt,name=pattern,proto3,oneof" json:"pattern,omitempty"`
}

func (x *ResourceFieldValidate) Reset() {
	*x = ResourceFieldValidate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resource_options_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResourceFieldValidate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResourceFieldValidate) ProtoMessage() {}

func (x *ResourceFieldValidate) ProtoReflect() protoreflect.Message {
	mi := &file_resource_options_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResourceFieldValidate.ProtoReflect.Descriptor instead.
func (*ResourceFieldValidate) Descriptor() ([]byte, []int) {
	return file_resource_options_proto_rawDescGZIP(), []int{3}
}

func (x *ResourceFieldValidate) GetNotNull() bool {
	if x != nil && x.NotNull != nil {
		return *x.NotNull
	}
	return false
}

func (x *ResourceFieldValidate) GetUniq() bool {
	if x != nil && x.Uniq != nil {
		return *x.Uniq
	}
	return false
}

func (x *ResourceFieldValidate) GetLength() string {
	if x != nil && x.Length != nil {
		return *x.Length
	}
	return ""
}

func (x *ResourceFieldValidate) GetRange() string {
	if x != nil && x.Range != nil {
		return *x.Range
	}
	return ""
}

func (x *ResourceFieldValidate) GetPattern() string {
	if x != nil && x.Pattern != nil {
		return *x.Pattern
	}
	return ""
}

var file_resource_options_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.FileOptions)(nil),
		ExtensionType: (*ResourceFileOpt)(nil),
		Field:         50100,
		Name:          "res_file_opt",
		Tag:           "bytes,50100,opt,name=res_file_opt",
		Filename:      "resource.options.proto",
	},
	{
		ExtendedType:  (*descriptorpb.MessageOptions)(nil),
		ExtensionType: (*ResourceTableOpt)(nil),
		Field:         50110,
		Name:          "res_table_opt",
		Tag:           "bytes,50110,opt,name=res_table_opt",
		Filename:      "resource.options.proto",
	},
	{
		ExtendedType:  (*descriptorpb.MessageOptions)(nil),
		ExtensionType: (*ResourceMsgOpt)(nil),
		Field:         50111,
		Name:          "res_msg_opt",
		Tag:           "bytes,50111,opt,name=res_msg_opt",
		Filename:      "resource.options.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*ResourceFieldValidate)(nil),
		Field:         50120,
		Name:          "res_validate",
		Tag:           "bytes,50120,opt,name=res_validate",
		Filename:      "resource.options.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*bool)(nil),
		Field:         50121,
		Name:          "res_one_column",
		Tag:           "varint,50121,opt,name=res_one_column",
		Filename:      "resource.options.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*bool)(nil),
		Field:         50122,
		Name:          "res_use_msg_key",
		Tag:           "varint,50122,opt,name=res_use_msg_key",
		Filename:      "resource.options.proto",
	},
}

// Extension fields to descriptorpb.FileOptions.
var (
	// optional ResourceFileOpt res_file_opt = 50100;
	E_ResFileOpt = &file_resource_options_proto_extTypes[0]
)

// Extension fields to descriptorpb.MessageOptions.
var (
	// optional ResourceTableOpt res_table_opt = 50110;
	E_ResTableOpt = &file_resource_options_proto_extTypes[1]
	// optional ResourceMsgOpt res_msg_opt = 50111;
	E_ResMsgOpt = &file_resource_options_proto_extTypes[2]
)

// Extension fields to descriptorpb.FieldOptions.
var (
	// optional ResourceFieldValidate res_validate = 50120;
	E_ResValidate = &file_resource_options_proto_extTypes[3]
	// optional bool res_one_column = 50121;
	E_ResOneColumn = &file_resource_options_proto_extTypes[4]
	// optional bool res_use_msg_key = 50122;
	E_ResUseMsgKey = &file_resource_options_proto_extTypes[5]
)

var File_resource_options_proto protoreflect.FileDescriptor

var file_resource_options_proto_rawDesc = []byte{
	0x0a, 0x16, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb9, 0x01, 0x0a, 0x0f, 0x52,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x4f, 0x70, 0x74, 0x12, 0x22,
	0x0a, 0x0a, 0x65, 0x78, 0x63, 0x65, 0x6c, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x00, 0x52, 0x09, 0x65, 0x78, 0x63, 0x65, 0x6c, 0x50, 0x61, 0x74, 0x68, 0x88,
	0x01, 0x01, 0x12, 0x26, 0x0a, 0x0c, 0x6d, 0x61, 0x72, 0x73, 0x68, 0x61, 0x6c, 0x5f, 0x70, 0x61,
	0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x0b, 0x6d, 0x61, 0x72, 0x73,
	0x68, 0x61, 0x6c, 0x50, 0x61, 0x74, 0x68, 0x88, 0x01, 0x01, 0x12, 0x3a, 0x0a, 0x0f, 0x6d, 0x61,
	0x72, 0x73, 0x68, 0x61, 0x6c, 0x5f, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x73, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x0e, 0x32, 0x11, 0x2e, 0x52, 0x65, 0x73, 0x4d, 0x61, 0x72, 0x73, 0x68, 0x61, 0x6c,
	0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x52, 0x0e, 0x6d, 0x61, 0x72, 0x73, 0x68, 0x61, 0x6c, 0x46,
	0x6f, 0x72, 0x6d, 0x61, 0x74, 0x73, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x65, 0x78, 0x63, 0x65, 0x6c,
	0x5f, 0x70, 0x61, 0x74, 0x68, 0x42, 0x0f, 0x0a, 0x0d, 0x5f, 0x6d, 0x61, 0x72, 0x73, 0x68, 0x61,
	0x6c, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x22, 0x95, 0x02, 0x0a, 0x10, 0x52, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x4f, 0x70, 0x74, 0x12, 0x34, 0x0a, 0x14, 0x65,
	0x78, 0x63, 0x65, 0x6c, 0x5f, 0x61, 0x6e, 0x64, 0x5f, 0x73, 0x68, 0x65, 0x65, 0x74, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x11, 0x65, 0x78, 0x63,
	0x65, 0x6c, 0x41, 0x6e, 0x64, 0x53, 0x68, 0x65, 0x65, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x88, 0x01,
	0x01, 0x12, 0x2a, 0x0a, 0x0e, 0x6d, 0x61, 0x72, 0x73, 0x68, 0x61, 0x6c, 0x5f, 0x70, 0x72, 0x65,
	0x66, 0x69, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x0d, 0x6d, 0x61, 0x72,
	0x73, 0x68, 0x61, 0x6c, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x88, 0x01, 0x01, 0x12, 0x21, 0x0a,
	0x0c, 0x6d, 0x61, 0x72, 0x73, 0x68, 0x61, 0x6c, 0x5f, 0x74, 0x61, 0x67, 0x73, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x0b, 0x6d, 0x61, 0x72, 0x73, 0x68, 0x61, 0x6c, 0x54, 0x61, 0x67, 0x73,
	0x12, 0x36, 0x0a, 0x15, 0x65, 0x78, 0x63, 0x65, 0x6c, 0x5f, 0x77, 0x69, 0x74, 0x68, 0x5f, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x48,
	0x02, 0x52, 0x12, 0x65, 0x78, 0x63, 0x65, 0x6c, 0x57, 0x69, 0x74, 0x68, 0x46, 0x69, 0x65, 0x6c,
	0x64, 0x54, 0x79, 0x70, 0x65, 0x88, 0x01, 0x01, 0x42, 0x17, 0x0a, 0x15, 0x5f, 0x65, 0x78, 0x63,
	0x65, 0x6c, 0x5f, 0x61, 0x6e, 0x64, 0x5f, 0x73, 0x68, 0x65, 0x65, 0x74, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x42, 0x11, 0x0a, 0x0f, 0x5f, 0x6d, 0x61, 0x72, 0x73, 0x68, 0x61, 0x6c, 0x5f, 0x70, 0x72,
	0x65, 0x66, 0x69, 0x78, 0x42, 0x18, 0x0a, 0x16, 0x5f, 0x65, 0x78, 0x63, 0x65, 0x6c, 0x5f, 0x77,
	0x69, 0x74, 0x68, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x22, 0x66,
	0x0a, 0x0e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4d, 0x73, 0x67, 0x4f, 0x70, 0x74,
	0x12, 0x2a, 0x0a, 0x11, 0x74, 0x61, 0x67, 0x5f, 0x69, 0x67, 0x6e, 0x6f, 0x72, 0x65, 0x5f, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0f, 0x74, 0x61, 0x67,
	0x49, 0x67, 0x6e, 0x6f, 0x72, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x12, 0x1c, 0x0a, 0x07,
	0x6d, 0x73, 0x67, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x48, 0x00, 0x52,
	0x06, 0x6d, 0x73, 0x67, 0x4b, 0x65, 0x79, 0x88, 0x01, 0x01, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x6d,
	0x73, 0x67, 0x5f, 0x6b, 0x65, 0x79, 0x22, 0xde, 0x01, 0x0a, 0x15, 0x52, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x12, 0x1e, 0x0a, 0x08, 0x6e, 0x6f, 0x74, 0x5f, 0x6e, 0x75, 0x6c, 0x6c, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x48, 0x00, 0x52, 0x07, 0x6e, 0x6f, 0x74, 0x4e, 0x75, 0x6c, 0x6c, 0x88, 0x01, 0x01,
	0x12, 0x17, 0x0a, 0x04, 0x75, 0x6e, 0x69, 0x71, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x48, 0x01,
	0x52, 0x04, 0x75, 0x6e, 0x69, 0x71, 0x88, 0x01, 0x01, 0x12, 0x1b, 0x0a, 0x06, 0x6c, 0x65, 0x6e,
	0x67, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x06, 0x6c, 0x65, 0x6e,
	0x67, 0x74, 0x68, 0x88, 0x01, 0x01, 0x12, 0x19, 0x0a, 0x05, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x03, 0x52, 0x05, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x88, 0x01,
	0x01, 0x12, 0x1d, 0x0a, 0x07, 0x70, 0x61, 0x74, 0x74, 0x65, 0x72, 0x6e, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x04, 0x52, 0x07, 0x70, 0x61, 0x74, 0x74, 0x65, 0x72, 0x6e, 0x88, 0x01, 0x01,
	0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x6e, 0x6f, 0x74, 0x5f, 0x6e, 0x75, 0x6c, 0x6c, 0x42, 0x07, 0x0a,
	0x05, 0x5f, 0x75, 0x6e, 0x69, 0x71, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x6c, 0x65, 0x6e, 0x67, 0x74,
	0x68, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x42, 0x0a, 0x0a, 0x08, 0x5f,
	0x70, 0x61, 0x74, 0x74, 0x65, 0x72, 0x6e, 0x2a, 0x2f, 0x0a, 0x10, 0x52, 0x65, 0x73, 0x4d, 0x61,
	0x72, 0x73, 0x68, 0x61, 0x6c, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x12, 0x07, 0x0a, 0x03, 0x42,
	0x69, 0x6e, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x54, 0x65, 0x78, 0x74, 0x10, 0x01, 0x12, 0x08,
	0x0a, 0x04, 0x4a, 0x73, 0x6f, 0x6e, 0x10, 0x02, 0x3a, 0x55, 0x0a, 0x0c, 0x72, 0x65, 0x73, 0x5f,
	0x66, 0x69, 0x6c, 0x65, 0x5f, 0x6f, 0x70, 0x74, 0x12, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x4f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xb4, 0x87, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10,
	0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x4f, 0x70, 0x74,
	0x52, 0x0a, 0x72, 0x65, 0x73, 0x46, 0x69, 0x6c, 0x65, 0x4f, 0x70, 0x74, 0x88, 0x01, 0x01, 0x3a,
	0x5b, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x5f, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x6f, 0x70, 0x74,
	0x12, 0x1f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x18, 0xbe, 0x87, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x52, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x4f, 0x70, 0x74, 0x52, 0x0b, 0x72, 0x65,
	0x73, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x4f, 0x70, 0x74, 0x88, 0x01, 0x01, 0x3a, 0x55, 0x0a, 0x0b,
	0x72, 0x65, 0x73, 0x5f, 0x6d, 0x73, 0x67, 0x5f, 0x6f, 0x70, 0x74, 0x12, 0x1f, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xbf, 0x87, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4d,
	0x73, 0x67, 0x4f, 0x70, 0x74, 0x52, 0x09, 0x72, 0x65, 0x73, 0x4d, 0x73, 0x67, 0x4f, 0x70, 0x74,
	0x88, 0x01, 0x01, 0x3a, 0x5d, 0x0a, 0x0c, 0x72, 0x65, 0x73, 0x5f, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x18, 0xc8, 0x87, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x52, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x52, 0x0b, 0x72, 0x65, 0x73, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x88,
	0x01, 0x01, 0x3a, 0x48, 0x0a, 0x0e, 0x72, 0x65, 0x73, 0x5f, 0x6f, 0x6e, 0x65, 0x5f, 0x63, 0x6f,
	0x6c, 0x75, 0x6d, 0x6e, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0xc9, 0x87, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x72, 0x65, 0x73,
	0x4f, 0x6e, 0x65, 0x43, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x88, 0x01, 0x01, 0x3a, 0x49, 0x0a, 0x0f,
	0x72, 0x65, 0x73, 0x5f, 0x75, 0x73, 0x65, 0x5f, 0x6d, 0x73, 0x67, 0x5f, 0x6b, 0x65, 0x79, 0x12,
	0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xca,
	0x87, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x55, 0x73, 0x65, 0x4d, 0x73,
	0x67, 0x4b, 0x65, 0x79, 0x88, 0x01, 0x01, 0x42, 0x2c, 0x5a, 0x2a, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x79, 0x61, 0x6f, 0x67, 0x75, 0x61, 0x6e, 0x67, 0x64, 0x75,
	0x61, 0x6e, 0x2f, 0x72, 0x65, 0x73, 0x6b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x2f, 0x72, 0x65, 0x73,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_resource_options_proto_rawDescOnce sync.Once
	file_resource_options_proto_rawDescData = file_resource_options_proto_rawDesc
)

func file_resource_options_proto_rawDescGZIP() []byte {
	file_resource_options_proto_rawDescOnce.Do(func() {
		file_resource_options_proto_rawDescData = protoimpl.X.CompressGZIP(file_resource_options_proto_rawDescData)
	})
	return file_resource_options_proto_rawDescData
}

var file_resource_options_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_resource_options_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_resource_options_proto_goTypes = []any{
	(ResMarshalFormat)(0),               // 0: ResMarshalFormat
	(*ResourceFileOpt)(nil),             // 1: ResourceFileOpt
	(*ResourceTableOpt)(nil),            // 2: ResourceTableOpt
	(*ResourceMsgOpt)(nil),              // 3: ResourceMsgOpt
	(*ResourceFieldValidate)(nil),       // 4: ResourceFieldValidate
	(*descriptorpb.FileOptions)(nil),    // 5: google.protobuf.FileOptions
	(*descriptorpb.MessageOptions)(nil), // 6: google.protobuf.MessageOptions
	(*descriptorpb.FieldOptions)(nil),   // 7: google.protobuf.FieldOptions
}
var file_resource_options_proto_depIdxs = []int32{
	0,  // 0: ResourceFileOpt.marshal_formats:type_name -> ResMarshalFormat
	5,  // 1: res_file_opt:extendee -> google.protobuf.FileOptions
	6,  // 2: res_table_opt:extendee -> google.protobuf.MessageOptions
	6,  // 3: res_msg_opt:extendee -> google.protobuf.MessageOptions
	7,  // 4: res_validate:extendee -> google.protobuf.FieldOptions
	7,  // 5: res_one_column:extendee -> google.protobuf.FieldOptions
	7,  // 6: res_use_msg_key:extendee -> google.protobuf.FieldOptions
	1,  // 7: res_file_opt:type_name -> ResourceFileOpt
	2,  // 8: res_table_opt:type_name -> ResourceTableOpt
	3,  // 9: res_msg_opt:type_name -> ResourceMsgOpt
	4,  // 10: res_validate:type_name -> ResourceFieldValidate
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	7,  // [7:11] is the sub-list for extension type_name
	1,  // [1:7] is the sub-list for extension extendee
	0,  // [0:1] is the sub-list for field type_name
}

func init() { file_resource_options_proto_init() }
func file_resource_options_proto_init() {
	if File_resource_options_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_resource_options_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*ResourceFileOpt); i {
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
		file_resource_options_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*ResourceTableOpt); i {
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
		file_resource_options_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*ResourceMsgOpt); i {
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
		file_resource_options_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*ResourceFieldValidate); i {
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
	file_resource_options_proto_msgTypes[0].OneofWrappers = []any{}
	file_resource_options_proto_msgTypes[1].OneofWrappers = []any{}
	file_resource_options_proto_msgTypes[2].OneofWrappers = []any{}
	file_resource_options_proto_msgTypes[3].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_resource_options_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 6,
			NumServices:   0,
		},
		GoTypes:           file_resource_options_proto_goTypes,
		DependencyIndexes: file_resource_options_proto_depIdxs,
		EnumInfos:         file_resource_options_proto_enumTypes,
		MessageInfos:      file_resource_options_proto_msgTypes,
		ExtensionInfos:    file_resource_options_proto_extTypes,
	}.Build()
	File_resource_options_proto = out.File
	file_resource_options_proto_rawDesc = nil
	file_resource_options_proto_goTypes = nil
	file_resource_options_proto_depIdxs = nil
}
