package request

import (
	"time"

	"github.com/phpgoc/zxqpro/model/entity"
)

type ProjectUpsertRole struct {
	UserID    uint            `json:"user_id"`
	ProjectID uint            `json:"project_id"`
	RoleType  entity.RoleType `json:"role_type"`
}

type ProjectDeleteRole struct {
	UserID    uint `json:"user_id"`
	ProjectID uint `json:"project_id"`
}

type ProjectUpdate struct {
	ID          uint                      `json:"id" binding:"required"`
	Name        string                    `json:"name" gorm:"unique;not null"`
	Description string                    `json:"description"`
	GitAddress  string                    `json:"git_address"`
	Config      entity.NoOrmProjectConfig `json:"config"`
}
type ProjectUpdateStatus struct {
	ID     uint                 `json:"id" binding:"required"`
	Status entity.ProjectStatus `json:"status" binding:"min=0,max=4"`
}

type ProjectUpdateConfig struct {
	ID     uint                      `json:"id" binding:"required"`
	Config entity.NoOrmProjectConfig `json:"config"`
}
type ProjectList struct {
	Page     int  `form:"page"  bindings:"min=1" default:"1"`
	PageSize int  `form:"page_size" bindings:"min=1,max=100" default:"10"`
	Status   byte `form:"status" bindings:"min=0,max=4" default:"0"`
	RoleType byte `form:"role_type" bindings:"min=0,max=5" default:"0"`
}

type TaskTimeEstimateCreate struct {
	TaskID                  uint           `json:"task_id" binding:"required,min=1"`
	TaskDuration            *time.Duration `json:"task_duration"`             // 预计完成时间
	EstimatedCompletionTime *time.Time     `json:"estimated_completion_time"` // 预计完成时间
}
