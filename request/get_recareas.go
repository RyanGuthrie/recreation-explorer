package request

import (
	"fmt"
	"net/http"
	"server/client/recreation_gov"
	"server/response"
)

const recAreasUrl = "/api/v1/recareas"

type GetRecAreasCmd struct {
	Request        *http.Request
	Response       *http.Response
	ParsedResponse response.GetRecAreasResponse
}

func NewGetRecAreasRequest(client recreation_gov.Client) (*Cmd[response.GetRecAreasResponse], error) {
	queryParams := make(map[string]string)
	queryParams["state"] = "CO"
	queryParams["sort"] = "Name"
	queryParams["activity"] = "CAMPING"
	queryParams["full"] = "true"

	uri := NewRequestUrl(recAreasUrl, queryParams)

	req, err := http.NewRequest(http.MethodGet, uri.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("failed creating request to %s", uri)
	}

	cmd := NewCmd[response.GetRecAreasResponse](client, req)

	return &cmd, nil
}
