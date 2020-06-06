package main

import (
	"gingo/common"
	"gingo/config"
	"gingo/router"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	common.InitDB()
	common.InitRedisPool()
}

// main 入口
func main() {
	conf := config.Conf

	r := gin.Default()
	r = router.CollectRoute(r)
	r.Run(conf.Server.Port)
}
