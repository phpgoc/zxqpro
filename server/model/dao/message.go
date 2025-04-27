package dao

import (
	"gorm.io/gorm"

	"github.com/phpgoc/zxqpro/model/entity"
	"github.com/phpgoc/zxqpro/my_runtime"
)

type MessageDAO struct {
	db *gorm.DB
}

func NewMessageDAO(db *gorm.DB) *MessageDAO {
	return &MessageDAO{db: db}
}

func CreateMessage(createUserID uint, toUserID []uint, action entity.Action, link, content *string) error {
	message := entity.Message{
		CreateUserID:   createUserID,
		Action:         action,
		Link:           link,
		MessageContent: content,
	}
	result := my_runtime.Db.Create(&message)
	if result.Error != nil {
		return result.Error
	}
	for _, id := range toUserID {
		messageTo := entity.MessageTo{
			UserID:    id,
			MessageID: message.ID,
			Read:      false,
		}
		result = my_runtime.Db.Create(&messageTo)
		if result.Error != nil {
			return result.Error
		}
	}
	return nil
}
