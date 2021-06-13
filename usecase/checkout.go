package usecase

import (
	"fmt"

	"github.com/pkg/errors"

	"github.com/fgunawan1995/bcg/model"
	"github.com/fgunawan1995/bcg/util"
)

//Checkout checkout current cart content
func (u *impl) Checkout(userID string) (string, error) {
	// get all necessary data
	cart, err := u.cacheDAL.GetCart(userID)
	if err != nil {
		return "", errors.WithStack(err)
	}
	if len(cart.Items) == 0 {
		return "", errors.WithStack(errors.New(model.EmptyCartMessage))
	}
	allItems, err := u.dbDAL.GetItemByIDs(append(cart.GetCartItemIDs(), cart.GetBonusItemIDs()...))
	if err != nil {
		return "", errors.WithStack(err)
	}
	cartItems := filterItems(cart.GetCartItemIDs(), allItems)
	promoMapped, err := u.common.GetPromoMappedByItemID(cart.GetCartItemIDs())
	if err != nil {
		return "", errors.WithStack(err)
	}

	// validation
	err = validateItemAvailability(cart, allItems)
	if err != nil {
		return "", errors.WithStack(err)
	}

	// reduce stock and empty cart
	err = util.WithTransaction(u.dbDAL.GetDB(), func(tx util.Transaction) error {
		for itemID, qty := range cart.GetQtyMap() {
			err = u.dbDAL.ReduceItemStock(tx, itemID, qty)
			if err != nil {
				return errors.WithStack(err)
			}
		}
		return nil
	})
	if err != nil {
		return "", errors.WithStack(err)
	}
	u.cacheDAL.EmptyCart(userID)

	// output message
	total := calculateTotal(cart, cartItems, promoMapped)
	return fmt.Sprintf(model.CheckoutMessage, total), nil
}

func filterItems(itemIDs []string, items []model.Item) []model.Item {
	result := make([]model.Item, 0)
	for _, itemID := range itemIDs {
		for _, item := range items {
			if itemID == item.ID {
				result = append(result, item)
			}
		}
	}
	return result
}

func calculateTotal(cart model.Cart, cartItems []model.Item, promoMapped map[string]*model.Promo) float64 {
	var result float64
	for _, cI := range cart.Items {
		for _, i := range cartItems {
			if cI.ItemID == i.ID {
				var iSubTotal, iDiscTotal, iTotal float64
				iSubTotal = i.Price * float64(cI.Qty)
				if promoMapped[cI.ItemID] != nil {
					if promoMapped[cI.ItemID].Type == model.PromoTypeDiscount {
						if cI.Qty >= promoMapped[cI.ItemID].Details.MinQty {
							iDiscTotal = iSubTotal * promoMapped[cI.ItemID].Details.DiscountPercentage
						}
					}
				}
				iTotal = iSubTotal - iDiscTotal
				result += iTotal
			}
		}
	}
	return result
}

func validateItemAvailability(cart model.Cart, items []model.Item) error {
	qtyMap := cart.GetQtyMap()
	for itemID, cartQty := range qtyMap {
		for _, item := range items {
			if itemID == item.ID && cartQty > item.Qty {
				return errors.WithStack(fmt.Errorf(model.OutOfStockMessage, item.Name, item.Qty))
			}
		}
	}
	return nil
}
