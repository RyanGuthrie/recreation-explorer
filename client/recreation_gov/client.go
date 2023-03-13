package recreation_gov

import (
	"net/http"
)

type Client struct {
	Clnt *http.Client
}

func NewHttpClient() Client {
	return Client{
		Clnt: http.DefaultClient,
	}
}
