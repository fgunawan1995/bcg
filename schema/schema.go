package schema

var Schema = `
    schema {
        query		: Query
        mutation	: Mutation
    }
    type Query {
        hello()					: String!
		cart(user_id:String!)	: Cart!
    }
    type Mutation {
		addToCart(add:AddItemToCart!)	: Cart!
		checkout(user_id:String!)		: String!
    }
	
	input AddItemToCart {
        user_id	: String!
        item_id	: String!
        qty		: Int!
    }
	type Cart {
		items		: [CartItem!]!
		bonus_items	: [BonusItem!]!
    }
	type CartItem {
		item		: Item!
		promo		: Promo
		cart_qty	: Int!
		sub_total	: Float!
		disc_total	: Float!
		total		: Float!
	}
	type BonusItem {
		item		: Item!
		bonus_qty	: Int!
	}
	type Item {
		id				: String!
		name			: String!
		sku				: String!
		price			: Float!
		current_stock	: Int!
	}

	type Promo {
		id		: String!
		name	: String!
	}
`
