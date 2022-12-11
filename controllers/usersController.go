package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jianrong/cvwo-be/initializers"
	"github.com/jianrong/cvwo-be/models"
	"github.com/jianrong/cvwo-be/utils"
)

func FetchAllUsers(c *gin.Context) {
	// Get all records
	var users []models.User
	initializers.DB.Model(&models.User{}).Preload("Posts").Find(&users)

	c.JSON(200, gin.H{
		"users": users,
	})
}

func FetchOneUser(c *gin.Context) {
	id := c.Param("id")

	var user models.User
	initializers.DB.First(&user, id)

	c.JSON(200, gin.H{
		"user": user,
	})
}

func Signup(c *gin.Context) {
	// get username/pass
	var body models.AuthenticationInput

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}
	user := models.User{
		Username: body.Username,
		Password: body.Password,
	}
	savedUser, err := user.Save()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create user",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": savedUser})
}

func Login(c *gin.Context) {
	// Get the email and pass off req body
	var body models.AuthenticationInput

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		
		return
	}
	// Look up requested user
	user, err := models.FindUserByUsername(body.Username)
	
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email or password",
		})

		return
	}
	// Compare sent in pass with saved user pass hash
	err = user.ValidatePassword(body.Password)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email or password",
		})

		return
	}

	
	// Generate a jwt token
	token, err := utils.GenerateJWT(user.ID)
	
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Unable to create token",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token, 
		"username": user.Username,
	})
	// c.SetSameSite(http.SameSiteLaxMode)
	// // set cookie for 20 minutes
	// c.SetCookie("Authorization", tokenString, 20 * 60, "", "", false, true)
	// // send it back
	// c.JSON(http.StatusOK, gin.H{
	// 	"token": tokenString,
	// 	"claims": token.Claims,
	// 	"username": user.Username,
	// })
}

func RefreshToken(c *gin.Context) {
	token, username, err := utils.RefreshToken(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Unable to refresh token",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"token": token,
		"username": username,
	})
}