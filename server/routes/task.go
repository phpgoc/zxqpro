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
	userID := service.GetUserIDFromAuthMiddleware(c)

	err := service.TaskCreateTop(userID, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.CreateResponseWithoutData(1, err.Error()))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponseWithoutData(0, "ok"))
}

// TaskUpdateTop  godoc
// @Summary task update top
// @Schemes
// @Description task update top
// @Tags Task
// @Accept json
// @Produce json
// @Param TaskUpdateTop body request.TaskUpdateTop true "AdminRegister"
// @Success 200 {object} response.CommonResponseWithoutData "成功响应"
// @Router /task/update_top [post]
func TaskUpdateTop(c *gin.Context) {
	var req request.TaskUpdateTop
	if success := utils.ValidateJson(c, &req); !success {
		return
	}
	userID := service.GetUserIDFromAuthMiddleware(c)
	err := service.TaskUpdateTop(userID, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.CreateResponseWithoutData(1, err.Error()))
		return
	}
	c.JSON(http.StatusOK, response.CreateResponseWithoutData(0, "ok"))
}

// TaskAssignTop  godoc
// @Summary task assign top
// @Schemes
// @Description task assign top
// @Tags Task
// @Accept json
// @Produce json
// @Param TaskAssignTop body request.CommonID true "AdminRegister"
// @Success 200 {object} response.CommonResponseWithoutData "成功响应"
// @Router /task/assign_top [post]
func TaskAssignTop(c *gin.Context) {
	var req request.CommonID
	if success := utils.ValidateJson(c, &req); !success {
		return
	}
	userID := service.GetUserIDFromAuthMiddleware(c)
	err := service.TaskAssignSelfToTop(userID, req.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.CreateResponseWithoutData(1, err.Error()))
		return
	}
	c.JSON(http.StatusOK, response.CreateResponseWithoutData(0, "ok"))
}
