package email

import (
	"context"
	"time"

	"gorm.io/datatypes"
)

type OutboxEmail struct {
	ID           int64          `gorm:"primaryKey;autoIncrement"`
	ToEmail      string         `gorm:"type:varchar(320);not null"`
	Template     string         `gorm:"type:varchar(128);not null"`
	Payload      datatypes.JSON `gorm:"type:json;not null"`
	Status       string         `gorm:"type:varchar(16);not null"`
	AttemptCount int            `gorm:"not null;default:0"`
	LastError    *string        `gorm:"type:text"`
	ScheduledAt  time.Time      `gorm:"type:datetime(3);not null"`
	LockedAt     *time.Time     `gorm:"type:datetime(3)"`
	LockedBy     *string        `gorm:"type:varchar(128)"`
	SentAt       *time.Time     `gorm:"type:datetime(3)"`
	CreatedAt    time.Time      `gorm:"type:datetime(3);not null"`
	UpdatedAt    time.Time      `gorm:"type:datetime(3);not null"`
}

func (OutboxEmail) TableName() string { return "email_outbox" }

const (
	StatusPending    = "pending"
	StatusProcessing = "processing"
	StatusSent       = "sent"
	StatusFailed     = "failed"
)

type Message struct {
	To      string
	Subject string
	Text    string
	HTML    string
}

type Sender interface {
	Send(ctx context.Context, m Message) error
}
