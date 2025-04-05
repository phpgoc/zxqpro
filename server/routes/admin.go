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
// @Param user body request.AdminRegister true "AdminRegister"
// @Success 200 {object} response.CommonResponseWithoutData "成功响应"
// @Router /admin/register [post]
func AdminRegister(g *gin.Context) {
	var req request.AdminRegister
	if success := utils.Validate(g, &req); !success {
		return
	}
	user := entity.User{Name: req.Name, Password: req.Password}
	if err := dao.CreateUser(&user); err != nil {
		g.JSON(http.StatusOK, response.CreateResponseWithoutData(1, err.Error()))
		return
	}
	g.JSON(http.StatusOK, response.CreateResponseWithoutData(0, "ok"))
}
