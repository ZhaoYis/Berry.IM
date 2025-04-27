package models

import "gorm.io/gorm"

type UserBasic struct {
	gorm.Model
	Uid           string `gorm:"column:uid;type:varchar(36);not null;unique"`
	Name          string `gorm:"column:name;type:varchar(128);not null"`
	PassWord      string `gorm:"column:pass_word;type:varchar(128);not null"`
	Phone         string `gorm:"column:phone;type:varchar(32)"`
	Email         string `gorm:"column:email;type:varchar(64)"`
	Identity      string `gorm:"column:identity;type:varchar(128)"`
	ClientIp      string `gorm:"column:client_ip;type:varchar(36)"`
	ClientPort    string `gorm:"column:client_port;type:varchar(8)"`
	Salt          string `json:"salt,type:varchar(128)"`
	LoginTime     uint64 `json:"login_time,type:uint64"`
	HeartbeatTime uint64 `json:"heartbeat_time,type:uint64"`
	LogoutTime    uint64 `json:"logout_time,type:uint64"`
	IsLogout      bool   `json:"is_logout,type:bool"`
	DeviceInfo    string `json:"device_info,type:varchar(128)"`
}

func (UserBasic) TableName() string {
	return "user_basic"
}
