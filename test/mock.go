package test

import (
	"fmt"
	"gingo/initialize"
	"gingo/model"
	"gingo/model/request"
)

func init() {
	//config.InitConf()
	//initialize.InitDB()
}

func flushTable(table string) error {
	db := initialize.DB
	sql := fmt.Sprintf("truncate %s", table)
	return db.Exec(sql).Error
}

var testUser1 = model.User{
	Name:      "Name1",
	Telephone: "11111111111",
	Password:  "Password1",
}

var testLoginRequest1 = request.LoginRequest{
	Telephone: "22222222222",
	Password:  "Password1",
}
