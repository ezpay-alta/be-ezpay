package response

import (
	"ezpay/features/users"
)

type UserResponse struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Fullname string `json:"fullname"`
	Phone    string `json:"phone"`
}

func ToUserResponse(user users.Core) UserResponse {
	return UserResponse{
		ID:       user.ID,
		Email:    user.Email,
		Fullname: user.Fullname,
		Phone:    user.Phone,
	}
}

func ToUserResponseList(userList []users.Core) []UserResponse {
	convertedUser := []UserResponse{}
	for _, user := range userList {
		convertedUser = append(convertedUser, ToUserResponse(user))
	}

	return convertedUser
}
