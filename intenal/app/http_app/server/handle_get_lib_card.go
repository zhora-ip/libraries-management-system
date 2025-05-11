package server

import (
	"log"
	"net/http"

	svc "github.com/zhora-ip/libraries-management-system/intenal/models/service"
)

func (s *Server) HandleGetLibCard() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			userID = r.Context().Value(ctxKeyUserID{}).(int64)
		)

		resp, err := s.uService.FindLibCard(r.Context(), &svc.FindLibCardRequest{UserID: userID})
		if err != nil {
			log.Print(err)
			switch {

			}
			s.error(w, http.StatusInternalServerError, nil)
			return
		}

		s.respond(w, http.StatusOK, resp)
	}
}
