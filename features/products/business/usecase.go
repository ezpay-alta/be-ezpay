package business

import (
	"ezpay/features/products"
)

type productUsecase struct {
	ProductData products.Data
}

func NewProductBusiness(productData products.Data) products.Business {
	return &productUsecase{ProductData: productData}
}

func (pu *productUsecase) CreateType(typeProduct products.TypeCore) error {
	err := pu.ProductData.CreateType(typeProduct)
	if err != nil {
		return err
	}

	return nil
}

func (pu *productUsecase) GetAllTypes() ([]products.TypeCore, error) {
	types, err := pu.ProductData.GetAllTypes()
	if err != nil {
		return nil, err
	}

	return types, nil
}

func (pu *productUsecase) GetTypeById(typeId int) (products.TypeCore, error) {
	productType, err := pu.ProductData.GetTypeById(typeId)
	if err != nil {
		return products.TypeCore{}, err
	}

	return productType, nil
}

func (pu *productUsecase) UpdateTypeById(typeId int, data products.TypeCore) error {
	err := pu.ProductData.UpdateTypeById(typeId, data)
	if err != nil {
		return err
	}

	return nil
}

func (pu *productUsecase) DeleteTypeById(typeId int) error {
	err := pu.ProductData.DeleteTypeById(typeId)
	if err != nil {
		return err
	}

	return nil
}

func (pu *productUsecase) CreateProduct(product products.Core) error {
	err := pu.ProductData.CreateProduct(product)
	if err != nil {
		return err
	}

	return nil
}

func (pu *productUsecase) GetAllProducts() ([]products.Core, error) {
	products, err := pu.ProductData.GetAllProducts()
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (pu *productUsecase) GetProductById(productId int) (products.Core, error) {
	product, err := pu.ProductData.GetProductById(productId)
	if err != nil {
		return products.Core{}, err
	}

	return product, nil
}

func (pu *productUsecase) UpdateProductById(productId int, data products.Core) error {
	err := pu.ProductData.UpdateProductById(productId, data)
	if err != nil {
		return err
	}

	return nil
}

func (pu *productUsecase) DeleteProductById(productId int) error {
	err := pu.ProductData.DeleteProductById(productId)
	if err != nil {
		return err
	}

	return nil
}
