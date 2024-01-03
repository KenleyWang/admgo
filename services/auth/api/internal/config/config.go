package config

import (
	db "github.com/admgo/admgo/pkg/db/config"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	DB db.Config
}
