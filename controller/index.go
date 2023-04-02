package controller

import (
	"encoding/json"
	"fmt"
	"github.com/chzyer/logex"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type availableRoutes struct {
	Methods map[string][]string `json:"method"`
}

func NewIndex(routes []Route) func(http.ResponseWriter, *http.Request, httprouter.Params) {
	notFoundHandler := http.NotFoundHandler()

	return func(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
		if request.URL.Path != "/" {
			notFoundHandler.ServeHTTP(writer, request)
			return
		}

		methodToPaths := make(map[string][]string)
		for _, route := range routes {
			paths, exists := methodToPaths[route.Verb]
			if !exists {
				paths = []string{}
			}
			methodToPaths[route.Verb] = append(paths, route.Path)
		}

		err := json.
			NewEncoder(writer).
			Encode(availableRoutes{Methods: methodToPaths})

		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			if _, err := writer.Write([]byte(fmt.Sprintf("Failed created list of available routes: %v", err))); err != nil {
				logex.Info(fmt.Sprintf("Failed responding to request: %v", err))
			}
		}
	}
}
