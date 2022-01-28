package response

import (
	productResponse "ezpay/features/products/presentation/response"
	"ezpay/features/transactions"
)

type TransactionResponse struct {
	ID         int                             `json:"id"`
	Product    productResponse.ProductResponse `json:"product"`
	Type       string                          `json:"type"`
	BillNumber int                             `json:"billNumber"`
	Nominal    int                             `json:"nominal"`
	Quantity   int                             `json:"quantity"`
	Total      int                             `json:"total"`
	Status     string                          `json:"status"`
}

func ToTransactionResponse(transaction transactions.Core) TransactionResponse {
	return TransactionResponse{
		ID:         transaction.ID,
		Product:    productResponse.ToProductResponse(transaction.Product),
		Type:       transaction.Type,
		BillNumber: transaction.BillNumber,
		Nominal:    transaction.Nominal,
		Quantity:   transaction.Quantity,
		Total:      transaction.Total,
		Status:     transaction.Status,
	}
}

func ToTransactionResponseList(transactionList []transactions.Core) []TransactionResponse {
	convertedTransaction := []TransactionResponse{}
	for _, transaction := range transactionList {
		convertedTransaction = append(convertedTransaction, ToTransactionResponse(transaction))
	}

	return convertedTransaction
}
