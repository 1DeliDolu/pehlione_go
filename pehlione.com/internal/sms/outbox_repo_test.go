package sms

import (
	"context"
	"database/sql"
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&OutboxMessage{})
	return db
}

func TestOutboxRepository_Enqueue(t *testing.T) {
	db := setupTestDB()
	defer func() {
		sqlDB, _ := db.DB()
		sqlDB.Close()
	}()

	repo := NewOutboxRepository(db)
	ctx := context.Background()

	t.Run("successfully enqueues an SMS message", func(t *testing.T) {
		payload := map[string]interface{}{
			"order_id": "123",
			"status":   "shipped",
		}
		opts := EnqueueOptions{
			ToPhoneE164: "+491701234567",
			Template:    "shipped",
			Payload:     payload,
		}

		id, err := repo.Enqueue(ctx, opts)
		require.NoError(t, err)
		assert.True(t, id > 0)

		var msg OutboxMessage
		require.NoError(t, db.First(&msg, id).Error)

		assert.Equal(t, opts.ToPhoneE164, msg.ToPhoneE164)
		assert.Equal(t, opts.Template, msg.Template)
		assert.Equal(t, "pending", msg.Status)
		assert.Equal(t, 0, msg.AttemptCount)
		assert.False(t, msg.CreatedAt.IsZero())
		assert.False(t, msg.UpdatedAt.IsZero())
		assert.False(t, msg.ScheduledAt.IsZero())

		var retrievedPayload map[string]interface{}
		require.NoError(t, json.Unmarshal(msg.Payload, &retrievedPayload))
		assert.Equal(t, payload["order_id"], retrievedPayload["order_id"])
		assert.Equal(t, payload["status"], retrievedPayload["status"])
	})

	t.Run("returns error on invalid payload", func(t *testing.T) {
		opts := EnqueueOptions{
			ToPhoneE164: "+491701234567",
			Template:    "invalid",
			Payload:     map[string]interface{}{
				"func": func() {}, // Unmarshalable type
			},
		}

		_, err := repo.Enqueue(ctx, opts)
		require.Error(t, err)
		assert.Contains(t, err.Error(), "unsupported type")
	})
}

func TestOutboxRepository_FetchAndLock(t *testing.T) {
	db := setupTestDB()
	defer func() {
		sqlDB, _ := db.DB()
		sqlDB.Close()
	}()

	repo := NewOutboxRepository(db)
	ctx := context.Background()

	// Prepare some messages
	msg1Payload, _ := json.Marshal(map[string]interface{}{"event": "first"})
	msg2Payload, _ := json.Marshal(map[string]interface{}{"event": "second"})
	msg3Payload, _ := json.Marshal(map[string]interface{}{"event": "third"})
	msg4Payload, _ := json.Marshal(map[string]interface{}{"event": "fourth"})

	db.Create(&OutboxMessage{ // ID 1
		ToPhoneE164: "+111", Template: "test", Payload: msg1Payload, Status: "pending", ScheduledAt: time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC),
	})
	db.Create(&OutboxMessage{ // ID 2
		ToPhoneE164: "+222", Template: "test", Payload: msg2Payload, Status: "pending", ScheduledAt: time.Date(2000, time.January, 1, 0, 30, 0, 0, time.UTC),
	})
	db.Create(&OutboxMessage{ // ID 3, Already processing
		ToPhoneE164: "+333", Template: "test", Payload: msg3Payload, Status: "processing", ScheduledAt: time.Date(2000, time.January, 1, 1, 0, 0, 0, time.UTC),
	})
	db.Create(&OutboxMessage{ // ID 4, Future scheduled
		ToPhoneE164: "+444", Template: "test", Payload: msg4Payload, Status: "pending", ScheduledAt: time.Date(3000, time.January, 1, 0, 0, 0, 0, time.UTC),
	})

	t.Run("fetches and locks pending jobs", func(t *testing.T) {
		workerID := "test-worker-1"
		lockedJobs, err := repo.FetchAndLock(ctx, workerID, 2) // Fetch up to 2 jobs
		require.NoError(t, err)
		assert.Len(t, lockedJobs, 2)

		// Verify jobs are locked and status changed
		for _, job := range lockedJobs {
			assert.Equal(t, "processing", job.Status)
			assert.Equal(t, workerID, job.LockedBy.String)
			assert.False(t, job.LockedAt.Time.IsZero())

			// Verify from DB directly
			var dbJob OutboxMessage
			require.NoError(t, db.First(&dbJob, job.ID).Error)
			assert.Equal(t, "processing", dbJob.Status)
			assert.Equal(t, workerID, dbJob.LockedBy.String)
			assert.False(t, dbJob.LockedAt.Time.IsZero())
		}
	})

	t.Run("does not fetch already locked or future scheduled jobs", func(t *testing.T) {
		time.Sleep(1 * time.Second) // Ensure datetime('now') is past the scheduled_at of the previous jobs
		// After the previous test, the first two pending jobs (ID 1, 2) are locked.
		// Now try to fetch again. Only ID 4 is still pending, but it's scheduled in the future.
		workerID := "test-worker-2"
		lockedJobs, err := repo.FetchAndLock(ctx, workerID, 2)
		require.NoError(t, err)
		assert.Len(t, lockedJobs, 0) // No new jobs should be fetched
	})
}

