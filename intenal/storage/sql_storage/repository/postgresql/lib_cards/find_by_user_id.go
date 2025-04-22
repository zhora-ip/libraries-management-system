package libcards

import (
	"context"
	"strings"

	"github.com/zhora-ip/libraries-management-system/intenal/models"
)

func (r *LibCardsRepo) FindByUserID(ctx context.Context, userID int64) (*models.LibCard, error) {
	var (
		libCard = &models.LibCard{}
		query   = `SELECT * 
				FROM 
					lib_cards
				WHERE
					user_id = $1;`
	)

	err := r.db.Get(ctx, libCard, query, userID)

	if err != nil {
		if strings.Contains(err.Error(), "no rows in result set") {
			return nil, models.ErrObjectNotFound
		}
		return nil, err
	}

	return libCard, nil
}
