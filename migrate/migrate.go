package main

import (
	"github.com/jianrong/cvwo-be/initializers"
	"github.com/jianrong/cvwo-be/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
  // Migrate the schema
	// initializers.DB.Model(&models.Post).r
  initializers.DB.AutoMigrate(&models.Post{}, &models.User{}, &models.Comment{})
	// initializers.DB.Model(&models.Post{})
}