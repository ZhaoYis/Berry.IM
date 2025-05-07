package ucenter

type BizUserUpdatedRequest struct {
	Uid        string `json:"uid"`
	Name       string `json:"name"`
	Password   string `json:"password"`
	Email      string `json:"email"`
	Phone      string `json:"phone"`
	ClientIp   string `json:"clientIp"`
	ClientPort string `json:"clientPort"`
	DeviceInfo string `json:"deviceInfo"`
}
