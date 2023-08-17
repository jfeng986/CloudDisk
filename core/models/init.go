package models

import (
	"log"
	"time"

	"go-zero-cloud-disk/core/internal/config"

	"github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func InitMysql(mysqlAddr string) *gorm.DB {
	conn := mysqlAddr
	var ormLogger logger.Interface
	ormLogger = logger.Default.LogMode(logger.Info)
	mdb, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       conn,
		DefaultStringSize:         256,
		DisableDatetimePrecision:  true,
		DontSupportRenameIndex:    true,
		DontSupportRenameColumn:   true,
		SkipInitializeWithVersion: false,
	}), &gorm.Config{
		Logger: ormLogger,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		panic(err)
	} else {
		log.Println("database connect success")
	}
	sqlDB, _ := mdb.DB()
	sqlDB.SetMaxIdleConns(20)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Second * 30)

	migration(mdb)

	return mdb
}

func migration(mdb *gorm.DB) {
	err := mdb.Set("gorm:table_options", "charset=utf8mb4").
		AutoMigrate(&UserBasic{}, &RepositoryPool{}, &UserRepository{})
	if err != nil {
		return
	}
}

func InitRedis(c config.Config) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     c.Redis.RedisAddr,
		Password: "", // no password set
		DB:       0,
	})
}
