package delivery

import (
	"net/http"
)

type (
	Routes struct {
		Collections []Attributes
	}

	Attributes struct {
		ID       int
		Root     string
		Full     string
		SubPaths []SubPath
		// method:HFunc
		Pairs map[string]http.Handler
		// Method   string
		// HFunc    http.Handler
	}

	SubPath struct {
		Index int
		Path  string // with param prefix
		Param Param
	}

	Param struct {
		Key string
		// type consists of regex pattern
		// (default)s:string, d:int, f:float
		// since it's not common to use ',' in a URI, then ->f: param digits should use 'c' instead
		// ex: ->f:39c12 means 39.12; this is overided in middleware's routing()
		Type string
	}

	Router interface {
		GET(path string, hFunc http.HandlerFunc)
		POST(path string, hFunc http.HandlerFunc)
		PATCH(path string, hFunc http.HandlerFunc)
		PUT(path string, hFunc http.HandlerFunc)
		DELETE(path string, hFunc http.HandlerFunc)
		GetCollections() []Attributes
	}
)
