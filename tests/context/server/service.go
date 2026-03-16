package server

import (
	"context"
)

type Service interface {
	Hello(ctx context.Context, msg string) string
	Error(ctx context.Context, msg string) error
	PkgError(ctx context.Context, msg string) error
	JoinedError(ctx context.Context, msg string) error
	WrappedError(ctx context.Context, msg string) error
	CustomError(ctx context.Context, msg string) error
}
