package middleware

import (
	"net/http"
	"time"

	"github.com/phpgoc/zxqpro/my_runtime"

	"github.com/gin-gonic/gin"
	"github.com/phpgoc/zxqpro/interfaces"
	"github.com/phpgoc/zxqpro/pro_types"
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
		cookie, err := c.Request.Cookie(my_runtime.CookieName)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code":    UnauthorizedStatus,
				"message": UnauthorizedMsg + "1",
			})
			c.Abort()
			return
		}
		cookieValue := cookie.Value
		var cookieData pro_types.Cookie
		has := interfaces.Cache.Get(cookieValue, &cookieData)
		if !has {
			c.JSON(http.StatusOK, gin.H{
				"code":    UnauthorizedStatus,
				"message": UnauthorizedMsg + "2",
			})
			c.Abort()
			return
		}

		if !cookieData.LongLogin {
			has := interfaces.Cache.GetAndRefresh(cookieValue, &cookieData, 30*time.Minute)
			if !has {
				c.JSON(InternalErrorStatus, gin.H{
					"code":    InternalErrorStatus,
					"message": InternalErrorMsg + ": Failed to refresh cookie",
				})
				c.Abort()
				return
			}
		}
		c.Set(my_runtime.UserIDInContextKey, cookieData.ID)
		// 验证通过，继续处理请求
		c.Next()
	}
}
