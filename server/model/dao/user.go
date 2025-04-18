package dao

import (
	"github.com/phpgoc/zxqpro/model/entity"
	"github.com/phpgoc/zxqpro/my_runtime"
)

func CreateUser(user *entity.User) error {
	result := my_runtime.Db.Create(&user)

	if result.RowsAffected != 1 {
		return result.Error
	}
	password := Md5Password(user.Password, user.ID)
	result = my_runtime.Db.Model(&user).Update("password", password)
	return result.Error
}

func GetUserById(id uint) (entity.User, error) {
	var user entity.User
	result := my_runtime.Db.First(&user, id)
	if result.Error != nil {
		return entity.User{}, result.Error
	}
	return user, nil
}

func IsAdmin(userId uint) bool {
	return userId == 1
}
