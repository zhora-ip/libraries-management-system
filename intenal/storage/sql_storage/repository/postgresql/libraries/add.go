package libraries

import (
	"context"

	"github.com/zhora-ip/libraries-management-system/intenal/models"
)

func (r *LibrariesRepo) Add(ctx context.Context, library *models.Library) (int64, error) {
	var ID int64
	err := r.db.ExecQueryRow(
		ctx,
		`INSERT INTO
			libraries(
				name,
				address,
				phone_number,
				lat,
				lng
			)
			VALUES($1,$2,$3,$4,$5)
			RETURNING id;`,
		library.Name,
		library.Address,
		library.PhoneNumber,
		library.Lat,
		library.Lng,
	).Scan(&ID)

	return ID, err
}
