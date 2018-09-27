package models

import (
	"github.com/utsav/gin-gonic-boilerplate/db"
	"github.com/utsav/gin-gonic-boilerplate/structs"
)

type Blog struct {
	Model
	structs.BlogFields
	User     User      `json:"user"`
	Comments []Comment `json:"comments"`
}

func init() {
	db.DBCon.AutoMigrate(&Blog{})
}
