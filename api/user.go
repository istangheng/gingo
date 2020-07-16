package api

import (
	"gingo/model"
	"gingo/model/request"
	"gingo/model/response"
	"gingo/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Register 用户注册
func Register(ctx *gin.Context) {
	var user model.User
	if err := ctx.ShouldBind(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := service.Register(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"token": token, "msg": "注册成功"})
}

// Login 登录
func Login(ctx *gin.Context) {
	var user request.LoginRequest
	if err := ctx.ShouldBind(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := service.Login(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"token": token, "msg": "登录成功"})
}

// Info 获取用户信息
func Info(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	u := response.UserResponse{
		Name:      user.(model.User).Name,
		Telephone: user.(model.User).Telephone,
	}
	ctx.JSON(http.StatusOK, gin.H{"user": u})
}
