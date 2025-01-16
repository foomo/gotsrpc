package gotsrpc

import (
	"compress/gzip"
	"context"
	"fmt"
	"io"
	"net/http"
	"sync"

	"github.com/golang/snappy"
	"github.com/pkg/errors"
	"github.com/ugorji/go/codec"
	"golang.org/x/sync/errgroup"
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

// WithClientHandle allows you to specify a custom clientHandle.
func WithClientHandle(h *clientHandle) ClientOption {
	return func(bc *BufferedClient) {
		bc.handle = h
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
		handle:     getHandleForType(EncodingJson),
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
	reader, writer := io.Pipe()
	defer reader.Close()
	// If no arguments are set, remove
	g, _ := errgroup.WithContext(ctx)

	if len(args) != 0 {
		g.Go(func() error {

			// Close piped writer after encoding
			defer writer.Close()

			var encodeWriter io.Writer
			switch c.compressor {
			case CompressorGZIP:
				gzipWriter := c.writerPoolMap[CompressorGZIP].Get().(*gzip.Writer)
				gzipWriter.Reset(writer)

				defer c.writerPoolMap[CompressorGZIP].Put(gzipWriter)

				encodeWriter = gzipWriter
				defer gzipWriter.Close()
			case CompressorSnappy:
				snappyWriter := c.writerPoolMap[CompressorSnappy].Get().(*snappy.Writer)
				snappyWriter.Reset(writer)

				defer c.writerPoolMap[CompressorSnappy].Put(snappyWriter)

				encodeWriter = snappyWriter
				defer snappyWriter.Close()
			case CompressorNone:
				encodeWriter = writer
			default:
				encodeWriter = writer
			}

			return codec.NewEncoder(encodeWriter, c.handle.handle).Encode(args)
		})
	} else {
		// Without arguments, skip the piping altogether
		writer.Close()
	}

	// Create post url
	postURL := fmt.Sprintf("%s%s/%s", url, endpoint, method)

	req, err := newRequest(ctx, postURL, c.handle.contentType, reader, c.headers.Clone())
	if err != nil {
		return NewClientError(errors.Wrap(err, "failed to create request"))
	}

	switch c.compressor {
	case CompressorGZIP:
		req.Header.Set("Content-Encoding", "gzip")
	case CompressorSnappy:
		req.Header.Set("Content-Encoding", "snappy")
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

	if len(args) != 0 {
		err = g.Wait()
		if err != nil {
			return NewClientError(errors.Wrap(err, "failed to send request data"))
		}
	}

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

	if err := codec.NewDecoder(resp.Body, clientHandle.handle).Decode(wrappedReply); err != nil {
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
