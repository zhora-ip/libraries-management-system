package server

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/zhora-ip/libraries-management-system/intenal/models"
	svc "github.com/zhora-ip/libraries-management-system/intenal/models/service"
)

func (s *Server) HandleSignIn() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			req = &svc.GenerateTokenRequest{}
		)

		err := json.NewDecoder(r.Body).Decode(req)
		if err != nil {
			log.Print(err)
			s.error(w, http.StatusBadRequest, nil)
			return
		}

		resp, err := s.uService.GenerateToken(r.Context(), req)
		if err != nil {
			log.Print(err)
			switch {
			case errors.Is(err, models.ErrValidationFailed):
				s.error(w, http.StatusBadRequest, models.ErrValidationFailed)
				return
			case errors.Is(err, models.ErrForbidden):
				s.error(w, http.StatusForbidden, nil)
				return
			}

			s.error(w, http.StatusInternalServerError, nil)
			return
		}

		s.respond(w, http.StatusOK, resp)
	}
}
