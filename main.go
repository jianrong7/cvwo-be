package main

import (
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jianrong/cvwo-be/initializers"
	"github.com/jianrong/cvwo-be/models"
	"github.com/jianrong/cvwo-be/routes"
)

// func init() {
// 	initializers.LoadEnvVariables()
// 	initializers.ConnectToDB()
// }

// func main() {
// 	gin.ForceConsoleColor()
// 	r := gin.New()
// 	r.Use(gin.Recovery())
// 	r.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
// 		// your custom format
// 		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
// 				param.ClientIP,
// 				param.TimeStamp.Format(time.RFC1123),
// 				param.Method,
// 				param.Path,
// 				param.Request.Proto,
// 				param.StatusCode,
// 				param.Latency,
// 				param.Request.UserAgent(),
// 				param.ErrorMessage,
// 		)
// 	}))

//   r.Run()
// }

func main() {
	initializers.LoadEnv()
	loadAndMigrateDB()
	serveApplication()
}


func loadAndMigrateDB() {
	initializers.ConnectToDB()
	// initializers.DB.Migrator().DropTable(&models.User{})
	// initializers.DB.Migrator().DropTable(&models.Post{})
	initializers.DB.AutoMigrate(&models.User{})
	initializers.DB.AutoMigrate(&models.Post{})
}

func serveApplication() {
	r := gin.Default()

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowCredentials = true
	// config.AddAllowHeaders("Access-")
	// config.AllowOrigins = []string{"http://localhost:8080", "http://localhost:3000"}
	// config.AllowCredentials = true
	config.AllowHeaders = []string{"Access-Control-Allow-Headers", "Access-Control-Allow-Origin", "Origin", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With"}
	r.Use(cors.New(config))
	
	v1 := r.Group("/api/v1")
	{
		routes.Posts(v1)
		routes.Users(v1)
	}

	r.Run()
	fmt.Println("Server running on port 3000")
}