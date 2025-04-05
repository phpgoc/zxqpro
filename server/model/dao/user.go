package dao

import (
	"github.com/phpgoc/zxqpro/model/entity"
)

func CreateUser(user *entity.User) error {
	result := Db.Create(&user)

	if result.RowsAffected != 1 {
		return result.Error
	}
	password := Md5Password(user.Password, user.ID)
	result = Db.Model(&user).Update("password", password)
	return result.Error
}

func GetUserById(id uint) (entity.User, error) {
	var user entity.User
	result := Db.First(&user, id)
	if result.Error != nil {
		return entity.User{}, result.Error
	}
	return user, nil
}
