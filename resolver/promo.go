package resolver

import "github.com/fgunawan1995/bcg/model"

type promoResolver struct {
	promo model.Promo
}

func (r promoResolver) ID() string {
	return r.promo.ID
}

func (r promoResolver) Name() string {
	return r.promo.Name
}
