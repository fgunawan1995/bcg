package usecase

import (
	"github.com/fgunawan1995/bcg/common"
	cachedal "github.com/fgunawan1995/bcg/dal/cache"
	dbdal "github.com/fgunawan1995/bcg/dal/db"
	"github.com/fgunawan1995/bcg/model"
)

type impl struct {
	common   common.Common
	dbDAL    dbdal.DBDAL
	cacheDAL cachedal.CacheDAL
}

type Usecase interface {
	//AddItemToCart add single item to cart (and give bonus items too)
	AddItemToCart(input model.AddItemToCart) error
	//Checkout checkout current cart content
	Checkout(userID string) (string, error)
}

func New(common common.Common, cache cachedal.CacheDAL, db dbdal.DBDAL) Usecase {
	return &impl{
		common:   common,
		dbDAL:    db,
		cacheDAL: cache,
	}
}
