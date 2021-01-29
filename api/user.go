package api

import (
	"gingo/initialize"
	"gingo/model"
	"gingo/model/request"
	"gingo/model/response"
	"gingo/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Register 用户注册
// Register 用户注册
// @Register 用户注册
// @Description Register 用户注册
// @Tags Admin
// @Accept  json
// @Produce json
// @Param login body request.RegisterRequest true "帐号注册"
// @Success 200 {object} response.JSONResult{code=int,msg=string,data=string} "注册成功"
// @Failure 400  "参数错误"
// @Failure 500  "系统异常"
// @Router /api/auth/register [post]
func Register(ctx *gin.Context) {
	var user request.RegisterRequest
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
		Data: token,
	})
}

// Login 登录
// Login 后台管理员登录
// @Login 后台管理员登录
// @Description Login 后台管理员登录
// @Tags Admin
// @Accept  json
// @Produce json
// @Param login body request.LoginRequest true "管理员帐号登录"
// @Success 200 {object} response.JSONResult{code=int,msg=string,data=string} "登陆成功"
// @Failure 400  "参数错误"
// @Failure 500  "系统异常"
// @Router /api/auth/login [post]
func Login(ctx *gin.Context) {
	var user request.LoginRequest
	if err := ctx.ShouldBind(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, response.JSONResult{
			Code: 400,
			Msg:  err.Error(),
			Data: "",
		})
		return
	}

	token, err := service.Login(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.JSONResult{
			Code: 500,
			Msg:  err.Error(),
			Data: "",
		})
		return
	}
	ctx.JSON(http.StatusOK, response.JSONResult{
		Code: 200,
		Msg:  "登录成功",
		Data: token,
	})
}

// Info 获取用户信息
// Info 获取用户信息
// @Info 获取用户信息
// @Description Info 获取用户信息
// @Tags Admin
// @Accept  json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Success 200 {object} response.JSONResult{code=int,msg=string,data=response.UserResp} "登陆成功"
// @Router /api/auth/info [get]
func Info(ctx *gin.Context) {
	userID, _ := ctx.Get("userID")
	DB := initialize.DB
	var user model.User
	DB.Model(&model.User{}).Where("id = ?", userID.(uint)).First(&user)
	u := response.UserResp{
		Name:      user.Name,
		Telephone: user.Telephone,
	}
	ctx.JSON(http.StatusOK, response.JSONResult{
		Code: 200,
		Msg:  "获取用户信息",
		Data: u,
	})
}
