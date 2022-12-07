package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jianrong/cvwo-be/initializers"
	"github.com/jianrong/cvwo-be/routes"
)


func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()
	r.Use(cors.Default())
	v1 := r.Group("/api/v1")
	{
		routes.Posts(v1)
		routes.Users(v1)
	}

  r.Run()
}