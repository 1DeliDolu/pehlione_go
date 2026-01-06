package fx

import (
	"context"
	"log"
	"strings"
	"time"
)

type Worker struct {
	Service  *Service
	Provider Provider
	Base     string
	Symbols  []string
	Interval time.Duration
}

func NewWorker(svc *Service, provider Provider, base string, symbols []string, interval time.Duration) *Worker {
	if interval <= 0 {
		interval = time.Hour
	}
	return &Worker{
		Service:  svc,
		Provider: provider,
		Base:     strings.ToUpper(strings.TrimSpace(base)),
		Symbols:  symbols,
		Interval: interval,
	}
}

func (w *Worker) Run(ctx context.Context) error {
	ticker := time.NewTicker(w.Interval)
	defer ticker.Stop()

	if err := w.tick(ctx); err != nil {
		log.Printf("fx worker initial sync failed: %v", err)
	}

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-ticker.C:
			if err := w.tick(ctx); err != nil {
				log.Printf("fx worker tick error: %v", err)
			}
		}
	}
}

func (w *Worker) tick(ctx context.Context) error {
	if w.Provider == nil || w.Service == nil {
		return nil
	}
	base := w.Base
	if base == "" {
		base = w.Service.BaseCurrency()
	}
	rates, err := w.Provider.FetchRates(ctx, base, w.Symbols)
	if err != nil {
		return err
	}
	if len(rates) == 0 {
		return nil
	}
	return w.Service.UpsertRates(ctx, w.Provider.Name(), time.Now(), rates)
}
