package delivery

import (
	"net/http"
)

type (
	Middleware struct {
		Router  Router
		Service MiddlewareService
	}

	Funcs func(next http.Handler, message string) http.Handler

	MiddlewareService interface {
		MethodNotAllowed(w http.ResponseWriter, message string)
		NotFound(w http.ResponseWriter, message string)
		BadRequest(w http.ResponseWriter, message string)
		Panic(w http.ResponseWriter, message string)
		CSP(w http.ResponseWriter, message string)
		HSTS(w http.ResponseWriter, message string)
	}

	Handler interface {
		ServeHTTP(w http.ResponseWriter, r *http.Request)
	}
)
