package server

import (
	"errors"
	"log"
	"net/http"

	"github.com/zhora-ip/libraries-management-system/intenal/models"
	svc "github.com/zhora-ip/libraries-management-system/intenal/models/service"
)

func (s *Server) HandleGetBooks() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := &svc.FindAllBooksRequest{}

		resp, err := s.bService.FindAll(r.Context(), req)
		if err != nil {
			log.Print(err)
			switch {
			case errors.Is(err, models.ErrObjectNotFound):
				s.error(w, http.StatusNotFound, models.ErrObjectNotFound)
				return
			}
			s.error(w, http.StatusInternalServerError, nil)
			return
		}

		s.respond(w, http.StatusOK, resp.Data)
	}
}
