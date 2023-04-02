package domain

import (
	"fmt"
	"github.com/manifoldco/promptui"
	"log"
	"server/model"
	"server/prompt"
	"server/request"
)

type Campsites struct {
	FacilityID string
}

func (c Campsites) Explore() error {
	var cursorPos = 0
	for {
		cmd, err := request.NewGetFacilityCampsitesRequest(client, c.FacilityID)
		if err != nil {
			return fmt.Errorf("failed creating request: %v", err)
		}

		err = cmd.Execute()
		if err != nil {
			log.Printf("Failed executing request: %v\n", err)
			continue
		}

		campsitesResponse := cmd.ParsedResponse

		selectTemplates := promptui.SelectTemplates{
			Active:   "→ {{ .CampsiteName | red }} ←",
			Inactive: "{{ .CampsiteName | cyan }}",
			Selected: "→ {{ .CampsiteName | blue }}",
		}

		choice, err := prompt.Select("Select Campsite", campsitesResponse.Campsites, &cursorPos, selectTemplates,
			func(item model.Campsite) string {
				return item.CampsiteName
			}, func(l, r int) bool {
				return campsitesResponse.Campsites[l].CampsiteName < campsitesResponse.Campsites[r].CampsiteName
			},
			model.Campsite{CampsiteName: "Exit Campsites"},
		)

		if err == prompt.ExitError {
			return nil
		} else if err != nil {
			return err
		}

		campsite := Campsite{
			FacilityID: c.FacilityID,
			CampsiteID: choice.CampsiteID,
			Campsite:   *choice,
		}

		if err = campsite.Explore(); err != nil {
			return err
		}
	}
}
