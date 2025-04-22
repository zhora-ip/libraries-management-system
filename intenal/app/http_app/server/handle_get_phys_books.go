package server

import (
	"context"
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/zhora-ip/libraries-management-system/intenal/models"
	svc "github.com/zhora-ip/libraries-management-system/intenal/models/service"
)

func (s *Server) HandleGetPhysBooks() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		strID := r.URL.Query().Get("book_id")

		ID, err := strconv.Atoi(strID)
		if err != nil {
			s.error(w, http.StatusBadRequest, nil)
			log.Print(err)
			return
		}
		resp, status, err := s.handleGetPhysBooksHelper(r.Context(), int64(ID))
		if resp != nil {
			s.respond(w, http.StatusOK, resp.Data)
			return
		}

		s.error(w, status, err)
	}
}

func (s *Server) handleGetPhysBooksHelper(ctx context.Context, ID int64) (*svc.FindPBookByBookIDResponse, int, error) {
	resp, err := s.pbService.FindByBookID(ctx, ID)
	if err != nil {
		log.Print(err)
		switch {
		case errors.Is(err, models.ErrObjectNotFound):
			return nil, http.StatusNotFound, models.ErrObjectNotFound
		}
		return nil, http.StatusInternalServerError, nil
	}

	return resp, http.StatusOK, nil
}
