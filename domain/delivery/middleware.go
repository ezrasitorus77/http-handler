package delivery

import (
	"net/http"
)

type (
	Middleware struct {
		Router  Router
		Service []Funcs
	}

	MiddlewareFuncs struct {
		IsMethodNotAllowed bool
		IsNotFound         bool
		IsBadRequest       bool
		IsPanic            bool
	}

	Funcs func(next http.Handler, message string) http.Handler

	MiddlewareService interface {
		MethodNotAllowed(next http.Handler, message string) http.Handler
		NotFound(next http.Handler, message string) http.Handler
		BadRequest(next http.Handler, message string) http.Handler
		Panic(next http.Handler, message string) http.Handler
		CSP(next http.Handler, message string) http.Handler
		HSTS(next http.Handler, message string) http.Handler
	}

	Handler interface {
		ServeHTTP(w http.ResponseWriter, r *http.Request)
	}
)
