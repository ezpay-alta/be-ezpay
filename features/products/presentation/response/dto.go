package response

import (
	"ezpay/features/products"
)

type ProductResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type TypeResponse struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func ToProductResponse(product products.Core) ProductResponse {
	return ProductResponse{
		ID:   product.ID,
		Name: product.Name,
	}
}

func ToTypeResponse(productType products.TypeCore) TypeResponse {
	return TypeResponse{
		ID:          productType.ID,
		Name:        productType.Name,
		Description: productType.Description,
	}
}

func ToProductResponseList(productList []products.Core) []ProductResponse {
	convertedProduct := []ProductResponse{}
	for _, product := range productList {
		convertedProduct = append(convertedProduct, ToProductResponse(product))
	}

	return convertedProduct
}

func ToTypeResponseList(typeList []products.TypeCore) []TypeResponse {
	convertedType := []TypeResponse{}
	for _, productType := range typeList {
		convertedType = append(convertedType, ToTypeResponse(productType))
	}

	return convertedType
}
