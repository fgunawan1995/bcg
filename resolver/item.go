package resolver

import "github.com/fgunawan1995/bcg/model"

type itemResolver struct {
	item model.Item
}

func (r itemResolver) ID() string {
	return r.item.ID
}

func (r itemResolver) Name() string {
	return r.item.Name
}

func (r itemResolver) SKU() string {
	return r.item.SKU
}

func (r itemResolver) Price() float64 {
	return r.item.Price
}

func (r itemResolver) CurrentStock() int32 {
	return r.item.Qty
}
