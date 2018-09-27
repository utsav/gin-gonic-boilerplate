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

//GetUsers ...
func GetUsers(c *gin.Context) {
	users := []models.User{}
	if err := db.DBCon.Find(&users).Error; err != nil {
		c.JSON(http.StatusNotFound, structs.Error{Code: http.StatusNotFound, Error: "getting error while processing db request"})
		return
	}

	c.JSON(http.StatusOK, users)
}

//GetUser ...
func GetUser(c *gin.Context) {
	ID, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, structs.Error{Code: http.StatusBadRequest, Error: "Please pass valid id"})
		return
	}

	user := models.User{}
	if err := db.DBCon.First(&user, ID).Error; err != nil {
		c.JSON(http.StatusNotFound, structs.Error{Code: http.StatusNotFound, Error: "record not found"})
		return
	}

	if err != nil {
		fmt.Println(err)
		return
	}

	c.JSON(http.StatusOK, user)
}

//CreateUser ...
func CreateUser(c *gin.Context) {

	var createUser (structs.UserFields)
	if err := c.ShouldBindWith(&createUser, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, structs.Error{Code: http.StatusBadRequest, Error: err.Error()})
		return
	}

	var user models.User
	user.UserFields = createUser
	if err := db.DBCon.Create(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, structs.Error{Code: http.StatusBadRequest, Error: "getting error while processing db request"})
		return
	}

	c.JSON(http.StatusOK, user)
}

//UpdateUser ...
func UpdateUser(c *gin.Context) {
	ID, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, structs.Error{Code: http.StatusBadRequest, Error: "Please pass valid id"})
		return
	}

	var user (models.User)
	user.ID = uint(ID)

	if err := db.DBCon.First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, structs.Error{Code: http.StatusBadRequest, Error: "no user find with this id"})
		return
	}

	var updateUser (structs.UpdateUserFields)

	if err := c.ShouldBindWith(&updateUser, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, structs.Error{Code: http.StatusBadRequest, Error: err.Error()})
		return
	}

	user.UpdateUserFields = updateUser

	if err := db.DBCon.Model(&user).Updates(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, structs.Error{Code: http.StatusBadRequest, Error: "getting error while processing db request"})
		return
	}

	c.JSON(http.StatusOK, user)
}

//DeleteUser ...
func DeleteUser(c *gin.Context) {
	ID, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, structs.Error{Code: http.StatusBadRequest, Error: "Please pass valid id"})
		return
	}

	var user (models.User)
	user.ID = uint(ID)
	if err := db.DBCon.Delete(&user).Error; err != nil {
		fmt.Println("err", err)
		c.JSON(http.StatusBadRequest, structs.Error{Code: http.StatusBadRequest, Error: "getting error while processing db request"})
		return
	}

	c.JSON(http.StatusOK, structs.Success{Code: http.StatusOK, Message: "user deleted successfully"})
}
