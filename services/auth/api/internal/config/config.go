package config

import (
	"github.com/admgo/admgo/lib/db/config"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	config.DBConfig
}
