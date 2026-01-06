package reviews

import (
	"time"
)

type Review struct {
	ID        string     `gorm:"type:char(36);primaryKey"`
	ProductID string     `gorm:"type:char(36);not null;index"`
	UserID    string     `gorm:"type:char(36);not null;index"`
	OrderID   string     `gorm:"type:char(36);not null"`
	Rating    int        `gorm:"type:int;not null"`
	Body      string     `gorm:"type:text"`
	Status    string     `gorm:"type:varchar(16);not null;default:pending"`
	CreatedAt time.Time  `gorm:"type:datetime(3);not null"`
	UpdatedAt time.Time  `gorm:"type:datetime(3);not null"`
	DeletedAt *time.Time `gorm:"type:datetime(3)"`
}

func (Review) TableName() string { return "product_reviews" }
