// Code generated by goctl. DO NOT EDIT.
// Source: cmdb.proto

package model

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

	Model interface {
		// 创建资源模型
		CreateResourceModel(ctx context.Context, in *CreateModelRequest, opts ...grpc.CallOption) (*CreateModelResponse, error)
	}

	defaultModel struct {
		cli zrpc.Client
	}
)

func NewModel(cli zrpc.Client) Model {
	return &defaultModel{
		cli: cli,
	}
}

// 创建资源模型
func (m *defaultModel) CreateResourceModel(ctx context.Context, in *CreateModelRequest, opts ...grpc.CallOption) (*CreateModelResponse, error) {
	client := pb.NewModelClient(m.cli.Conn())
	return client.CreateResourceModel(ctx, in, opts...)
}