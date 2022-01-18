package request

import (
	"ezpay/features/promos"
	"time"
)

type PromoRequest struct {
	ProductID int       `json:"product_id"`
	Label     string    `json:"label"`
	Amount    int       `json:"amount"`
	Expires   time.Time `json:"expires"`
}

func (requestData *PromoRequest) ToPromoCore() promos.Core {
	return promos.Core{
		ProductID: requestData.ProductID,
		Label:     requestData.Label,
		Amount:    requestData.Amount,
		Expires:   requestData.Expires,
	}
}
