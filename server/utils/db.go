package utils

import (
	"context"

	"github.com/phpgoc/zxqpro/my_runtime"
	"gorm.io/gorm"
)

// Transaction 使用这个函数调用DAO层时，不要使用Service本身的DAO，而要使用闭包函数传入的db来初始化DAO
func Transaction(c context.Context, dbFunc func(tx *gorm.DB) error) error {
	var tx *gorm.DB
	if c != nil {
		tx = my_runtime.Db.WithContext(c).Begin()
	} else {
		tx = my_runtime.Db.Begin()
	}
	defer func() {
		if r := recover(); r != nil {
			if rollbackErr := tx.Rollback().Error; rollbackErr != nil {
				LogErrorWithUpLevel(rollbackErr.Error(), 2)
			}
		}
	}()

	if err := dbFunc(tx); err != nil {
		if rollbackErr := tx.Rollback().Error; rollbackErr != nil {
			LogWarnWithUpLevel(rollbackErr.Error(), 1)
		}
		LogInfoWithUpLevel(err.Error(), 1)
		return err
	}

	if err := tx.Commit().Error; err != nil {
		LogErrorWithUpLevel(err.Error(), 1)
		return err
	}

	return nil
}
