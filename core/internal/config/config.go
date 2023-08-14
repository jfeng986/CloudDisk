package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf
	Mysql struct {
		MysqlAddr string
	}
	Redis struct {
		RedisAddr string
	}
}
