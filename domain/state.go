package domain

import (
	"github.com/manifoldco/promptui"
	"log"
	"server/client/recreation_gov"
	"server/prompt"
	"server/request"
)

type State struct {
	// Full state name e.g. "Colorado"
	Name string

	// The states two character code e.g. "CO"
	Code string
}

var client = recreation_gov.NewHttpClient()

func GetState(cursorPos *int) (State, error) {
	var states []State
	for state, code := range StateToCode {
		states = append(states, State{Name: state, Code: code})
	}

	selectTemplates := promptui.SelectTemplates{
		Active:   "→ {{ .Name | red }}{{if .Code}} ({{ .Code | red }}){{end}} ←",
		Inactive: "{{ .Name | cyan }}",
		Selected: "→ {{ .Name | blue }}",
	}

	state, err := prompt.Select(
		"Select State",
		states, cursorPos,
		selectTemplates,
		func(item State) string {
			return item.Name
		}, func(l, r int) bool {
			return states[l].Code < states[r].Code
		}, State{Name: "Quit Program"})
	if err != nil {
		return State{}, err
	}

	return *state, nil
}

func (state State) Explore() error {
	var cursorPos = 0
	cmd, err := request.NewGetFacilitiesRequest(client, state.Code)
	if err != nil {
		log.Fatalln(err)
	}

	if err = cmd.Execute(); err != nil {
		log.Fatalf("Failed executing req [%v]: %s", cmd.RawReq.URL, err)
	}

	for {
		facility, err := GetFacility(state, cmd.ParsedResponse.Facilities, &cursorPos)
		if err == prompt.ExitError {
			return err
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

var StateToCode = map[string]string{
	"Alabama":        "AL",
	"Alaska":         "AK",
	"Arizona":        "AZ",
	"Arkansas":       "AR",
	"California":     "CA",
	"Colorado":       "CO",
	"Connecticut":    "CT",
	"Delaware":       "DE",
	"Florida":        "FL",
	"Georgia":        "GA",
	"Hawaii":         "HI",
	"Idaho":          "ID",
	"Illinois":       "IL",
	"Indiana":        "IN",
	"Iowa":           "IA",
	"Kansas":         "KS",
	"Kentucky":       "KY",
	"Louisiana":      "LA",
	"Maine":          "ME",
	"Maryland":       "MD",
	"Massachusetts":  "MA",
	"Michigan":       "MI",
	"Minnesota":      "MN",
	"Mississippi":    "MS",
	"Missouri":       "MO",
	"Montana":        "MT",
	"Nebraska":       "NE",
	"Nevada":         "NV",
	"New Hampshire":  "NH",
	"New Jersey":     "NJ",
	"New Mexico":     "NM",
	"New York":       "NY",
	"North Carolina": "NC",
	"North Dakota":   "ND",
	"Ohio":           "OH",
	"Oklahoma":       "OK",
	"Oregon":         "OR",
	"Pennsylvania":   "PA",
	"Rhode Island":   "RI",
	"South Carolina": "SC",
	"South Dakota":   "SD",
	"Tennessee":      "TN",
	"Texas":          "TX",
	"Utah":           "UT",
	"Vermont":        "VT",
	"Virginia":       "VA",
	"Washington":     "WA",
	"West Virginia":  "WV",
	"Wisconsin":      "WI",
	"Wyoming":        "WY",
	// Territories
	"American Samoa":                 "AS",
	"District of Columbia":           "DC",
	"Federated States of Micronesia": "FM",
	"Guam":                           "GU",
	"Marshall Islands":               "MH",
	"Northern Mariana Islands":       "MP",
	"Palau":                          "PW",
	"Puerto Rico":                    "PR",
	"Virgin Islands":                 "VI",
	// Armed Forces (AE includes Europe, Africa, Canada, and the Middle East)
	"Armed Forces Americas": "AA",
	"Armed Forces Europe":   "AE",
	"Armed Forces Pacific":  "AP",
}
