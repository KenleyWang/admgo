package svc

import (
	"fmt"
	"github.com/admgo/admgo/lib/db"
	"github.com/admgo/admgo/services/auth/api/internal/config"
	"github.com/admgo/admgo/services/auth/api/internal/middleware"
	"github.com/zeromicro/go-zero/rest"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config          config.Config
	AuthInterceptor rest.Middleware
	DBW             *gorm.DB
	DBR             *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	conf, ok := c.DB["default"]
	if !ok {
		panic(fmt.Sprintf("db config %s not found", "default"))
	}
	writeCachekey := fmt.Sprintf("%s_write", "default")
	readCachekey := fmt.Sprintf("%s_read", "default")
	return &ServiceContext{
		Config:          c,
		DBW:             db.Load(conf.Write, conf.Log, writeCachekey, c.Mode, "/Users/kenley/Documents/个人项目/go-project/admgo/services/auth/api"),
		DBR:             db.Load(conf.Read, conf.Log, readCachekey, c.Mode, "/Users/kenley/Documents/个人项目/go-project/admgo/services/auth/api"),
		AuthInterceptor: middleware.NewAuthInterceptorMiddleware().Handle,
	}
}
