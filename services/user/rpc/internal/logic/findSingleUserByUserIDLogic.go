package logic

import (
	"context"

	"github.com/admgo/admgo/services/user/rpc/internal/svc"
	"github.com/admgo/admgo/services/user/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindSingleUserByUserIDLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindSingleUserByUserIDLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindSingleUserByUserIDLogic {
	return &FindSingleUserByUserIDLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// FindSingleUserByUserID 根据用户ID查找指定用户
func (l *FindSingleUserByUserIDLogic) FindSingleUserByUserID(in *pb.FindSingleUserByUserIDRequest) (*pb.FindSingleUserByUserIDResponse, error) {

	return &pb.FindSingleUserByUserIDResponse{}, nil
}
