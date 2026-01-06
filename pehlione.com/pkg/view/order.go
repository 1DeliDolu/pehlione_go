package view

type OrderItem struct {
	ProductName string
	SKU         string
	Options     string
	Qty         int
	PriceEach   string
	LineTotal   string
}

type OrderDetail struct {
	ID        string
	Status    string
	Currency  string
	Subtotal  string
	Shipping  string
	Tax       string
	Discount  string
	Total     string
	Items     []OrderItem
	Shipments []OrderShipment
}

type OrderShipment struct {
	Carrier        string
	Status         string
	TrackingNumber string
	TrackingURL    string
}
