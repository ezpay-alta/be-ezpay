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

	return n

}
