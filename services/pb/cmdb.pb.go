// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.19.4
// source: cmdb.proto

package pb

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

// ==============================================ResourceGroup(资源组)===================================================
type CreateResourceGroupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateResourceGroupRequest) Reset() {
	*x = CreateResourceGroupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmdb_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateResourceGroupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateResourceGroupRequest) ProtoMessage() {}

func (x *CreateResourceGroupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cmdb_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateResourceGroupRequest.ProtoReflect.Descriptor instead.
func (*CreateResourceGroupRequest) Descriptor() ([]byte, []int) {
	return file_cmdb_proto_rawDescGZIP(), []int{0}
}

type CreateResourceGroupResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateResourceGroupResponse) Reset() {
	*x = CreateResourceGroupResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmdb_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateResourceGroupResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateResourceGroupResponse) ProtoMessage() {}

func (x *CreateResourceGroupResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cmdb_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateResourceGroupResponse.ProtoReflect.Descriptor instead.
func (*CreateResourceGroupResponse) Descriptor() ([]byte, []int) {
	return file_cmdb_proto_rawDescGZIP(), []int{1}
}

// ==============================================ResourceType(资源类型)==================================================
type CreateResourceTypeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateResourceTypeRequest) Reset() {
	*x = CreateResourceTypeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmdb_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateResourceTypeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateResourceTypeRequest) ProtoMessage() {}

func (x *CreateResourceTypeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cmdb_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateResourceTypeRequest.ProtoReflect.Descriptor instead.
func (*CreateResourceTypeRequest) Descriptor() ([]byte, []int) {
	return file_cmdb_proto_rawDescGZIP(), []int{2}
}

type CreateResourceTypeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateResourceTypeResponse) Reset() {
	*x = CreateResourceTypeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmdb_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateResourceTypeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateResourceTypeResponse) ProtoMessage() {}

func (x *CreateResourceTypeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cmdb_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateResourceTypeResponse.ProtoReflect.Descriptor instead.
func (*CreateResourceTypeResponse) Descriptor() ([]byte, []int) {
	return file_cmdb_proto_rawDescGZIP(), []int{3}
}

// ==============================================ResourceModel(资源模型)=================================================
type CreateResourceModelRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateResourceModelRequest) Reset() {
	*x = CreateResourceModelRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmdb_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateResourceModelRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateResourceModelRequest) ProtoMessage() {}

func (x *CreateResourceModelRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cmdb_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateResourceModelRequest.ProtoReflect.Descriptor instead.
func (*CreateResourceModelRequest) Descriptor() ([]byte, []int) {
	return file_cmdb_proto_rawDescGZIP(), []int{4}
}

type CreateResourceModelResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateResourceModelResponse) Reset() {
	*x = CreateResourceModelResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmdb_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateResourceModelResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateResourceModelResponse) ProtoMessage() {}

func (x *CreateResourceModelResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cmdb_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateResourceModelResponse.ProtoReflect.Descriptor instead.
func (*CreateResourceModelResponse) Descriptor() ([]byte, []int) {
	return file_cmdb_proto_rawDescGZIP(), []int{5}
}

// ==============================================ResourceAttribute(资源属性)=============================================
type CreateResourceAttributeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateResourceAttributeRequest) Reset() {
	*x = CreateResourceAttributeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmdb_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateResourceAttributeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateResourceAttributeRequest) ProtoMessage() {}

func (x *CreateResourceAttributeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cmdb_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateResourceAttributeRequest.ProtoReflect.Descriptor instead.
func (*CreateResourceAttributeRequest) Descriptor() ([]byte, []int) {
	return file_cmdb_proto_rawDescGZIP(), []int{6}
}

type CreateResourceAttributeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateResourceAttributeResponse) Reset() {
	*x = CreateResourceAttributeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmdb_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateResourceAttributeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateResourceAttributeResponse) ProtoMessage() {}

