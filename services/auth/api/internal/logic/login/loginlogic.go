package login

import (
	"context"
	"github.com/admgo/admgo/services/user/rpc/user"

	"github.com/admgo/admgo/services/auth/api/internal/svc"
	"github.com/admgo/admgo/services/auth/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginRsp, err error) {
	_, err = l.svcCtx.UserRPC.FindSingleUser(l.ctx, &user.FindSingleUserRequest{UserId: 1})

	rsp := types.BaseRsp{
		Msg:  "success",
		Code: 0,
	}

	return &types.LoginRsp{
		BaseRsp: rsp,
	}, err
}
