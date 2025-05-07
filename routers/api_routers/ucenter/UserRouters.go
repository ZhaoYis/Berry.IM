package ucenter

import (
	"Berry_IM/controller/v1/ucenter"
	"fmt"
	"github.com/gin-gonic/gin"
)

// InitUserControllerRouters 初始化用户模块路由
func InitUserControllerRouters(r *gin.Engine) {
	// 初始化控制器
	ucenterController := ucenter.NewUserController()

	v1 := r.Group("/api/v1")
	{
		userGroup := v1.Group("/ucenter")
		{
			//api/v1/ucenter路由全局中间件
			userGroup.Use(func(c *gin.Context) {
				fmt.Printf("[Middleware]Befor-/api/v1/ucenter路由全局中间件...\n")
				c.Next()
				fmt.Printf("[Middleware]After-/api/v1/ucenter路由全局中间件...\n")
			})

			userGroup.POST("/create", func(c *gin.Context) {
				// 路由中间件
				fmt.Printf("[Middleware]Befor-create user...\n")

				// 获取手动设置的变量
				env := c.GetString("env")
				fmt.Printf("[Middleware]env: %v\n", env)

				c.Next()
				fmt.Printf("[Middleware]After-create user...\n")
			}, ucenterController.CreateUser)

			userGroup.GET("/getUserById/:uid", ucenterController.GetUserById)
			//userGroup.POST("/uploadAvatar", ucenterController.UploadAvatar)
		}
	}
}
