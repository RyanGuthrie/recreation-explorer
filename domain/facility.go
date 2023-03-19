package domain

import (
	"fmt"
	"github.com/manifoldco/promptui"
	"log"
	"server/model"
	"server/prompt"
	"strings"
)

func GetFacility(state State, facilities []model.Facility, cursorPos *int) (model.Facility, error) {
	selectTemplates := promptui.SelectTemplates{
		Active:   "→ {{ .FacilityName | red }} ←",
		Inactive: "{{ .FacilityName | cyan }}",
		Selected: "→ {{ .FacilityName | blue }}",
	}

	facility, err := prompt.Select(
		"Select Facility",
		facilities, cursorPos, selectTemplates, func(item model.Facility) string {
			return item.FacilityName
		}, func(l, r int) bool {
			return facilities[l].FacilityName < facilities[r].FacilityName
		}, model.Facility{FacilityName: fmt.Sprintf("Exit %s facilities", state.Name)})

	if err != nil {
		return model.Facility{}, err
	}

	return *facility, nil
}

func Explore(facility model.Facility) error {
	var cursorPos int = 1
	for {
		items := []string{
			"details",
			"address",
			"reservation",
			"activities",
			"events",
			"campsites",
		}

		selectTemplates := promptui.SelectTemplates{
			Active:   "→ {{ . | red }} ←",
			Inactive: "{{ . | cyan }}",
			Selected: "→ {{ . | blue }}",
		}

		choice, err := prompt.Select(
			"Select Option", items, &cursorPos, selectTemplates,
			func(item string) string {
				return item
			}, func(l, r int) bool {
				return items[l] < items[r]
			},
			fmt.Sprintf("Exit %s", facility.FacilityName),
		)

		if err != nil {
			return err
		}

		switch *choice {
		case "details":
			log.Printf("%s\n%s\n%s\n%s\n%s\n",
				fmt.Sprintf("Name: %s (%s) - (Enalbed: %v)", facility.FacilityName, facility.FacilityID, facility.Enabled),
				fmt.Sprintf("Description: %s", facility.FacilityDescription),
				fmt.Sprintf("Campsites: %d, Activities: %d, Events: %d, Tours: %d",
					len(facility.Campsites),
					len(facility.Activities),
					len(facility.Events),
					len(facility.Tours)),
				fmt.Sprintf("Type: %s, UseFee: %s", facility.FacilityTypeDescription, facility.FacilityUseFeeDescription),
				fmt.Sprintf("Keywords: %s", facility.Keywords),
			)
		case "address":
			var addresses []string
			for _, address := range facility.FacilityAddresses {
				addresses = append(addresses, address.String())
			}
			log.Printf("%s\n", strings.Join(addresses, "\n"))
		case "reservation":
			log.Printf("Reservable: %v\nReservation URL: %s", facility.Reservable, facility.FacilityReservationURL)
		case "activities":
			var activities []string
			for _, activity := range facility.Activities {
				activities = append(activities, activity.String())
			}

			log.Printf(strings.Join(activities, "\n"))
		case "events":
			var events []string
			for _, event := range facility.Events {
				events = append(events, event.String())
			}

			log.Printf(strings.Join(events, "\n"))
		case "campsites":
			campsite := Campsites{FacilityID: facility.FacilityID}

			err := campsite.Explore()
			if err != nil {
				return err
			}
		case "exit":
			return nil
		}
	}
}
