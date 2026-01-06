package emails

import (
	"strings"

	"pehlione.com/app/internal/modules/orders"
	"pehlione.com/app/pkg/view"
)

// BuildOrderPayload prepares a common payload for transactional order emails.
func BuildOrderPayload(baseURL string, order orders.Order, items []orders.OrderItem, statusLabel string, reason string) map[string]any {
	data := map[string]any{
		"OrderID":     order.ID,
		"OrderURL":    strings.TrimRight(baseURL, "/") + "/orders/" + order.ID,
		"StatusLabel": statusLabel,
		"Total":       view.MoneyFromCents(order.TotalCents, order.Currency),
	}
	if reason != "" {
		data["Reason"] = reason
	}

	if len(items) > 0 {
		out := make([]map[string]any, 0, len(items))
		for _, it := range items {
			out = append(out, map[string]any{
				"Name":  it.ProductName,
				"Qty":   it.Quantity,
				"Price": view.MoneyFromCents(it.LineTotalCents, it.Currency),
			})
		}
		data["Items"] = out
	}
	return data
}
