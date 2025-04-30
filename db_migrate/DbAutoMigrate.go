package db_migrate

import (
	"Berry_IM/models"
	"Berry_IM/utils"
)

type DbAutoMigrate struct {
}

func NewDbAutoMigrate() *DbAutoMigrate {
	return &DbAutoMigrate{}
}

// AutoMigrate 自动迁移数据库
func (dam *DbAutoMigrate) AutoMigrate() {
	db := utils.GetMySqlDB()
	// 迁移 UserBasic
	err := db.AutoMigrate(&models.UserBasic{})
	if err != nil {
		panic(err)
	}
}
