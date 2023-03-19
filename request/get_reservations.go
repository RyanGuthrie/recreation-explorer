package request

import (
	"fmt"
	"net/http"
	"server/client/recreation_gov"
	"server/response"
)

const reservationsUrl = "/api/v1/reservations"

type GetReservationsCmd struct {
	Request        *http.Request
	Response       *http.Response
	ParsedResponse response.GetReservationsResponse
}

func NewGetReservationsRequest(client recreation_gov.Client) (*Cmd[response.GetReservationsResponse], error) {
	queryParams := make(map[string]string)
	queryParams["dataFrom"] = "2023-02-06"
	queryParams["dateTo"] = "2024-02-06"

	uri := NewRequestUrl(reservationsUrl, queryParams)

	req, err := http.NewRequest(http.MethodGet, uri.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("failed creating request to %s", uri)
	}

	cmd := NewCmd[response.GetReservationsResponse](client, req)

	return &cmd, nil
}
