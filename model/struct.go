package model

type Item struct {
	ID    string  `db:"id"`
	SKU   string  `db:"sku"`
	Name  string  `db:"name"`
	Price float64 `db:"price"`
	Qty   int32   `db:"qty"`
}

type Promo struct {
	ID      string      `db:"id"`
	Name    string      `db:"name"`
	Type    int32       `db:"type"`
	Details PromoDetail `db:"details"`
}

type PromoDetail struct {
	BonusItemID        string  `json:"bonus_item_id"`
	BonusQty           int32   `json:"bonus_qty"`
	PerQty             int32   `json:"per_qty"`
	MinQty             int32   `json:"min_qty"`
	DiscountPercentage float64 `json:"discount_percentage"`
}

type ItemPromo struct {
	ID      string `db:"id"`
	ItemID  string `db:"item_id"`
	PromoID string `db:"promo_id"`
}

type Cart struct {
	Items []CartItem  `json:"items"`
	Bonus []BonusItem `json:"bonus"`
}

type CartItem struct {
	ItemID string
	Qty    int32
}

type BonusItem struct {
	ItemID string
	Qty    int32
}

type AddItemToCart struct {
	UserID string
	ItemID string
	Qty    int32
}
