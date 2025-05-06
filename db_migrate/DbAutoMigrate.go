package db_migrate

import (
	"Berry_IM/models"
	"gorm.io/gorm"
)

type DbAutoMigrate struct {
	db *gorm.DB
}

func NewDbAutoMigrate(gormDb *gorm.DB) *DbAutoMigrate {
	return &DbAutoMigrate{
		db: gormDb,
	}
}

// AutoMigrate 自动迁移数据库
func (dam *DbAutoMigrate) AutoMigrate() {
	// 迁移 UserBasic
	err := dam.db.AutoMigrate(&models.UserBasic{})
	if err != nil {
		panic(err)
	}
}
