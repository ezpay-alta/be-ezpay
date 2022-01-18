package data

import (
	"ezpay/features/products"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	ID      int
	Name    string
	TypeID  int
	Type    Type
	Nominal int
	Price   int
}

type Type struct {
	gorm.Model
	ID          int
	Name        string `gorm:"unique"`
	Description string
}

func ToTypeRecord(productType products.TypeCore) Type {
	return Type{
		ID:          productType.ID,
		Name:        productType.Name,
		Description: productType.Description,
	}
}

func ToProductRecord(product products.Core) Product {
	return Product{
		ID:      product.ID,
		Name:    product.Name,
		TypeID:  product.TypeID,
		Type:    ToTypeRecord(product.Type),
		Nominal: product.Nominal,
		Price:   product.Price,
	}
}

func ToProductCore(product Product) products.Core {
	return products.Core{
		ID:      product.ID,
		Name:    product.Name,
		TypeID:  product.TypeID,
		Type:    ToTypeCore(product.Type),
		Nominal: product.Nominal,
		Price:   product.Price,
	}
}

func ToTypeCore(productType Type) products.TypeCore {
	return products.TypeCore{
		ID:          productType.ID,
		Name:        productType.Name,
		Description: productType.Description,
	}
}

func ToProductCoreList(aList []Product) []products.Core {
	convertedProduct := []products.Core{}

	for _, product := range aList {
		convertedProduct = append(convertedProduct, ToProductCore(product))
	}

	return convertedProduct
}

func ToTypeProductCoreList(aList []Type) []products.TypeCore {
	convertedTypeProduct := []products.TypeCore{}

	for _, productType := range aList {
		convertedTypeProduct = append(convertedTypeProduct, ToTypeCore(productType))
	}

	return convertedTypeProduct
}
