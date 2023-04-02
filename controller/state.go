package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"server/domain"
)

type StateCodes struct {
	States []State `json:"states"`
}

func makeResponse(stateCodes map[string]string) StateCodes {
	states := make([]State, len(stateCodes))
	var i = 0
	for state, code := range stateCodes {
		states[i] = State{Name: state, Code: code}
		i++
	}

	return StateCodes{
		States: states,
	}
}

type State struct {
	Name string `json:"name"`
	Code string `json:"code"`
}

func StateIndex(writer http.ResponseWriter, _ *http.Request) {
	err := json.NewEncoder(writer).Encode(makeResponse(domain.StateToCode))
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)

		if _, err := writer.Write([]byte(fmt.Sprintf("Failed to write states: %v", err))); err != nil {
			log.Printf("Failed writing response: %v\n", err)
		}
	}
}
