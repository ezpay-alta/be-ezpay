package factory

import (

	// user domain
	"ezpay/db"
	userBusiness "ezpay/features/users/business"
	userData "ezpay/features/users/data"
	userPresentation "ezpay/features/users/presentation"
)

type Presenter struct {
	UserHandler userPresentation.UserHandler
}

func Init() Presenter {

	// users layer
	userData := userData.NewMysqlUserRepository(db.DB)
	userBusiness := userBusiness.NewUserBusiness(userData)
	userPresentation := userPresentation.NewUserHandler(userBusiness)

	return Presenter{
		UserHandler: *userPresentation,
	}
}
