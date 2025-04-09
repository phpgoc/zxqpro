package entity

import "gorm.io/gorm"

type Project struct {
	gorm.Model
	Name        string        `json:"name" gorm:"unique;not null"`
	OwnerID     uint          `json:"owner_id" gorm:"not null"`
	Description string        `json:"description"`
	GitAddress  string        `json:"git_address"`
	Status      ProjectStatus `json:"status" gorm:"type:tinyint;default:0;min=0;max=3"`
	Member      []Role        `gorm:"foreignKey:ProjectID;references:ID"` // Many-to-Many relationship with Role]
}
type ProjectStatus byte

const (
	Draft ProjectStatus = iota + 1
	Active
	Completed
	Archived
)
