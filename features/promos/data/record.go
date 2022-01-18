package data

import (
	productData "ezpay/features/products/data"
	"ezpay/features/promos"
	"time"

	"gorm.io/gorm"
)

type Promo struct {
	gorm.Model
	ID        int
	ProductID int
	Product   productData.Product
	Label     string
	Amount    int
	Expires   time.Time
}

func ToPromoRecord(promo promos.Core) Promo {
	return Promo{
		ID:        promo.ID,
		ProductID: promo.ProductID,
		Product:   productData.ToProductRecord(promo.Product),
		Label:     promo.Label,
		Amount:    promo.Amount,
		Expires:   promo.Expires,
	}
}

func ToPromoCore(promo Promo) promos.Core {
	return promos.Core{
		ID:        promo.ID,
		ProductID: promo.ProductID,
		Product:   productData.ToProductCore(promo.Product),
		Label:     promo.Label,
		Amount:    promo.Amount,
		Expires:   promo.Expires,
	}
}

func ToPromoCoreList(pList []Promo) []promos.Core {
	convertedPromo := []promos.Core{}

	for _, promo := range pList {
		convertedPromo = append(convertedPromo, ToPromoCore(promo))
	}

	return convertedPromo
}
