package main

import (
	"gingo/controller"
	"gingo/middleware"

	"github.com/gin-gonic/gin"
)

// CollectRoute 路由集合
func CollectRoute(r *gin.Engine) *gin.Engine {
	r.Use(middleware.CorsMiddleware())
	r.POST("/api/auth/register", controller.Register)
	r.POST("/api/auth/login", controller.Login)
	r.GET("/api/auth/info", middleware.AuthMiddleware(), controller.Info)
	return r
}
