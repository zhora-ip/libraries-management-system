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

func (s *Server) HandleIssueOrder() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			userID = r.Context().Value(ctxKeyUserID{}).(int64)
			req    = &svc.IssueOrderRequest{}
		)

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			s.error(w, http.StatusBadRequest, nil)
			log.Print(err)
			return
		}
		req.UserID = userID

		resp, status, err := s.handleIssueOrderHelper(r.Context(), req)
		if resp != nil {
			s.respond(w, status, resp)
			return
		}

		s.error(w, status, err)

	}
}

func (s *Server) handleIssueOrderHelper(ctx context.Context, req *svc.IssueOrderRequest) (*svc.IssueOrderResponse, int, error) {

	resp, err := s.oService.Issue(ctx, req)
	if err != nil {
		log.Print(err)
		switch {
		case errors.Is(err, models.ErrObjectNotFound):
			return nil, http.StatusNotFound, models.ErrObjectNotFound
		case errors.Is(err, models.ErrNoRows):
			return nil, http.StatusBadRequest, models.ErrNoRows
		case errors.Is(err, models.ErrIncorrectOrderStatus):
			return nil, http.StatusBadRequest, models.ErrIncorrectOrderStatus
		case errors.Is(err, models.ErrAlreadyExpired):
			return nil, http.StatusBadRequest, models.ErrAlreadyExpired
		case errors.Is(err, models.ErrForbidden):
			return nil, http.StatusForbidden, models.ErrForbidden
		}
		return nil, http.StatusInternalServerError, nil
	}
	return resp, http.StatusOK, nil

}
