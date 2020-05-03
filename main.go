package main

import (
	"gingo/config"
	"gingo/router"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

// main 入口
func main() {
	// 初始化配置
	config.Init()

	r := gin.Default()
	r = router.CollectRoute(r)
	r.Run()
}
