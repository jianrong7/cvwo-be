package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jianrong/cvwo-be/initializers"
	"github.com/jianrong/cvwo-be/models"
	"github.com/jianrong/cvwo-be/utils"
)

func CreatePost(c *gin.Context) {
	var body models.Post

	if c.ShouldBindJSON(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid body request",
		})
		return
	}
	// determine current user from context
	user, err := utils.CurrentUser(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	
	body.UserId = user.ID
	savedPost, err := body.Save()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"post": savedPost,
	})
}

func GetAllPosts(c *gin.Context) {
  var posts []models.Post
	var err error
	sqlStatement := `SELECT * FROM "posts" WHERE "posts"."deleted_at" IS NULL `

	tags := c.Query("tags")
	fmt.Println(tags)
	if len(tags) != 0 {
		sqlStatement += `AND tags @> '{` + tags + `}' `
	}

	sort := c.Query("sort")
	if sort == "upvotes" {
		sqlStatement += `ORDER BY "upvotes" `
	} else if sort == "downvotes" {
		sqlStatement += `ORDER BY "downvotes" `
	} else {
		sqlStatement += `ORDER BY "created_at" `
	}

	order := c.Query("order")
	if order == "asc" {
		sqlStatement += `asc`
	} else {
		sqlStatement += `desc`
	}
	
	initializers.DB.Raw(sqlStatement).Scan(&posts)
  if err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }
  c.JSON(http.StatusOK, gin.H{"posts": posts})
}

func GetAllPostsFromUser(c *gin.Context) {
  user, err := utils.CurrentUser(c)

  if err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }
  c.JSON(http.StatusOK, gin.H{"posts": user.Posts})
}

func GetOnePost(c *gin.Context) {
	id := c.Param("id")

	var post models.Post
	initializers.DB.First(&post, id)
	
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsUpdate(c *gin.Context) {
	id := c.Param("id")

	var body struct{
		Title string
		Content string
	}
	
	c.Bind(&body)

	var post models.Post
	initializers.DB.First(&post, id)

	initializers.DB.Model(&post).Updates(models.Post{
		Title: body.Title,
		Content: body.Content,
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