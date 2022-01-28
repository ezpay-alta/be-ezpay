package request

import (
	"ezpay/features/transactions"
	"strconv"
)

type TransactionRequest struct {
	Product string `json:"product"`
	Jenis   string `json:"jenis"`
	Nomor   int    `json:"nomor"`
	Wilayah string `json:"wilayah"`
	Nominal int    `json:"nominal"`
	Total   int    `json:"total"`
	Bulan   int    `json:"bulan"`
}

type XenditRequest struct {
	ID     string `json:"external_id"`
	Status string `json:"status"`
}

func (requestData *TransactionRequest) ToTransactionCore(userId int, productId int) transactions.Core {
	return transactions.Core{
		ProductID: productId,
		UserID:    userId,
		Total:     requestData.Total,
		Nomor:     requestData.Nomor,
		Wilayah:   requestData.Wilayah,
		Nominal:   requestData.Nominal,
		Bulan:     requestData.Bulan,
		Jenis:     requestData.Jenis,
	}
}

func (requestData *XenditRequest) ToTransactionCore() transactions.Core {
	transactionId, _ := strconv.Atoi(requestData.ID)
	return transactions.Core{
		ID:     transactionId,
		Status: requestData.Status,
	}
}
