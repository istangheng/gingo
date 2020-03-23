package main

import (
	"gingo/config"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

// main 入口
func main() {
	// 初始化配置
	config.Init()

	r := gin.Default()
	r = CollectRoute(r)
	r.Run()
}
