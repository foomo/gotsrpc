package gotsrpc

import (
	"encoding/json"
	"net/http"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/foomo/gotsrpc/v3/semconv/httpconv"
	"github.com/pkg/errors"
)

// GetCalledFunc returns the method name from the request path, that is the
// path with endPoint and the following slash removed.
func GetCalledFunc(r *http.Request, endPoint string) string {
	return strings.TrimPrefix(r.URL.Path, endPoint+"/")
}

// ErrorFuncNotFound replies with 404 Not Found for an unknown method.
func ErrorFuncNotFound(w http.ResponseWriter) {
	http.Error(w, "method not found", http.StatusNotFound)
}

// ErrorCouldNotReply replies with 500 Internal Server Error when the response
// could not be written.
func ErrorCouldNotReply(w http.ResponseWriter) {
	http.Error(w, "could not reply", http.StatusInternalServerError)
}

// ErrorCouldNotLoadArgs replies with 400 Bad Request when the arguments could
// not be decoded.
func ErrorCouldNotLoadArgs(w http.ResponseWriter) {
	http.Error(w, "could not load args", http.StatusBadRequest)
}

// ErrorMethodNotAllowed replies with 405 Method Not Allowed for requests that
// are not POST.
func ErrorMethodNotAllowed(w http.ResponseWriter) {
	http.Error(w, "you gotta POST", http.StatusMethodNotAllowed)
}

// LoadArgs decodes the request body into args, which must be a pointer. The
// codec is selected by the Content-Type header and falls back to JSON. If
// call is not nil, the decode time and size are recorded on it. Generated
// proxies call LoadArgs; application code does not need to.
func LoadArgs(args any, call *httpconv.ServerCall, r *http.Request) error {
	start := time.Now()

	ch := getHandlerForContentType(r.Header.Get("Content-Type"))
	dec := ch.getDecoder(r.Body)
	errDecode := dec.Decode(args)
	ch.putDecoder(dec)

	if errDecode != nil {
		return errors.Wrap(errDecode, "could not decode arguments")
	}

	if call != nil {
		call.RecordUnmarshal(time.Since(start), int(r.ContentLength))
	}

	return nil
}

func loadArgs(args any, jsonBytes []byte) error {
	if err := json.Unmarshal(jsonBytes, &args); err != nil {
		return err
	}

	return nil
}

// Reply encodes response in the codec selected by the request Content-Type
// and writes it to w. Elements wrapped with [ErrorReply] are sent as [Error].
// If call is not nil, the encode time, size and non-nil errors are recorded
// on it.
//
// Reply is exported for generated proxies only. Do not call it directly.
func Reply(response []any, call *httpconv.ServerCall, r *http.Request, w http.ResponseWriter) error {
	var errorIndices []int

	for i, v := range response {
		if er, ok := v.(*errorReply); ok {
			errorIndices = append(errorIndices, i)
			response[i] = er.err
		}
	}

	serializationStart := time.Now()

	ch := getHandlerForContentType(r.Header.Get("Content-Type"))

	w.Header().Set("Content-Type", ch.contentType)

	if ch.beforeEncodeReply != nil {
		if err := ch.beforeEncodeReply(&response, errorIndices); err != nil {
			return errors.Wrap(err, "error during before encoder reply")
		}
	}

	buf := getBuffer()
	defer putBuffer(buf)

	enc := ch.getEncoder(buf)
	err := enc.Encode(response)
	ch.putEncoder(enc)

	if err != nil {
		return errors.Wrap(err, "could not encode data to accepted format")
	}

	w.Header().Set("Content-Length", strconv.Itoa(buf.Len()))

	if _, err := w.Write(buf.Bytes()); err != nil {
		return errors.Wrap(err, "could not write response")
	}

	if call != nil {
		call.RecordMarshal(time.Since(serializationStart), buf.Len())

		for _, i := range errorIndices {
			if v, ok := response[i].(error); ok && v != nil {
				if !reflect.ValueOf(v).IsZero() {
					code := 1
					if c, ok := v.(interface {
						ErrorCode() int
					}); ok {
						code = c.ErrorCode()
					}

					call.RecordError(v, code)
				}
			}
		}
	}

	return nil
}
