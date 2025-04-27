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

type ProjectHandler struct {
	projectService *service.ProjectService // 注入 service 实例
}

func NewProjectHandler(projectService *service.ProjectService) *ProjectHandler {
	return &ProjectHandler{
		projectService: projectService,
	}
}

// @BasePath /api

// CreateRole  godoc
// CreateRole  godoc
// @Summary project create role
// @Schemes
// @Description project create role
// @Tags Project
// @Accept json
// @Produce json
// @Param ProjectUpsertRole body request.ProjectUpsertRole true "AdminRegister"
// @Success 200 {object} response.CommonResponseWithoutData "成功响应"
// @Router /project/create_role [post]
func (h *ProjectHandler) CreateRole(c *gin.Context) {
	var req request.ProjectUpsertRole
	if success := utils.ValidateJson(c, &req); !success {
		return
	}
	userId := service.GetUserIDFromAuthMiddleware(c)
	var err error
	if err = h.projectService.HasOwnPermission(userId, req.ProjectID); err != nil {
		c.JSON(http.StatusOK, response.CreateResponseWithoutData(1, err.Error()))
		return
	}
	if err = dao.CreateRole(req.UserID, req.ProjectID, req.RoleType); err != nil {
		c.JSON(http.StatusOK, response.CreateResponseWithoutData(1, err.Error()))
		return
	}
	c.JSON(http.StatusOK, response.CreateResponseWithoutData(0, "ok"))
}

// DeleteRole  godoc
// @Summary project delete role
// @Schemes
// @Description project delete role
// @Tags Project
// @Accept json
// @Produce json
// @Param ProjectDeleteRole body request.ProjectDeleteRole true "ProjectDeleteRole"
// @Success 200 {object} response.CommonResponseWithoutData "成功响应"
// @Router /project/delete_role [post]
func (h *ProjectHandler) DeleteRole(c *gin.Context) {
	var req request.ProjectDeleteRole
	if success := utils.ValidateJson(c, &req); !success {
		return
	}
	userId := service.GetUserIDFromAuthMiddleware(c)
	var err error
	if err = h.projectService.HasOwnPermission(userId, req.ProjectID); err != nil {
		c.JSON(http.StatusOK, response.CreateResponseWithoutData(1, err.Error()))
		return
	}
	if err := dao.DeleteRole(req.UserID, req.ProjectID); err != nil {
		c.JSON(http.StatusOK, response.CreateResponseWithoutData(1, err.Error()))
		return
	}
	c.JSON(http.StatusOK, response.CreateResponseWithoutData(0, "ok"))
}

// UpdateRole  godoc
// @Summary project update role
// @Schemes
// @Description project update role
// @Tags Project
// @Accept json
// @Produce json
// @Param ProjectUpsertRole body request.ProjectUpsertRole true "AdminRegister"
// @Success 200 {object} response.CommonResponseWithoutData "成功响应"
// @Router /project/update_role [post]
func (h *ProjectHandler) UpdateRole(c *gin.Context) {
	var req request.ProjectUpsertRole
	if success := utils.ValidateJson(c, &req); !success {
		return
	}
	userId := service.GetUserIDFromAuthMiddleware(c)
	var err error
	if err = h.projectService.HasOwnPermission(userId, req.ProjectID); err != nil {
		c.JSON(http.StatusOK, response.CreateResponseWithoutData(1, err.Error()))
		return
	}
	if err := dao.UpdateRole(req.UserID, req.ProjectID, req.RoleType); err != nil {
		c.JSON(http.StatusOK, response.CreateResponseWithoutData(1, err.Error()))
		return
	}
	c.JSON(http.StatusOK, response.CreateResponseWithoutData(0, "ok"))
}

