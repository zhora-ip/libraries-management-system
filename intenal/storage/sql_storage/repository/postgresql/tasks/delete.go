package tasks

import (
	"context"

	"github.com/zhora-ip/libraries-management-system/intenal/models"
)

func (r *TasksRepo) Delete(ctx context.Context, ID int) error {
	res, err := r.db.Exec(ctx, "DELETE FROM tasks WHERE id=$1", ID)

	if res.RowsAffected() == 0 {
		return models.ErrObjectNotFound
	}

	return err
}
