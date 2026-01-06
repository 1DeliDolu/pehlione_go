package sms

import (
	"context"
	"fmt"
	"log/slog"
	"time"
)

type Worker struct {
	provider    SMSProvider
	outboxRepo  *OutboxRepository
	logger      *slog.Logger
	workerID    string
	batchSize   int
	pollRate    time.Duration
}

func NewWorker(provider SMSProvider, outboxRepo *OutboxRepository, logger *slog.Logger, workerID string) *Worker {
	return &Worker{
		provider:    provider,
		outboxRepo:  outboxRepo,
		logger:      logger,
		workerID:    workerID,
		batchSize:   10,
		pollRate:    5 * time.Second,
	}
}

func (w *Worker) Run(ctx context.Context) {
	ticker := time.NewTicker(w.pollRate)
	defer ticker.Stop()

	w.logger.Info("SMS worker started", "worker_id", w.workerID)

	for {
		select {
		case <-ctx.Done():
			w.logger.Info("SMS worker shutting down")
			return
		case <-ticker.C:
			w.processJobs(ctx)
		}
	}
}

func (w *Worker) processJobs(ctx context.Context) {
	jobs, err := w.outboxRepo.FetchAndLock(ctx, w.workerID, w.batchSize)
	if err != nil {
		w.logger.Error("Failed to fetch jobs from outbox", "error", err)
		return
	}

	if len(jobs) > 0 {
		w.logger.Info(fmt.Sprintf("Processing %d jobs", len(jobs)))
	}

	for _, job := range jobs {
		// Placeholder for rendering and sending logic
		body := fmt.Sprintf("Message for %s with template %s", job.ToPhoneE164, job.Template)

		providerMessageID, err := w.provider.Send(ctx, job.ToPhoneE164, body, fmt.Sprintf("sms-%d", job.ID))
		if err != nil {
			w.handleFailedJob(ctx, job, err)
			continue
		}

		w.handleSuccessfulJob(ctx, job, providerMessageID)
	}
}

func (w *Worker) handleSuccessfulJob(ctx context.Context, job OutboxMessage, providerMessageID string) {

	if err := w.outboxRepo.MarkSent(ctx, job.ID, providerMessageID); err != nil {

		w.logger.Error("Failed to mark job as sent", "job_id", job.ID, "error", err)

	}

}



func (w *Worker) handleFailedJob(ctx context.Context, job OutboxMessage, jobErr error) {

	attemptCount := job.AttemptCount + 1

	backoffDuration := time.Duration(30*attemptCount) * time.Second // Example exponential backoff



	if err := w.outboxRepo.MarkFailed(ctx, job.ID, jobErr.Error(), attemptCount, time.Now().Add(backoffDuration)); err != nil {

		w.logger.Error("Failed to mark job as failed", "job_id", job.ID, "error", err)

	}

}
