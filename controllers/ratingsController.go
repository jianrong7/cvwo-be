package controllers

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jianrong/cvwo-be/initializers"
	"github.com/jianrong/cvwo-be/models"
	"github.com/jianrong/cvwo-be/utils"
	"gorm.io/gorm"
)

func CreateRating(c *gin.Context) {
	// get req body
	var body models.Rating
	var rating models.Rating
	
	if c.ShouldBindJSON(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid body request",
		})
		return
	}
	// identify whether post or comment
	// determine current user from context
	user, err := utils.CurrentUser(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	
	body.UserID = user.ID

	// check if user has rated this post/comment before
	err = initializers.DB.Where("entry_id = ?", body.EntryID).Where("entry_type = ?", body.EntryType).Where("user_id = ?", user.ID).First(&rating).Error
	fmt.Print(err)
	
	if errors.Is(err, gorm.ErrRecordNotFound) {
		// if user has not rated this post before, create rating
		savedRating, err := body.Save()

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
	
		c.JSON(http.StatusCreated, gin.H{
			"rating": savedRating,
		})

	} else {
		// if user has rated this post before, update rating
		fmt.Println("SEEN BEFORE", rating)
		err = initializers.DB.Model(&rating).Update("value", body.Value).Error

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"rating": rating,
		})
	}
}
