package entity

import "gorm.io/gorm"

type Project struct {
	gorm.Model
	Name        string `json:"name" gorm:"unique;not null"`
	OwnerID     uint   `json:"owner_id" gorm:"not null"`
	Description string `json:"description"`
	GitAddress  string `json:"git_address"`
	Status      byte   `json:"status" gorm:"type:tinyint;default:0"` // 0: Draft, 1: Active, 2: Inactive
	Member      []Role `gorm:"foreignKey:ProjectID;references:ID"`   // Many-to-Many relationship with Role]
}
