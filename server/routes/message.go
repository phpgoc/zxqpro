package routes

import (
	"encoding/json"
	"net/http"

	"gorm.io/gorm/clause"

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
	sseManager := c.MustGet("sseManager").(*utils.SSEManager)
	user, _ := dao.GetUserById(userId)
	toUser, _ := dao.GetUserById(req.ToUserId)
	sseManager.SendMessageToUser(req.ToUserId,
		utils.SSEMessage{
			Message: dao.JoinReceiveMessage(user.UserName, toUser.UserName, entity.ActionnShareLink),
			Link:    &(req.Link),
		})

	c.JSON(http.StatusOK, response.CreateResponseWithoutData(0, "ok"))
}

// MessageReceiveList  godoc
// @Summary message receive_list
// @Schemes
// @Description message receive_list
// @Tags Message
// @Accept */*
// @Produce json
// @Param MessageReceiveList query request.MessageList true "MessageReceiveList"
// @Success 200 {object} response.CommonResponse[Data=response.MessageList] "成功响应"
// @Router /message/receive_list [get]
func MessageReceiveList(c *gin.Context) {
	var req request.MessageList
	if success := utils.ValidateQuery(c, &req); !success {
		return
	}
	userId := middleware.GetUserIdFromAuthMiddleware(c)
	var res response.MessageList
	var messageToList []entity.MessageTo
	model := dao.Db.Model(entity.MessageTo{}).Preload(clause.Associations).Preload("Message.CreateUser").Where(entity.MessageTo{UserId: userId, Read: req.Read})
	result := model.Count(&res.Total)
	if result.Error != nil {
		c.JSON(http.StatusOK, response.CreateResponseWithoutData(1, result.Error.Error()))
		return
	}
	// model.DryRun = true
	result = model.Offset((req.Page - 1) * req.PageSize).Limit(req.PageSize).Order("message_id desc").Find(&messageToList)
	// utils.LogError(result.Statement.SQL.String())
	if result.Error != nil {
		c.JSON(http.StatusOK, response.CreateResponseWithoutData(1, result.Error.Error()))
		return
	}
	for _, messageTo := range messageToList {
		j, _ := json.Marshal(messageTo.Message.CreateUser)
		utils.LogError(string(j))
		res.List = append(res.List, response.Message{
			Id:      messageTo.ID,
			Link:    messageTo.Message.Link,
			Message: dao.JoinReceiveMessage(messageTo.Message.CreateUser.UserName, messageTo.User.UserName, messageTo.Message.Action),
			Read:    messageTo.Read,
		})
	}
	c.JSON(http.StatusOK, response.CreateResponse(0, "ok", res))
}

// MessageSendList  godoc
// @Summary message send_list
// @Schemes
// @Description message send_list
// @Tags Message
// @Accept */*
// @Produce json
// @Param MessageSendList query request.Page true "MessageSendList"
// @Success 200 {object} response.CommonResponse[Data=response.MessageList] "成功响应"
// @Router /message/send_list [get]
func MessageSendList(c *gin.Context) {
	var req request.Page
	if success := utils.ValidateQuery(c, &req); !success {
		return
	}
	userId := middleware.GetUserIdFromAuthMiddleware(c)
	var res response.MessageList
	var messageToList []entity.MessageTo
	model := dao.Db.Model(entity.MessageTo{}).Preload("Message", "create_user_id = ?", userId).Preload("User")
	result := model.Count(&res.Total)
	if result.Error != nil {
		c.JSON(http.StatusOK, response.CreateResponseWithoutData(1, result.Error.Error()))
		return
	}
	result = model.Offset((req.Page - 1) * req.PageSize).Limit(req.PageSize).Order("message_id desc").Find(&messageToList)

	utils.LogError(result.Statement.SQL.String())

	if result.Error != nil {
		c.JSON(http.StatusOK, response.CreateResponseWithoutData(1, result.Error.Error()))
		return
	}
	for _, messageTo := range messageToList {
		res.List = append(res.List, response.Message{
			Id:      messageTo.ID,
			Link:    messageTo.Message.Link,
			Message: dao.JoinSendMessage(messageTo.Message.CreateUser.UserName, messageTo.User.UserName, messageTo.Message.Action),
			Read:    messageTo.Read,
		})
	}
	c.JSON(http.StatusOK, response.CreateResponse(0, "ok", res))
}
