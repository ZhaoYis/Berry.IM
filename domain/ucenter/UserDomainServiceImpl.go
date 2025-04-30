package ucenter

import (
	"Berry_IM/utils"
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
