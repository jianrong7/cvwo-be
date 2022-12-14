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

	posts.POST("/", middleware.RequireAuth(), controllers.CreatePost)

	posts.PUT("/:id", middleware.RequireAuth(), controllers.PostsUpdate)
	
	posts.DELETE("/:id", middleware.RequireAuth(), controllers.PostsDelete)
	posts.GET("/comments/:id", controllers.GetAllCommentsFromPost)
}