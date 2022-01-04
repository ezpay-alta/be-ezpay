package main

import (
	"ezpay/db"
	"ezpay/migrate"
	"ezpay/routes"
)

func main() {
	db.InitDB()
	migrate.AutoMigrate()
	e := routes.New()
	e.Start(":8000")
}
