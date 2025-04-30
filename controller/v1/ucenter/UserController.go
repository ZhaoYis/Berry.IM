package ucenter

import (
	"Berry_IM/controller"
	webRequest "Berry_IM/controller/models/request/ucenter"
	webResponse "Berry_IM/controller/models/response/ucenter"
	bizRequest "Berry_IM/service/models/request/ucenter"
	userService "Berry_IM/service/ucenter"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	// 继承 BaseController
	controller.BaseController
	userService userService.UserService
}

func NewUserController() *UserController {
	return &UserController{
		userService: userService.NewUserService(),
	}
}

func (uc *UserController) CreateUser(c *gin.Context) {
	var webReq webRequest.WebUserCreatedRequest
	if err := c.ShouldBindJSON(&webReq); err != nil {
		uc.Error(c, err.Error())
		return
	}
	// 调用服务层
	bizReq := &bizRequest.BizUserCreatedRequest{
		Name:     webReq.Name,
		Password: webReq.Password,
		Email:    webReq.Email,
		Phone:    webReq.Phone,
	}
	bo, err := uc.userService.CreateUser(bizReq)
	if err != nil {
		uc.Error(c, err.Error())
	}
	uc.SuccessWithData(c, &webResponse.UserCreatedResultVO{
		Uid: bo.Uid,
	})
}
