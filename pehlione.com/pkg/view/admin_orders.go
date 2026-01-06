package view

type AdminOrderListItem struct {
	ID         string
	Status     string
	Total      string
	CreatedAt  string
	UserID     string
	GuestEmail string
}

type AdminOrdersListPage struct {
	Items      []AdminOrderListItem
	Q          string
	Status     string
	Page       int
	TotalPages int
}

type AdminOrderItem struct {
	ProductName string
	SKU         string
	Options     string
	Qty         int
	Unit        string
	Line        string
}

type AdminOrderEvent struct {
	Action      string
	From        string
	To          string
	ActorUserID string
	Note        string
	At          string
}

type AdminOrderDetail struct {
	ID         string
	Status     string
	Currency   string
	UserID     string
	GuestEmail string
	CreatedAt  string

	Subtotal string
	Shipping string
	Tax      string
	Discount string
	Total    string

	Items             []AdminOrderItem
	Events            []AdminOrderEvent
	Shipments         []AdminShipment
	Financial         []AdminOrderFinancialEntry
	ShippingAvailable bool
}

type AdminOrderFinancialEntry struct {
	Event       string
	AmountCents int
	AmountStr   string
	Currency    string
	RefType     string
	RefID       string
	At          string
}

type AdminShipment struct {
	ID             string
	Carrier        string
	Status         string
	TrackingNumber string
	TrackingURL    string
	LabelURL       string
	Note           string
	ShippedAt      string
	DeliveredAt    string
	Error          string
}
