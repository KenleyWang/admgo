package svc

import (
	"github.com/admgo/admgo/pkg/db"
	dbconfig "github.com/admgo/admgo/pkg/db/config"
	"github.com/admgo/admgo/services/auth/api/internal/config"
	"github.com/admgo/admgo/services/auth/api/internal/middleware"
	"github.com/admgo/admgo/services/user/rpc/user"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config          config.Config
	AuthInterceptor rest.Middleware
	DB              *db.DB
	UserRPC         user.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	db := db.MustNewMysql(&dbconfig.Config{
		Username:     c.DB.Username,
		Password:     c.DB.Password,
		Host:         c.DB.Host,
		Port:         c.DB.Port,
		Database:     c.DB.Database,
		Charset:      c.DB.Charset,
		MaxOpenConns: c.DB.MaxOpenConns,
		MaxIdleConns: c.DB.MaxIdleConns,
		MaxLifetime:  c.DB.MaxLifetime,
		TimeOut:      c.DB.TimeOut,
		ReadTimeOut:  c.DB.ReadTimeOut,
		WriteTimeOut: c.DB.WriteTimeOut,
	})
	userRPC := zrpc.MustNewClient(c.Rpc["user.rpc"])
	return &ServiceContext{
		Config:          c,
		DB:              db,
		UserRPC:         user.NewUser(userRPC),
		AuthInterceptor: middleware.NewAuthInterceptorMiddleware().Handle,
	}
}
