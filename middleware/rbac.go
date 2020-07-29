package middleware

import (
	"gingo/initialize"
	"gingo/model"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

var excludepath = []string{"login", "logout"}

// RbacMiddleware 权限控制middle
func RbacMiddleware() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		if ctx.Request.URL.Path == "/" {
			ctx.Next()
			return
		}
		for _, url := range excludepath {
			if strings.Contains(ctx.Request.URL.Path, url) {
				ctx.Next()
				return
			}
		}

		user, _ := ctx.Get("user")
		role := user.(model.User).Role
		requestURL := ctx.Request.URL.Path
		method := ctx.Request.Method
		pass, err := initialize.Enforcer.Enforce(role, requestURL, method)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "message": "权限不足"})
			ctx.Abort()
			return
		}
		if !pass {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "message": "权限不足"})
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}
