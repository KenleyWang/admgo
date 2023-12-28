package logic

import (
	"context"

	"github.com/admgo/admgo/services/cmdb/api/internal/svc"
	"github.com/admgo/admgo/services/cmdb/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CmdbLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCmdbLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CmdbLogic {
	return &CmdbLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CmdbLogic) Cmdb(req *types.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}
