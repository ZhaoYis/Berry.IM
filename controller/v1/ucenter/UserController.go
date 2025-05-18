package ucenter

import (
	"Berry_IM/controller"
	webRequest "Berry_IM/controller/models/request/ucenter"
	webResponse "Berry_IM/controller/models/response/ucenter"
	bizRequest "Berry_IM/service/models/request/ucenter"
	userService "Berry_IM/service/ucenter"
	"github.com/asaskevich/govalidator"
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

// CreateUser 创建用户
//
//	@Summary		创建用户
//	@Description	创建用户
//	@Tags			ucenter
//	@Accept			json
//	@Produce		json
//	@Param			webReq	body		webRequest.WebUserCreatedRequest	true	"请求参数"
//	@Success		200		{object}	webResponse.UserCreatedResultVO
//	@Router			/ucenter/create [post]
func (uc *UserController) CreateUser(c *gin.Context) {
	var webReq webRequest.WebUserCreatedRequest
	if err := c.ShouldBindJSON(&webReq); err != nil {
		uc.Error(c, err.Error())
		return
	}
	// 校验参数
	_, err := govalidator.ValidateStruct(&webReq)
	if err != nil {
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

// UpdateUser 更新用户信息
//
//	@Summary		更新用户信息
//	@Description	更新用户信息
//	@Tags			ucenter
//	@Accept			json
//	@Produce		json
//	@Param			uid		path		string								true	"用户ID"
//	@Param			webReq	body		webRequest.WebUserUpdatedRequest	true	"请求参数"
//	@Success		200		{object}	webResponse.UserUpdatedResultVO
//	@Router			/ucenter/update/{uid} [put]
func (uc *UserController) UpdateUser(c *gin.Context) {
	var webReq webRequest.WebUserUpdatedRequest
	if err := c.ShouldBindJSON(&webReq); err != nil {
		uc.Error(c, err.Error())
	}
	// 校验参数
	_, err := govalidator.ValidateStruct(&webReq)
	if err != nil {
		uc.Error(c, err.Error())
		return
	}
	bo, err := uc.userService.UpdateUser(&bizRequest.BizUserUpdatedRequest{
		Uid:        c.Param("uid"),
		Name:       webReq.Name,
		Password:   webReq.Password,
		Email:      webReq.Email,
		Phone:      webReq.Phone,
		ClientIp:   webReq.ClientIp,
		ClientPort: webReq.ClientPort,
		DeviceInfo: webReq.DeviceInfo,
	})
	if err != nil {
		uc.Error(c, err.Error())
	}
	uc.SuccessWithData(c, &webResponse.UserUpdatedResultVO{
		Uid: bo.Uid,
	})
}

// GetUserById 根据用户ID获取用户信息
//
//	@Summary		获取用户信息
//	@Description	获取用户信息
//	@Tags			ucenter
//	@Accept			json
//	@Produce		json
//	@Param			uid	path		string	true	"用户ID"
//	@Success		200	{object}	webResponse.UserBasicInfoResultVO
//	@Router			/ucenter/getUserById/{uid} [get]
func (uc *UserController) GetUserById(c *gin.Context) {
	uid := c.Param("uid")
	bo, err := uc.userService.GetUserById(uid)
	if err != nil {
		uc.Error(c, err.Error())
	}
	uc.SuccessWithData(c, &webResponse.UserBasicInfoResultVO{
		Uid:   bo.Uid,
		Name:  bo.Name,
		Email: bo.Email,
		Phone: bo.Phone,
	})
}
