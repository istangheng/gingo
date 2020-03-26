package model

import "github.com/jinzhu/gorm"

// User 用户
type User struct {
	gorm.Model
	Name      string `gorm:"type:varchar(20);not null" form:"name" json:"name"`
	Telephone string `gorm:"type:varchar(11);not null;unique" form:"telephone" json:"telephone"`
	Password  string `grom:"type:varchar(20):not null" form:"password" json:"password"`
}
