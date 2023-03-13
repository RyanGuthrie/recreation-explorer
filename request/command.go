package request

import (
	"encoding/json"
	"fmt"
	"net/http"
	"server/client/recreation_gov"
)

type Cmd[T any] struct {
	client         recreation_gov.Client
	RawReq         *http.Request
	RawResp        *http.Response
	ParsedResponse T
}

func NewCmd[T any](client recreation_gov.Client, req *http.Request) Cmd[T] {
	cmd := Cmd[T]{
		client: client,
		RawReq: req,
	}

	cmd.Initialize()

	return cmd
}

func (cmd *Cmd[any]) Initialize() {
	cmd.RawReq.Header.Set("accept", "application/json")
	cmd.RawReq.Header.Set("apikey", RecreationGovBearer)
}

var count = 0

func (cmd *Cmd[any]) Execute() error {
	rawResponse, err := cmd.client.Clnt.Do(cmd.RawReq)
	if err != nil {
		return fmt.Errorf("failed executing request [%v]: %s", cmd.RawReq, err)
	}

	cmd.RawResp = rawResponse

	if count > 0 {
		//toString := util.ReaderToString(rawResponse.Body)
		//fmt.Println(toString)
	}

	count++

	decoder := json.NewDecoder(rawResponse.Body)
	decoder.DisallowUnknownFields()

	err = decoder.Decode(&cmd.ParsedResponse)
	if err != nil {
		return fmt.Errorf("failed deserializing response: %s", err)
	}

	return nil
}
