package request

import (
	"fmt"
	"net/http"
	"server/client/recreation_gov"
	"server/response"
)

const facilitiesUrl = "/api/v1/facilities"

type GetFacilitiesCmd struct {
	Request        *http.Request
	Response       *http.Response
	ParsedResponse response.GetFacilitiesResponse
}

func NewGetFacilitiesRequest(client recreation_gov.Client, stateCode string) (*Cmd[response.GetFacilitiesResponse], error) {
	queryParams := make(map[string]string)
	queryParams["limit"] = "1000"
	queryParams["state"] = stateCode
	queryParams["sort"] = "Name"
	queryParams["activity"] = "CAMPING"
	queryParams["full"] = "true"

	uri := NewRequestUrl(facilitiesUrl, queryParams)

	req, err := http.NewRequest(http.MethodGet, uri.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("failed creating request to %s", uri)
	}

	cmd := NewCmd[response.GetFacilitiesResponse](client, req)

	return &cmd, nil
}
