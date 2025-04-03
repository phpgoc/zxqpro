package middle_ware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// 定义验证器实例
var validate = validator.New()

// ValidationMiddleware 验证中间件
func ValidationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 假设请求体是一个 JSON 数据，绑定到结构体
		var requestData interface{}
		// 这里可以根据实际情况修改绑定的结构体类型
		// 例如：var requestData YourRequestStruct
		if err := c.ShouldBindJSON(&requestData); err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code":    400,
				"message": err.Error(),
			})
			c.Abort()
			return
		}

		// 验证请求数据
		if err := validate.Struct(requestData); err != nil {
			// 验证失败，返回错误信息
			if _, ok := err.(*validator.InvalidValidationError); ok {
				c.JSON(http.StatusInternalServerError, gin.H{
					"code":    400,
					"message": err.Error(),
				})
			} else {
				var validationErrors []string
				for _, err := range err.(validator.ValidationErrors) {
					validationErrors = append(validationErrors, err.Error())
				}
				c.JSON(http.StatusBadRequest, gin.H{
					"code":    400,
					"message": validationErrors,
				})
			}
			c.Abort()
			return
		}

		// 验证通过，继续处理请求
		c.Next()
	}
}
