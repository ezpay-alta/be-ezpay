package transactions

import (
	"ezpay/features/products"
	"ezpay/features/users"
)

type Core struct {
	ID         int
	UserID     int
	User       users.Core
	ProductID  int
	Product    products.Core
	Total      int
	Status     string
	BillNumber int
	Nominal    int
	Quantity   int
	Type       string
	Region     string
}

type Business interface {
	CreateTransaction(data Core) (int, error)
	GetAllTransactions() ([]Core, error)
	GetTransactionById(transactionId int) (Core, error)
	GetTransactionByUserId(userId int) ([]Core, error)
	UpdateTransactionById(transactionId int, data Core) error
	DeleteTransactionById(transactionId int) error
}

type Data interface {
	CreateTransaction(data Core) (int, error)
	GetAllTransactions() ([]Core, error)
	GetTransactionById(transactionId int) (Core, error)
	GetTransactionByUserId(userId int) ([]Core, error)
	UpdateTransactionById(transactionId int, data Core) error
	DeleteTransactionById(transactionId int) error
}
