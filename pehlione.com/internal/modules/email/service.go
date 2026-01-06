package email

import (
	"context"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type OutboxService struct {
	db *gorm.DB
}

func NewService(db *gorm.DB) *OutboxService {
	return &OutboxService{db: db}
}

// Enqueue adds an email to the outbox for async processing
func (s *OutboxService) Enqueue(ctx context.Context, to, subject, text, html string) error {
	now := time.Now()
	var tptr, hptr *string
	if text != "" {
		tptr = &text
	}
	if html != "" {
		hptr = &html
	}

	e := OutboxEmail{
		ID:            uuid.NewString(),
		ToEmail:       to,
		Subject:       subject,
		BodyText:      tptr,
		BodyHTML:      hptr,
		Status:        StatusPending,
		Attempts:      0,
		LastError:     nil,
		NextAttemptAt: now,
		CreatedAt:     now,
		SentAt:        nil,
	}
	return s.db.WithContext(ctx).Create(&e).Error
}

// GetPending returns pending emails ready to be sent
func (s *OutboxService) GetPending(ctx context.Context, limit int) ([]OutboxEmail, error) {
	var emails []OutboxEmail
	err := s.db.WithContext(ctx).
		Where("status = ? AND next_attempt_at <= ?", StatusPending, time.Now()).
		Order("created_at ASC").
		Limit(limit).
		Find(&emails).Error
	return emails, err
}

// UpdateStatus updates the status of an email
func (s *OutboxService) UpdateStatus(ctx context.Context, id, status string, sentAt *time.Time, lastError *string) error {
	return s.db.WithContext(ctx).Model(&OutboxEmail{}).
		Where("id = ?", id).
		Updates(map[string]any{
			"status":     status,
			"sent_at":    sentAt,
			"last_error": lastError,
		}).Error
}

// UpdateRetry updates retry info for a failed email
func (s *OutboxService) UpdateRetry(ctx context.Context, id string, attempts int, nextAttempt time.Time, lastError *string, status string) error {
	return s.db.WithContext(ctx).Model(&OutboxEmail{}).
		Where("id = ?", id).
		Updates(map[string]any{
			"attempts":        attempts,
			"last_error":      lastError,
			"next_attempt_at": nextAttempt,
			"status":          status,
		}).Error
}
