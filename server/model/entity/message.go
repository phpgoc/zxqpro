package entity

import "time"

type Action byte

const (
	ActionnShareLink Action = iota + 1
)

type Message struct {
	ID           uint `gorm:"primarykey"`
	CreatedAt    time.Time
	CreateUserId uint
	CreateUser   User        `gorm:"foreignKey:CreateUserId;references:ID"`
	ToList       []MessageTo `gorm:"foreignKey:MessageId"`
	Action       Action
	Link         *string
}

type MessageTo struct {
	ID        uint    `gorm:"primarykey"`
	UserId    uint    `gorm:"index:idx_user_read_message,priority:1"`
	User      User    `gorm:"foreignKey:UserId;references:ID"`
	MessageId uint    `gorm:"index:idx_user_read_message,priority:3"`
	Message   Message `gorm:"foreignKey:MessageId;references:ID"`
	Read      bool    `gorm:"index:idx_user_read_message,priority:2"`
	ReadAt    time.Time
}
