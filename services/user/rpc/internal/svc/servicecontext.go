package svc

import (
	"github.com/admgo/admgo/pkg/db"
	"github.com/admgo/admgo/services/user/rpc/internal/config"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

type ServiceContext struct {
	Config   config.Config
	BizRedis *redis.Redis
	DB       *db.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
