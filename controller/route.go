package controller

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type Route struct {
	Verb    string
	Path    string
	Handler func(http.ResponseWriter, *http.Request, httprouter.Params)
}

func (r Route) asTuple() (string, string) {
	return r.Verb, r.Path
}
