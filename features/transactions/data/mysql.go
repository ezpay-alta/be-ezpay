package data

import (
	"ezpay/features/transactions"

	"gorm.io/gorm"
)

type mysqlTransactionRepository struct {
	Conn *gorm.DB
}

func NewMysqlTransactionRepository(conn *gorm.DB) transactions.Data {
	return &mysqlTransactionRepository{
		Conn: conn,
	}
}

func (ar *mysqlTransactionRepository) CreateTransaction(transaction transactions.Core) error {
	err := ar.Conn.Create(ToTransactionRecord(transaction)).Error
	if err != nil {
		return err
	}

	return nil
}

func (ar *mysqlTransactionRepository) GetAllTransaction() ([]transactions.Core, error) {

	transactions := []Transaction{}
	err := ar.Conn.Find(&transactions).Error
	if err != nil {
		return nil, err
	}
	return ToTransactionCoreList(transactions), nil
}

func (ar *mysqlTransactionRepository) GetTransactionById(transactionId int) (transactions.Core, error) {

	transaction := Transaction{}
	err := ar.Conn.First(&transaction, transactionId).Error
	if err != nil {
		return ToTransactionCore(Transaction{}), err
	}
	return ToTransactionCore(transaction), nil
}

func (ar *mysqlTransactionRepository) UpdateTransactionById(transactionId int, data transactions.Core) error {

	transaction := ToTransactionRecord(data)
	transaction.ID = transactionId

	err := ar.Conn.First(&transaction, transactionId).Error
	if err != nil {
		return err
	}

	if data.Status != "" {
		transaction.Status = data.Status
	}
	if data.UserID != 0 {
		transaction.UserID = data.UserID
	}
	if data.ProductID != 0 {
		transaction.ProductID = data.ProductID
	}
	if data.PromoID != 0 {
		transaction.PromoID = data.PromoID
	}
	if data.Total != 0 {
		transaction.Total = data.Total
	}

	err = ar.Conn.Save(&transaction).Error
	if err != nil {
		return err
	}
	return nil
}

func (ar *mysqlTransactionRepository) DeleteTransactionById(transactionId int) error {
	err := ar.Conn.Delete(&Transaction{}, transactionId).Error
	if err != nil {
		return err
	}
	return nil
}
