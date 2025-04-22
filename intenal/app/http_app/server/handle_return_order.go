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

func (s *Server) HandleReturnOrder() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := &svc.ReturnOrderRequest{}

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			s.error(w, http.StatusBadRequest, nil)
			log.Print(err)
			return
		}

		resp, status, err := s.handleReturnOrderHelper(r.Context(), req)
		if resp != nil {
			s.respond(w, status, resp)
			return
		}

		s.error(w, status, err)

	}
}

func (s *Server) handleReturnOrderHelper(ctx context.Context, req *svc.ReturnOrderRequest) (*svc.ReturnOrderResponse, int, error) {

	resp, err := s.oService.Return(ctx, req)
	if err != nil {
		log.Print(err)
		switch {
		case errors.Is(err, models.ErrObjectNotFound):
			return nil, http.StatusNotFound, models.ErrObjectNotFound
		case errors.Is(err, models.ErrNoRows):
			return nil, http.StatusBadRequest, models.ErrNoRows
		case errors.Is(err, models.ErrIncorrectOrderStatus):
			return nil, http.StatusBadRequest, models.ErrIncorrectOrderStatus
		case errors.Is(err, models.ErrForbidden):
			return nil, http.StatusForbidden, models.ErrForbidden
		}
		return nil, http.StatusInternalServerError, nil
	}
	return resp, http.StatusOK, nil

}