func TestOutboxRepository_MarkSent(t *testing.T) {
	db := setupTestDB()
	defer func() {
		sqlDB, _ := db.DB()
		sqlDB.Close()
	}()

	repo := NewOutboxRepository(db)
	ctx := context.Background()

	// Create a message that is "processing"
	msgPayload, _ := json.Marshal(map[string]interface{}{"event": "to_send"})
	msg := OutboxMessage{
		ToPhoneE164: "+111", Template: "test", Payload: msgPayload, Status: "processing",
		LockedAt: sql.NullTime{Time: time.Now().Add(-time.Minute), Valid: true},
		LockedBy: sql.NullString{String: "worker-1", Valid: true},
	}
	require.NoError(t, db.Create(&msg).Error)

	t.Run("marks message as sent", func(t *testing.T) {
		providerMessageID := "provider-msg-id-123"
		err := repo.MarkSent(ctx, msg.ID, providerMessageID)
		require.NoError(t, err)

		var updatedMsg OutboxMessage
		require.NoError(t, db.First(&updatedMsg, msg.ID).Error)

		assert.Equal(t, "sent", updatedMsg.Status)
		assert.False(t, updatedMsg.SentAt.Time.IsZero())
		assert.Equal(t, providerMessageID, updatedMsg.ProviderMessageID.String)
		assert.True(t, updatedMsg.ProviderMessageID.Valid)
		assert.False(t, updatedMsg.LockedAt.Valid)
		assert.False(t, updatedMsg.LockedBy.Valid)
	})
}

func TestOutboxRepository_MarkFailed(t *testing.T) {
	db := setupTestDB()
	defer func() {
		sqlDB, _ := db.DB()
		sqlDB.Close()
	}()

	repo := NewOutboxRepository(db)
	ctx := context.Background()

	// Create a message that is "processing"
	msgPayload, _ := json.Marshal(map[string]interface{}{"event": "to_fail"})
	msg := OutboxMessage{
		ToPhoneE164: "+111", Template: "test", Payload: msgPayload, Status: "processing", AttemptCount: 0,
		LockedAt: sql.NullTime{Time: time.Now().Add(-time.Minute), Valid: true},
		LockedBy: sql.NullString{String: "worker-1", Valid: true},
	}
	require.NoError(t, db.Create(&msg).Error)

	t.Run("marks message as pending with incremented attempt and scheduled_at for retry", func(t *testing.T) {
		lastError := "provider timeout"
		attemptCount := msg.AttemptCount + 1
		scheduledAt := time.Now().Add(time.Minute)

		err := repo.MarkFailed(ctx, msg.ID, lastError, attemptCount, scheduledAt)
		require.NoError(t, err)

		var updatedMsg OutboxMessage
		require.NoError(t, db.First(&updatedMsg, msg.ID).Error)

		assert.Equal(t, "pending", updatedMsg.Status)
		assert.Equal(t, attemptCount, updatedMsg.AttemptCount)
		assert.Equal(t, lastError, updatedMsg.LastError.String)
		assert.True(t, updatedMsg.LastError.Valid)
		assert.WithinDuration(t, scheduledAt, updatedMsg.ScheduledAt, time.Second)
		assert.False(t, updatedMsg.LockedAt.Valid)
		assert.False(t, updatedMsg.LockedBy.Valid)
	})

	t.Run("marks message as failed after max attempts", func(t *testing.T) {
		// Reset message to processing with max attempts - 1
		msg.Status = "processing"
		msg.AttemptCount = 7 // Max attempts is 8 (0-indexed 0 to 7)
		require.NoError(t, db.Save(&msg).Error)

		lastError := "permanent error"
		attemptCount := msg.AttemptCount + 1 // This will be 8
		scheduledAt := time.Now().Add(time.Hour)

		err := repo.MarkFailed(ctx, msg.ID, lastError, attemptCount, scheduledAt)
		require.NoError(t, err)

		var updatedMsg OutboxMessage
		require.NoError(t, db.First(&updatedMsg, msg.ID).Error)

		assert.Equal(t, "failed", updatedMsg.Status)
		assert.Equal(t, attemptCount, updatedMsg.AttemptCount)
		assert.Equal(t, lastError, updatedMsg.LastError.String)
		assert.True(t, updatedMsg.LastError.Valid)
		assert.WithinDuration(t, scheduledAt, updatedMsg.ScheduledAt, time.Second)
		assert.False(t, updatedMsg.LockedAt.Valid)
		assert.False(t, updatedMsg.LockedBy.Valid)
	})
}