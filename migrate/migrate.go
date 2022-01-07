package migrate

import (
	"ezpay/db"
	productData "ezpay/features/products/data"
	userData "ezpay/features/users/data"
)

func AutoMigrate() {
	db.DB.AutoMigrate(
		&userData.User{},
		&productData.Product{},
		&productData.Type{},
	)
}
