// Code generated by goctl. DO NOT EDIT.
// Source: cmdb.proto

package server

import (
	"context"

	"github.com/admgo/admgo/services/cmdb/rpc/internal/logic/attributetype"
	"github.com/admgo/admgo/services/cmdb/rpc/internal/svc"
	"github.com/admgo/admgo/services/cmdb/rpc/pb"
)

type AttributeTypeServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedAttributeTypeServer
}

func NewAttributeTypeServer(svcCtx *svc.ServiceContext) *AttributeTypeServer {
	return &AttributeTypeServer{
		svcCtx: svcCtx,
	}
}

// 创建资源属性类型
func (s *AttributeTypeServer) CreateAttributeType(ctx context.Context, in *pb.CreateAttributeTypeRequest) (*pb.CreateAttributeTypeResponse, error) {
	l := attributetypelogic.NewCreateAttributeTypeLogic(ctx, s.svcCtx)
	return l.CreateAttributeType(in)
}