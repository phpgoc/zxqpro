package entity

import "time"

type TaskTimeEstimate struct {
	ID                      uint           `json:"id" gorm:"primaryKey"`
	CreatedAt               time.Time      `json:"created_at"`
	UserID                  uint           `json:"user_id" gorm:"not null"`
	User                    User           `json:"user" gorm:"foreignKey:UserID;references:ID"`
	TaskID                  uint           `json:"task_id" gorm:"not null"`
	Comment                 string         `json:"comment"`                   // 备注
	TaskDuration            *time.Duration `json:"task_duration"`             // 预计完成时间
	EstimatedCompletionTime *time.Time     `json:"estimated_completion_time"` // 预计完成时间
	FreeTimeBegin           *time.Time     `json:"free_time_begin"`           // 空闲时间开始时间
	FreeTimeEnd             *time.Time     `json:"free_time_end"`             // 空闲时间结束时间
}
