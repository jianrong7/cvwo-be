package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jianrong/cvwo-be/controllers"
	"github.com/jianrong/cvwo-be/middleware"
)


func Posts(route *gin.RouterGroup) {
	posts := route.Group("/posts")
	
	posts.GET("/", controllers.GetAllPosts)
	posts.GET("/:id", controllers.GetOnePost)
	posts.GET("/user", controllers.GetAllPostsFromUser)
	posts.GET("/comments/:id", controllers.GetAllCommentsFromPost)

	posts.POST("/", middleware.RequireAuth(), controllers.CreatePost)
	posts.POST("/ai", middleware.RequireAuth(), controllers.CreatePostFromOpenAI)

	posts.PUT("/:id", middleware.RequireAuth(), controllers.PostsUpdate)
	
	posts.DELETE("/:id", middleware.RequireAuth(), controllers.PostsDelete)
}