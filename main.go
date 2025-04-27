package main

import (
	"Berry_IM/db_migrate"
	"fmt"
)

func init() {
	fmt.Println("[main]Let's go！")
	//自动迁移数据库
	dam := db_migrate.NewDbAutoMigrate()
	dam.AutoMigrate()
	fmt.Println("[main]迁移数据库完成！")
}

func main() {
	fmt.Println("Hello Berry IM...")
}
