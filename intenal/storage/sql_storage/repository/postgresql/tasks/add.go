package tasks

import (
	"context"
	"time"

	"github.com/zhora-ip/libraries-management-system/intenal/models"
)

func (r *TasksRepo) Add(ctx context.Context, payload string) (int, error) {
	var ID int

	err := r.db.ExecQueryRow(
		ctx,
		`INSERT INTO
			tasks (
				created_at,
				updated_at,
				status,
				type,
				payload
			)
			VALUES ($1,$2,$3,$4,$5)
			RETURNING id;`,
		time.Now(),
		time.Now(),
		models.TaskStatusCreated,
		models.TaskTypeAuditLog,
		[]byte(payload),
	).Scan(&ID)

	return ID, err
}
