package dao

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"

	"github.com/phpgoc/zxqpro/model/entity"
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
	_ = Db.AutoMigrate(&entity.User{})
	_ = Db.AutoMigrate(&entity.Project{})
	// 如果数据库没有数据就插入一条数据
	var count int64
	Db.Model(entity.User{}).Count(&count)
	if count == 0 {
		defaultAdmin := entity.User{
			Name:     "admin",
			Password: "Aa123456",
		}
		_ = CreateUser(&defaultAdmin)
	}
}

func Md5Password(password string, id uint) string {
	combined := fmt.Sprintf("%s%d", password, id)
	// 计算 MD5 哈希值
	hash := md5.Sum([]byte(combined))
	// 将哈希值转换为十六进制字符串
	return hex.EncodeToString(hash[:])
}
