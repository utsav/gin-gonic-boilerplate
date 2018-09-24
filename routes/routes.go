package routes

import (
	"github.com/gin-gonic/gin"
	ctrl "github.com/utsav/gin-gonic-boilerplate/controllers"
)

func Run() {
	r := gin.Default()

	r.Use(func(c *gin.Context) {
		// add header Access-Control-Allow-Origin
		c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Add("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Add("Access-Control-Allow-Methods", "GET,HEAD,OPTIONS,POST,PUT,DELETE")
		c.Writer.Header().Add("Access-Control-Allow-Headers", "Access-Control-Allow-Headers, Origin,Accept, X-Requested-With, Content-Type, Access-Control-Request-Method, Access-Control-Request-Headers")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	})

	r.GET("/todos", ctrl.GetTodos)
	r.GET("/todo/:id", ctrl.GetTodo)
	r.POST("/todo", ctrl.CreateTodo)
	r.PUT("/todo/:id", ctrl.UpdateTodo)
	r.DELETE("/todo/:id", ctrl.DeleteTodo)
	r.Run(":3030") // listen and serve on 0.0.0.0:8080
}
