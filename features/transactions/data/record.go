package data

import (
	productsData "ezpay/features/products/data"
	promosData "ezpay/features/promos/data"
	"ezpay/features/transactions"
	usersData "ezpay/features/users/data"

	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	ID            int
	UserID        int
	User          usersData.User
	ProductID     int
	Product       productsData.Product
	PromoID       int
	Promo         promosData.Promo
	Total         int
	Status        string
	PaymentMethod PaymentMethod
}

type PaymentMethod struct {
	gorm.Model
	ID          int
	Code        string
	Type        string
	Name        string
	StepPayment StepPayment
}

type StepPayment struct {
	gorm.Model
	ID     int
	Label  string
	Detail string
}

func ToTransactionRecord(transaction transactions.Core) Transaction {
	return Transaction{
		ID:            transaction.ID,
		UserID:        transaction.UserID,
		User:          usersData.ToUserRecord(transaction.User),
		ProductID:     transaction.ProductID,
		Product:       productsData.ToProductRecord(transaction.Product),
		PromoID:       transaction.PromoID,
		Promo:         promosData.ToPromoRecord(transaction.Promo),
		Total:         transaction.Total,
		Status:        transaction.Status,
		PaymentMethod: ToPaymentMethodRecord(transaction.PaymentMethod),
	}
}

func ToPaymentMethodRecord(paymentMethod transactions.PaymentMethodCore) PaymentMethod {
	return PaymentMethod{
		ID:          paymentMethod.ID,
		Code:        paymentMethod.Code,
		Type:        paymentMethod.Type,
		Name:        paymentMethod.Name,
		StepPayment: ToStepPaymentRecord(paymentMethod.StepPayment),
	}
}

func ToStepPaymentRecord(stepPayment transactions.StepPaymentCore) StepPayment {
	return StepPayment{
		ID:     stepPayment.ID,
		Label:  stepPayment.Label,
		Detail: stepPayment.Detail,
	}
}

func ToTransactionCore(transaction Transaction) transactions.Core {
	return transactions.Core{
		ID:            transaction.ID,
		UserID:        transaction.UserID,
		User:          usersData.ToUserCore(transaction.User),
		ProductID:     transaction.ProductID,
		Product:       productsData.ToProductCore(transaction.Product),
		PromoID:       transaction.PromoID,
		Promo:         promosData.ToPromoCore(transaction.Promo),
		Total:         transaction.Total,
		Status:        transaction.Status,
		PaymentMethod: ToPaymentMethodCore(transaction.PaymentMethod),
	}
}

func ToPaymentMethodCore(paymentMethod PaymentMethod) transactions.PaymentMethodCore {
	return transactions.PaymentMethodCore{
		ID:          paymentMethod.ID,
		Code:        paymentMethod.Code,
		Type:        paymentMethod.Type,
		Name:        paymentMethod.Name,
		StepPayment: ToStepPaymentCore(paymentMethod.StepPayment),
	}
}

func ToStepPaymentCore(stepPayment StepPayment) transactions.StepPaymentCore {
	return transactions.StepPaymentCore{
		ID:     stepPayment.ID,
		Label:  stepPayment.Label,
		Detail: stepPayment.Detail,
	}
}

func ToTransactionCoreList(aList []Transaction) []transactions.Core {
	convertedTransaction := []transactions.Core{}

	for _, transaction := range aList {
		convertedTransaction = append(convertedTransaction, ToTransactionCore(transaction))
	}

	return convertedTransaction
}
