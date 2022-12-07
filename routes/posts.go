package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jianrong/cvwo-be/controllers"
)


func Posts(route *gin.RouterGroup) {
	posts := route.Group("/posts")
	
	posts.POST("/", controllers.PostsCreate)
	posts.GET("/", controllers.PostsIndex)
	posts.GET("/:id", controllers.PostsShow)
	posts.PUT("/:id", controllers.PostsUpdate)
	posts.DELETE("/:id", controllers.PostsDelete)
}