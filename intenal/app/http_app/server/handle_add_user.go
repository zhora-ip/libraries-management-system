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

func (s *Server) HandleAddUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := &svc.AddUserRequest{}

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			s.error(w, http.StatusBadRequest, nil)
			log.Print(err)
			return
		}

		resp, status, err := s.handleAddUserHelper(r.Context(), req)
		if resp != nil {
			s.respond(w, status, resp)
			return
		}

		s.error(w, status, err)

	}
}

func (s *Server) handleAddUserHelper(ctx context.Context, req *svc.AddUserRequest) (*svc.AddUserResponse, int, error) {

	resp, err := s.uService.Add(ctx, req)
	if err != nil {
		switch {
		case errors.Is(err, models.ErrValidationFailed):
			return nil, http.StatusBadRequest, err
		case errors.Is(err, models.ErrEncryptionFailed):
			return nil, http.StatusInternalServerError, err
		case errors.Is(err, models.ErrEmailAlreadyExists):
			return nil, http.StatusBadRequest, err
		case errors.Is(err, models.ErrPhoneNumberAlreadyExists):
			return nil, http.StatusBadRequest, err
		case errors.Is(err, models.ErrLoginAlreadyExists):
			return nil, http.StatusBadRequest, err
		}
		log.Print(err)
		return nil, http.StatusInternalServerError, nil
	}
	return resp, http.StatusOK, nil

}
