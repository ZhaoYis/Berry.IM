package ucenter

import (
	domainRequest "Berry_IM/domain/models/request/ucenter"
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
	userCreatedModel := &domainRequest.UserCreatedModel{
		Name:     request.Name,
		Password: request.Password,
		Email:    request.Email,
		Phone:    request.Phone,
	}
	userBasicModel, err := us.userDomainService.Create(userCreatedModel)
	if err != nil {
		return nil, err
	}
	return &bizResponse.UserCreatedResultBO{
		Uid: userBasicModel.Uid,
	}, nil
}
