package server

import (
	"net/http"
)

func (s *Server) configureRouter() {
	s.router.HandleFunc("/books", s.HandleAddBook()).Methods(http.MethodPost)
	s.router.HandleFunc("/books", s.HandleGetBooks()).Methods(http.MethodGet)

	s.router.HandleFunc("/reg", s.HandleAddUser()).Methods(http.MethodPost)

	s.router.HandleFunc("/physbooks", s.HandleGetPhysBooks()).Methods(http.MethodGet)

	s.router.HandleFunc("/orders", s.HandleAddOrder()).Methods(http.MethodPost)
	s.router.HandleFunc("/issue", s.HandleIssueOrder()).Methods(http.MethodPost)
	s.router.HandleFunc("/return", s.HandleReturnOrder()).Methods(http.MethodPost)
	s.router.HandleFunc("/history", s.HandleGetHistory()).Methods(http.MethodGet)
}
