package wishlist

import "time"

type Item struct {
	ID        string    `gorm:"type:char(36);primaryKey"`
	UserID    string    `gorm:"type:char(36);not null;index"`
	ProductID string    `gorm:"type:char(36);not null;index"`
	CreatedAt time.Time `gorm:"type:datetime(3);not null"`
}

func (Item) TableName() string { return "wishlist_items" }
