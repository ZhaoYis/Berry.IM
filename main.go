package main

import (
	"Berry_IM/routers"
	"Berry_IM/utils"
	"fmt"
)

func init() {
	fmt.Println("[main]Let's go！")
	//初始化应用
	utils.InitApp()
}

func main() {
	fmt.Println("Hello Berry IM...")
	r := routers.Router()
	// 启动服务器
	err := r.Run(":" + utils.GetAppConfig().Server.Port)
	if err != nil {
		panic(err)
		return
	}
}
