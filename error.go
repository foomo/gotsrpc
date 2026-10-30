package gotsrpc

import (
	"fmt"
	"io"
	"reflect"
	"strings"

	"github.com/mitchellh/mapstructure"
	"github.com/pkg/errors"
)

// Error is the wire representation of an error returned by a service method.
// It keeps the message, the Go type and package of the original error, the
// error value as Data, and the cause chain, so that [errors.Is] and
// [errors.As] keep working on the client side.
type Error struct {
	Msg       string   `json:"m"`
	Pkg       string   `json:"p"`
	Type      string   `json:"t"`
	Data      any      `json:"d,omitempty"`
	ErrCause  *Error   `json:"c,omitempty"`
	ErrCauses []*Error `json:"cs,omitempty"`
}

// NewError converts err into an [Error]. It returns nil for a nil or typed nil
// err and returns err unchanged if it already is an *Error. Wrapped causes
// become ErrCause, and errors joined with [errors.Join] become ErrCauses.
func NewError(err error) *Error {
	// guard against typed nil (e.g. (*MyError)(nil) assigned to error interface)
	if reflect.ValueOf(err).IsZero() {
		return nil
	}

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
	if errType.Kind() == reflect.Pointer {
		errElem = errType.Elem()
	}

	inst := &Error{
		Msg:  err.Error(),
		Type: errType.String(),
		Pkg:  errElem.PkgPath(),
		Data: err,
	}

	// check for joined errors (errors.Join)
	if joinedErr, ok := err.(interface{ Unwrap() []error }); ok {
		errs := joinedErr.Unwrap()
		if len(errs) > 0 {
			inst.ErrCauses = make([]*Error, len(errs))
			for i, e := range errs {
				inst.ErrCauses[i] = NewError(e)
			}

			return inst
		}
	}

	// unwrap single error
	if unwrappedErr := errors.Unwrap(err); unwrappedErr != nil {
		inst.ErrCause = NewError(unwrappedErr)
		inst.Msg = strings.TrimSuffix(inst.Msg, ": "+unwrappedErr.Error())
	}

	return inst
}

// As decodes Data into err if the type of *err matches the original error
// type. It reports whether the decode succeeded. err must be a non-nil
// pointer.
func (e *Error) As(err any) bool {
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

// Cause returns the direct cause, or nil if there is none. It never returns e,
// because pkg/errors.Cause would then loop forever.
func (e *Error) Cause() error {
	if e.ErrCause != nil {
		return e.ErrCause
	}

	return nil
}

// Format implements [fmt.Formatter]. The %+v verb prints the package, type
// and data before the message.
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

// Unwrap returns the single cause and all joined causes, or nil if there are
// none.
func (e *Error) Unwrap() []error {
	if e == nil {
		return nil
	}

	var errs []error
	if e.ErrCause != nil {
		errs = append(errs, e.ErrCause)
	}

	for _, c := range e.ErrCauses {
		if c != nil {
			errs = append(errs, c)
		}
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Is reports whether err has the same message, type and package as the
// original error.
func (e *Error) Is(err error) bool {
	if e == nil || err == nil {
		return false
	}

	errType := reflect.TypeOf(err)

	errElem := errType
	if errType.Kind() == reflect.Pointer {
		errElem = errType.Elem()
	}

	if e.Msg == err.Error() &&
		errType.String() == e.Type &&
		errElem.PkgPath() == e.Pkg {
		return true
	}

	return false
}

// Error returns the message followed by the messages of the cause chain.
func (e *Error) Error() string {
	msg := e.Msg
	if e.ErrCause != nil {
		msg += ": " + e.ErrCause.Error()
	}

	return msg
}
