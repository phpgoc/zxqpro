package routes

import (
	"net/http"
	"time"

	"github.com/phpgoc/zxqpro/model/service"

	"github.com/phpgoc/zxqpro/interfaces"
	"github.com/phpgoc/zxqpro/my_runtime"

	"github.com/gin-gonic/gin"
	"github.com/phpgoc/zxqpro/model/dao"
	"github.com/phpgoc/zxqpro/model/entity"
	"github.com/phpgoc/zxqpro/routes/request"
	"github.com/phpgoc/zxqpro/routes/response"
	"github.com/phpgoc/zxqpro/utils"
)

// @BasePath /api

// ProjectCreateRole  godoc
// @Summary project create role
// @Schemes
// @Description project create role
// @Tags Project
// @Accept json
// @Produce json
// @Param ProjectUpsertRole body request.ProjectUpsertRole true "AdminRegister"
// @Success 200 {object} response.CommonResponseWithoutData "成功响应"
// @Router /project/create_role [post]
func ProjectCreateRole(c *gin.Context) {
	var req request.ProjectUpsertRole
	if success := utils.ValidateJson(c, &req); !success {
		return
	}
	if !service.HasOwnPermission(c, req.ProjectID) {
		return
	}
	if err := dao.CreateRole(req.UserID, req.ProjectID, req.RoleType); err != nil {
		c.JSON(http.StatusOK, response.CreateResponseWithoutData(1, err.Error()))
		return
	}
	c.JSON(http.StatusOK, response.CreateResponseWithoutData(0, "ok"))
}

// ProjectDeleteRole  godoc
// @Summary project delete role
// @Schemes
// @Description project delete role
// @Tags Project
// @Accept json
// @Produce json
// @Param ProjectDeleteRole body request.ProjectDeleteRole true "ProjectDeleteRole"
// @Success 200 {object} response.CommonResponseWithoutData "成功响应"
// @Router /project/delete_role [post]
func ProjectDeleteRole(c *gin.Context) {
	var req request.ProjectDeleteRole
	if success := utils.ValidateJson(c, &req); !success {
		return
	}
	if !service.HasOwnPermission(c, req.ProjectID) {
		return
	}
	if err := dao.DeleteRole(req.UserID, req.ProjectID); err != nil {
		c.JSON(http.StatusOK, response.CreateResponseWithoutData(1, err.Error()))
		return
	}
	c.JSON(http.StatusOK, response.CreateResponseWithoutData(0, "ok"))
}

// ProjectUpdateRole  godoc
// @Summary project update role
// @Schemes
// @Description project update role
// @Tags Project
// @Accept json
// @Produce json
// @Param ProjectUpsertRole body request.ProjectUpsertRole true "AdminRegister"
// @Success 200 {object} response.CommonResponseWithoutData "成功响应"
// @Router /project/update_role [post]
func ProjectUpdateRole(c *gin.Context) {
	var req request.ProjectUpsertRole
	if success := utils.ValidateJson(c, &req); !success {
		return
	}
	if !service.HasOwnPermission(c, req.ProjectID) {
		return
	}
	if err := dao.UpdateRole(req.UserID, req.ProjectID, req.RoleType); err != nil {
		c.JSON(http.StatusOK, response.CreateResponseWithoutData(1, err.Error()))
		return
	}
	c.JSON(http.StatusOK, response.CreateResponseWithoutData(0, "ok"))
}

