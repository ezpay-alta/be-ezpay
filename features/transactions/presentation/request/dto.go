package request

import (
	"ezpay/features/transactions"
)

type TransactionRequest struct {
	UserID          int    `json:"user_id"`
	ProductID       int    `json:"product_id"`
	PromoID         int    `json:"promo_id"`
	Total           int    `json:"total"`
	Status          string `json:"status"`
	PaymentMethodID int    `json:"paymentMethod_id"`
}

func (requestData *TransactionRequest) ToTransactionCore(userId int) transactions.Core {
	return transactions.Core{
		UserID:          userId,
		ProductID:       requestData.ProductID,
		PromoID:         requestData.PromoID,
		Total:           requestData.Total,
		Status:          requestData.Status,
		PaymentMethodID: requestData.PaymentMethodID,
	}
}
