package ucenter

import (
	bizRequest "Berry_IM/service/models/request/ucenter"
	bizResponse "Berry_IM/service/models/response/ucenter"
)

type UserServiceImpl struct {
}

func NewUserService() *UserServiceImpl {
	return &UserServiceImpl{}
}

func (us *UserServiceImpl) CreateUser(request *bizRequest.BizUserCreatedRequest) (*bizResponse.UserCreatedResultBO, error) {
	return &bizResponse.UserCreatedResultBO{
		Uid: "1",
	}, nil
}
