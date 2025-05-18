package ucenter

type WebUserCreatedRequest struct {
	Name     string `json:"name" binding:"required" valid:"minstringlength(5)"`
	Password string `json:"password" binding:"required" valid:"stringlength(6|10)"`
	Email    string `json:"email" binding:"required" valid:"email"`
	Phone    string `json:"phone" binding:"required" valid:"matches(^1[3-9]\\d{9}$)"`
}
