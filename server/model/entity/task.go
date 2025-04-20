package entity

import (
	"time"

	"gorm.io/gorm"
)

type TaskStatus byte

const (
	TaskStatusCreated TaskStatus = iota + 1
	TaskStarted
	TaskStatusCompleted
	TaskStatusArchived
	TaskFailed
)

type Task struct {
	gorm.Model
	Name                   string         `json:"name" gorm:"not null"`
	Description            string         `json:"description" gorm:"type:text"`
	CreateUserID           uint           `json:"create_user_id"`
	CreateUser             User           `json:"create_user" gorm:"foreignKey:CreateUserID;references:ID"`
	ExpectCompleteDuration *time.Duration `json:"expect_complete_duration"` // 预计完成时间
	ExpectCompleteTime     *time.Time     `json:"expect_complete_time"`
	AssignUserID           uint           `json:"assign_user_id"` // 如果是顶级任务，不使用这个字典，非顶级任务，指定的用户是一对一的。顶级任务是一对多的
	AssignUser             User           `json:"assign_user" gorm:"foreignKey:AssignUserID;references:ID"`
	TopTaskAssignUsers     []User         `gorm:"many2many:task_assign_users;"` // 顶级任务使用这个，即使只有一个人，也要使用这个
	TesterID               uint           `json:"tester_id"`
	Tester                 User           `json:"tester" gorm:"foreignKey:TesterID;references:ID"`
	ProjectID              uint           `json:"project_id" gorm:"not null,index"`
	Project                Project        `json:"project" gorm:"foreignKey:ProjectID;references:ID"`
	ParentID               uint           `json:"parent_id"` // 0表示顶级任务
	ParentTask             *Task          `json:"parent_task" gorm:"foreignKey:ParentID;references:ID"`
	HierarchyPath          string         `json:"hierarchy_path"` // 以冒号结尾
	Status                 TaskStatus
	StartedAt              *time.Time         `json:"started_at"`
	CompletedAt            *time.Time         `json:"completed_at"`
	ArchivedAt             *time.Time         `json:"archived_at"`
	Steps                  []Step             `json:"steps" gorm:"foreignKey:TaskID;references:ID"`               // 一对多关系
	TaskTimeEstimates      []TaskTimeEstimate `json:"task_time_estimates" gorm:"foreignKey:TaskID;references:ID"` // 一对多关系
}

type TaskTimeEstimate struct {
	ID                      uint           `json:"id" gorm:"primaryKey"`
	CreatedAt               time.Time      `json:"created_at"`
	UserID                  uint           `json:"user_id" gorm:"not null"`
	User                    User           `json:"user" gorm:"foreignKey:UserID;references:ID"`
	TaskID                  uint           `json:"task_id" gorm:"not null"`
	TaskDuration            *time.Duration `json:"task_duration"`             // 预计完成时间
	EstimatedCompletionTime *time.Time     `json:"estimated_completion_time"` // 预计完成时间
}

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
