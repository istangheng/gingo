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
		ctx.JSON(http.StatusBadRequest, response.JSONResult{
			Code: 400,
			Msg:  err.Error(),
		})
		return
	}

	token, err := service.Register(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.JSONResult{
			Code: 500,
			Msg:  err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, response.JSONResult{
		Code: 200,
		Msg:  "注册成功",
		Data: map[string]string{"token": token},
	})
}

// Login 登录
func Login(ctx *gin.Context) {
	var user request.LoginRequest
	if err := ctx.ShouldBind(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, response.JSONResult{
			Code: 400,
			Msg:  err.Error(),
		})
		return
	}

	token, err := service.Login(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.JSONResult{
			Code: 500,
			Msg:  err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, response.JSONResult{
		Code: 200,
		Msg:  "登录成功",
		Data: map[string]string{"token": token},
	})
}

// Info 获取用户信息
func Info(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	u := response.UserResp{
		Name:      user.(string),
		Telephone: user.(string),
	}
	ctx.JSON(http.StatusOK, response.JSONResult{
		Code: 200,
		Msg:  "获取用户信息",
		Data: u,
	})
}
