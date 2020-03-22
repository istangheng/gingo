package main

import (
	"gingo/controller"

	"github.com/gin-gonic/gin"
)

// CollectRoute 路由集合
func CollectRoute(r *gin.Engine) *gin.Engine {
	r.POST("/api/auth/register", controller.Register)
	r.POST("/api/auth/login", controller.Login)
	return r
}
