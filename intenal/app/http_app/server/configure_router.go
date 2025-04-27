package server

import (
	"net/http"
)

func (s *Server) configureRouter() {
	r := s.router.NewRoute().Subrouter()
	r.Use(s.logger)

	p := r.NewRoute().Subrouter()
	p.Use(s.userIdentity)

	p.HandleFunc("/books", s.HandleAddBook()).Methods(http.MethodPost)
	r.HandleFunc("/books", s.HandleGetBooks()).Methods(http.MethodGet)

	r.HandleFunc("/sign-up", s.HandleAddUser()).Methods(http.MethodPost)
	r.HandleFunc("/sign-in", s.HandleSignIn()).Methods(http.MethodPost)

	r.HandleFunc("/physbooks", s.HandleGetPhysBooks()).Methods(http.MethodGet)

	p.HandleFunc("/orders", s.HandleAddOrder()).Methods(http.MethodPost)
	p.HandleFunc("/issue", s.HandleIssueOrder()).Methods(http.MethodPost)
	p.HandleFunc("/return", s.HandleReturnOrder()).Methods(http.MethodPost)
	p.HandleFunc("/history", s.HandleGetHistory()).Methods(http.MethodGet)
}
