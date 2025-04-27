package service

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/phpgoc/zxqpro/model/dao"
	"github.com/phpgoc/zxqpro/model/entity"
	"github.com/phpgoc/zxqpro/routes/request"
	"github.com/phpgoc/zxqpro/utils"
)

type MessageService struct {
	messageDAO *dao.MessageDAO
	userDAO    *dao.UserDAO
}

func NewMessageService(messageDAO *dao.MessageDAO, userDAO *dao.UserDAO) *MessageService {
	return &MessageService{
		messageDAO: messageDAO,
		userDAO:    userDAO,
	}
}

func (s *MessageService) ShareLink(c *gin.Context, req *request.MessageShareLink) error {
	userID := GetUserIDFromAuthMiddleware(c)
	if err := dao.CreateMessage(userID, []uint{req.ToUserID}, entity.ActionShareLink, &req.Link, nil); err != nil {
		return err
	}
	sseManager := c.MustGet("sseManager").(*utils.SSEManager)
	user, _ := s.userDAO.GetByID(userID)
	toUser, err := s.userDAO.GetByID(req.ToUserID)
	if err != nil {
		return err
	}

	sseManager.SendMessageToUser(req.ToUserID,
		utils.SSEMessage{
			Message: s.JoinReceiveMessage(user.UserName, toUser.UserName, entity.ActionShareLink, nil),
			Link:    &(req.Link),
		})
	return nil
}

func (s *MessageService) JoinReceiveMessage(fromUserName, toUserName string, action entity.Action, messageContent *string) string {
	switch action {
	case entity.ActionShareLink:
		return fmt.Sprintf("%s 共享了一个链接给你.", fromUserName)
	case entity.ActionManual:
		return fmt.Sprintf("%s 发送了一个消息给你： %s", fromUserName, *messageContent)
	default:
		return ""

	}
}

func JoinSendMessage(fromUserName string, toUserName string, action entity.Action, messageContent *string) string {
	switch action {
	case entity.ActionShareLink:
		return fmt.Sprintf("你共享了一个链接给 %s.", toUserName)
	case entity.ActionManual:
		return fmt.Sprintf("你发送了一个消息给 %s： %s", toUserName, *messageContent)
	default:
		return ""
	}
}
