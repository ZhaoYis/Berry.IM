package ucenter

type WebUserUpdatedRequest struct {
	Name       string `json:"name" binding:"required" valid:"minstringlength(5)"`
	Password   string `json:"password" binding:"required" valid:"stringlength(6|10)"`
	Email      string `json:"email" binding:"required" valid:"email"`
	Phone      string `json:"phone" binding:"required" valid:"matches(^1[3-9]\\d{9}$)"`
	ClientIp   string `json:"clientIp" binding:"required" valid:"ipv4"`
	ClientPort string `json:"clientPort" binding:"required" valid:"port"`
	DeviceInfo string `json:"deviceInfo" binding:"required"`
}
