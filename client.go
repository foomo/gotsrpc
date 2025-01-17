package gotsrpc

import (
	"bytes"
	"compress/gzip"
	"context"
	"fmt"
	"io"
	"net/http"
	"sync"

	"github.com/golang/snappy"
	"github.com/pkg/errors"
	"github.com/ugorji/go/codec"
)

const (
	HeaderServiceToService = "X-Foomo-S2S"
)

type Compressor int

const (
	CompressorNone Compressor = iota
	CompressorGZIP
	CompressorSnappy
)

func (c Compressor) String() string {
	switch c {
	case CompressorNone:
		return "none"
	case CompressorGZIP:
		return "gzip"
	case CompressorSnappy:
		return "snappy"
	default:
		return "unknown"
	}
}

// ClientTransport to use for calls
// var ClientTransport = &http.Transport{}

var _ Client = &BufferedClient{}

type Client interface {
	Call(ctx context.Context, url string, endpoint string, method string, args []interface{}, reply []interface{}) (err error)
}

func NewClientWithHttpClient(client *http.Client) Client { //nolint:stylecheck
	return NewBufferedClient(WithHTTPClient(client))
}

func newRequest(ctx context.Context, url string, contentType string, reader io.Reader, headers http.Header) (r *http.Request, err error) {
	request, errRequest := http.NewRequestWithContext(ctx, http.MethodPost, url, reader)
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

type BufferedClient struct {
	client        *http.Client
	handle        *clientHandle
	headers       http.Header
	compressor    Compressor
	writerPoolMap map[Compressor]*sync.Pool
}

// ClientOption is a function that configures a BufferedClient.
type ClientOption func(*BufferedClient)

// WithHTTPClient allows you to specify a custom *http.Client.
func WithHTTPClient(c *http.Client) ClientOption {
	return func(bc *BufferedClient) {
		if c == nil {
			bc.client = defaultHttpFactory()
		} else {
			bc.client = c
		}
	}
}

func WithClientEncoding(encoding ClientEncoding) ClientOption {
	return func(bc *BufferedClient) {
		bc.handle = getHandleForType(encoding)
	}
}

// WithHeaders allows you to specify custom HTTP headers.
func WithHeaders(h http.Header) ClientOption {
	return func(bc *BufferedClient) {
		bc.headers = h
	}
}

func WithCompressor(compressor Compressor) ClientOption {
	return func(bc *BufferedClient) {
		bc.compressor = compressor
	}
}

// NewBufferedClient is the constructor that applies all functional options.
func NewBufferedClient(opts ...ClientOption) *BufferedClient {
	// Set reasonable defaults here
	bc := &BufferedClient{
		client:     defaultHttpFactory(),
		headers:    make(http.Header),
		handle:     getHandleForType(EncodingMsgpack),
		compressor: CompressorNone,
		writerPoolMap: map[Compressor]*sync.Pool{
			CompressorGZIP: {
				New: func() interface{} { return gzip.NewWriter(nil) },
			},
			CompressorSnappy: {
				New: func() interface{} { return snappy.NewBufferedWriter(nil) },
			},
		},
	}

	// Apply each option
	for _, opt := range opts {
		opt(bc)
	}
	return bc
}

// Call calls a method on the remove service
func (c *BufferedClient) Call(ctx context.Context, url string, endpoint string, method string, args []interface{}, reply []interface{}) error {
	// Marshall args
	buffer := &bytes.Buffer{}

	// If no arguments are set, remove

	var encodeWriter io.Writer
	switch c.compressor {
	case CompressorGZIP:
		gzipWriter := c.writerPoolMap[CompressorGZIP].Get().(*gzip.Writer)
		gzipWriter.Reset(buffer)

		defer c.writerPoolMap[CompressorGZIP].Put(gzipWriter)

		encodeWriter = gzipWriter
	case CompressorSnappy:
		snappyWriter := c.writerPoolMap[CompressorSnappy].Get().(*snappy.Writer)
		snappyWriter.Reset(buffer)

		defer c.writerPoolMap[CompressorSnappy].Put(snappyWriter)
		encodeWriter = snappyWriter
	case CompressorNone:
		encodeWriter = buffer
	default:
		encodeWriter = buffer
	}

	err := codec.NewEncoder(encodeWriter, c.handle.handle).Encode(args)
	if err != nil {
		return errors.Wrap(err, "could not encode data")
	}

	if writer, ok := encodeWriter.(io.Closer); ok {
		if err = writer.Close(); err != nil {
			return errors.Wrap(err, "failed to write to request body")
		}
	}

	// Create post url
	postURL := fmt.Sprintf("%s%s/%s", url, endpoint, method)
	req, err := newRequest(ctx, postURL, c.handle.contentType, buffer, c.headers.Clone())
	if err != nil {
		return NewClientError(errors.Wrap(err, "failed to create request"))
	}

	switch c.compressor {
	case CompressorGZIP:
		req.Header.Set("Content-Encoding", "gzip")
		req.Header.Set("Accept-Encoding", "gzip")
	case CompressorSnappy:
		req.Header.Set("Content-Encoding", "snappy")
		req.Header.Set("Accept-Encoding", "snappy")
	case CompressorNone:
		// uncompressed, nothing to do
	default:
		// uncompressed, nothing to do
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return NewClientError(errors.Wrap(err, "failed to send request"))
	}
	defer resp.Body.Close()

	// Check status
	if resp.StatusCode != http.StatusOK {
		var msg string
		if value, err := io.ReadAll(resp.Body); err != nil {
			msg = "failed to read response body: " + err.Error()
		} else {
			msg = string(value)
		}
		return NewClientError(NewHTTPError(msg, resp.StatusCode))
	}
	clientHandle := getHandlerForContentType(resp.Header.Get("Content-Type"))

	wrappedReply := reply
	if clientHandle.beforeDecodeReply != nil {
		if value, err := clientHandle.beforeDecodeReply(reply); err != nil {
			return NewClientError(errors.Wrap(err, "failed to call beforeDecodeReply hook"))
		} else {
			wrappedReply = value
		}
	}

	var responseBodyReader io.Reader

	switch resp.Header.Get("Content-Encoding") {
	case "snappy":
		responseBodyReader = snappy.NewReader(resp.Body)
	case "gzip":
		gzipReader, err := gzip.NewReader(resp.Body)
		if err != nil {
			return NewClientError(errors.Wrap(err, "could not create gzip reader"))
		}
		responseBodyReader = gzipReader
		defer gzipReader.Close()
	default:
		responseBodyReader = resp.Body
	}

	if err := codec.NewDecoder(responseBodyReader, clientHandle.handle).Decode(wrappedReply); err != nil {
		return NewClientError(errors.Wrap(err, "failed to decode response"))
	}

	// replace error
	if clientHandle.afterDecodeReply != nil {
		if err := clientHandle.afterDecodeReply(&reply, wrappedReply); err != nil {
			return NewClientError(errors.Wrap(err, "failed to call afterDecodeReply hook"))
		}
	}

	return nil
}
