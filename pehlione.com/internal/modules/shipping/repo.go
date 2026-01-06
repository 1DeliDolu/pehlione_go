package shipping

import (
	"context"

	"gorm.io/gorm"
)

type Repo struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) *Repo { return &Repo{db: db} }

func (r *Repo) ListByOrder(ctx context.Context, orderID string) ([]Shipment, error) {
	var shipments []Shipment
	err := r.db.WithContext(ctx).
		Order("created_at ASC").
		Find(&shipments, "order_id = ?", orderID).Error
	return shipments, err
}
