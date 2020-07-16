package router

import (
	"gingo/api"
	"gingo/middleware"

	"github.com/gin-gonic/gin"
)

// CollectRoute 路由集合
func CollectRoute(r *gin.Engine) *gin.Engine {
	r.Use(middleware.CorsMiddleware())
	r.POST("/api/auth/register", api.Register)
	r.POST("/api/auth/login", api.Login)
	r.GET("/api/auth/info", middleware.AuthMiddleware(), api.Info)
	return r
}
