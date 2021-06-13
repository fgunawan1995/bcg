package resolver

import (
	"github.com/fgunawan1995/bcg/model"
	"github.com/fgunawan1995/bcg/util"
)

func (r *Resolver) AddToCart(args *struct {
	Add model.AddItemToCart
}) (cartResolver, error) {
	err := r.Usecase.AddItemToCart(args.Add)
	if err != nil {
		return cartResolver{}, util.FormatGQLError(err)
	}
	cart, err := r.CacheDAL.GetCart(args.Add.UserID)
	return cartResolver{
		dbDAL:  r.DBDAL,
		common: r.Common,
		cart:   cart,
	}, util.FormatGQLError(err)
}
