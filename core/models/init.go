package models

import (
	"log"

	"github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql"

	"xorm.io/xorm"
)

var (
	Engine = InitMysql("root:root@tcp(localhost:3306)/cloud_disk?charset=utf8mb4&parseTime=True&loc=Local")
	RDB    = InitRedis("localhost:6379", 1, "")
)

func InitMysql(dataSource string) *xorm.Engine {
	engine, err := xorm.NewEngine("mysql", dataSource)
	if err != nil {
		log.Printf("Xorm New Engine Error:%v", err)
		return nil
	}

	return engine
}

func InitRedis(addr string, db int, password string) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,
	})
}