// ProjectList godoc
// @Summary project list
// @Schemes
// @Description project list
// @Tags Project
// @Accept */*
// @Produce json
// @Param ProjectList query request.ProjectList true "ProjectList"
// @Success 200 {object} response.ProjectList "成功响应"
// @Router /project/list [get]
func ProjectList(c *gin.Context) {
	var req request.ProjectList
	if success := utils.ValidateQuery(c, &req); !success {
		return
	}

	userID := service.GetUserIDFromAuthMiddleware(c)

	responseProjectList, err := service.GetProjectList(userID, req.Status, req.RoleType, req.Page, req.PageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.CreateResponse(1, "error", nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(0, "ok", responseProjectList))
}

// ProjectUpdate  godoc
// @Summary project update
// @Schemes
// @Description project update
// @Tags Project
// @Accept json
// @Produce json
// @Param ProjectUpdate body request.ProjectUpdate true "ProjectUpdate"
// @Success 200 {object} response.CommonResponseWithoutData "成功响应"
// @Router /project/update [post]
func ProjectUpdate(c *gin.Context) {
	var req request.ProjectUpdate
	if success := utils.ValidateJson(c, &req); !success {
		return
	}
	if !service.HasOwnPermission(c, req.ID) {
		return
	}
	project := entity.Project{
		Name:        req.Name,
		Description: req.Description,
		GitAddress:  req.GitAddress,
		Config:      req.Config,
	}
	if err := dao.UpdateProject(req.ID, project); err != nil {
		c.JSON(http.StatusOK, response.CreateResponseWithoutData(1, err.Error()))
		return
	}
	c.JSON(http.StatusOK, response.CreateResponseWithoutData(0, "ok"))
}

// ProjectUpdateStatus  godoc
// @Summary project update status
// @Schemes
// @Description project update status
// @Tags Project
// @Accept json
// @Produce json
// @Param ProjectUpdateStatus body request.ProjectUpdateStatus true "ProjectUpdateStatus"
// @Success 200 {object} response.CommonResponseWithoutData "成功响应"
// @Router /project/update_status [post]
func ProjectUpdateStatus(c *gin.Context) {
	var req request.ProjectUpdateStatus
	if success := utils.ValidateJson(c, &req); !success {
		return
	}
	if !service.HasOwnPermission(c, req.ID) {
		return
	}
	if err := dao.UpdateProjectStatus(req.ID, req.Status); err != nil {
		c.JSON(http.StatusOK, response.CreateResponseWithoutData(1, err.Error()))
		return
	}
	c.JSON(http.StatusOK, response.CreateResponseWithoutData(0, "ok"))
}

// ProjectInfo godoc
// @Summary project info
// @Schemes
// @Description project info
// @Tags Project
// @Accept */*
// @Produce json
// @Param CommonID query request.CommonID true "CommonID"
// @Success 200 {object} response.CommonResponse[data=response.ProjectInfo] "成功响应"
// @Router /project/info [get]
func ProjectInfo(c *gin.Context) {
	var req request.CommonID
	if success := utils.ValidateQuery(c, &req); !success {
		return
	}
	project, err := dao.GetProjectByID(req.ID)
	if err != nil {
		c.JSON(http.StatusOK, response.CreateResponseWithoutData(1, err.Error()))
		return
	}
	projectInfo := response.ProjectInfo{
		ID:          project.ID,
		Name:        project.Name,
		OwnerID:     project.OwnerID,
		OwnerName:   project.Owner.UserName,
		Description: project.Description,
		GitAddress:  project.GitAddress,
		Config:      project.Config,
		Status:      project.Status,
	}
	c.JSON(http.StatusOK, response.CreateResponse(0, "ok", projectInfo))
}

// ProjectRoleIn godoc
// @Summary project role in
// @Schemes
// @Description project role
// @Tags Project
// @Accept */*
// @Produce json
// @Param CommonID query request.CommonID true "CommonID"
// @Success 200 {object} response.CommonResponse[data=response.ProjectRole] "成功响应"
// @Router /project/role_in [get]
func ProjectRoleIn(c *gin.Context) {
	var req request.CommonID
	if success := utils.ValidateQuery(c, &req); !success {
		return
	}
	userID := service.GetUserIDFromAuthMiddleware(c)
	key := utils.JoinCacheKey(my_runtime.PrefixUseridProjectRole, userID, req.ID)
	roleType := interfaces.GetOrSet(interfaces.Cache, key, func() entity.RoleType {
		roleType, _ := service.GetRoleType(userID, req.ID)
		return roleType
	}, time.Hour)

	c.JSON(http.StatusOK, response.CreateResponse(0, "ok", response.ProjectRole{
		RoleType: roleType,
	}))
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
func ProjectTaskList(c *gin.Context) {
	var req request.ProjectTaskList
	if success := utils.ValidateJson(c, &req); !success {
		return
	}

	taskList, err := service.GetTaskList(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.CreateResponse(1, "error", nil))
		return
	}
	c.JSON(http.StatusOK, response.CreateResponse(0, "ok", taskList))
}
