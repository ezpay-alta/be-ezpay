package request

import (
	"ezpay/features/transactions"
	"strconv"
)

type TransactionRequest struct {
	UserID          int    `json:"user_id"`
	ProductID       int    `json:"product_id"`
	PromoID         int    `json:"promo_id"`
	Total           int    `json:"total"`
	Status          string `json:"status"`
	PaymentMethodID int    `json:"paymentMethod_id"`
}

type XenditRequest struct {
	ID     string `json:"external_id"`
	Status string `json:"status"`
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

func (requestData *XenditRequest) ToTransactionCore() transactions.Core {
	transactionId, _ := strconv.Atoi(requestData.ID)
	return transactions.Core{
		ID:     transactionId,
		Status: requestData.Status,
	}
}
