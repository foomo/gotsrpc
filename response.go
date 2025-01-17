package gotsrpc

import (
	"compress/gzip"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"slices"
	"sync"
	"time"

	"github.com/golang/snappy"
	"github.com/pkg/errors"
	"github.com/ugorji/go/codec"
)

var (
	responseCompressors = map[Compressor]*sync.Pool{
		CompressorGZIP:   {New: func() interface{} { return gzip.NewWriter(io.Discard) }},
		CompressorSnappy: {New: func() interface{} { return snappy.NewBufferedWriter(io.Discard) }},
	}
)

// Reply despite the fact, that this is a public method - do not call it, it will be called by generated code
func Reply(response []interface{}, stats *CallStats, r *http.Request, w http.ResponseWriter) error {

	responseWriter := newResponseWriterWithLength(w)
	serializationStart := time.Now()
	var responseBody io.Writer = responseWriter

	clientHandle := getHandlerForContentType(r.Header.Get("Content-Type"))
	responseWriter.Header().Set("Content-Type", clientHandle.contentType)
	// TODO: Add weighted compression support based on Accepted-Encoding header
	switch {
	case slices.Contains(r.Header.Values("Accept-Encoding"), "snappy"):
		responseWriter.Header().Set("Content-Encoding", "snappy")
		responseWriter.Header().Set("Vary", "Accept-Encoding")

		snappyWriter := responseCompressors[CompressorSnappy].Get().(*snappy.Writer)
		snappyWriter.Reset(responseWriter)

		defer responseCompressors[CompressorSnappy].Put(snappyWriter)
		responseBody = snappyWriter
	case slices.Contains(r.Header.Values("Accept-Encoding"), "gzip"):
		responseWriter.Header().Set("Content-Encoding", "gzip")
		responseWriter.Header().Set("Vary", "Accept-Encoding")

		gzipWriter := responseCompressors[CompressorGZIP].Get().(*gzip.Writer)
		gzipWriter.Reset(responseWriter)

		defer responseCompressors[CompressorGZIP].Put(gzipWriter)
		responseBody = gzipWriter
	}

	if clientHandle.beforeEncodeReply != nil {
		if err := clientHandle.beforeEncodeReply(&response); err != nil {
			return fmt.Errorf("error during before encoder reply: %w", err)
		}
	}

	if err := codec.NewEncoder(responseBody, clientHandle.handle).Encode(response); err != nil {
		return fmt.Errorf("could not encode data to accepted format: %w", err)
	}

	if writer, ok := responseBody.(io.Closer); ok {
		if err := writer.Close(); err != nil {
			return errors.Wrap(err, "failed to write to response body")
		}
	}

	if stats != nil {
		stats.ResponseSize = responseWriter.length
		stats.Marshalling = time.Since(serializationStart)
		if len(response) > 0 {
			errResp := response[len(response)-1]
			if v, ok := errResp.(error); ok && v != nil {
				if !reflect.ValueOf(v).IsNil() {
					stats.ErrorCode = 1
					stats.ErrorType = fmt.Sprintf("%T", v)
					stats.ErrorMessage = v.Error()
					if v, ok := v.(interface {
						ErrorCode() int
					}); ok {
						stats.ErrorCode = v.ErrorCode()
					}
				}
			}
		}
	}
	return nil
}
