package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/phpgoc/zxqpro/model/service"
	"github.com/phpgoc/zxqpro/routes/request"
	"github.com/phpgoc/zxqpro/routes/response"
	"github.com/phpgoc/zxqpro/utils"
)

// TaskPublicInfo  godoc
// @Summary task
// @Schemes
// @Description task info
// @Tags Task
// @Accept */*
// @Produce json
// @Param CommonID query request.CommonID true "CommonID"
// @Success 200 {object} response.CommonResponse[data=response.TaskInfo] "成功响应"
// @Router /task/public_info [get]
func TaskPublicInfo(c *gin.Context) {
	var req request.CommonID
	if success := utils.ValidateQuery(c, &req); !success {
		return
	}
	taskInfo, err := service.TaskInfo(req.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.CreateResponseWithoutData(1, err.Error()))
		return
	}
	c.JSON(http.StatusOK, response.CreateResponse(0, "ok", taskInfo))
}
