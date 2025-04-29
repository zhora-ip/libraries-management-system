package books

import (
	"context"
	"strings"

	"github.com/zhora-ip/libraries-management-system/intenal/models"
	svc "github.com/zhora-ip/libraries-management-system/intenal/models/service"
)

func (r *BooksRepo) FindAll(ctx context.Context, req *svc.FindAllBooksRequest) (*svc.FindAllBooksResponse, error) {
	var (
		resp  = &svc.FindAllBooksResponse{}
		query = `SELECT * FROM books`
		args  = []any{}
	)

	if req.ID != nil {
		query += ` WHERE id = $1`
		args = append(args, req.ID)
	}

	err := r.db.Select(ctx, &resp.Data, query, args...)

	if err != nil {
		if strings.Contains(err.Error(), "no rows in result set") {
			return nil, models.ErrObjectNotFound
		}
		return nil, err
	}

	return resp, nil
}
