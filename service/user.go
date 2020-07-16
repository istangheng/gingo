package service

import (
	"errors"
	"gingo/initialize"
	"gingo/model"
	"gingo/model/request"
	"gingo/tools"
)

// Register 用户注册
func Register(user model.User) (token string, err error) {
	db := initialize.DB
	// 数据验证
	if len(user.Telephone) != 11 {
		return "", errors.New("手机号必须为11位")
	}
	if len(user.Password) < 6 {
		return "", errors.New("密码必须大于6位")
	}

	// 如果名字没有传入，随机生成10位字符串名字
	if len(user.Name) == 0 {
		user.Name = tools.RandomString(10)
	}

	// 判断手机号是否存在
	count := 0
	db.Model(&model.User{}).Where("telephone = ?", user.Telephone).Count(&count)
	if count > 0 {
		return "", errors.New("用户已存在")
	}

	// 创建用户
	hashPassword, err := tools.EncryptedPwd(user.Password)
	if err != nil {
		return "", err
	}
	user.Password = hashPassword
	db.Create(&user)

	// 发放token
	token, err = tools.ReleaseToken(user)
	if err != nil {
		return "", errors.New("系统异常")
	}
	return token, nil
}

// Login 用户登录
func Login(loginRequest request.LoginRequest) (token string, err error) {
	db := initialize.DB
	// 数据验证
	if len(loginRequest.Telephone) != 11 {
		return "", errors.New("手机号必须为11位")
	}
	if len(loginRequest.Password) < 6 {
		return "", errors.New("密码必须大于6位")
	}
	// 判断手机号是否存在
	var user model.User
	db.Where("telephone = ?", loginRequest.Telephone).Find(&user)
	if user.ID == 0 {
		return "", errors.New("用户已存在")
	}
	// 判断密码是否存在
	if err := tools.ComparePwd(user.Password, loginRequest.Password); err != nil {
		return "", errors.New("密码错误")
	}
	// 发放token
	token, err = tools.ReleaseToken(user)
	if err != nil {
		return "", errors.New("系统异常")
	}
	return
}
