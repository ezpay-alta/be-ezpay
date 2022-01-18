package data

import (
	"ezpay/features/promos"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type mysqlPromoRepository struct {
	Conn *gorm.DB
}

func NewMysqlPromoRepository(conn *gorm.DB) promos.Data {
	return &mysqlPromoRepository{
		Conn: conn,
	}
}

func (ar *mysqlPromoRepository) CreatePromo(data promos.Core) error {
	recordData := ToPromoRecord(data)
	err := ar.Conn.Create(&recordData).Error
	if err != nil {
		return err
	}

	return nil
}

func (ar *mysqlPromoRepository) GetAllPromos() ([]promos.Core, error) {

	promos := []Promo{}
	err := ar.Conn.Preload("Product.Type").Preload(clause.Associations).Find(&promos).Error
	if err != nil {
		return nil, err
	}
	return ToPromoCoreList(promos), nil
}

func (ar *mysqlPromoRepository) GetPromoById(promoId int) (promos.Core, error) {

	promo := Promo{}
	err := ar.Conn.Preload("Product.Type").Preload(clause.Associations).First(&promo, promoId).Error
	if err != nil {
		return ToPromoCore(Promo{}), err
	}
	return ToPromoCore(promo), nil
}

func (ar *mysqlPromoRepository) UpdatePromoById(promoId int, data promos.Core) error {

	promo := ToPromoRecord(data)
	promo.ID = promoId

	err := ar.Conn.First(&promo, promoId).Error
	if err != nil {
		return err
	}

	if data.ProductID != 0 {
		promo.ProductID = data.ProductID
	}
	if data.Label != "" {
		promo.Label = data.Label
	}
	if data.Amount != 0 {
		promo.Amount = data.Amount
	}

	err = ar.Conn.Save(&promo).Error
	if err != nil {
		return err
	}
	return nil
}

func (ar *mysqlPromoRepository) DeletePromoById(promoId int) error {
	err := ar.Conn.Delete(&Promo{}, promoId).Error
	if err != nil {
		return err
	}
	return nil
}
