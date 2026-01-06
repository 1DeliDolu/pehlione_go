package shipping

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
)

// MockProvider simulates carrier label creation without external calls.
type MockProvider struct {
	BaseURL string
}

func NewMockProvider(baseURL string) MockProvider {
	return MockProvider{BaseURL: strings.TrimRight(baseURL, "/")}
}

func (p MockProvider) Name() string { return "mockship" }

func (p MockProvider) CreateLabel(ctx context.Context, req LabelRequest) (LabelResponse, error) {
	_ = ctx
	tracking := strings.ToUpper(strings.ReplaceAll(uuid.NewString(), "-", ""))[:12]
	tracking = fmt.Sprintf("TRK-%s", tracking)

	base := p.BaseURL
	trackingURL := ""
	labelURL := ""
	if base != "" {
		trackingURL = fmt.Sprintf("%s/track/%s", base, tracking)
		labelURL = fmt.Sprintf("%s/labels/%s.pdf", base, req.ShipmentID)
	} else {
		trackingURL = fmt.Sprintf("https://mockship.local/track/%s", tracking)
		labelURL = fmt.Sprintf("https://mockship.local/labels/%s.pdf", req.ShipmentID)
	}

	meta := map[string]any{
		"estimated_delivery": time.Now().Add(72 * time.Hour).Format(time.RFC3339),
	}

	return LabelResponse{
		Carrier:        req.Carrier,
		TrackingNumber: tracking,
		TrackingURL:    trackingURL,
		LabelURL:       labelURL,
		Meta:           meta,
	}, nil
}
