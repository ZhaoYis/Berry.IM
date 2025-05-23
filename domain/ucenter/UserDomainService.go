package ucenter

import (
	domainRequest "Berry_IM/domain/models/request/ucenter"
	domainResponse "Berry_IM/domain/models/response/ucenter"
)

type UserDomainService interface {
	/**
	 * Create 创建用户
	 * @param request
	 * @return
	 */
	Create(request *domainRequest.UserCreatedModel) (*domainResponse.UserBasicModel, error)
	/**
	 * Update 更新用户信息
	 * @param request
	 * @return
	 */
	Update(id string, request *domainRequest.UserUpdatedModel) (bool, error)
	/**
	 * GetUserById 根据用户ID获取用户信息
	 * @param uid
	 * @return
	 */
	GetUserById(uid string) (*domainResponse.UserBasicModel, error)
}
