package currency

import (
	"context"
	"strings"

	"pehlione.com/app/internal/modules/fx"
)

type Config struct {
	BaseCurrency      string
	DefaultDisplay    string
	DisplayCurrencies []string
	ChargeCurrencies  []string
}

type Service struct {
	fx             *fx.Service
	base           string
	defaultDisplay string
	displayAllowed map[string]struct{}
	displayOptions []string
	chargeAllowed  map[string]struct{}
}

func NewService(fxSvc *fx.Service, cfg Config) *Service {
	base := strings.ToUpper(strings.TrimSpace(cfg.BaseCurrency))
	if base == "" {
		base = fxSvc.BaseCurrency()
	}
	display := strings.ToUpper(strings.TrimSpace(cfg.DefaultDisplay))
	if display == "" {
		display = base
	}
	allowed := make(map[string]struct{})
	options := make([]string, 0, len(cfg.DisplayCurrencies))
	for _, c := range cfg.DisplayCurrencies {
		code := strings.ToUpper(strings.TrimSpace(c))
		if code == "" {
			continue
		}
		if _, ok := allowed[code]; ok {
			continue
		}
		allowed[code] = struct{}{}
		options = append(options, code)
	}
	if len(options) == 0 {
		allowed[base] = struct{}{}
		options = []string{base}
	}
	if _, ok := allowed[display]; !ok {
		display = options[0]
	}

	chargeAllowed := make(map[string]struct{})
	for _, c := range cfg.ChargeCurrencies {
		code := strings.ToUpper(strings.TrimSpace(c))
		if code == "" {
			continue
		}
		chargeAllowed[code] = struct{}{}
	}
	if len(chargeAllowed) == 0 {
		chargeAllowed[base] = struct{}{}
	}

	return &Service{
		fx:             fxSvc,
		base:           base,
		defaultDisplay: display,
		displayAllowed: allowed,
		displayOptions: options,
		chargeAllowed:  chargeAllowed,
	}
}

func (s *Service) BaseCurrency() string {
	if s.base != "" {
		return s.base
	}
	return s.fx.BaseCurrency()
}

func (s *Service) DefaultDisplayCurrency() string {
	if s.defaultDisplay != "" {
		return s.defaultDisplay
	}
	return s.BaseCurrency()
}

func (s *Service) DisplayOptions() []string {
	return append([]string{}, s.displayOptions...)
}

func (s *Service) NormalizeDisplay(code string) (string, bool) {
	c := strings.ToUpper(strings.TrimSpace(code))
	if c == "" {
		return s.DefaultDisplayCurrency(), true
	}
	if _, ok := s.displayAllowed[c]; ok {
		return c, true
	}
	return "", false
}

func (s *Service) ConvertDisplay(ctx context.Context, baseCents int, target string) (int, fx.Rate, error) {
	code, ok := s.NormalizeDisplay(target)
	if !ok {
		code = s.DefaultDisplayCurrency()
	}
	return s.fx.ConvertFromBase(ctx, baseCents, code)
}

func (s *Service) DisplayRate(ctx context.Context, target string) (fx.Rate, error) {
	code, ok := s.NormalizeDisplay(target)
	if !ok {
		code = s.DefaultDisplayCurrency()
	}
	return s.fx.Rate(ctx, code)
}

func (s *Service) ConvertCharge(ctx context.Context, baseCents int, chargeCurrency string) (int, fx.Rate, error) {
	code := strings.ToUpper(strings.TrimSpace(chargeCurrency))
	if code == "" {
		code = s.BaseCurrency()
	}
	return s.fx.ConvertFromBase(ctx, baseCents, code)
}

func (s *Service) ChooseChargeCurrency(display string) string {
	code := strings.ToUpper(strings.TrimSpace(display))
	if code == "" {
		code = s.DefaultDisplayCurrency()
	}
	if _, ok := s.chargeAllowed[code]; ok {
		return code
	}
	return s.BaseCurrency()
}

func (s *Service) CanCharge(code string) bool {
	_, ok := s.chargeAllowed[strings.ToUpper(strings.TrimSpace(code))]
	return ok
}
