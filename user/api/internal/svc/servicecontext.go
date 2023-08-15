package svc

import (
	"github.com/admgo/admgo/user/api/internal/config"
	"github.com/admgo/admgo/user/api/internal/middleware"
	"github.com/zeromicro/go-zero/rest"
)

type ServiceContext struct {
	Config config.Config
	Num    rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Num:    middleware.NewNumMiddleware().Handle,
	}
}
