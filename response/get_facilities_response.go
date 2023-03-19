package response

import (
	"fmt"
	"server/model"
	"strings"
)

type GetFacilitiesResponse struct {
	Facilities []model.Facility `json:"RECDATA"`
	MetaData   model.Metadata   `json:"METADATA"`
}

func (resp GetFacilitiesResponse) String() string {
	summary := make([]string, len(resp.Facilities))
	for i, a := range resp.Facilities {
		summary[i] = a.String()
	}

	return fmt.Sprintf("Metadata: %v\nFacilities: \n  -%v\n", resp.MetaData.String(), strings.Join(summary, "\n  -"))
}
