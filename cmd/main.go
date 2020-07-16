package main

import (
	"gingo/initialize"
	"gingo/internal/config"
	"gingo/router"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	config.InitConf()
	initialize.InitDB()
	initialize.InitRedisPool()
}

// main 入口
func main() {
	conf := config.Conf

	r := gin.Default()
	r = router.CollectRoute(r)
	r.Run(conf.Server.Port)
}
