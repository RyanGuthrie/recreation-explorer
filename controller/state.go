package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
	"server/domain"
)

type State struct {
	Name string `json:"name"`
	Code string `json:"code"`
}

type StateCodes struct {
	States []State `json:"states"`
}

func StateIndex(writer http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	var randomDuration time.Duration = time.Duration(rand.Int63n(500)) * time.Millisecond
	time.Sleep(randomDuration)

	err := json.NewEncoder(writer).Encode(domainStateToJson(domain.States))
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)

		if _, err := writer.Write([]byte(fmt.Sprintf("Failed to write states: %v", err))); err != nil {
			log.Printf("Failed writing response: %v\n", err)
		}
	}
}

func domainStateToJson(stateCodes []domain.State) StateCodes {
	states := make([]State, len(stateCodes))
	for i, state := range stateCodes {
		states[i] = State{Name: state.Name, Code: state.Code}
		i++
	}

	return StateCodes{
		States: states,
	}
}
