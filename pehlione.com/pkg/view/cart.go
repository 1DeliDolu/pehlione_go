package view

type CartItem struct {
	ProductName string
	ProductSlug string
	ImageURL    string

	VariantID string
	Qty       int

	UnitPrice string
	LineTotal string

	UnitPriceCents int
	LineTotalCents int
}

type CartPage struct {
	Items []CartItem

	Currency string
	Count    int

	SubtotalCents int
	Subtotal      string

	TotalCents int
	Total      string
}
