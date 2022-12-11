package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jianrong/cvwo-be/utils"
)

func RequireAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
			err := utils.ValidateJWT(c)
			if err != nil {
					c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication required"})
					c.Abort()
					return
			}
			c.Next()
	}
}
