package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jianrong/cvwo-be/controllers"
)


func Posts(route *gin.RouterGroup) {
	posts := route.Group("/posts")
	
	posts.POST("/", controllers.CreatePost)
	posts.GET("/", controllers.FetchAllPosts)
	posts.GET("/:id", controllers.FetchOnePost)
	posts.PUT("/:id", controllers.PostsUpdate)
	posts.DELETE("/:id", controllers.PostsDelete)
}