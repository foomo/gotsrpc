package service

import (
	"github.com/pkg/errors"
)

var ErrSomething = errors.New("something")

type MyError struct {
	Payload string `json:"payload"`
	err     error
}

func (e *MyError) Error() string {
	return e.Payload + ": " + e.err.Error()
}

func (e *MyError) Unwrap() error {
	return e.err
}

func NewMyError(msg string, err error) error {
	return &MyError{Payload: msg, err: err}
}
