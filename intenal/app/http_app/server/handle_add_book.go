package server

import (
	"encoding/json"
	"net/http"

	bookservice "github.com/zhora-ip/libraries-management-system/intenal/app/service/book"
)

func (s *Server) HandleAddBook() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := &bookservice.AddBookRequest{}

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			s.error(w, http.StatusBadRequest, err)
		}

		resp, err := s.bService.Add(r.Context(), req)
		if err != nil {
			s.error(w, http.StatusInternalServerError, err)
		}
		s.respond(w, http.StatusOK, resp)
	}
}
