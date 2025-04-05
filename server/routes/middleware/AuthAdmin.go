package middleware

import (
	"github.com/gin-gonic/gin"
)

func AuthAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := GetUserIdFromAuthMiddleware(c)
		if userId != 1 {
			c.JSON(403, gin.H{
				"code":    403,
				"message": "Forbidden",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
