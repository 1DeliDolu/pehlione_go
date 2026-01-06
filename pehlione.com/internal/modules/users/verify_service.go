package users

import (
	"bytes"
	"context"
	"crypto/rand"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"pehlione.com/app/internal/modules/auth"
	"pehlione.com/app/internal/modules/email"
)

type VerifyService struct {
	db         *gorm.DB
	emailSvc   *email.OutboxService
	appBaseURL string
}

func NewVerifyService(db *gorm.DB, emailSvc *email.OutboxService, appBaseURL string) *VerifyService {
	return &VerifyService{
		db:         db,
		emailSvc:   emailSvc,
		appBaseURL: appBaseURL,
	}
}

// StartEmailVerification creates a verification code and sends it to the user
func (s *VerifyService) StartEmailVerification(ctx context.Context, userID, userEmail string) error {
	code := new6DigitCode()
	hash := sha256.Sum256([]byte(code))
	now := time.Now()
	exp := now.Add(15 * time.Minute)

	log.Printf("VerifyService: starting email verification for user %s, email %s, code=%s", userID, userEmail, code)

	return s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// Mark old active verifications as used
		_ = tx.WithContext(ctx).Model(&EmailVerification{}).
			Where("user_id = ? AND used_at IS NULL", userID).
			Update("used_at", now).Error

		ev := EmailVerification{
			ID:         uuid.NewString(),
			UserID:     userID,
			CodeHash:   hash[:],
			ExpiresAt:  exp,
			UsedAt:     nil,
			Attempts:   0,
			LastSentAt: &now,
			CreatedAt:  now,
			UpdatedAt:  now,
		}
		if err := tx.WithContext(ctx).Create(&ev).Error; err != nil {
			return err
		}

		subject := "Pehlione - E-posta doğrulama kodunuz"
		text := fmt.Sprintf("Doğrulama kodunuz: %s\nBu kod 15 dakika geçerlidir.\n", code)
		html := fmt.Sprintf(`
<html>
  <body style="font-family: Arial, sans-serif;">
    <h2>E-posta Doğrulama</h2>
    <p>Doğrulama kodunuz: <b>%s</b></p>
    <p>Bu kod 15 dakika geçerlidir.</p>
    <p>Kodu doğrulama sayfasında girin.</p>
  </body>
</html>`, code)

		log.Printf("VerifyService: enqueueing verification email for %s", userEmail)
		if err := s.emailSvc.Enqueue(ctx, userEmail, subject, text, html); err != nil {
			log.Printf("VerifyService: failed to enqueue email: %v", err)
			return err
		}
		log.Printf("VerifyService: verification email enqueued successfully")
		return nil
	})
}

// ConfirmWithCode verifies the code and marks the email as verified
func (s *VerifyService) ConfirmWithCode(ctx context.Context, userID, code string) error {
	hash := sha256.Sum256([]byte(code))
	now := time.Now()

	return s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		var ev EmailVerification
		if err := tx.WithContext(ctx).
			Clauses(clause.Locking{Strength: "UPDATE"}).
			Where("user_id = ? AND used_at IS NULL", userID).
			Order("created_at DESC").
			First(&ev).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				return fmt.Errorf("no active verification found")
			}
			return err
		}

		if now.After(ev.ExpiresAt) {
			return fmt.Errorf("code expired")
		}

		// Increment attempts
		_ = tx.WithContext(ctx).Model(&EmailVerification{}).
			Where("id = ?", ev.ID).
			Update("attempts", gorm.Expr("attempts + 1")).Error

		// Compare hashes
		if len(ev.CodeHash) != 32 || !bytes.Equal(ev.CodeHash, hash[:]) {
			return fmt.Errorf("invalid code")
		}

		// Mark as used
		if err := tx.WithContext(ctx).Model(&EmailVerification{}).
			Where("id = ?", ev.ID).
			Update("used_at", now).Error; err != nil {
			return err
		}

		// Mark user as verified
		return tx.WithContext(ctx).Model(&auth.User{}).
			Where("id = ?", userID).
			Updates(map[string]any{
				"email_verified_at": &now,
				"status":            "active",
			}).Error
	})
}

func new6DigitCode() string {
	var b [8]byte
	_, _ = rand.Read(b[:])
	n := binary.LittleEndian.Uint64(b[:])
	return fmt.Sprintf("%06d", int(n%1_000_000))
}
