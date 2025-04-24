package audit

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/zhora-ip/libraries-management-system/intenal/models"
)

func (p *AuditProducer) Produce(ctx context.Context) {

	ticker := time.NewTicker(time.Second * timeoutFindLogs)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			if err := p.processBatchOfTasks(ctx); err != nil {
				return
			}

		}
	}
}

func (p *AuditProducer) processBatchOfTasks(ctx context.Context) error {
	tasks, err := p.tasksRepo.FindTasks(ctx)
	if err != nil {
		if errors.Is(err, models.ErrObjectNotFound) {
			return nil
		}
		log.Print(err)
		return err
	}

	return p.iterateTasks(ctx, tasks)
}

func (p *AuditProducer) iterateTasks(ctx context.Context, tasks []*models.Task) error {
	for _, t := range tasks {
		p.processTask(ctx, *t)

		if ctx.Err() != nil {
			return ctx.Err()
		}
	}
	return nil
}

func (p *AuditProducer) processTask(ctx context.Context, task models.Task) {

	var (
		now = time.Now()
	)

	if task.Type != models.TaskTypeAuditLog {
		return
	}

	if task.Status == models.TaskStatusFailed && time.Since(*task.UpdatedAt) < timeoutFailed*time.Second {
		return
	}

	task.Status = models.TaskStatusProcessing
	if err := p.tasksRepo.Update(ctx, task); err != nil {
		log.Print(err)
		return
	}

	if err := p.producer.Produce(task.Payload); err != nil {
		task.Status = models.TaskStatusFailed
		task.AttemptCount++

		if task.AttemptCount >= maxAttempts {
			task.Status = models.TaskStatusNoAttemptsLeft
			task.FinishedAt = &now
		}

		if updateErr := p.tasksRepo.Update(ctx, task); updateErr != nil {
			log.Print(updateErr)
		}
		return
	}

	if err := p.tasksRepo.Delete(ctx, task.ID); err != nil {
		log.Print(err)
	}

}
