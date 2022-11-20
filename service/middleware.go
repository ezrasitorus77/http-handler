package service

import (
	"net/http"

	"github.com/ezrasitorus77/http-handler/domain/delivery"
	log "github.com/ezrasitorus77/http-handler/helper"
	response "github.com/ezrasitorus77/http-handler/helper"
	"github.com/ezrasitorus77/http-handler/internal/consts"
)

type middlewareService struct{}

var MiddlewareService delivery.MiddlewareService

func init() {
	MiddlewareService = &middlewareService{}
}

func (obj *middlewareService) MethodNotAllowed(w http.ResponseWriter, message string) {
	var respData delivery.ResponseData = delivery.ResponseData{
		Description:  consts.NotAllowedMessage,
		ErrorMessage: message,
	}

	log.ERROR(message)

	response.Send(w, http.StatusMethodNotAllowed, consts.RCMethodNotAllowed, respData)
}

func (obj *middlewareService) NotFound(w http.ResponseWriter, message string) {
	var respData delivery.ResponseData = delivery.ResponseData{
		Description:  consts.NotFoundMessage,
		ErrorMessage: message,
	}

	log.ERROR(message)

	response.Send(w, http.StatusNotFound, consts.RCNotFound, respData)
}

func (obj *middlewareService) BadRequest(w http.ResponseWriter, message string) {
	var respData delivery.ResponseData = delivery.ResponseData{
		Description:  consts.BadRequestMessage,
		ErrorMessage: message,
	}

	log.ERROR(message)

	response.Send(w, http.StatusInternalServerError, consts.RCInternalServerError, respData)
}

func (obj *middlewareService) Panic(w http.ResponseWriter, message string) {
	var respData delivery.ResponseData = delivery.ResponseData{
		Description: message,
	}

	log.ERROR(message)

	response.Send(w, http.StatusInternalServerError, consts.RCInternalServerError, respData)
}

func (obj *middlewareService) CSP(w http.ResponseWriter, script string) {
	w.Header().Set(consts.CSP, script)
}

func (obj *middlewareService) HSTS(w http.ResponseWriter, script string) {
	w.Header().Set(consts.HSTS, script)
}
