package config

type DBConfig struct {
	DB map[string]dbitem `yaml:"db" json:",optional"`
}
