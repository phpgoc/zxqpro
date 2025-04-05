package routes

import (
	"net/http"

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
	user := entity.User{Name: req.Name, Password: req.Password}
	if err := dao.CreateUser(&user); err != nil {
		g.JSON(http.StatusOK, response.CreateResponseWithoutData(1, err.Error()))
		return
	}
	g.JSON(http.StatusOK, response.CreateResponseWithoutData(0, "ok"))
}

// AdminUserList  godoc
// @Summary admin user_list
// @Schemes
// @Description admin user_list
// @Tags Admin
// @Accept */*
// @Produce json
// @Param page query request.Page true "AdminUserList"
// @Success 200 {object} response.CommonResponse[data=response.UserList] "成功响应"
// @Router /admin/user_list [get]
func AdminUserList(g *gin.Context) {
	var req request.Page
	if success := utils.ValidateQuery(g, &req); !success {
		return
	}
	var total int64
	dao.Db.Model(&entity.User{}).Count(&total)
	total = total - 1
	var responseUsers []response.User
	dao.Db.Model(entity.User{}).Where("deleted_at IS NULL").Where("id != ?", 1).Offset((req.Page - 1) * req.PageSize).Limit(req.PageSize).Select("id, name, user_name, email, avatar").Find(&responseUsers)
	g.JSON(http.StatusOK, response.CreateResponse(0, "ok", response.UserList{
		Total: total,
		Users: responseUsers,
	}))
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
	result := dao.Db.Where("id = ?", req.ID).First(&user)
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
