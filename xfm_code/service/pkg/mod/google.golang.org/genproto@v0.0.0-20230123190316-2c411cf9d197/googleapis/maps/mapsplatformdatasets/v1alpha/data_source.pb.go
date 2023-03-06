// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.12.2
// source: google/maps/mapsplatformdatasets/v1alpha/data_source.proto

package mapsplatformdatasets

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// The format of the file being uploaded.
type FileFormat int32

const (
	// Unspecified file format.
	FileFormat_FILE_FORMAT_UNSPECIFIED FileFormat = 0
	// GeoJson file.
	FileFormat_FILE_FORMAT_GEOJSON FileFormat = 1
	// KML file.
	FileFormat_FILE_FORMAT_KML FileFormat = 2
	// CSV file.
	FileFormat_FILE_FORMAT_CSV FileFormat = 3
	// Protobuf file.
	FileFormat_FILE_FORMAT_PROTO FileFormat = 4
	// KMZ file.
	FileFormat_FILE_FORMAT_KMZ FileFormat = 5
)

// Enum value maps for FileFormat.
var (
	FileFormat_name = map[int32]string{
		0: "FILE_FORMAT_UNSPECIFIED",
		1: "FILE_FORMAT_GEOJSON",
		2: "FILE_FORMAT_KML",
		3: "FILE_FORMAT_CSV",
		4: "FILE_FORMAT_PROTO",
		5: "FILE_FORMAT_KMZ",
	}
	FileFormat_value = map[string]int32{
		"FILE_FORMAT_UNSPECIFIED": 0,
		"FILE_FORMAT_GEOJSON":     1,
		"FILE_FORMAT_KML":         2,
		"FILE_FORMAT_CSV":         3,
		"FILE_FORMAT_PROTO":       4,
		"FILE_FORMAT_KMZ":         5,
	}
)

func (x FileFormat) Enum() *FileFormat {
	p := new(FileFormat)
	*p = x
	return p
}

func (x FileFormat) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (FileFormat) Descriptor() protoreflect.EnumDescriptor {
	return file_google_maps_mapsplatformdatasets_v1alpha_data_source_proto_enumTypes[0].Descriptor()
}

func (FileFormat) Type() protoreflect.EnumType {
	return &file_google_maps_mapsplatformdatasets_v1alpha_data_source_proto_enumTypes[0]
}

func (x FileFormat) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use FileFormat.Descriptor instead.
func (FileFormat) EnumDescriptor() ([]byte, []int) {
	return file_google_maps_mapsplatformdatasets_v1alpha_data_source_proto_rawDescGZIP(), []int{0}
}

// The details about the data source when it is a local file.
type LocalFileSource struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The file name and extension of the uploaded file.
	Filename string `protobuf:"bytes,1,opt,name=filename,proto3" json:"filename,omitempty"`
	// The format of the file that is being uploaded.
	FileFormat FileFormat `protobuf:"varint,2,opt,name=file_format,json=fileFormat,proto3,enum=google.maps.mapsplatformdatasets.v1alpha.FileFormat" json:"file_format,omitempty"`
}

