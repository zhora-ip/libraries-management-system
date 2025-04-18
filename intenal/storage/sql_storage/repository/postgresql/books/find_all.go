package books

import (
	"context"
	"errors"

	"github.com/jackc/pgx"
	"github.com/zhora-ip/libraries-management-system/intenal/models"
	svc "github.com/zhora-ip/libraries-management-system/intenal/models/service"
)

func (r *BooksRepo) FindAll(ctx context.Context, req *svc.FindAllRequest) (*svc.FindAllResponse, error) {
	resp := &svc.FindAllResponse{}
	err := r.db.Select(ctx, &resp.Data, `SELECT * FROM books;`)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, models.ErrObjectNotFound
		}
		return nil, err
	}

	return resp, nil
}
