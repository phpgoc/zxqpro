package service

import (
	"errors"
	"fmt"
	"time"

	"github.com/phpgoc/zxqpro/model/dao"

	"github.com/phpgoc/zxqpro/utils"

	"github.com/phpgoc/zxqpro/routes/response"

	"github.com/phpgoc/zxqpro/model/entity"
	"github.com/phpgoc/zxqpro/my_runtime"
	"github.com/phpgoc/zxqpro/routes/request"
)

type TaskService struct {
	taskDAO        *dao.TaskDAO
	projectService *ProjectService
}

func NewTaskService(taskDAO *dao.TaskDAO, projectService *ProjectService) *TaskService {
	return &TaskService{
		taskDAO:        taskDAO,
		projectService: projectService,
	}
}

func (s *TaskService) CanCreateTop(userID, projectID uint) bool {
	if IsAdmin(userID) {
		return true
	}
	roleTypeInProject, err := s.projectService.GetRoleType(userID, projectID)
	if err != nil {
		return false
	}
	return roleTypeInProject == entity.RoleTypeOwner || roleTypeInProject == entity.RoleTypeProducter
}

func (s *TaskService) CanBeAssignedDeveloper(userID, projectID uint) bool {
	if IsAdmin(userID) {
		return false
	}
	role, err := s.projectService.GetRoleType(userID, projectID)
	if err != nil {
		return false
	}
	if role == entity.RoleTypeOwner || role == entity.RoleTypeDeveloper || role == entity.RoleTypeProducter {
		return true
	} else {
		return false
	}
}

func (s *TaskService) CanBeAssignedTester(userID, projectID uint) bool {
	if IsAdmin(userID) {
		return false
	}
	role, err := s.projectService.GetRoleType(userID, projectID)
	if err != nil {
		return false
	}
	if role == entity.RoleTypeOwner || role == entity.RoleTypeTester || role == entity.RoleTypeProducter {
		return true
	} else {
		return false
	}
}

// GetChildrenTaskList 只包括儿子，不包括孙子
func (s *TaskService) GetChildrenTaskList(taskID uint) ([]response.TaskOneForList, error) {
	var tasks []entity.Task
	err := my_runtime.Db.Preload("CreateUser").Preload("TestUser").Preload("TopTaskAssignUsers").Where("parent_id = ?", taskID).Find(&tasks).Error
	if err != nil {
		return nil, err
	}
	var testUser *response.CommonIDAndName
	var topTaskAssignUsers []response.CommonIDAndName
	var subAssignUser *response.CommonIDAndName
	var taskList []response.TaskOneForList
	for _, task := range tasks {
		testUser = nil
		subAssignUser = nil
		topTaskAssignUsers = make([]response.CommonIDAndName, 0)
		if task.TesterID != 0 {
			testUser = &response.CommonIDAndName{
				ID:   task.TesterID,
				Name: task.Tester.Name,
			}
		}
		if task.AssignUserID != 0 {
			subAssignUser = &response.CommonIDAndName{
				ID:   task.AssignUserID,
				Name: task.AssignUser.Name,
			}
		}
		for _, u := range task.TopTaskAssignUsers {
			if !s.CanBeAssignedDeveloper(u.ID, task.ID) {
				return nil, errors.New(fmt.Sprintf("%d cannot be assigned to develper", u.ID))
			}
			topTaskAssignUsers = append(topTaskAssignUsers, response.CommonIDAndName{
				ID:   u.ID,
				Name: u.Name,
			})
		}

		taskList = append(taskList, response.TaskOneForList{
			ID:   task.ID,
			Name: task.Name,
			CreateUser: response.CommonIDAndName{
				ID:   task.CreateUserID,
				Name: task.CreateUser.Name,
			},
			Status:             task.Status,
			TestUser:           testUser,
			CompletedAt:        task.CompletedAt,
			SubAssignUser:      subAssignUser,
			TopTaskAssignUsers: topTaskAssignUsers,
			ExpectCompleteTime: task.ExpectCompleteTime,
		})
	}
	return taskList, nil
}

func (s *TaskService) TaskCreateTop(userID uint, req request.TaskCreateTop) error {
	var expectCompleteTime *time.Time = nil
	if req.ExpectCompleteTime != nil {
		t, err := time.Parse("2006-01-02", *req.ExpectCompleteTime)
		if err != nil {
			return errors.New("ExpectCompleteTime格式错误")
		}
		expectCompleteTime = &time.Time{}
		*expectCompleteTime = t
	}
	if !s.CanCreateTop(userID, req.ProjectID) {
		return errors.New("无权限")
	}
	// 可以为空的，为空也不会进入循环
	for _, u := range req.AssignUsers {
		if !s.CanBeAssignedDeveloper(u, req.ProjectID) {
			return errors.New(fmt.Sprintf("%d cannot be assigned to develper", u))
		}
	}

	if !s.CanBeAssignedTester(req.TesterID, req.ProjectID) {
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
		Status:             entity.TaskStatusStarted, // 顶级任务没有创建后的状态，直接就开始了，非顶级任务才有
	}
	err := my_runtime.Db.Create(&task).Error
	if err != nil {
		return err
	}
	task.HierarchyPath = fmt.Sprintf("%d:", task.ID)
	err = my_runtime.Db.Save(&task).Error
	return err
}

