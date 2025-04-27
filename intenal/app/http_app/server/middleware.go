package server

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/zhora-ip/libraries-management-system/intenal/models"
)

type (
	ctxKeyUserID   struct{}
	ctxKeyUserRole struct{}
)

func (s *Server) logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		bodyBytes, _ := io.ReadAll(r.Body)
		r.Body.Close()
		r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

		recorder := NewResponseRecorder(w)

		auditRequest := &models.AuditRequest{
			Method: r.Method,
			URL:    r.URL.Path,
		}

		if len(r.URL.Query()) > 0 {
			auditRequest.Query = fmt.Sprintf("%v", r.URL.Query())
		}

		if len(bodyBytes) > 0 {
			auditRequest.Body = string(bodyBytes)
		}

		s.oService.Submit(auditRequest, nil)

		next.ServeHTTP(recorder, r)

		auditResponse := &models.AuditResponse{
			Code: strconv.Itoa(recorder.StatusCode()),
		}

		if r.Method != http.MethodGet && len(recorder.Body()) > 0 {
			auditResponse.Body = strings.TrimSuffix(recorder.Body(), "\n")
		}

		s.oService.Submit(auditResponse, nil)
	})
}

func (s *Server) userIdentity(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		token := r.Header.Get("Authorization")

		user, role, err := s.tkManager.Parse(token)
		if err != nil {
			log.Print(err)
			s.error(w, http.StatusBadRequest, nil)
			return
		}

		ctx := context.WithValue(context.Background(), ctxKeyUserID{}, user)
		ctx = context.WithValue(ctx, ctxKeyUserRole{}, role)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
