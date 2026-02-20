package server

import (
	"context"
	"errors"
	"fmt"

	pkgerrors "github.com/pkg/errors"
)

var (
	ErrGo  = pkgerrors.New("go")
	ErrPkg = pkgerrors.New("pkg error")
)

type Handler struct{}

func (h *Handler) Hello(ctx context.Context, args string) string {
	fmt.Println(args)
	return "Hello " + args
}

func (h *Handler) Error(ctx context.Context, args string) error {
	return ErrGo
}

func (h *Handler) JoinedError(ctx context.Context, args string) error {
	return errors.Join(ErrGo, ErrPkg)
}

func (h *Handler) PkgError(ctx context.Context, args string) error {
	return ErrPkg
}

func (h *Handler) WrappedError(ctx context.Context, args string) error {
	return pkgerrors.Wrap(ErrPkg, "wrapped error")
}

func (h *Handler) CustomError(ctx context.Context, msg string) error {
	return NewMyError(msg)
}
