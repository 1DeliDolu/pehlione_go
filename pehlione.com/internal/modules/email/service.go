package email

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Job struct {
	To       string
	Template string
	Payload  map[string]any
}

type OutboxService struct {
	db *gorm.DB
}

func NewService(db *gorm.DB) *OutboxService {
	return &OutboxService{db: db}
}

func (s *OutboxService) Enqueue(ctx context.Context, job Job) error {
	return s.enqueue(ctx, s.db, job)
}

func (s *OutboxService) EnqueueTx(ctx context.Context, tx *gorm.DB, job Job) error {
	target := tx
	if target == nil {
		target = s.db
	}
	return s.enqueue(ctx, target, job)
}

func (s *OutboxService) enqueue(ctx context.Context, db *gorm.DB, job Job) error {
	to := strings.TrimSpace(job.To)
	if to == "" || job.Template == "" {
		return fmt.Errorf("invalid email job")
	}
	payload := job.Payload
	if payload == nil {
		payload = map[string]any{}
	}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	now := time.Now()
	e := OutboxEmail{
		ToEmail:      to,
		Template:     job.Template,
		Payload:      datatypes.JSON(data),
		Status:       StatusPending,
		AttemptCount: 0,
		LastError:    nil,
		ScheduledAt:  now,
		LockedAt:     nil,
		LockedBy:     nil,
		SentAt:       nil,
		CreatedAt:    now,
		UpdatedAt:    now,
	}
	return db.WithContext(ctx).Create(&e).Error
}

// LockBatch claims a batch of pending emails for processing using SKIP LOCKED semantics.
func (s *OutboxService) LockBatch(ctx context.Context, workerID string, limit int) ([]OutboxEmail, error) {
	if limit <= 0 {
		limit = 10
	}

	var locked []OutboxEmail
	err := s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.
			Clauses(clause.Locking{Strength: "UPDATE", Options: "SKIP LOCKED"}).
			Where("status = ? AND scheduled_at <= ?", StatusPending, time.Now()).
			Order("scheduled_at ASC").
			Limit(limit).
			Find(&locked).Error; err != nil {
			return err
		}

		if len(locked) == 0 {
			return nil
		}

		now := time.Now()
		ids := make([]int64, len(locked))
		for i, e := range locked {
			ids[i] = e.ID
		}

		if err := tx.Model(&OutboxEmail{}).
			Where("id IN ?", ids).
			Updates(map[string]any{
				"status":     StatusProcessing,
				"locked_at":  now,
				"locked_by":  workerID,
				"updated_at": now,
			}).Error; err != nil {
			return err
		}

		return tx.Where("id IN ?", ids).Find(&locked).Error
	})

	return locked, err
}

// MarkSent marks an email as sent successfully.
func (s *OutboxService) MarkSent(ctx context.Context, id int64) error {
	now := time.Now()
	return s.db.WithContext(ctx).Model(&OutboxEmail{}).
		Where("id = ?", id).
		Updates(map[string]any{
			"status":       StatusSent,
			"sent_at":      now,
			"last_error":   nil,
			"locked_at":    gorm.Expr("NULL"),
			"locked_by":    gorm.Expr("NULL"),
			"updated_at":   now,
			"scheduled_at": now,
		}).Error
}

// MarkRetry schedules an email for another attempt with exponential backoff.
func (s *OutboxService) MarkRetry(ctx context.Context, id int64, errMsg string, delay time.Duration) error {
	if delay <= 0 {
		delay = time.Minute
	}
	now := time.Now()
	next := now.Add(delay)

	return s.db.WithContext(ctx).Model(&OutboxEmail{}).
		Where("id = ?", id).
		Updates(map[string]any{
			"status":        StatusPending,
			"last_error":    errMsg,
			"scheduled_at":  next,
			"locked_at":     gorm.Expr("NULL"),
			"locked_by":     gorm.Expr("NULL"),
			"updated_at":    now,
			"attempt_count": gorm.Expr("attempt_count + 1"),
		}).Error
}

// MarkFailed marks an email as permanently failed.
func (s *OutboxService) MarkFailed(ctx context.Context, id int64, errMsg string) error {
	now := time.Now()
	return s.db.WithContext(ctx).Model(&OutboxEmail{}).
		Where("id = ?", id).
		Updates(map[string]any{
			"status":        StatusFailed,
			"last_error":    errMsg,
			"locked_at":     gorm.Expr("NULL"),
			"locked_by":     gorm.Expr("NULL"),
			"updated_at":    now,
			"attempt_count": gorm.Expr("attempt_count + 1"),
		}).Error
}
