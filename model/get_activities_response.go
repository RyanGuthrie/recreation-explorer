package model

import (
	"fmt"
	"strings"
)

type GetActivitiesResponse struct {
	RecData  []Activity `json:"RECDATA"`
	MetaData Metadata   `json:"METADATA"`
}

func (resp GetActivitiesResponse) String() string {
	summary := make([]string, len(resp.RecData))
	for i, a := range resp.RecData {
		summary[i] = a.String()
	}

	return fmt.Sprintf("Metadata: %v\nActivities: \n  -%v\n", resp.MetaData.String(), strings.Join(summary, "\n  -"))
}
