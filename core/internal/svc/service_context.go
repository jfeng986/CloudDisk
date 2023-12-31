package svc

import (
	"go-zero-cloud-disk/core/internal/config"
	"go-zero-cloud-disk/core/internal/middleware"
	"go-zero-cloud-disk/core/models"

	"github.com/go-redis/redis/v8"
	"github.com/zeromicro/go-zero/rest"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config config.Config
	MDB    *gorm.DB
	RDB    *redis.Client
	Auth   rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		MDB:    models.InitMysql(c.Mysql.MysqlAddr),
		RDB:    models.InitRedis(c),
		Auth:   middleware.NewAuthMiddleware().Handle,
	}
}
