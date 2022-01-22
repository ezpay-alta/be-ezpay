package presentation

import (
	"ezpay/features/middlewares"
	"ezpay/features/products"
	"ezpay/features/products/presentation/request"
	"ezpay/features/products/presentation/response"
	"net/http"

	"strconv"

	"github.com/labstack/echo/v4"
)

type ProductHandler struct {
	ProductBusiness products.Business
}

func NewProductHandler(productBusiness products.Business) *ProductHandler {
	return &ProductHandler{ProductBusiness: productBusiness}
}

func (ph *ProductHandler) CreateProductHandler(e echo.Context) error {
	_, role, err := middlewares.VerifyAccessToken(e)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "fail",
			"message": "can not create product",
			"err":     "token is invalid",
		})
	}
	if role != "admin" {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "fail",
			"message": "can not create product",
			"err":     "you are not an admin",
		})
	}

	productData := request.ProductRequest{}

	e.Bind(&productData)

	err = ph.ProductBusiness.CreateProduct(productData.ToProductCore())
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "fail",
			"message": "can not create product",
			"err":     err.Error(),
		})
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"status":  "success",
		"message": "new product is created",
	})
}

func (ph *ProductHandler) GetAllProductHandler(e echo.Context) error {
	data, err := ph.ProductBusiness.GetAllProducts()
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "fail",
			"message": "can not get all products",
			"err":     err.Error(),
		})
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"data":   response.ToProductResponseList(data),
	})

}

func (ph *ProductHandler) GetProductByIdHandler(e echo.Context) error {
	productId, err := strconv.Atoi(e.Param("productId"))
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "fail",
			"message": "can not get product",
			"err":     err.Error(),
		})
	}

	data, err := ph.ProductBusiness.GetProductById(productId)
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "fail",
			"message": "can not get product",
			"err":     err.Error(),
		})
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"data":   response.ToProductResponse(data),
	})

}

func (ah *ProductHandler) UpdateProductByIdHandler(e echo.Context) error {
	_, role, err := middlewares.VerifyAccessToken(e)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "fail",
			"message": "can not create product",
			"err":     "token is invalid",
		})
	}
	if role != "admin" {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "fail",
			"message": "can not create product",
			"err":     "you are not an admin",
		})
	}

	productId, err := strconv.Atoi(e.Param("productId"))
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "fail",
			"message": "can not update product",
			"err":     err.Error(),
		})
	}

	productData := request.ProductRequest{}
	e.Bind(&productData)

	err = ah.ProductBusiness.UpdateProductById(productId, productData.ToProductCore())
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "fail",
			"message": "can not update product",
			"err":     err.Error(),
		})
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"status":  "success",
		"message": "update product",
	})

}

func (uh *ProductHandler) DeleteProductByIdHandler(e echo.Context) error {
	_, role, err := middlewares.VerifyAccessToken(e)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "fail",
			"message": "can not create product",
			"err":     "token is invalid",
		})
	}
	if role != "admin" {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "fail",
			"message": "can not create product",
			"err":     "you are not an admin",
		})
	}

	productId, err := strconv.Atoi(e.Param("productId"))
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "fail",
			"message": "can not delete product",
			"err":     err.Error(),
		})
	}

	err = uh.ProductBusiness.DeleteProductById(productId)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "fail",
			"message": "can not delete product",
			"err":     err.Error(),
		})
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"status":  "success",
		"message": "delete product",
	})

}

func (ph *ProductHandler) CreateProductTypeHandler(e echo.Context) error {
	_, role, err := middlewares.VerifyAccessToken(e)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "fail",
			"message": "can not create product type",
			"err":     "token is invalid",
		})
	}
	if role != "admin" {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "fail",
			"message": "can not create product type",
			"err":     "you are not an admin",
		})
	}

	typeData := request.TypeRequest{}

	e.Bind(&typeData)

	err = ph.ProductBusiness.CreateType(typeData.ToTypeCore())
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "fail",
			"message": "can not create product type",
			"err":     err.Error(),
		})
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"status":  "success",
		"message": "new product type is created",
	})
}

func (ph *ProductHandler) GetAllTypeProductHandler(e echo.Context) error {
	data, err := ph.ProductBusiness.GetAllTypes()
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "fail",
			"message": "can not get all products type",
			"err":     err.Error(),
		})
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"data":   response.ToTypeResponseList(data),
	})

}

func (ph *ProductHandler) GetProductTypeByIdHandler(e echo.Context) error {
	typeId, err := strconv.Atoi(e.Param("typeId"))
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "fail",
			"message": "can not get product type",
			"err":     err.Error(),
		})
	}

	data, err := ph.ProductBusiness.GetTypeById(typeId)
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "fail",
			"message": "can not get product type",
			"err":     err.Error(),
		})
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"data":   response.ToTypeResponse(data),
	})

}

func (ah *ProductHandler) UpdateProductTypeByIdHandler(e echo.Context) error {

	typeId, err := strconv.Atoi(e.Param("typeId"))
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "fail",
			"message": "can not update product type",
			"err":     err.Error(),
		})
	}

	productData := request.TypeRequest{}
	e.Bind(&productData)

	_, role, err := middlewares.VerifyAccessToken(e)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "fail",
			"message": "can not create product type",
			"err":     "token is invalid",
		})
	}
	if role != "admin" {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "fail",
			"message": "can not create product type",
			"err":     "you are not an admin",
		})
	}

	err = ah.ProductBusiness.UpdateTypeById(typeId, productData.ToTypeCore())
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "fail",
			"message": "can not update product type",
			"err":     err.Error(),
		})
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"status":  "success",
		"message": "update product type",
	})

}

func (uh *ProductHandler) DeleteProductTypeByIdHandler(e echo.Context) error {
	_, role, err := middlewares.VerifyAccessToken(e)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "fail",
			"message": "can not create product type",
			"err":     "token is invalid",
		})
	}
	if role != "admin" {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "fail",
			"message": "can not create product type",
			"err":     "you are not an admin",
		})
	}

	typeId, err := strconv.Atoi(e.Param("typeId"))
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "fail",
			"message": "can not delete product type",
			"err":     err.Error(),
		})
	}

	err = uh.ProductBusiness.DeleteTypeById(typeId)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "fail",
			"message": "can not delete product type",
			"err":     err.Error(),
		})
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"status":  "success",
		"message": "delete product type",
	})

}
