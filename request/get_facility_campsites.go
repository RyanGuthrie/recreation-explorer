package request

import (
	"fmt"
	"net/http"
	"server/client/recreation_gov"
	"server/response"
)

const facilityCampsitesURL = "/api/v1/facilities/%s/campsites"

type GetFacilityCampsitesCmd struct {
	Request        *http.Request
	Response       *http.Response
	ParsedResponse response.GetFacilityCampsitesResponse
}

func NewGetFacilityCampsitesRequest(client recreation_gov.Client, facilityID string) (*Cmd[response.GetFacilityCampsitesResponse], error) {
	queryParams := make(map[string]string)
	queryParams["state"] = "CO"
	queryParams["sort"] = "Name"
	//queryParams["activity"] = "CAMPING"
	queryParams["full"] = "true"

	URL := fmt.Sprintf(facilityCampsitesURL, facilityID)
	uri := NewRequestUrl(URL, queryParams)

	req, err := http.NewRequest(http.MethodGet, uri.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("failed creating request to %s", uri)
	}

	cmd := NewCmd[response.GetFacilityCampsitesResponse](client, req)

	return &cmd, nil
}
