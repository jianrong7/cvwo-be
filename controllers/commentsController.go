package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jianrong/cvwo-be/initializers"
	"github.com/jianrong/cvwo-be/models"
	"github.com/jianrong/cvwo-be/utils"
)

// func GetOneComment(c *gin.Context) {
// 	id := c.Param("id")

// 	// var post models.Post
// 	var comment models.Comment
// 	var upvotes []models.Rating
// 	var downvotes []models.Rating

// 	initializers.DB.Preload("Ratings", "entry_type = 'post'").Preload("User").Preload("Comments", func(db *gorm.DB) *gorm.DB {
// 		return db.Order("comments.created_at DESC")
// 	}).First(&post, id)

// 	initializers.DB.Where(map[string]interface{}{"entry_id": id, "entry_type": "post", "value": 1}).Find(&upvotes)
// 	initializers.DB.Where(map[string]interface{}{"entry_id": id, "entry_type": "post", "value": -1}).Find(&downvotes)

// 	c.JSON(200, gin.H{
// 		"comment": comment,
// 		"upvotes": upvotes,
// 		"downvotes": downvotes,
// 	})
// }

func CreateComment(c *gin.Context) {
	var body models.Comment

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
	savedComment, err := body.Save()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"comment": savedComment,
	})
}

func CommentUpdate(c *gin.Context) {
	id := c.Param("id")

	var body struct{
		Content string
	}
	
	c.Bind(&body)

	var comment models.Comment

	initializers.DB.Where("ID = ?", id).Model(&comment).Update("content", body.Content)
	
	c.JSON(200, gin.H{
		"comment": comment,
	})
}

func CommentDelete(c *gin.Context) {
	id := c.Param("id")

	initializers.DB.Delete(&models.Comment{}, id)

	c.Status(200)
}
