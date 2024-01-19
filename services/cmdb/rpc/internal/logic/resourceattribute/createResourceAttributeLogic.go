package resourceattributelogic

import (
	"context"

	"github.com/admgo/admgo/services/cmdb/rpc/internal/svc"
	"github.com/admgo/admgo/services/cmdb/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateResourceAttributeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateResourceAttributeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateResourceAttributeLogic {
	return &CreateResourceAttributeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 创建资源属性
func (l *CreateResourceAttributeLogic) CreateResourceAttribute(in *pb.CreateResourceAttributeRequest) (*pb.CreateResourceAttributeResponse, error) {
	// todo: add your logic here and delete this line

	return &pb.CreateResourceAttributeResponse{}, nil
}