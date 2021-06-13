package cache

import (
	"fmt"

	"github.com/pkg/errors"

	"github.com/fgunawan1995/bcg/model"
)

//GetCart get cart content per user_id
func (dal *impl) GetCart(userID string) (model.Cart, error) {
	var result model.Cart
	temp, found := dal.cacheObj.Get(fmt.Sprintf(model.KeyCart, userID))
	if !found {
		return result, nil
	}
	result, ok := temp.(model.Cart)
	if !ok {
		return result, errors.WithStack(errors.New("failed conversion"))
	}
	return result, nil
}

//SaveCart save cart content to cache
func (dal *impl) SaveCart(userID string, cart model.Cart) {
	dal.cacheObj.SetDefault(fmt.Sprintf(model.KeyCart, userID), cart)
}

//EmptyCart save empty cart content to cache
func (dal *impl) EmptyCart(userID string) {
	dal.cacheObj.SetDefault(fmt.Sprintf(model.KeyCart, userID), model.Cart{})
}
