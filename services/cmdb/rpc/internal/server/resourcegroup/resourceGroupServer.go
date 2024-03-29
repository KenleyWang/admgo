// Code generated by goctl. DO NOT EDIT.
// Source: cmdb.proto

package server

import (
	"context"

	"github.com/admgo/admgo/services/cmdb/rpc/internal/logic/resourcegroup"
	"github.com/admgo/admgo/services/cmdb/rpc/internal/svc"
	"github.com/admgo/admgo/services/cmdb/rpc/pb/pb"
)

type ResourceGroupServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedResourceGroupServer
}

func NewResourceGroupServer(svcCtx *svc.ServiceContext) *ResourceGroupServer {
	return &ResourceGroupServer{
		svcCtx: svcCtx,
	}
}

// 创建资源组
func (s *ResourceGroupServer) CreateResourceGroup(ctx context.Context, in *pb.CreateResourceGroupRequest) (*pb.CreateResourceGroupResponse, error) {
	l := resourcegrouplogic.NewCreateResourceGroupLogic(ctx, s.svcCtx)
	return l.CreateResourceGroup(in)
}
