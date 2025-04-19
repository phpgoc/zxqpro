package request

import (
	"time"

	"github.com/phpgoc/zxqpro/model/entity"
)

type TaskCreateTop struct {
	ProjectId          uint       `json:"project_id" binding:"required;min=1"`
	Name               string     `json:"name" binding:"required;max=20"`
	Description        string     `json:"description" binding:"required"`
	ExpectCompleteTime *time.Time `json:"expect_complete_time" binding:"required;datetime=2006-01-02 15:04:05"`
	AssignUsers        []uint     `json:"assign_users"`
	TesterId           uint       `json:"tester_id" binding:"required;min=1"`
}

type TaskUpdateTop struct {
	Id                 uint               `json:"id" binding:"required;min=1"`
	Name               *string            `json:"name" binding:"required;max=20"`
	Description        *string            `json:"description" binding:"required"`
	ExpectCompleteTime *time.Time         `json:"expect_complete_time" binding:"omitempty;datetime=2006-01-02 15:04:05"`
	AssignUsers        []uint             `json:"assign_users"`
	TesterId           *uint              `json:"tester_id" binding:"required;min=1"`
	Status             *entity.TaskStatus `json:"status" binding:"min=1;max=5"`
}

type TaskCreateSub struct {
	Parent                 uint           `json:"id" binding:"required;min=1"`
	Name                   string         `json:"name" binding:"required;max=20"`
	Description            string         `json:"description" binding:"required"`
	ExpectCompleteDuration *time.Duration `json:"expect_complete_duration" binding:"omitempty"` // 预计完成时间
	ExpectCompleteTime     *time.Time     `json:"expect_complete_time" binding:"omitempty;datetime=2006-01-02 15:04:05"`
	TesterId               uint           `json:"tester_id" binding:"required;min=0"`
}
type TaskUpdateSub struct {
	Parent                 uint               `json:"id" binding:"required;min=1"`
	Name                   *string            `json:"name" binding:"required;max=20"`
	Description            *string            `json:"description" binding:"required"`
	ExpectCompleteDuration *time.Duration     `json:"expect_complete_duration" binding:"omitempty"` // 预计完成时间
	ExpectCompleteTime     *time.Time         `json:"expect_complete_time" binding:"omitempty;datetime=2006-01-02 15:04:05"`
	TesterId               *uint              `json:"tester_id" binding:"required;min=0"`
	Status                 *entity.TaskStatus `json:"status" binding:"min=1;max=5"`
}

type TaskTimeEstimateCreate struct {
	TaskId                  uint           `json:"task_id" binding:"required;min=1"`
	TaskDuration            *time.Duration `json:"task_duration"`             // 预计完成时间
	EstimatedCompletionTime *time.Time     `json:"estimated_completion_time"` // 预计完成时间
}