func (s *TaskService) TaskUpdateTop(userID uint, req request.TaskUpdateTop) error {
	var expectCompleteTime *time.Time = nil
	if req.ExpectCompleteTime != nil {
		t, err := time.Parse("2006-01-02", *req.ExpectCompleteTime)
		if err != nil {
			return errors.New("ExpectCompleteTime格式错误")
		}
		expectCompleteTime = &time.Time{}
		*expectCompleteTime = t
	}
	var task entity.Task
	var err error
	if err = my_runtime.Db.Preload("TopTaskAssignUsers").Where("id = ?", req.ID).First(&task).Error; err != nil {
		return err
	}
	if !IsAdmin(userID) && task.CreateUserID != userID {
		return errors.New("you are not the owner of this task")
	}
	if task.ExpectCompleteTime != nil && req.ExpectCompleteTime != nil {
		return errors.New("can not update expect_complete_time after first set")
	}
	if task.ExpectCompleteDuration != nil && req.ExpectCompleteTime != nil {
		return errors.New("can not update expect_complete_duration after first set")
	}
	updateTask := make(map[string]interface{})
	if req.Name != nil {
		updateTask["Name"] = *req.Name
	}
	if req.Description != nil {
		updateTask["Description"] = *req.Description
	}
	if req.ExpectCompleteTime != nil {
		updateTask["ExpectCompleteTime"] = expectCompleteTime
	}
	if req.Status != nil {
		if *req.Status == entity.TaskStatusCompleted || *req.Status == entity.TaskStatusArchived {
			// 空的逻辑正常，可以完成或关闭吧
			childTasks, _ := s.GetChildrenTaskList(task.ID)
			allArchived := true
			for _, childTask := range childTasks {
				allArchived = allArchived && childTask.Status == entity.TaskStatusArchived
			}
			if !allArchived {
				return errors.New("there are child tasks that are not archived")
			}
		}
		updateTask["Status"] = *req.Status
	}

	if req.TesterID != nil {
		if !s.CanBeAssignedTester(*req.TesterID, task.ProjectID) {
			return errors.New(fmt.Sprintf("%d cannot be assigned to tester", *req.TesterID))
		}
		updateTask["TesterID"] = *req.TesterID
	}
	if req.AssignUsers != nil {
		for _, u := range req.AssignUsers {
			if !s.CanBeAssignedDeveloper(u, task.ProjectID) {
				return errors.New(fmt.Sprintf("%d cannot be assigned to develper", u))
			}
		}
		var users []entity.User
		// 根据 user_id 查询对应的 User 对象
		if err := my_runtime.Db.Find(&users, req.AssignUsers).Error; err != nil {
			return err
		}
		updateTask["TopTaskAssignUsers"] = users
	}
	if err := my_runtime.Db.Model(&task).Updates(updateTask).Error; err != nil {
		return err
	}

	return nil
}

func (s *TaskService) TaskInfo(id uint) (response.TaskInfo, error) {
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

func (s *TaskService) TaskAssignSelfToTop(userID, TaskID uint) error {
	if IsAdmin(userID) {
		return errors.New("admin can not assign self to top task")
	}

	var task entity.Task
	if err := my_runtime.Db.Preload("Project").Preload("TopTaskAssignUsers").First(&task, TaskID).Error; err != nil {
		return err
	}
	if task.ParentID != 0 {
		return errors.New("task is not top task")
	}

	if task.Project.Config.JoinBySelf == false {
		return errors.New("project self assign to top task is disabled")
	}

	roleTypeInProject, err := s.projectService.GetRoleType(userID, task.ProjectID)
	if err != nil {
		return err
	}

	// 所有者和 产品经理都可以将自己分配到顶级任务中。 create user 也可以
	// create user 理论不该显示这个接口，但是为了兼容性，这里不做限制。
	if roleTypeInProject == entity.RoleTypeNone || roleTypeInProject == entity.RoleTypeTester {
		return errors.New("you are not allowed to assign self to top task")
	}
	var userIDs []uint
	for _, user := range task.TopTaskAssignUsers {
		userIDs = append(userIDs, user.ID)
	}
	if utils.Contains(userIDs, userID) {
		return errors.New("you have already assigned to this task")
	}
	var user entity.User
	_ = my_runtime.Db.First(&user, userID).Error
	task.TopTaskAssignUsers = append(task.TopTaskAssignUsers, user)
	if err = my_runtime.Db.Save(&task).Error; err != nil {
		return err
	}

	return nil
}
