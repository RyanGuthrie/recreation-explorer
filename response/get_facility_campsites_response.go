package response

import (
	"fmt"
	"server/model"
	"strings"
)

type GetFacilityCampsitesResponse struct {
	Campsites []model.Campsite `json:"RECDATA"`
	MetaData  model.Metadata   `json:"METADATA"`
}

func (resp GetFacilityCampsitesResponse) String() string {
	summary := CampsiteSummary(resp.Campsites)

	return fmt.Sprintf("Metadata: %v\nCampsites: \n  -%v\n", resp.MetaData.String(), strings.Join(summary, "\n  -"))
}

func CampsiteSummary(campsites []model.Campsite) []string {
	summary := make([]string, len(campsites))
	for i, a := range campsites {
		summary[i] = a.String()
	}
	return summary
}
