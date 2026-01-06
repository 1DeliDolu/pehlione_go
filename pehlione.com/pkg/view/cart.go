package view

type CartItem struct {
	ProductName string
	ProductSlug string
	ImageURL    string

	VariantID string
	Qty       int

	UnitPrice string
	LineTotal string

	UnitPriceCents     int
	LineTotalCents     int
	BaseUnitPriceCents int
	BaseLineTotalCents int
}

type CartPage struct {
	Items         []CartItem
	Currency      string
	BaseCurrency  string
	Count         int
	SubtotalCents int
	DisplaySubtotalCents int
	Subtotal      string
	TotalCents    int
	Total         string
	DisplayTotalCents int
	BaseSubtotalCents int
	BaseTotalCents    int
	CSRFToken     string
}
