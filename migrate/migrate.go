package migrate

import (
	"ezpay/db"
	"ezpay/features/users/data"
)

func AutoMigrate() {
	db.DB.AutoMigrate(&data.User{})
}
