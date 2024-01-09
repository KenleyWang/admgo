package code

import "github.com/admgo/admgo/pkg/errorx"

var (
	UserEmpty                   = errorx.NewRpcStatus(110101, "no user")
	IncorrectUserNameORPassword = errorx.NewRpcStatus(110102, "incorrect user name or password")
)
