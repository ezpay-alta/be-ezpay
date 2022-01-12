package presentation

import (
	"ezpay/features/promos"
	"ezpay/features/promos/presentation/request"
	"ezpay/features/promos/presentation/response"
	"net/http"

	"strconv"

	"github.com/labstack/echo/v4"
)

type PromoHandler struct {
	PromoBusiness promos.Business
}

func NewPromoHandler(promoBusiness promos.Business) *PromoHandler {
	return &PromoHandler{PromoBusiness: promoBusiness}
}

func (ph *PromoHandler) CreatePromoHandler(e echo.Context) error {
	promoData := request.PromoRequest{}

	e.Bind(&promoData)

	err := ph.PromoBusiness.CreatePromo(promoData.ToPromoCore())
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "fail",
			"message": "can not create promo",
			"err":     err.Error(),
		})
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"status":  "success",
		"message": "new promo is created",
	})
}

func (ph *PromoHandler) GetAllPromoHandler(e echo.Context) error {
	data, err := ph.PromoBusiness.GetAllPromos()
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "fail",
			"message": "can not get all promos",
			"err":     err.Error(),
		})
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"data":   response.ToProductResponseList(data),
	})

}

func (ph *PromoHandler) GetPromoByIdHandler(e echo.Context) error {
	promoId, err := strconv.Atoi(e.Param("promoId"))
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "fail",
			"message": "can not get promo",
			"err":     err.Error(),
		})
	}

	data, err := ph.PromoBusiness.GetPromoById(promoId)
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "fail",
			"message": "can not get promo",
			"err":     err.Error(),
		})
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"data":   response.ToPromoResponse(data),
	})

}

func (ah *PromoHandler) UpdatePromoByIdHandler(e echo.Context) error {
	promoId, err := strconv.Atoi(e.Param("promoId"))
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "fail",
			"message": "can not update promo",
			"err":     err.Error(),
		})
	}

	promoData := request.PromoRequest{}
	e.Bind(&promoData)

	err = ah.PromoBusiness.UpdatePromoById(promoId, promoData.ToPromoCore())
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "fail",
			"message": "can not update promo",
			"err":     err.Error(),
		})
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"status":  "success",
		"message": "update promo",
	})

}

func (uh *PromoHandler) DeletePromoByIdHandler(e echo.Context) error {
	promoId, err := strconv.Atoi(e.Param("promoId"))
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "fail",
			"message": "can not delete promo",
			"err":     err.Error(),
		})
	}

	err = uh.PromoBusiness.DeletePromoById(promoId)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "fail",
			"message": "can not delete promo",
			"err":     err.Error(),
		})
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"status":  "success",
		"message": "delete product",
	})

}
