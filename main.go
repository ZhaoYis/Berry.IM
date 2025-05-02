package main

import (
	"Berry_IM/db_migrate"
	"Berry_IM/routers"
	"Berry_IM/utils"
	"fmt"
	"github.com/gin-gonic/gin"
)

func init() {
	fmt.Println("[main]Let's go！")
	//初始化应用
	utils.InitApp()

	//自动迁移数据库
	dam := db_migrate.NewDbAutoMigrate()
	dam.AutoMigrate()
	fmt.Println("[main]迁移数据库完成！")
}

func main() {
	fmt.Println("Hello Berry IM...")

	r := gin.Default()
	// 注册路由
	routers.InitApiRouters(r)

	// 启动服务器
	err := r.Run(":" + utils.GetAppConfig().Server.Port)
	if err != nil {
		panic(err)
		return
	}
}
