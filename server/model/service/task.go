package service

import (
	"errors"
	"fmt"
	"time"

	"github.com/phpgoc/zxqpro/model/dao"
	"github.com/phpgoc/zxqpro/model/entity"
	"github.com/phpgoc/zxqpro/my_runtime"
	"github.com/phpgoc/zxqpro/routes/request"
)

func CanCreateTop(userId, projectId uint) bool {
	if IsAdmin(userId) {
		return true
	}
	roleTypeInProject, err := GetRoleType(userId, projectId)
	if err != nil {
		return false
	}
	return roleTypeInProject == entity.RoleTypeOwner || roleTypeInProject == entity.RoleTypeProducter
}

func CanAssignSelfToTop(userId, projectId uint) bool {
	if IsAdmin(userId) {
		return false
	}
	roleTypeInProject, err := GetRoleType(userId, projectId)
	if err != nil {
		return false
	}

	project, _ := dao.GetProjectById(projectId)
	if project.OwnerID == userId {
		return true
	}
	return project.Config.JoinBySelf && (roleTypeInProject == entity.RoleTypeOwner || roleTypeInProject == entity.RoleTypeProducter || roleTypeInProject == entity.RoleTypeDeveloper)
}

func CanCreateTopTask(userId, projectId uint) bool {
	if IsAdmin(userId) {
		return true
	}
	role, err := GetUserRoleInProject(userId, projectId)
	if err != nil {
		return false
	}
	if role.RoleType == entity.RoleTypeOwner || role.RoleType == entity.RoleTypeAdmin {
		return true
	} else {
		return false
	}
}

func CanBeAssignedDeveloper(userId, projectId uint) bool {
	if IsAdmin(userId) {
		return false
	}
	role, err := GetUserRoleInProject(userId, projectId)
	if err != nil {
		return false
	}
	if role.RoleType == entity.RoleTypeOwner || role.RoleType == entity.RoleTypeDeveloper || role.RoleType == entity.RoleTypeProducter {
		return true
	} else {
		return false
	}
}

func CanBeAssignedTester(userId, projectId uint) bool {
	if IsAdmin(userId) {
		return false
	}
	role, err := GetUserRoleInProject(userId, projectId)
	if err != nil {
		return false
	}
	if role.RoleType == entity.RoleTypeOwner || role.RoleType == entity.RoleTypeTester || role.RoleType == entity.RoleTypeProducter {
		return true
	} else {
		return false
	}
}

func TaskCreateTop(userId uint, req request.TaskCreateTop) error {
	var expectCompleteTime *time.Time = nil
	if req.ExpectCompleteTime != nil {
		t, err := time.Parse("2006-01-02", *req.ExpectCompleteTime)
		if err != nil {
			return errors.New("ExpectCompleteTime格式错误")
		}
		expectCompleteTime = &time.Time{}
		*expectCompleteTime = t
	}
	if !CanCreateTop(userId, req.ProjectId) {
		return errors.New("无权限")
	}
	//if req.AssignUsers == nil || len(req.AssignUsers) == 0 {
	//	return errors.New("AssignUsers不能为空")
	//}
	for _, u := range req.AssignUsers {
		if !CanBeAssignedDeveloper(u, req.ProjectId) {
			return errors.New(fmt.Sprintf("%d cannot be assigned to develper", u))
		}
	}

	if !CanBeAssignedTester(req.TesterId, req.ProjectId) {
		return errors.New(fmt.Sprintf("%d cannot be assigned to tester", req.TesterId))
	}
	var users []entity.User
	// 根据 user_id 查询对应的 User 对象
	if err := my_runtime.Db.Find(&users, req.AssignUsers).Error; err != nil {
		return err
	}

	task := entity.Task{
		Name:               req.Name,
		ProjectId:          req.ProjectId,
		Description:        req.Description,
		ParentId:           0,
		CreateUserId:       userId,
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
