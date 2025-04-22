package orders

import (
	"context"
	"strings"

	"github.com/zhora-ip/libraries-management-system/intenal/models"
)

func (r *OrdersRepo) FindByID(ctx context.Context, ID int64) (*models.Order, error) {
	var (
		order = &models.Order{}
		query = `SELECT * FROM orders WHERE id = $1`
	)
	err := r.db.Get(ctx, order, query, ID)

	if err != nil {
		if strings.Contains(err.Error(), "no rows in result set") {
			return nil, models.ErrObjectNotFound
		}
		return nil, err
	}

	return order, nil
}
