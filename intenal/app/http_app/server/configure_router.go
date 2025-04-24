package server

import (
	"net/http"
)

func (s *Server) configureRouter() {
	r := s.router.NewRoute().Subrouter()
	r.Use(s.logger)

	r.HandleFunc("/books", s.HandleAddBook()).Methods(http.MethodPost)
	r.HandleFunc("/books", s.HandleGetBooks()).Methods(http.MethodGet)

	r.HandleFunc("/reg", s.HandleAddUser()).Methods(http.MethodPost)

	r.HandleFunc("/physbooks", s.HandleGetPhysBooks()).Methods(http.MethodGet)

	r.HandleFunc("/orders", s.HandleAddOrder()).Methods(http.MethodPost)
	r.HandleFunc("/issue", s.HandleIssueOrder()).Methods(http.MethodPost)
	r.HandleFunc("/return", s.HandleReturnOrder()).Methods(http.MethodPost)
	r.HandleFunc("/history", s.HandleGetHistory()).Methods(http.MethodGet)
}
