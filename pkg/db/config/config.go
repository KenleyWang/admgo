package config

type Config struct {
	Username     string `json:"username"`
	Password     string `json:"password"`
	Host         string `json:"host"`
	Port         string `json:"port"`
	Database     string `json:"database"`
	Charset      string `json:"charset"`
	MaxOpenConns int    `json:"max_open_conns"`
	MaxIdleConns int    `json:"max_idle_conns"`
	MaxLifetime  int    `json:"max_lifetime"`
	TimeOut      int    `json:"timeout"`
	WriteTimeOut int    `json:"write_timeout"`
	ReadTimeOut  int    `json:"read_timeout"`
}
