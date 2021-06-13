package common

import (
	cachedal "github.com/fgunawan1995/bcg/dal/cache"
	dbdal "github.com/fgunawan1995/bcg/dal/db"
	"github.com/fgunawan1995/bcg/model"
)

type impl struct {
	dbDAL    dbdal.DBDAL
	cacheDAL cachedal.CacheDAL
}

type Common interface {
	//GetPromoMappedByItemID helper method to easily get promo mapped by item_id
	GetPromoMappedByItemID(itemIDs []string) (map[string]*model.Promo, error)
}

func New(cache cachedal.CacheDAL, db dbdal.DBDAL) Common {
	return &impl{
		dbDAL:    db,
		cacheDAL: cache,
	}
}
