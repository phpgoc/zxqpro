package dao

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"fmt"
	"log"
	"time"

	"github.com/phpgoc/zxqpro/my_runtime"

	"gorm.io/gorm/logger"

	"github.com/phpgoc/zxqpro/model/entity"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	_ "modernc.org/sqlite"
)

func newSQLiteDialector(dsn string) gorm.Dialector {
	db, err := sql.Open("sqlite", dsn)
	if err != nil {
		panic(err)
	}
	return sqlite.Dialector{
		Conn: db,
	}
}

func InitDb() {
	newLogger := logger.New(
		log.New(my_runtime.GormLogWriter, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second,             // Slow SQL threshold
			LogLevel:                  my_runtime.GormLogLevel, // Log level
			IgnoreRecordNotFoundError: true,                    // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      false,
			Colorful:                  true, // Disable color
		},
	)
	var err error
	my_runtime.Db, err = gorm.Open(newSQLiteDialector("zxqpro.db"), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	_ = my_runtime.Db.AutoMigrate(&entity.User{})
	_ = my_runtime.Db.AutoMigrate(&entity.Project{})
	_ = my_runtime.Db.AutoMigrate(&entity.TaskTimeEstimate{}, &entity.Task{}, &entity.Step{})
	_ = my_runtime.Db.AutoMigrate(&entity.Role{})
	_ = my_runtime.Db.AutoMigrate(&entity.Message{}, &entity.MessageTo{})
	// 如果数据库没有数据就插入一条数据
	var count int64
	my_runtime.Db.Model(entity.User{}).Count(&count)
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
