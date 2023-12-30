package logic

import (
	"context"
	"github.com/admgo/admgo/services/user/rpc/pb"

	"github.com/admgo/admgo/services/user/rpc/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

type FindSingleUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindSingleUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindSingleUserLogic {
	return &FindSingleUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查找指定用户
func (l *FindSingleUserLogic) FindSingleUser(in *pb.FindSingleUserRequest) (*pb.FindSingleUserResponse, error) {
	// todo: add your logic here and delete this line

	return &pb.FindSingleUserResponse{}, nil
}
