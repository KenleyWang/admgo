package svc

import (
	"github.com/admgo/admgo/pkg/db"
	dbconfig "github.com/admgo/admgo/pkg/db/config"
	"github.com/admgo/admgo/services/user/rpc/internal/config"
	"github.com/admgo/admgo/services/user/rpc/internal/model"
)

type ServiceContext struct {
	Config    config.Config
	UserModel model.UserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	dbInstance := db.MustNewMysql(&dbconfig.Config{
		Username:     c.DB.Username,
		Password:     c.DB.Password,
		Host:         c.DB.Host,
		Port:         c.DB.Port,
		Database:     c.DB.Database,
		Charset:      c.DB.Charset,
		MaxOpenConns: c.DB.MaxOpenConns,
		MaxLifetime:  c.DB.MaxLifetime,
		TimeOut:      c.DB.TimeOut,
		WriteTimeOut: c.DB.WriteTimeOut,
		ReadTimeOut:  c.DB.ReadTimeOut,
	})
	return &ServiceContext{
		Config:    c,
		UserModel: model.NewUserModel(dbInstance.DB),
	}
}
