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
	Name               string     `json:"name" gorm:"not null"`
	Description        string     `json:"description" gorm:"type:text"`
	CreateUserId       uint       `json:"create_user_id"`
	CreateUser         User       `json:"create_user" gorm:"foreignKey:CreateUserId;references:ID"`
	ExpectCompleteTime *time.Time `json:"expect_complete_time"`
	AssignUserId       uint       `json:"assign_user_id"` // 如果是顶级任务，不使用这个字典，非顶级任务，指定的用户是一对一的。顶级任务是一对多的
	AssignUser         User       `json:"assign_user" gorm:"foreignKey:AssignUserId;references:ID"`
	TopTaskAssignUsers []User     `gorm:"many2many:task_assign_users;"` // 顶级任务使用这个，即使只有一个人，也要使用这个
	TesterId           uint       `json:"tester_id"`
	Tester             User       `json:"tester" gorm:"foreignKey:TesterId;references:ID"`
	ProjectId          uint       `json:"project_id" gorm:"not null,index"`
	Project            Project    `json:"project" gorm:"foreignKey:ProjectId;references:ID"`
	ParentId           uint       `json:"parent_id"` // 0表示顶级任务
	ParentTask         *Task      `json:"parent_task" gorm:"foreignKey:ParentId;references:ID"`
	HierarchyPath      uint       `json:"hierarchy_path"` // 以冒号结尾
	Status             TaskStatus
	StartedAt          time.Time          `json:"started_at"`
	CompletedAt        time.Time          `json:"completed_at"`
	ArchivedAt         time.Time          `json:"archived_at"`
	Steps              []Step             `json:"steps" gorm:"foreignKey:TaskId;references:ID"`               // 一对多关系
	TaskTimeEstimates  []TaskTimeEstimate `json:"task_time_estimates" gorm:"foreignKey:TaskId;references:ID"` // 一对多关系
}

type TaskTimeEstimate struct {
	gorm.Model
	UserId                  uint           `json:"user_id" gorm:"not null"`
	User                    User           `json:"user" gorm:"foreignKey:UserId;references:ID"`
	TaskId                  uint           `json:"task_id" gorm:"not null"`
	TaskDuration            *time.Duration `json:"task_duration"`             // 预计完成时间
	EstimatedCompletionTime *time.Time     `json:"estimated_completion_time"` // 预计完成时间
}

type Step struct {
	gorm.Model
	TaskId      uint     `json:"task_id" gorm:"not null"`
	Description string   `json:"description"`
	Percent     *byte    `json:"percent"`
	DeveloperId uint     `json:"developer_id"`
	Developer   User     `json:"developer" gorm:"foreignKey:DeveloperId;references:ID"`
	TesterId    uint     `json:"tester_id"`
	Tester      *User    `json:"tester" gorm:"foreignKey:TesterId;references:ID"`
	Commits     []string `json:"commits" gorm:"type:text"` // 提交的commit id
}
