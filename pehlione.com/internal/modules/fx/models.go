package fx

import "time"

// Rate represents a cached FX rate relative to the base currency.
// Example: if base=TRY and currency=USD, rate=0.030 means 1 TRY = 0.03 USD.
type Rate struct {
	Currency  string    `gorm:"type:char(3);primaryKey"`
	Rate      float64   `gorm:"type:decimal(18,8);not null"`
	Source    string    `gorm:"type:varchar(32);not null"`
	FetchedAt time.Time `gorm:"type:datetime(3);not null"`
	CreatedAt time.Time `gorm:"type:datetime(3);not null"`
	UpdatedAt time.Time `gorm:"type:datetime(3);not null"`
}

func (Rate) TableName() string { return "fx_rates" }
