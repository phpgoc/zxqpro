package entity

import (
	"database/sql"
	"time"
)

type RoleType byte

const (
	RoleTypeOwner RoleType = iota + 1
	RoleTypeProducter
	RoleTypeDeveloper
	RoleTypeTester
	RoleTypeAdmin
)

type Role struct {
	UserID    uint    `gorm:"primaryKey"`
	User      User    `gorm:"foreignKey:UserID;references:ID"`
	ProjectID uint    `gorm:"primaryKey"`
	Project   Project `gorm:"foreignKey:ProjectID;references:ID"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime `gorm:"index"`
	RoleType  RoleType     `gorm:"type:tinyint;default:0"` // 0: Manager, 1: Developer, 2: Tester
}
