package common

import (
	"gingo/model"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

// DB 注释
var DB *gorm.DB

// InitDB 初始化数据库
func InitDB() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	mysqlDSN := os.Getenv("MYSQL_DSN")
	db, err := gorm.Open("mysql", mysqlDSN)
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
