package service

import (
	"context"
)

type Service interface {
	Hello(ctx context.Context, args string) string
	TypedError(ctx context.Context, msg string) error
	CustomError(ctx context.Context, msg string) error
	WrappedError(ctx context.Context, msg string) error
	StandardError(ctx context.Context, msg string) error
}
