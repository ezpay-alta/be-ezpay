package business

import (
	"ezpay/features/promos"
)

type promoUsecase struct {
	PromoData promos.Data
}

func NewPromoBusiness(promoData promos.Data) promos.Business {
	return &promoUsecase{PromoData: promoData}
}

func (pu *promoUsecase) CreatePromo(data promos.Core) error {
	err := pu.PromoData.CreatePromo(data)
	if err != nil {
		return err
	}

	return nil
}

func (pu *promoUsecase) GetAllPromos() ([]promos.Core, error) {
	promos, err := pu.PromoData.GetAllPromos()
	if err != nil {
		return nil, err
	}

	return promos, nil
}

func (pu *promoUsecase) GetPromoById(promoId int) (promos.Core, error) {
	promo, err := pu.PromoData.GetPromoById(promoId)
	if err != nil {
		return promos.Core{}, err
	}

	return promo, nil
}

func (pu *promoUsecase) UpdatePromoById(promoId int, data promos.Core) error {
	err := pu.PromoData.UpdatePromoById(promoId, data)
	if err != nil {
		return err
	}

	return nil
}

func (pu *promoUsecase) DeletePromoById(promoId int) error {
	err := pu.PromoData.DeletePromoById(promoId)
	if err != nil {
		return err
	}

	return nil
}
