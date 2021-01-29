package router

import (
	"gingo/api"
	"gingo/middleware"

	_ "gingo/docs" // swagger docs

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// CollectRoute 路由集合
func CollectRoute(r *gin.Engine) *gin.Engine {
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Use(middleware.CorsMiddleware())
	r.POST("/api/auth/register", api.Register)
	r.POST("/api/auth/login", api.Login)
	r.GET("/api/auth/info", middleware.AuthMiddleware(), api.Info)
	r.GET("/ping", middleware.AuthMiddleware(), middleware.RbacMiddleware(), api.Ping)
	return r
}
