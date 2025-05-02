package books

import (
	"context"
	"errors"
	"fmt"
	"log"
	"slices"

	"github.com/jackc/pgx"
	"github.com/zhora-ip/libraries-management-system/intenal/models"
	svc "github.com/zhora-ip/libraries-management-system/intenal/models/service"
	"github.com/zhora-ip/libraries-management-system/pkg"
)

func (r *BooksRepo) FindAll(ctx context.Context, req *svc.FindAllBooksRequest) ([]*models.Book, error) {
	var (
		books           = []*models.Book{}
		baseQuery       = `SELECT * FROM books WHERE`
		timeCondition   = " updated_at < $1"
		orderBy         = " ORDER BY updated_at DESC"
		limit           = " LIMIT $2"
		i               = 3
		args            = []any{req.Cursor, req.Limit}
		idCondition     = pkg.AddCondition(req.ID, &args, " AND id = $%d", &i)
		titleCondition  string
		authorCondition string
		genreCondition  string
		ageCondition    string
	)

	if req.Backward {
		timeCondition = " updated_at > $1"
		orderBy = " ORDER BY updated_at ASC"
	}

	if req.Title != nil {
		titleCondition = fmt.Sprintf(" AND title = $%d", i)
		args = append(args, *req.Title)
		i++
	}

	if req.Author != nil {
		authorCondition = fmt.Sprintf(" AND author = $%d", i)
		args = append(args, *req.Author)
		i++
	}

	if req.Genre != nil {
		genreCondition = fmt.Sprintf(" AND genre = $%d", i)
		args = append(args, *req.Genre)
		i++
	}

	if req.AgeLimit != nil {
		ageCondition = fmt.Sprintf(" AND age_limit >= $%d", i)
		args = append(args, *req.AgeLimit)
		i++
	}

	query := fmt.Sprintf(
		`%s %s %s %s %s %s %s %s %s`,
		baseQuery,
		timeCondition,
		idCondition,
		titleCondition,
		authorCondition,
		genreCondition,
		ageCondition,
		orderBy,
		limit,
	)

	log.Print(query)
	err := r.db.Select(ctx, &books, query, args...)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, models.ErrObjectNotFound
		}
		return nil, err
	}

	if len(books) == 0 {
		return nil, models.ErrObjectNotFound
	}

	if req.Backward {
		slices.Reverse(books)
	}
	return books, nil
}
