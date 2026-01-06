package users

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"strings"
	"time"

	"gorm.io/gorm"

	"pehlione.com/app/internal/modules/email"
)

type EmailVerification struct {
	ID        int64      `gorm:"primaryKey"`
	UserID    string     `gorm:"column:user_id"`
	TokenHash string     `gorm:"column:token_hash"`
	ExpiresAt time.Time  `gorm:"column:expires_at"`
	UsedAt    *time.Time `gorm:"column:used_at"`
	CreatedAt time.Time  `gorm:"column:created_at"`
}

func (EmailVerification) TableName() string { return "email_verifications" }

type VerifyService struct {
	db         *gorm.DB
	emailSvc   *email.OutboxService
	appBaseURL string
	fromName   string
}

func NewVerifyService(db *gorm.DB, emailSvc *email.OutboxService, appBaseURL, fromName string) *VerifyService {
	return &VerifyService{
		db:         db,
		emailSvc:   emailSvc,
		appBaseURL: appBaseURL,
		fromName:   fromName,
	}
}

func (s *VerifyService) StartEmailVerification(ctx context.Context, userID, userEmail string) error {
	if s.emailSvc == nil {
		return nil
	}

	return s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		rawToken, err := randomToken(32)
		if err != nil {
			return err
		}
		hash := sha256.Sum256([]byte(rawToken))
		hashHex := hex.EncodeToString(hash[:])

		_ = tx.WithContext(ctx).Where("user_id = ?", userID).Delete(&EmailVerification{}).Error

		ev := EmailVerification{
			UserID:    userID,
			TokenHash: hashHex,
			ExpiresAt: time.Now().Add(30 * time.Minute),
			CreatedAt: time.Now(),
		}
		if err := tx.WithContext(ctx).Create(&ev).Error; err != nil {
			return err
		}

		verifyURL := strings.TrimRight(s.appBaseURL, "/") + "/verify-email?token=" + rawToken
		return s.emailSvc.EnqueueTx(ctx, tx, email.Job{
			To:       userEmail,
			Template: "verify_email",
			Payload: map[string]any{
				"verify_url": verifyURL,
				"from_name":  s.fromName,
			},
		})
	})
}

func randomToken(n int) (string, error) {
	b := make([]byte, n)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return hex.EncodeToString(b), nil
}
