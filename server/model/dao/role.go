package dao

import (
	"github.com/phpgoc/zxqpro/model/entity"
	"github.com/phpgoc/zxqpro/my_runtime"
	"gorm.io/gorm"
)

type RoleDAO struct {
	db *gorm.DB
}

// NewRoleDAO 创建一个新的 RoleDAO 实例
func NewRoleDAO(db *gorm.DB) *RoleDAO {
	return &RoleDAO{db: db}
}

func (d *RoleDAO) GetAllUserByProjectID(projectID uint) ([]entity.Role, error) {
	var roles []entity.Role
	result := d.db.Preload("User").Where("project_id = ?", projectID).Find(&roles)
	if result.Error != nil {
		return nil, result.Error
	}
	return roles, nil
}

func (d *RoleDAO) GetRole(userID, projectID uint) (entity.Role, error) {
	role := entity.Role{}
	res := my_runtime.Db.Where("user_id = ? and project_id = ?", userID, projectID).First(&role).Error
	return role, res
}

func CreateRole(userID, roleID uint, roleType entity.RoleType) error {
	role := entity.Role{
		UserID:    userID,
		ProjectID: roleID,
		RoleType:  roleType,
	}
	result := my_runtime.Db.Model(&entity.Role{}).Where("user_id = ? and project_id = ?", userID, roleID).UpdateColumns(map[string]interface{}{
		"deleted_at": nil,
		"role_type":  roleType,
	})
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected > 0 {
		return nil // 如果存在并且已经设置了 delete_at 不为 null，则更新成功，直接返回 nil。
	}
	result = my_runtime.Db.Create(&role)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func DeleteRole(userID, projectID uint) error {
	role := entity.Role{
		UserID:    userID,
		ProjectID: projectID,
	}
	result := my_runtime.Db.Delete(&role)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func UpdateRole(userID, projectID uint, roleType entity.RoleType) error {
	role := entity.Role{
		UserID:    userID,
		ProjectID: projectID,
	}
	result := my_runtime.Db.Model(&role).Where("user_id = ? and project_id = ?", userID, projectID).Update("role_type", roleType)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
