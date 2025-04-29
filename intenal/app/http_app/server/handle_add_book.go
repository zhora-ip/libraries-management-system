package server

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/zhora-ip/libraries-management-system/intenal/models"
	svc "github.com/zhora-ip/libraries-management-system/intenal/models/service"
)

func (s *Server) HandleAddBook() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var (
			role = r.Context().Value(ctxKeyUserRole{}).(int32)
		)
		if role != int32(models.UserRoleLibrarian) && role != int32(models.UserRoleAdmin) {
			s.error(w, http.StatusForbidden, models.ErrInvalidRole)
			return
		}

		req := &svc.AddBookRequest{}

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			s.error(w, http.StatusBadRequest, nil)
			log.Print(err)
			return
		}

		resp, status, err := s.handleAddBookHelper(r.Context(), req)
		if resp != nil {
			s.respond(w, status, resp)
			return
		}

		s.error(w, status, err)

	}
}

func (s *Server) handleAddBookHelper(ctx context.Context, req *svc.AddBookRequest) (*svc.AddBookResponse, int, error) {

	resp, err := s.bService.Add(ctx, req)
	if err != nil {
		log.Print(err)
		switch {
		case errors.Is(err, models.ErrValidationFailed):
			return nil, http.StatusBadRequest, models.ErrValidationFailed
		}
		return nil, http.StatusInternalServerError, nil
	}
	return resp, http.StatusOK, nil

}
