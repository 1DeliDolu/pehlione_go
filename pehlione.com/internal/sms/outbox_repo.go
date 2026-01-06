package sms

import (
	"context"
	"database/sql"
	"encoding/json"
	"time"

	"gorm.io/gorm"
)

type OutboxMessage struct {
	ID                int64
	ToPhoneE164       string
	Template          string
	Payload           json.RawMessage
	Status            string
	AttemptCount      int
	LastError         sql.NullString
	ScheduledAt       time.Time
	LockedAt          sql.NullTime
	LockedBy          sql.NullString
	SentAt            sql.NullTime
	ProviderMessageID sql.NullString
	CreatedAt         time.Time
	UpdatedAt         time.Time
}

func (OutboxMessage) TableName() string {
	return "sms_outbox"
}

type OutboxRepository struct {
	db *gorm.DB
}

func NewOutboxRepository(db *gorm.DB) *OutboxRepository {
	return &OutboxRepository{db: db}
}

type EnqueueOptions struct {
	ToPhoneE164 string
	Template    string
	Payload     map[string]interface{}
}

func (r *OutboxRepository) Enqueue(ctx context.Context, opts EnqueueOptions) (int64, error) {
	payloadBytes, err := json.Marshal(opts.Payload)
	if err != nil {
		return 0, err
	}

	msg := OutboxMessage{
		ToPhoneE164: opts.ToPhoneE164,
		Template:    opts.Template,
		Payload:     payloadBytes,
		Status:      "pending",
		ScheduledAt: time.Now(),
	}

	if err := r.db.WithContext(ctx).Create(&msg).Error; err != nil {
		return 0, err
	}
		return msg.ID, nil
	}
	
	func (r *OutboxRepository) FetchAndLock(ctx context.Context, workerID string, limit int) ([]OutboxMessage, error) {
		var jobs []OutboxMessage
		err := r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
			// Find and lock pending jobs
			var ids []int64
			query := getFetchAndLockQuery(tx.Dialector.Name())
				if err := tx.Raw(query, limit).Pluck("id", &ids).Error; err != nil {
					return err
				}
			
				if len(ids) == 0 {
					return nil // No jobs to process
				}
			
				// For SQLite, we also need to explicitly update the status in the same transaction
				// to simulate the locking behavior for tests.
				if tx.Dialector.Name() == "sqlite" {
					now := time.Now()
					if err := tx.Model(&OutboxMessage{}).Where("id IN (?)", ids).Updates(map[string]interface{}{
						"status":    "processing",
						"locked_at": now,
						"locked_by": workerID,
					}).Error; err != nil {
						return err
					}
				}
				
			// Update locked jobs
			now := time.Now()
			if err := tx.Model(&OutboxMessage{}).Where("id IN (?)", ids).Updates(map[string]interface{}{
				"status":    "processing",
				"locked_at": now,
				"locked_by": workerID,
			}).Error; err != nil {
				return err
			}
	
			// Now fetch the full job details
			return tx.Where("id IN (?)", ids).Find(&jobs).Error
		})
	
			return jobs, err
		}

func getFetchAndLockQuery(dialect string) string {
	switch dialect {
	case "mysql":
		return `
			SELECT id FROM sms_outbox
			WHERE status = 'pending' AND scheduled_at <= NOW()
			ORDER BY scheduled_at, id
			LIMIT ?
			FOR UPDATE SKIP LOCKED
		`
	case "sqlite":
		// SQLite doesn't support FOR UPDATE SKIP LOCKED directly.
		// For testing, we'll just select pending jobs.
		// In a real SQLite setup, you'd need a different locking strategy or avoid this pattern.
		return `
			SELECT id FROM sms_outbox
			WHERE status = 'pending' AND scheduled_at <= datetime('now')
			ORDER BY scheduled_at, id
			LIMIT ?
		`
	default:
		// Fallback for other databases or when dialect is unknown
		return `
			SELECT id FROM sms_outbox
			WHERE status = 'pending' AND scheduled_at <= NOW()
			ORDER BY scheduled_at, id
			LIMIT ?
			FOR UPDATE
		`
	}
}
		
		func (r *OutboxRepository) MarkSent(ctx context.Context, id int64, providerMessageID string) error {
			return r.db.WithContext(ctx).Model(&OutboxMessage{}).Where("id = ?", id).Updates(map[string]interface{}{
				"status":              "sent",
				"sent_at":             time.Now(),
				"provider_message_id": providerMessageID,
				"locked_at":           nil,
				"locked_by":           nil,
			}).Error
		}
		
		func (r *OutboxRepository) MarkFailed(ctx context.Context, id int64, lastError string, attemptCount int, scheduledAt time.Time) error {
			status := "pending"
			if attemptCount >= 8 { // Hardcoded max attempts for now
				status = "failed"
			}
			return r.db.WithContext(ctx).Model(&OutboxMessage{}).Where("id = ?", id).Updates(map[string]interface{}{
				"status":      status,
				"attempt_count": attemptCount,
				"last_error":  lastError,
				"scheduled_at":  scheduledAt,
				"locked_at":   nil,
				"locked_by":   nil,
			}).Error
		}
		