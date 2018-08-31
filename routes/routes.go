package routes

import "github.com/gin-gonic/gin"
import ctrl "github.com/utsav/gin-gonic-boilerplate/controllers"

func Run() {
	r := gin.Default()
	r.GET("/todos", ctrl.GetTodos)
	r.GET("/todo", ctrl.GetTodo)
	r.POST("/todo", ctrl.PostTodo)
	r.Run(":3030") // listen and serve on 0.0.0.0:8080
}
