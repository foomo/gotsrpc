package gotsrpc

import (
	"compress/gzip"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"slices"
	"time"

	"github.com/golang/snappy"
	"github.com/pkg/errors"
	"github.com/ugorji/go/codec"
)

// Reply despite the fact, that this is a public method - do not call it, it will be called by generated code
func Reply(response []interface{}, stats *CallStats, r *http.Request, w http.ResponseWriter) error {
	responseWriter := newResponseWriterWithLength(w)
	defer recordStats(stats, response, responseWriter)()

	var responseBody io.Writer

	clientHandle := getHandlerForContentType(r.Header.Get("Content-Type"))
	responseWriter.Header().Set("Content-Type", clientHandle.contentType)
	// TODO: Add weighted compression support based on Accepted-Encoding header
	switch {
	case slices.Contains(r.Header.Values("Accept-Encoding"), "snappy"):
		responseWriter.Header().Set("Content-Encoding", "snappy")
		responseWriter.Header().Set("Vary", "Accept-Encoding")
		if snappyWriter, ok := globalCompressorWriterPools[CompressorSnappy].Get().(*snappy.Writer); ok {
			snappyWriter.Reset(responseWriter)

			defer globalCompressorWriterPools[CompressorSnappy].Put(snappyWriter)
			responseBody = snappyWriter
		}

	case slices.Contains(r.Header.Values("Accept-Encoding"), "gzip"):
		responseWriter.Header().Set("Content-Encoding", "gzip")
		responseWriter.Header().Set("Vary", "Accept-Encoding")

		if gzipWriter, ok := globalCompressorWriterPools[CompressorGZIP].Get().(*gzip.Writer); ok {
			gzipWriter.Reset(responseWriter)

			defer globalCompressorWriterPools[CompressorGZIP].Put(gzipWriter)
			responseBody = gzipWriter
		}
	default:
		responseBody = responseWriter
	}
	if clientHandle.beforeEncodeReply != nil {
		if err := clientHandle.beforeEncodeReply(&response); err != nil {
			return fmt.Errorf("error during before encoder reply: %w", err)
		}
	}

	if err := codec.NewEncoder(responseBody, clientHandle.handle).Encode(response); err != nil {
		return fmt.Errorf("could not encode data to accepted format: %w", err)
	}

	// We need to close the response body writer, otherwise the client will hang if it's compressed
	if writer, ok := responseBody.(io.Closer); ok {
		if err := writer.Close(); err != nil {
			return errors.Wrap(err, "failed to write to response body")
		}
	}
	return nil
}

func recordStats(stats *CallStats, response []interface{}, responseWriter *responseWriterWithLength) func() {
	if stats == nil {
		return func() {}
	}

	start := time.Now()
	return func() {
		stats.ResponseSize = responseWriter.length
		stats.Marshalling = time.Since(start)
		if len(response) == 0 {
			return
		}

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
