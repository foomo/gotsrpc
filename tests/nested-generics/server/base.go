package server

import (
	"context"
)

type Base[T any] interface {
	GetValue(ctx context.Context) T
}
