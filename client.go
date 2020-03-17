package gotsrpc

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"
	"github.com/ugorji/go/codec"
)

const (
	HeaderServiceToService = "X-Foomo-S2S"
)

// ClientTransport to use for calls
// var ClientTransport = &http.Transport{}

var _ Client = &bufferedClient{}

type Client interface {
	Call(url string, endpoint string, method string, args []interface{}, reply []interface{}) (err error)
	SetClientEncoding(encoding ClientEncoding)
	SetTransportHttpClient(client *http.Client)
	SetDefaultHeaders(headers http.Header)
}

func NewClient() Client {
	return &bufferedClient{client: defaultHttpFactory(), handle: getHandleForEncoding(EncodingMsgpack), headers: nil}
}

func NewClientWithHttpClient(client *http.Client) Client {
	if client != nil {
		return &bufferedClient{client: client, handle: getHandleForEncoding(EncodingMsgpack), headers: nil}
	} else {
		return &bufferedClient{client: defaultHttpFactory(), handle: getHandleForEncoding(EncodingMsgpack), headers: nil}
	}
}

func newRequest(url string, contentType string, buffer *bytes.Buffer, headers http.Header) (r *http.Request, err error) {
	if buffer == nil {
		buffer = &bytes.Buffer{}
	}
	request, errRequest := http.NewRequest("POST", url, buffer)
	if errRequest != nil {
		return nil, errors.Wrap(errRequest, "could not create a request")
	}
	if len(headers) > 0 {
		request.Header = headers
	}
	request.Header.Set("Content-Type", contentType)
	request.Header.Set("Accept", contentType)
	request.Header.Set(HeaderServiceToService, "true")

	return request, nil
}

type bufferedClient struct {
	client  *http.Client
	handle  *clientHandle
	headers http.Header
}

func (c *bufferedClient) SetDefaultHeaders(headers http.Header) {
	c.headers = headers
}

func (c *bufferedClient) SetClientEncoding(encoding ClientEncoding) {
	c.handle = getHandleForEncoding(encoding)
}

func (c *bufferedClient) SetTransportHttpClient(client *http.Client) {
	c.client = client
}

// CallClient calls a method on the remove service
func (c *bufferedClient) Call(url string, endpoint string, method string, args []interface{}, reply []interface{}) (err error) {
	// Marshall args
	b := new(bytes.Buffer)

	// If no arguments are set, remove
	if len(args) > 0 {
		if err := codec.NewEncoder(b, c.handle.handle).Encode(args); err != nil {
			return errors.Wrap(err, "could not encode argument")
		}
	}

	// Create request
	// Create post url
	postURL := fmt.Sprintf("%s%s/%s", url, endpoint, method)

	request, errRequest := newRequest(postURL, c.handle.contentType, b, c.headers.Clone())
	if errRequest != nil {
		return errRequest
	}

	resp, errDo := c.client.Do(request)
	if errDo != nil {
		return errors.Wrap(errDo, "could not execute request")
	}
	defer resp.Body.Close()

	// Check status
	if resp.StatusCode != http.StatusOK {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		return fmt.Errorf("%s: %s", resp.Status, string(body))
	}

	responseHandle := getHandlerForContentType(resp.Header.Get("Content-Type")).handle
	if err := codec.NewDecoder(resp.Body, responseHandle).Decode(reply); err != nil {
		return errors.Wrap(err, "could not decode response from client")
	}

	return err
}
