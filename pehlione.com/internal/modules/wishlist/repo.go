package wishlist

import (
	"context"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Repo struct{ db *gorm.DB }

func NewRepo(db *gorm.DB) *Repo { return &Repo{db: db} }

func (r *Repo) Add(ctx context.Context, userID, productID string) error {
	item := Item{
		ID:        uuid.NewString(),
		UserID:    userID,
		ProductID: productID,
		CreatedAt: time.Now(),
	}
	return r.db.WithContext(ctx).Create(&item).Error
}

func (r *Repo) Remove(ctx context.Context, userID, productID string) error {
	return r.db.WithContext(ctx).
		Where("user_id = ? AND product_id = ?", userID, productID).
		Delete(&Item{}).Error
}

func (r *Repo) ListByUser(ctx context.Context, userID string) ([]Item, error) {
	var items []Item
	err := r.db.WithContext(ctx).
		Where("user_id = ?", userID).
		Order("created_at DESC").
		Find(&items).Error
	return items, err
}

func (r *Repo) Contains(ctx context.Context, userID, productID string) (bool, error) {
	var count int64
	err := r.db.WithContext(ctx).
		Model(&Item{}).
		Where("user_id = ? AND product_id = ?", userID, productID).
		Count(&count).Error
	return count > 0, err
}
