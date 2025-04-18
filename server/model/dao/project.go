package dao

import (
	"errors"

	"github.com/phpgoc/zxqpro/model/entity"
	"github.com/phpgoc/zxqpro/my_runtime"
	"github.com/phpgoc/zxqpro/utils"
)

func GetUserRoleInProject(userId, projectId uint) (entity.Role, error) {
	res := entity.Role{}
	err := my_runtime.Db.Model(&entity.Role{}).Preload("Project").Where("project_id = ? and user_id = ?", projectId, userId).First(&res).Error
	if err != nil {
		return res, err
	} else {
		return res, nil
	}
}

func UpdateProject(projectId uint, project entity.Project) error {
	originalProject := entity.Project{}
	res := my_runtime.Db.Model(&entity.Project{}).Where("id = ?", projectId).First(&originalProject)
	if res.Error != nil {
		return res.Error
	}

	if originalProject.GitAddress == "" && project.GitAddress != "" {
		if !utils.IsGitRepository(project.GitAddress) {
			return errors.New("not a git repository")
		}
		// 如果新的GitAddress不为空，而原来的为空，做一些操作
		project.Status = entity.ProjectStatusActive
		my_runtime.GitPathList.Add(project.GitAddress)
	} else if originalProject.GitAddress != project.GitAddress {
		if !utils.IsGitRepository(project.GitAddress) {
			return errors.New("not a git repository")
		}
		my_runtime.GitPathList.Remove(originalProject.GitAddress)
		my_runtime.GitPathList.Add(project.GitAddress)
	}

	res = my_runtime.Db.Model(&entity.Project{}).Where("id = ?", projectId).Updates(project)
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected != 1 {
		return errors.New("not found")
	}
	return nil
}

func UpdateProjectStatus(projectId uint, status entity.ProjectStatus) error {
	// todo 如果要更新状态为已完成，检查是否所有的任务都已完成
	res := my_runtime.Db.Model(&entity.Project{}).Where("id = ?", projectId).Update("status", status)
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected != 1 {
		return errors.New("not found")
	}
	return nil
}

func GetOneProject(projectId uint) (entity.Project, error) {
	project := entity.Project{}
	res := my_runtime.Db.Preload("Owner").Where("id = ?", projectId).First(&project)
	if res.Error != nil {
		return project, res.Error
	}
	return project, nil
}
