package config

type Server struct {
	Port int    `mapstructure:"port"`
	Mode string `mapstructure:"mode"`
}

type Mysql struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Dbname   string `mapstructure:"dbname"`
	Config   string `mapstructure:"config"`
}

type JWT struct {
	SigningKey string `mapstructure:"signing_key"`
}

// 汇总结构体
type Config struct {
	Server Server `mapstructure:"server"`
	Mysql  Mysql  `mapstructure:"mysql"`
	JWT    JWT    `mapstructure:"jwt"`
}
