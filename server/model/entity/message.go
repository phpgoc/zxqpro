package entity

import "time"

type Action byte

const (
	ActionShareLink Action = iota + 1
	ActionManual
)

type Message struct {
	ID             uint `gorm:"primarykey"`
	CreatedAt      time.Time
	CreateUserID   uint
	CreateUser     User `gorm:"foreignKey:CreateUserID;references:ID"`
	MessageContent *string
	ToList         []MessageTo `gorm:"foreignKey:MessageID"`
	Action         Action
	Link           *string
}

type MessageTo struct {
	ID        uint    `gorm:"primarykey"`
	UserID    uint    `gorm:"index:idx_user_read_message,priority:1"`
	User      User    `gorm:"foreignKey:UserID;references:ID"`
	MessageID uint    `gorm:"index:idx_user_read_message,priority:3"`
	Message   Message `gorm:"foreignKey:MessageID;references:ID"`
	Read      bool    `gorm:"index:idx_user_read_message,priority:2"`
	ReadAt    time.Time
}
