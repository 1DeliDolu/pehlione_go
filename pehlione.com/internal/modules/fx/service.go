package fx

import (
	"context"
	"math"
	"strings"
	"time"
)

type Service struct {
	repo *Repo
	base string
}

func NewService(repo *Repo, base string) *Service {
	return &Service{
		repo: repo,
		base: strings.ToUpper(strings.TrimSpace(base)),
	}
}

// UpsertRates proxies to the underlying repo.
func (s *Service) UpsertRates(ctx context.Context, source string, fetchedAt time.Time, rates map[string]float64) error {
	if s.base != "" {
		rates[strings.ToUpper(s.base)] = 1.0
	}
	return s.repo.UpsertRates(ctx, source, fetchedAt, rates)
}

func (s *Service) BaseCurrency() string {
	if s.base == "" {
		return "TRY"
	}
	return s.base
}

func (s *Service) Rate(ctx context.Context, target string) (Rate, error) {
	target = strings.ToUpper(strings.TrimSpace(target))
	if target == "" || target == s.BaseCurrency() {
		return Rate{
			Currency:  s.BaseCurrency(),
			Rate:      1,
			Source:    "base",
			FetchedAt: time.Now(),
		}, nil
	}
	return s.repo.GetRate(ctx, target)
}

func (s *Service) ConvertFromBase(ctx context.Context, cents int, target string) (int, Rate, error) {
	rate, err := s.Rate(ctx, target)
	if err != nil {
		return 0, Rate{}, err
	}
	converted := roundCents(float64(cents) * rate.Rate)
	return converted, rate, nil
}

func roundCents(val float64) int {
	if val >= 0 {
		return int(math.Round(val))
	}
	return -int(math.Round(math.Abs(val)))
}
