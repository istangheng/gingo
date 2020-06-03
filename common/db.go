package common

import (
	"gingo/config"
	"gingo/model"

	"github.com/jinzhu/gorm"
)

// DB 实例
var DB *gorm.DB

// InitDB 初始化数据库
func InitDB() *gorm.DB {
	conf := config.Conf
	db, err := gorm.Open("mysql", conf.MysqlDsn)
	if err != nil {
		panic("failed to connect database,err:" + err.Error())
	}

	db.AutoMigrate(&model.User{})

	DB = db

	return db
}

// GetDB 获取数据库实例
func GetDB() *gorm.DB {
	return DB
}
