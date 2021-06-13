package usecase

import (
	"github.com/fgunawan1995/bcg/model"
	"github.com/pkg/errors"
)

//AddItemToCart add single item to cart (and give bonus items too)
func (u *impl) AddItemToCart(input model.AddItemToCart) error {
	// get current data
	cart, err := u.cacheDAL.GetCart(input.UserID)
	if err != nil {
		return errors.WithStack(err)
	}
	cart = cart.AddToCart(input)
	promoMapped, err := u.common.GetPromoMappedByItemID(cart.GetCartItemIDs())
	if err != nil {
		return errors.WithStack(err)
	}

	// adding bonus items
	cart = cart.ResetBonusItem()
	for _, c := range cart.Items {
		if promoMapped[c.ItemID] != nil {
			if promoMapped[c.ItemID].Type == model.PromoTypeBonus {
				cart = cart.AddBonusItem(model.BonusItem{
					ItemID: promoMapped[c.ItemID].Details.BonusItemID,
					Qty:    (c.Qty / promoMapped[c.ItemID].Details.PerQty) * promoMapped[c.ItemID].Details.BonusQty,
				})
			}
		}
	}

	// save the cart
	u.cacheDAL.SaveCart(input.UserID, cart)

	return nil
}
