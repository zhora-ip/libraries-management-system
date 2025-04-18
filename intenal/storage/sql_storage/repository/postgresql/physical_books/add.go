package physbooks

import (
	"context"

	"github.com/zhora-ip/libraries-management-system/intenal/models"
)

func (r *PhysBooksRepo) Add(ctx context.Context, book *models.PhysBook) (int64, error) {
	var ID int64
	err := r.db.ExecQueryRow(
		ctx,
		`INSERT INTO
			physical_books(
				library_id,
				book_id,
				is_available
			)
			VALUES($1,$2)
			RETURNING id;`,
		book.LibraryID,
		book.BookID,
		book.IsAvailable,
	).Scan(&ID)

	return ID, err
}
