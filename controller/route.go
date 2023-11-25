package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Route struct {
	Verb    string
	Path    string
	Handler func(http.ResponseWriter, *http.Request, httprouter.Params)
}
