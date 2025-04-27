package service

import (
	"fmt"
	"strings"

	"github.com/phpgoc/zxqpro/routes/response"

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
	if err := s.messageDAO.CreateMessage(userID, []uint{req.ToUserID}, entity.ActionShareLink, &req.Link, nil); err != nil {
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

func (s *MessageService) SendList(c *gin.Context, req *request.Page) (*response.MessageList, error) {
	userID := GetUserIDFromAuthMiddleware(c)
	var res response.MessageList
	var err error
	var messageList []entity.Message
	res.Total, messageList, err = s.messageDAO.GetSendListAndCount(userID, req.Page, req.PageSize)
	if err != nil {
		return nil, err
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
			Message:  s.JoinSendMessage(message.CreateUser.UserName, joinedNames, message.Action, message.MessageContent),
			Time:     message.CreatedAt.Format("2006-01-02 15:04:05"),
			Read:     allRead,
		})
	}
	return &res, nil
}

func (s *MessageService) ReceiveList(c *gin.Context, req *request.MessageList) (*response.MessageList, error) {
	userID := GetUserIDFromAuthMiddleware(c)
	var res response.MessageList
	var messageToList []entity.MessageTo
	var err error
	res.Total, messageToList, err = s.messageDAO.GetReceiveListAndCount(userID, req.Page, req.PageSize, req.Read)
	if err != nil {
		return nil, err
	}

	for _, messageTo := range messageToList {
		res.List = append(res.List, response.Message{
			ID:       messageTo.ID,
			Link:     messageTo.Message.Link,
			UserName: messageTo.Message.CreateUser.UserName,
			Message:  s.JoinReceiveMessage(messageTo.Message.CreateUser.UserName, messageTo.User.UserName, messageTo.Message.Action, messageTo.Message.MessageContent),
			Time:     messageTo.Message.CreatedAt.Format("2006-01-02 15:04:05"),
			Read:     messageTo.Read,
		})
	}
	return &res, nil
}

func (s *MessageService) Read(c *gin.Context, req *request.MessageRead) error {
	userID := GetUserIDFromAuthMiddleware(c)
	return s.messageDAO.UpdateMessageToRead(userID, req.ID)
}

func (s *MessageService) Manual(c *gin.Context, req *request.ManualMessage) error {
	userID := GetUserIDFromAuthMiddleware(c)

	if err := s.messageDAO.CreateMessage(userID, req.UserIDs, entity.ActionManual, req.Link, &req.Content); err != nil {
		return err
	}

	sseManager := c.MustGet("sseManager").(*utils.SSEManager)
	user, _ := dao.GetUserByID(userID)
	for _, id := range req.UserIDs {
		toUser, _ := dao.GetUserByID(id)
		sseManager.SendMessageToUser(id,
			utils.SSEMessage{
				Code:    0,
				Message: s.JoinReceiveMessage(user.UserName, toUser.UserName, entity.ActionManual, &req.Content),
			})
	}
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

func (s *MessageService) JoinSendMessage(fromUserName string, toUserName string, action entity.Action, messageContent *string) string {
	switch action {
	case entity.ActionShareLink:
		return fmt.Sprintf("你共享了一个链接给 %s.", toUserName)
	case entity.ActionManual:
		return fmt.Sprintf("你发送了一个消息给 %s： %s", toUserName, *messageContent)
	default:
		return ""
	}
}
