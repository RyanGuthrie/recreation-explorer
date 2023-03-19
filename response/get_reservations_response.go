package response

import (
	"fmt"
	"server/model"
	"strings"
)

type GetReservationsResponse struct {
	Reservations []model.Reservation `json:"data"`
	TotalResults int                 `json:"total_results"`
}

func (resp GetReservationsResponse) String() string {
	summary := make([]string, len(resp.Reservations))
	for i, a := range resp.Reservations {
		summary[i] = a.String()
	}

	return fmt.Sprintf("Reservations [%d]: %s\n", resp.TotalResults, strings.Join(summary, "\n  -"))
}
