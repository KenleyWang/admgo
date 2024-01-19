package resourcetypelogic

import (
	"context"

	"github.com/admgo/admgo/services/cmdb/rpc/internal/svc"
	"github.com/admgo/admgo/services/cmdb/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateResourceTypeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateResourceTypeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateResourceTypeLogic {
	return &CreateResourceTypeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 创建资源类型
func (l *CreateResourceTypeLogic) CreateResourceType(in *pb.CreateResourceTypeRequest) (*pb.CreateResourceTypeResponse, error) {
	// todo: add your logic here and delete this line

	return &pb.CreateResourceTypeResponse{}, nil
}
