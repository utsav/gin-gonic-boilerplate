package routes

import "github.com/gin-gonic/gin"
import ctrl "github.com/utsav/gin-gonic-boilerplate/controllers"

func Run() {
	r := gin.Default()
	r.GET("/todos", ctrl.GetTodos)
	r.GET("/todo/:id", ctrl.GetTodo)
	r.POST("/todo", ctrl.CreateTodo)
	r.PUT("/todo/:id", ctrl.UpdateTodo)
	r.DELETE("/todo/:id", ctrl.DeleteTodo)
	r.Run(":3030") // listen and serve on 0.0.0.0:8080
}
