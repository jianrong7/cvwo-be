package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jianrong/cvwo-be/initializers"
	"github.com/jianrong/cvwo-be/models"
	"github.com/jianrong/cvwo-be/utils"
)

func GetOneComment(c *gin.Context) {
	id := c.Param("id")
	var comment models.Comment

	err := initializers.DB.First(&comment, id).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"comment": comment,
	})
}

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
	
	if c.ShouldBindJSON(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to read body",
		})
		return
	}

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
