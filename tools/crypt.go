package tools

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
)

// EncryptedPwd 加密密码
func EncryptedPwd(password string) (hashPwd string, err error) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", errors.New("加密错误")
	}
	return string(hashPassword), nil
}

// ComparePwd 密码校验
func ComparePwd(hashedPassword, password string) (err error) {
	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return
}
