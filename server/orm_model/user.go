package orm_model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `gorm:"unique;not null"  json:"name"`
	Password string `gorm:"not null" json:"password"`
	Email    string `gorm:"unique" json:"email"`
}
