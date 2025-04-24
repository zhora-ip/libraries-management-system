package orders

import (
	"context"
	"errors"
	"fmt"
	"log"
	"slices"

	"github.com/jackc/pgx"
	"github.com/zhora-ip/libraries-management-system/intenal/models"
	svc "github.com/zhora-ip/libraries-management-system/intenal/models/service"
)

func (r *OrdersRepo) FindAll(ctx context.Context, req *svc.FindAllOrdersRequest) ([]*models.Order, error) {

	var (
		orders                = []*models.Order{}
		baseQuery             = " SELECT * FROM orders WHERE"
		timeCondition         = " updated_at < $1"
		orderBy               = " ORDER BY updated_at DESC"
		limit                 = " LIMIT $2"
		i                     = 3
		idCondition           = ""
		occurrenceIdCondition = ""
		userIdCondition       = ""
		libraryIdCondition    = ""
		bookIdCondition       = ""
		args                  = []any{req.Cursor, req.Limit}
	)

	if req.Backward {
		timeCondition = " updated_at > $1"
		orderBy = " ORDER BY updated_at ASC"
	}

	idCondition = addCondition(req.ID, &args, " AND id = $%d", &i)
	occurrenceIdCondition = addCondition(req.OccurrenceID, &args, " AND CAST(id AS TEXT) ILIKE '%%' || $%d || '%%'", &i)
	bookIdCondition = addCondition(req.BookID, &args, " AND physical_book_id = $%d", &i)
	userIdCondition = addCondition(req.UserID, &args, " AND user_id = $%d", &i)
	libraryIdCondition = addCondition(req.LibraryID, &args, " AND library_id = $%d", &i)

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

	log.Print(query)

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

func addCondition(v *int64, args *[]any, str string, i *int) string {
	if v != nil {
		cond := fmt.Sprintf(str, *i)
		*i++
		*args = append(*args, fmt.Sprint(*v))
		return cond
	}
	return ""
}
