package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/phpgoc/zxqpro/model/dao"
	"github.com/phpgoc/zxqpro/model/entity"
	"github.com/phpgoc/zxqpro/routes/middleware"
	"github.com/phpgoc/zxqpro/routes/request"
	"github.com/phpgoc/zxqpro/routes/response"
	"github.com/phpgoc/zxqpro/utils"
	"gorm.io/gorm"
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
	if !hasPermission(c, req.ProjectId) {
		return
	}
	if err := dao.CreateRole(req.UserId, req.ProjectId, req.RoleType); err != nil {
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
	if !hasPermission(c, req.ProjectId) {
		return
	}
	if err := dao.DeleteRole(req.UserId, req.ProjectId); err != nil {
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
	if !hasPermission(c, req.ProjectId) {
		return
	}
	if err := dao.UpdateRole(req.UserId, req.ProjectId, req.RoleType); err != nil {
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
// @Produce json
// @Success 200 {object} response.ProjectList "成功响应"
// @Router /project/list [get]
func ProjectList(c *gin.Context) {
	userId := middleware.GetUserIdFromAuthMiddleware(c)

	var result *gorm.DB
	responseProjectList := response.ProjectList{}

	if userId == 1 {
		var projects []entity.Project
		result = dao.Db.Find(&projects)
		for _, project := range projects {
			responseProjectList.Projects = append(responseProjectList.Projects, response.Project{
				ID:       project.ID,
				Name:     project.Name,
				RoleType: entity.RoleTypeAdmin,
			})
		}

	} else {
		var roles []entity.Role
		result = dao.Db.Preload("Project").Where("user_id = ?", userId).Find(&roles)
		for _, role := range roles {
			responseProjectList.Projects = append(responseProjectList.Projects, response.Project{
				ID:       role.Project.ID,
				Name:     role.Project.Name,
				RoleType: role.RoleType,
			})
		}
	}
	_ = result

	c.JSON(http.StatusOK, response.CreateResponse(0, "ok", responseProjectList))
}

func hasPermission(c *gin.Context, projectId uint) bool {
	userId := middleware.GetUserIdFromAuthMiddleware(c)
	project := entity.Project{}
	result := dao.Db.First(&project, projectId)

	if result.Error != nil {
		c.JSON(http.StatusOK, response.CreateResponseWithoutData(1, "项目不存在"))
		return false
	}
	if userId == 1 {
		return true
	}
	if project.OwnerID != userId {
		c.JSON(http.StatusOK, response.CreateResponseWithoutData(1, "没有权限"))
		return false
	}
	return true
}
