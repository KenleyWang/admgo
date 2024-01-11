package logic

import (
	"context"
	"github.com/admgo/admgo/services/user/rpc/internal/db/models"

	"github.com/admgo/admgo/services/user/rpc/internal/svc"
	"github.com/admgo/admgo/services/user/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateUserLogic {
	return &CreateUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// CreateUser 创建用户
func (l *CreateUserLogic) CreateUser(in *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	// todo: add your logic here and delete this line

	err := l.svcCtx.UserModel.Insert(l.ctx, &models.User{
		UserName: in.UserName,
		Name:     in.Name,
	})

	return &pb.CreateUserResponse{}, err
}
