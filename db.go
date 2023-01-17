package corelibs

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DBConn *gorm.DB
)

func DBEstablish() *gorm.DB {
	dsn := "host=" + GetConfig("DB_HOST") + " user=" + GetConfig("DB_USER") + " password=" + GetConfig("DB_PASS") + " dbname=" + GetConfig("DB_NAME") + " port=" + GetConfig("DB_PORT") + " sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		CommonLogger().Error(err)
		panic("failed to connect database")
	}

	return db
}

func DBConnection() *gorm.DB {
	return DBConn
}
