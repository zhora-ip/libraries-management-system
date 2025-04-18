package server

import (
	"net/http"
)

func (s *Server) configureRouter() {
	s.router.HandleFunc("/books", s.HandleAddBook()).Methods(http.MethodPost)
}
