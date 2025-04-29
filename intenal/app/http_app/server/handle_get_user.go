package server

import (
	"errors"
	"log"
	"net/http"

	"github.com/zhora-ip/libraries-management-system/intenal/models"
	"github.com/zhora-ip/libraries-management-system/intenal/models/service"
)

func (s *Server) HandleGetProfile() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			userID = r.Context().Value(ctxKeyUserID{}).(int64)
		)

		resp, err := s.uService.FindByID(r.Context(), &service.FindUserByIDRequest{ID: userID})
		if err != nil {
			log.Print(err)
			switch {
			case errors.Is(err, models.ErrObjectNotFound):
				s.error(w, http.StatusNotFound, nil)
				return
			}
			s.error(w, http.StatusInternalServerError, nil)
			return
		}

		s.respond(w, http.StatusOK, resp)
	}
}
