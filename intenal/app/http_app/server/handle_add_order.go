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

func (s *Server) HandleAddOrder() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := &svc.AddOrderRequest{}

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			s.error(w, http.StatusBadRequest, nil)
			log.Print(err)
			return
		}

		resp, status, err := s.handleAddOrderHelper(r.Context(), req)
		if resp != nil {
			s.respond(w, status, resp)
			return
		}

		s.error(w, status, err)

	}
}

func (s *Server) handleAddOrderHelper(ctx context.Context, req *svc.AddOrderRequest) (*svc.AddOrderResponse, int, error) {

	resp, err := s.oService.Add(ctx, req)
	if err != nil {
		log.Print(err)
		switch {
		case errors.Is(err, models.ErrLibCardExpired):
			return nil, http.StatusBadRequest, models.ErrLibCardExpired
		case errors.Is(err, models.ErrObjectNotFound):
			return nil, http.StatusNotFound, models.ErrObjectNotFound
		case errors.Is(err, models.ErrAlreadyUnavailable):
			return nil, http.StatusBadRequest, models.ErrAlreadyUnavailable
		}
		return nil, http.StatusInternalServerError, nil
	}
	return resp, http.StatusOK, nil

}
