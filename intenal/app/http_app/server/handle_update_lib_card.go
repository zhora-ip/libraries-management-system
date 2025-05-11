package server

import (
	"errors"
	"log"
	"net/http"

	"github.com/zhora-ip/libraries-management-system/intenal/models"
	"github.com/zhora-ip/libraries-management-system/intenal/models/service"
)

func (s *Server) HandleUpdateLibCard() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			userID = r.Context().Value(ctxKeyUserID{}).(int64)
		)

		err := s.uService.ExtendLibCard(r.Context(), &service.ExtendLibCardRequest{UserID: userID})
		if err != nil {
			log.Print(err)
			switch {
			case errors.Is(err, models.ErrCardNotExpired):
				s.error(w, http.StatusBadRequest, err)
				return
			}
			s.error(w, http.StatusInternalServerError, nil)
			return
		}

		s.respond(w, http.StatusOK, nil)
	}
}
