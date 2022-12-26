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
		users.POST("/refresh", controllers.RefreshToken)
		users.GET("/comments/:id", controllers.GetAllCommentsFromUser)
		users.POST("/selected", controllers.GetAllSelectedEntries)
		users.POST("/:id", middleware.RequireAuth(), controllers.UploadImageToS3)
		users.PUT("/:id", middleware.RequireAuth(), controllers.UpdateUser)
	}
}