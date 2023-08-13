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
	/*
		xormLogFile, err := os.OpenFile("logs/xorm_sql.log", os.O_APPEND|os.O_WRONLY, 6)
		if err != nil {
			logx.Errorf("Open xorm_sql.log failed:%v", err)
			return nil
		}
		engine.SetLogger(xlog.NewSimpleLogger(xormLogFile)) // 将日志重定向到文件中
		engine.ShowSQL(true)                                // 打印出生成的SQL语句
		engine.Logger().SetLevel(xlog.LOG_DEBUG)            // 打印调试及以上的信息
	*/
	return engine
}
