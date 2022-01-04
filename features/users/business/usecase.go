package business

import (
	"ezpay/features/users"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type userUsecase struct {
	UserData users.Data
}

func NewUserBusiness(userData users.Data) users.Business {
	return &userUsecase{UserData: userData}
}

func (uu *userUsecase) CreateUser(data users.Core) (userId int, err error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		return 0, err
	}

	data.Password = string(hashedPassword)
	fmt.Println(data.Phone)

	userId, err = uu.UserData.CreateUser(data)

	if err != nil {
		return 0, err
	}

	return userId, nil
}

func (uu *userUsecase) GetAllUsers() ([]users.Core, error) {
	users, err := uu.UserData.GetAllUsers()
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (us *userUsecase) GetUserById(id int) (users.Core, error) {
	userData, err := us.UserData.GetUserById(id)

	if err != nil {
		return users.Core{}, err
	}

	return userData, nil
}

func (uu *userUsecase) UpdateUserById(userId int, data users.Core) error {
	err := uu.UserData.UpdateUserById(userId, data)
	if err != nil {
		return err
	}

	return nil
}

func (uu *userUsecase) DeleteUserById(userId int) error {
	err := uu.UserData.DeleteUserById(userId)
	if err != nil {
		return err
	}

	return nil
}
