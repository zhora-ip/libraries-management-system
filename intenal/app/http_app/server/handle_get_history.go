package server

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/zhora-ip/libraries-management-system/intenal/models"
	svc "github.com/zhora-ip/libraries-management-system/intenal/models/service"
)

const (
	getHistoryLimit = 10
)

func (s *Server) HandleGetHistory() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var (
			userID = r.Context().Value(ctxKeyUserID{}).(int64)
			role   = r.Context().Value(ctxKeyUserRole{}).(int32)
			req    = &svc.FindAllOrdersRequest{
				Cursor:   time.Now(),
				Limit:    getHistoryLimit,
				Backward: false,
			}
			err error
		)

		err = json.NewDecoder(r.Body).Decode(req)
		switch {
		case err != nil:
			log.Print(err)
			s.respond(w, http.StatusBadRequest, nil)
			return
		case role == int32(models.UserRoleReader):
			req.UserID = &userID
		}

		resp, status, err := s.getHistoryHelper(r.Context(), req)
		if err != nil {
			s.error(w, status, nil)
			return
		}

		if resp == nil || len(resp.Data) == 0 {
			s.respond(w, http.StatusNoContent, &svc.FindAllBooksResponse{})
			return
		}

		s.respond(
			w,
			http.StatusOK,
			map[string]any{
				"orders":       resp.Data,
				"first_cursor": resp.Data[0].UpdatedAt.UTC(),
				"last_cursor":  resp.Data[len(resp.Data)-1].UpdatedAt.UTC(),
			},
		)

	}
}

func (s *Server) getHistoryHelper(ctx context.Context, req *svc.FindAllOrdersRequest) (*svc.FindAllOrdersResponse, int, error) {

	resp, err := s.oService.FindAll(ctx, req)
	if err != nil {
		log.Print(err)
		switch {
		case errors.Is(err, models.ErrObjectNotFound):
			return nil, http.StatusNoContent, models.ErrObjectNotFound
		}
		return nil, http.StatusInternalServerError, models.ErrInternal
	}

	return resp, http.StatusOK, nil
}
