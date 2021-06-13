package resolver

import (
	"github.com/fgunawan1995/bcg/model"
)

type bonusItemResolver struct {
	item     model.Item
	bonusQty int32
}

func (r bonusItemResolver) Item() (itemResolver, error) {
	return itemResolver{r.item}, nil
}

func (r bonusItemResolver) BonusQty() (int32, error) {
	return r.bonusQty, nil
}
