package routers

import (
	"Berry_IM/routers/api_routers"
	"Berry_IM/routers/swagger"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()
	// 初始化swagger
	swagger.InitSwaggerInfo()
	// 注册路由
	api_routers.InitApiRouters(r)
	// 注册swagger路由
	swagger.RegisterSwaggerRouter(r)
	return r
}
