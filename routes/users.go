package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jianrong/cvwo-be/controllers"
)


func Users(route *gin.RouterGroup) {
	users := route.Group("/users")
	{
		users.GET("/:id", controllers.FetchOneUser)
		users.GET("/", controllers.FetchAllUsers)
		users.POST("/signup", controllers.Signup)
		users.POST("/login", controllers.Login)
		users.GET("/refresh", controllers.RefreshToken)
	}
}