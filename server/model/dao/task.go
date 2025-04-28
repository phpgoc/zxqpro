package dao

import (
	"errors"

	"github.com/phpgoc/zxqpro/routes/request"
	"gorm.io/gorm/clause"

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

// Create 创建一个新的任务
func (d *TaskDAO) Create(task *entity.Task) error {
	result := d.db.Create(task)
	return result.Error
}

// GetByID 根据 ID 获取任务
func (d *TaskDAO) GetByID(id uint) (*entity.Task, error) {
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

// GetWithoutPreloadByID 根据 ID 获取任务，但不预加载任何关联
func (d *TaskDAO) GetWithoutPreloadByID(id uint) (*entity.Task, error) {
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

// Update 更新任务信息
func (d *TaskDAO) Update(task *entity.Task) error {
	result := d.db.Save(task)
	return result.Error
}

func (d *TaskDAO) Updates(task *entity.Task, values interface{}) error {
	result := d.db.Model(task).Updates(values)
	return result.Error
}

// Delete 根据 ID 删除任务
func (d *TaskDAO) Delete(id uint) error {
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

func (d *TaskDAO) GetProjectTaskListAndCount(req request.ProjectTaskList) (int64, []entity.Task, error) {
	var total int64
	var taskList []entity.Task

	model := d.db.Model(&entity.Task{}).Preload(clause.Associations).Where("project_id = ?", req.ID)
	if req.CreateUserID != 0 {
		model = model.Where("create_user_id = ?", req.CreateUserID)
	}
	if req.Status != 0 {
		model = model.Where("status = ?", req.Status)
	}
	if req.TopStatus != 0 {
		if req.TopStatus == 1 {
			model = model.Where("parent_id = ?", 0)
		} else {
			model = model.Where("parent_id != ?", 0)
		}
	}
	for _, order := range req.OrderList {
		if order.Desc {
			model = model.Order(order.Field + " desc")
		} else {
			model = model.Order(order.Field + " asc")
		}
	}

	err := model.Count(&total).Error
	if err != nil {
		return 0, nil, err
	}

	err = model.Limit(req.PageSize).Offset((req.Page - 1) * req.PageSize).Find(&taskList).Error
	return total, taskList, err
}