func (x *CreateResourceAttributeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cmdb_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateResourceAttributeResponse.ProtoReflect.Descriptor instead.
func (*CreateResourceAttributeResponse) Descriptor() ([]byte, []int) {
	return file_cmdb_proto_rawDescGZIP(), []int{7}
}

// ==============================================AttributeType(资源属性类型)==============================================
type CreateAttributeTypeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateAttributeTypeRequest) Reset() {
	*x = CreateAttributeTypeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmdb_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAttributeTypeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAttributeTypeRequest) ProtoMessage() {}

func (x *CreateAttributeTypeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cmdb_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAttributeTypeRequest.ProtoReflect.Descriptor instead.
func (*CreateAttributeTypeRequest) Descriptor() ([]byte, []int) {
	return file_cmdb_proto_rawDescGZIP(), []int{8}
}

type CreateAttributeTypeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateAttributeTypeResponse) Reset() {
	*x = CreateAttributeTypeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmdb_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAttributeTypeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAttributeTypeResponse) ProtoMessage() {}

func (x *CreateAttributeTypeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cmdb_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAttributeTypeResponse.ProtoReflect.Descriptor instead.
func (*CreateAttributeTypeResponse) Descriptor() ([]byte, []int) {
	return file_cmdb_proto_rawDescGZIP(), []int{9}
}

// ==============================================Resource(资源)==============================================
type CreateResourceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateResourceRequest) Reset() {
	*x = CreateResourceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmdb_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateResourceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateResourceRequest) ProtoMessage() {}

func (x *CreateResourceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cmdb_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateResourceRequest.ProtoReflect.Descriptor instead.
func (*CreateResourceRequest) Descriptor() ([]byte, []int) {
	return file_cmdb_proto_rawDescGZIP(), []int{10}
}

type CreateResourceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateResourceResponse) Reset() {
	*x = CreateResourceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmdb_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateResourceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateResourceResponse) ProtoMessage() {}

func (x *CreateResourceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cmdb_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateResourceResponse.ProtoReflect.Descriptor instead.
func (*CreateResourceResponse) Descriptor() ([]byte, []int) {
	return file_cmdb_proto_rawDescGZIP(), []int{11}
}

var File_cmdb_proto protoreflect.FileDescriptor

var file_cmdb_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x63, 0x6d, 0x64, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0x1c, 0x0a, 0x1a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x22, 0x1d, 0x0a, 0x1b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x1b, 0x0a, 0x19, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22,
	0x1c, 0x0a, 0x1a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1c, 0x0a,
	0x1a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4d,
	0x6f, 0x64, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x1d, 0x0a, 0x1b, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4d, 0x6f, 0x64,
	0x65, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x20, 0x0a, 0x1e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x41, 0x74, 0x74, 0x72,
	0x69, 0x62, 0x75, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x21, 0x0a, 0x1f,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x41, 0x74,
	0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x1c, 0x0a, 0x1a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75,
	0x74, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x1d, 0x0a,
	0x1b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x17, 0x0a, 0x15,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x18, 0x0a, 0x16, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32,
	0x71, 0x0a, 0x0d, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x12, 0x60, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x23, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x32, 0x6d, 0x0a, 0x0c, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x5d, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x22, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x32, 0x71, 0x0a, 0x0d, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4d, 0x6f, 0x64,
	0x65, 0x6c, 0x12, 0x60, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x23, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x32, 0x81, 0x01, 0x0a, 0x11, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x12, 0x6c, 0x0a, 0x17, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x41, 0x74, 0x74, 0x72,
	0x69, 0x62, 0x75, 0x74, 0x65, 0x12, 0x27, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x41, 0x74,
	0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x71, 0x0a, 0x0d, 0x41, 0x74, 0x74, 0x72,
	0x69, 0x62, 0x75, 0x74, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x60, 0x0a, 0x13, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x23, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x5d, 0x0a, 0x08, 0x52,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x51, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x1e, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x07, 0x5a, 0x05, 0x2e, 0x2e,
	0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cmdb_proto_rawDescOnce sync.Once
	file_cmdb_proto_rawDescData = file_cmdb_proto_rawDesc
)

func file_cmdb_proto_rawDescGZIP() []byte {
	file_cmdb_proto_rawDescOnce.Do(func() {
		file_cmdb_proto_rawDescData = protoimpl.X.CompressGZIP(file_cmdb_proto_rawDescData)
	})
	return file_cmdb_proto_rawDescData
}

