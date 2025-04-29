package server

import (
	"net/http"
)

func (s *Server) configureRouter() {
	r := s.router.NewRoute().Subrouter()
	r.Use(s.logger)
	r.Use(s.cors)

	p := r.NewRoute().Subrouter()
	p.Use(s.userIdentity)

	// TODO: update profile, update libcard, book, library
	// delete

	p.HandleFunc("/profile", s.HandleGetProfile()).Methods(http.MethodGet)

	p.HandleFunc("/books", s.HandleAddBook()).Methods(http.MethodPost)
	p.HandleFunc("/books", s.HandleGetBooks()).Methods(http.MethodGet)
	p.HandleFunc("/books/{id}", s.HandleGetBooks()).Methods(http.MethodGet)

	r.HandleFunc("/sign-up", s.HandleAddUser()).Methods(http.MethodPost, http.MethodOptions)
	r.HandleFunc("/sign-in", s.HandleSignIn()).Methods(http.MethodPost, http.MethodOptions)

	p.HandleFunc("/physbooks", s.HandleGetPhysBooks()).Methods(http.MethodGet, http.MethodOptions)

	p.HandleFunc("/orders", s.HandleAddOrder()).Methods(http.MethodPost)
	p.HandleFunc("/issue", s.HandleIssueOrder()).Methods(http.MethodPost)
	p.HandleFunc("/return", s.HandleReturnOrder()).Methods(http.MethodPost)
	p.HandleFunc("/history", s.HandleGetHistory()).Methods(http.MethodPost, http.MethodOptions)
}
