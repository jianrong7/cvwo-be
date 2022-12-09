package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/jianrong/cvwo-be/initializers"
	"github.com/jianrong/cvwo-be/models"
)

// gorm.Model
// Title string
// Body string
// // Tags pq.StringArray `gorm:type:text[]"`
// UserId uint
// User User `gorm:"foreignKey:UserId"`
// Upvotes uint
// Downvotes uint
// Comments []Comment
// }

func CreatePost(c *gin.Context) {
	var body struct{
		Title string
		Body string
		UserId uint
	}
	
	c.Bind(&body)

	// Create a post
	postData := models.Post{Title: body.Title, Body: body.Body, UserId: body.UserId, Upvotes: 0, Downvotes: 0}

	result := initializers.DB.Create(&postData)

	if result.Error != nil {
		c.Status(400)
		return
	}

	var post models.Post
	initializers.DB.Joins("User").Find(&post)
	
	c.JSON(200, gin.H{
		"post": post,
	})
}

func FetchAllPosts(c *gin.Context) {
	var posts []models.Post
	// initializers.DB.Find(&posts)
	initializers.DB.Model(&models.Post{}).Preload("User").Find(&posts)


	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func FetchOnePost(c *gin.Context) {
	id := c.Param("id")

	var post models.Post
	initializers.DB.Model(&models.Post{}).Preload("User").First(&post, id)
	
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsUpdate(c *gin.Context) {
	id := c.Param("id")

	var body struct{
		Title string
		Body string
	}
	
	c.Bind(&body)

	var post models.Post
	initializers.DB.First(&post, id)

	initializers.DB.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body: body.Body,
	})
	
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsDelete(c *gin.Context) {
	id := c.Param("id")

	initializers.DB.Delete(&models.Post{}, id)

	c.Status(200)
}