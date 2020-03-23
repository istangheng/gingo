package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// CorsMiddleware 跨域中间件
func CorsMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		config := cors.DefaultConfig()
		config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"}
		config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Cookie"}

		if gin.Mode() == gin.ReleaseMode {
			// 生产环境需要配置跨域域名，否则403
			config.AllowOrigins = []string{"http://www.example.com"}
		} else {
			config.AllowAllOrigins = true
		}

		config.AllowCredentials = true
		cors.New(config)
	}
}
