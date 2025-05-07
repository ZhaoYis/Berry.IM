package models

import "gorm.io/gorm"

type UserBasic struct {
	gorm.Model
	Uid           string `gorm:"column:uid;type:varchar(36);not null;unique"`
	Name          string `gorm:"column:name;type:varchar(128);not null"`
	Password      string `gorm:"column:password;type:varchar(128);not null"`
	Phone         string `gorm:"column:phone;type:varchar(32)"`
	Email         string `gorm:"column:email;type:varchar(64)"`
	Identity      string `gorm:"column:identity;type:varchar(128)"`
	ClientIp      string `gorm:"column:client_ip;type:varchar(36)"`
	ClientPort    string `gorm:"column:client_port;type:varchar(8)"`
	Salt          string `gorm:"column:salt;type:varchar(128)"`
	LoginTime     uint64 `gorm:"column:login_time;type:int"`
	HeartbeatTime uint64 `gorm:"column:heartbeat_time;type:int"`
	LogoutTime    uint64 `gorm:"column:logout_time;type:int"`
	IsLogout      bool   `gorm:"column:is_logout;type:bit"`
	DeviceInfo    string `gorm:"column:device_info;type:varchar(128)"`
}

func (UserBasic) TableName() string {
	return "user_basic"
}
