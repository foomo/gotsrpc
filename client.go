package gotsrpc

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"
	"github.com/ugorji/go/codec"
)

// ClientTransport to use for calls
// var ClientTransport = &http.Transport{}

var _ Client = &bufferedClient{}

type Client interface {
	Call(url string, endpoint string, method string, args []interface{}, reply []interface{}) (err error)
}

func NewClient(httpClient *http.Client) Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	return &bufferedClient{client: httpClient}
}

func newRequest(url string, contentType string, reader io.Reader) (r *http.Request, err error) {
	request, errRequest := http.NewRequest("POST", url, reader)
	if errRequest != nil {
		return nil, errors.Wrap(errRequest, "could not create a request")
	}
	request.Header.Set("Content-Type", contentType)
	request.Header.Set("Accept", contentType)
	return request, nil
}

type bufferedClient struct {
	client *http.Client
}

// CallClient calls a method on the remove service
func (c *bufferedClient) Call(url string, endpoint string, method string, args []interface{}, reply []interface{}) (err error) {
	// Marshall args

	b := new(bytes.Buffer)
	errEncode := codec.NewEncoder(b, msgpackHandle).Encode(args)
	if errEncode != nil {
		return errors.Wrap(errEncode, "could not encode argument")
	}

	// Create request
	// Create post url
	postURL := fmt.Sprintf("%s%s/%s", url, endpoint, method)
	// Post

	request, errRequest := newRequest(postURL, msgpackContentType, b)
	if errRequest != nil {
		return errRequest
	}

	resp, errDo := c.client.Do(request)
	if errDo != nil {
		return errors.Wrap(err, "could not execute request")
	}

	// Check status
	if resp.StatusCode != http.StatusOK {
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		return fmt.Errorf("%s: %s", resp.Status, string(body))
	}

	var errDecode error
	switch resp.Header.Get("Content-Type") {
	case msgpackContentType:
		errDecode = codec.NewDecoder(resp.Body, msgpackHandle).Decode(reply)
	case jsonContentType:
		errDecode = codec.NewDecoder(resp.Body, jsonHandle).Decode(reply)
	default:
		errDecode = codec.NewDecoder(resp.Body, jsonHandle).Decode(reply)
	}

	// Unmarshal reply
	if errDecode != nil {
		return errors.Wrap(errDecode, "could not decode response from client")
	}
	return err
}
