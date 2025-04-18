package books

import (
	"context"

	"github.com/zhora-ip/libraries-management-system/intenal/models"
)

func (r *BooksRepo) Add(ctx context.Context, book *models.Book) (int64, error) {
	var ID int64
	err := r.db.ExecQueryRow(
		ctx,
		`INSERT INTO
			books(
				title,
				author,
				genre,
				description,
				age_limit
			)
			VALUES($1,$2,$3,$4,$5)
			RETURNING id;`,
		book.Title,
		book.Author,
		book.Genre,
		book.Description,
		book.AgeLimit,
	).Scan(&ID)

	return ID, err
}
