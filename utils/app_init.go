package utils

import (
	"Berry_IM/db_migrate"
	"fmt"
)

func InitApp() {
	// 1、初始化配置文件
	InitAppConfig()
	// 2、初始化数据库
	InitMySQL()
	// 3、自动迁移数据库
	dam := db_migrate.NewDbAutoMigrate()
	dam.AutoMigrate()
	fmt.Println("[main]迁移数据库完成！")
}
