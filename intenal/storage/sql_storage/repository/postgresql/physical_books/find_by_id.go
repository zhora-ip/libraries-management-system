package physbooks

import (
	"context"
	"strings"

	"github.com/zhora-ip/libraries-management-system/intenal/models"
)

func (r *PhysBooksRepo) FindByID(ctx context.Context, ID int64) (*models.PhysBook, error) {
	var (
		book  = &models.PhysBook{}
		query = `SELECT * 
				FROM 
					physical_books
				WHERE
					id = $1 AND is_available;`
	)

	err := r.db.Get(ctx, book, query, ID)

	if err != nil {
		if strings.Contains(err.Error(), "no rows in result set") {
			return nil, models.ErrObjectNotFound
		}
		return nil, err
	}

	return book, nil
}
