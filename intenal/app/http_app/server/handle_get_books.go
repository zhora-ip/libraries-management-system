package server

import (
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/zhora-ip/libraries-management-system/intenal/models"
	svc "github.com/zhora-ip/libraries-management-system/intenal/models/service"
)

func (s *Server) HandleGetBooks() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			mp  = mux.Vars(r)
			req = &svc.FindAllBooksRequest{}
			err error
		)

		strID, ok := mp["id"]
		if ok {
			ID, err := strconv.ParseInt(strID, 10, 64)
			if err != nil {
				log.Print(err)
				s.error(w, http.StatusBadRequest, nil)
				return
			}

			req.ID = &ID
		}

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
