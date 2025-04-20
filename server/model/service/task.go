package service

import (
	"errors"
	"fmt"
	"time"

	"github.com/phpgoc/zxqpro/routes/response"

	"github.com/phpgoc/zxqpro/model/dao"
	"github.com/phpgoc/zxqpro/model/entity"
	"github.com/phpgoc/zxqpro/my_runtime"
	"github.com/phpgoc/zxqpro/routes/request"
)

func CanCreateTop(userID, projectID uint) bool {
	if IsAdmin(userID) {
		return true
	}
	roleTypeInProject, err := GetRoleType(userID, projectID)
	if err != nil {
		return false
	}
	return roleTypeInProject == entity.RoleTypeOwner || roleTypeInProject == entity.RoleTypeProducter
}

func CanAssignSelfToTop(userID, projectID uint) bool {
	if IsAdmin(userID) {
		return false
	}
	roleTypeInProject, err := GetRoleType(userID, projectID)
	if err != nil {
		return false
	}

	project, _ := dao.GetProjectByID(projectID)
	if project.OwnerID == userID {
		return true
	}
	return project.Config.JoinBySelf && (roleTypeInProject == entity.RoleTypeOwner || roleTypeInProject == entity.RoleTypeProducter || roleTypeInProject == entity.RoleTypeDeveloper)
}

func CanCreateTopTask(userID, projectID uint) bool {
	if IsAdmin(userID) {
		return true
	}
	role, err := GetUserRoleInProject(userID, projectID)
	if err != nil {
		return false
	}
	if role.RoleType == entity.RoleTypeOwner || role.RoleType == entity.RoleTypeAdmin {
		return true
	} else {
		return false
	}
}

func CanBeAssignedDeveloper(userID, projectID uint) bool {
	if IsAdmin(userID) {
		return false
	}
	role, err := GetUserRoleInProject(userID, projectID)
	if err != nil {
		return false
	}
	if role.RoleType == entity.RoleTypeOwner || role.RoleType == entity.RoleTypeDeveloper || role.RoleType == entity.RoleTypeProducter {
		return true
	} else {
		return false
	}
}

func CanBeAssignedTester(userID, projectID uint) bool {
	if IsAdmin(userID) {
		return false
	}
	role, err := GetUserRoleInProject(userID, projectID)
	if err != nil {
		return false
	}
	if role.RoleType == entity.RoleTypeOwner || role.RoleType == entity.RoleTypeTester || role.RoleType == entity.RoleTypeProducter {
		return true
	} else {
		return false
	}
}

func TaskCreateTop(userID uint, req request.TaskCreateTop) error {
	var expectCompleteTime *time.Time = nil
	if req.ExpectCompleteTime != nil {
		t, err := time.Parse("2006-01-02", *req.ExpectCompleteTime)
		if err != nil {
			return errors.New("ExpectCompleteTime格式错误")
		}
		expectCompleteTime = &time.Time{}
		*expectCompleteTime = t
	}
	if !CanCreateTop(userID, req.ProjectID) {
		return errors.New("无权限")
	}
	//if req.AssignUsers == nil || len(req.AssignUsers) == 0 {
	//	return errors.New("AssignUsers不能为空")
	//}
	for _, u := range req.AssignUsers {
		if !CanBeAssignedDeveloper(u, req.ProjectID) {
			return errors.New(fmt.Sprintf("%d cannot be assigned to develper", u))
		}
	}

	if !CanBeAssignedTester(req.TesterID, req.ProjectID) {
		return errors.New(fmt.Sprintf("%d cannot be assigned to tester", req.TesterID))
	}
	var users []entity.User
	// 根据 user_id 查询对应的 User 对象
	if err := my_runtime.Db.Find(&users, req.AssignUsers).Error; err != nil {
		return err
	}

	task := entity.Task{
		Name:               req.Name,
		ProjectID:          req.ProjectID,
		Description:        req.Description,
		ParentID:           0,
		CreateUserID:       userID,
		TopTaskAssignUsers: users,
		ExpectCompleteTime: expectCompleteTime,
	}
	err := my_runtime.Db.Create(&task).Error
	if err != nil {
		return err
	}
	task.HierarchyPath = fmt.Sprintf("%d:", task.ID)
	err = my_runtime.Db.Save(&task).Error
	return err
}

func TaskUpdateTop(id uint, req request.TaskUpdateTop) error {
	return nil
}

func TaskInfo(id uint) (response.TaskInfo, error) {
	// 没有权限问题，任何人都可以查看任务信息
	var task entity.Task
	err := my_runtime.Db.Preload("CreateUser").Preload("AssignUser").Preload("Tester").Preload("Steps").Preload("Steps.Developer").Preload("Project").Preload("TopTaskAssignUsers").Where("id = ?", id).First(&task).Error
	if err != nil {
		return response.TaskInfo{}, err
	}
	var taskTimeEstimates []entity.TaskTimeEstimate
	err = my_runtime.Db.Where("task_id = ?", id).Find(&taskTimeEstimates).Error
	if err != nil {
		return response.TaskInfo{}, err
	}

	var testUser *response.CommonIDAndName = nil
	if task.TesterID != 0 {
		testUser = &response.CommonIDAndName{
			ID:   task.TesterID,
			Name: task.Tester.Name,
		}
	}
	var subAssignUser *response.CommonIDAndName = nil
	if task.AssignUserID != 0 {
		subAssignUser = &response.CommonIDAndName{
			ID:   task.AssignUserID,
			Name: task.AssignUser.Name,
		}
	}
	var topTaskAssignUsers []response.CommonIDAndName = make([]response.CommonIDAndName, len(task.TopTaskAssignUsers))
	for i := range task.TopTaskAssignUsers {
		topTaskAssignUsers[i] = response.CommonIDAndName{
			ID:   task.TopTaskAssignUsers[i].ID,
			Name: task.TopTaskAssignUsers[i].Name,
		}
	}

	taskInfo := response.TaskInfo{
		ProjectID:   task.ProjectID,
		ParentID:    task.ParentID,
		Name:        task.Name,
		Description: task.Description,
		Status:      task.Status,
		CreateUser: response.CommonIDAndName{
			ID:   task.CreateUserID,
			Name: task.CreateUser.Name,
		},
		ExpectCompleteTime:   task.ExpectCompleteTime,
		TestUser:             testUser,
		SubAssignUser:        subAssignUser,
		TopTaskAssignUsers:   topTaskAssignUsers,
		StartedAt:            task.StartedAt,
		CompletedAt:          task.CompletedAt,
		ArchivedAt:           task.ArchivedAt,
		TaskTimeEstimateList: task.TaskTimeEstimates,
		StepList:             task.Steps,
	}
	return taskInfo, nil
}
