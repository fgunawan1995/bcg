package db

import (
	"github.com/fgunawan1995/bcg/model"
	"github.com/fgunawan1995/bcg/util"
	"github.com/lib/pq"
	"github.com/pkg/errors"
)

//GetItemPromoByItemIDs get itempromo relation table for item-promo
func (dal *impl) GetItemPromoByItemIDs(itemIDs []string) ([]model.ItemPromo, error) {
	var result []model.ItemPromo
	err := dal.db.Select(&result, getItemPromosByItemIDs, pq.Array(util.Distinct(itemIDs)))
	if err != nil {
		return result, errors.WithStack(err)
	}
	return result, err
}

//GetPromoByIDs get promo by ids given and give back array of promo
func (dal *impl) GetPromoByIDs(promoIDs []string) ([]model.Promo, error) {
	var result []model.Promo
	err := dal.db.Select(&result, getPromosByIDs, pq.Array(util.Distinct(promoIDs)))
	if err != nil {
		return result, errors.WithStack(err)
	}
	return result, err
}
