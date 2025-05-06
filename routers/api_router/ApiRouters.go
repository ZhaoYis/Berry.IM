package api_router

import (
	"Berry_IM/controller/v1/ucenter"
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
	initUserControllerRouters(r)
}

/**
 * @Description: 初始化用户模块路由
 * @param r
 */
func initUserControllerRouters(r *gin.Engine) {
	// 初始化控制器
	ucenterController := ucenter.NewUserController()

	userGroup := r.Group("/api/v1")
	{
		//api/v2/users路由全局中间件
		userGroup.Use(func(c *gin.Context) {
			fmt.Printf("[Middleware]Befor-/api/v1/ucenter路由全局中间件...\n")
			c.Next()
			fmt.Printf("[Middleware]After-/api/v1/ucenter路由全局中间件...\n")
		})

		userGroup.POST("/ucenter/create", func(c *gin.Context) {
			// 路由中间件
			fmt.Printf("[Middleware]Befor-create user...\n")

			// 获取手动设置的变量
			env := c.GetString("env")
			fmt.Printf("[Middleware]env: %v\n", env)

			c.Next()
			fmt.Printf("[Middleware]After-create user...\n")
		}, ucenterController.CreateUser)

		//userGroup.GET("/getUserList", ucenterController.GetUserList)
		//userGroup.GET("/getUserById", ucenterController.GetUserById)
		//userGroup.POST("/uploadAvatar", ucenterController.UploadAvatar)
	}
}
