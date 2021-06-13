package model

// cache keys
const (
	KeyCart = "cart:%s"
)

// promo types
const (
	PromoTypeDiscount = 1
	PromoTypeBonus    = 2
)

// messages
const (
	// success
	CheckoutMessage = "Your checkout total is $%.2f"
	// error
	EmptyCartMessage      = OutputErrorIdentifier + "Your cart is empty"
	OutOfStockMessage     = OutputErrorIdentifier + "%s current qty is only %d"
	DefaultErrorMessage   = "An error has occurred"
	OutputErrorIdentifier = "ERROR_GQL_OUTPUT"
)

// misc
const (
	EnvKey = "ENV"
)
