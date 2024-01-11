package logic

import (
	"context"

	"github.com/admgo/admgo/services/user/rpc/internal/svc"
	"github.com/admgo/admgo/services/user/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindAllUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindAllUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindAllUserLogic {
	return &FindAllUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// FindAllUser 查找所有用户
func (l *FindAllUserLogic) FindAllUser(in *pb.FindAllUserRequest) (*pb.FindAllUserResponse, error) {
	// todo: add your logic here and delete this line

	return &pb.FindAllUserResponse{}, nil
}
