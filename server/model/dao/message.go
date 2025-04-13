package dao

import (
	"fmt"

	"github.com/phpgoc/zxqpro/model/entity"
)

func CreateMessage(createUserId uint, toUserId []uint, action entity.Action, link *string) error {
	message := entity.Message{
		CreateUserId: createUserId,
		Action:       action,
		Link:         link,
	}
	result := Db.Create(&message)
	if result.Error != nil {
		return result.Error
	}
	for _, id := range toUserId {
		messageTo := entity.MessageTo{
			UserId:    id,
			MessageId: message.ID,
			Read:      false,
		}
		result = Db.Create(&messageTo)
		if result.Error != nil {
			return result.Error
		}
	}
	return nil
}

func JoinReceiveMessage(fromUserName string, toUserName string, action entity.Action) string {
	switch action {
	case entity.ActionnShareLink:
		return fmt.Sprintf("%s 共享了一个链接给你.", fromUserName)
	default:
		return ""

	}
}

func JoinSendMessage(fromUserName string, toUserName string, action entity.Action) string {
	switch action {
	case entity.ActionnShareLink:
		return fmt.Sprintf("你共享了一个链接给 %s.", toUserName)
	default:
		return ""
	}
}
