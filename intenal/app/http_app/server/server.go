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
	FindAll(context.Context, *svc.FindAllBooksRequest) (*svc.FindAllBooksResponse, error)
}

type userService interface {
	Add(context.Context, *svc.AddUserRequest) (*svc.AddUserResponse, error)
	GenerateToken(context.Context, *svc.GenerateTokenRequest) (*svc.GenerateTokenResponse, error)
	FindByID(context.Context, *svc.FindUserByIDRequest) (*svc.FindUserByIDResponse, error)
	Delete(context.Context, *svc.DeleteUserRequest) (*svc.DeleteUserResponse, error)
	Update(context.Context, *svc.UpdateUserRequest) (*svc.UpdateUserResponse, error)
	FindLibCard(context.Context, *svc.FindLibCardRequest) (*svc.FindLibCardResponse, error)
	ExtendLibCard(context.Context, *svc.ExtendLibCardRequest) error
}

type physBookService interface {
	FindByBookID(context.Context, int64) (*svc.FindPBookByBookIDResponse, error)
}

type orderService interface {
	Add(context.Context, *svc.AddOrderRequest) (*svc.AddOrderResponse, error)
	Issue(context.Context, *svc.IssueOrderRequest) (*svc.IssueOrderResponse, error)
	Return(context.Context, *svc.ReturnOrderRequest) (*svc.ReturnOrderResponse, error)
	Accept(context.Context, *svc.AcceptOrderRequest) (*svc.AcceptOrderResponse, error)
	FindAll(context.Context, *svc.FindAllOrdersRequest) (*svc.FindAllOrdersResponse, error)
	Submit(any, chan<- error)
}

type cache interface {
	Set(context.Context, string, *models.Response) error
	Get(context.Context, string) (*models.Response, error)
}

type tkManager interface {
	Parse(accessToken string) (int64, int32, error)
}

type Server struct {
	router    *mux.Router
	bService  bookService
	uService  userService
	pbService physBookService
	oService  orderService
	tkManager tkManager
	rCache    cache
}

func New(bs bookService, us userService, pbs physBookService, os orderService, tm tkManager, ch cache) *Server {
	srv := &Server{
		router:    mux.NewRouter(),
		bService:  bs,
		uService:  us,
		pbService: pbs,
		oService:  os,
		tkManager: tm,
		rCache:    ch,
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
