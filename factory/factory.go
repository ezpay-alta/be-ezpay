package factory

import (

	// user domain
	"ezpay/db"
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
)

type Presenter struct {
	UserHandler    userPresentation.UserHandler
	ProductHandler productPresentation.ProductHandler
	PromoHandler   promoPresentation.PromoHandler
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

	// products layer
	promoData := promoData.NewMysqlPromoRepository(db.DB)
	promoBusiness := promoBusiness.NewPromoBusiness(promoData)
	promoPresentation := promoPresentation.NewPromoHandler(promoBusiness)

	return Presenter{
		UserHandler:    *userPresentation,
		ProductHandler: *productPresentation,
		PromoHandler:   *promoPresentation,
	}
}
