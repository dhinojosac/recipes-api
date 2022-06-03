package middleware

import (
	"os"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware is a middleware to authenticate requests with API KEY
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.GetHeader("X-API-KEY") !=
			os.Getenv("X_API_KEY") {
			c.AbortWithStatus(401)
		}
		c.Next()
	}
}
