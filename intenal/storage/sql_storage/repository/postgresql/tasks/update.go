package tasks

import (
	"context"
	"time"

	"github.com/zhora-ip/libraries-management-system/intenal/models"
)

func (r *TasksRepo) Update(ctx context.Context, task models.Task) error {

	_, err := r.db.Exec(
		ctx,
		`UPDATE tasks
		 SET
			status = $1,
			updated_at = $2,
			finished_at = $3,
			attempt_count = $4
		 WHERE id = $5`,
		task.Status,
		time.Now(),
		task.FinishedAt,
		task.AttemptCount,
		task.ID,
	)

	return err

}
