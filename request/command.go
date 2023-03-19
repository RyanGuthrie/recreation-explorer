package request

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"io/fs"
	"net/http"
	"os"
	"server/client/recreation_gov"
	"strings"
)

const tempFilename = "/tmp/failedBody.json"

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

func (cmd *Cmd[any]) Execute() error {
	rawResponse, err := cmd.client.Clnt.Do(cmd.RawReq)
	if err != nil {
		return fmt.Errorf("failed executing request [%v]: %s", cmd.RawReq, err)
	}

	cmd.RawResp = rawResponse

	bodyBuf := new(strings.Builder)
	bodyReader := io.TeeReader(rawResponse.Body, bufio.NewWriter(bodyBuf))

	decoder := json.NewDecoder(bodyReader)
	decoder.DisallowUnknownFields()

	err = decoder.Decode(&cmd.ParsedResponse)

	// The body must be calculated here because we are Tee-ing the original reader
	body := bodyBuf.String()

	if rawResponse.StatusCode != http.StatusOK {
		return fmt.Errorf("received failed request [%v]: %v", rawResponse.StatusCode, body)
	}

	if err != nil {
		if err := cmd.saveTo(body, tempFilename); err != nil {
			return err
		}

		return fmt.Errorf("failed deserializing response (file://%s): %s", tempFilename, err)
	}

	return nil
}

func (cmd *Cmd[any]) saveTo(body, filename string) error {
	if err := os.WriteFile(filename, []byte(body), fs.ModePerm); err != nil {
		return fmt.Errorf("failed writing body to temp file [file://%s]: %s", filename, err)
	}

	return nil
}
