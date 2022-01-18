package data

import (
	"ezpay/features/products"

	"gorm.io/gorm"
)

type mysqlProductRepository struct {
	Conn *gorm.DB
}

func NewMysqlProductRepository(conn *gorm.DB) products.Data {
	return &mysqlProductRepository{
		Conn: conn,
	}
}

func (ar *mysqlProductRepository) CreateType(productType products.TypeCore) error {
	recordData := ToTypeRecord(productType)
	err := ar.Conn.Create(&recordData).Error
	if err != nil {
		return err
	}

	return nil
}

func (ar *mysqlProductRepository) GetAllTypes() ([]products.TypeCore, error) {

	types := []Type{}
	err := ar.Conn.Find(&types).Error
	if err != nil {
		return nil, err
	}
	return ToTypeProductCoreList(types), nil
}

func (ar *mysqlProductRepository) GetTypeById(typeId int) (products.TypeCore, error) {

	productType := Type{}
	err := ar.Conn.First(&productType, typeId).Error
	if err != nil {
		return ToTypeCore(Type{}), err
	}
	return ToTypeCore(productType), nil
}

func (ar *mysqlProductRepository) UpdateTypeById(typeId int, data products.TypeCore) error {

	productType := ToTypeRecord(data)
	productType.ID = typeId

	err := ar.Conn.First(&productType, typeId).Error
	if err != nil {
		return err
	}

	if data.Name != "" {
		productType.Name = data.Name
	}
	if data.Description != "" {
		productType.Description = data.Description
	}

	err = ar.Conn.Save(&productType).Error
	if err != nil {
		return err
	}
	return nil
}

func (ar *mysqlProductRepository) DeleteTypeById(typeId int) error {
	err := ar.Conn.Delete(&Type{}, typeId).Error
	if err != nil {
		return err
	}
	return nil
}

func (ar *mysqlProductRepository) CreateProduct(data products.Core) error {

	recordData := ToProductRecord(data)
	err := ar.Conn.Create(&recordData).Error
	if err != nil {
		return err
	}
	return nil
}

func (ar *mysqlProductRepository) GetAllProducts() ([]products.Core, error) {

	products := []Product{}
	err := ar.Conn.Joins("Type").Find(&products).Error
	if err != nil {
		return nil, err
	}
	return ToProductCoreList(products), nil
}

func (ar *mysqlProductRepository) GetProductById(productId int) (products.Core, error) {

	product := Product{}
	err := ar.Conn.Joins("Type").First(&product, productId).Error
	if err != nil {
		return ToProductCore(Product{}), err
	}
	return ToProductCore(product), nil
}

func (ar *mysqlProductRepository) UpdateProductById(productId int, data products.Core) error {

	product := ToProductRecord(data)
	product.ID = productId

	err := ar.Conn.First(&product, productId).Error
	if err != nil {
		return err
	}

	if data.Name != "" {
		product.Name = data.Name
	}
	if data.Nominal != 0 {
		product.Nominal = data.Nominal
	}
	if data.Price != 0 {
		product.Price = data.Price
	}
	if data.TypeID != 0 {
		product.TypeID = data.TypeID
	}

	err = ar.Conn.Save(&product).Error
	if err != nil {
		return err
	}
	return nil
}

func (ar *mysqlProductRepository) DeleteProductById(productId int) error {

	err := ar.Conn.Delete(&Product{}, productId).Error
	if err != nil {
		return err
	}
	return nil
}
