package dao

import "github.com/phpgoc/zxqpro/model/entity"

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
