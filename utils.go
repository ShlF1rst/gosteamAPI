package steamAPI

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

type InvalidResponseError struct {
	error string
}

func (e InvalidResponseError) Error() string {
	return e.error
}

func NewInvalidResponseError(resp *http.Response) *InvalidResponseError {
	bytes, err := ioutil.ReadAll(resp.Body)
	msg := ""
	if err == nil {
		msg = string(bytes)
	}
	return &InvalidResponseError{
		error: fmt.Sprintf("invalid response error with status code %d and message %s", resp.StatusCode, msg),
	}
}

func (client *client) processAPIRequest(httpMethod string, inter string, method string, version int, params map[string]string) ([]byte, error) {
	var req *http.Request
	var resp *http.Response
	var err error
	var bodyBytes []byte

	url := fmt.Sprintf("http://api.steampowered.com/%s/%s/v%d/", inter, method, version)

	if req, err = http.NewRequest(httpMethod, url, nil); err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("key", client.APIKey)
	query.Set("format", "json")
	for key, value := range params {
		query.Set(key, value)
	}
	req.URL.RawQuery = query.Encode()
	if resp, err = client.httpClient.Do(req); err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, NewInvalidResponseError(resp)
	}

	if bodyBytes, err = io.ReadAll(resp.Body); err != nil {
		return nil, err
	}

	return bodyBytes, nil
}
