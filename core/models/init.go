package models

import (
	"log"

	_ "github.com/go-sql-driver/mysql"

	"xorm.io/xorm"
)

var Engine = Init("root:root@tcp(localhost:3306)/cloud_disk?charset=utf8mb4&parseTime=True&loc=Local")

func Init(dataSource string) *xorm.Engine {
	engine, err := xorm.NewEngine("mysql", dataSource)
	if err != nil {
		log.Printf("Xorm New Engine Error:%v", err)
		return nil
	}

	return engine
}
