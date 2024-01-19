package resourcemodellogic

import (
	"context"

	"github.com/admgo/admgo/services/cmdb/rpc/internal/svc"
	"github.com/admgo/admgo/services/cmdb/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateResourceModelLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateResourceModelLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateResourceModelLogic {
	return &CreateResourceModelLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 创建资源模型
func (l *CreateResourceModelLogic) CreateResourceModel(in *pb.CreateResourceModelRequest) (*pb.CreateResourceModelResponse, error) {
	// todo: add your logic here and delete this line

	return &pb.CreateResourceModelResponse{}, nil
}
