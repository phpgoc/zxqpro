package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/phpgoc/zxqpro/interfaces"
	"github.com/phpgoc/zxqpro/pro_types"
	"github.com/phpgoc/zxqpro/utils"
)

const (
	UnauthorizedStatus  = 401
	UnauthorizedMsg     = "Unauthorized"
	InternalErrorStatus = 500
	InternalErrorMsg    = "Internal Server Error"
)

func AuthLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取当前路由，不验证一些路由
		if c.Request.URL.Path == "/api/user/login" {
			c.Next()
			return
		}
		cookie, err := c.Request.Cookie(utils.CookieName)
		if err != nil {
			c.JSON(UnauthorizedStatus, gin.H{
				"code":    UnauthorizedStatus,
				"message": UnauthorizedMsg,
			})
			c.Abort()
			return
		}
		cookieValue := cookie.Value
		cookieStruct, has := interfaces.Cache.Get(cookieValue)
		if !has {
			c.JSON(UnauthorizedStatus, gin.H{
				"code":    UnauthorizedStatus,
				"message": UnauthorizedMsg,
			})
			c.Abort()
			return
		}

		// 类型断言并处理错误
		cookieData, ok := cookieStruct.(pro_types.Cookie)
		if !ok {
			c.JSON(InternalErrorStatus, gin.H{
				"code":    InternalErrorStatus,
				"message": InternalErrorMsg + ": Invalid cookie data",
			})
			c.Abort()
			return
		}

		if !cookieData.UseMobile {
			_, has := interfaces.Cache.GetAndRefresh(cookieValue, 30*time.Minute)
			if !has {
				c.JSON(InternalErrorStatus, gin.H{
					"code":    InternalErrorStatus,
					"message": InternalErrorMsg + ": Failed to refresh cookie",
				})
				c.Abort()
				return
			}
		}

		// 验证通过，继续处理请求
		c.Next()
	}
}
