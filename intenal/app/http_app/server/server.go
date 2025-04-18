package server

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/zhora-ip/libraries-management-system/intenal/models"
	svc "github.com/zhora-ip/libraries-management-system/intenal/models/service"
)

type bookService interface {
	Add(context.Context, *svc.AddBookRequest) (*svc.AddBookResponse, error)
	FindAll(context.Context, *svc.FindAllRequest) (*svc.FindAllResponse, error)
}

type userService interface {
	Add(context.Context, *svc.AddUserRequest) (*svc.AddUserResponse, error)
}

type Server struct {
	router   *mux.Router
	bService bookService
	uService userService
}

func New(bs bookService, us userService) *Server {
	srv := &Server{
		router:   mux.NewRouter(),
		bService: bs,
		uService: us,
	}

	srv.configureRouter()
	return srv
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *Server) error(w http.ResponseWriter, code int, err error) *models.Response {
	if err != nil {
		return s.respond(w, code, map[string]string{"error": err.Error()})
	}
	return s.respond(w, code, nil)
}

func (s *Server) respond(w http.ResponseWriter, code int, data interface{}) *models.Response {
	w.WriteHeader(code)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
	return &models.Response{
		Code: code,
		Data: data,
	}
}
