package service

import (
	"fmt"
	"net/http"

	"github.com/ezrasitorus77/http-handler/domain/delivery"
	"github.com/ezrasitorus77/http-handler/internal/consts"
	log "github.com/ezrasitorus77/http-handler/internal/helper"
	response "github.com/ezrasitorus77/http-handler/internal/helper"

	h "github.com/julienschmidt/httprouter"
)

type middlewareService delivery.MiddlewareFuncs

var MiddlewareService delivery.MiddlewareService

func init() {
	fmt.Println(h.New())
	MiddlewareService = &middlewareService{}
}

func (obj *middlewareService) MethodNotAllowed(next http.Handler, message string) http.Handler {
	var respData delivery.ResponseData = delivery.ResponseData{
		Description:  consts.NotAllowedMessage,
		ErrorMessage: message,
	}

	obj.IsMethodNotAllowed = true

	log.ERROR(message)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response.Send(w, http.StatusMethodNotAllowed, consts.RCMethodNotAllowed, respData)

		next.ServeHTTP(w, r)
	})
}

func (obj *middlewareService) NotFound(next http.Handler, message string) http.Handler {
	var respData delivery.ResponseData = delivery.ResponseData{
		Description:  consts.NotFoundMessage,
		ErrorMessage: message,
	}

	obj.IsNotFound = true

	log.ERROR(message)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response.Send(w, http.StatusNotFound, consts.RCNotFound, respData)

		next.ServeHTTP(w, r)
	})
}

func (obj *middlewareService) BadRequest(next http.Handler, message string) http.Handler {
	var respData delivery.ResponseData = delivery.ResponseData{
		Description:  consts.BadRequestMessage,
		ErrorMessage: message,
	}

	obj.IsBadRequest = true

	log.ERROR(message)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response.Send(w, http.StatusInternalServerError, consts.RCInternalServerError, respData)

		next.ServeHTTP(w, r)
	})
}

func (obj *middlewareService) Panic(next http.Handler, message string) http.Handler {
	var respData delivery.ResponseData = delivery.ResponseData{
		Description: message,
	}

	obj.IsPanic = true

	log.ERROR(message)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response.Send(w, http.StatusInternalServerError, consts.RCInternalServerError, respData)

		next.ServeHTTP(w, r)
	})
}

func (obj *middlewareService) CSP(next http.Handler, script string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(consts.CSP, script)
		next.ServeHTTP(w, r)
	})
}

func (obj *middlewareService) HSTS(next http.Handler, script string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(consts.HSTS, script)
		next.ServeHTTP(w, r)
	})
}
