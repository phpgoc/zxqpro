package service

import (
	"github.com/gin-gonic/gin"
	"github.com/phpgoc/zxqpro/my_runtime"
)

func IsAdmin(userId uint) bool {
	return userId == 1
}

func GetUserIdFromAuthMiddleware(c *gin.Context) uint {
	// 不需要处理异常，一定有
	userId, _ := c.Get(my_runtime.UserIdInContextKey)
	return userId.(uint)
}
