package db

import (
	"github.com/fgunawan1995/bcg/model"
	"github.com/fgunawan1995/bcg/util"
	"github.com/jmoiron/sqlx"
)

type DBDAL interface {
	//GetItemByIDs from db and return array of items
	GetItemByIDs(itemIDs []string) ([]model.Item, error)
	//GetItemPromoByItemIDs get itempromo relation table for item-promo
	GetItemPromoByItemIDs(itemIDs []string) ([]model.ItemPromo, error)
	//GetPromoByIDs get promo by ids given and give back array of promo
	GetPromoByIDs(promoIDs []string) ([]model.Promo, error)
	//ReduceItemStock substract qty from table 'items'
	ReduceItemStock(tx util.Transaction, itemID string, qty int32) error
	//GetDB for used withing util.WithTransaction
	GetDB() *sqlx.DB
}

type impl struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) DBDAL {
	return &impl{
		db: db,
	}
}

//GetDB for used withing util.WithTransaction
func (dal *impl) GetDB() *sqlx.DB {
	return dal.db
}
