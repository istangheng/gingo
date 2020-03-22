package main

import (
	"gingo/common"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

// main 入口
func main() {
	db := common.InitDB()
	defer db.Close()
	r := gin.Default()
	r = CollectRoute(r)
	r.Run()
}
