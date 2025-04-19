package service

import (
	"github.com/phpgoc/zxqpro/model/dao"
	"github.com/phpgoc/zxqpro/model/entity"
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
