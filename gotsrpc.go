package gotsrpc

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/pkg/errors"
)

func GetCalledFunc(r *http.Request, endPoint string) string {
	return strings.TrimPrefix(r.URL.Path, endPoint+"/")
}

func ErrorFuncNotFound(w http.ResponseWriter) {
	http.Error(w, "method not found", http.StatusNotFound)
}

func ErrorCouldNotReply(w http.ResponseWriter) {
	http.Error(w, "could not reply", http.StatusInternalServerError)
}

func ErrorCouldNotLoadArgs(w http.ResponseWriter) {
	http.Error(w, "could not load args", http.StatusBadRequest)
}

func ErrorMethodNotAllowed(w http.ResponseWriter) {
	http.Error(w, "you gotta POST", http.StatusMethodNotAllowed)
}

func LoadArgs(args interface{}, callStats *CallStats, r *http.Request) error {
	var start time.Time
	if callStats != nil {
		start = time.Now()
	}

	ch := getHandlerForContentType(r.Header.Get("Content-Type"))
	dec := ch.getDecoder(r.Body)
	errDecode := dec.Decode(args)
	ch.putDecoder(dec)

	if errDecode != nil {
		return errors.Wrap(errDecode, "could not decode arguments")
	}

	if callStats != nil {
		callStats.Unmarshalling = time.Since(start)
		callStats.RequestSize = int(r.ContentLength)
	}

	return nil
}

func loadArgs(args interface{}, jsonBytes []byte) error {
	if err := json.Unmarshal(jsonBytes, &args); err != nil {
		return err
	}

	return nil
}

// Reply although this is a public method - do not call it, it will be called by generated code
func Reply(response []interface{}, stats *CallStats, r *http.Request, w http.ResponseWriter) error {
	var errorIndices []int

	for i, v := range response {
		if er, ok := v.(*errorReply); ok {
			errorIndices = append(errorIndices, i)
			response[i] = er.err
		}
	}

	var serializationStart time.Time
	if stats != nil {
		serializationStart = time.Now()
	}

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

	if stats != nil {
		stats.ResponseSize = buf.Len()
		stats.Marshalling = time.Since(serializationStart)

		for _, i := range errorIndices {
			if v, ok := response[i].(error); ok && v != nil {
				if !reflect.ValueOf(v).IsZero() {
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
