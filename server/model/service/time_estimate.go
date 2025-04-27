package service

import (
	"errors"

	"github.com/phpgoc/zxqpro/model/dao"
	"github.com/phpgoc/zxqpro/model/entity"
)

type TaskTimeEstimateService struct {
	taskTimeEstimateDAO *dao.TaskTimeEstimateDAO
	taskDAO             *dao.TaskDAO
	projectService      *ProjectService
}

func NewTaskTimeEstimateService(taskTimeEstimateDAO *dao.TaskTimeEstimateDAO, taskDAO *dao.TaskDAO, projectService *ProjectService) *TaskTimeEstimateService {
	return &TaskTimeEstimateService{
		taskTimeEstimateDAO: taskTimeEstimateDAO,
		taskDAO:             taskDAO,
		projectService:      projectService,
	}
}

// CanEstimateTimeWithFree   任何人都可以评估，但是这个项目中的测试不能提供空闲时间
func (s *TaskTimeEstimateService) CanEstimateTimeWithFree(userID, taskID uint) bool {
	roleType := s.projectService.GetRoleType(userID, taskID)
	return roleType == entity.RoleTypeTester
}

// CanAddEstimateTime 估时不可修改，不可删除，最多添加3次,
// 不能对状态为完成，归档，失败的任务进行估时
func (s *TaskTimeEstimateService) CanAddEstimateTime(userID, taskID uint) error {
	task, err := s.taskDAO.GetTaskWithoutPreloadByID(taskID)
	if err != nil {
		return err
	}
	if task.Status == entity.TaskStatusCompleted || task.Status == entity.TaskStatusArchived || task.Status == entity.TaskStatusFailed {
		return errors.New("if task status is completed, archived or failed, you can not add time estimate")
	}
	if s.taskTimeEstimateDAO.GetTaskTimeEstimateCountByUserAndTask(userID, taskID) >= 3 {
		return errors.New("you can only add time estimate up to three times")
	}
	return nil
}
