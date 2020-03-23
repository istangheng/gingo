package config

import "gingo/common"

// Init 初始化配置
func Init() {
	common.InitDB()
	// defer db.Close()

	common.InitRedis()
}
