package server

import (
	"context"
)

type Middle[T any] interface {
	Base[T]
	GetWrapped(ctx context.Context) Response[T]
}

type Keyed[K comparable, V any] interface {
	GetByKey(ctx context.Context, key K) V
}

type Service interface {
	Middle[Item]
	Keyed[string, int]
	GetName(ctx context.Context) string
}
