package users

import "time"

type EmailVerification struct {
	ID         string     `gorm:"type:char(36);primaryKey"`
	UserID     string     `gorm:"type:char(36);index;not null"`
	CodeHash   []byte     `gorm:"type:varbinary(32);not null"`
	ExpiresAt  time.Time  `gorm:"type:datetime(3);not null"`
	UsedAt     *time.Time `gorm:"type:datetime(3)"`
	Attempts   int        `gorm:"not null"`
	LastSentAt *time.Time `gorm:"type:datetime(3)"`
	CreatedAt  time.Time  `gorm:"type:datetime(3);not null"`
	UpdatedAt  time.Time  `gorm:"type:datetime(3);not null"`
}

func (EmailVerification) TableName() string { return "email_verifications" }
