package login

import (
	"context"
	"database/sql"
	"time"

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
	// todo: add your logic here and delete this line

	type User struct {
		Id       int64          `db:"id"`
		Name     sql.NullString `db:"name"`     // The username
		Password string         `db:"password"` // The user password
		Mobile   string         `db:"mobile"`   // The mobile phone number
		Gender   string         `db:"gender"`   // gender,male|female|unknown
		Nickname string         `db:"nickname"` // The nickname
		Type     int64          `db:"type"`     // The user type, 0:normal,1:vip, for test golang keyword
		CreateAt sql.NullTime   `db:"create_at"`
		UpdateAt time.Time      `db:"update_at"`
	}

	var result User

	// Raw SQL
	l.svcCtx.DB.Raw("SELECT name FROM users WHERE name = ?", "users").Scan(&result)

	return

}
