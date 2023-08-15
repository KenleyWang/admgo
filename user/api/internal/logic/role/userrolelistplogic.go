package role

import (
	"context"

	"github.com/admgo/admgo/user/api/internal/svc"
	"github.com/admgo/admgo/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserRoleListPLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserRoleListPLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRoleListPLogic {
	return &UserRoleListPLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserRoleListPLogic) UserRoleListP() (resp []types.UserRoleResp, err error) {
	// todo: add your logic here and delete this line

	return
}
