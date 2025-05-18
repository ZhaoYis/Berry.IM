package ucenter

import (
	"Berry_IM/dao"
	domainRequest "Berry_IM/domain/models/request/ucenter"
	domainResponse "Berry_IM/domain/models/response/ucenter"
	"Berry_IM/models"
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserDomainServiceImpl struct {
	db *gorm.DB
}

func NewUserDomainService() *UserDomainServiceImpl {
	return &UserDomainServiceImpl{
		db: dao.GetMySqlDB(),
	}
}

/**
 * Create 创建用户
 * @param request
 * @return
 */
func (us *UserDomainServiceImpl) Create(request *domainRequest.UserCreatedModel) (*domainResponse.UserBasicModel, error) {
	userBasic := &models.UserBasic{
		Uid:      uuid.NewString(),
		Name:     request.Name,
		Password: request.Password,
		Phone:    request.Phone,
		Email:    request.Email,
		Identity: uuid.NewString(),
	}
	us.db.Create(userBasic)
	return &domainResponse.UserBasicModel{
		Uid:      userBasic.Uid,
		Name:     userBasic.Name,
		Password: userBasic.Password,
		Phone:    userBasic.Phone,
		Email:    userBasic.Email,
	}, nil
}

/**
 * Update 更新用户信息
 * @param request
 * @return
 */
func (us *UserDomainServiceImpl) Update(id string, request *domainRequest.UserUpdatedModel) (bool, error) {
	us.db.Model(&models.UserBasic{}).Where("uid = ?", id).Updates(request)
	return true, nil
}

/**
 * GetUserById 根据用户id获取用户信息
 * @param uid
 * @return
 */
func (us *UserDomainServiceImpl) GetUserById(uid string) (*domainResponse.UserBasicModel, error) {
	var userBasic models.UserBasic
	us.db.Where("uid = ?", uid).First(&userBasic)
	if (userBasic == models.UserBasic{} || userBasic.Uid == "") {
		return nil, errors.New("用户不存在")
	}
	return &domainResponse.UserBasicModel{
		Uid:        userBasic.Uid,
		Name:       userBasic.Name,
		Password:   userBasic.Password,
		Phone:      userBasic.Phone,
		Email:      userBasic.Email,
		Identity:   userBasic.Identity,
		ClientIp:   userBasic.ClientIp,
		LoginTime:  userBasic.LoginTime,
		LogoutTime: userBasic.LogoutTime,
		IsLogout:   userBasic.IsLogout,
		DeviceInfo: userBasic.DeviceInfo,
	}, nil
}
