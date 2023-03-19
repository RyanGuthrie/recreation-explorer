package response

import (
	"fmt"
	"server/model"
	"strings"
)

type GetActivitiesResponse struct {
	RecData  []model.Activity `json:"RECDATA"`
	MetaData model.Metadata   `json:"METADATA"`
}

func (resp GetActivitiesResponse) String() string {
	summary := make([]string, len(resp.RecData))
	for i, a := range resp.RecData {
		summary[i] = a.String()
	}

	return fmt.Sprintf("Metadata: %v\nActivities: \n  -%v\n", resp.MetaData.String(), strings.Join(summary, "\n  -"))
}
