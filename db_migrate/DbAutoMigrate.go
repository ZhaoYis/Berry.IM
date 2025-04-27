package db_migrate

import (
	"Berry_IM/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

type DbAutoMigrate struct {
}

func NewDbAutoMigrate() *DbAutoMigrate {
	return &DbAutoMigrate{}
}

// AutoMigrate 自动迁移数据库
func (dam *DbAutoMigrate) AutoMigrate() {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      true,        // Don't include params in the SQL log
			Colorful:                  false,       // Disable color
		},
	)

	dsn := "root:1qaz2wsx.com@tcp(127.0.0.1:3306)/berry_im?charset=utf8mb4&parseTime=True&loc=Local"
	var DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic("failed to connect database")
	}
	// 迁移 UserBasic
	err = DB.AutoMigrate(&models.UserBasic{})
	if err != nil {
		panic(err)
	}
}
