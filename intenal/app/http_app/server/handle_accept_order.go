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

func (s *Server) HandleAcceptOrder() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			role = r.Context().Value(ctxKeyUserRole{}).(int32)
			req  = &svc.AcceptOrderRequest{}
		)

		if role != int32(models.UserRoleLibrarian) && role != int32(models.UserRoleAdmin) {
			s.error(w, http.StatusForbidden, models.ErrInvalidRole)
			return
		}

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			s.error(w, http.StatusBadRequest, nil)
			log.Print(err)
			return
		}

		resp, status, err := s.handleAcceptOrderHelper(r.Context(), req)
		if resp != nil {
			s.respond(w, status, resp)
			return
		}

		s.error(w, status, err)

	}
}

func (s *Server) handleAcceptOrderHelper(ctx context.Context, req *svc.AcceptOrderRequest) (*svc.AcceptOrderResponse, int, error) {

	resp, err := s.oService.Accept(ctx, req)
	if err != nil {
		log.Print(err)
		switch {
		case errors.Is(err, models.ErrObjectNotFound):
			return nil, http.StatusNotFound, models.ErrObjectNotFound
		case errors.Is(err, models.ErrNoRows):
			return nil, http.StatusBadRequest, models.ErrNoRows
		case errors.Is(err, models.ErrIncorrectOrderStatus):
			return nil, http.StatusBadRequest, models.ErrIncorrectOrderStatus
		}
		return nil, http.StatusInternalServerError, nil
	}
	return resp, http.StatusOK, nil

}
