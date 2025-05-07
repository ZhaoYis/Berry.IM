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

func (us *UserServiceImpl) UpdateUser(request *bizRequest.BizUserUpdatedRequest) (*bizResponse.UserUpdatedResultBO, error) {
	userUpdatedModel := &domainRequest.UserUpdatedModel{
		Name:       request.Name,
		Password:   request.Password,
		Email:      request.Email,
		Phone:      request.Phone,
		ClientIp:   request.ClientIp,
		ClientPort: request.ClientPort,
		DeviceInfo: request.DeviceInfo,
	}
	_, err := us.userDomainService.Update(request.Uid, userUpdatedModel)
	if err != nil {
		return nil, err
	}
	return &bizResponse.UserUpdatedResultBO{
		Uid: request.Uid,
	}, nil
}

func (us *UserServiceImpl) GetUserById(uid string) (*bizResponse.UserBasicInfoResultBO, error) {
	userBasicModel, err := us.userDomainService.GetUserById(uid)
	if err != nil {
		return nil, err
	}
	return &bizResponse.UserBasicInfoResultBO{
		Uid:   userBasicModel.Uid,
		Name:  userBasicModel.Name,
		Email: userBasicModel.Email,
		Phone: userBasicModel.Phone,
	}, nil
}
