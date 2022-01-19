package request

import (
	"ezpay/features/auth"
	"ezpay/features/users"
)

type UserRequest struct {
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

type TokenRequest struct {
	RefreshToken string `json:"refreshToken"`
}

func (requestData *UserRequest) ToUserCore() users.Core {
	return users.Core{
		Email:    requestData.Email,
		Password: requestData.Password,
	}
}

func (requestData *TokenRequest) ToTokenCore() auth.Core {
	return auth.Core{
		Token: requestData.RefreshToken,
	}
}
