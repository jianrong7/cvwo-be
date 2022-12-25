package controllers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jianrong/cvwo-be/initializers"
	"github.com/jianrong/cvwo-be/models"
	"github.com/jianrong/cvwo-be/utils"
	gogpt "github.com/sashabaranov/go-gpt3"
	"gorm.io/gorm"
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
	
	body.UserID = user.ID
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

	tags := c.Query("tags")

	search := c.Query("search")

	// sort := c.Query("sort")
	// if sort == "downvotes" {
	// 	initializers.DB.
	// 	Where("UPPER(title) LIKE UPPER('%" + search + "%') OR UPPER(content) LIKE UPPER('%" + search + "%')").
	// 	Where(`tags @> '{` + tags + `}'`).
	// 	Preload("User").
	// 	Preload("Comments").
	// 	Preload("Upvotes", "entry_type = 'post' AND value = ?", 1).
	// 	Preload("Downvotes", "entry_type = 'post' AND value = ?", -1, func(db *gorm.DB) *gorm.DB {
	// 		return db.Group("ratings.id").Group("ratings.value").Order("COUNT(ratings.value) DESC")
	// 	}).
	// 	Find(&posts)
	// } else if sort == "upvotes" {
	// 	initializers.DB.
	// 	Where("UPPER(title) LIKE UPPER('%" + search + "%') OR UPPER(content) LIKE UPPER('%" + search + "%')").
	// 	Where(`tags @> '{` + tags + `}'`).
	// 	Preload("User").
	// 	Preload("Comments").
	// 	Preload("Upvotes", "entry_type = 'post' AND value = ?", 1, func(db *gorm.DB) *gorm.DB {
	// 		return db.Group("ratings.id").Group("ratings.value").Order("COUNT(value) DESC")
	// 	}).
	// 	Preload("Downvotes", "entry_type = 'post' AND value = ?", -1).
	// 	Order("downvotes").
	// 	Find(&posts)
	// } else {
		initializers.DB.
		Where("UPPER(title) LIKE UPPER('%" + search + "%') OR UPPER(content) LIKE UPPER('%" + search + "%')").
		Where(`tags @> '{` + tags + `}'`).
		Order("created_at desc").
		Preload("User").
		Preload("Comments").
		Preload("Upvotes", "entry_type = 'post' AND value = ?", 1).
		Preload("Downvotes", "entry_type = 'post' AND value = ?", -1).
		Find(&posts)
	// }

	// order := c.Query("order")
	// if order == "" {
	// 	order = "desc"
	// }

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
	var upvotes []models.Rating
	var downvotes []models.Rating
	
	initializers.DB.Preload("Upvotes", "entry_type = 'post' AND value = ?", 1).
	Preload("Downvotes", "entry_type = 'post' AND value = ?", -1).
	Preload("User").
	Preload("Comments", func(db *gorm.DB) *gorm.DB {
		return db.Order("comments.created_at DESC")
	}).
	First(&post, id)

	initializers.DB.Where(map[string]interface{}{"entry_id": id, "entry_type": "post", "value": 1}).Find(&upvotes)
	initializers.DB.Where(map[string]interface{}{"entry_id": id, "entry_type": "post", "value": -1}).Find(&downvotes)
	
	c.JSON(200, gin.H{
		"post": post,
		"upvotes": upvotes,
		"downvotes": downvotes,
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

	initializers.DB.Model(&post).
	Updates(map[string]interface{}{"title": body.Title, "content": body.Content})
	
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsDelete(c *gin.Context) {
	id := c.Param("id")

	initializers.DB.Delete(&models.Post{}, id)

	c.Status(200)
}

func GetAllCommentsFromPost(c *gin.Context) {
	postId := c.Param("id")
  var comments []models.Comment

	err := initializers.DB.
	Order("created_at desc").
	Where("post_id = ?", postId).
	Preload("Upvotes", "entry_type = 'comment' AND value = ?", 1).
	Preload("Downvotes", "entry_type = 'comment' AND value = ?", -1).
	Preload("User").
	Find(&comments).Error

  if err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }
  c.JSON(http.StatusOK, gin.H{"comments": comments})
}

func CreatePostFromOpenAI(c *gin.Context) {
	var body models.AiInput
	fmt.Print(body)

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	gpt := initializers.OpenAiClient()
	ctx := context.Background()

	req := gogpt.CompletionRequest{
		Model: "text-davinci-003",
		MaxTokens: body.MaxTokens,
		Prompt:    body.Prompt,
		Temperature: 0.5,
	}

	resp, err := gpt.CreateCompletion(ctx, req)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": resp.Choices[0].Text})
}