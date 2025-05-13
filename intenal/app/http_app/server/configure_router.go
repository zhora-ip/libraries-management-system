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

	p.HandleFunc("/user", s.HandleGetUser()).Methods(http.MethodGet, http.MethodOptions)
	p.HandleFunc("/user", s.HandleUpdateUser()).Methods(http.MethodPatch, http.MethodOptions)
	p.HandleFunc("/user", s.HandleDeleteUser()).Methods(http.MethodDelete, http.MethodOptions)

	p.HandleFunc("/books/create", s.HandleAddBook()).Methods(http.MethodPost, http.MethodOptions)
	p.HandleFunc("/books", s.HandleGetBooks()).Methods(http.MethodPost, http.MethodOptions)
	p.HandleFunc("/books/{id}", s.HandleGetBooks()).Methods(http.MethodGet, http.MethodOptions)

	p.HandleFunc("/libcard", s.HandleGetLibCard()).Methods(http.MethodGet, http.MethodOptions)
	p.HandleFunc("/libcard", s.HandleUpdateLibCard()).Methods(http.MethodPatch, http.MethodOptions)

	r.HandleFunc("/sign-up", s.HandleAddUser()).Methods(http.MethodPost, http.MethodOptions)
	r.HandleFunc("/sign-in", s.HandleSignIn()).Methods(http.MethodPost, http.MethodOptions)

	p.HandleFunc("/physbooks", s.HandleGetPhysBooks()).Methods(http.MethodGet, http.MethodOptions)

	p.HandleFunc("/orders", s.HandleAddOrder()).Methods(http.MethodPost, http.MethodOptions)
	p.HandleFunc("/issue", s.HandleIssueOrder()).Methods(http.MethodPatch, http.MethodOptions)
	p.HandleFunc("/return", s.HandleReturnOrder()).Methods(http.MethodPatch, http.MethodOptions)
	p.HandleFunc("/accept", s.HandleAcceptOrder()).Methods(http.MethodPatch, http.MethodOptions)
	p.HandleFunc("/history", s.HandleGetHistory()).Methods(http.MethodPost, http.MethodOptions)

	r.HandleFunc("/verify", s.HandleEmailVerification()).Methods(http.MethodGet, http.MethodOptions)
}
