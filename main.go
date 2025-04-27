package main

import (
	"Berry_IM/db_migrate"
	"fmt"
)

func init() {
	fmt.Println("[main]Let's go！")
}

func main() {
	fmt.Println("Hello Berry IM...")

	//自动迁移数据库
	dam := db_migrate.NewDbAutoMigrate()
	dam.AutoMigrate()
}
