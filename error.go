package gotsrpc

import (
	"fmt"
	"io"
	"reflect"
	"strings"

	"github.com/mitchellh/mapstructure"
	"github.com/pkg/errors"
)

type Error struct {
	Msg      string      `json:"m"`
	Pkg      string      `json:"p"`
	Type     string      `json:"t"`
	Data     interface{} `json:"d,omitempty"`
	ErrCause *Error      `json:"c,omitempty"`
}

// NewError returns a new instance
func NewError(err error) *Error {
	// check if already transformed
	if v, ok := err.(*Error); ok { //nolint:errorlint
		return v
	}

	// skip *withStack error type
	if _, ok := err.(interface {
		StackTrace() errors.StackTrace
	}); ok && errors.Unwrap(err) != nil {
		err = errors.Unwrap(err)
	}

	// retrieve error details
	errType := reflect.TypeOf(err)
	errElem := errType
	if errType.Kind() == reflect.Ptr {
		errElem = errType.Elem()
	}

	inst := &Error{
		Msg:  err.Error(),
		Type: errType.String(),
		Pkg:  errElem.PkgPath(),
		Data: err,
	}

	// unwrap error
	if unwrappedErr := errors.Unwrap(err); unwrappedErr != nil {
		inst.ErrCause = NewError(unwrappedErr)
		inst.Msg = strings.TrimSuffix(inst.Msg, ": "+unwrappedErr.Error())
	}

	return inst
}

// As interface
func (e *Error) As(err interface{}) bool {
	if e == nil || err == nil {
		return false
	}
	if reflect.TypeOf(err).Elem().String() == e.Type {
		if decodeErr := mapstructure.Decode(e.Data, &err); decodeErr != nil {
			fmt.Printf("ERROR: failed to decode error data\n%+v", decodeErr)
			return false
		} else {
			return true
		}
	}
	return false
}

// Cause interface
func (e *Error) Cause() error {
	if e.ErrCause != nil {
		return e.ErrCause
	}
	return e
}

// Format interface
func (e *Error) Format(s fmt.State, verb rune) {
	switch verb {
	case 'v':
		if s.Flag('+') {
			_, _ = fmt.Fprintf(s, "%s.(%s)\n", e.Pkg, e.Type)
			if e.Data != nil {
				_, _ = fmt.Fprintf(s, "Data: %v\n", e.Data)
			}
		}
		fallthrough
	case 's', 'q':
		_, _ = io.WriteString(s, e.Error())
	}
}

// Unwrap interface
func (e *Error) Unwrap() error {
	if e != nil && e.ErrCause != nil {
		return e.ErrCause
	}
	return nil
}

// Is interface
func (e *Error) Is(err error) bool {
	if e == nil || err == nil {
		return false
	}

	errType := reflect.TypeOf(err)
	errElem := errType
	if errType.Kind() == reflect.Ptr {
		errElem = errType.Elem()
	}

	if e.Msg == err.Error() &&
		errType.String() == e.Type &&
		errElem.PkgPath() == e.Pkg {
		return true
	}

	return false
}

// Error interface
func (e *Error) Error() string {
	msg := e.Msg
	if e.ErrCause != nil {
		msg += ": " + e.ErrCause.Error()
	}
	return msg
}
