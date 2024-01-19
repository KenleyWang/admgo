package resourcelogic

import (
	"context"

	"github.com/admgo/admgo/services/cmdb/rpc/internal/svc"
	"github.com/admgo/admgo/services/cmdb/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateResourceLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateResourceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateResourceLogic {
	return &CreateResourceLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 创建资源
func (l *CreateResourceLogic) CreateResource(in *pb.CreateResourceRequest) (*pb.CreateResourceResponse, error) {
	// todo: add your logic here and delete this line

	return &pb.CreateResourceResponse{}, nil
}
