package server

import (
	"net/http"
)

func (s *Server) configureRouter() {
	s.router.HandleFunc("/books", s.HandleAddBook()).Methods(http.MethodPost)
	s.router.HandleFunc("/books", s.HandleGetBooks()).Methods(http.MethodGet)

	s.router.HandleFunc("/reg", s.HandleAddUser()).Methods(http.MethodPost)
}
