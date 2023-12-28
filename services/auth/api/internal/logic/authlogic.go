package logic

import (
	"context"
	"github.com/admgo/admgo/services/auth/api/internal/svc"
	"github.com/admgo/admgo/services/auth/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AuthLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAuthLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AuthLogic {
	return &AuthLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AuthLogic) Auth(req *types.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line
	type Result struct {
		Host string
		user string
	}
	var result Result
	l.svcCtx.DBW.Raw("SELECT host, user FROM user WHERE user = ?", "addmgo").Scan(&result)
	return
}
