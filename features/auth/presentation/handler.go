package presentation

import (
	"ezpay/features/auth"
	"ezpay/features/auth/presentation/request"
	"ezpay/features/auth/presentation/response"
	"ezpay/features/middlewares"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	AuthBusiness auth.Business
}

func NewAuthHandler(authBusiness auth.Business) *AuthHandler {
	return &AuthHandler{AuthBusiness: authBusiness}
}

func (ah *AuthHandler) LoginHandler(e echo.Context) error {

	user := request.UserRequest{}
	e.Bind(&user)

	userId, role, err := ah.AuthBusiness.VerifyUserCredential(user.ToUserCore())
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "fail",
			"message": "can not login",
			"err":     err.Error(),
		})
	}

	accessToken, err := middlewares.CreateToken(userId, role)
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "fail",
			"message": "can not login",
			"err":     err.Error(),
		})
	}

	refreshToken, err := middlewares.CreateRefreshToken(userId, role)
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "fail",
			"message": "can not login",
			"err":     err.Error(),
		})
	}

	err = ah.AuthBusiness.AddRefreshToken(auth.Core{Token: refreshToken})
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "fail",
			"message": "can not login",
			"err":     err.Error(),
		})
	}

	return e.JSON(http.StatusCreated, map[string]interface{}{
		"status":  "success",
		"message": "login",
		"data": response.AuthResponse{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		},
	})
}

func (ah *AuthHandler) ReLoginHandler(e echo.Context) error {
	auth := request.TokenRequest{}
	e.Bind(&auth)
	err := ah.AuthBusiness.VerifyRefreshToken(auth.ToTokenCore())
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "fail",
			"message": "can not login",
			"err":     err.Error(),
		})
	}

	userId, role, err := middlewares.VerifyRefreshToken(auth.RefreshToken)
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "fail",
			"message": "can not login",
			"err":     err.Error(),
		})
	}

	accessToken, err := middlewares.CreateToken(userId, role)
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "fail",
			"message": "can not login",
			"err":     err.Error(),
		})
	}
	return e.JSON(http.StatusBadRequest, map[string]interface{}{
		"status":  "success",
		"message": "login",
		"data": response.AuthRefreshResponse{
			AccessToken: accessToken,
		},
	})

}

func (ah *AuthHandler) LogoutHandler(e echo.Context) error {
	auth := request.TokenRequest{}
	e.Bind(&auth)

	err := ah.AuthBusiness.VerifyRefreshToken(auth.ToTokenCore())
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "fail",
			"message": "cannot logout",
			"err":     err.Error(),
		})
	}

	err = ah.AuthBusiness.DeleteRefreshToken(auth.ToTokenCore())
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "fail",
			"message": "cannot logout",
			"err":     err.Error(),
		})
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"status":  "success",
		"message": "logout",
	})

}
