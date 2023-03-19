package response

import (
	"fmt"
	"server/model"
	"strings"
)

type GetRecAreasResponse struct {
	RecData  []model.RecreationArea `json:"RECDATA"`
	MetaData model.Metadata         `json:"METADATA"`
}

func (resp GetRecAreasResponse) String() string {
	summary := make([]string, len(resp.RecData))
	for i, a := range resp.RecData {
		summary[i] = a.String()
	}

	return fmt.Sprintf("Metadata: %v\nRecAreas: \n  -%v\n", resp.MetaData.String(), strings.Join(summary, "\n  -"))
}
