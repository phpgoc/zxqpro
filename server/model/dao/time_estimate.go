package dao

import (
	"errors"

	"github.com/phpgoc/zxqpro/model/entity"
	"gorm.io/gorm"
)

type TaskTimeEstimateDAO struct {
	db *gorm.DB
}

// NewTaskTimeEstimateDAO 创建一个新的 entity.taskTimeEstimateDAO 实例
func NewTaskTimeEstimateDAO(db *gorm.DB) *TaskTimeEstimateDAO {
	return &TaskTimeEstimateDAO{
		db: db,
	}
}

// CreateTaskTimeEstimate 创建一个新的 entity.TaskTimeEstimate 记录
func (d *TaskTimeEstimateDAO) CreateTaskTimeEstimate(estimate *entity.TaskTimeEstimate) error {
	result := d.db.Create(estimate)
	return result.Error
}

// GetTaskTimeEstimateByID 根据 ID 获取 entity.TaskTimeEstimate 记录
func (d *TaskTimeEstimateDAO) GetTaskTimeEstimateByID(id uint) (*entity.TaskTimeEstimate, error) {
	var estimate entity.TaskTimeEstimate
	result := d.db.Preload("User").First(&estimate, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return &estimate, nil
}

func (d *TaskTimeEstimateDAO) GetTaskTimeEstimateCountByUserAndTask(userID, taskID uint) int64 {
	var count int64
	result := d.db.Model(&entity.TaskTimeEstimate{}).Where("user_id = ? AND task_id = ?", userID, taskID).Count(&count)
	if result.Error != nil {
		return 0
	}
	return count
}

// UpdateTaskTimeEstimate 更新 entity.TaskTimeEstimate 记录
//func (dao *taskTimeEstimateDAO) UpdateTaskTimeEstimate(estimate *entity.TaskTimeEstimate) error {
//	result := dao.db.Save(estimate)
//	return result.Error
//}

// DeleteTaskTimeEstimate 根据 ID 删除 entity.TaskTimeEstimate 记录
//func (dao *taskTimeEstimateDAO) DeleteTaskTimeEstimate(id uint) error {
//	result := dao.db.Delete(&entity.TaskTimeEstimate{}, id)
//	return result.Error
//}
