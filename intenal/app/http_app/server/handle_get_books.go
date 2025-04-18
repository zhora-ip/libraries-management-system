package server

import (
	"log"
	"net/http"

	svc "github.com/zhora-ip/libraries-management-system/intenal/models/service"
)

func (s *Server) HandleGetBooks() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := &svc.FindAllRequest{}

		resp, err := s.bService.FindAll(r.Context(), req)
		if err != nil {
			s.error(w, http.StatusInternalServerError, nil)
			log.Print(err)
			return
		}

		s.respond(w, http.StatusOK, resp.Data)
	}
}
