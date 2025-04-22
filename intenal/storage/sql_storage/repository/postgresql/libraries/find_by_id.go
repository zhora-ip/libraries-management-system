package libraries

import (
	"context"
	"strings"

	"github.com/zhora-ip/libraries-management-system/intenal/models"
)

func (r *LibrariesRepo) FindByID(ctx context.Context, ID int64) (*models.Library, error) {
	var lib = &models.Library{}

	err := r.db.Get(ctx, lib, "SELECT * FROM libraries WHERE id = $1", ID)

	if err != nil {
		if strings.Contains(err.Error(), "no rows in result set") {
			return nil, models.ErrObjectNotFound
		}
		return nil, err
	}

	return lib, nil
}
