package fx

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type Provider interface {
	Name() string
	FetchRates(ctx context.Context, base string, symbols []string) (map[string]float64, error)
}

type ExchangeRateHostProvider struct {
	client  *http.Client
	Symbols []string
}

func NewExchangeRateHostProvider(symbols []string) *ExchangeRateHostProvider {
	return &ExchangeRateHostProvider{
		client:  &http.Client{Timeout: 10 * time.Second},
		Symbols: symbols,
	}
}

func (p *ExchangeRateHostProvider) Name() string { return "exchange_rate_host" }

func (p *ExchangeRateHostProvider) FetchRates(ctx context.Context, base string, symbols []string) (map[string]float64, error) {
	if len(symbols) == 0 {
		symbols = p.Symbols
	}
	params := url.Values{}
	params.Set("base", strings.ToUpper(base))
	if len(symbols) > 0 {
		params.Set("symbols", strings.Join(symbols, ","))
	}
	reqURL := "https://api.exchangerate.host/latest?" + params.Encode()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := p.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("fx provider error: %s", resp.Status)
	}

	var data struct {
		Success bool               `json:"success"`
		Rates   map[string]float64 `json:"rates"`
		Base    string             `json:"base"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	if len(data.Rates) == 0 {
		return nil, fmt.Errorf("fx provider returned no rates")
	}

	normalized := make(map[string]float64, len(data.Rates))
	for k, v := range data.Rates {
		normalized[strings.ToUpper(k)] = v
	}
	return normalized, nil
}
