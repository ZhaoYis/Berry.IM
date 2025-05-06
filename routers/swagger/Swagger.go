package swagger

import (
	"Berry_IM/docs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitSwaggerInfo() {
	// set swagger info
	docs.SwaggerInfo.Title = "Berry IM API"
	docs.SwaggerInfo.Description = "This is berry im server."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:9191"
	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
}

func RegisterSwaggerRouter(r *gin.Engine) {
	// 注册swagger路由
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
