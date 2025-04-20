package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/phpgoc/zxqpro/model/service"
)

func AuthAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := service.GetUserIDFromAuthMiddleware(c)
		if userID != 1 {
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
