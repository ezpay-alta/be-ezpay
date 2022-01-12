package data

import (
	productData "ezpay/features/products/data"
	"ezpay/features/promos"

	"gorm.io/gorm"
)

type Promo struct {
	gorm.Model
	ID        int
	ProductID int
	Product   productData.Product
	Label     string
	Amount    int
}

func ToPromoRecord(promo promos.Core) Promo {
	return Promo{
		ID:        promo.ID,
		ProductID: promo.ProductID,
		Product:   productData.ToProductRecord(promo.Product),
		Label:     promo.Label,
		Amount:    promo.Amount,
	}
}

func ToPromoCore(promo Promo) promos.Core {
	return promos.Core{
		ID:        promo.ID,
		ProductID: promo.ProductID,
		Product:   productData.ToProductCore(promo.Product),
		Label:     promo.Label,
		Amount:    promo.Amount,
	}
}

func ToPromoCoreList(pList []Promo) []promos.Core {
	convertedPromo := []promos.Core{}

	for _, promo := range pList {
		convertedPromo = append(convertedPromo, ToPromoCore(promo))
	}

	return convertedPromo
}
