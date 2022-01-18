package routes

import (
	"ezpay/factory"
	"net/http"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	n := echo.New()

	e := n.Group("/v1")
	// admin := e.Group("admin")

	e.GET("", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	presenter := factory.Init()

	eUsers := e.Group("/users")
	eUsers.POST("", presenter.UserHandler.RegisterUserHandler)
	eUsers.GET("", presenter.UserHandler.GetAllUsersHandler)
	eUsers.GET("/:userId", presenter.UserHandler.GetUserByIdHandler)
	eUsers.PATCH("/:userId", presenter.UserHandler.UpdateUserHandler)
	eUsers.DELETE("/:userId", presenter.UserHandler.DeleteUserHandler)

	eProduct := e.Group("/products")
	eProduct.POST("", presenter.ProductHandler.CreateProductHandler)
	eProduct.GET("", presenter.ProductHandler.GetAllProductHandler)
	eProduct.GET("/:productId", presenter.ProductHandler.GetProductByIdHandler)
	eProduct.PATCH("/:productId", presenter.ProductHandler.UpdateProductByIdHandler)
	eProduct.DELETE("/:productId", presenter.ProductHandler.DeleteProductByIdHandler)

	eType := eProduct.Group("/type")
	eType.POST("", presenter.ProductHandler.CreateProductTypeHandler)
	eType.GET("", presenter.ProductHandler.GetAllTypeProductHandler)
	eType.GET("/:typeId", presenter.ProductHandler.GetProductTypeByIdHandler)
	eType.PATCH("/:typeId", presenter.ProductHandler.UpdateProductTypeByIdHandler)
	eType.DELETE("/:typeId", presenter.ProductHandler.DeleteProductTypeByIdHandler)

	ePromo := e.Group("/promos")
	ePromo.POST("", presenter.PromoHandler.CreatePromoHandler)
	ePromo.GET("", presenter.PromoHandler.GetAllPromoHandler)
	ePromo.GET("/:promoId", presenter.PromoHandler.GetPromoByIdHandler)
	ePromo.PATCH("/:promoId", presenter.PromoHandler.UpdatePromoByIdHandler)
	ePromo.DELETE("/:promoId", presenter.PromoHandler.DeletePromoByIdHandler)

	eTransaction := eProduct.Group("/transaction")
	eTransaction.POST("", presenter.TransactionHandler.CreateTransactionHandler)
	eTransaction.GET("", presenter.TransactionHandler.GetAllTransactionsHandler)
	eTransaction.GET("/:transactionId", presenter.TransactionHandler.GetTransactionByIdHandler)
	eTransaction.PATCH("/:transactionId", presenter.TransactionHandler.UpdateTransactionByIdHandler)
	eTransaction.DELETE("/:transactionId", presenter.TransactionHandler.DeleteTransactionByIdHandler)

	return n

}
