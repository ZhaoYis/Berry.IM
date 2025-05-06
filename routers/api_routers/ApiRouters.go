package api_routers

import (
	ucenterRouter "Berry_IM/routers/api_routers/ucenter"
	"Berry_IM/utils"
	"fmt"
	"github.com/gin-gonic/gin"
)

func InitApiRouters(r *gin.Engine) {
	// 设置文件大小限制（默认是32MiB）
	r.MaxMultipartMemory = 8 << 20 // 8 MiB

	//全局中间件
	r.Use(func(c *gin.Context) {
		// 设置变量，整个请求都可以获取到
		// 注意：这里设置的变量只在当前请求的上下文中有效，不会影响其他请求的变量
		c.Set("env", utils.GetEnv())

		fmt.Printf("[Middleware]Befor-全局中间件...\n")
		c.Next()
		fmt.Printf("[Middleware]After-全局中间件...\n")
	})

	// 初始化用户模块路由
	ucenterRouter.InitUserControllerRouters(r)
}
