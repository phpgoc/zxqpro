package dao

import (
	"errors"

	"github.com/phpgoc/zxqpro/model/entity"
	"gorm.io/gorm"
)

type TaskDAO struct {
	db *gorm.DB
}

// NewTaskDAO 创建一个新的 TaskDAO 实例
func NewTaskDAO(db *gorm.DB) *TaskDAO {
	return &TaskDAO{
		db: db,
	}
}

// CreateTask 创建一个新的任务
func (d *TaskDAO) CreateTask(task *entity.Task) error {
	result := d.db.Create(task)
	return result.Error
}

// GetTaskByID 根据 ID 获取任务
func (d *TaskDAO) GetTaskByID(id uint) (*entity.Task, error) {
	var task entity.Task
	result := d.db.Preload("Create").Preload("AssignUser").Preload("Tester").Preload("Project").Preload("ParentTask").Preload("Steps").Preload("TaskTimeEstimates").Preload("TopTaskAssignUsers").First(&task, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return &task, nil
}

// GetTaskWithoutPreloadByID 根据 ID 获取任务，但不预加载任何关联
func (d *TaskDAO) GetTaskWithoutPreloadByID(id uint) (*entity.Task, error) {
	var task entity.Task
	result := d.db.First(&task, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return &task, nil
}

// UpdateTask 更新任务信息
func (d *TaskDAO) UpdateTask(task *entity.Task) error {
	result := d.db.Save(task)
	return result.Error
}

// DeleteTask 根据 ID 删除任务
func (d *TaskDAO) DeleteTask(id uint) error {
	result := d.db.Delete(&entity.Task{}, id)
	return result.Error
}

func (d *TaskDAO) GetChildrenTasksByParentID(parentID uint) ([]entity.Task, error) {
	var tasks []entity.Task
	// 预加载所有需要的关联数据（与原Service中的Preload一致）
	err := d.db.Preload("Create").Preload("Tester").Preload("AssignUser").Preload("TopTaskAssignUsers").
		Where("parent_id = ?", parentID).Find(&tasks).Error
	return tasks, err
}
