package db

import (
	"github.com/fgunawan1995/bcg/model"
	"github.com/fgunawan1995/bcg/util"
	"github.com/lib/pq"
	"github.com/pkg/errors"
)

//GetItemByIDs from db and return array of items
func (dal *impl) GetItemByIDs(itemIDs []string) ([]model.Item, error) {
	var result []model.Item
	err := dal.db.Select(&result, getItemsByIDs, pq.Array(util.Distinct(itemIDs)))
	if err != nil {
		return result, errors.WithStack(err)
	}
	return result, nil
}

//ReduceItemStock substract qty from table 'items'
func (dal *impl) ReduceItemStock(tx util.Transaction, itemID string, qty int32) error {
	_, err := tx.Exec(reduceItemStockByID, itemID, qty)
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}
