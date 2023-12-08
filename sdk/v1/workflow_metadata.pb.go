// The MIT License
//
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.20.3
// source: temporal/api/sdk/v1/workflow_metadata.proto

package sdk

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

// The name of the query to retrieve this information is `__temporal_getWorkflowMetadata`.
type WorkflowMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Metadata provided at declaration or creation time.
	Definition *WorkflowDefinition `protobuf:"bytes,1,opt,name=definition,proto3" json:"definition,omitempty"`
}

func (x *WorkflowMetadata) Reset() {
	*x = WorkflowMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_temporal_api_sdk_v1_workflow_metadata_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorkflowMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkflowMetadata) ProtoMessage() {}

func (x *WorkflowMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_temporal_api_sdk_v1_workflow_metadata_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkflowMetadata.ProtoReflect.Descriptor instead.
func (*WorkflowMetadata) Descriptor() ([]byte, []int) {
	return file_temporal_api_sdk_v1_workflow_metadata_proto_rawDescGZIP(), []int{0}
}

func (x *WorkflowMetadata) GetDefinition() *WorkflowDefinition {
	if x != nil {
		return x.Definition
	}
	return nil
}

// (-- api-linter: core::0203::optional=disabled --)
type WorkflowDefinition struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A name scoped by the task queue that maps to this workflow definition.
	// If missing, this workflow is a dynamic workflow.
	Type string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	// An optional workflow description provided by the application.
	// By convention, external tools may interpret its first part,
	// i.e., ending with a line break, as a summary of the description.
	Description       string                           `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	QueryDefinitions  []*WorkflowInteractionDefinition `protobuf:"bytes,3,rep,name=query_definitions,json=queryDefinitions,proto3" json:"query_definitions,omitempty"`
	SignalDefinitions []*WorkflowInteractionDefinition `protobuf:"bytes,4,rep,name=signal_definitions,json=signalDefinitions,proto3" json:"signal_definitions,omitempty"`
	UpdateDefinitions []*WorkflowInteractionDefinition `protobuf:"bytes,5,rep,name=update_definitions,json=updateDefinitions,proto3" json:"update_definitions,omitempty"`
}

func (x *WorkflowDefinition) Reset() {
	*x = WorkflowDefinition{}
	if protoimpl.UnsafeEnabled {
		mi := &file_temporal_api_sdk_v1_workflow_metadata_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorkflowDefinition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkflowDefinition) ProtoMessage() {}

func (x *WorkflowDefinition) ProtoReflect() protoreflect.Message {
	mi := &file_temporal_api_sdk_v1_workflow_metadata_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkflowDefinition.ProtoReflect.Descriptor instead.
func (*WorkflowDefinition) Descriptor() ([]byte, []int) {
	return file_temporal_api_sdk_v1_workflow_metadata_proto_rawDescGZIP(), []int{1}
}

func (x *WorkflowDefinition) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *WorkflowDefinition) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *WorkflowDefinition) GetQueryDefinitions() []*WorkflowInteractionDefinition {
	if x != nil {
		return x.QueryDefinitions
	}
	return nil
}

func (x *WorkflowDefinition) GetSignalDefinitions() []*WorkflowInteractionDefinition {
	if x != nil {
		return x.SignalDefinitions
	}
	return nil
}

func (x *WorkflowDefinition) GetUpdateDefinitions() []*WorkflowInteractionDefinition {
	if x != nil {
		return x.UpdateDefinitions
	}
	return nil
}

// (-- api-linter: core::0123::resource-annotation=disabled
//
//	aip.dev/not-precedent: The `name` field is optional. --)
//
// (-- api-linter: core::0203::optional=disabled --)
type WorkflowInteractionDefinition struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// An optional name for the handler. If missing, it represents
	// a dynamic handler that processes any interactions not handled by others.
	// There is at most one dynamic handler per workflow and interaction kind.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// An optional interaction description provided by the application.
	// By convention, external tools may interpret its first part,
	// i.e., ending with a line break, as a summary of the description.
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *WorkflowInteractionDefinition) Reset() {
	*x = WorkflowInteractionDefinition{}
	if protoimpl.UnsafeEnabled {
		mi := &file_temporal_api_sdk_v1_workflow_metadata_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorkflowInteractionDefinition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkflowInteractionDefinition) ProtoMessage() {}

func (x *WorkflowInteractionDefinition) ProtoReflect() protoreflect.Message {
	mi := &file_temporal_api_sdk_v1_workflow_metadata_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkflowInteractionDefinition.ProtoReflect.Descriptor instead.
func (*WorkflowInteractionDefinition) Descriptor() ([]byte, []int) {
	return file_temporal_api_sdk_v1_workflow_metadata_proto_rawDescGZIP(), []int{2}
}

func (x *WorkflowInteractionDefinition) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *WorkflowInteractionDefinition) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

var File_temporal_api_sdk_v1_workflow_metadata_proto protoreflect.FileDescriptor

var file_temporal_api_sdk_v1_workflow_metadata_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73,
	0x64, 0x6b, 0x2f, 0x76, 0x31, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x5f, 0x6d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x74,
	0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x64, 0x6b, 0x2e,
	0x76, 0x31, 0x22, 0x5b, 0x0a, 0x10, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x47, 0x0a, 0x0a, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x74, 0x65, 0x6d,
	0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x76, 0x31,
	0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x22,
	0xf1, 0x02, 0x0a, 0x12, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x44, 0x65, 0x66, 0x69,
	0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x5f, 0x0a, 0x11,
	0x71, 0x75, 0x65, 0x72, 0x79, 0x5f, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72,
	0x61, 0x6c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x57, 0x6f,
	0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x10, 0x71, 0x75, 0x65,
	0x72, 0x79, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x61, 0x0a,
	0x12, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x5f, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x74, 0x65, 0x6d, 0x70,
	0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x76, 0x31, 0x2e,
	0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x11, 0x73,
	0x69, 0x67, 0x6e, 0x61, 0x6c, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x12, 0x61, 0x0a, 0x12, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x64, 0x65, 0x66, 0x69, 0x6e,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x74,
	0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x64, 0x6b, 0x2e,
	0x76, 0x31, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x49, 0x6e, 0x74, 0x65, 0x72,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x11, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x22, 0x55, 0x0a, 0x1d, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x49,
	0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x83, 0x01, 0x0a, 0x16, 0x69,
	0x6f, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73,
	0x64, 0x6b, 0x2e, 0x76, 0x31, 0x42, 0x15, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x1d,
	0x67, 0x6f, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x69, 0x6f, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x73, 0x64, 0x6b, 0x2f, 0x76, 0x31, 0x3b, 0x73, 0x64, 0x6b, 0xaa, 0x02, 0x15,
	0x54, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x69, 0x6f, 0x2e, 0x41, 0x70, 0x69, 0x2e, 0x53,
	0x64, 0x6b, 0x2e, 0x56, 0x31, 0xea, 0x02, 0x18, 0x54, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c,
	0x69, 0x6f, 0x3a, 0x3a, 0x41, 0x70, 0x69, 0x3a, 0x3a, 0x53, 0x64, 0x6b, 0x3a, 0x3a, 0x56, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_temporal_api_sdk_v1_workflow_metadata_proto_rawDescOnce sync.Once
	file_temporal_api_sdk_v1_workflow_metadata_proto_rawDescData = file_temporal_api_sdk_v1_workflow_metadata_proto_rawDesc
)

func file_temporal_api_sdk_v1_workflow_metadata_proto_rawDescGZIP() []byte {
	file_temporal_api_sdk_v1_workflow_metadata_proto_rawDescOnce.Do(func() {
		file_temporal_api_sdk_v1_workflow_metadata_proto_rawDescData = protoimpl.X.CompressGZIP(file_temporal_api_sdk_v1_workflow_metadata_proto_rawDescData)
	})
	return file_temporal_api_sdk_v1_workflow_metadata_proto_rawDescData
}

var file_temporal_api_sdk_v1_workflow_metadata_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_temporal_api_sdk_v1_workflow_metadata_proto_goTypes = []interface{}{
	(*WorkflowMetadata)(nil),              // 0: temporal.api.sdk.v1.WorkflowMetadata
	(*WorkflowDefinition)(nil),            // 1: temporal.api.sdk.v1.WorkflowDefinition
	(*WorkflowInteractionDefinition)(nil), // 2: temporal.api.sdk.v1.WorkflowInteractionDefinition
}
var file_temporal_api_sdk_v1_workflow_metadata_proto_depIdxs = []int32{
	1, // 0: temporal.api.sdk.v1.WorkflowMetadata.definition:type_name -> temporal.api.sdk.v1.WorkflowDefinition
	2, // 1: temporal.api.sdk.v1.WorkflowDefinition.query_definitions:type_name -> temporal.api.sdk.v1.WorkflowInteractionDefinition
	2, // 2: temporal.api.sdk.v1.WorkflowDefinition.signal_definitions:type_name -> temporal.api.sdk.v1.WorkflowInteractionDefinition
	2, // 3: temporal.api.sdk.v1.WorkflowDefinition.update_definitions:type_name -> temporal.api.sdk.v1.WorkflowInteractionDefinition
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_temporal_api_sdk_v1_workflow_metadata_proto_init() }
func file_temporal_api_sdk_v1_workflow_metadata_proto_init() {
	if File_temporal_api_sdk_v1_workflow_metadata_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_temporal_api_sdk_v1_workflow_metadata_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorkflowMetadata); i {
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
		file_temporal_api_sdk_v1_workflow_metadata_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorkflowDefinition); i {
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
		file_temporal_api_sdk_v1_workflow_metadata_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorkflowInteractionDefinition); i {
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
			RawDescriptor: file_temporal_api_sdk_v1_workflow_metadata_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_temporal_api_sdk_v1_workflow_metadata_proto_goTypes,
		DependencyIndexes: file_temporal_api_sdk_v1_workflow_metadata_proto_depIdxs,
		MessageInfos:      file_temporal_api_sdk_v1_workflow_metadata_proto_msgTypes,
	}.Build()
	File_temporal_api_sdk_v1_workflow_metadata_proto = out.File
	file_temporal_api_sdk_v1_workflow_metadata_proto_rawDesc = nil
	file_temporal_api_sdk_v1_workflow_metadata_proto_goTypes = nil
	file_temporal_api_sdk_v1_workflow_metadata_proto_depIdxs = nil
}
