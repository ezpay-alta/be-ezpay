package auth

import "ezpay/features/users"

type Core struct {
	Token string
}

type Business interface {
	AddRefreshToken(data Core) error
	VerifyRefreshToken(data Core) error
	DeleteRefreshToken(data Core) error
	VerifyUserCredential(data users.Core) (userId int, role string, err error)
}

type Data interface {
	AddRefreshToken(data Core) error
	VerifyRefreshToken(data Core) error
	DeleteRefreshToken(data Core) error
	VerifyUserCredential(data users.Core) (userId int, role string, err error)
}
