package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jianrong/cvwo-be/controllers"
	"github.com/jianrong/cvwo-be/middleware"
)


func Comments(route *gin.RouterGroup) {
	comments := route.Group("/comments")
	comments.GET("/:id", controllers.GetOneComment)
	
	comments.POST("/", middleware.RequireAuth(), controllers.CreateComment)

	comments.PUT("/:id", middleware.RequireAuth(), controllers.CommentUpdate)
	
	comments.DELETE("/:id", middleware.RequireAuth(), controllers.CommentDelete)
}