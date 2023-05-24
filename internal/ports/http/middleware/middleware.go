package middleware

import (
	"net/http"
	"strings"
)

type Middleware interface {
	Auth(next http.HandlerFunc) http.HandlerFunc
	MatchPattern(next http.HandlerFunc, pattern string) http.HandlerFunc
	POST(next http.HandlerFunc) http.HandlerFunc
	GET(next http.HandlerFunc) http.HandlerFunc
	PUT(next http.HandlerFunc) http.HandlerFunc
	DELETE(next http.HandlerFunc) http.HandlerFunc
}

type middleware struct {
}

func New() Middleware {
	return &middleware{}
}

func (m *middleware) Auth(next http.HandlerFunc) http.HandlerFunc {
	// TODO: implement
	return next
}

func (m *middleware) MatchPattern(next http.HandlerFunc, pattern string) http.HandlerFunc {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			if strings.Contains(strings.TrimPrefix(r.URL.Path, pattern), "/") {
				w.WriteHeader(http.StatusNotFound)
				return
			}
			next.ServeHTTP(w, r)
		})
}

func (m *middleware) POST(next http.HandlerFunc) http.HandlerFunc {
	return methodMiddleware(http.MethodPost, next)
}

func (m *middleware) GET(next http.HandlerFunc) http.HandlerFunc {
	return methodMiddleware(http.MethodGet, next)
}

func (m *middleware) PUT(next http.HandlerFunc) http.HandlerFunc {
	return methodMiddleware(http.MethodPut, next)
}

func (m *middleware) DELETE(next http.HandlerFunc) http.HandlerFunc {
	return methodMiddleware(http.MethodDelete, next)
}

func methodMiddleware(method string, next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			if r.Method != method {
				w.WriteHeader(http.StatusMethodNotAllowed)
				return
			}
			next.ServeHTTP(w, r)
		})
}
