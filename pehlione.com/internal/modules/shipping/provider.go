package shipping

import (
	"context"
)

// Address describes shipment destination details.
type Address struct {
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Address1   string `json:"address1"`
	Address2   string `json:"address2,omitempty"`
	City       string `json:"city"`
	PostalCode string `json:"postal_code"`
	Country    string `json:"country"`
	Phone      string `json:"phone"`
	Email      string `json:"email"`
}

// Item represents a single line in the shipment for provider integrations.
type Item struct {
	Name  string `json:"name"`
	SKU   string `json:"sku"`
	Qty   int    `json:"qty"`
	Price int    `json:"price_cents"`
}

type LabelRequest struct {
	ShipmentID string
	OrderID    string
	Carrier    string
	Service    string
	Reference  string
	Recipient  Address
	Items      []Item
	Note       string
}

type LabelResponse struct {
	Carrier        string
	TrackingNumber string
	TrackingURL    string
	LabelURL       string
	Meta           map[string]any
}

type Provider interface {
	Name() string
	CreateLabel(ctx context.Context, req LabelRequest) (LabelResponse, error)
}
