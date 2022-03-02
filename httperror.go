package gotsrpc

import (
	"errors"
	"fmt"
)

type HTTPError struct {
	StatusCode int
	Body       string
	error
}

func NewHTTPError(msg string, code int) *HTTPError {
	return &HTTPError{
		error:      errors.New(msg),
		StatusCode: code,
	}
}

func (e *HTTPError) Error() string {
	return fmt.Sprintf("[%d] %s", e.StatusCode, e.error.Error())
}
