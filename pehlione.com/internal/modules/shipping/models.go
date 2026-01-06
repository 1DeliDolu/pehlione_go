package shipping

import (
	"time"

	"gorm.io/datatypes"
)

// Shipment statuses.
const (
	StatusPending   = "pending"
	StatusQueued    = "queued"
	StatusShipped   = "shipped"
	StatusDelivered = "delivered"
	StatusFailed    = "failed"
)

// Job statuses for asynchronous label creation.
const (
	JobPending    = "pending"
	JobProcessing = "processing"
	JobSucceeded  = "succeeded"
	JobFailed     = "failed"
)

type Shipment struct {
	ID       string `gorm:"type:char(36);primaryKey"`
	OrderID  string `gorm:"type:char(36);not null;index:idx_shipments_order_created,priority:1"`
	Provider string `gorm:"type:varchar(32);not null"`
	Carrier  string `gorm:"type:varchar(64);not null"`
	Service  *string
	Note     *string

	TrackingNumber *string `gorm:"column:tracking_no"`
	TrackingURL    *string
	LabelURL       *string

	Status       string     `gorm:"type:varchar(32);not null"`
	ErrorMessage *string    `gorm:"type:text"`
	ShippedAt    *time.Time `gorm:"type:datetime(3)"`
	DeliveredAt  *time.Time `gorm:"type:datetime(3)"`

	CreatedAt time.Time `gorm:"type:datetime(3);not null"`
	UpdatedAt time.Time `gorm:"type:datetime(3);not null"`
}

func (Shipment) TableName() string { return "shipments" }

type ShipmentJob struct {
	ID           int64          `gorm:"primaryKey;autoIncrement"`
	ShipmentID   string         `gorm:"type:char(36);not null;uniqueIndex:ux_shipment_jobs_shipment"`
	Status       string         `gorm:"type:varchar(16);not null"`
	Payload      datatypes.JSON `gorm:"type:json;not null"`
	AttemptCount int            `gorm:"column:attempt_count;not null"`
	LastError    *string        `gorm:"type:text"`
	ScheduledAt  time.Time      `gorm:"type:datetime(3);not null"`
	LockedAt     *time.Time     `gorm:"type:datetime(3)"`
	LockedBy     *string        `gorm:"type:varchar(64)"`
	CreatedAt    time.Time      `gorm:"type:datetime(3);not null"`
	UpdatedAt    time.Time      `gorm:"type:datetime(3);not null"`
}

func (ShipmentJob) TableName() string { return "shipment_jobs" }
