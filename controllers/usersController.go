package controllers

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/jianrong/cvwo-be/initializers"
	"github.com/jianrong/cvwo-be/models"
	"golang.org/x/crypto/bcrypt"
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
	// get email/pass 
	var body struct {
		Username string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		
		return
	}
	// hash password
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to hash password",
		})

		return
	}
	// create user
	user := models.User{Username: body.Username, Password: string(hash)}
	result := initializers.DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create user",
		})
		
		return
	}
	// respond
	c.JSON(http.StatusOK, gin.H{})
}

func Login(c *gin.Context) {
	// Get the email and pass off req body
	var body struct {
		Username string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		
		return
	}
	// Look up requested user
	var user models.User
	initializers.DB.First(&user, "username = ?", body.Username)
	// fmt.Println(user)
	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email or password",
		})

		return
	}

	// Compare sent in pass with saved user pass hash
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email or password",
		})

		return
	}
	// Generate a jwt token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		// expire token after 20 minutes
		"exp": time.Now().Add(time.Minute * 20).Unix(),
	})
	
	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Unable to create token",
		})

		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	// set cookie for 20 minutes
	c.SetCookie("Authorization", tokenString, 20 * 60, "", "", false, true)
	// send it back
	c.JSON(http.StatusOK, gin.H{
		"token": tokenString,
		"claims": token.Claims,
		"username": user.Username,
	})
}

func Validate(c *gin.Context) {
	userData, _ := c.Get("user")
	user := userData.(models.User)

	tokenData, _ := c.Get("token")
	token := tokenData.(*jwt.Token)

	tokenString, _ := c.Cookie("Authorization")

	c.JSON(http.StatusOK, gin.H{
		"token": tokenString,
		"claims": token.Claims,
		"username": user.Username,
	})
}