package routes

import (
	"github.com/gin-gonic/gin"
	ctrl "github.com/utsav/gin-gonic-boilerplate/controllers"
)

//Router ...
func Router(r *gin.Engine) {

	//user routes
	r.GET("/users", ctrl.GetUsers)
	r.GET("/user/:id", ctrl.GetUser)
	r.POST("/user", ctrl.CreateUser)
	r.PUT("/user/:id", ctrl.UpdateUser)
	r.DELETE("/user/:id", ctrl.DeleteUser)

	//blog routes
	r.GET("/blogs", ctrl.GetBlogs)
	r.GET("/blog/:id", ctrl.GetBlog)
	r.POST("/blog", ctrl.CreateBlog)
	r.PUT("/blog/:id", ctrl.UpdateBlog)
	r.DELETE("/blog/:id", ctrl.DeleteBlog)

	//comment routes
	r.POST("/comment", ctrl.CreateComment)
}
