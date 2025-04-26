package dao

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"

	"github.com/phpgoc/zxqpro/model/entity"
	"github.com/phpgoc/zxqpro/my_runtime"
	"gorm.io/gorm"
)

type UserDAO struct {
	db *gorm.DB
}

func NewUserDAO(db *gorm.DB) *UserDAO {
	return &UserDAO{
		db: db,
	}
}

func (d *UserDAO) GetByID(id uint) (entity.User, error) {
	var user entity.User
	result := d.db.First(&user, id)
	if result.Error != nil {
		return entity.User{}, result.Error
	}
	return user, nil
}

func (d *UserDAO) Create(user *entity.User) error {
	result := d.db.Create(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (d *UserDAO) UpdatePassword(user *entity.User, password string) error {
	return d.db.Model(user).Update("password", password).Error
}

func (d *UserDAO) Update(user *entity.User) error {
	result := d.db.Save(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (d *UserDAO) Sha1Password(password string, id uint) string {
	combined := fmt.Sprintf("%s%d", password, id)
	// 计算 MD5 哈希值
	hash := sha1.Sum([]byte(combined))
	// 将哈希值转换为十六进制字符串
	return hex.EncodeToString(hash[:])
}

func GetUserByID(id uint) (entity.User, error) {
	var user entity.User
	result := my_runtime.Db.First(&user, id)
	if result.Error != nil {
		return entity.User{}, result.Error
	}
	return user, nil
}

func (d *UserDAO) GetByEntity(user *entity.User) (entity.User, error) {
	result := my_runtime.Db.Where(user).First(user)
	if result.Error != nil {
		return entity.User{}, result.Error
	}
	return *user, nil
}

// ListUsers 列出无项目关联的用户
func (d *UserDAO) ListUsers(includeAdmin bool) ([]entity.User, error) {
	var users []entity.User
	model := d.db.Model(&entity.User{})
	if !includeAdmin {
		model = model.Where("id != ?", 1)
	}
	err := model.Find(&users).Error
	return users, err
}
