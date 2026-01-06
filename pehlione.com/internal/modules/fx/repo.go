package fx

import (
	"context"
	"strings"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Repo struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) *Repo { return &Repo{db: db} }

// UpsertRates stores the latest rates for the provided currencies.
func (r *Repo) UpsertRates(ctx context.Context, source string, fetchedAt time.Time, rates map[string]float64) error {
	if len(rates) == 0 {
		return nil
	}
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		for currency, rate := range rates {
			currency = strings.ToUpper(strings.TrimSpace(currency))
			if currency == "" {
				continue
			}
			rec := Rate{
				Currency:  currency,
				Rate:      rate,
				Source:    source,
				FetchedAt: fetchedAt,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			}
			if err := tx.Clauses(clause.OnConflict{
				Columns:   []clause.Column{{Name: "currency"}},
				DoUpdates: clause.AssignmentColumns([]string{"rate", "source", "fetched_at", "updated_at"}),
			}).Create(&rec).Error; err != nil {
				return err
			}
		}
		return nil
	})
}

func (r *Repo) GetRate(ctx context.Context, currency string) (Rate, error) {
	var rate Rate
	err := r.db.WithContext(ctx).First(&rate, "currency = ?", strings.ToUpper(currency)).Error
	return rate, err
}
