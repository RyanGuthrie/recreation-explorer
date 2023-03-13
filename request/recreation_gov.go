package request

import (
	"log"
	"net/url"
	"os"
)

const baseUrl = "https://ridb.recreation.gov"

var RecreationGovBearer = os.Getenv("RECREATION_GOV_BEARER")
var recreationGovUrl url.URL

func init() {
	parsedUrl, err := url.Parse(baseUrl)
	recreationGovUrl = *parsedUrl

	if err != nil {
		log.Fatalf("failed creating URL from %v: %v", baseUrl, err)
	}
}

func NewRequestUrl(path string, queryParams map[string]string) url.URL {
	recUrl := recreationGovUrl

	recUrl.Path += path
	encodeQueryParameters(queryParams, &recUrl)

	return recUrl
}

func encodeQueryParameters(queryParams map[string]string, recUrl *url.URL) {
	if queryParams == nil {
		return
	}

	params := url.Values{}

	for k, v := range queryParams {
		params.Add(k, v)

	}

	recUrl.RawQuery = params.Encode()
}
