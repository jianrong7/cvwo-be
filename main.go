package main

import (
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/logger"
	"github.com/gin-gonic/gin"
	"github.com/jianrong/cvwo-be/initializers"
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

func CORSConfig() cors.Config {
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{
		"http://cvwo-fe.s3-website-ap-southeast-1.amazonaws.com/",
		"https://cvwo-fe.s3.ap-southeast-1.amazonaws.com",
		"https://d3mj3t330xelda.cloudfront.net",
		"http://localhost:8080",
	}
	corsConfig.AllowCredentials = true
	corsConfig.AddAllowHeaders("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers", "Content-Type", "X-XSRF-TOKEN", "Accept", "Origin", "X-Requested-With", "Authorization")
	corsConfig.AddAllowMethods("GET, HEAD, POST, PUT, DELETE, OPTIONS, PATCH")
	return corsConfig
}

func serveApplication() {
	r := gin.Default()
	// r.Use(middleware.CORSMiddleware())
	r.Use(logger.SetLogger())
	r.Use(cors.New(CORSConfig()))


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