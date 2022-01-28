package data

import (
	productsData "ezpay/features/products/data"
	"ezpay/features/transactions"
	usersData "ezpay/features/users/data"

	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	ID         int
	UserID     int
	User       usersData.User
	ProductID  int
	Product    productsData.Product
	BillNumber int
	Region     string
	Quantity   int
	Nominal    int
	Total      int
	Status     string
	Type       string
}

func ToTransactionRecord(transaction transactions.Core) Transaction {
	return Transaction{
		ID:         transaction.ID,
		UserID:     transaction.UserID,
		User:       usersData.ToUserRecord(transaction.User),
		ProductID:  transaction.ProductID,
		Product:    productsData.ToProductRecord(transaction.Product),
		Type:       transaction.Type,
		Total:      transaction.Total,
		Status:     transaction.Status,
		BillNumber: transaction.BillNumber,
		Nominal:    transaction.Nominal,
		Quantity:   transaction.Quantity,
		Region:     transaction.Region,
	}
}

func ToTransactionCore(transaction Transaction) transactions.Core {
	return transactions.Core{
		ID:         transaction.ID,
		UserID:     transaction.UserID,
		User:       usersData.ToUserCore(transaction.User),
		ProductID:  transaction.ProductID,
		Product:    productsData.ToProductCore(transaction.Product),
		Type:       transaction.Type,
		Total:      transaction.Total,
		Status:     transaction.Status,
		BillNumber: transaction.BillNumber,
		Nominal:    transaction.Nominal,
		Quantity:   transaction.Quantity,
		Region:     transaction.Region,
	}
}

func ToTransactionCoreList(aList []Transaction) []transactions.Core {
	convertedTransaction := []transactions.Core{}

	for _, transaction := range aList {
		convertedTransaction = append(convertedTransaction, ToTransactionCore(transaction))
	}

	return convertedTransaction
}
