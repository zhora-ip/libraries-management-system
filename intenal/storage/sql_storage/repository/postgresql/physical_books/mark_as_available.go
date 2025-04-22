package physbooks

import (
	"context"

	"github.com/zhora-ip/libraries-management-system/intenal/models"
)

func (r *PhysBooksRepo) MarkAsAvailable(ctx context.Context, ID int64) error {
	var (
		query = `UPDATE   
					physical_books
				SET
					is_available = true,
					updated_at = now()
				WHERE
					id = $1;`
	)

	tag, err := r.db.Exec(ctx, query, ID)
	if tag.RowsAffected() == 0 {
		return models.ErrAlreadyAvailable
	}
	return err
}
