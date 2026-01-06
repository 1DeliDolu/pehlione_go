package sms

import (
	"context"
	"log/slog"
)

// MockProvider is a mock implementation of the SMSProvider interface.
type MockProvider struct {
	logger *slog.Logger
}

// NewMockProvider creates a new MockProvider.
func NewMockProvider(logger *slog.Logger) *MockProvider {
	return &MockProvider{logger: logger}
}

// Send logs the SMS message instead of sending it.
func (p *MockProvider) Send(ctx context.Context, toE164 string, body string, idempotencyKey string) (string, error) {
	p.logger.Info("Sending SMS",
		"to", toE164,
		"body", body,
		"idempotencyKey", idempotencyKey,
	)
	// In a real mock, you might want to simulate different outcomes
	// based on the input or some internal state.
	return "mock_message_id_" + idempotencyKey, nil
}
