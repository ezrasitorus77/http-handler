package controller

import (
	"context"
	"fmt"
	"net/http"
	"regexp"
	"strings"

	err "github.com/ezrasitorus77/http-handler/domain"
	"github.com/ezrasitorus77/http-handler/domain/delivery"
	log "github.com/ezrasitorus77/http-handler/helper"
	"github.com/ezrasitorus77/http-handler/internal/consts"
)

type middlewareController delivery.Middleware

var MiddlewareController delivery.Handler

func NewMiddleware(router delivery.Router, service delivery.MiddlewareService) delivery.Handler {
	return &middlewareController{
		Router:  router,
		Service: service,
	}
}

func (obj *middlewareController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var (
		e              error
		hFunc          http.Handler
		params         map[string]string
		attr           delivery.Attributes
		allowedMethods string
	)

	log.INFO(r.Method + ":" + r.URL.Path)

	// default HSTS
	w.Header().Set(consts.HSTS, consts.DefaultHSTS)

	// default CSP
	w.Header().Set(consts.CSP, consts.DefaultCSP)

	// default to JSON
	w.Header().Set("Content-Type", "application/json")

	if e, attr, params = obj.getAttrAndParams(w, r); e != nil {
		if rc, msg := e.(*err.Error).ResponseCode, e.(*err.Error).Reason; rc == consts.RCMethodNotAllowed {
			obj.Service.MethodNotAllowed(w, msg)

			return
		} else if rc == consts.RCNotFound {
			obj.Service.NotFound(w, msg)

			return
		} else {
			obj.Service.BadRequest(w, msg)

			return
		}
	} else {
		// CORS
		if r.Method == "OPTIONS" {
			allowedMethods = obj.getAllowedMethods(attr)

			w.Header().Set(consts.AccesControlAllowMethods, allowedMethods)
			w.WriteHeader(http.StatusOK)

			return
		}

		hFunc = obj.getHandlerFunc(attr, r.Method)

		r = r.WithContext(context.WithValue(r.Context(), consts.ContextParamsKey, params))

		hFunc.ServeHTTP(w, r)
	}
}

func (obj *middlewareController) getAttrAndParams(w http.ResponseWriter, r *http.Request) (e error, attribute delivery.Attributes, params map[string]string) {
	var (
		path         string = r.URL.Path
		re           *regexp.Regexp
		typePattern  string
		typeString   string
		method       string = r.Method
		root         string
		splittedPath []string
		routes       []delivery.Attributes
		errFlag      int    // 0 = no error, 1 = page not found, 2 = method not allowed, 3 = invalid param type
		errSubPath   string // sub path index for invalid param type
	)

	defer func() {
		if e == nil {
			switch errFlag {
			case 1:
				e = &err.Error{
					HttpStatus:   http.StatusNotFound,
					ResponseCode: consts.RCNotFound,
					Reason:       consts.PageNotFoundMessage,
				}
			case 2:
				e = &err.Error{
					HttpStatus:   http.StatusMethodNotAllowed,
					ResponseCode: consts.RCMethodNotAllowed,
					Reason:       consts.NotAllowedMessage,
				}
			case 3:
				e = &err.Error{
					HttpStatus:   http.StatusBadRequest,
					ResponseCode: consts.RCInvalidParamType,
					Reason:       consts.InvalidParamTypeMessage + errSubPath,
				}
			default:
			}
		}
	}()

	routes = obj.Router.GetCollections()
	params = make(map[string]string)

	for _, attr := range routes {
		if attr.Full == path {
			for m, _ := range attr.Pairs {
				if m == method {
					attribute = attr

					return
				}
			}
		} else {
			path = strings.Trim(path, "/")
			splittedPath = strings.Split(path, "/")
			root = splittedPath[0]

			if attr.Root == root {
				if len(attr.SubPaths) == len(splittedPath[1:]) {
					for m, _ := range attr.Pairs {
						// CORS
						if method != "OPTIONS" && m != method {
							errFlag = 2

							continue
						}

						for isp, sp := range splittedPath[1:] {
							errFlag = 0

							for _, sub := range attr.SubPaths {
								if sub.Index == isp {
									if sub.Path != sp {
										errFlag = 1

										if sub.Param.Key != "" {
											typePattern, typeString = strings.Split(sub.Param.Type, ":")[0], strings.Split(sub.Param.Type, ":")[1]

											re, e = regexp.Compile(typePattern)
											if e != nil {
												return
											}

											if re.MatchString(sp) {
												if typeString == "float" {
													re, e = regexp.Compile(`c`)
													if e != nil {
														return
													}

													params[sub.Param.Key] = re.ReplaceAllString(sp, ".")
												} else {
													params[sub.Param.Key] = sp
												}

												errFlag = 0
												attribute = attr

												break
											} else {
												errSubPath = fmt.Sprintf("%d (want: %s)", sub.Index+1, typeString)
												errFlag = 3

												return
											}
										}
									}
								}
							}

							if attribute.ID == 0 {
								break
							}
						}

						if attribute.ID != 0 {
							break
						}
					}
				} else {
					errFlag = 1
				}
			} else {
				errFlag = 1
			}
		}

		if attribute.ID != 0 {
			return
		}
	}

	return
}

func (obj *middlewareController) getAllowedMethods(attr delivery.Attributes) (allowedMethods string) {
	var methods []string

	for m, _ := range attr.Pairs {
		methods = append(methods, m)
	}

	allowedMethods = strings.Join(methods, ", ")

	return
}

func (obj *middlewareController) getHandlerFunc(attr delivery.Attributes, method string) (hFunc http.Handler) {
	for m, f := range attr.Pairs {
		if m == method {
			hFunc = f
		}
	}

	return
}
