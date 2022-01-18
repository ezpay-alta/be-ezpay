package response

import (
	"ezpay/features/products"
)

type ProductResponse struct {
	ID      int          `json:"id"`
	Name    string       `json:"name"`
	Type    TypeResponse `json:"type"`
	Nominal int          `json:"nominal"`
	Price   int          `json:"price"`
}

type TypeResponse struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func ToProductResponse(product products.Core) ProductResponse {
	return ProductResponse{
		ID:      product.ID,
		Name:    product.Name,
		Type:    ToTypeResponse(product.Type),
		Nominal: product.Nominal,
		Price:   product.Price,
	}
}

func ToTypeResponse(typeProduct products.TypeCore) TypeResponse {
	return TypeResponse{
		ID:          typeProduct.ID,
		Name:        typeProduct.Name,
		Description: typeProduct.Description,
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
	for _, typeProduct := range typeList {
		convertedType = append(convertedType, ToTypeResponse(typeProduct))
	}

	return convertedType
}
