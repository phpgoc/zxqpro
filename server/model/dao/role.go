package dao

import "github.com/phpgoc/zxqpro/model/entity"

func CreateRole(userId, roleId uint, roleType entity.RoleType) error {
	role := entity.Role{
		UserID:    userId,
		ProjectID: roleId,
		RoleType:  roleType,
	}
	result := Db.Model(&entity.Role{}).Where("user_id = ? and project_id = ?", userId, roleId).UpdateColumns(map[string]interface{}{
		"deleted_at": nil,
		"role_type":  roleType,
	})
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected > 0 {
		return nil // 如果存在并且已经设置了 delete_at 不为 null，则更新成功，直接返回 nil。
	}
	result = Db.Create(&role)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func DeleteRole(userId, projectId uint) error {
	role := entity.Role{
		UserID:    userId,
		ProjectID: projectId,
	}
	result := Db.Delete(&role)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func UpdateRole(userId, projectId uint, roleType entity.RoleType) error {
	role := entity.Role{
		UserID:    userId,
		ProjectID: projectId,
	}
	result := Db.Model(&role).Where("user_id = ? and project_id = ?", userId, projectId).Update("role_type", roleType)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
