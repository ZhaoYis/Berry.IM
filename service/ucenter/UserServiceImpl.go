package ucenter

import (
	userDomainService "Berry_IM/domain/ucenter"
	bizRequest "Berry_IM/service/models/request/ucenter"
	bizResponse "Berry_IM/service/models/response/ucenter"
)

type UserServiceImpl struct {
	userDomainService userDomainService.UserDomainService
}

func NewUserService() *UserServiceImpl {
	return &UserServiceImpl{
		userDomainService: userDomainService.NewUserDomainService(),
	}
}

func (us *UserServiceImpl) CreateUser(request *bizRequest.BizUserCreatedRequest) (*bizResponse.UserCreatedResultBO, error) {
	return &bizResponse.UserCreatedResultBO{
		Uid: "1",
	}, nil
}
