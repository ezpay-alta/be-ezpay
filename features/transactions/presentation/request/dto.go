package request

import (
	"ezpay/features/transactions"
	"strconv"
)

type TransactionRequest struct {
	Product    string `json:"product"`
	Type       string `json:"type"`
	BillNumber int    `json:"billNumber"`
	Region     string `json:"region"`
	Nominal    int    `json:"nominal"`
	Total      int    `json:"total"`
	Quantity   int    `json:"quantity"`
}

type XenditRequest struct {
	ID     string `json:"external_id"`
	Status string `json:"status"`
}

func (requestData *TransactionRequest) ToTransactionCore(userId int, productId int) transactions.Core {
	return transactions.Core{
		ProductID:  productId,
		UserID:     userId,
		Total:      requestData.Total,
		BillNumber: requestData.BillNumber,
		Region:     requestData.Region,
		Nominal:    requestData.Nominal,
		Quantity:   requestData.Quantity,
		Type:       requestData.Type,
	}
}

func (requestData *XenditRequest) ToTransactionCore() transactions.Core {
	transactionId, _ := strconv.Atoi(requestData.ID)
	return transactions.Core{
		ID:     transactionId,
		Status: requestData.Status,
	}
}
