package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jianrong/cvwo-be/initializers"
	"github.com/jianrong/cvwo-be/models"
	"github.com/jianrong/cvwo-be/utils"
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

	sort := c.Query("sort")
	if sort == "" {
		sort = "created_at"
	}

	order := c.Query("order")
	if order == "" {
		order = "desc"
	}
	
	initializers.DB.Where("UPPER(title) LIKE UPPER('%" + search + "%') OR UPPER(content) LIKE UPPER('%" + search + "%')").Where(`tags @> '{` + tags + `}'`).Order(sort + " " + order).Preload("User").Preload("Comments").Find(&posts)
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
	
	initializers.DB.Preload("User").Preload("Comments", func(db *gorm.DB) *gorm.DB {
		return db.Order("comments.created_at DESC")
	}).First(&post, id)
	
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

func GetAllCommentsFromPost(c *gin.Context) {
	postId := c.Param("id")
  var comments []models.Comment

	err := initializers.DB.Order("created_at desc").Where("post_id = ?", postId).Preload("User").Find(&comments).Error
  if err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }
  c.JSON(http.StatusOK, gin.H{"comments": comments})
}