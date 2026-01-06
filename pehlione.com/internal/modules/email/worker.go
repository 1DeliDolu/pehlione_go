package email

import (
	"context"
	"errors"
	"log"
	"time"

	"gorm.io/gorm"
)

type Worker struct {
	db     *gorm.DB
	sender Sender
}

func NewWorker(db *gorm.DB, sender Sender) *Worker {
	return &Worker{db: db, sender: sender}
}

// Run starts the email worker loop
func (w *Worker) Run(ctx context.Context) error {
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-ticker.C:
			_ = w.tick(ctx)
		}
	}
}

func (w *Worker) tick(ctx context.Context) error {
	svc := NewService(w.db)

	// Get batch of pending emails
	batch, err := svc.GetPending(ctx, 10)
	if err != nil {
		log.Printf("Email worker GetPending error: %v", err)
		return err
	}
	if len(batch) == 0 {
		return nil
	}

	log.Printf("Email worker: processing %d pending emails", len(batch))

	for _, e := range batch {
		if err := w.sendOne(ctx, svc, e.ID); err != nil {
			log.Printf("Email worker: failed to send %s: %v", e.ID, err)
		}
	}
	return nil
}

func (w *Worker) sendOne(ctx context.Context, svc *OutboxService, id string) error {
	var e OutboxEmail
	if err := w.db.WithContext(ctx).First(&e, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		}
		return err
	}

	if e.Status != StatusPending {
		return nil
	}

	text := ""
	html := ""
	if e.BodyText != nil {
		text = *e.BodyText
	}
	if e.BodyHTML != nil {
		html = *e.BodyHTML
	}

	log.Printf("Email worker: sending email %s to %s (attempt %d)", id, e.ToEmail, e.Attempts+1)

	err := w.sender.Send(ctx, Message{
		To:      e.ToEmail,
		Subject: e.Subject,
		Text:    text,
		HTML:    html,
	})

	now := time.Now()
	if err == nil {
		log.Printf("Email worker: successfully sent %s to %s", id, e.ToEmail)
		return svc.UpdateStatus(ctx, id, StatusSent, &now, nil)
	}

	log.Printf("Email worker: failed to send %s to %s: %v", id, e.ToEmail, err)

	// Retry logic
	attempts := e.Attempts + 1
	nextAttempt := now.Add(backoff(attempts))
	errMsg := truncate(err.Error(), 250)
	status := StatusPending

	if attempts >= 8 {
		status = StatusFailed
		log.Printf("Email %s failed after %d attempts: %v", id, attempts, err)
	}

	return svc.UpdateRetry(ctx, id, attempts, nextAttempt, &errMsg, status)
}

func backoff(attempt int) time.Duration {
	switch attempt {
	case 1:
		return 1 * time.Minute
	case 2:
		return 5 * time.Minute
	case 3:
		return 15 * time.Minute
	default:
		return 30 * time.Minute
	}
}

func truncate(s string, n int) string {
	if n <= 0 || len(s) <= n {
		return s
	}
	return s[:n]
}
