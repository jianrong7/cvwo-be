package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jianrong/cvwo-be/controllers"
	"github.com/jianrong/cvwo-be/middleware"
)


func Users(route *gin.RouterGroup) {
	users := route.Group("/users")
	{
		users.GET("/:id", controllers.FetchOneUser)
		users.GET("/", controllers.FetchAllUsers)
		users.POST("/signup", controllers.Signup)
		users.POST("/login", controllers.Login)
		users.GET("/validate", middleware.RequireAuth, controllers.Validate)
		// posts.GET("/", controllers.PostsIndex)
		// posts.GET("/:id", controllers.PostsShow)
		// posts.PUT("/:id", controllers.PostsUpdate)
		// posts.DELETE("/:id", controllers.PostsDelete)
	}
}