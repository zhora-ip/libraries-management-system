package orders

import (
	"context"
	"errors"
	"fmt"
	"slices"

	"github.com/jackc/pgx"
	"github.com/zhora-ip/libraries-management-system/intenal/models"
	svc "github.com/zhora-ip/libraries-management-system/intenal/models/service"
	"github.com/zhora-ip/libraries-management-system/pkg"
)

func (r *OrdersRepo) FindAll(ctx context.Context, req *svc.FindAllOrdersRequest) ([]*models.Order, error) {

	var (
		orders                = []*models.Order{}
		baseQuery             = " SELECT * FROM orders WHERE"
		timeCondition         = " updated_at < $1"
		orderBy               = " ORDER BY updated_at DESC"
		limit                 = " LIMIT $2"
		i                     = 3
		args                  = []any{req.Cursor, req.Limit}
		idCondition           = pkg.AddCondition(req.ID, &args, " AND id = $%d", &i)
		occurrenceIdCondition = pkg.AddCondition(req.OccurrenceID, &args, " AND CAST(id AS TEXT) ILIKE '%%' || $%d || '%%'", &i)
		bookIdCondition       = pkg.AddCondition(req.BookID, &args, " AND physical_book_id = $%d", &i)
		userIdCondition       = pkg.AddCondition(req.UserID, &args, " AND user_id = $%d", &i)
		libraryIdCondition    = pkg.AddCondition(req.LibraryID, &args, " AND library_id = $%d", &i)
	)

	if req.Backward {
		timeCondition = " updated_at > $1"
		orderBy = " ORDER BY updated_at ASC"
	}

	query := fmt.Sprintf(
		`%s %s %s %s %s %s %s %s %s`,
		baseQuery,
		timeCondition,
		idCondition,
		occurrenceIdCondition,
		bookIdCondition,
		userIdCondition,
		libraryIdCondition,
		orderBy,
		limit,
	)

	err := r.db.Select(
		ctx,
		&orders,
		query,
		args...,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, models.ErrObjectNotFound
		}
		return nil, err
	}

	if len(orders) == 0 {
		return nil, models.ErrObjectNotFound
	}

	if req.Backward {
		slices.Reverse(orders)
	}

	return orders, nil

}
