package gotsrpc

import (
	"errors"
	"fmt"
)

// HTTPError reports a response with a status code other than 200 OK.
// [Client.Call] returns it wrapped in a [ClientError].
type HTTPError struct {
	// StatusCode is the HTTP status code of the response.
	StatusCode int
	// Body is reserved for the response body. [NewHTTPError] does not set it.
	Body string
	error
}

// NewHTTPError returns an [HTTPError] with status code and message msg.
func NewHTTPError(msg string, code int) *HTTPError {
	return &HTTPError{
		error:      errors.New(msg),
		StatusCode: code,
	}
}

// Error returns the message prefixed with the status code, for example
// "[404] not found".
func (e *HTTPError) Error() string {
	return fmt.Sprintf("[%d] %s", e.StatusCode, e.error.Error())
}
