package utils

import (
	"githutb.com/phpgoc/zxqpro/server/orm_model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Db *gorm.DB

func InitDb() {
	var err error
	Db, err = gorm.Open(sqlite.Open("zxqpro.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	_ = Db.AutoMigrate(&orm_model.User{})
}
