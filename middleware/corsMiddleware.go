package middleware

import "github.com/gin-gonic/gin"

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

			c.Header("Access-Control-Allow-Origin", "https://cvwo-fe.s3.ap-southeast-1.amazonaws.com")
			c.Header("Access-Control-Allow-Origin", "http://localhost:8080")
			c.Header("Access-Control-Allow-Credentials", "true")
			c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
			c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, HEAD, POST, PUT, DELETE, OPTIONS, PATCH")

			if c.Request.Method == "OPTIONS" {
					c.IndentedJSON(204, "")
					return
			}

			c.Next()
	}
}