// List godoc
// @Summary project list
// @Schemes
// @Description project list
// @Tags Project
// @Accept */*
// @Produce json
// @Param ProjectList query request.ProjectList true "ProjectList"
// @Success 200 {object} response.ProjectList "成功响应"
// @Router /project/list [get]
func (h *ProjectHandler) List(c *gin.Context) {
	var req request.ProjectList
	if success := utils.ValidateQuery(c, &req); !success {
		return
	}

	userID := service.GetUserIDFromAuthMiddleware(c)

	responseProjectList, err := h.projectService.GetProjectList(userID, req.Status, req.RoleType, req.Page, req.PageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.CreateResponse(1, "error", nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(0, "ok", responseProjectList))
}

// Update  godoc
// @Summary project update
// @Schemes
// @Description project update
// @Tags Project
// @Accept json
// @Produce json
// @Param ProjectUpdate body request.ProjectUpdate true "ProjectUpdate"
// @Success 200 {object} response.CommonResponseWithoutData "成功响应"
// @Router /project/update [post]
func (h *ProjectHandler) Update(c *gin.Context) {
	var req request.ProjectUpdate
	if success := utils.ValidateJson(c, &req); !success {
		return
	}
	userId := service.GetUserIDFromAuthMiddleware(c)
	var err error
	if err = h.projectService.HasOwnPermission(userId, req.ID); err != nil {
		c.JSON(http.StatusOK, response.CreateResponseWithoutData(1, err.Error()))
		return
	}
	project := entity.Project{
		Name:        req.Name,
		Description: req.Description,
		GitAddress:  req.GitAddress,
		Config:      req.Config,
	}
	if err := h.projectService.UpdateProject(req.ID, project); err != nil {
		c.JSON(http.StatusOK, response.CreateResponseWithoutData(1, err.Error()))
		return
	}
	c.JSON(http.StatusOK, response.CreateResponseWithoutData(0, "ok"))
}

// UpdateStatus  godoc
// @Summary project update status
// @Schemes
// @Description project update status
// @Tags Project
// @Accept json
// @Produce json
// @Param ProjectUpdateStatus body request.ProjectUpdateStatus true "ProjectUpdateStatus"
// @Success 200 {object} response.CommonResponseWithoutData "成功响应"
// @Router /project/update_status [post]
func (h *ProjectHandler) UpdateStatus(c *gin.Context) {
	var req request.ProjectUpdateStatus
	if success := utils.ValidateJson(c, &req); !success {
		return
	}
	userId := service.GetUserIDFromAuthMiddleware(c)
	var err error
	if err = h.projectService.HasOwnPermission(userId, req.ID); err != nil {
		c.JSON(http.StatusOK, response.CreateResponseWithoutData(1, err.Error()))
		return
	}
	if err := h.projectService.UpdateProjectStatus(req.ID, req.Status); err != nil {
		c.JSON(http.StatusOK, response.CreateResponseWithoutData(1, err.Error()))
		return
	}
	c.JSON(http.StatusOK, response.CreateResponseWithoutData(0, "ok"))
}

// Info godoc
// @Summary project info
// @Schemes
// @Description project info
// @Tags Project
// @Accept */*
// @Produce json
// @Param CommonID query request.CommonID true "CommonID"
// @Success 200 {object} response.CommonResponse[data=response.ProjectInfo] "成功响应"
// @Router /project/info [get]
func (h *ProjectHandler) Info(c *gin.Context) {
	var req request.CommonID
	if success := utils.ValidateQuery(c, &req); !success {
		return
	}
	projectInfo, err := h.projectService.ProjectInfo(req.ID)
	if err != nil {
		c.JSON(http.StatusOK, response.CreateResponseWithoutData(1, err.Error()))
		return
	}
	c.JSON(http.StatusOK, response.CreateResponse(0, "ok", projectInfo))
}

// RoleIn godoc
// @Summary project role in
// @Schemes
// @Description project role
// @Tags Project
// @Accept */*
// @Produce json
// @Param CommonID query request.CommonID true "CommonID"
// @Success 200 {object} response.CommonResponse[data=response.ProjectRole] "成功响应"
// @Router /project/role_in [get]
func (h *ProjectHandler) RoleIn(c *gin.Context) {
	var req request.CommonID
	if success := utils.ValidateQuery(c, &req); !success {
		return
	}
	userID := service.GetUserIDFromAuthMiddleware(c)
	key := utils.JoinCacheKey(my_runtime.PrefixUseridProjectRole, userID, req.ID)
	roleType := interfaces.GetOrSet(interfaces.Cache, key, func() entity.RoleType {
		return h.projectService.GetRoleType(userID, req.ID)
	}, time.Hour)

	c.JSON(http.StatusOK, response.CreateResponse(0, "ok", response.ProjectRole{
		RoleType: roleType,
	}))
}

// TaskList godoc
// @Summary project task_list
// @Description project task_list
// @Tags Project
// @Accept */*
// @Produce json
// @Param ProjectTaskList body request.ProjectTaskList true "ProjectTaskList"
// @Success 200 {object} response.CommonResponse[data=response.TaskList] "成功响应"
// @Router /project/task_list [post]
func (h *ProjectHandler) TaskList(c *gin.Context) {
	var req request.ProjectTaskList
	if success := utils.ValidateJson(c, &req); !success {
		return
	}

	taskList, err := h.projectService.GetTaskList(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.CreateResponse(1, err.Error(), nil))
		return
	}
	c.JSON(http.StatusOK, response.CreateResponse(0, "ok", taskList))
}

// UserList
// @Summary  user_list
// @Description user_list
// @Tags Project
// @Accept */*
// @Produce json
// @Param CommonID query request.CommonID true "CommonID"
// @Success 200 {object} response.CommonResponse[data=response.UserList] "成功响应"
// @Router /project/user_list [get]
func (h *ProjectHandler) UserList(c *gin.Context) {
	var req request.CommonID
	if success := utils.ValidateQuery(c, &req); !success {
		return
	}
	res, err := h.projectService.UserList(req.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.CreateResponse(1, "error", nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(0, "ok", res))
}
