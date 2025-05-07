package ucenter

type WebUserUpdatedRequest struct {
	Name       string `json:"name" binding:"required"`
	Password   string `json:"password" binding:"required"`
	Email      string `json:"email" binding:"required"`
	Phone      string `json:"phone" binding:"required"`
	ClientIp   string `json:"clientIp" binding:"required"`
	ClientPort string `json:"clientPort" binding:"required"`
	DeviceInfo string `json:"deviceInfo" binding:"required"`
}
