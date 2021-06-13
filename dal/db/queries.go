package db

const (
	getItemsByIDs = `
	SELECT * FROM items WHERE id = any($1)
	`
	getItemPromosByItemIDs = `
	SELECT * FROM item_promo WHERE item_id = any($1)	
	`
	getPromosByIDs = `
	SELECT * FROM promos WHERE id = any($1)
	`
	reduceItemStockByID = `
	UPDATE items SET 
		qty = qty - $2
	WHERE id = $1
	`
)
