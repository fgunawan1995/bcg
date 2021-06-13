package resolver

import (
	"github.com/fgunawan1995/bcg/model"
)

type cartItemResolver struct {
	item    model.Item
	promo   *model.Promo
	cartQty int32
}

func (r cartItemResolver) Item() (itemResolver, error) {
	return itemResolver{r.item}, nil
}

func (r cartItemResolver) Promo() (*promoResolver, error) {
	if r.promo == nil {
		return nil, nil
	}
	return &promoResolver{*r.promo}, nil
}

func (r cartItemResolver) CartQty() (int32, error) {
	return r.cartQty, nil
}

func (r cartItemResolver) SubTotal() (float64, error) {
	return r.item.Price * float64(r.cartQty), nil
}

func (r cartItemResolver) DiscTotal() (float64, error) {
	var discTotal float64
	subTotal := r.item.Price * float64(r.cartQty)
	if r.promo != nil {
		if r.promo.Type == model.PromoTypeDiscount {
			if r.cartQty >= r.promo.Details.MinQty {
				discTotal = subTotal * r.promo.Details.DiscountPercentage
			}
		}
	}
	return discTotal, nil
}

func (r cartItemResolver) Total() (float64, error) {
	var discTotal float64
	subTotal := r.item.Price * float64(r.cartQty)
	if r.promo != nil {
		if r.promo.Type == model.PromoTypeDiscount {
			if r.cartQty >= r.promo.Details.MinQty {
				discTotal = subTotal * r.promo.Details.DiscountPercentage
			}
		}
	}
	return subTotal - discTotal, nil
}
