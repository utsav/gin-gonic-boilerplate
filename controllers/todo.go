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

func GetTodos(c *gin.Context) {
	todos := []models.Todo{}
	if err := db.DBCon.Find(&todos).Error; err != nil {
		c.JSON(http.StatusNotFound, structs.Error{http.StatusNotFound, "getting error while processing db request"})
		return
	}

	c.JSON(http.StatusOK, todos)
}

func GetTodo(c *gin.Context) {
	ID, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, structs.Error{http.StatusBadRequest, "Please pass valid id"})
		return
	}

	todo := models.Todo{}
	if err := db.DBCon.First(&todo, ID).Error; err != nil {
		c.JSON(http.StatusNotFound, structs.Error{http.StatusNotFound, "record not found"})
		return
	}

	c.JSON(http.StatusOK, todo)
}

func CreateTodo(c *gin.Context) {

	var todo (models.Todo)
	if err := c.ShouldBindWith(&todo, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, structs.Error{http.StatusBadRequest, err.Error()})
		return
	}

	if err := db.DBCon.Create(&todo).Error; err != nil {
		c.JSON(http.StatusBadRequest, structs.Error{http.StatusBadRequest, "getting error while processing db request"})
		return
	}

	c.JSON(http.StatusOK, todo)
}

func UpdateTodo(c *gin.Context) {
	ID, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, structs.Error{http.StatusBadRequest, "Please pass valid id"})
		return
	}

	var todo (models.Todo)
	todo.ID = uint(ID)

	if err := db.DBCon.First(&todo).Error; err != nil {
		c.JSON(http.StatusBadRequest, structs.Error{http.StatusBadRequest, "no todo find with this id"})
		return
	}

	if err := c.ShouldBindWith(&todo, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, structs.Error{http.StatusBadRequest, err.Error()})
		return
	}

	if err := db.DBCon.Model(&todo).Updates(&todo).Error; err != nil {
		c.JSON(http.StatusBadRequest, structs.Error{http.StatusBadRequest, "getting error while processing db request"})
		return
	}

	c.JSON(http.StatusOK, todo)
}

func DeleteTodo(c *gin.Context) {
	ID, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, structs.Error{http.StatusBadRequest, "Please pass valid id"})
		return
	}

	var todo (models.Todo)
	todo.ID = uint(ID)
	if err := db.DBCon.Delete(&todo).Error; err != nil {
		fmt.Println("err", err)
		c.JSON(http.StatusBadRequest, structs.Error{http.StatusBadRequest, "getting error while processing db request"})
		return
	}

	c.JSON(http.StatusOK, structs.Success{http.StatusOK, "todo deleted successfully"})
}
