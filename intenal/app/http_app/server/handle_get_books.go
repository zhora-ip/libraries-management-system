package server

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/zhora-ip/libraries-management-system/intenal/models"
	svc "github.com/zhora-ip/libraries-management-system/intenal/models/service"
)

const (
	getBooksLimit = 10
)

func (s *Server) HandleGetBooks() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var (
			mp  = mux.Vars(r)
			req = &svc.FindAllBooksRequest{
				Cursor:   time.Now(),
				Limit:    getBooksLimit,
				Backward: false,
			}
			rCache    = &models.Response{}
			bodyBytes = []byte{}
			key       string
			err       error
		)

		bodyBytes, err = io.ReadAll(r.Body)
		if err != nil {
			log.Print(err)
			s.error(w, http.StatusBadRequest, nil)
			return
		}
		key = r.URL.String() + string(bodyBytes)

		defer func() {
			s.rCache.Set(r.Context(), key, rCache)
		}()

		rCache, err = s.rCache.Get(r.Context(), key)
		if err == nil {
			s.respond(w, rCache.Code, rCache.Data)
			return
		}

		err = json.Unmarshal(bodyBytes, req)
		switch {
		case mp != nil:
		case err != nil:
			log.Print(err)
			rCache = s.error(w, http.StatusBadRequest, nil)
			return
		}

		resp, status, err := s.getBooksHelper(r.Context(), req, mp)

		switch {
		case err != nil:
			rCache = s.error(w, status, err)
			return
		case resp == nil || len(resp.Data) == 0:
			rCache = s.respond(w, http.StatusNoContent, nil)
			return
		}

		rCache = s.respond(
			w,
			http.StatusOK,
			map[string]any{
				"books":        resp.Data,
				"first_cursor": resp.Data[0].UpdatedAt.UTC(),
				"last_cursor":  resp.Data[len(resp.Data)-1].UpdatedAt.UTC(),
			},
		)
	}
}

func (s *Server) getBooksHelper(ctx context.Context, req *svc.FindAllBooksRequest, mp map[string]string) (*svc.FindAllBooksResponse, int, error) {

	strID, ok := mp["id"]
	if ok {
		ID, err := strconv.ParseInt(strID, 10, 64)
		if err != nil {
			log.Print(err)
			return nil, http.StatusInternalServerError, models.ErrInternal
		}
		req.ID = &ID
	}

	resp, err := s.bService.FindAll(ctx, req)
	if err != nil {
		log.Print(err)
		switch {
		case errors.Is(err, models.ErrObjectNotFound):
			return nil, http.StatusNoContent, models.ErrObjectNotFound
		}
		return nil, http.StatusInternalServerError, models.ErrInternal
	}

	return resp, http.StatusOK, nil
}
