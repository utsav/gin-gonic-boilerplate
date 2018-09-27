package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/utsav/gin-gonic-boilerplate/db"
	"github.com/utsav/gin-gonic-boilerplate/models"
	"github.com/utsav/gin-gonic-boilerplate/structs"
)

//GetBlogs ...
func GetBlogs(c *gin.Context) {
	blogs := []models.Blog{}
	if err := db.DBCon.Preload("Comments").Preload("User").Find(&blogs).Error; err != nil {
		c.JSON(http.StatusNotFound, structs.Error{Code: http.StatusNotFound, Error: "getting error while processing db request"})
		return
	}

	c.JSON(http.StatusOK, blogs)
}

//GetBlog ...
func GetBlog(c *gin.Context) {
	ID, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, structs.Error{Code: http.StatusBadRequest, Error: "Please pass valid id"})
		return
	}

	blog := models.Blog{}
	if err := db.DBCon.Preload("Comments").Preload("User").First(&blog, ID).Error; err != nil {
		c.JSON(http.StatusNotFound, structs.Error{Code: http.StatusNotFound, Error: "record not found"})
		return
	}

	c.JSON(http.StatusOK, blog)
}

//CreateBlog ...
func CreateBlog(c *gin.Context) {

	var createBlog (structs.BlogFields)
	if err := c.ShouldBindWith(&createBlog, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, structs.Error{Code: http.StatusBadRequest, Error: err.Error()})
		return
	}

	var blog models.Blog
	blog.BlogFields = createBlog
	if err := db.DBCon.Create(&blog).Error; err != nil {
		c.JSON(http.StatusBadRequest, structs.Error{Code: http.StatusBadRequest, Error: "getting error while processing db request"})
		return
	}

	c.JSON(http.StatusOK, blog)
}

//UpdateBlog ...
func UpdateBlog(c *gin.Context) {
	ID, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, structs.Error{Code: http.StatusBadRequest, Error: "Please pass valid id"})
		return
	}

	var blog (models.Blog)
	blog.ID = uint(ID)

	if err := db.DBCon.First(&blog).Error; err != nil {
		c.JSON(http.StatusBadRequest, structs.Error{Code: http.StatusBadRequest, Error: "no blog find with this id"})
		return
	}

	var updateBlog structs.UpdateBlogFields
	if err := c.ShouldBindWith(&updateBlog, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, structs.Error{Code: http.StatusBadRequest, Error: err.Error()})
		return
	}

	blog.UpdateBlogFields = updateBlog
	if err := db.DBCon.Model(&blog).Updates(&blog).Error; err != nil {
		c.JSON(http.StatusBadRequest, structs.Error{Code: http.StatusBadRequest, Error: "getting error while processing db request"})
		return
	}

	c.JSON(http.StatusOK, blog)
}

//DeleteBlog ...
func DeleteBlog(c *gin.Context) {
	ID, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, structs.Error{Code: http.StatusBadRequest, Error: "Please pass valid id"})
		return
	}

	var blog (models.Blog)
	blog.ID = uint(ID)
	if err := db.DBCon.Delete(&blog).Error; err != nil {
		fmt.Println("err", err)
		c.JSON(http.StatusBadRequest, structs.Error{Code: http.StatusBadRequest, Error: "getting error while processing db request"})
		return
	}

	c.JSON(http.StatusOK, structs.Success{Code: http.StatusOK, Message: "blog deleted successfully"})
}
