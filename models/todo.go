package models

import (
	"time"

	"github.com/utsav/gin-gonic-boilerplate/db"
)

type Todo struct {
	ID          uint       `json:"id" gorm:"primary_key"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at" sql:"index"`
	Title       string     `json:"title" binding:"required"`
	Description string     `json:"description" binding:"required"`
}

func init() {
	db.DBCon.AutoMigrate(&Todo{})
}
