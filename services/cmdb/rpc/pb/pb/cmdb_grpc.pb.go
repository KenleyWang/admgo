// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.19.4
// source: cmdb.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	ResourceGroup_CreateResourceGroup_FullMethodName = "/service.ResourceGroup/CreateResourceGroup"
)

// ResourceGroupClient is the client API for ResourceGroup service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ResourceGroupClient interface {
	// 创建资源组
	CreateResourceGroup(ctx context.Context, in *CreateResourceGroupRequest, opts ...grpc.CallOption) (*CreateResourceGroupResponse, error)
}

type resourceGroupClient struct {
	cc grpc.ClientConnInterface
}

func NewResourceGroupClient(cc grpc.ClientConnInterface) ResourceGroupClient {
	return &resourceGroupClient{cc}
}

func (c *resourceGroupClient) CreateResourceGroup(ctx context.Context, in *CreateResourceGroupRequest, opts ...grpc.CallOption) (*CreateResourceGroupResponse, error) {
	out := new(CreateResourceGroupResponse)
	err := c.cc.Invoke(ctx, ResourceGroup_CreateResourceGroup_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ResourceGroupServer is the server API for ResourceGroup service.
// All implementations must embed UnimplementedResourceGroupServer
// for forward compatibility
type ResourceGroupServer interface {
	// 创建资源组
	CreateResourceGroup(context.Context, *CreateResourceGroupRequest) (*CreateResourceGroupResponse, error)
	mustEmbedUnimplementedResourceGroupServer()
}

// UnimplementedResourceGroupServer must be embedded to have forward compatible implementations.
type UnimplementedResourceGroupServer struct {
}

func (UnimplementedResourceGroupServer) CreateResourceGroup(context.Context, *CreateResourceGroupRequest) (*CreateResourceGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateResourceGroup not implemented")
}
func (UnimplementedResourceGroupServer) mustEmbedUnimplementedResourceGroupServer() {}

// UnsafeResourceGroupServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ResourceGroupServer will
// result in compilation errors.
type UnsafeResourceGroupServer interface {
	mustEmbedUnimplementedResourceGroupServer()
}

func RegisterResourceGroupServer(s grpc.ServiceRegistrar, srv ResourceGroupServer) {
	s.RegisterService(&ResourceGroup_ServiceDesc, srv)
}

func _ResourceGroup_CreateResourceGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateResourceGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceGroupServer).CreateResourceGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ResourceGroup_CreateResourceGroup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceGroupServer).CreateResourceGroup(ctx, req.(*CreateResourceGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ResourceGroup_ServiceDesc is the grpc.ServiceDesc for ResourceGroup service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ResourceGroup_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "service.ResourceGroup",
	HandlerType: (*ResourceGroupServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateResourceGroup",
			Handler:    _ResourceGroup_CreateResourceGroup_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cmdb.proto",
}

const (
	Category_CreateResourceType_FullMethodName = "/service.Category/CreateResourceType"
)

// CategoryClient is the client API for Category service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CategoryClient interface {
	// 创建资源类型
	CreateResourceType(ctx context.Context, in *CreateCategoryRequest, opts ...grpc.CallOption) (*CreateCategoryResponse, error)
}

type categoryClient struct {
	cc grpc.ClientConnInterface
}

func NewCategoryClient(cc grpc.ClientConnInterface) CategoryClient {
	return &categoryClient{cc}
}

func (c *categoryClient) CreateResourceType(ctx context.Context, in *CreateCategoryRequest, opts ...grpc.CallOption) (*CreateCategoryResponse, error) {
	out := new(CreateCategoryResponse)
	err := c.cc.Invoke(ctx, Category_CreateResourceType_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CategoryServer is the server API for Category service.
// All implementations must embed UnimplementedCategoryServer
// for forward compatibility
type CategoryServer interface {
	// 创建资源类型
	CreateResourceType(context.Context, *CreateCategoryRequest) (*CreateCategoryResponse, error)
	mustEmbedUnimplementedCategoryServer()
}

// UnimplementedCategoryServer must be embedded to have forward compatible implementations.
type UnimplementedCategoryServer struct {
}

func (UnimplementedCategoryServer) CreateResourceType(context.Context, *CreateCategoryRequest) (*CreateCategoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateResourceType not implemented")
}
func (UnimplementedCategoryServer) mustEmbedUnimplementedCategoryServer() {}

// UnsafeCategoryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CategoryServer will
// result in compilation errors.
type UnsafeCategoryServer interface {
	mustEmbedUnimplementedCategoryServer()
}

func RegisterCategoryServer(s grpc.ServiceRegistrar, srv CategoryServer) {
	s.RegisterService(&Category_ServiceDesc, srv)
}

func _Category_CreateResourceType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCategoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CategoryServer).CreateResourceType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Category_CreateResourceType_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CategoryServer).CreateResourceType(ctx, req.(*CreateCategoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Category_ServiceDesc is the grpc.ServiceDesc for Category service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Category_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "service.Category",
	HandlerType: (*CategoryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateResourceType",
			Handler:    _Category_CreateResourceType_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cmdb.proto",
}

const (
	Model_CreateResourceModel_FullMethodName = "/service.Model/CreateResourceModel"
)

// ModelClient is the client API for Model service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ModelClient interface {
	// 创建资源模型
	CreateResourceModel(ctx context.Context, in *CreateModelRequest, opts ...grpc.CallOption) (*CreateModelResponse, error)
}

type modelClient struct {
	cc grpc.ClientConnInterface
}

func NewModelClient(cc grpc.ClientConnInterface) ModelClient {
	return &modelClient{cc}
}

func (c *modelClient) CreateResourceModel(ctx context.Context, in *CreateModelRequest, opts ...grpc.CallOption) (*CreateModelResponse, error) {
	out := new(CreateModelResponse)
	err := c.cc.Invoke(ctx, Model_CreateResourceModel_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ModelServer is the server API for Model service.
// All implementations must embed UnimplementedModelServer
// for forward compatibility
type ModelServer interface {
	// 创建资源模型
	CreateResourceModel(context.Context, *CreateModelRequest) (*CreateModelResponse, error)
	mustEmbedUnimplementedModelServer()
}

// UnimplementedModelServer must be embedded to have forward compatible implementations.
type UnimplementedModelServer struct {
}

func (UnimplementedModelServer) CreateResourceModel(context.Context, *CreateModelRequest) (*CreateModelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateResourceModel not implemented")
}
func (UnimplementedModelServer) mustEmbedUnimplementedModelServer() {}

// UnsafeModelServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ModelServer will
// result in compilation errors.
type UnsafeModelServer interface {
	mustEmbedUnimplementedModelServer()
}

func RegisterModelServer(s grpc.ServiceRegistrar, srv ModelServer) {
	s.RegisterService(&Model_ServiceDesc, srv)
}

func _Model_CreateResourceModel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateModelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModelServer).CreateResourceModel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Model_CreateResourceModel_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModelServer).CreateResourceModel(ctx, req.(*CreateModelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Model_ServiceDesc is the grpc.ServiceDesc for Model service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Model_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "service.Model",
	HandlerType: (*ModelServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateResourceModel",
			Handler:    _Model_CreateResourceModel_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cmdb.proto",
}

const (
	Attribute_CreateResourceAttribute_FullMethodName = "/service.Attribute/CreateResourceAttribute"
)

// AttributeClient is the client API for Attribute service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AttributeClient interface {
	// 创建资源属性
	CreateResourceAttribute(ctx context.Context, in *CreateAttributeRequest, opts ...grpc.CallOption) (*CreateAttributeResponse, error)
}

type attributeClient struct {
	cc grpc.ClientConnInterface
}

func NewAttributeClient(cc grpc.ClientConnInterface) AttributeClient {
	return &attributeClient{cc}
}

func (c *attributeClient) CreateResourceAttribute(ctx context.Context, in *CreateAttributeRequest, opts ...grpc.CallOption) (*CreateAttributeResponse, error) {
	out := new(CreateAttributeResponse)
	err := c.cc.Invoke(ctx, Attribute_CreateResourceAttribute_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AttributeServer is the server API for Attribute service.
// All implementations must embed UnimplementedAttributeServer
// for forward compatibility
type AttributeServer interface {
	// 创建资源属性
	CreateResourceAttribute(context.Context, *CreateAttributeRequest) (*CreateAttributeResponse, error)
	mustEmbedUnimplementedAttributeServer()
}

// UnimplementedAttributeServer must be embedded to have forward compatible implementations.
type UnimplementedAttributeServer struct {
}

func (UnimplementedAttributeServer) CreateResourceAttribute(context.Context, *CreateAttributeRequest) (*CreateAttributeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateResourceAttribute not implemented")
}
func (UnimplementedAttributeServer) mustEmbedUnimplementedAttributeServer() {}

// UnsafeAttributeServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AttributeServer will
// result in compilation errors.
type UnsafeAttributeServer interface {
	mustEmbedUnimplementedAttributeServer()
}

func RegisterAttributeServer(s grpc.ServiceRegistrar, srv AttributeServer) {
	s.RegisterService(&Attribute_ServiceDesc, srv)
}

func _Attribute_CreateResourceAttribute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAttributeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AttributeServer).CreateResourceAttribute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Attribute_CreateResourceAttribute_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AttributeServer).CreateResourceAttribute(ctx, req.(*CreateAttributeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Attribute_ServiceDesc is the grpc.ServiceDesc for Attribute service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Attribute_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "service.Attribute",
	HandlerType: (*AttributeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateResourceAttribute",
			Handler:    _Attribute_CreateResourceAttribute_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cmdb.proto",
}

const (
	AttributeType_CreateAttributeType_FullMethodName = "/service.AttributeType/CreateAttributeType"
)

// AttributeTypeClient is the client API for AttributeType service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AttributeTypeClient interface {
	// 创建资源属性类型
	CreateAttributeType(ctx context.Context, in *CreateAttributeTypeRequest, opts ...grpc.CallOption) (*CreateAttributeTypeResponse, error)
}

type attributeTypeClient struct {
	cc grpc.ClientConnInterface
}

func NewAttributeTypeClient(cc grpc.ClientConnInterface) AttributeTypeClient {
	return &attributeTypeClient{cc}
}

func (c *attributeTypeClient) CreateAttributeType(ctx context.Context, in *CreateAttributeTypeRequest, opts ...grpc.CallOption) (*CreateAttributeTypeResponse, error) {
	out := new(CreateAttributeTypeResponse)
	err := c.cc.Invoke(ctx, AttributeType_CreateAttributeType_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AttributeTypeServer is the server API for AttributeType service.
// All implementations must embed UnimplementedAttributeTypeServer
// for forward compatibility
type AttributeTypeServer interface {
	// 创建资源属性类型
	CreateAttributeType(context.Context, *CreateAttributeTypeRequest) (*CreateAttributeTypeResponse, error)
	mustEmbedUnimplementedAttributeTypeServer()
}

// UnimplementedAttributeTypeServer must be embedded to have forward compatible implementations.
type UnimplementedAttributeTypeServer struct {
}

func (UnimplementedAttributeTypeServer) CreateAttributeType(context.Context, *CreateAttributeTypeRequest) (*CreateAttributeTypeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAttributeType not implemented")
}
func (UnimplementedAttributeTypeServer) mustEmbedUnimplementedAttributeTypeServer() {}

// UnsafeAttributeTypeServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AttributeTypeServer will
// result in compilation errors.
type UnsafeAttributeTypeServer interface {
	mustEmbedUnimplementedAttributeTypeServer()
}

func RegisterAttributeTypeServer(s grpc.ServiceRegistrar, srv AttributeTypeServer) {
	s.RegisterService(&AttributeType_ServiceDesc, srv)
}

func _AttributeType_CreateAttributeType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAttributeTypeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AttributeTypeServer).CreateAttributeType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AttributeType_CreateAttributeType_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AttributeTypeServer).CreateAttributeType(ctx, req.(*CreateAttributeTypeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AttributeType_ServiceDesc is the grpc.ServiceDesc for AttributeType service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AttributeType_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "service.AttributeType",
	HandlerType: (*AttributeTypeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateAttributeType",
			Handler:    _AttributeType_CreateAttributeType_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cmdb.proto",
}

const (
	Resource_CreateResource_FullMethodName = "/service.Resource/CreateResource"
)

// ResourceClient is the client API for Resource service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ResourceClient interface {
	// 创建资源
	CreateResource(ctx context.Context, in *CreateResourceRequest, opts ...grpc.CallOption) (*CreateResourceResponse, error)
}

type resourceClient struct {
	cc grpc.ClientConnInterface
}

func NewResourceClient(cc grpc.ClientConnInterface) ResourceClient {
	return &resourceClient{cc}
}

func (c *resourceClient) CreateResource(ctx context.Context, in *CreateResourceRequest, opts ...grpc.CallOption) (*CreateResourceResponse, error) {
	out := new(CreateResourceResponse)
	err := c.cc.Invoke(ctx, Resource_CreateResource_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ResourceServer is the server API for Resource service.
// All implementations must embed UnimplementedResourceServer
// for forward compatibility
type ResourceServer interface {
	// 创建资源
	CreateResource(context.Context, *CreateResourceRequest) (*CreateResourceResponse, error)
	mustEmbedUnimplementedResourceServer()
}

// UnimplementedResourceServer must be embedded to have forward compatible implementations.
type UnimplementedResourceServer struct {
}

func (UnimplementedResourceServer) CreateResource(context.Context, *CreateResourceRequest) (*CreateResourceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateResource not implemented")
}
func (UnimplementedResourceServer) mustEmbedUnimplementedResourceServer() {}

// UnsafeResourceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ResourceServer will
// result in compilation errors.
type UnsafeResourceServer interface {
	mustEmbedUnimplementedResourceServer()
}

func RegisterResourceServer(s grpc.ServiceRegistrar, srv ResourceServer) {
	s.RegisterService(&Resource_ServiceDesc, srv)
}

func _Resource_CreateResource_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateResourceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceServer).CreateResource(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Resource_CreateResource_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceServer).CreateResource(ctx, req.(*CreateResourceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Resource_ServiceDesc is the grpc.ServiceDesc for Resource service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Resource_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "service.Resource",
	HandlerType: (*ResourceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateResource",
			Handler:    _Resource_CreateResource_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cmdb.proto",
}