var file_cmdb_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_cmdb_proto_goTypes = []interface{}{
	(*CreateResourceGroupRequest)(nil),      // 0: service.CreateResourceGroupRequest
	(*CreateResourceGroupResponse)(nil),     // 1: service.CreateResourceGroupResponse
	(*CreateResourceTypeRequest)(nil),       // 2: service.CreateResourceTypeRequest
	(*CreateResourceTypeResponse)(nil),      // 3: service.CreateResourceTypeResponse
	(*CreateResourceModelRequest)(nil),      // 4: service.CreateResourceModelRequest
	(*CreateResourceModelResponse)(nil),     // 5: service.CreateResourceModelResponse
	(*CreateResourceAttributeRequest)(nil),  // 6: service.CreateResourceAttributeRequest
	(*CreateResourceAttributeResponse)(nil), // 7: service.CreateResourceAttributeResponse
	(*CreateAttributeTypeRequest)(nil),      // 8: service.CreateAttributeTypeRequest
	(*CreateAttributeTypeResponse)(nil),     // 9: service.CreateAttributeTypeResponse
	(*CreateResourceRequest)(nil),           // 10: service.CreateResourceRequest
	(*CreateResourceResponse)(nil),          // 11: service.CreateResourceResponse
}
var file_cmdb_proto_depIdxs = []int32{
	0,  // 0: service.ResourceGroup.CreateResourceGroup:input_type -> service.CreateResourceGroupRequest
	2,  // 1: service.ResourceType.CreateResourceType:input_type -> service.CreateResourceTypeRequest
	4,  // 2: service.ResourceModel.CreateResourceModel:input_type -> service.CreateResourceModelRequest
	6,  // 3: service.ResourceAttribute.CreateResourceAttribute:input_type -> service.CreateResourceAttributeRequest
	8,  // 4: service.AttributeType.CreateAttributeType:input_type -> service.CreateAttributeTypeRequest
	10, // 5: service.Resource.CreateResource:input_type -> service.CreateResourceRequest
	1,  // 6: service.ResourceGroup.CreateResourceGroup:output_type -> service.CreateResourceGroupResponse
	3,  // 7: service.ResourceType.CreateResourceType:output_type -> service.CreateResourceTypeResponse
	5,  // 8: service.ResourceModel.CreateResourceModel:output_type -> service.CreateResourceModelResponse
	7,  // 9: service.ResourceAttribute.CreateResourceAttribute:output_type -> service.CreateResourceAttributeResponse
	9,  // 10: service.AttributeType.CreateAttributeType:output_type -> service.CreateAttributeTypeResponse
	11, // 11: service.Resource.CreateResource:output_type -> service.CreateResourceResponse
	6,  // [6:12] is the sub-list for method output_type
	0,  // [0:6] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_cmdb_proto_init() }
func file_cmdb_proto_init() {
	if File_cmdb_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cmdb_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateResourceGroupRequest); i {
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
		file_cmdb_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateResourceGroupResponse); i {
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
		file_cmdb_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateResourceTypeRequest); i {
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
		file_cmdb_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateResourceTypeResponse); i {
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
		file_cmdb_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateResourceModelRequest); i {
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
		file_cmdb_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateResourceModelResponse); i {
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
		file_cmdb_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateResourceAttributeRequest); i {
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
		file_cmdb_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateResourceAttributeResponse); i {
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
		file_cmdb_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateAttributeTypeRequest); i {
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
		file_cmdb_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateAttributeTypeResponse); i {
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
		file_cmdb_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateResourceRequest); i {
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
		file_cmdb_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateResourceResponse); i {
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
			RawDescriptor: file_cmdb_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   6,
		},
		GoTypes:           file_cmdb_proto_goTypes,
		DependencyIndexes: file_cmdb_proto_depIdxs,
		MessageInfos:      file_cmdb_proto_msgTypes,
	}.Build()
	File_cmdb_proto = out.File
	file_cmdb_proto_rawDesc = nil
	file_cmdb_proto_goTypes = nil
	file_cmdb_proto_depIdxs = nil
}
