package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/utsav/gin-gonic-boilerplate/db"
	"github.com/utsav/gin-gonic-boilerplate/models"
	"github.com/utsav/gin-gonic-boilerplate/structs"
)

//CreateComment ...
func CreateComment(c *gin.Context) {

	var createComment structs.CommentFields
	if err := c.ShouldBindWith(&createComment, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, structs.Error{Code: http.StatusBadRequest, Error: err.Error()})
		return
	}

	var comment models.Comment
	comment.CommentFields = createComment
	if err := db.DBCon.Create(&comment).Error; err != nil {
		c.JSON(http.StatusBadRequest, structs.Error{Code: http.StatusBadRequest, Error: "getting error while processing db request"})
		return
	}

	c.JSON(http.StatusOK, comment)
}
