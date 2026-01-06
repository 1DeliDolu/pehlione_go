package sms

import (
	"context"
	"time"
)

type SMSProvider interface {
	Send(ctx context.Context, toE164 string, body string, idempotencyKey string) (providerMessageID string, err error)
}

type PhoneVerification struct {
	ID        uint `gorm:"primarykey"`
	UserID    string
	PhoneE164 string
	CodeHash  string
	ExpiresAt time.Time
	UsedAt    *time.Time
	CreatedAt time.Time
}

func (PhoneVerification) TableName() string {
	return "phone_verifications"
}

type Consent struct {
	ID        uint      `gorm:"primarykey"`
	UserID    string    `gorm:"type:char(36);not null"`
	PhoneE164 string    `gorm:"type:varchar(32);not null"`
	Action    string    `gorm:"type:varchar(16);not null"` // "opt_in" or "opt_out"
	Source    string    `gorm:"type:varchar(32);not null"` // "profile"
	IP        string    `gorm:"type:varchar(64)"`
	UserAgent string    `gorm:"type:varchar(255)"`
	CreatedAt time.Time `gorm:"not null"`
}

func (Consent) TableName() string {
	return "sms_consents"
}
