package models

import (
	"fmt"
	"log"

	"go-zero-cloud-disk/core/internal/config"

	"github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql"

	"xorm.io/xorm"
)

func InitMysql(mysqlAddr string) *xorm.Engine {
	engine, err := xorm.NewEngine("mysql", mysqlAddr)
	if err != nil {
		log.Printf("Xorm New Engine Error:%v", err)
		return nil
	}

	return engine
}

func InitRedis(c config.Config) *redis.Client {
	fmt.Println(c)
	return redis.NewClient(&redis.Options{
		Addr:     c.Redis.RedisAddr,
		Password: "", // no password set
		DB:       0,
	})
}
