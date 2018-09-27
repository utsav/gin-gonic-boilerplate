package models

import (
	"github.com/utsav/gin-gonic-boilerplate/db"
	"github.com/utsav/gin-gonic-boilerplate/structs"
)

type Comment struct {
	Model
	structs.CommentFields
}

func init() {
	db.DBCon.AutoMigrate(&Comment{})
}
