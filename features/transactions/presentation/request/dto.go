package request

import (
	"ezpay/features/transactions"
)

type TransactionRequest struct {
	ID              int    `json:"id"`
	UserID          int    `json:"user_id"`
	ProductID       int    `json:"product_id"`
	PromoID         int    `json:"promo_id"`
	Total           int    `json:"total"`
	Status          string `json:"status"`
	PaymentMethodID int    `json:"paymentMethod_id"`
}

func (requestData *TransactionRequest) ToTransactionCore() transactions.Core {
	return transactions.Core{
		ID:              requestData.ID,
		UserID:          requestData.UserID,
		ProductID:       requestData.ProductID,
		PromoID:         requestData.PromoID,
		Total:           requestData.Total,
		Status:          requestData.Status,
		PaymentMethodID: requestData.PaymentMethodID,
	}
}
