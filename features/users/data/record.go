package data

import (
	"ezpay/features/users"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       int
	Fullname string
	Phone    string `gorm:"unique"`
	Email    string `gorm:"unique"`
	Password string
}

func ToUserRecord(user users.Core) User {
	return User{
		ID:       user.ID,
		Email:    user.Email,
		Password: user.Password,
		Fullname: user.Fullname,
		Phone:    user.Phone,
	}
}

func ToUserCore(user User) users.Core {
	return users.Core{
		ID:       user.ID,
		Email:    user.Email,
		Fullname: user.Fullname,
	}
}

func ToUserCoreList(uList []User) []users.Core {
	convertedUser := []users.Core{}

	for _, user := range uList {
		convertedUser = append(convertedUser, ToUserCore(user))
	}

	return convertedUser
}

func ToUserRecordList(uList []users.Core) []User {
	convertedUser := []User{}

	for _, user := range uList {
		convertedUser = append(convertedUser, ToUserRecord(users.Core{
			ID:       user.ID,
			Email:    user.Email,
			Fullname: user.Fullname,
		}))
	}

	return convertedUser
}
