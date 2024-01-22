// Code generated by goctl. DO NOT EDIT.
// Source: cmdb.proto

package server

import (
	"context"

	"github.com/admgo/admgo/services/cmdb/rpc/internal/logic/resourceattribute"
	"github.com/admgo/admgo/services/cmdb/rpc/internal/svc"
	"github.com/admgo/admgo/services/cmdb/rpc/pb/pb"
)

type ResourceAttributeServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedResourceAttributeServer
}

func NewResourceAttributeServer(svcCtx *svc.ServiceContext) *ResourceAttributeServer {
	return &ResourceAttributeServer{
		svcCtx: svcCtx,
	}
}

// 创建资源属性
func (s *ResourceAttributeServer) CreateResourceAttribute(ctx context.Context, in *pb.CreateAttributeRequest) (*pb.CreateAttributeResponse, error) {
	l := resourceattributelogic.NewCreateResourceAttributeLogic(ctx, s.svcCtx)
	return l.CreateResourceAttribute(in)
}
