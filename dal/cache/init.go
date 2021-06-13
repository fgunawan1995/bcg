package cache

import (
	"github.com/fgunawan1995/bcg/model"
	cache "github.com/patrickmn/go-cache"
)

type CacheDAL interface {
	//GetCart get cart content per user_id
	GetCart(userID string) (model.Cart, error)
	//SaveCart save cart content to cache
	SaveCart(userID string, cart model.Cart)
	//EmptyCart save empty cart content to cache
	EmptyCart(userID string)
}

type impl struct {
	cacheObj *cache.Cache
}

func New(cache *cache.Cache) CacheDAL {
	return &impl{
		cacheObj: cache,
	}
}
