package api

import (
	"gingo/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Ping 测试
func Ping(ctx *gin.Context) {
	resp := service.Ping()
	ctx.JSON(http.StatusOK, gin.H{"msg": resp})
}
