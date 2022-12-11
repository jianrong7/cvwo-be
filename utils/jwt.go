package utils

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/jianrong/cvwo-be/models"
)

var privateKey = []byte(os.Getenv("JWT_SECRET"))
// generate JWT when logging in
func GenerateJWT(userId uint) (string, error) {
		// Generate a jwt token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": userId,
		
		"iat": time.Now().Unix(),
		// expire token after 20 minutes
		"exp": time.Now().Add(time.Minute * 20).Unix(),
	})
	return token.SignedString(privateKey)
}

func CurrentUser(c *gin.Context) (models.User, error) {
	// validate JWT
	err := ValidateJWT(c)
	if err != nil {
		return models.User{}, err
	}
	// then get the user
	token, _ := getToken(c)
	claims, _ := token.Claims.(jwt.MapClaims)
	userId := uint(claims["sub"].(float64))

	user, err := models.FindUserById(userId)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func ValidateJWT(c *gin.Context) error {
	// get token first
	token, err := getToken(c)
	if err != nil {
		return err
	}

	// then validate that its claims are valid
	_, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		return nil
	}
	return errors.New("invalid token provided")
}

func getToken(c *gin.Context) (*jwt.Token, error) {
	tokenString := getTokenFromRequest(c)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return privateKey, nil
	})
	return token, err
}

func getTokenFromRequest(c *gin.Context) string {
	bearerToken := c.Request.Header.Get("Authorization")
	splitToken := strings.Split(bearerToken, " ")
	if len(splitToken) == 2 {
			return splitToken[1]
	}
	return ""
}

func RefreshToken(c *gin.Context) (string, error) {
		// get stale token first
		staleToken, err := getToken(c)
		if err != nil {
			return "", err
		}
		claims, _ := staleToken.Claims.(jwt.MapClaims)
		userId := uint(claims["sub"].(float64))
		// Generate a jwt token
		token, err := GenerateJWT(userId)
			
		if err != nil {
			return "", err
		}
		return token, nil
}