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
}
