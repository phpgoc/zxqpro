package routes

import (
	"net/http"

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
// @Success 200 {object} response.response.CommonResponseWithoutData "成功响应"
// @Router /task/create_top [post]
func TaskCreateTop(c *gin.Context) {
	var req request.TaskCreateTop
	if err := c.ShouldBindJSON(&req); err != nil {
		return
	}
	userId := service.GetUserIdFromAuthMiddleware(c)
	if !service.CanCreateTop(userId, req.ProjectId) {
		c.JSON(http.StatusOK, response.CreateResponseWithoutData(http.StatusForbidden, "无权限"))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponseWithoutData(200, "ok"))
}
