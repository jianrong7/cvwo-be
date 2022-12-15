package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jianrong/cvwo-be/controllers"
	"github.com/jianrong/cvwo-be/middleware"
)


func Comments(route *gin.RouterGroup) {
	comments := route.Group("/comments")
	
	
	comments.POST("/", middleware.RequireAuth(), controllers.CreateComment)
	// posts.GET("/:id", controllers.GetOnePost)
	// posts.GET("/user", middleware.RequireAuth(), controllers.GetAllPostsFromUser)

	// posts.POST("/", middleware.RequireAuth(), controllers.CreatePost)

	// posts.PUT("/:id", middleware.RequireAuth(), controllers.PostsUpdate)
	
	comments.DELETE("/:id", middleware.RequireAuth(), controllers.CommentDelete)
}