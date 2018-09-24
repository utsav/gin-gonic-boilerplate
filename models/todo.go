package models

import (
	"github.com/jinzhu/gorm"
	"github.com/utsav/gin-gonic-boilerplate/db"
)

type Todo struct {
	gorm.Model
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
}

func init() {
	db.DBCon.AutoMigrate(&Todo{})
}
