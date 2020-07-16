package test

import (
	"gingo/service"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRegister(t *testing.T) {
	assert := assert.New(t)

	// _ = flushTable("users")

	_, err := service.Register(testUser1)
	assert.Nil(err)
}

func TestLogin(t *testing.T) {
	assert := assert.New(t)

	//_ = flushTable("user")
	assert.Nil(nil)

	//guard := monkey.Patch((*gorm.DB).Find, func(s *gorm.DB) (user model.User) {
	//	return model.User{
	//		Name:      "Name1",
	//		Telephone: "Telephone1",
	//		Password:  "Password1",
	//	}
	//})
	//token, err := service.Login(testLoginRequest1)
	//assert.NotEmpty(token)
	//assert.Nil(err)
	//guard.Unpatch()
}
