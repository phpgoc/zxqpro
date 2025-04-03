package routes

import (
	"net/http"

	"github.com/phpgoc/zxqpro/orm_model"
	"github.com/phpgoc/zxqpro/utils"

	"github.com/gin-gonic/gin"
	"github.com/phpgoc/zxqpro/request"
	"github.com/phpgoc/zxqpro/response"
)

// @BasePath /api

// UserRegister  godoc
// @Summary user register
// @Schemes
// @Description do hello
// @Tags User
// @Accept json
// @Produce json
// @Param user body request.Register true "UserRegister"
// @Success 200 {object} response.CommonResponseWithoutData "成功响应"
// @Router /user/register [post]
func UserRegister(g *gin.Context) {
	var req request.Register
	if success := utils.Validate(g, &req); !success {
		return
	}
	result := utils.Db.Create(&orm_model.User{Name: req.Name, Password: req.Password, Email: req.Email})
	if result.RowsAffected == 1 {
		g.JSON(http.StatusOK, response.CreateResponseWithoutData(0, "ok"))
	} else {
		g.JSON(http.StatusOK, response.CreateResponseWithoutData(1, result.Error.Error()))
	}
}
