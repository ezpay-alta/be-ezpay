package factory

import (
	"ezpay/db"

	// user domain
	userBusiness "ezpay/features/users/business"
	userData "ezpay/features/users/data"
	userPresentation "ezpay/features/users/presentation"

	// product domain
	productBusiness "ezpay/features/products/business"
	productData "ezpay/features/products/data"
	productPresentation "ezpay/features/products/presentation"

	// promo domain
	promoBusiness "ezpay/features/promos/business"
	promoData "ezpay/features/promos/data"
	promoPresentation "ezpay/features/promos/presentation"

	// transaction domain
	transactionBusiness "ezpay/features/transactions/business"
	transactionData "ezpay/features/transactions/data"
	transactionPresentation "ezpay/features/transactions/presentation"

	// auth domain
	authBusiness "ezpay/features/auth/business"
	authData "ezpay/features/auth/data"
	authPresentation "ezpay/features/auth/presentation"
)

type Presenter struct {
	UserHandler        userPresentation.UserHandler
	ProductHandler     productPresentation.ProductHandler
	PromoHandler       promoPresentation.PromoHandler
	TransactionHandler transactionPresentation.TransactionHandler
	AuthHandler        authPresentation.AuthHandler
}

func Init() Presenter {

	// users layer
	userData := userData.NewMysqlUserRepository(db.DB)
	userBusiness := userBusiness.NewUserBusiness(userData)
	userPresentation := userPresentation.NewUserHandler(userBusiness)

	// products layer
	productData := productData.NewMysqlProductRepository(db.DB)
	productBusiness := productBusiness.NewProductBusiness(productData)
	productPresentation := productPresentation.NewProductHandler(productBusiness)

	// promo layer
	promoData := promoData.NewMysqlPromoRepository(db.DB)
	promoBusiness := promoBusiness.NewPromoBusiness(promoData)
	promoPresentation := promoPresentation.NewPromoHandler(promoBusiness)

	// transaction layer
	transactionData := transactionData.NewMysqlTransactionRepository(db.DB)
	transactionBusiness := transactionBusiness.NewTransactionBusiness(transactionData, productData, promoData)
	transactionPresentation := transactionPresentation.NewTransactionHandler(transactionBusiness, productBusiness)

	// auth layer
	authData := authData.NewMysqlAuthRepository(db.DB)
	authBusiness := authBusiness.NewAuthBusiness(authData)
	authPresentation := authPresentation.NewAuthHandler(authBusiness)

	return Presenter{
		UserHandler:        *userPresentation,
		ProductHandler:     *productPresentation,
		PromoHandler:       *promoPresentation,
		TransactionHandler: *transactionPresentation,
		AuthHandler:        *authPresentation,
	}
}
