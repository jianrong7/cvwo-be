package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jianrong/cvwo-be/initializers"
	"github.com/jianrong/cvwo-be/middleware"
	"github.com/jianrong/cvwo-be/models"
	"github.com/jianrong/cvwo-be/routes"
)

func main() {
	initializers.LoadEnv()
	loadAndMigrateDB()
	serveApplication()
}


func loadAndMigrateDB() {
	initializers.ConnectToDB()
	initializers.DB.AutoMigrate(&models.Post{})
	initializers.DB.AutoMigrate(&models.Rating{})
	initializers.DB.AutoMigrate(&models.Comment{})
	initializers.DB.AutoMigrate(&models.User{})
}

func serveApplication() {
	r := gin.Default()
	r.Use(middleware.CORSMiddleware())

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	
	v1 := r.Group("/api/v1")
	{
		routes.Comments(v1)
		routes.Posts(v1)
		routes.Users(v1)
		routes.Ratings(v1)
	}

	r.Run()
	fmt.Println("Server running on port 3000")
}