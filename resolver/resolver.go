package resolver

import (
	"github.com/fgunawan1995/bcg/common"
	cachedal "github.com/fgunawan1995/bcg/dal/cache"
	dbdal "github.com/fgunawan1995/bcg/dal/db"
	"github.com/fgunawan1995/bcg/usecase"
)

type Resolver struct {
	DBDAL    dbdal.DBDAL
	CacheDAL cachedal.CacheDAL
	Usecase  usecase.Usecase
	Common   common.Common
}

func New(usecase usecase.Usecase, common common.Common, cache cachedal.CacheDAL, db dbdal.DBDAL) *Resolver {
	return &Resolver{
		DBDAL:    db,
		CacheDAL: cache,
		Usecase:  usecase,
		Common:   common,
	}
}

func (r *Resolver) Hello() string {
	return "Hello, world!"
}
