package ucenter

type UserBasicInfoResultBO struct {
	Uid   string `json:"uid"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}
