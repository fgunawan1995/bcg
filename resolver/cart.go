package resolver

import (
	"github.com/fgunawan1995/bcg/common"
	dbdal "github.com/fgunawan1995/bcg/dal/db"
	"github.com/fgunawan1995/bcg/model"
	"github.com/fgunawan1995/bcg/util"
)

type cartResolver struct {
	common common.Common
	dbDAL  dbdal.DBDAL
	cart   model.Cart
}

func (r cartResolver) Items() ([]cartItemResolver, error) {
	var l []cartItemResolver
	result, err := r.dbDAL.GetItemByIDs(r.cart.GetCartItemIDs())
	if err != nil {
		return l, util.FormatGQLError(err)
	}
	promoMapped, err := r.common.GetPromoMappedByItemID(r.cart.GetCartItemIDs())
	if err != nil {
		return l, util.FormatGQLError(err)
	}
	for _, v := range result {
		l = append(l, cartItemResolver{
			item:    v,
			cartQty: r.cart.GetCartItemQty(v.ID),
			promo:   promoMapped[v.ID],
		})
	}
	return l, nil
}

func (r cartResolver) BonusItems() ([]bonusItemResolver, error) {
	var l []bonusItemResolver
	var err error
	result, err := r.dbDAL.GetItemByIDs(r.cart.GetBonusItemIDs())
	if err != nil {
		return l, util.FormatGQLError(err)
	}
	for _, v := range result {
		l = append(l, bonusItemResolver{
			item:     v,
			bonusQty: r.cart.GetBonusItemQty(v.ID),
		})
	}
	return l, nil
}

func (r *Resolver) Cart(args *struct {
	UserID string
}) (cartResolver, error) {
	cart, err := r.CacheDAL.GetCart(args.UserID)
	return cartResolver{
		dbDAL:  r.DBDAL,
		common: r.Common,
		cart:   cart,
	}, util.FormatGQLError(err)
}
