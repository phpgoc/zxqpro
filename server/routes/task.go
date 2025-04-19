package routes

import (
	"net/http"

	"github.com/phpgoc/zxqpro/utils"

	"github.com/gin-gonic/gin"
	"github.com/phpgoc/zxqpro/model/service"
	"github.com/phpgoc/zxqpro/routes/request"
	"github.com/phpgoc/zxqpro/routes/response"
)

// TaskCreateTop  godoc
// @Summary task create top
// @Schemes
// @Description task create top
// @Tags Task
// @Accept json
// @Produce json
// @Param TaskCreateTop body request.TaskCreateTop true "AdminRegister"
// @Success 200 {object} response.CommonResponseWithoutData "成功响应"
// @Router /task/create_top [post]
func TaskCreateTop(c *gin.Context) {
	var req request.TaskCreateTop
	if success := utils.ValidateJson(c, &req); !success {
		return
	}
	userId := service.GetUserIdFromAuthMiddleware(c)

	err := service.TaskCreateTop(userId, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.CreateResponseWithoutData(1, err.Error()))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponseWithoutData(0, "ok"))
}
