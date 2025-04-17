package routes

import (
	"fmt"
	"net/http"

	"github.com/phpgoc/zxqpro/routes/middleware"

	"github.com/gin-gonic/gin"
	"github.com/phpgoc/zxqpro/model/dao"
	"github.com/phpgoc/zxqpro/model/entity"
	"github.com/phpgoc/zxqpro/routes/request"
	"github.com/phpgoc/zxqpro/routes/response"
	"github.com/phpgoc/zxqpro/utils"
)

// @BasePath /api

// AdminRegister  godoc
// @Summary admin register
// @Schemes
// @Description admin register
// @Tags Admin
// @Accept json
// @Produce json
// @Param AdminRegister body request.AdminRegister true "AdminRegister"
// @Success 200 {object} response.CommonResponseWithoutData "成功响应"
// @Router /admin/register [post]
func AdminRegister(g *gin.Context) {
	var req request.AdminRegister
	if success := utils.ValidateJson(g, &req); !success {
		return
	}
	user := entity.User{Name: req.Name, UserName: req.Name, Password: req.Password}
	if err := dao.CreateUser(&user); err != nil {
		g.JSON(http.StatusOK, response.CreateResponseWithoutData(1, err.Error()))
		return
	}
	g.JSON(http.StatusOK, response.CreateResponseWithoutData(0, "ok"))
}

// AdminUpdatePassword  godoc
// @Summary admin update_password
// @Schemes
// @Description admin update_password
// @Tags Admin
// @Accept json
// @Produce json
// @Param AdminUpdatePassword body request.AdminUpdatePassword true "AdminUpdatePassword"
// @Success 200 {object} response.CommonResponseWithoutData "成功响应"
// @Router /admin/update_password [post]
func AdminUpdatePassword(c *gin.Context) {
	var req request.AdminUpdatePassword
	if success := utils.ValidateJson(c, &req); !success {
		return
	}
	user := entity.User{}
	result := dao.Db.Where("id = ?", req.UserId).First(&user)
	if result.Error != nil {
		c.JSON(http.StatusOK, response.CreateResponseWithoutData(1, result.Error.Error()))
		return
	}
	user.Password = dao.Md5Password(req.Password, user.ID)
	if err := dao.Db.Save(&user).Error; err != nil {
		c.JSON(http.StatusOK, response.CreateResponseWithoutData(1, err.Error()))
		return
	}
	c.JSON(http.StatusOK, response.CreateResponseWithoutData(0, "ok"))
}

// AdminCreateProject  godoc
// @Summary admin create_project
// @Schemes
// @Description admin create_project
// @Tags Admin
// @Accept json
// @Produce json
// @Param AdminCreateProject body request.AdminCreateProject true "AdminCreateProject"
// @Success 200 {object} response.CommonResponseWithoutData "成功响应"
// @Router /admin/create_project [post]
func AdminCreateProject(c *gin.Context) {
	var req request.AdminCreateProject
	if success := utils.ValidateJson(c, &req); !success {
		return
	}
	// 判断id是否存在 没删除的
	var user entity.User
	result := dao.Db.Where("id = ?", req.OwnerID).Where("deleted_at IS NULL").First(&user)
	if result.Error != nil {
		c.JSON(http.StatusOK, response.CreateResponseWithoutData(1, result.Error.Error()))
		return
	}

	project := entity.Project{
		Name:        req.Name,
		OwnerID:     req.OwnerID,
		Description: req.Description,
		Status:      entity.ProjectStatusInActive,
		Config:      entity.DefaultProjectConfig(),
	}
	if err := dao.Db.Create(&project).Error; err != nil {
		c.JSON(http.StatusOK, response.CreateResponseWithoutData(1, err.Error()))
		return
	}
	_ = dao.CreateRole(project.OwnerID, project.ID, entity.RoleTypeOwner)
	c.JSON(http.StatusOK, response.CreateResponseWithoutData(0, "ok"))
}

// AdminResetRateLimit  godoc
// @Summary admin reset_rate_limit
// @Schemes
// @Description admin reset_rate_limit
// @Tags Admin
// @Produce json
// @Success 200 {object} response.CommonResponseWithoutData "成功响应"
// @Router /admin/reset_rate_limit [post]
func AdminResetRateLimit(c *gin.Context) {
	sum := middleware.CleanAllMap()
	c.JSON(http.StatusOK, response.CreateResponseWithoutData(0, fmt.Sprintf("ok,cleaned %d ", sum)))
}
