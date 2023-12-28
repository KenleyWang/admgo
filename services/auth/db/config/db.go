package config

type DBItemConf struct {
	Host         string `yaml:"host" json:",optional"`
	Port         string `yaml:"port" json:",optional"`
	Database     string `yaml:"database" json:",optional"`
	Username     string `yaml:"username" json:",optional"`
	Password     string `yaml:"password" json:",optional"`
	Charset      string `yaml:"charset" json:",optional"`
	TimeOut      int    `yaml:"timeout" json:",optional"`
	WriteTimeOut int    `yaml:"write_time_out" json:",optional"`
	ReadTimeOut  int    `yaml:"read_time_out" json:",optional"`
	MaxIdleConns int    `yaml:"max_idle_conns" json:",optional"`
	MaxOpenConns int    `yaml:"max_open_conns" json:",optional"`
}

type DBLog struct {
	Enable bool   `yaml:"enable" json:",optional"`
	Level  string `yaml:"level" json:",optional"`
	Path   string `yaml:"path" json:",optional"`
	Type   string `yaml:"type" json:",optional"`
	Format string `yaml:"format" json:",optional"`
}

type dbitem struct {
	Log   DBLog      `yaml:"log" json:",optional"`
	Write DBItemConf `yaml:"write" json:",optional"`
	Read  DBItemConf `yaml:"read" json:",optional"`
}
