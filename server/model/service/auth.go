package service

import (
	"github.com/gin-gonic/gin"
	"github.com/phpgoc/zxqpro/my_runtime"
)

func IsAdmin(userID uint) bool {
	return userID == 1
}

func GetUserIDFromAuthMiddleware(c *gin.Context) uint {
	// 不需要处理异常，一定有
	userID := c.MustGet(my_runtime.UserIDInContextKey)
	return userID.(uint)
}
