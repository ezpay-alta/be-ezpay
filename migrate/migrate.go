package migrate

import (
	"ezpay/db"
	productData "ezpay/features/products/data"
	promoData "ezpay/features/promos/data"
	userData "ezpay/features/users/data"
)

func AutoMigrate() {
	db.DB.AutoMigrate(
		&userData.User{},
		&productData.Product{},
		&productData.Type{},
		&promoData.Promo{},
	)
}
