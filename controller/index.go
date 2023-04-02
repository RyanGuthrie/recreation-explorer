package controller

import (
	"encoding/json"
	"fmt"
	"github.com/chzyer/logex"
	"net/http"
)

type availableRoutes struct {
	Routes []string `json:"routes"`
}

func NewIndex(routes []string) func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, _ *http.Request) {
		err := json.
			NewEncoder(writer).
			Encode(availableRoutes{Routes: routes})

		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			if _, err := writer.Write([]byte(fmt.Sprintf("Failed created list of available routes: %v", err))); err != nil {
				logex.Info(fmt.Sprintf("Failed responding to request: %v", err))
			}
		}
	}
}
