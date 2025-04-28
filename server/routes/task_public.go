package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/phpgoc/zxqpro/routes/request"
	"github.com/phpgoc/zxqpro/routes/response"
	"github.com/phpgoc/zxqpro/utils"
)

// PublicInfo  godoc
// @Summary task
// @Schemes
// @Description task info
// @Tags Task
// @Accept */*
// @Produce json
// @Param CommonID query request.CommonID true "CommonID"
// @Success 200 {object} response.CommonResponse[data=response.TaskInfo] "成功响应"
// @Router /task/public_info [get]
func (h *TaskHandler) PublicInfo(c *gin.Context) {
	var req request.CommonID
	if success := utils.ValidateQuery(c, &req); !success {
		return
	}
	taskInfo, err := h.taskService.TaskInfo(req.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.CreateResponseWithoutData(1, err.Error()))
		return
	}
	c.JSON(http.StatusOK, response.CreateResponse(0, "ok", taskInfo))
}

// ProjectTaskList godoc
// @Summary project task_list
// @Description project task_list
// @Tags Project
// @Accept */*
// @Produce json
// @Param ProjectTaskList body request.ProjectTaskList true "ProjectTaskList"
// @Success 200 {object} response.CommonResponse[data=response.TaskList] "成功响应"
// @Router /project/task_list [post]
func (h *TaskHandler) ProjectTaskList(c *gin.Context) {
	var req request.ProjectTaskList
	if success := utils.ValidateJson(c, &req); !success {
		return
	}

	taskList, err := h.taskService.GetProjectTaskList(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.CreateResponse(1, err.Error(), nil))
		return
	}
	c.JSON(http.StatusOK, response.CreateResponse(0, "ok", taskList))
}
