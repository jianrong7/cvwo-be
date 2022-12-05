package main

import (
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

	routes.Posts(r)
	routes.Users(r)

  r.Run()
}