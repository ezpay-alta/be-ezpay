package request

import (
	"ezpay/features/products"
)

type ProductRequest struct {
	Name    string `json:"name"`
	TypeID  int    `json:"type_id"`
	Nominal int    `json:"nominal"`
	Price   int    `json:"price"`
}

type TypeRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (requestData *ProductRequest) ToProductCore() products.Core {
	return products.Core{
		Name:    requestData.Name,
		TypeID:  requestData.TypeID,
		Nominal: requestData.Nominal,
		Price:   requestData.Price,
	}
}

func (requestData *TypeRequest) ToTypeCore() products.TypeCore {
	return products.TypeCore{
		Name:        requestData.Name,
		Description: requestData.Description,
	}
}
