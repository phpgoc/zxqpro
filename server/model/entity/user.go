package entity

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string  `gorm:"unique;not null"  json:"name"`
	UserName string  `json:"user_name"`
	Password string  `gorm:"not null" json:"password"`
	Email    *string `gorm:"unique" json:"email"`
	Avatar   byte    `gorm:"size:255;default:0" json:"avatar"`
}
