package server

import (
	"log"
	"net/http"

	svc "github.com/zhora-ip/libraries-management-system/intenal/models/service"
)

func (s *Server) HandleEmailVerification() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.URL.Query().Get("token")
		if token == "" {
			s.error(w, http.StatusBadRequest, nil)
			return
		}

		_, err := s.uService.ConfirmEmail(r.Context(), &svc.ConfirmEmailRequest{Token: token})
		if err != nil {
			log.Print(err)
			s.error(w, http.StatusInternalServerError, nil)
			return
		}
		s.respond(w, http.StatusOK, nil)
	}
}
