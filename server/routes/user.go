package routes

import (
	"net/http"

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
	if err := g.ShouldBindJSON(&req); err != nil {
		g.JSON(400, gin.H{
			"code":    400,
			"message": err.Error(),
		})
		return
	}
	g.JSON(http.StatusOK, response.CreateResponseWithoutData(0, "ok"))
}