func (x *LocalFileSource) Reset() {
	*x = LocalFileSource{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_maps_mapsplatformdatasets_v1alpha_data_source_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LocalFileSource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LocalFileSource) ProtoMessage() {}

func (x *LocalFileSource) ProtoReflect() protoreflect.Message {
	mi := &file_google_maps_mapsplatformdatasets_v1alpha_data_source_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LocalFileSource.ProtoReflect.Descriptor instead.
func (*LocalFileSource) Descriptor() ([]byte, []int) {
	return file_google_maps_mapsplatformdatasets_v1alpha_data_source_proto_rawDescGZIP(), []int{0}
}

func (x *LocalFileSource) GetFilename() string {
	if x != nil {
		return x.Filename
	}
	return ""
}

func (x *LocalFileSource) GetFileFormat() FileFormat {
	if x != nil {
		return x.FileFormat
	}
	return FileFormat_FILE_FORMAT_UNSPECIFIED
}

// The details about the data source when it is in Google Cloud Storage.
type GcsSource struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Source data URI. For example, `gs://my_bucket/my_object`.
	InputUri string `protobuf:"bytes,1,opt,name=input_uri,json=inputUri,proto3" json:"input_uri,omitempty"`
	// The file format of the Google Cloud Storage object. This is used mainly for
	// validation.
	FileFormat FileFormat `protobuf:"varint,2,opt,name=file_format,json=fileFormat,proto3,enum=google.maps.mapsplatformdatasets.v1alpha.FileFormat" json:"file_format,omitempty"`
}

func (x *GcsSource) Reset() {
	*x = GcsSource{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_maps_mapsplatformdatasets_v1alpha_data_source_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GcsSource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GcsSource) ProtoMessage() {}

func (x *GcsSource) ProtoReflect() protoreflect.Message {
	mi := &file_google_maps_mapsplatformdatasets_v1alpha_data_source_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GcsSource.ProtoReflect.Descriptor instead.
func (*GcsSource) Descriptor() ([]byte, []int) {
	return file_google_maps_mapsplatformdatasets_v1alpha_data_source_proto_rawDescGZIP(), []int{1}
}

func (x *GcsSource) GetInputUri() string {
	if x != nil {
		return x.InputUri
	}
	return ""
}

func (x *GcsSource) GetFileFormat() FileFormat {
	if x != nil {
		return x.FileFormat
	}
	return FileFormat_FILE_FORMAT_UNSPECIFIED
}

var File_google_maps_mapsplatformdatasets_v1alpha_data_source_proto protoreflect.FileDescriptor

var file_google_maps_mapsplatformdatasets_v1alpha_data_source_proto_rawDesc = []byte{
	0x0a, 0x3a, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x6d, 0x61, 0x70, 0x73, 0x2f, 0x6d, 0x61,
	0x70, 0x73, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65,
	0x74, 0x73, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x5f,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x28, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6d, 0x61, 0x70, 0x73, 0x2e, 0x6d, 0x61, 0x70, 0x73, 0x70, 0x6c,
	0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x73, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x22, 0x84, 0x01, 0x0a, 0x0f, 0x4c, 0x6f, 0x63, 0x61, 0x6c,
	0x46, 0x69, 0x6c, 0x65, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69,
	0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69,
	0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x55, 0x0a, 0x0b, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x66,
	0x6f, 0x72, 0x6d, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x34, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6d, 0x61, 0x70, 0x73, 0x2e, 0x6d, 0x61, 0x70, 0x73, 0x70, 0x6c,
	0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x73, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x46, 0x6f, 0x72, 0x6d, 0x61,
	0x74, 0x52, 0x0a, 0x66, 0x69, 0x6c, 0x65, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x22, 0x7f, 0x0a,
	0x09, 0x47, 0x63, 0x73, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x6e,
	0x70, 0x75, 0x74, 0x5f, 0x75, 0x72, 0x69, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69,
	0x6e, 0x70, 0x75, 0x74, 0x55, 0x72, 0x69, 0x12, 0x55, 0x0a, 0x0b, 0x66, 0x69, 0x6c, 0x65, 0x5f,
	0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x34, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6d, 0x61, 0x70, 0x73, 0x2e, 0x6d, 0x61, 0x70, 0x73, 0x70,
	0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x73, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x46, 0x6f, 0x72, 0x6d,
	0x61, 0x74, 0x52, 0x0a, 0x66, 0x69, 0x6c, 0x65, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x2a, 0x98,
	0x01, 0x0a, 0x0a, 0x46, 0x69, 0x6c, 0x65, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x12, 0x1b, 0x0a,
	0x17, 0x46, 0x49, 0x4c, 0x45, 0x5f, 0x46, 0x4f, 0x52, 0x4d, 0x41, 0x54, 0x5f, 0x55, 0x4e, 0x53,
	0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x17, 0x0a, 0x13, 0x46, 0x49,
	0x4c, 0x45, 0x5f, 0x46, 0x4f, 0x52, 0x4d, 0x41, 0x54, 0x5f, 0x47, 0x45, 0x4f, 0x4a, 0x53, 0x4f,
	0x4e, 0x10, 0x01, 0x12, 0x13, 0x0a, 0x0f, 0x46, 0x49, 0x4c, 0x45, 0x5f, 0x46, 0x4f, 0x52, 0x4d,
	0x41, 0x54, 0x5f, 0x4b, 0x4d, 0x4c, 0x10, 0x02, 0x12, 0x13, 0x0a, 0x0f, 0x46, 0x49, 0x4c, 0x45,
	0x5f, 0x46, 0x4f, 0x52, 0x4d, 0x41, 0x54, 0x5f, 0x43, 0x53, 0x56, 0x10, 0x03, 0x12, 0x15, 0x0a,
	0x11, 0x46, 0x49, 0x4c, 0x45, 0x5f, 0x46, 0x4f, 0x52, 0x4d, 0x41, 0x54, 0x5f, 0x50, 0x52, 0x4f,
	0x54, 0x4f, 0x10, 0x04, 0x12, 0x13, 0x0a, 0x0f, 0x46, 0x49, 0x4c, 0x45, 0x5f, 0x46, 0x4f, 0x52,
	0x4d, 0x41, 0x54, 0x5f, 0x4b, 0x4d, 0x5a, 0x10, 0x05, 0x42, 0xf5, 0x01, 0x0a, 0x2c, 0x63, 0x6f,
	0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6d, 0x61, 0x70, 0x73, 0x2e, 0x6d, 0x61,
	0x70, 0x73, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65,
	0x74, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x42, 0x0f, 0x44, 0x61, 0x74, 0x61,
	0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x5c, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67,
	0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x6d, 0x61, 0x70, 0x73, 0x2f, 0x6d, 0x61, 0x70, 0x73, 0x70, 0x6c,
	0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x73, 0x2f, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x3b, 0x6d, 0x61, 0x70, 0x73, 0x70, 0x6c, 0x61, 0x74, 0x66,
	0x6f, 0x72, 0x6d, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x73, 0xaa, 0x02, 0x28, 0x47, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x4d, 0x61, 0x70, 0x73, 0x2e, 0x4d, 0x61, 0x70, 0x73, 0x50, 0x6c,
	0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x73, 0x2e, 0x56,
	0x31, 0x41, 0x6c, 0x70, 0x68, 0x61, 0xca, 0x02, 0x28, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c,
	0x4d, 0x61, 0x70, 0x73, 0x5c, 0x4d, 0x61, 0x70, 0x73, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72,
	0x6d, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x73, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_maps_mapsplatformdatasets_v1alpha_data_source_proto_rawDescOnce sync.Once
	file_google_maps_mapsplatformdatasets_v1alpha_data_source_proto_rawDescData = file_google_maps_mapsplatformdatasets_v1alpha_data_source_proto_rawDesc
)

func file_google_maps_mapsplatformdatasets_v1alpha_data_source_proto_rawDescGZIP() []byte {
	file_google_maps_mapsplatformdatasets_v1alpha_data_source_proto_rawDescOnce.Do(func() {
		file_google_maps_mapsplatformdatasets_v1alpha_data_source_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_maps_mapsplatformdatasets_v1alpha_data_source_proto_rawDescData)
	})
	return file_google_maps_mapsplatformdatasets_v1alpha_data_source_proto_rawDescData
}

var file_google_maps_mapsplatformdatasets_v1alpha_data_source_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_maps_mapsplatformdatasets_v1alpha_data_source_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_google_maps_mapsplatformdatasets_v1alpha_data_source_proto_goTypes = []interface{}{
	(FileFormat)(0),         // 0: google.maps.mapsplatformdatasets.v1alpha.FileFormat
	(*LocalFileSource)(nil), // 1: google.maps.mapsplatformdatasets.v1alpha.LocalFileSource
	(*GcsSource)(nil),       // 2: google.maps.mapsplatformdatasets.v1alpha.GcsSource
}
var file_google_maps_mapsplatformdatasets_v1alpha_data_source_proto_depIdxs = []int32{
	0, // 0: google.maps.mapsplatformdatasets.v1alpha.LocalFileSource.file_format:type_name -> google.maps.mapsplatformdatasets.v1alpha.FileFormat
	0, // 1: google.maps.mapsplatformdatasets.v1alpha.GcsSource.file_format:type_name -> google.maps.mapsplatformdatasets.v1alpha.FileFormat
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_google_maps_mapsplatformdatasets_v1alpha_data_source_proto_init() }
func file_google_maps_mapsplatformdatasets_v1alpha_data_source_proto_init() {
	if File_google_maps_mapsplatformdatasets_v1alpha_data_source_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_maps_mapsplatformdatasets_v1alpha_data_source_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LocalFileSource); i {
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
		file_google_maps_mapsplatformdatasets_v1alpha_data_source_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GcsSource); i {
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
			RawDescriptor: file_google_maps_mapsplatformdatasets_v1alpha_data_source_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_maps_mapsplatformdatasets_v1alpha_data_source_proto_goTypes,
		DependencyIndexes: file_google_maps_mapsplatformdatasets_v1alpha_data_source_proto_depIdxs,
		EnumInfos:         file_google_maps_mapsplatformdatasets_v1alpha_data_source_proto_enumTypes,
		MessageInfos:      file_google_maps_mapsplatformdatasets_v1alpha_data_source_proto_msgTypes,
	}.Build()
	File_google_maps_mapsplatformdatasets_v1alpha_data_source_proto = out.File
	file_google_maps_mapsplatformdatasets_v1alpha_data_source_proto_rawDesc = nil
	file_google_maps_mapsplatformdatasets_v1alpha_data_source_proto_goTypes = nil
	file_google_maps_mapsplatformdatasets_v1alpha_data_source_proto_depIdxs = nil
}
