package models

import (
	"github.com/utsav/gin-gonic-boilerplate/db"
	"github.com/utsav/gin-gonic-boilerplate/services"
	"github.com/utsav/gin-gonic-boilerplate/structs"
)

type User struct {
	Model
	structs.UserFields
}

func init() {
	db.DBCon.AutoMigrate(&User{})
}

func (u *User) BeforeCreate() (err error) {
	if u.Password != "" {
		u.Password = services.HashAndSalt([]byte(u.Password))
	}
	return
}
