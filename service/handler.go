package service

import (
	"net/http"
)

type handlerFunc func(w http.ResponseWriter, r *http.Request)

func (f handlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f(w, r)
}
