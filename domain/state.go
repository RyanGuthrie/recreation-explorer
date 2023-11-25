package domain

import (
	"fmt"
	"github.com/manifoldco/promptui"
	"log"
	"server/client/recreation_gov"
	"server/prompt"
	"server/request"
	"server/response"
	"strings"
)

type State struct {
	// Full state name e.g. "Colorado"
	Name string

	// The states two character code e.g. "CO"
	Code string
}

var client = recreation_gov.NewHttpClient()

func GetState(cursorPos *int) (State, error) {
	selectTemplates := promptui.SelectTemplates{
		Active:   "→ {{ .Name | red }}{{if .Code}} ({{ .Code | red }}){{end}} ←",
		Inactive: "{{ .Name | cyan }}",
		Selected: "→ {{ .Name | blue }}",
	}

	state, err := prompt.Select(
		"Select State",
		States,
		cursorPos,
		selectTemplates,
		func(item State) string {
			return item.Name
		}, func(l, r int) bool {
			return States[l].Code < States[r].Code
		}, State{Name: "Quit Program"})
	if err != nil {
		return State{}, err
	}

	return *state, nil
}

func ExecuteGetFacilitiesCmd(state State) (*request.Cmd[response.GetFacilitiesResponse], error) {
	cmd, err := request.NewGetFacilitiesRequest(client, state.Code)
	if err != nil {
		log.Fatalln(err)
	}

	if err = cmd.Execute(); err != nil {
		return nil, fmt.Errorf("failed executing req [%v]: %s", cmd.RawReq.URL, err)
	}
	return cmd, nil
}

func (state State) Explore() error {
	var cursorPos = 0

	cmd, err := ExecuteGetFacilitiesCmd(state)
	if err != nil {
		return err
	}

	for {
		facility, err := GetFacility(state, cmd.ParsedResponse.Facilities, &cursorPos)
		if err == prompt.ExitError {
			return nil
		} else if err != nil {
			log.Println("Failed getting facility: ", err)
			continue
		}

		log.Printf("Facility: %s", facility.String())

		if err = Explore(facility); err != nil {
			return err
		}
	}
}

var States = []State{
	{Name: "Alabama", Code: "AL"},
	{Name: "Alaska", Code: "AK"},
	{Name: "Arizona", Code: "AZ"},
	{Name: "Arkansas", Code: "AR"},
	{Name: "California", Code: "CA"},
	{Name: "Colorado", Code: "CO"},
	{Name: "Connecticut", Code: "CT"},
	{Name: "Delaware", Code: "DE"},
	{Name: "Florida", Code: "FL"},
	{Name: "Georgia", Code: "GA"},
	{Name: "Hawaii", Code: "HI"},
	{Name: "Idaho", Code: "ID"},
	{Name: "Illinois", Code: "IL"},
	{Name: "Indiana", Code: "IN"},
	{Name: "Iowa", Code: "IA"},
	{Name: "Kansas", Code: "KS"},
	{Name: "Kentucky", Code: "KY"},
	{Name: "Louisiana", Code: "LA"},
	{Name: "Maine", Code: "ME"},
	{Name: "Maryland", Code: "MD"},
	{Name: "Massachusetts", Code: "MA"},
	{Name: "Michigan", Code: "MI"},
	{Name: "Minnesota", Code: "MN"},
	{Name: "Mississippi", Code: "MS"},
	{Name: "Missouri", Code: "MO"},
	{Name: "Montana", Code: "MT"},
	{Name: "Nebraska", Code: "NE"},
	{Name: "Nevada", Code: "NV"},
	{Name: "New Hampshire", Code: "NH"},
	{Name: "New Jersey", Code: "NJ"},
	{Name: "New Mexico", Code: "NM"},
	{Name: "New York", Code: "NY"},
	{Name: "North Carolina", Code: "NC"},
	{Name: "North Dakota", Code: "ND"},
	{Name: "Ohio", Code: "OH"},
	{Name: "Oklahoma", Code: "OK"},
	{Name: "Oregon", Code: "OR"},
	{Name: "Pennsylvania", Code: "PA"},
	{Name: "Rhode Island", Code: "RI"},
	{Name: "South Carolina", Code: "SC"},
	{Name: "South Dakota", Code: "SD"},
	{Name: "Tennessee", Code: "TN"},
	{Name: "Texas", Code: "TX"},
	{Name: "Utah", Code: "UT"},
	{Name: "Vermont", Code: "VT"},
	{Name: "Virginia", Code: "VA"},
	{Name: "Washington", Code: "WA"},
	{Name: "West Virginia", Code: "WV"},
	{Name: "Wisconsin", Code: "WI"},
	{Name: "Wyoming", Code: "WY"},
	{Name: "American Samoa", Code: "AS"},
	{Name: "District of Columbia", Code: "DC"},
	{Name: "Federated States of Micronesia", Code: "FM"},
	{Name: "Guam", Code: "GU"},
	{Name: "Marshall Islands", Code: "MH"},
	{Name: "Northern Mariana Islands", Code: "MP"},
	{Name: "Palau", Code: "PW"},
	{Name: "Puerto Rico", Code: "PR"},
	{Name: "Virgin Islands", Code: "VI"},
}

func FindState(name string) (State, error) {
	for i := range States {
		if strings.ToLower(States[i].Name) == strings.ToLower(name) {
			return States[i], nil
		}
	}

	return State{}, fmt.Errorf("invalid state: %v", name)
}
