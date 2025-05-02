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

func (s *Server) HandleUpdateUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			req    = &svc.UpdateUserRequest{}
			userID = r.Context().Value(ctxKeyUserID{}).(int64)
		)

		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, http.StatusBadRequest, nil)
			return
		}
		req.ID = userID

		_, status, err := s.handleUpdateUserHelper(r.Context(), req)
		if err != nil {
			s.error(w, status, err)
			return
		}

		s.respond(w, http.StatusOK, nil)
	}
}

func (s *Server) handleUpdateUserHelper(ctx context.Context, req *svc.UpdateUserRequest) (*svc.UpdateUserResponse, int, error) {

	resp, err := s.uService.Update(ctx, req)
	if err != nil {
		log.Print(err)
		switch {
		case errors.Is(err, models.ErrRepeatedPassword):
			return nil, http.StatusBadRequest, err
		case errors.Is(err, models.ErrObjectNotFound):
			return nil, http.StatusNotFound, models.ErrObjectNotFound
		case errors.Is(err, models.ErrValidationFailed):
			return nil, http.StatusBadRequest, models.ErrValidationFailed
		case errors.Is(err, models.ErrEncryptionFailed):
			return nil, http.StatusInternalServerError, models.ErrEncryptionFailed
		case errors.Is(err, models.ErrEmailAlreadyExists):
			return nil, http.StatusBadRequest, models.ErrEmailAlreadyExists
		case errors.Is(err, models.ErrPhoneNumberAlreadyExists):
			return nil, http.StatusBadRequest, models.ErrPhoneNumberAlreadyExists
		case errors.Is(err, models.ErrLoginAlreadyExists):
			return nil, http.StatusBadRequest, models.ErrLoginAlreadyExists
		}
		return nil, http.StatusInternalServerError, models.ErrInternal
	}
	return resp, http.StatusOK, nil
}
