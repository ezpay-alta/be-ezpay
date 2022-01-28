package routes

import (
	"ezpay/factory"
	"ezpay/features/middlewares"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var ACCESS_TOKEN_KEY string = os.Getenv("ACCESS_TOKEN_KEY")

func New() *echo.Echo {
	n := echo.New()

	e := n.Group("/v1")
	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	configJWT := middleware.JWTConfig{
		SigningKey: []byte(ACCESS_TOKEN_KEY),
		Claims:     &middlewares.JwtCustomClaims{},
	}

	presenter := factory.Init()

	eUsers := e.Group("/users")
	eUsers.POST("", presenter.UserHandler.RegisterUserHandler)
	eUsers.GET("", presenter.UserHandler.GetAllUsersHandler)
	eUsers.GET("/:userId", presenter.UserHandler.GetUserByIdHandler)
	eUsers.PATCH("/:userId", presenter.UserHandler.UpdateUserHandler)
	eUsers.DELETE("/:userId", presenter.UserHandler.DeleteUserHandler)

	eProduct := e.Group("/products")
	eProduct.POST("", presenter.ProductHandler.CreateProductHandler, middleware.JWTWithConfig(configJWT))
	// eProduct.GET("", presenter.ProductHandler.GetAllProductHandler)
	eProduct.GET("", presenter.ProductHandler.GetProductByTypeHandler)
	eProduct.GET("/:productId", presenter.ProductHandler.GetProductByIdHandler)
	eProduct.PATCH("/:productId", presenter.ProductHandler.UpdateProductByIdHandler, middleware.JWTWithConfig(configJWT))
	eProduct.DELETE("/:productId", presenter.ProductHandler.DeleteProductByIdHandler, middleware.JWTWithConfig(configJWT))

	eType := eProduct.Group("/type")
	eType.POST("", presenter.ProductHandler.CreateProductTypeHandler, middleware.JWTWithConfig(configJWT))
	eType.GET("", presenter.ProductHandler.GetAllTypeProductHandler)
	eType.GET("/:typeId", presenter.ProductHandler.GetProductTypeByIdHandler)
	eType.PATCH("/:typeId", presenter.ProductHandler.UpdateProductTypeByIdHandler, middleware.JWTWithConfig(configJWT))
	eType.DELETE("/:typeId", presenter.ProductHandler.DeleteProductTypeByIdHandler, middleware.JWTWithConfig(configJWT))

	ePromo := e.Group("/promos")
	ePromo.POST("", presenter.PromoHandler.CreatePromoHandler, middleware.JWTWithConfig(configJWT))
	ePromo.GET("", presenter.PromoHandler.GetAllPromoHandler)
	ePromo.GET("/:promoId", presenter.PromoHandler.GetPromoByIdHandler)
	ePromo.PATCH("/:promoId", presenter.PromoHandler.UpdatePromoByIdHandler, middleware.JWTWithConfig(configJWT))
	ePromo.DELETE("/:promoId", presenter.PromoHandler.DeletePromoByIdHandler, middleware.JWTWithConfig(configJWT))

	eTransaction := e.Group("/transactions")
	eTransaction.POST("", presenter.TransactionHandler.CreateTransactionHandler)
	eTransaction.POST("/xendit", presenter.TransactionHandler.UpdateTransactionByXenditHandler)
	eTransaction.GET("", presenter.TransactionHandler.GetAllTransactionsHandler)
	eTransaction.GET("/:transactionId", presenter.TransactionHandler.GetTransactionByIdHandler)
	eTransaction.PATCH("/:transactionId", presenter.TransactionHandler.UpdateTransactionByIdHandler)
	eTransaction.DELETE("/:transactionId", presenter.TransactionHandler.DeleteTransactionByIdHandler)

	eAuth := e.Group("/auth")
	eAuth.POST("", presenter.AuthHandler.LoginHandler)
	eAuth.PUT("", presenter.AuthHandler.ReLoginHandler)
	eAuth.DELETE("", presenter.AuthHandler.LogoutHandler)

	return n

}
