package entity

import "gorm.io/gorm"

type Step struct {
	gorm.Model
	TaskID      uint     `json:"task_id" gorm:"not null"`
	Description string   `json:"description"`
	Percent     *byte    `json:"percent"`
	DeveloperID uint     `json:"developer_id"`
	Developer   User     `json:"developer" gorm:"foreignKey:DeveloperID;references:ID"`
	TesterID    uint     `json:"tester_id"`
	Tester      *User    `json:"tester" gorm:"foreignKey:TesterID;references:ID"`
	Commits     []string `json:"commits" gorm:"type:text"` // 提交的commit id
}
