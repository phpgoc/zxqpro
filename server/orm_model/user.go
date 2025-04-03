package orm_model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	name     string `gorm:"unique;not null"`
	password string `gorm:"not null"`
	email    string `gorm:"unique"`
}
