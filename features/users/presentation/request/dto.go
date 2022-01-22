package request

import "ezpay/features/users"

type UserRequest struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Fullname string `json:"fullname"`
	Phone    string `json:"phone"`
	Role     string `json:"role"`
}

func (requestData *UserRequest) ToUserCore() users.Core {
	return users.Core{
		ID:       requestData.ID,
		Email:    requestData.Email,
		Password: requestData.Password,
		Fullname: requestData.Fullname,
		Phone:    requestData.Phone,
		Role:     requestData.Role,
	}
}
