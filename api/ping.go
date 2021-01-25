package api

import (
	"gingo/model/response"
	"gingo/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Ping 测试
func Ping(ctx *gin.Context) {
	resp := service.Ping()
	ctx.JSON(http.StatusOK, response.JSONResult{
		Code: 200,
		Msg:  "ping...",
		Data: resp,
	})
}
