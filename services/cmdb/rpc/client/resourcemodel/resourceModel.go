// Code generated by goctl. DO NOT EDIT.
// Source: cmdb.proto

package resourcemodel

import (
	"context"

	"github.com/admgo/admgo/services/cmdb/rpc/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	CreateAttributeTypeRequest      = pb.CreateAttributeTypeRequest
	CreateAttributeTypeResponse     = pb.CreateAttributeTypeResponse
	CreateResourceAttributeRequest  = pb.CreateResourceAttributeRequest
	CreateResourceAttributeResponse = pb.CreateResourceAttributeResponse
	CreateResourceGroupRequest      = pb.CreateResourceGroupRequest
	CreateResourceGroupResponse     = pb.CreateResourceGroupResponse
	CreateResourceModelRequest      = pb.CreateResourceModelRequest
	CreateResourceModelResponse     = pb.CreateResourceModelResponse
	CreateResourceRequest           = pb.CreateResourceRequest
	CreateResourceResponse          = pb.CreateResourceResponse
	CreateResourceTypeRequest       = pb.CreateResourceTypeRequest
	CreateResourceTypeResponse      = pb.CreateResourceTypeResponse

	ResourceModel interface {
		// 创建资源模型
		CreateResourceModel(ctx context.Context, in *CreateResourceModelRequest, opts ...grpc.CallOption) (*CreateResourceModelResponse, error)
	}

	defaultResourceModel struct {
		cli zrpc.Client
	}
)

func NewResourceModel(cli zrpc.Client) ResourceModel {
	return &defaultResourceModel{
		cli: cli,
	}
}

// 创建资源模型
func (m *defaultResourceModel) CreateResourceModel(ctx context.Context, in *CreateResourceModelRequest, opts ...grpc.CallOption) (*CreateResourceModelResponse, error) {
	client := pb.NewResourceModelClient(m.cli.Conn())
	return client.CreateResourceModel(ctx, in, opts...)
}
