package ucenter

import (
	bizRequest "Berry_IM/service/models/request/ucenter"
	bizResponse "Berry_IM/service/models/response/ucenter"
)

type UserService interface {
	/**
	 * 创建用户, 返回用户ID
	 * @param req *ucenter.BizUserCreatedRequest
	 * @return string 用户ID
	 * @return error 错误信息
	 */
	CreateUser(*bizRequest.BizUserCreatedRequest) (*bizResponse.UserCreatedResultBO, error)
	/**
	 * 根据用户ID获取用户信息
	 * @param uid string 用户ID
	 * @return *ucenter.BizUserBasicInfoResultBO 用户信息
	 * @return error 错误信息
	 */
	GetUserById(uid string) (*bizResponse.UserBasicInfoResultBO, error)
	/**
	 * 更新用户信息
	 * @param b *ucenter.BizUserUpdatedRequest
	 * @return *ucenter.BizUserUpdatedResultBO
	 * @return error
	 */
	UpdateUser(b *bizRequest.BizUserUpdatedRequest) (*bizResponse.UserUpdatedResultBO, error)
}
