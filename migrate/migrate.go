package migrate

import (
	"ezpay/db"
	authData "ezpay/features/auth/data"
	productData "ezpay/features/products/data"
	promoData "ezpay/features/promos/data"
	transactionData "ezpay/features/transactions/data"
	userData "ezpay/features/users/data"
)

func AutoMigrate() {
	db.DB.AutoMigrate(
		&userData.User{},
		&productData.Product{},
		&productData.Type{},
		&promoData.Promo{},
		&transactionData.Transaction{},
		&authData.Authentication{},
		// &transactionData.PaymentMethod{},
		// &transactionData.StepPayment{},
	)
}
