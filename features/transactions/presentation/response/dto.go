package response

import (
	productResponse "ezpay/features/products/presentation/response"
	promoResponse "ezpay/features/promos/presentation/response"
	"ezpay/features/transactions"
	userResponse "ezpay/features/users/presentation/response"
)

type TransactionResponse struct {
	ID            int                             `json:"id"`
	User          userResponse.UserResponse       `json:"user"`
	Product       productResponse.ProductResponse `json:"product"`
	Promo         promoResponse.PromoResponse     `json:"promo"`
	Total         int                             `json:"total"`
	Status        string                          `json:"status"`
	PaymentMethod PaymentMethodResponse           `json:"paymentMethod"`
}

type PaymentMethodResponse struct {
	ID          int                 `json:"id"`
	Code        string              `json:"code"`
	Type        string              `json:"type"`
	Name        string              `json:"name"`
	StepPayment StepPaymentResponse `json:"stepPayment"`
}

type StepPaymentResponse struct {
	ID     int    `json:"id"`
	Label  string `json:"label"`
	Detail string `json:"detail"`
}

func ToTransactionResponse(transaction transactions.Core) TransactionResponse {
	return TransactionResponse{
		ID:            transaction.ID,
		User:          userResponse.ToUserResponse(transaction.User),
		Product:       productResponse.ToProductResponse(transaction.Product),
		Promo:         promoResponse.ToPromoResponse(transaction.Promo),
		Total:         transaction.Total,
		Status:        transaction.Status,
		PaymentMethod: ToPaymentMethodResponse(transaction.PaymentMethod),
	}
}

func ToPaymentMethodResponse(paymentMethod transactions.PaymentMethodCore) PaymentMethodResponse {
	return PaymentMethodResponse{
		ID:          paymentMethod.ID,
		Code:        paymentMethod.Code,
		Type:        paymentMethod.Type,
		Name:        paymentMethod.Name,
		StepPayment: ToStepPaymentResponse(paymentMethod.StepPayment),
	}
}

func ToStepPaymentResponse(stepPayment transactions.StepPaymentCore) StepPaymentResponse {
	return StepPaymentResponse{
		ID:     stepPayment.ID,
		Label:  stepPayment.Label,
		Detail: stepPayment.Detail,
	}
}

func ToTransactionResponseList(transactionList []transactions.Core) []TransactionResponse {
	convertedTransaction := []TransactionResponse{}
	for _, transaction := range transactionList {
		convertedTransaction = append(convertedTransaction, ToTransactionResponse(transaction))
	}

	return convertedTransaction
}
