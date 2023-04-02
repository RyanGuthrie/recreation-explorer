package controller

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"server/domain"
	"server/model"
)

type Facility struct {
	Name string
}

type Facilities struct {
	Facilities []Facility `json:"facilities"`
}

func facilitiesToJson(modelFacilities []model.Facility) Facilities {
	facilities := make([]Facility, len(modelFacilities))

	for i, facility := range modelFacilities {
		facilities[i] = Facility{
			Name: facility.FacilityName,
		}
	}

	return Facilities{
		Facilities: facilities,
	}
}

func FacilityIndex(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	stateName := params.ByName("state")
	log.Printf("Got state %v", stateName)

	state, err := domain.FindState(stateName)
	if err != nil {
		message := fmt.Sprintf("%v\n", err.Error())
		_, _ = writer.Write([]byte(message))
		http.NotFound(writer, request)
		return
	}

	cmd, err := domain.ExecuteGetFacilitiesCmd(state)
	if err != nil {
		writer.WriteHeader(http.StatusServiceUnavailable)
		_, _ = writer.Write([]byte(err.Error()))
		return
	}

	if err := json.NewEncoder(writer).Encode(facilitiesToJson(cmd.ParsedResponse.Facilities)); err != nil {
		writer.WriteHeader(http.StatusInternalServerError)

		if _, err := writer.Write([]byte(fmt.Sprintf("Failed to write states: %v", err))); err != nil {
			log.Printf("Failed writing response: %v\n", err)
		}
	}
}
