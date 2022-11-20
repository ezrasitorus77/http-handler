package service

import (
	"fmt"
	"net/http"
	"regexp"
	"strings"

	err "github.com/ezrasitorus77/http-handler/domain"
	"github.com/ezrasitorus77/http-handler/domain/delivery"
	log "github.com/ezrasitorus77/http-handler/helper"
	"github.com/ezrasitorus77/http-handler/internal/consts"
)

type routerService delivery.Routes

var RouterService delivery.Router

func init() {
	RouterService = &routerService{}
}

func (obj *routerService) GET(path string, hFunc http.HandlerFunc) {
	if e := obj.addRoute(path, http.MethodGet, hFunc); e != nil {
		log.PANIC(e.Error())
	}
}

func (obj *routerService) POST(path string, hFunc http.HandlerFunc) {
	if e := obj.addRoute(path, http.MethodPost, hFunc); e != nil {
		log.PANIC(e.Error())
	}

	return
}

func (obj *routerService) PATCH(path string, hFunc http.HandlerFunc) {
	if e := obj.addRoute(path, http.MethodPatch, hFunc); e != nil {
		log.PANIC(e.Error())
	}

	return
}

func (obj *routerService) PUT(path string, hFunc http.HandlerFunc) {
	if e := obj.addRoute(path, http.MethodPut, hFunc); e != nil {
		log.PANIC(e.Error())
	}

	return
}

func (obj *routerService) DELETE(path string, hFunc http.HandlerFunc) {
	if e := obj.addRoute(path, http.MethodDelete, hFunc); e != nil {
		log.PANIC(e.Error())
	}

	return
}

func (obj *routerService) addRoute(path, method string, hFunc http.HandlerFunc) (e error) {
	var (
		splittedPath []string
		attributes   delivery.Attributes
		param        delivery.Param
	)

	if e = validatePath(path, true); e != nil {
		return
	}

	attributes.ID = len(obj.Collections) + 1
	attributes.Full = path

	if e = obj.checkRoutes(path, method); e != nil {
		return
	}

	path = strings.Trim(path, "/")
	splittedPath = strings.Split(path, "/")

	if attributes.Pairs == nil {
		attributes.Pairs = make(map[string]http.Handler)
	}
	attributes.Pairs[method] = hFunc

	if splittedPath[0] == "" {
		attributes.Root = "/"
	} else {
		attributes.Root = splittedPath[0]

		for i, sub := range splittedPath[1:] {
			if e = validatePath(sub, false); e != nil {
				return
			}

			param, e = checkParams(sub)
			if e != nil {
				return
			}

			attributes.SubPaths = append(attributes.SubPaths, delivery.SubPath{
				Index: i, // without root
				Path:  sub,
				Param: param,
			})
		}
	}

	obj.Collections = append(obj.Collections, attributes)

	return
}

// checkParams returns Param{} if param exists
func checkParams(path string) (param delivery.Param, e error) {
	var (
		re          *regexp.Regexp
		paramFormat int
		typ         string // param data type
	)

	re, e = regexp.Compile(consts.ParamPrefixRegex)
	if e != nil {
		return
	}

	paramFormat = len(re.FindAllString(path, -1))
	if paramFormat == 0 {
		return
	} else {
		typ = string(path[2])

		re, e = regexp.Compile(consts.ParamKeyRegex)
		if e != nil {
			return
		}

		param.Key = re.FindString(path)

		switch typ {
		case "d":
			param.Type = fmt.Sprintf("%s:int", consts.IntRegex)
		case "f":
			param.Type = fmt.Sprintf("%s:float", consts.FloatRegex)
		default:
			param.Type = fmt.Sprintf("%s:string", consts.StringRegex)
		}
	}

	return
}

// validatePath validates trimmed path
func validatePath(path string, isFullPath bool) (e error) {
	var (
		re *regexp.Regexp
	)

	if path == "" {
		return &err.Error{
			HttpStatus:   http.StatusBadRequest,
			ResponseCode: consts.RCInvalidPathFormat,
			Reason:       consts.InvalidPathFormat + path,
		}
	}

	if isFullPath {
		if path[0] != '/' {
			return &err.Error{
				HttpStatus:   http.StatusBadRequest,
				ResponseCode: consts.RCInvalidPathFormat,
				Reason:       consts.InvalidPathFormat + path,
			}
		}

		re, e = regexp.Compile(consts.FullPathRegex)
	} else {
		re, e = regexp.Compile(consts.SubPathRegex)
	}

	if e != nil {
		return
	}

	if re.MatchString(path) {
		return nil
	}

	return &err.Error{
		HttpStatus:   http.StatusBadRequest,
		ResponseCode: consts.RCInvalidPathFormat,
		Reason:       consts.InvalidPathFormat + path,
	}
}

func (obj *routerService) checkRoutes(path, method string) error {
	var (
		splittedPath    []string
		root            string
		errInvalidRoute err.Error = err.Error{
			HttpStatus:   http.StatusBadRequest,
			ResponseCode: consts.RCInvalidAddRoute,
			Reason:       consts.InvalidAddRoute + path + "with method " + method,
		}
	)

	for _, attr := range obj.Collections {
		if attr.Full == path {
			for m := range attr.Pairs {
				if m == method {
					return &errInvalidRoute
				}
			}
		} else {
			splittedPath = strings.Split(strings.Trim(path, "/"), "/")
			root = splittedPath[0]

			if attr.Root == root {
				if len(attr.SubPaths) == len(splittedPath[1:]) {
					var sameNextSub int

					for isp, sp := range splittedPath[1:] {
						for isub, sub := range attr.SubPaths {
							if sub.Path == sp {
								for m := range attr.Pairs {
									if m == method {
										if isub != isp {
											// neglects the rest of subpaths because they must inherit last subpath with current attr.Method
											// ex: GET = /root/sub1/->d:id/sub2 must not overlap GET = /root/sub1/->d:p/sub3
											return &err.Error{
												HttpStatus:   http.StatusBadRequest,
												ResponseCode: consts.RCOverlappingPathFormat,
												Reason:       consts.OverlappingPathFormat + path,
											}
										}

										return &errInvalidRoute
									}
								}
							} else {
								if len(splittedPath)-1 >= isp+2 && len(attr.SubPaths)-1 >= isub+1 {
									if splittedPath[isp+2] == attr.SubPaths[isub+1].Path {
										sameNextSub++
										if sameNextSub == len(splittedPath[1:]) {
											for m := range attr.Pairs {
												if m == method {
													return &err.Error{
														HttpStatus:   http.StatusBadRequest,
														ResponseCode: consts.RCOverlappingPathFormat,
														Reason:       consts.OverlappingPathFormat + path + " with different param type",
													}
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}

	return nil
}

func (obj *routerService) GetCollections() []delivery.Attributes {
	return obj.Collections
}
