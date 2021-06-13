package common

import (
	"github.com/fgunawan1995/bcg/model"
	"github.com/pkg/errors"
)

//GetPromoMappedByItemID helper method to easily get promo mapped by item_id
func (c *impl) GetPromoMappedByItemID(itemIDs []string) (map[string]*model.Promo, error) {
	result := make(map[string]*model.Promo)
	itemPromos, err := c.dbDAL.GetItemPromoByItemIDs(itemIDs)
	if err != nil {
		return result, errors.WithStack(err)
	}
	promoIDs := make([]string, 0)
	for _, val := range itemPromos {
		promoIDs = append(promoIDs, val.PromoID)
	}
	promos, err := c.dbDAL.GetPromoByIDs(promoIDs)
	if err != nil {
		return result, errors.WithStack(err)
	}
	for _, itemPromo := range itemPromos {
		for _, promo := range promos {
			temp := promo
			if itemPromo.PromoID == promo.ID {
				result[itemPromo.ItemID] = &temp
			}
		}
	}
	return result, nil
}
