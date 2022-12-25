package controllers

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jianrong/cvwo-be/initializers"
	"github.com/jianrong/cvwo-be/models"
	"github.com/jianrong/cvwo-be/utils"
)

func FetchAllUsers(c *gin.Context) {
	// Get all records
	var users []models.User
	err := initializers.DB.Model(&models.User{}).Preload("Posts").Find(&users).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"users": users,
	})
}

func FetchOneUser(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid ID",
		})
		return
	}
	user, err := models.FindUserById(uint(id))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "User not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}

func Signup(c *gin.Context) {
	// get username/pass
	var body models.AuthenticationInput

	if c.ShouldBindJSON(&body) != nil {
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

	if c.ShouldBindJSON(&body) != nil {
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
		"ID": user.ID,
		"profilePicture": user.ProfilePicture,
	})
}

func RefreshToken(c *gin.Context) {
	token, user, err := utils.RefreshToken(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Unable to refresh token",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"token": token,
		"username": user.Username,
		"ID": user.ID,
		"profilePicture": user.ProfilePicture,
	})
}

func GetAllCommentsFromUser(c *gin.Context) {
	userId := c.Param("id")
	var comments []models.Comment

	err := initializers.DB.Where("user_id = ?", userId).Preload("User").Find(&comments).Error
  if err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }
  c.JSON(http.StatusOK, gin.H{"comments": comments})

}

func GetAllSelectedEntries(c *gin.Context) {
  var posts []models.Post
	var comments []models.Comment
	var postIds []int
	var commentIds []int

	var body models.SelectedEntries
	
	if c.ShouldBindJSON(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	for _, tag := range body.PostIds {
		postIds = append(postIds, int(tag))
	}

	for _, tag := range body.CommentIds {
		commentIds = append(commentIds, int(tag))
	}

	err := initializers.DB.Where("id IN ?", postIds).
	Where("user_id = ?", body.UserId).
	Order("created_at desc").
	Find(&posts).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	err = initializers.DB.Where("id IN ?", commentIds).
	Where("user_id = ?", body.UserId).
	Order("created_at desc").
	Find(&comments).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Print(comments)
	c.JSON(http.StatusOK, gin.H{"posts": posts, "comments": comments})
}

func UpdateUser(c *gin.Context) {
	var body models.ImageInput
	if c.ShouldBindJSON(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{ "error": "Failed to read body" })
		return
	}

	var user models.User

	err := initializers.DB.Model(&user).Where("id = ?", body.UserId).Update("profile_picture", body.ImageName).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{ "error": err })
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}

func UploadImageToS3(c *gin.Context) {
	form, _ := c.MultipartForm()
	images := form.File["image"]
	id := uuid.New()
	var imageName string

	for _, image := range images {
		f, err := image.Open()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err,
			})
			return
		}
		defer f.Close()
		size := image.Size
		buffer := make([]byte, size)
		f.Read(buffer)
    fileBytes := bytes.NewReader(buffer)
    fileType := http.DetectContentType(buffer)
    path := id.String() + image.Filename
		imageName = path

		awsClient := initializers.AwsClient()
		uploader := manager.NewUploader(awsClient)
		_, err = uploader.Upload(context.TODO(), &s3.PutObjectInput{
			Bucket:        aws.String("cvwo-user-profiles"),
      Key:           aws.String(path),
      Body:          fileBytes,
      ContentLength: size,
      ContentType:   aws.String(fileType),
    })
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err,
			})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"imageName": imageName,
	})
}