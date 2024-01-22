// Code generated by goctl. DO NOT EDIT.
// Source: cmdb.proto

package resource

import (
	"context"

	"github.com/admgo/admgo/services/cmdb/rpc/pb/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	CreateAttributeRequest      = pb.CreateAttributeRequest
	CreateAttributeResponse     = pb.CreateAttributeResponse
	CreateAttributeTypeRequest  = pb.CreateAttributeTypeRequest
	CreateAttributeTypeResponse = pb.CreateAttributeTypeResponse
	CreateCategoryRequest       = pb.CreateCategoryRequest
	CreateCategoryResponse      = pb.CreateCategoryResponse
	CreateModelRequest          = pb.CreateModelRequest
	CreateModelResponse         = pb.CreateModelResponse
	CreateResourceGroupRequest  = pb.CreateResourceGroupRequest
	CreateResourceGroupResponse = pb.CreateResourceGroupResponse
	CreateResourceRequest       = pb.CreateResourceRequest
	CreateResourceResponse      = pb.CreateResourceResponse

	Resource interface {
		// 创建资源
		CreateResource(ctx context.Context, in *CreateResourceRequest, opts ...grpc.CallOption) (*CreateResourceResponse, error)
	}

	defaultResource struct {
		cli zrpc.Client
	}
)

func NewResource(cli zrpc.Client) Resource {
	return &defaultResource{
		cli: cli,
	}
}

// 创建资源
func (m *defaultResource) CreateResource(ctx context.Context, in *CreateResourceRequest, opts ...grpc.CallOption) (*CreateResourceResponse, error) {
	client := pb.NewResourceClient(m.cli.Conn())
	return client.CreateResource(ctx, in, opts...)
}
