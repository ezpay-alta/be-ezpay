package business

import (
	"ezpay/features/products"
	"ezpay/features/promos"
	"ezpay/features/transactions"
)

type transactionUsecase struct {
	ProductData     products.Data
	PromoData       promos.Data
	TransactionData transactions.Data
}

func NewTransactionBusiness(transactionData transactions.Data, productData products.Data, promoData promos.Data) transactions.Business {
	return &transactionUsecase{
		TransactionData: transactionData,
		ProductData:     productData,
		PromoData:       promoData,
	}
}

func (tu *transactionUsecase) CreateTransaction(transaction transactions.Core) error {
	product, _ := tu.ProductData.GetProductById(transaction.ProductID)
	promo, _ := tu.PromoData.GetPromoById(transaction.PromoID)
	total := product.Price - promo.Amount
	transaction.Total = total
	err := tu.TransactionData.CreateTransaction(transaction)
	if err != nil {
		return err
	}

	return nil
}

func (pu *transactionUsecase) GetAllTransactions() ([]transactions.Core, error) {
	transactions, err := pu.TransactionData.GetAllTransactions()
	if err != nil {
		return nil, err
	}

	return transactions, nil
}

func (pu *transactionUsecase) GetTransactionById(transactionId int) (transactions.Core, error) {
	transaction, err := pu.TransactionData.GetTransactionById(transactionId)
	if err != nil {
		return transactions.Core{}, err
	}

	return transaction, nil
}

func (pu *transactionUsecase) UpdateTransactionById(transactionId int, data transactions.Core) error {
	err := pu.TransactionData.UpdateTransactionById(transactionId, data)
	if err != nil {
		return err
	}

	return nil
}

func (pu *transactionUsecase) DeleteTransactionById(transactionId int) error {
	err := pu.TransactionData.DeleteTransactionById(transactionId)
	if err != nil {
		return err
	}

	return nil
}
