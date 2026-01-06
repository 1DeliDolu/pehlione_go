package email

import (
	"context"
	"time"
)

type OutboxEmail struct {
	ID            string     `gorm:"type:char(36);primaryKey"`
	ToEmail       string     `gorm:"type:varchar(255);not null"`
	Subject       string     `gorm:"type:varchar(255);not null"`
	BodyText      *string    `gorm:"type:text"`
	BodyHTML      *string    `gorm:"type:mediumtext"`
	Status        string     `gorm:"type:varchar(16);not null"`
	Attempts      int        `gorm:"not null"`
	LastError     *string    `gorm:"type:varchar(255)"`
	NextAttemptAt time.Time  `gorm:"type:datetime(3);not null"`
	CreatedAt     time.Time  `gorm:"type:datetime(3);not null"`
	SentAt        *time.Time `gorm:"type:datetime(3)"`
}

func (OutboxEmail) TableName() string { return "outbox_emails" }

const (
	StatusPending = "pending"
	StatusSent    = "sent"
	StatusFailed  = "failed"
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
