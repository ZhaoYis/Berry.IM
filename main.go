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

//	@title			Berry IM API
//	@version		1.0
//	@description	This is berry im server.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host		localhost:9191
//	@BasePath	/api/v1

//	@securityDefinitions.basic	BasicAuth

// @externalDocs.description	OpenAPI
// @externalDocs.url			https://swagger.io/resources/open-api/
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
