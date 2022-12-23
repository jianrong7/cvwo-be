package main

import (
	"fmt"

	"github.com/gin-contrib/cors"
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
	// initializers.DB.Migrator().DropTable(&models.User{})
	// initializers.DB.Migrator().DropTable(&models.Post{})
	// initializers.DB.Migrator().DropTable(&models.Rating{})
	// initializers.DB.Migrator().DropTable(&models.Comment{})
	initializers.DB.AutoMigrate(&models.Post{})
	initializers.DB.AutoMigrate(&models.Rating{})
	initializers.DB.AutoMigrate(&models.Comment{})
	initializers.DB.AutoMigrate(&models.User{})
}

func serveApplication() {
	r := gin.Default()
	r.Use(middleware.CORSMiddleware())

	// config := cors.DefaultConfig()
	// // config.AllowAllOrigins = true
	// config.AllowCredentials = true
	// // config.AddAllowHeaders("Access-")
	// config.AllowOrigins = []string{"http://localhost:8080", "http://localhost:3000", "http://cvwo-fe.s3.ap-southeast-1.amazonaws.com"}
	// // config.AllowCredentials = true
	// config.AllowHeaders = []string{"Access-Control-Allow-Headers", "Access-Control-Allow-Origin", "Origin", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With"}
	// r.Use(cors.New(config))
	r.Use(cors.Default())

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "It is working",
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