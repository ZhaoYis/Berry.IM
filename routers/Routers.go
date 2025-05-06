package routers

import "github.com/gin-gonic/gin"

func Router() *gin.Engine {
	r := gin.Default()
	// 注册路由
	initApiRouters(r)
	return r
}
