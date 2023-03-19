package request

import (
	"fmt"
	"net/http"
	"server/client/recreation_gov"
	"server/response"
)

const campsitesURL = "/api/v1/campsites"

type GetCampsitesCmd struct {
	Request        *http.Request
	Response       *http.Response
	ParsedResponse response.GetCampsitesResponse
}

const (
	Overnight string = "Overnight"
	Day              = "Day"
)

func NewGetCampsitesCmd(client recreation_gov.Client, campsiteType string) (*Cmd[response.GetCampsitesResponse], error) {
	queryParams := make(map[string]string)
	queryParams["limit"] = "50"
	queryParams["full"] = "true"
	queryParams["query"] = campsiteType
	uri := NewRequestUrl(campsitesURL, queryParams)

	req, err := http.NewRequest(http.MethodGet, uri.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("failed creating request to %s", uri)
	}

	cmd := NewCmd[response.GetCampsitesResponse](client, req)

	return &cmd, nil
}
