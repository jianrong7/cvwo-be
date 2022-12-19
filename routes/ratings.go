package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jianrong/cvwo-be/controllers"
	"github.com/jianrong/cvwo-be/middleware"
)


func Ratings(route *gin.RouterGroup) {
	ratings := route.Group("/ratings")
	
	// using PUT method here guarantees indempotence. Create rating if it does not exist and also updates it if it exists.
	ratings.PUT("/", middleware.RequireAuth(), controllers.CreateRating)

	// posts.GET("/:id", controllers.GetOnePost)
	// posts.GET("/user", middleware.RequireAuth(), controllers.GetAllPostsFromUser)

	// posts.POST("/", middleware.RequireAuth(), controllers.CreatePost)

	// comments.PUT("/:id", middleware.RequireAuth(), controllers.CommentUpdate)
	
	// comments.DELETE("/:id", middleware.RequireAuth(), controllers.CommentDelete)
}