package service

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
)

type Handler struct{}

func (h *Handler) Hello(ctx context.Context, args string) string {
	fmt.Println(args)
	return "Hello " + args
}

func (h *Handler) TypedError(ctx context.Context, msg string) error {
	return ErrSomething
}

func (h *Handler) WrappedError(ctx context.Context, msg string) error {
	return errors.Wrap(ErrSomething, msg)
}

func (h *Handler) CustomError(ctx context.Context, msg string) error {
	return NewMyError(msg, ErrSomething)
}

func (h *Handler) StandardError(ctx context.Context, msg string) error {
	return fmt.Errorf("something went wrong: %s", msg)
}
