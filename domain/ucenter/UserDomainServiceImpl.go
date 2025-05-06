package ucenter

import (
	domainRequest "Berry_IM/domain/models/request/ucenter"
	domainResponse "Berry_IM/domain/models/response/ucenter"
	"Berry_IM/models"
	"Berry_IM/utils"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserDomainServiceImpl struct {
	db *gorm.DB
}

func NewUserDomainService() *UserDomainServiceImpl {
	return &UserDomainServiceImpl{
		db: utils.GetMySqlDB(),
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
		PassWord: request.Password,
		Phone:    request.Phone,
		Email:    request.Email,
	}
	us.db.Create(userBasic)
	return &domainResponse.UserBasicModel{
		Uid:      userBasic.Uid,
		Name:     userBasic.Name,
		PassWord: userBasic.PassWord,
		Phone:    userBasic.Phone,
		Email:    userBasic.Email,
	}, nil
}
