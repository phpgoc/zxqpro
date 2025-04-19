package routes

import (
	"net/http"
	"strings"

	"github.com/phpgoc/zxqpro/model/service"

	"github.com/phpgoc/zxqpro/my_runtime"

	"gorm.io/gorm/clause"

	"github.com/gin-gonic/gin"
	"github.com/phpgoc/zxqpro/model/dao"
	"github.com/phpgoc/zxqpro/model/entity"
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
	userId := service.GetUserIdFromAuthMiddleware(c)

	if err := dao.CreateMessage(userId, []uint{req.ToUserId}, entity.ActionShareLink, &req.Link, nil); err != nil {
		c.JSON(http.StatusOK, response.CreateResponseWithoutData(1, err.Error()))
		return
	}
	sseManager := c.MustGet("sseManager").(*utils.SSEManager)
	user, _ := dao.GetUserById(userId)
	toUser, _ := dao.GetUserById(req.ToUserId)
	sseManager.SendMessageToUser(req.ToUserId,
		utils.SSEMessage{
			Message: dao.JoinReceiveMessage(user.UserName, toUser.UserName, entity.ActionShareLink, nil),
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
// @Success 200 {object} response.CommonResponse[data=response.MessageList] "成功响应"
// @Router /message/receive_list [get]
func MessageReceiveList(c *gin.Context) {
	var req request.MessageList
	if success := utils.ValidateQuery(c, &req); !success {
		return
	}
	userId := service.GetUserIdFromAuthMiddleware(c)
	var res response.MessageList
	var messageToList []entity.MessageTo
	model := my_runtime.Db.Model(entity.MessageTo{}).Preload(clause.Associations).Preload("Message.CreateUser").Where("message_tos.user_id = ?", userId).Where("message_tos.read = ?", req.Read)

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
		res.List = append(res.List, response.Message{
			Id:       messageTo.ID,
			Link:     messageTo.Message.Link,
			UserName: messageTo.Message.CreateUser.UserName,
			Message:  dao.JoinReceiveMessage(messageTo.Message.CreateUser.UserName, messageTo.User.UserName, messageTo.Message.Action, messageTo.Message.MessageContent),
			Time:     messageTo.Message.CreatedAt.Format("2006-01-02 15:04:05"),
			Read:     messageTo.Read,
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
// @Success 200 {object} response.CommonResponse[data=response.MessageList] "成功响应"
// @Router /message/send_list [get]
func MessageSendList(c *gin.Context) {
	var req request.Page
	if success := utils.ValidateQuery(c, &req); !success {
		return
	}
	userId := service.GetUserIdFromAuthMiddleware(c)
	var res response.MessageList
	var messageList []entity.Message
	model := my_runtime.Db.Model(entity.Message{}).Where("create_user_id = ?", userId).Preload("ToList").Preload("ToList.User")
	result := model.Count(&res.Total)
	if result.Error != nil {
		c.JSON(http.StatusOK, response.CreateResponseWithoutData(1, result.Error.Error()))
		return
	}
	result = model.Offset((req.Page - 1) * req.PageSize).Limit(req.PageSize).Order("id desc").Find(&messageList)

	if result.Error != nil {
		c.JSON(http.StatusOK, response.CreateResponseWithoutData(1, result.Error.Error()))
		return
	}
	for _, message := range messageList {
		var names []string
		for _, messageToUser := range message.ToList {
			names = append(names, messageToUser.User.UserName)
		}
		allRead := true
		for _, messageToUser := range message.ToList {
			allRead = allRead && messageToUser.Read
		}
		joinedNames := strings.Join(names, ",")
		res.List = append(res.List, response.Message{
			Id:       message.ID,
			UserName: joinedNames,
			Link:     message.Link,
			Message:  dao.JoinSendMessage(message.CreateUser.UserName, joinedNames, message.Action, message.MessageContent),
			Time:     message.CreatedAt.Format("2006-01-02 15:04:05"),
			Read:     allRead,
		})
	}
	c.JSON(http.StatusOK, response.CreateResponse(0, "ok", res))
}

// MessageRead  godoc
// @Summary message read
// @Schemes
// @Description message read
// @Tags Message
// @Accept json
// @Produce json
// @Param MessageRead body request.MessageRead true "MessageRead"
// @Success 200 {object} response.CommonResponseWithoutData "成功响应"
// @Router /message/read [post]
func MessageRead(c *gin.Context) {
	var req request.MessageRead
	if success := utils.ValidateJson(c, &req); !success {
		return
	}
	userId := service.GetUserIdFromAuthMiddleware(c)
	var messageTo entity.MessageTo
	result := my_runtime.Db.Model(entity.MessageTo{}).Where("id = ? and user_id = ?", req.Id, userId).First(&messageTo)
	if result.Error != nil {
		c.JSON(http.StatusOK, response.CreateResponseWithoutData(1, result.Error.Error()))
		return
	}
	messageTo.Read = true
	my_runtime.Db.Save(&messageTo)

	c.JSON(http.StatusOK, response.CreateResponseWithoutData(0, "ok"))
}

// MessageManual  godoc
// @Summary message manual
// @Schemes
// @Description message manual
// @Tags Message
// @Accept json
// @Produce json
// @Param MessageRead body request.ManualMessage true "MessageRead"
// @Success 200 {object} response.CommonResponseWithoutData "成功响应"
// @Router /message/manual [post]
func MessageManual(c *gin.Context) {
	var req request.ManualMessage
	if success := utils.ValidateJson(c, &req); !success {
		return
	}
	userId := service.GetUserIdFromAuthMiddleware(c)

	if err := dao.CreateMessage(userId, req.UserIds, entity.ActionManual, req.Link, &req.Content); err != nil {
		c.JSON(http.StatusOK, response.CreateResponseWithoutData(1, err.Error()))
		return
	}

	sseManager := c.MustGet("sseManager").(*utils.SSEManager)
	user, _ := dao.GetUserById(userId)
	for _, id := range req.UserIds {
		toUser, _ := dao.GetUserById(id)
		sseManager.SendMessageToUser(id,
			utils.SSEMessage{
				Code:    0,
				Message: dao.JoinReceiveMessage(user.UserName, toUser.UserName, entity.ActionManual, &req.Content),
			})
	}
	c.JSON(http.StatusOK, response.CreateResponseWithoutData(0, "ok"))
}
