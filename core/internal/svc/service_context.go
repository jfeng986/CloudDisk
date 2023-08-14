package svc

import (
	"go-zero-cloud-disk/core/internal/config"
	"go-zero-cloud-disk/core/models"

	"github.com/go-redis/redis/v8"
	"xorm.io/xorm"
)

type ServiceContext struct {
	Config config.Config
	Engine *xorm.Engine
	RDB    *redis.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Engine: models.InitMysql(c.Mysql.MysqlAddr),
		RDB:    models.InitRedis(c),
	}
}
