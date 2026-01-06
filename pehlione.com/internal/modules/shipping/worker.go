package shipping

import (
	"context"
	"log"
	"time"

	"github.com/google/uuid"
)

type Worker struct {
	svc       *Service
	workerID  string
	batchSize int
	interval  time.Duration
}

func NewWorker(svc *Service) *Worker {
	return &Worker{
		svc:       svc,
		workerID:  "ship-" + uuid.NewString(),
		batchSize: 5,
		interval:  2 * time.Second,
	}
}

func (w *Worker) Run(ctx context.Context) error {
	ticker := time.NewTicker(w.interval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-ticker.C:
			if err := w.tick(ctx); err != nil {
				log.Printf("shipping worker %s tick error: %v", w.workerID, err)
			}
		}
	}
}

func (w *Worker) tick(ctx context.Context) error {
	if w.svc == nil {
		return nil
	}
	jobs, err := w.svc.LockJobBatch(ctx, w.workerID, w.batchSize)
	if err != nil {
		return err
	}
	if len(jobs) == 0 {
		return nil
	}

	for _, job := range jobs {
		if err := w.svc.ProcessJob(ctx, job); err != nil {
			log.Printf("shipping worker %s: job %d failed: %v", w.workerID, job.ID, err)
		}
	}
	return nil
}
