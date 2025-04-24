package tasks

import (
	"context"
	"errors"

	"github.com/jackc/pgx"
	"github.com/zhora-ip/libraries-management-system/intenal/models"
)

func (r *TasksRepo) FindTasks(ctx context.Context) ([]*models.Task, error) {
	var (
		tasks []*models.Task
	)

	err := r.db.Select(
		ctx,
		&tasks,
		`SELECT * FROM tasks
		WHERE
			status = $1 OR
			status = $2
		ORDER BY
			updated_at ASC;`,
		models.TaskStatusCreated,
		models.TaskStatusFailed,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, models.ErrObjectNotFound
		}
		return nil, err
	}

	return tasks, nil
}
