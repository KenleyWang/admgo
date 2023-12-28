package svc

import (
	"github.com/admgo/admgo/services/storager/storager/internal/config"
)

type ServiceContext struct {
	Config config.Config
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
