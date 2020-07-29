package initialize

import (
	"fmt"
	"gingo/internal/config"
	"gingo/model"
	gormadapter "github.com/casbin/gorm-adapter/v2"
	"github.com/jinzhu/gorm"
)

// DB 实例
var DB *gorm.DB

// InitDB 初始化数据库
func InitDB() *gorm.DB {
	conf := config.Conf
	fmt.Println(conf.MysqlDsn)
	db, err := gorm.Open("mysql", conf.MysqlDsn)
	if err != nil {
		panic("failed to connect database,err:" + err.Error())
	}

	db.AutoMigrate(&gormadapter.CasbinRule{})
	db.AutoMigrate(&model.User{})
	DB = db
	return db
}
