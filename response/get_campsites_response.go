package response

import (
	"fmt"
	"server/model"
	"strings"
)

type GetCampsitesResponse struct {
	RecData  []model.Campsite `json:"RECDATA"`
	MetaData model.Metadata   `json:"METADATA"`
}

func (resp GetCampsitesResponse) String() string {
	summary := make([]string, len(resp.RecData))
	for i, a := range resp.RecData {
		summary[i] = a.String()
	}

	return fmt.Sprintf("Metadata: %v\nCampsites: \n  -%v\n", resp.MetaData.String(), strings.Join(summary, "\n  -"))
}
