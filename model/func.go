package model

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
)

func (p Cart) GetCartItemIDs() []string {
	result := make([]string, 0)
	for _, v := range p.Items {
		result = append(result, v.ItemID)
	}
	return result
}

func (p Cart) GetBonusItemIDs() []string {
	result := make([]string, 0)
	for _, v := range p.Bonus {
		result = append(result, v.ItemID)
	}
	return result
}

func (p Cart) GetCartItemQty(itemID string) int32 {
	for _, v := range p.Items {
		if v.ItemID == itemID {
			return v.Qty
		}
	}
	return 0
}

func (p Cart) GetBonusItemQty(itemID string) int32 {
	for _, v := range p.Bonus {
		if v.ItemID == itemID {
			return v.Qty
		}
	}
	return 0
}

func (p Cart) AddToCart(input AddItemToCart) Cart {
	found := false
	for i, v := range p.Items {
		if v.ItemID == input.ItemID {
			p.Items[i].Qty += input.Qty
			found = true
		}
	}
	if !found {
		p.Items = append(p.Items, CartItem{
			ItemID: input.ItemID,
			Qty:    input.Qty,
		})
	}
	for i, v := range p.Items {
		if v.Qty <= 0 {
			p.Items = removeCartItem(p.Items, i)
		}
	}
	return p
}

func (p Cart) ResetBonusItem() Cart {
	p.Bonus = make([]BonusItem, 0)
	return p
}

func (p Cart) AddBonusItem(input BonusItem) Cart {
	found := false
	for i, v := range p.Bonus {
		if v.ItemID == input.ItemID {
			p.Bonus[i].Qty += input.Qty
			found = true
		}
	}
	if !found {
		p.Bonus = append(p.Bonus, BonusItem{
			ItemID: input.ItemID,
			Qty:    input.Qty,
		})
	}
	for i, v := range p.Bonus {
		if v.Qty <= 0 {
			p.Bonus = removeBonusItem(p.Bonus, i)
		}
	}
	return p
}

func (p Cart) GetQtyMap() map[string]int32 {
	result := make(map[string]int32)
	for _, cI := range p.Items {
		result[cI.ItemID] += cI.Qty
	}
	for _, cB := range p.Bonus {
		result[cB.ItemID] += cB.Qty
	}
	return result
}

func removeCartItem(slice []CartItem, s int) []CartItem {
	return append(slice[:s], slice[s+1:]...)
}

func removeBonusItem(slice []BonusItem, s int) []BonusItem {
	return append(slice[:s], slice[s+1:]...)
}

func (pd *PromoDetail) Scan(val interface{}) error {
	switch v := val.(type) {
	case []byte:
		json.Unmarshal(v, &pd)
		return nil
	case string:
		json.Unmarshal([]byte(v), &pd)
		return nil
	default:
		return errors.New(fmt.Sprintf("unsupported type: %T", v))
	}
}
func (pd *PromoDetail) Value() (driver.Value, error) {
	return json.Marshal(pd)
}
