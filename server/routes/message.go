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

type MessageHandler struct {
	messageService *service.MessageService
}

func NewMessageHandler(messageService *service.MessageService) *MessageHandler {
	return &MessageHandler{
		messageService: messageService,
	}
}

// ShareLink  godoc
// @Summary message share link
// @Schemes
// @Description message share link
// @Tags Message
// @Accept json
// @Produce json
// @Param ShareLink body request.MessageShareLink true "ShareLink"
// @Success 200 {object} response.CommonResponseWithoutData "成功响应"
// @Router /message/share_link [post]
func (h *MessageHandler) ShareLink(c *gin.Context) {
	var req request.MessageShareLink
	if success := utils.ValidateJson(c, &req); !success {
		return
	}
	err := h.messageService.ShareLink(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.CreateResponseWithoutData(1, "error"))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponseWithoutData(0, "ok"))
}

// ReceiveList  godoc
// @Summary message receive_list
// @Schemes
// @Description message receive_list
// @Tags Message
// @Accept */*
// @Produce json
// @Param ReceiveList query request.MessageList true "ReceiveList"
// @Success 200 {object} response.CommonResponse[data=response.MessageList] "成功响应"
// @Router /message/receive_list [get]
func (h *MessageHandler) ReceiveList(c *gin.Context) {
	var req request.MessageList
	if success := utils.ValidateQuery(c, &req); !success {
		return
	}
	userID := service.GetUserIDFromAuthMiddleware(c)
	var res response.MessageList
	var messageToList []entity.MessageTo
	model := my_runtime.Db.Model(entity.MessageTo{}).Preload(clause.Associations).Preload("Message.Create").Where("message_tos.user_id = ?", userID).Where("message_tos.read = ?", req.Read)

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
			ID:       messageTo.ID,
			Link:     messageTo.Message.Link,
			UserName: messageTo.Message.CreateUser.UserName,
			Message:  h.messageService.JoinReceiveMessage(messageTo.Message.CreateUser.UserName, messageTo.User.UserName, messageTo.Message.Action, messageTo.Message.MessageContent),
			Time:     messageTo.Message.CreatedAt.Format("2006-01-02 15:04:05"),
			Read:     messageTo.Read,
		})
	}
	c.JSON(http.StatusOK, response.CreateResponse(0, "ok", res))
}

// SendList  godoc
// @Summary message send_list
// @Schemes
// @Description message send_list
// @Tags Message
// @Accept */*
// @Produce json
// @Param SendList query request.Page true "SendList"
// @Success 200 {object} response.CommonResponse[data=response.MessageList] "成功响应"
// @Router /message/send_list [get]
func (h *MessageHandler) SendList(c *gin.Context) {
	var req request.Page
	if success := utils.ValidateQuery(c, &req); !success {
		return
	}
	userID := service.GetUserIDFromAuthMiddleware(c)
	var res response.MessageList
	var messageList []entity.Message
	model := my_runtime.Db.Model(entity.Message{}).Where("create_user_id = ?", userID).Preload("ToList").Preload("ToList.User")
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
			ID:       message.ID,
			UserName: joinedNames,
			Link:     message.Link,
			Message:  service.JoinSendMessage(message.CreateUser.UserName, joinedNames, message.Action, message.MessageContent),
			Time:     message.CreatedAt.Format("2006-01-02 15:04:05"),
			Read:     allRead,
		})
	}
	c.JSON(http.StatusOK, response.CreateResponse(0, "ok", res))
}

// Read  godoc
// @Summary message read
// @Schemes
// @Description message read
// @Tags Message
// @Accept json
// @Produce json
// @Param Read body request.MessageRead true "Read"
// @Success 200 {object} response.CommonResponseWithoutData "成功响应"
// @Router /message/read [post]
func (h *MessageHandler) Read(c *gin.Context) {
	var req request.MessageRead
	if success := utils.ValidateJson(c, &req); !success {
		return
	}
	userID := service.GetUserIDFromAuthMiddleware(c)
	var messageTo entity.MessageTo
	result := my_runtime.Db.Model(entity.MessageTo{}).Where("id = ? and user_id = ?", req.ID, userID).First(&messageTo)
	if result.Error != nil {
		c.JSON(http.StatusOK, response.CreateResponseWithoutData(1, result.Error.Error()))
		return
	}
	messageTo.Read = true
	my_runtime.Db.Save(&messageTo)

	c.JSON(http.StatusOK, response.CreateResponseWithoutData(0, "ok"))
}

// Manual  godoc
// @Summary message manual
// @Schemes
// @Description message manual
// @Tags Message
// @Accept json
// @Produce json
// @Param Read body request.ManualMessage true "Read"
// @Success 200 {object} response.CommonResponseWithoutData "成功响应"
// @Router /message/manual [post]
func (h *MessageHandler) Manual(c *gin.Context) {
	var req request.ManualMessage
	if success := utils.ValidateJson(c, &req); !success {
		return
	}
	userID := service.GetUserIDFromAuthMiddleware(c)

	if err := dao.CreateMessage(userID, req.UserIDs, entity.ActionManual, req.Link, &req.Content); err != nil {
		c.JSON(http.StatusOK, response.CreateResponseWithoutData(1, err.Error()))
		return
	}

	sseManager := c.MustGet("sseManager").(*utils.SSEManager)
	user, _ := dao.GetUserByID(userID)
	for _, id := range req.UserIDs {
		toUser, _ := dao.GetUserByID(id)
		sseManager.SendMessageToUser(id,
			utils.SSEMessage{
				Code:    0,
				Message: h.messageService.JoinReceiveMessage(user.UserName, toUser.UserName, entity.ActionManual, &req.Content),
			})
	}
	c.JSON(http.StatusOK, response.CreateResponseWithoutData(0, "ok"))
}
