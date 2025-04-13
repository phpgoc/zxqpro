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
)

// MessageShareLink  godoc
// @Summary message share link
// @Schemes
// @Description message share link
// @Tags Message
// @Accept json
// @Produce json
// @Param MessageShareLink body request.MessageShareLink true "MessageShareLink"
// @Success 200 {object} response.CommonResponseWithoutData "成功响应"
// @Router /message/share_link [post]
func MessageShareLink(c *gin.Context) {
	var req request.MessageShareLink
	if success := utils.ValidateJson(c, &req); !success {
		return
	}
	userId := middleware.GetUserIdFromAuthMiddleware(c)

	if err := dao.CreateMessage(userId, []uint{req.ToUserId}, entity.ActionnShareLink, &req.Link); err != nil {
		c.JSON(http.StatusOK, response.CreateResponseWithoutData(1, err.Error()))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponseWithoutData(0, "ok"))
}
