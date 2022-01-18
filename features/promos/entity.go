package promos

import (
	"ezpay/features/products"
	"time"
)

type Core struct {
	ID        int
	ProductID int
	Product   products.Core
	Label     string
	Amount    int
	Expires   time.Time
}

type Business interface {
	CreatePromo(data Core) error
	GetAllPromos() ([]Core, error)
	GetPromoById(promoId int) (Core, error)
	UpdatePromoById(promoId int, data Core) error
	DeletePromoById(promoId int) error
}

type Data interface {
	CreatePromo(data Core) error
	GetAllPromos() ([]Core, error)
	GetPromoById(promoId int) (Core, error)
	UpdatePromoById(promoId int, data Core) error
	DeletePromoById(promoId int) error
}
