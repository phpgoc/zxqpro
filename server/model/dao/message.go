package dao

import (
	"fmt"

	"github.com/phpgoc/zxqpro/model/entity"
	"github.com/phpgoc/zxqpro/my_runtime"
)

func CreateMessage(createUserId uint, toUserId []uint, action entity.Action, link, content *string) error {
	message := entity.Message{
		CreateUserId:   createUserId,
		Action:         action,
		Link:           link,
		MessageContent: content,
	}
	result := my_runtime.Db.Create(&message)
	if result.Error != nil {
		return result.Error
	}
	for _, id := range toUserId {
		messageTo := entity.MessageTo{
			UserId:    id,
			MessageId: message.ID,
			Read:      false,
		}
		result = my_runtime.Db.Create(&messageTo)
		if result.Error != nil {
			return result.Error
		}
	}
	return nil
}

func JoinReceiveMessage(fromUserName, toUserName string, action entity.Action, messageContent *string) string {
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
