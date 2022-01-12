package response

import (
	productResponse "ezpay/features/products/presentation/response"
	"ezpay/features/promos"
)

type PromoResponse struct {
	ID      int    `json:"id"`
	Label   string `json:"label"`
	Product productResponse.ProductResponse
	Amount  int `json:"amount"`
}

func ToPromoResponse(promo promos.Core) PromoResponse {
	return PromoResponse{
		ID:      promo.ID,
		Label:   promo.Label,
		Product: productResponse.ToProductResponse(promo.Product),
		Amount:  promo.Amount,
	}
}

func ToProductResponseList(promoList []promos.Core) []PromoResponse {
	convertedPromo := []PromoResponse{}
	for _, promo := range promoList {
		convertedPromo = append(convertedPromo, ToPromoResponse(promo))
	}

	return convertedPromo
}
