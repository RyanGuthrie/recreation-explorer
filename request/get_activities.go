package request

import (
	"fmt"
	"net/http"
	"server/client/recreation_gov"
	"server/response"
)

const activitiesUrl = "/api/v1/activities"

type GetActivitiesCmd struct {
	Request        *http.Request
	Response       *http.Response
	ParsedResponse response.GetActivitiesResponse
}

func NewGetActivitiesCmd(client recreation_gov.Client) (*Cmd[response.GetActivitiesResponse], error) {
	uri := NewRequestUrl(activitiesUrl, nil)

	req, err := http.NewRequest(http.MethodGet, uri.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("failed creating request to %s", uri)
	}

	cmd := NewCmd[response.GetActivitiesResponse](client, req)

	return &cmd, nil
}
