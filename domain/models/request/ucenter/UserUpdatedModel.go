package ucenter

type UserUpdatedModel struct {
	Name          string `json:"name"`
	Password      string `json:"password"`
	Email         string `json:"email"`
	Phone         string `json:"phone"`
	ClientIp      string `json:"clientIp"`
	ClientPort    string `json:"clientPort"`
	HeartbeatTime uint64 `json:"heartbeatTime"`
	DeviceInfo    string `json:"deviceInfo"`
}
