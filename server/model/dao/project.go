package dao

import (
	"errors"

	"github.com/phpgoc/zxqpro/model/entity"
	"github.com/phpgoc/zxqpro/my_runtime"
)

func UpdateProjectStatus(projectID uint, status entity.ProjectStatus) error {
	// todo 如果要更新状态为已完成，检查是否所有的任务都已完成
	res := my_runtime.Db.Model(&entity.Project{}).Where("id = ?", projectID).Update("status", status)
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected != 1 {
		return errors.New("not found")
	}
	return nil
}

func GetProjectByID(projectID uint) (entity.Project, error) {
	project := entity.Project{}
	res := my_runtime.Db.Preload("Owner").Where("id = ?", projectID).First(&project)
	if res.Error != nil {
		return project, res.Error
	}
	return project, nil
}

func GetProjectsForAdmin(status byte, page, pageSize int) ([]entity.Project, int64, error) {
	var projects []entity.Project
	var total int64
	model := my_runtime.Db.Model(&entity.Project{}).Preload("Owner")
	if status != 0 {
		model = model.Where("status = ?", status)
	}
	err := model.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}
	err = model.Offset((page - 1) * pageSize).Limit(pageSize).Find(&projects).Error
	return projects, total, err
}

func GetProjectsForUser(userID uint, status, roleType byte, page, pageSize int) ([]entity.Role, int64, error) {
	var roles []entity.Role
	var total int64
	model := my_runtime.Db.Model(&entity.Role{}).Preload("Project").Where("user_id = ?", userID).Preload("Project.Owner")
	if status != 0 {
		model = model.Where("status = ?", status)
	}
	if roleType != 0 {
		model = model.Where("role_type = ?", roleType)
	}
	err := model.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}
	err = model.Offset((page - 1) * pageSize).Limit(pageSize).Find(&roles).Error
	return roles, total, err
}
