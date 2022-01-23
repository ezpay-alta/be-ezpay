package transactions

import (
	"ezpay/features/products"
	"ezpay/features/promos"
	"ezpay/features/users"
)

type Core struct {
	ID              int
	UserID          int
	User            users.Core
	ProductID       int
	Product         products.Core
	PromoID         int
	Promo           promos.Core
	Total           int
	Status          string
	PaymentMethodID int
	// PaymentMethod   PaymentMethodCore
}

// type PaymentMethodCore struct {
// 	ID   int
// 	Code string
// 	Type string
// 	Name string
// 	// StepPayment StepPaymentCore
// }

// type StepPaymentCore struct {
// 	ID     int
// 	Label  string
// 	Detail string
// }

type Business interface {
	CreateTransaction(data Core) (int, error)
	GetAllTransactions() ([]Core, error)
	GetTransactionById(transactionId int) (Core, error)
	UpdateTransactionById(transactionId int, data Core) error
	DeleteTransactionById(transactionId int) error
}

type Data interface {
	CreateTransaction(data Core) (int, error)
	GetAllTransactions() ([]Core, error)
	GetTransactionById(transactionId int) (Core, error)
	UpdateTransactionById(transactionId int, data Core) error
	DeleteTransactionById(transactionId int) error
}
