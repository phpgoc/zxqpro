package request

import (
	"time"

	"github.com/phpgoc/zxqpro/model/entity"
)

type TaskCreateTop struct {
	ProjectID          uint    `json:"project_id" binding:"required,min=1" default:"1"`
	Name               string  `json:"name" binding:"required,max=20" default:"任务名称"`
	Description        string  `json:"description" binding:"required" default:"任务描述"`
	ExpectCompleteTime *string `json:"expect_complete_time" binding:"omitempty" default:"2006-01-02"`
	AssignUsers        []uint  `json:"assign_users"`
	TesterID           uint    `json:"tester_id" binding:"required,min=1" default:"1"`
}

type TaskUpdateTop struct {
	ID                 uint               `json:"id" binding:"required,min=1"`
	Name               *string            `json:"name" binding:"omitempty,max=20"`
	Description        *string            `json:"description" binding:"omitempty"`
	ExpectCompleteTime *string            `json:"expect_complete_time" binding:"omitempty"`
	AssignUsers        []uint             `json:"assign_users"  binding:"omitempty"`
	TesterID           *uint              `json:"tester_id"  binding:"omitempty,min=1"`
	Status             *entity.TaskStatus `json:"status" binding:"omitempty,min=1,max=5"`
}

type TaskCreateSub struct {
	Parent                 uint           `json:"id" binding:"required,min=1"`
	Name                   string         `json:"name" binding:"required,max=20"`
	Description            string         `json:"description" binding:"required"`
	ExpectCompleteDuration *time.Duration `json:"expect_complete_duration" binding:"omitempty"` // 预计完成时间
	ExpectCompleteTime     *time.Time     `json:"expect_complete_time" binding:"omitempty,datetime=2006-01-02 15:04:05"`
	TesterID               uint           `json:"tester_id" binding:"required,min=0"`
}
type TaskUpdateSub struct {
	Parent                 uint               `json:"id" binding:"required,min=1"`
	Name                   *string            `json:"name" binding:"required,max=20"`
	Description            *string            `json:"description" binding:"required"`
	ExpectCompleteDuration *time.Duration     `json:"expect_complete_duration" binding:"omitempty"` // 预计完成时间
	ExpectCompleteTime     *time.Time         `json:"expect_complete_time" binding:"omitempty,datetime=2006-01-02 15:04:05"`
	TesterID               *uint              `json:"tester_id" binding:"required,min=0"`
	Status                 *entity.TaskStatus `json:"status" binding:"min=1,max=5"`
}
