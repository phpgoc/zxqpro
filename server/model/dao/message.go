package dao

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/phpgoc/zxqpro/model/entity"
)

type MessageDAO struct {
	db *gorm.DB
}

func NewMessageDAO(db *gorm.DB) *MessageDAO {
	return &MessageDAO{db: db}
}

func (d *MessageDAO) GetSendListAndCount(userID uint, page, pageSize int) (int64, []entity.Message, error) {
	var total int64
	var messageList []entity.Message

	// 先获取消息总数
	countResult := d.db.Model(&entity.Message{}).Where("create_user_id = ?", userID).Count(&total)
	if countResult.Error != nil {
		return 0, nil, countResult.Error
	}

	// 再获取分页消息列表
	listResult := d.db.Model(&entity.Message{}).
		Where("create_user_id = ?", userID).
		Preload("ToList").
		Preload("ToList.User").
		Offset((page - 1) * pageSize).
		Limit(pageSize).
		Order("id desc").
		Find(&messageList)

	return total, messageList, listResult.Error
}

func (d *MessageDAO) GetReceiveListAndCount(userID uint, page, pageSize int, read bool) (int64, []entity.MessageTo, error) {
	var total int64
	var messageToList []entity.MessageTo

	// 先获取消息总数
	countResult := d.db.Model(&entity.MessageTo{}).Where("user_id = ?", userID).Where("read = ?", read).Count(&total)
	if countResult.Error != nil {
		return 0, nil, countResult.Error
	}

	// 再获取分页消息列表
	listResult := d.db.Model(&entity.MessageTo{}).
		Where("user_id = ?", userID).
		Where("read = ?", read).
		Preload(clause.Associations).Preload("Message.Create").
		Order("id desc").
		Offset((page - 1) * pageSize).
		Limit(pageSize).
		Find(&messageToList)

	return total, messageToList, listResult.Error
}

func (d *MessageDAO) CreateMessage(createUserID uint, toUserID []uint, action entity.Action, link, content *string) error {
	message := entity.Message{
		CreateUserID:   createUserID,
		Action:         action,
		Link:           link,
		MessageContent: content,
	}
	result := d.db.Create(&message)
	if result.Error != nil {
		return result.Error
	}
	for _, id := range toUserID {
		messageTo := entity.MessageTo{
			UserID:    id,
			MessageID: message.ID,
			Read:      false,
		}
		result = d.db.Create(&messageTo)
		if result.Error != nil {
			return result.Error
		}
	}
	return nil
}

func (d *MessageDAO) UpdateMessageToRead(userID, messageToID uint) error {
	messageTo := entity.MessageTo{
		ID:     messageToID,
		UserID: userID,
	}
	return d.db.Model(&messageTo).Updates(map[string]interface{}{"read": true}).Error
}
