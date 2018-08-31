package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github/utsav/gogin/services"
	"github/utsav/gogin/structs"
)

func GetTodos(c *gin.Context) {
	todos := []structs.Todo{}
	todos = append(todos, services.TodoData());
	c.JSON(http.StatusOK, todos);
}

func GetTodo(c *gin.Context) {
	c.JSON(http.StatusOK, services.TodoData());
}

func PostTodo(c *gin.Context) {
	var todo (structs.Todo)
	err := c.ShouldBindWith(&todo, binding.JSON)
	if err != nil {
		c.JSON(http.StatusBadRequest, structs.Error{http.StatusBadRequest, err.Error()} )
		return
	}

	c.JSON(http.StatusOK, services.TodoData());
}
