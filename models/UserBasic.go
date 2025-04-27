package models

import "gorm.io/gorm"

type UserBasic struct {
	gorm.Model
	Name          string
	PassWord      string
	Phone         string
	Email         string
	Identity      string
	ClientIp      string
	ClientPort    string
	Salt          string
	LoginTime     uint64
	HeartbeatTime uint64
	LogoutTime    uint64
	IsLogout      bool
	DeviceInfo    string
}

func (UserBasic) TableName() string {
	return "user_basic"
}
