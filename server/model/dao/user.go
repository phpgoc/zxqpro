package dao

import (
	"github.com/phpgoc/zxqpro/model/entity"
	"github.com/phpgoc/zxqpro/utils"
)

func GetUserById(id uint) (entity.User, error) {
	var user entity.User
	result := utils.Db.First(&user, id)
	if result.Error != nil {
		return entity.User{}, result.Error
	}
	return user, nil
}
