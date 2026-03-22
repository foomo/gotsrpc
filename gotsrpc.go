package gotsrpc

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/pkg/errors"
	"github.com/ugorji/go/codec"
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
	start := time.Now()

	handle := getHandlerForContentType(r.Header.Get("Content-Type")).handle
	if errDecode := codec.NewDecoder(r.Body, handle).Decode(args); errDecode != nil {
		_, _ = fmt.Fprintln(os.Stderr, errDecode.Error())
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

// Reply despite the fact, that this is a public method - do not call it, it will be called by generated code
func Reply(response []interface{}, stats *CallStats, r *http.Request, w http.ResponseWriter) error {
	serializationStart := time.Now()

	clientHandle := getHandlerForContentType(r.Header.Get("Content-Type"))

	w.Header().Set("Content-Type", clientHandle.contentType)

	if clientHandle.beforeEncodeReply != nil {
		if err := clientHandle.beforeEncodeReply(&response); err != nil {
			_, _ = fmt.Fprintln(os.Stderr, err.Error())
			return errors.Wrap(err, "error during before encoder reply")
		}
	}

	buf := new(bytes.Buffer)
	if err := codec.NewEncoder(buf, clientHandle.handle).Encode(response); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err.Error())
		return errors.Wrap(err, "could not encode data to accepted format")
	}

	w.Header().Set("Content-Length", strconv.Itoa(buf.Len()))

	if _, err := w.Write(buf.Bytes()); err != nil {
		return errors.Wrap(err, "could not write response")
	}

	if stats != nil {
		stats.ResponseSize = buf.Len()
		stats.Marshalling = time.Since(serializationStart)
		if len(response) > 0 {
			errResp := response[len(response)-1]
			if v, ok := errResp.(error); ok && v != nil {
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
